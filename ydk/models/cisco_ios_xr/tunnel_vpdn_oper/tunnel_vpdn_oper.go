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

// Vpdn
// VPDN operational data
type Vpdn struct {
    parent types.Entity
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

func (vpdn *Vpdn) GetFilter() yfilter.YFilter { return vpdn.YFilter }

func (vpdn *Vpdn) SetFilter(yf yfilter.YFilter) { vpdn.YFilter = yf }

func (vpdn *Vpdn) GetGoName(yname string) string {
    if yname == "sessions" { return "Sessions" }
    if yname == "tunnel-destinations" { return "TunnelDestinations" }
    if yname == "vpdn-mirroring" { return "VpdnMirroring" }
    if yname == "vpdn-redundancy" { return "VpdnRedundancy" }
    if yname == "history-failures" { return "HistoryFailures" }
    return ""
}

func (vpdn *Vpdn) GetSegmentPath() string {
    return "Cisco-IOS-XR-tunnel-vpdn-oper:vpdn"
}

func (vpdn *Vpdn) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sessions" {
        return &vpdn.Sessions
    }
    if childYangName == "tunnel-destinations" {
        return &vpdn.TunnelDestinations
    }
    if childYangName == "vpdn-mirroring" {
        return &vpdn.VpdnMirroring
    }
    if childYangName == "vpdn-redundancy" {
        return &vpdn.VpdnRedundancy
    }
    if childYangName == "history-failures" {
        return &vpdn.HistoryFailures
    }
    return nil
}

func (vpdn *Vpdn) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sessions"] = &vpdn.Sessions
    children["tunnel-destinations"] = &vpdn.TunnelDestinations
    children["vpdn-mirroring"] = &vpdn.VpdnMirroring
    children["vpdn-redundancy"] = &vpdn.VpdnRedundancy
    children["history-failures"] = &vpdn.HistoryFailures
    return children
}

func (vpdn *Vpdn) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vpdn *Vpdn) GetBundleName() string { return "cisco_ios_xr" }

func (vpdn *Vpdn) GetYangName() string { return "vpdn" }

func (vpdn *Vpdn) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vpdn *Vpdn) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vpdn *Vpdn) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vpdn *Vpdn) SetParent(parent types.Entity) { vpdn.parent = parent }

func (vpdn *Vpdn) GetParent() types.Entity { return vpdn.parent }

func (vpdn *Vpdn) GetParentYangName() string { return "Cisco-IOS-XR-tunnel-vpdn-oper" }

// Vpdn_Sessions
// VPDN session list
type Vpdn_Sessions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VPDN session information. The type is slice of Vpdn_Sessions_Session.
    Session []Vpdn_Sessions_Session
}

func (sessions *Vpdn_Sessions) GetFilter() yfilter.YFilter { return sessions.YFilter }

func (sessions *Vpdn_Sessions) SetFilter(yf yfilter.YFilter) { sessions.YFilter = yf }

func (sessions *Vpdn_Sessions) GetGoName(yname string) string {
    if yname == "session" { return "Session" }
    return ""
}

func (sessions *Vpdn_Sessions) GetSegmentPath() string {
    return "sessions"
}

func (sessions *Vpdn_Sessions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "session" {
        for _, c := range sessions.Session {
            if sessions.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Vpdn_Sessions_Session{}
        sessions.Session = append(sessions.Session, child)
        return &sessions.Session[len(sessions.Session)-1]
    }
    return nil
}

func (sessions *Vpdn_Sessions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sessions.Session {
        children[sessions.Session[i].GetSegmentPath()] = &sessions.Session[i]
    }
    return children
}

func (sessions *Vpdn_Sessions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sessions *Vpdn_Sessions) GetBundleName() string { return "cisco_ios_xr" }

func (sessions *Vpdn_Sessions) GetYangName() string { return "sessions" }

func (sessions *Vpdn_Sessions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sessions *Vpdn_Sessions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sessions *Vpdn_Sessions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sessions *Vpdn_Sessions) SetParent(parent types.Entity) { sessions.parent = parent }

func (sessions *Vpdn_Sessions) GetParent() types.Entity { return sessions.parent }

func (sessions *Vpdn_Sessions) GetParentYangName() string { return "vpdn" }

// Vpdn_Sessions_Session
//  VPDN session information
type Vpdn_Sessions_Session struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Session label. The type is string with pattern:
    // [0-9a-fA-F]{1,8}.
    SessionLabel interface{}

    // Time to setup session in sec:msec. The type is interface{} with range:
    // 0..4294967295.
    SetupTime interface{}

    // Parent interface name. The type is string with pattern: [a-zA-Z0-9./-]+.
    ParentInterfaceName interface{}

    // Session data.
    Session Vpdn_Sessions_Session_Session

    // L2TP data.
    L2Tp Vpdn_Sessions_Session_L2Tp

    // Subscriber data.
    Subscriber Vpdn_Sessions_Session_Subscriber

    // Configuration data.
    Configuration Vpdn_Sessions_Session_Configuration
}

func (session *Vpdn_Sessions_Session) GetFilter() yfilter.YFilter { return session.YFilter }

func (session *Vpdn_Sessions_Session) SetFilter(yf yfilter.YFilter) { session.YFilter = yf }

func (session *Vpdn_Sessions_Session) GetGoName(yname string) string {
    if yname == "session-label" { return "SessionLabel" }
    if yname == "setup-time" { return "SetupTime" }
    if yname == "parent-interface-name" { return "ParentInterfaceName" }
    if yname == "session" { return "Session" }
    if yname == "l2tp" { return "L2Tp" }
    if yname == "subscriber" { return "Subscriber" }
    if yname == "configuration" { return "Configuration" }
    return ""
}

func (session *Vpdn_Sessions_Session) GetSegmentPath() string {
    return "session" + "[session-label='" + fmt.Sprintf("%v", session.SessionLabel) + "']"
}

func (session *Vpdn_Sessions_Session) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "session" {
        return &session.Session
    }
    if childYangName == "l2tp" {
        return &session.L2Tp
    }
    if childYangName == "subscriber" {
        return &session.Subscriber
    }
    if childYangName == "configuration" {
        return &session.Configuration
    }
    return nil
}

func (session *Vpdn_Sessions_Session) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["session"] = &session.Session
    children["l2tp"] = &session.L2Tp
    children["subscriber"] = &session.Subscriber
    children["configuration"] = &session.Configuration
    return children
}

func (session *Vpdn_Sessions_Session) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["session-label"] = session.SessionLabel
    leafs["setup-time"] = session.SetupTime
    leafs["parent-interface-name"] = session.ParentInterfaceName
    return leafs
}

func (session *Vpdn_Sessions_Session) GetBundleName() string { return "cisco_ios_xr" }

func (session *Vpdn_Sessions_Session) GetYangName() string { return "session" }

func (session *Vpdn_Sessions_Session) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (session *Vpdn_Sessions_Session) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (session *Vpdn_Sessions_Session) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (session *Vpdn_Sessions_Session) SetParent(parent types.Entity) { session.parent = parent }

func (session *Vpdn_Sessions_Session) GetParent() types.Entity { return session.parent }

func (session *Vpdn_Sessions_Session) GetParentYangName() string { return "sessions" }

// Vpdn_Sessions_Session_Session
// Session data
type Vpdn_Sessions_Session_Session struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Elapsed time since last change in hh:mm:ss format. The type is string.
    LastChange interface{}

    // Session interface name. The type is string with pattern: [a-zA-Z0-9./-]+.
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

func (session *Vpdn_Sessions_Session_Session) GetFilter() yfilter.YFilter { return session.YFilter }

func (session *Vpdn_Sessions_Session_Session) SetFilter(yf yfilter.YFilter) { session.YFilter = yf }

func (session *Vpdn_Sessions_Session_Session) GetGoName(yname string) string {
    if yname == "last-change" { return "LastChange" }
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "username" { return "Username" }
    if yname == "domain-name" { return "DomainName" }
    if yname == "state" { return "State" }
    if yname == "l2tp-session-id" { return "L2TpSessionId" }
    if yname == "l2tp-tunnel-id" { return "L2TpTunnelId" }
    if yname == "srg-slave" { return "SrgSlave" }
    return ""
}

func (session *Vpdn_Sessions_Session_Session) GetSegmentPath() string {
    return "session"
}

func (session *Vpdn_Sessions_Session_Session) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (session *Vpdn_Sessions_Session_Session) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (session *Vpdn_Sessions_Session_Session) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["last-change"] = session.LastChange
    leafs["interface-name"] = session.InterfaceName
    leafs["username"] = session.Username
    leafs["domain-name"] = session.DomainName
    leafs["state"] = session.State
    leafs["l2tp-session-id"] = session.L2TpSessionId
    leafs["l2tp-tunnel-id"] = session.L2TpTunnelId
    leafs["srg-slave"] = session.SrgSlave
    return leafs
}

func (session *Vpdn_Sessions_Session_Session) GetBundleName() string { return "cisco_ios_xr" }

func (session *Vpdn_Sessions_Session_Session) GetYangName() string { return "session" }

func (session *Vpdn_Sessions_Session_Session) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (session *Vpdn_Sessions_Session_Session) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (session *Vpdn_Sessions_Session_Session) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (session *Vpdn_Sessions_Session_Session) SetParent(parent types.Entity) { session.parent = parent }

func (session *Vpdn_Sessions_Session_Session) GetParent() types.Entity { return session.parent }

func (session *Vpdn_Sessions_Session_Session) GetParentYangName() string { return "session" }

// Vpdn_Sessions_Session_L2Tp
// L2TP data
type Vpdn_Sessions_Session_L2Tp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Local endpoint IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    LocalEndpoint interface{}

    // Remote endpoint IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
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

func (l2Tp *Vpdn_Sessions_Session_L2Tp) GetFilter() yfilter.YFilter { return l2Tp.YFilter }

func (l2Tp *Vpdn_Sessions_Session_L2Tp) SetFilter(yf yfilter.YFilter) { l2Tp.YFilter = yf }

func (l2Tp *Vpdn_Sessions_Session_L2Tp) GetGoName(yname string) string {
    if yname == "local-endpoint" { return "LocalEndpoint" }
    if yname == "remote-endpoint" { return "RemoteEndpoint" }
    if yname == "call-serial-number" { return "CallSerialNumber" }
    if yname == "is-l2tp-class-attribute-mask-set" { return "IsL2TpClassAttributeMaskSet" }
    if yname == "local-tunnel-id" { return "LocalTunnelId" }
    if yname == "remote-tunnel-id" { return "RemoteTunnelId" }
    if yname == "local-session-id" { return "LocalSessionId" }
    if yname == "remote-session-id" { return "RemoteSessionId" }
    if yname == "remote-port" { return "RemotePort" }
    if yname == "tunnel-client-authentication-id" { return "TunnelClientAuthenticationId" }
    if yname == "tunnel-server-authentication-id" { return "TunnelServerAuthenticationId" }
    if yname == "tunnel-assignment-id" { return "TunnelAssignmentId" }
    if yname == "is-tunnel-authentication-enabled" { return "IsTunnelAuthenticationEnabled" }
    return ""
}

func (l2Tp *Vpdn_Sessions_Session_L2Tp) GetSegmentPath() string {
    return "l2tp"
}

func (l2Tp *Vpdn_Sessions_Session_L2Tp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (l2Tp *Vpdn_Sessions_Session_L2Tp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (l2Tp *Vpdn_Sessions_Session_L2Tp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["local-endpoint"] = l2Tp.LocalEndpoint
    leafs["remote-endpoint"] = l2Tp.RemoteEndpoint
    leafs["call-serial-number"] = l2Tp.CallSerialNumber
    leafs["is-l2tp-class-attribute-mask-set"] = l2Tp.IsL2TpClassAttributeMaskSet
    leafs["local-tunnel-id"] = l2Tp.LocalTunnelId
    leafs["remote-tunnel-id"] = l2Tp.RemoteTunnelId
    leafs["local-session-id"] = l2Tp.LocalSessionId
    leafs["remote-session-id"] = l2Tp.RemoteSessionId
    leafs["remote-port"] = l2Tp.RemotePort
    leafs["tunnel-client-authentication-id"] = l2Tp.TunnelClientAuthenticationId
    leafs["tunnel-server-authentication-id"] = l2Tp.TunnelServerAuthenticationId
    leafs["tunnel-assignment-id"] = l2Tp.TunnelAssignmentId
    leafs["is-tunnel-authentication-enabled"] = l2Tp.IsTunnelAuthenticationEnabled
    return leafs
}

func (l2Tp *Vpdn_Sessions_Session_L2Tp) GetBundleName() string { return "cisco_ios_xr" }

func (l2Tp *Vpdn_Sessions_Session_L2Tp) GetYangName() string { return "l2tp" }

func (l2Tp *Vpdn_Sessions_Session_L2Tp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (l2Tp *Vpdn_Sessions_Session_L2Tp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (l2Tp *Vpdn_Sessions_Session_L2Tp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (l2Tp *Vpdn_Sessions_Session_L2Tp) SetParent(parent types.Entity) { l2Tp.parent = parent }

func (l2Tp *Vpdn_Sessions_Session_L2Tp) GetParent() types.Entity { return l2Tp.parent }

func (l2Tp *Vpdn_Sessions_Session_L2Tp) GetParentYangName() string { return "session" }

// Vpdn_Sessions_Session_Subscriber
// Subscriber data
type Vpdn_Sessions_Session_Subscriber struct {
    parent types.Entity
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

func (subscriber *Vpdn_Sessions_Session_Subscriber) GetFilter() yfilter.YFilter { return subscriber.YFilter }

func (subscriber *Vpdn_Sessions_Session_Subscriber) SetFilter(yf yfilter.YFilter) { subscriber.YFilter = yf }

func (subscriber *Vpdn_Sessions_Session_Subscriber) GetGoName(yname string) string {
    if yname == "nas-port-type" { return "NasPortType" }
    if yname == "physical-channel-id" { return "PhysicalChannelId" }
    if yname == "receive-connect-speed" { return "ReceiveConnectSpeed" }
    if yname == "transmit-connect-speed" { return "TransmitConnectSpeed" }
    if yname == "nas-port" { return "NasPort" }
    return ""
}

func (subscriber *Vpdn_Sessions_Session_Subscriber) GetSegmentPath() string {
    return "subscriber"
}

func (subscriber *Vpdn_Sessions_Session_Subscriber) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (subscriber *Vpdn_Sessions_Session_Subscriber) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (subscriber *Vpdn_Sessions_Session_Subscriber) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["nas-port-type"] = subscriber.NasPortType
    leafs["physical-channel-id"] = subscriber.PhysicalChannelId
    leafs["receive-connect-speed"] = subscriber.ReceiveConnectSpeed
    leafs["transmit-connect-speed"] = subscriber.TransmitConnectSpeed
    leafs["nas-port"] = subscriber.NasPort
    return leafs
}

func (subscriber *Vpdn_Sessions_Session_Subscriber) GetBundleName() string { return "cisco_ios_xr" }

func (subscriber *Vpdn_Sessions_Session_Subscriber) GetYangName() string { return "subscriber" }

func (subscriber *Vpdn_Sessions_Session_Subscriber) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (subscriber *Vpdn_Sessions_Session_Subscriber) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (subscriber *Vpdn_Sessions_Session_Subscriber) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (subscriber *Vpdn_Sessions_Session_Subscriber) SetParent(parent types.Entity) { subscriber.parent = parent }

func (subscriber *Vpdn_Sessions_Session_Subscriber) GetParent() types.Entity { return subscriber.parent }

func (subscriber *Vpdn_Sessions_Session_Subscriber) GetParentYangName() string { return "session" }

// Vpdn_Sessions_Session_Configuration
// Configuration data
type Vpdn_Sessions_Session_Configuration struct {
    parent types.Entity
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

func (configuration *Vpdn_Sessions_Session_Configuration) GetFilter() yfilter.YFilter { return configuration.YFilter }

func (configuration *Vpdn_Sessions_Session_Configuration) SetFilter(yf yfilter.YFilter) { configuration.YFilter = yf }

func (configuration *Vpdn_Sessions_Session_Configuration) GetGoName(yname string) string {
    if yname == "template-name" { return "TemplateName" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "l2tp-busy-timeout" { return "L2TpBusyTimeout" }
    if yname == "tos-mode" { return "TosMode" }
    if yname == "tos" { return "Tos" }
    if yname == "dsl-line-forwarding" { return "DslLineForwarding" }
    if yname == "vpn-id" { return "VpnId" }
    return ""
}

func (configuration *Vpdn_Sessions_Session_Configuration) GetSegmentPath() string {
    return "configuration"
}

func (configuration *Vpdn_Sessions_Session_Configuration) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vpn-id" {
        return &configuration.VpnId
    }
    return nil
}

func (configuration *Vpdn_Sessions_Session_Configuration) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["vpn-id"] = &configuration.VpnId
    return children
}

func (configuration *Vpdn_Sessions_Session_Configuration) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["template-name"] = configuration.TemplateName
    leafs["vrf-name"] = configuration.VrfName
    leafs["l2tp-busy-timeout"] = configuration.L2TpBusyTimeout
    leafs["tos-mode"] = configuration.TosMode
    leafs["tos"] = configuration.Tos
    leafs["dsl-line-forwarding"] = configuration.DslLineForwarding
    return leafs
}

func (configuration *Vpdn_Sessions_Session_Configuration) GetBundleName() string { return "cisco_ios_xr" }

func (configuration *Vpdn_Sessions_Session_Configuration) GetYangName() string { return "configuration" }

func (configuration *Vpdn_Sessions_Session_Configuration) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (configuration *Vpdn_Sessions_Session_Configuration) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (configuration *Vpdn_Sessions_Session_Configuration) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (configuration *Vpdn_Sessions_Session_Configuration) SetParent(parent types.Entity) { configuration.parent = parent }

func (configuration *Vpdn_Sessions_Session_Configuration) GetParent() types.Entity { return configuration.parent }

func (configuration *Vpdn_Sessions_Session_Configuration) GetParentYangName() string { return "session" }

// Vpdn_Sessions_Session_Configuration_VpnId
// VPN ID
type Vpdn_Sessions_Session_Configuration_VpnId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OUI. The type is interface{} with range: 0..4294967295.
    Oui interface{}

    // Index. The type is interface{} with range: 0..4294967295.
    Index interface{}
}

func (vpnId *Vpdn_Sessions_Session_Configuration_VpnId) GetFilter() yfilter.YFilter { return vpnId.YFilter }

func (vpnId *Vpdn_Sessions_Session_Configuration_VpnId) SetFilter(yf yfilter.YFilter) { vpnId.YFilter = yf }

func (vpnId *Vpdn_Sessions_Session_Configuration_VpnId) GetGoName(yname string) string {
    if yname == "oui" { return "Oui" }
    if yname == "index" { return "Index" }
    return ""
}

func (vpnId *Vpdn_Sessions_Session_Configuration_VpnId) GetSegmentPath() string {
    return "vpn-id"
}

func (vpnId *Vpdn_Sessions_Session_Configuration_VpnId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (vpnId *Vpdn_Sessions_Session_Configuration_VpnId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (vpnId *Vpdn_Sessions_Session_Configuration_VpnId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["oui"] = vpnId.Oui
    leafs["index"] = vpnId.Index
    return leafs
}

func (vpnId *Vpdn_Sessions_Session_Configuration_VpnId) GetBundleName() string { return "cisco_ios_xr" }

func (vpnId *Vpdn_Sessions_Session_Configuration_VpnId) GetYangName() string { return "vpn-id" }

func (vpnId *Vpdn_Sessions_Session_Configuration_VpnId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vpnId *Vpdn_Sessions_Session_Configuration_VpnId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vpnId *Vpdn_Sessions_Session_Configuration_VpnId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vpnId *Vpdn_Sessions_Session_Configuration_VpnId) SetParent(parent types.Entity) { vpnId.parent = parent }

func (vpnId *Vpdn_Sessions_Session_Configuration_VpnId) GetParent() types.Entity { return vpnId.parent }

func (vpnId *Vpdn_Sessions_Session_Configuration_VpnId) GetParentYangName() string { return "configuration" }

// Vpdn_TunnelDestinations
// VPDN tunnel Destinations
type Vpdn_TunnelDestinations struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VPDN tunnel destination information. The type is slice of
    // Vpdn_TunnelDestinations_TunnelDestination.
    TunnelDestination []Vpdn_TunnelDestinations_TunnelDestination
}

func (tunnelDestinations *Vpdn_TunnelDestinations) GetFilter() yfilter.YFilter { return tunnelDestinations.YFilter }

func (tunnelDestinations *Vpdn_TunnelDestinations) SetFilter(yf yfilter.YFilter) { tunnelDestinations.YFilter = yf }

func (tunnelDestinations *Vpdn_TunnelDestinations) GetGoName(yname string) string {
    if yname == "tunnel-destination" { return "TunnelDestination" }
    return ""
}

func (tunnelDestinations *Vpdn_TunnelDestinations) GetSegmentPath() string {
    return "tunnel-destinations"
}

func (tunnelDestinations *Vpdn_TunnelDestinations) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tunnel-destination" {
        for _, c := range tunnelDestinations.TunnelDestination {
            if tunnelDestinations.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Vpdn_TunnelDestinations_TunnelDestination{}
        tunnelDestinations.TunnelDestination = append(tunnelDestinations.TunnelDestination, child)
        return &tunnelDestinations.TunnelDestination[len(tunnelDestinations.TunnelDestination)-1]
    }
    return nil
}

func (tunnelDestinations *Vpdn_TunnelDestinations) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range tunnelDestinations.TunnelDestination {
        children[tunnelDestinations.TunnelDestination[i].GetSegmentPath()] = &tunnelDestinations.TunnelDestination[i]
    }
    return children
}

func (tunnelDestinations *Vpdn_TunnelDestinations) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (tunnelDestinations *Vpdn_TunnelDestinations) GetBundleName() string { return "cisco_ios_xr" }

func (tunnelDestinations *Vpdn_TunnelDestinations) GetYangName() string { return "tunnel-destinations" }

func (tunnelDestinations *Vpdn_TunnelDestinations) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tunnelDestinations *Vpdn_TunnelDestinations) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tunnelDestinations *Vpdn_TunnelDestinations) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tunnelDestinations *Vpdn_TunnelDestinations) SetParent(parent types.Entity) { tunnelDestinations.parent = parent }

func (tunnelDestinations *Vpdn_TunnelDestinations) GetParent() types.Entity { return tunnelDestinations.parent }

func (tunnelDestinations *Vpdn_TunnelDestinations) GetParentYangName() string { return "vpdn" }

// Vpdn_TunnelDestinations_TunnelDestination
// VPDN tunnel destination information
type Vpdn_TunnelDestinations_TunnelDestination struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VRF name. The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
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

func (tunnelDestination *Vpdn_TunnelDestinations_TunnelDestination) GetFilter() yfilter.YFilter { return tunnelDestination.YFilter }

func (tunnelDestination *Vpdn_TunnelDestinations_TunnelDestination) SetFilter(yf yfilter.YFilter) { tunnelDestination.YFilter = yf }

func (tunnelDestination *Vpdn_TunnelDestinations_TunnelDestination) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "address" { return "Address" }
    if yname == "vrf-name-xr" { return "VrfNameXr" }
    if yname == "load" { return "Load" }
    if yname == "status" { return "Status" }
    if yname == "connects" { return "Connects" }
    if yname == "disconnects" { return "Disconnects" }
    if yname == "retry" { return "Retry" }
    if yname == "status-change-time" { return "StatusChangeTime" }
    return ""
}

func (tunnelDestination *Vpdn_TunnelDestinations_TunnelDestination) GetSegmentPath() string {
    return "tunnel-destination"
}

func (tunnelDestination *Vpdn_TunnelDestinations_TunnelDestination) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (tunnelDestination *Vpdn_TunnelDestinations_TunnelDestination) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (tunnelDestination *Vpdn_TunnelDestinations_TunnelDestination) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = tunnelDestination.VrfName
    leafs["address"] = tunnelDestination.Address
    leafs["vrf-name-xr"] = tunnelDestination.VrfNameXr
    leafs["load"] = tunnelDestination.Load
    leafs["status"] = tunnelDestination.Status
    leafs["connects"] = tunnelDestination.Connects
    leafs["disconnects"] = tunnelDestination.Disconnects
    leafs["retry"] = tunnelDestination.Retry
    leafs["status-change-time"] = tunnelDestination.StatusChangeTime
    return leafs
}

func (tunnelDestination *Vpdn_TunnelDestinations_TunnelDestination) GetBundleName() string { return "cisco_ios_xr" }

func (tunnelDestination *Vpdn_TunnelDestinations_TunnelDestination) GetYangName() string { return "tunnel-destination" }

func (tunnelDestination *Vpdn_TunnelDestinations_TunnelDestination) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tunnelDestination *Vpdn_TunnelDestinations_TunnelDestination) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tunnelDestination *Vpdn_TunnelDestinations_TunnelDestination) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tunnelDestination *Vpdn_TunnelDestinations_TunnelDestination) SetParent(parent types.Entity) { tunnelDestination.parent = parent }

func (tunnelDestination *Vpdn_TunnelDestinations_TunnelDestination) GetParent() types.Entity { return tunnelDestination.parent }

func (tunnelDestination *Vpdn_TunnelDestinations_TunnelDestination) GetParentYangName() string { return "tunnel-destinations" }

// Vpdn_VpdnMirroring
// VPDN Mirroring Statistics
type Vpdn_VpdnMirroring struct {
    parent types.Entity
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

func (vpdnMirroring *Vpdn_VpdnMirroring) GetFilter() yfilter.YFilter { return vpdnMirroring.YFilter }

func (vpdnMirroring *Vpdn_VpdnMirroring) SetFilter(yf yfilter.YFilter) { vpdnMirroring.YFilter = yf }

func (vpdnMirroring *Vpdn_VpdnMirroring) GetGoName(yname string) string {
    if yname == "sync-not-conn-cnt" { return "SyncNotConnCnt" }
    if yname == "sso-err-cnt" { return "SsoErrCnt" }
    if yname == "sso-batch-err-cnt" { return "SsoBatchErrCnt" }
    if yname == "alloc-err-cnt" { return "AllocErrCnt" }
    if yname == "alloc-cnt" { return "AllocCnt" }
    if yname == "qad-send-stats" { return "QadSendStats" }
    if yname == "qad-recv-stats" { return "QadRecvStats" }
    if yname == "qad-send-stats-last-clear" { return "QadSendStatsLastClear" }
    if yname == "qad-recv-stats-last-clear" { return "QadRecvStatsLastClear" }
    return ""
}

func (vpdnMirroring *Vpdn_VpdnMirroring) GetSegmentPath() string {
    return "vpdn-mirroring"
}

func (vpdnMirroring *Vpdn_VpdnMirroring) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "qad-send-stats" {
        return &vpdnMirroring.QadSendStats
    }
    if childYangName == "qad-recv-stats" {
        return &vpdnMirroring.QadRecvStats
    }
    if childYangName == "qad-send-stats-last-clear" {
        return &vpdnMirroring.QadSendStatsLastClear
    }
    if childYangName == "qad-recv-stats-last-clear" {
        return &vpdnMirroring.QadRecvStatsLastClear
    }
    return nil
}

func (vpdnMirroring *Vpdn_VpdnMirroring) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["qad-send-stats"] = &vpdnMirroring.QadSendStats
    children["qad-recv-stats"] = &vpdnMirroring.QadRecvStats
    children["qad-send-stats-last-clear"] = &vpdnMirroring.QadSendStatsLastClear
    children["qad-recv-stats-last-clear"] = &vpdnMirroring.QadRecvStatsLastClear
    return children
}

func (vpdnMirroring *Vpdn_VpdnMirroring) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sync-not-conn-cnt"] = vpdnMirroring.SyncNotConnCnt
    leafs["sso-err-cnt"] = vpdnMirroring.SsoErrCnt
    leafs["sso-batch-err-cnt"] = vpdnMirroring.SsoBatchErrCnt
    leafs["alloc-err-cnt"] = vpdnMirroring.AllocErrCnt
    leafs["alloc-cnt"] = vpdnMirroring.AllocCnt
    return leafs
}

func (vpdnMirroring *Vpdn_VpdnMirroring) GetBundleName() string { return "cisco_ios_xr" }

func (vpdnMirroring *Vpdn_VpdnMirroring) GetYangName() string { return "vpdn-mirroring" }

func (vpdnMirroring *Vpdn_VpdnMirroring) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vpdnMirroring *Vpdn_VpdnMirroring) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vpdnMirroring *Vpdn_VpdnMirroring) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vpdnMirroring *Vpdn_VpdnMirroring) SetParent(parent types.Entity) { vpdnMirroring.parent = parent }

func (vpdnMirroring *Vpdn_VpdnMirroring) GetParent() types.Entity { return vpdnMirroring.parent }

func (vpdnMirroring *Vpdn_VpdnMirroring) GetParentYangName() string { return "vpdn" }

// Vpdn_VpdnMirroring_QadSendStats
// qad send stats
type Vpdn_VpdnMirroring_QadSendStats struct {
    parent types.Entity
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

func (qadSendStats *Vpdn_VpdnMirroring_QadSendStats) GetFilter() yfilter.YFilter { return qadSendStats.YFilter }

func (qadSendStats *Vpdn_VpdnMirroring_QadSendStats) SetFilter(yf yfilter.YFilter) { qadSendStats.YFilter = yf }

func (qadSendStats *Vpdn_VpdnMirroring_QadSendStats) GetGoName(yname string) string {
    if yname == "msgs-sent" { return "MsgsSent" }
    if yname == "acks-sent" { return "AcksSent" }
    if yname == "no-partner" { return "NoPartner" }
    if yname == "sends-failed" { return "SendsFailed" }
    if yname == "acks-failed" { return "AcksFailed" }
    if yname == "pending-acks" { return "PendingAcks" }
    if yname == "timeouts" { return "Timeouts" }
    if yname == "suspends" { return "Suspends" }
    if yname == "resumes" { return "Resumes" }
    if yname == "sends-fragment" { return "SendsFragment" }
    if yname == "qad-last-seq-number" { return "QadLastSeqNumber" }
    if yname == "qad-frag-count" { return "QadFragCount" }
    if yname == "qad-ack-count" { return "QadAckCount" }
    if yname == "qad-unknown-acks" { return "QadUnknownAcks" }
    if yname == "qad-timeouts" { return "QadTimeouts" }
    if yname == "qad-rx-count" { return "QadRxCount" }
    if yname == "qad-rx-list-count" { return "QadRxListCount" }
    if yname == "qad-rx-list-q-size" { return "QadRxListQSize" }
    if yname == "qad-rx-first-seq-number" { return "QadRxFirstSeqNumber" }
    return ""
}

func (qadSendStats *Vpdn_VpdnMirroring_QadSendStats) GetSegmentPath() string {
    return "qad-send-stats"
}

func (qadSendStats *Vpdn_VpdnMirroring_QadSendStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (qadSendStats *Vpdn_VpdnMirroring_QadSendStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (qadSendStats *Vpdn_VpdnMirroring_QadSendStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["msgs-sent"] = qadSendStats.MsgsSent
    leafs["acks-sent"] = qadSendStats.AcksSent
    leafs["no-partner"] = qadSendStats.NoPartner
    leafs["sends-failed"] = qadSendStats.SendsFailed
    leafs["acks-failed"] = qadSendStats.AcksFailed
    leafs["pending-acks"] = qadSendStats.PendingAcks
    leafs["timeouts"] = qadSendStats.Timeouts
    leafs["suspends"] = qadSendStats.Suspends
    leafs["resumes"] = qadSendStats.Resumes
    leafs["sends-fragment"] = qadSendStats.SendsFragment
    leafs["qad-last-seq-number"] = qadSendStats.QadLastSeqNumber
    leafs["qad-frag-count"] = qadSendStats.QadFragCount
    leafs["qad-ack-count"] = qadSendStats.QadAckCount
    leafs["qad-unknown-acks"] = qadSendStats.QadUnknownAcks
    leafs["qad-timeouts"] = qadSendStats.QadTimeouts
    leafs["qad-rx-count"] = qadSendStats.QadRxCount
    leafs["qad-rx-list-count"] = qadSendStats.QadRxListCount
    leafs["qad-rx-list-q-size"] = qadSendStats.QadRxListQSize
    leafs["qad-rx-first-seq-number"] = qadSendStats.QadRxFirstSeqNumber
    return leafs
}

func (qadSendStats *Vpdn_VpdnMirroring_QadSendStats) GetBundleName() string { return "cisco_ios_xr" }

func (qadSendStats *Vpdn_VpdnMirroring_QadSendStats) GetYangName() string { return "qad-send-stats" }

func (qadSendStats *Vpdn_VpdnMirroring_QadSendStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (qadSendStats *Vpdn_VpdnMirroring_QadSendStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (qadSendStats *Vpdn_VpdnMirroring_QadSendStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (qadSendStats *Vpdn_VpdnMirroring_QadSendStats) SetParent(parent types.Entity) { qadSendStats.parent = parent }

func (qadSendStats *Vpdn_VpdnMirroring_QadSendStats) GetParent() types.Entity { return qadSendStats.parent }

func (qadSendStats *Vpdn_VpdnMirroring_QadSendStats) GetParentYangName() string { return "vpdn-mirroring" }

// Vpdn_VpdnMirroring_QadRecvStats
// qad recv stats
type Vpdn_VpdnMirroring_QadRecvStats struct {
    parent types.Entity
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

func (qadRecvStats *Vpdn_VpdnMirroring_QadRecvStats) GetFilter() yfilter.YFilter { return qadRecvStats.YFilter }

func (qadRecvStats *Vpdn_VpdnMirroring_QadRecvStats) SetFilter(yf yfilter.YFilter) { qadRecvStats.YFilter = yf }

func (qadRecvStats *Vpdn_VpdnMirroring_QadRecvStats) GetGoName(yname string) string {
    if yname == "msgs-recvd" { return "MsgsRecvd" }
    if yname == "acks-recvd" { return "AcksRecvd" }
    if yname == "recvd-acks-failed" { return "RecvdAcksFailed" }
    if yname == "init-drops" { return "InitDrops" }
    if yname == "msg-drops" { return "MsgDrops" }
    if yname == "ooo-drops" { return "OooDrops" }
    if yname == "stale-msgs" { return "StaleMsgs" }
    return ""
}

func (qadRecvStats *Vpdn_VpdnMirroring_QadRecvStats) GetSegmentPath() string {
    return "qad-recv-stats"
}

func (qadRecvStats *Vpdn_VpdnMirroring_QadRecvStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (qadRecvStats *Vpdn_VpdnMirroring_QadRecvStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (qadRecvStats *Vpdn_VpdnMirroring_QadRecvStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["msgs-recvd"] = qadRecvStats.MsgsRecvd
    leafs["acks-recvd"] = qadRecvStats.AcksRecvd
    leafs["recvd-acks-failed"] = qadRecvStats.RecvdAcksFailed
    leafs["init-drops"] = qadRecvStats.InitDrops
    leafs["msg-drops"] = qadRecvStats.MsgDrops
    leafs["ooo-drops"] = qadRecvStats.OooDrops
    leafs["stale-msgs"] = qadRecvStats.StaleMsgs
    return leafs
}

func (qadRecvStats *Vpdn_VpdnMirroring_QadRecvStats) GetBundleName() string { return "cisco_ios_xr" }

func (qadRecvStats *Vpdn_VpdnMirroring_QadRecvStats) GetYangName() string { return "qad-recv-stats" }

func (qadRecvStats *Vpdn_VpdnMirroring_QadRecvStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (qadRecvStats *Vpdn_VpdnMirroring_QadRecvStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (qadRecvStats *Vpdn_VpdnMirroring_QadRecvStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (qadRecvStats *Vpdn_VpdnMirroring_QadRecvStats) SetParent(parent types.Entity) { qadRecvStats.parent = parent }

func (qadRecvStats *Vpdn_VpdnMirroring_QadRecvStats) GetParent() types.Entity { return qadRecvStats.parent }

func (qadRecvStats *Vpdn_VpdnMirroring_QadRecvStats) GetParentYangName() string { return "vpdn-mirroring" }

// Vpdn_VpdnMirroring_QadSendStatsLastClear
// qad send stats last clear
type Vpdn_VpdnMirroring_QadSendStatsLastClear struct {
    parent types.Entity
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

func (qadSendStatsLastClear *Vpdn_VpdnMirroring_QadSendStatsLastClear) GetFilter() yfilter.YFilter { return qadSendStatsLastClear.YFilter }

func (qadSendStatsLastClear *Vpdn_VpdnMirroring_QadSendStatsLastClear) SetFilter(yf yfilter.YFilter) { qadSendStatsLastClear.YFilter = yf }

func (qadSendStatsLastClear *Vpdn_VpdnMirroring_QadSendStatsLastClear) GetGoName(yname string) string {
    if yname == "msgs-sent" { return "MsgsSent" }
    if yname == "acks-sent" { return "AcksSent" }
    if yname == "no-partner" { return "NoPartner" }
    if yname == "sends-failed" { return "SendsFailed" }
    if yname == "acks-failed" { return "AcksFailed" }
    if yname == "pending-acks" { return "PendingAcks" }
    if yname == "timeouts" { return "Timeouts" }
    if yname == "suspends" { return "Suspends" }
    if yname == "resumes" { return "Resumes" }
    if yname == "sends-fragment" { return "SendsFragment" }
    if yname == "qad-last-seq-number" { return "QadLastSeqNumber" }
    if yname == "qad-frag-count" { return "QadFragCount" }
    if yname == "qad-ack-count" { return "QadAckCount" }
    if yname == "qad-unknown-acks" { return "QadUnknownAcks" }
    if yname == "qad-timeouts" { return "QadTimeouts" }
    if yname == "qad-rx-count" { return "QadRxCount" }
    if yname == "qad-rx-list-count" { return "QadRxListCount" }
    if yname == "qad-rx-list-q-size" { return "QadRxListQSize" }
    if yname == "qad-rx-first-seq-number" { return "QadRxFirstSeqNumber" }
    return ""
}

func (qadSendStatsLastClear *Vpdn_VpdnMirroring_QadSendStatsLastClear) GetSegmentPath() string {
    return "qad-send-stats-last-clear"
}

func (qadSendStatsLastClear *Vpdn_VpdnMirroring_QadSendStatsLastClear) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (qadSendStatsLastClear *Vpdn_VpdnMirroring_QadSendStatsLastClear) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (qadSendStatsLastClear *Vpdn_VpdnMirroring_QadSendStatsLastClear) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["msgs-sent"] = qadSendStatsLastClear.MsgsSent
    leafs["acks-sent"] = qadSendStatsLastClear.AcksSent
    leafs["no-partner"] = qadSendStatsLastClear.NoPartner
    leafs["sends-failed"] = qadSendStatsLastClear.SendsFailed
    leafs["acks-failed"] = qadSendStatsLastClear.AcksFailed
    leafs["pending-acks"] = qadSendStatsLastClear.PendingAcks
    leafs["timeouts"] = qadSendStatsLastClear.Timeouts
    leafs["suspends"] = qadSendStatsLastClear.Suspends
    leafs["resumes"] = qadSendStatsLastClear.Resumes
    leafs["sends-fragment"] = qadSendStatsLastClear.SendsFragment
    leafs["qad-last-seq-number"] = qadSendStatsLastClear.QadLastSeqNumber
    leafs["qad-frag-count"] = qadSendStatsLastClear.QadFragCount
    leafs["qad-ack-count"] = qadSendStatsLastClear.QadAckCount
    leafs["qad-unknown-acks"] = qadSendStatsLastClear.QadUnknownAcks
    leafs["qad-timeouts"] = qadSendStatsLastClear.QadTimeouts
    leafs["qad-rx-count"] = qadSendStatsLastClear.QadRxCount
    leafs["qad-rx-list-count"] = qadSendStatsLastClear.QadRxListCount
    leafs["qad-rx-list-q-size"] = qadSendStatsLastClear.QadRxListQSize
    leafs["qad-rx-first-seq-number"] = qadSendStatsLastClear.QadRxFirstSeqNumber
    return leafs
}

func (qadSendStatsLastClear *Vpdn_VpdnMirroring_QadSendStatsLastClear) GetBundleName() string { return "cisco_ios_xr" }

func (qadSendStatsLastClear *Vpdn_VpdnMirroring_QadSendStatsLastClear) GetYangName() string { return "qad-send-stats-last-clear" }

func (qadSendStatsLastClear *Vpdn_VpdnMirroring_QadSendStatsLastClear) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (qadSendStatsLastClear *Vpdn_VpdnMirroring_QadSendStatsLastClear) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (qadSendStatsLastClear *Vpdn_VpdnMirroring_QadSendStatsLastClear) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (qadSendStatsLastClear *Vpdn_VpdnMirroring_QadSendStatsLastClear) SetParent(parent types.Entity) { qadSendStatsLastClear.parent = parent }

func (qadSendStatsLastClear *Vpdn_VpdnMirroring_QadSendStatsLastClear) GetParent() types.Entity { return qadSendStatsLastClear.parent }

func (qadSendStatsLastClear *Vpdn_VpdnMirroring_QadSendStatsLastClear) GetParentYangName() string { return "vpdn-mirroring" }

// Vpdn_VpdnMirroring_QadRecvStatsLastClear
// qad recv stats last clear
type Vpdn_VpdnMirroring_QadRecvStatsLastClear struct {
    parent types.Entity
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

func (qadRecvStatsLastClear *Vpdn_VpdnMirroring_QadRecvStatsLastClear) GetFilter() yfilter.YFilter { return qadRecvStatsLastClear.YFilter }

func (qadRecvStatsLastClear *Vpdn_VpdnMirroring_QadRecvStatsLastClear) SetFilter(yf yfilter.YFilter) { qadRecvStatsLastClear.YFilter = yf }

func (qadRecvStatsLastClear *Vpdn_VpdnMirroring_QadRecvStatsLastClear) GetGoName(yname string) string {
    if yname == "msgs-recvd" { return "MsgsRecvd" }
    if yname == "acks-recvd" { return "AcksRecvd" }
    if yname == "recvd-acks-failed" { return "RecvdAcksFailed" }
    if yname == "init-drops" { return "InitDrops" }
    if yname == "msg-drops" { return "MsgDrops" }
    if yname == "ooo-drops" { return "OooDrops" }
    if yname == "stale-msgs" { return "StaleMsgs" }
    return ""
}

func (qadRecvStatsLastClear *Vpdn_VpdnMirroring_QadRecvStatsLastClear) GetSegmentPath() string {
    return "qad-recv-stats-last-clear"
}

func (qadRecvStatsLastClear *Vpdn_VpdnMirroring_QadRecvStatsLastClear) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (qadRecvStatsLastClear *Vpdn_VpdnMirroring_QadRecvStatsLastClear) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (qadRecvStatsLastClear *Vpdn_VpdnMirroring_QadRecvStatsLastClear) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["msgs-recvd"] = qadRecvStatsLastClear.MsgsRecvd
    leafs["acks-recvd"] = qadRecvStatsLastClear.AcksRecvd
    leafs["recvd-acks-failed"] = qadRecvStatsLastClear.RecvdAcksFailed
    leafs["init-drops"] = qadRecvStatsLastClear.InitDrops
    leafs["msg-drops"] = qadRecvStatsLastClear.MsgDrops
    leafs["ooo-drops"] = qadRecvStatsLastClear.OooDrops
    leafs["stale-msgs"] = qadRecvStatsLastClear.StaleMsgs
    return leafs
}

func (qadRecvStatsLastClear *Vpdn_VpdnMirroring_QadRecvStatsLastClear) GetBundleName() string { return "cisco_ios_xr" }

func (qadRecvStatsLastClear *Vpdn_VpdnMirroring_QadRecvStatsLastClear) GetYangName() string { return "qad-recv-stats-last-clear" }

func (qadRecvStatsLastClear *Vpdn_VpdnMirroring_QadRecvStatsLastClear) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (qadRecvStatsLastClear *Vpdn_VpdnMirroring_QadRecvStatsLastClear) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (qadRecvStatsLastClear *Vpdn_VpdnMirroring_QadRecvStatsLastClear) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (qadRecvStatsLastClear *Vpdn_VpdnMirroring_QadRecvStatsLastClear) SetParent(parent types.Entity) { qadRecvStatsLastClear.parent = parent }

func (qadRecvStatsLastClear *Vpdn_VpdnMirroring_QadRecvStatsLastClear) GetParent() types.Entity { return qadRecvStatsLastClear.parent }

func (qadRecvStatsLastClear *Vpdn_VpdnMirroring_QadRecvStatsLastClear) GetParentYangName() string { return "vpdn-mirroring" }

// Vpdn_VpdnRedundancy
// Show VPDN Redundancy information
type Vpdn_VpdnRedundancy struct {
    parent types.Entity
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

func (vpdnRedundancy *Vpdn_VpdnRedundancy) GetFilter() yfilter.YFilter { return vpdnRedundancy.YFilter }

func (vpdnRedundancy *Vpdn_VpdnRedundancy) SetFilter(yf yfilter.YFilter) { vpdnRedundancy.YFilter = yf }

func (vpdnRedundancy *Vpdn_VpdnRedundancy) GetGoName(yname string) string {
    if yname == "session-total" { return "SessionTotal" }
    if yname == "session-synced" { return "SessionSynced" }
    if yname == "state" { return "State" }
    if yname == "start-time" { return "StartTime" }
    if yname == "finish-time" { return "FinishTime" }
    if yname == "abort-time" { return "AbortTime" }
    return ""
}

func (vpdnRedundancy *Vpdn_VpdnRedundancy) GetSegmentPath() string {
    return "vpdn-redundancy"
}

func (vpdnRedundancy *Vpdn_VpdnRedundancy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (vpdnRedundancy *Vpdn_VpdnRedundancy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (vpdnRedundancy *Vpdn_VpdnRedundancy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["session-total"] = vpdnRedundancy.SessionTotal
    leafs["session-synced"] = vpdnRedundancy.SessionSynced
    leafs["state"] = vpdnRedundancy.State
    leafs["start-time"] = vpdnRedundancy.StartTime
    leafs["finish-time"] = vpdnRedundancy.FinishTime
    leafs["abort-time"] = vpdnRedundancy.AbortTime
    return leafs
}

func (vpdnRedundancy *Vpdn_VpdnRedundancy) GetBundleName() string { return "cisco_ios_xr" }

func (vpdnRedundancy *Vpdn_VpdnRedundancy) GetYangName() string { return "vpdn-redundancy" }

func (vpdnRedundancy *Vpdn_VpdnRedundancy) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vpdnRedundancy *Vpdn_VpdnRedundancy) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vpdnRedundancy *Vpdn_VpdnRedundancy) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vpdnRedundancy *Vpdn_VpdnRedundancy) SetParent(parent types.Entity) { vpdnRedundancy.parent = parent }

func (vpdnRedundancy *Vpdn_VpdnRedundancy) GetParent() types.Entity { return vpdnRedundancy.parent }

func (vpdnRedundancy *Vpdn_VpdnRedundancy) GetParentYangName() string { return "vpdn" }

// Vpdn_HistoryFailures
// VPDN history failure list
type Vpdn_HistoryFailures struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VPDN history failure information. The type is slice of
    // Vpdn_HistoryFailures_HistoryFailure.
    HistoryFailure []Vpdn_HistoryFailures_HistoryFailure
}

func (historyFailures *Vpdn_HistoryFailures) GetFilter() yfilter.YFilter { return historyFailures.YFilter }

func (historyFailures *Vpdn_HistoryFailures) SetFilter(yf yfilter.YFilter) { historyFailures.YFilter = yf }

func (historyFailures *Vpdn_HistoryFailures) GetGoName(yname string) string {
    if yname == "history-failure" { return "HistoryFailure" }
    return ""
}

func (historyFailures *Vpdn_HistoryFailures) GetSegmentPath() string {
    return "history-failures"
}

func (historyFailures *Vpdn_HistoryFailures) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "history-failure" {
        for _, c := range historyFailures.HistoryFailure {
            if historyFailures.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Vpdn_HistoryFailures_HistoryFailure{}
        historyFailures.HistoryFailure = append(historyFailures.HistoryFailure, child)
        return &historyFailures.HistoryFailure[len(historyFailures.HistoryFailure)-1]
    }
    return nil
}

func (historyFailures *Vpdn_HistoryFailures) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range historyFailures.HistoryFailure {
        children[historyFailures.HistoryFailure[i].GetSegmentPath()] = &historyFailures.HistoryFailure[i]
    }
    return children
}

func (historyFailures *Vpdn_HistoryFailures) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (historyFailures *Vpdn_HistoryFailures) GetBundleName() string { return "cisco_ios_xr" }

func (historyFailures *Vpdn_HistoryFailures) GetYangName() string { return "history-failures" }

func (historyFailures *Vpdn_HistoryFailures) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (historyFailures *Vpdn_HistoryFailures) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (historyFailures *Vpdn_HistoryFailures) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (historyFailures *Vpdn_HistoryFailures) SetParent(parent types.Entity) { historyFailures.parent = parent }

func (historyFailures *Vpdn_HistoryFailures) GetParent() types.Entity { return historyFailures.parent }

func (historyFailures *Vpdn_HistoryFailures) GetParentYangName() string { return "vpdn" }

// Vpdn_HistoryFailures_HistoryFailure
// VPDN history failure information
type Vpdn_HistoryFailures_HistoryFailure struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Username. The type is string.
    Username interface{}

    // Remote name. The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
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
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    DestinationAddress interface{}

    // Remote client ID. The type is interface{} with range: 0..65535.
    RemoteClientId interface{}

    // Home gateway. The type is string.
    HomeGateway interface{}

    // Source IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
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

func (historyFailure *Vpdn_HistoryFailures_HistoryFailure) GetFilter() yfilter.YFilter { return historyFailure.YFilter }

func (historyFailure *Vpdn_HistoryFailures_HistoryFailure) SetFilter(yf yfilter.YFilter) { historyFailure.YFilter = yf }

func (historyFailure *Vpdn_HistoryFailures_HistoryFailure) GetGoName(yname string) string {
    if yname == "username" { return "Username" }
    if yname == "remote-name" { return "RemoteName" }
    if yname == "username-xr" { return "UsernameXr" }
    if yname == "domain-name" { return "DomainName" }
    if yname == "mid" { return "Mid" }
    if yname == "nas" { return "Nas" }
    if yname == "destination-address" { return "DestinationAddress" }
    if yname == "remote-client-id" { return "RemoteClientId" }
    if yname == "home-gateway" { return "HomeGateway" }
    if yname == "source-address" { return "SourceAddress" }
    if yname == "local-client-id" { return "LocalClientId" }
    if yname == "event-time" { return "EventTime" }
    if yname == "error-repeat-count" { return "ErrorRepeatCount" }
    if yname == "failure-type" { return "FailureType" }
    return ""
}

func (historyFailure *Vpdn_HistoryFailures_HistoryFailure) GetSegmentPath() string {
    return "history-failure"
}

func (historyFailure *Vpdn_HistoryFailures_HistoryFailure) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (historyFailure *Vpdn_HistoryFailures_HistoryFailure) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (historyFailure *Vpdn_HistoryFailures_HistoryFailure) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["username"] = historyFailure.Username
    leafs["remote-name"] = historyFailure.RemoteName
    leafs["username-xr"] = historyFailure.UsernameXr
    leafs["domain-name"] = historyFailure.DomainName
    leafs["mid"] = historyFailure.Mid
    leafs["nas"] = historyFailure.Nas
    leafs["destination-address"] = historyFailure.DestinationAddress
    leafs["remote-client-id"] = historyFailure.RemoteClientId
    leafs["home-gateway"] = historyFailure.HomeGateway
    leafs["source-address"] = historyFailure.SourceAddress
    leafs["local-client-id"] = historyFailure.LocalClientId
    leafs["event-time"] = historyFailure.EventTime
    leafs["error-repeat-count"] = historyFailure.ErrorRepeatCount
    leafs["failure-type"] = historyFailure.FailureType
    return leafs
}

func (historyFailure *Vpdn_HistoryFailures_HistoryFailure) GetBundleName() string { return "cisco_ios_xr" }

func (historyFailure *Vpdn_HistoryFailures_HistoryFailure) GetYangName() string { return "history-failure" }

func (historyFailure *Vpdn_HistoryFailures_HistoryFailure) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (historyFailure *Vpdn_HistoryFailures_HistoryFailure) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (historyFailure *Vpdn_HistoryFailures_HistoryFailure) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (historyFailure *Vpdn_HistoryFailures_HistoryFailure) SetParent(parent types.Entity) { historyFailure.parent = parent }

func (historyFailure *Vpdn_HistoryFailures_HistoryFailure) GetParent() types.Entity { return historyFailure.parent }

func (historyFailure *Vpdn_HistoryFailures_HistoryFailure) GetParentYangName() string { return "history-failures" }

