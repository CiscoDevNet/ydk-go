// This module contains a collection of YANG definitions
// for Cisco IOS-XR tunnel-vpdn package operational data.
// 
// This module contains definitions
// for the following management objects:
//   vpdn: VPDN operational data
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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

    // List of nodes for which subscriber data is collected.
    Nodes Vpdn_Nodes
}

func (vpdn *Vpdn) GetEntityData() *types.CommonEntityData {
    vpdn.EntityData.YFilter = vpdn.YFilter
    vpdn.EntityData.YangName = "vpdn"
    vpdn.EntityData.BundleName = "cisco_ios_xr"
    vpdn.EntityData.ParentYangName = "Cisco-IOS-XR-tunnel-vpdn-oper"
    vpdn.EntityData.SegmentPath = "Cisco-IOS-XR-tunnel-vpdn-oper:vpdn"
    vpdn.EntityData.AbsolutePath = vpdn.EntityData.SegmentPath
    vpdn.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vpdn.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vpdn.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vpdn.EntityData.Children = types.NewOrderedMap()
    vpdn.EntityData.Children.Append("nodes", types.YChild{"Nodes", &vpdn.Nodes})
    vpdn.EntityData.Leafs = types.NewOrderedMap()

    vpdn.EntityData.YListKeys = []string {}

    return &(vpdn.EntityData)
}

// Vpdn_Nodes
// List of nodes for which subscriber data is
// collected
type Vpdn_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Subscriber data for a particular node. The type is slice of
    // Vpdn_Nodes_Node.
    Node []*Vpdn_Nodes_Node
}

func (nodes *Vpdn_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "vpdn"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.AbsolutePath = "Cisco-IOS-XR-tunnel-vpdn-oper:vpdn/" + nodes.EntityData.SegmentPath
    nodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodes.EntityData.Children = types.NewOrderedMap()
    nodes.EntityData.Children.Append("node", types.YChild{"Node", nil})
    for i := range nodes.Node {
        nodes.EntityData.Children.Append(types.GetSegmentPath(nodes.Node[i]), types.YChild{"Node", nodes.Node[i]})
    }
    nodes.EntityData.Leafs = types.NewOrderedMap()

    nodes.EntityData.YListKeys = []string {}

    return &(nodes.EntityData)
}

// Vpdn_Nodes_Node
// Subscriber data for a particular node
type Vpdn_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Node name. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeName interface{}

    // VPDN session list.
    Sessions Vpdn_Nodes_Node_Sessions

    // VPDN tunnel Destinations.
    TunnelDestinations Vpdn_Nodes_Node_TunnelDestinations

    // VPDN Mirroring Statistics.
    VpdnMirroring Vpdn_Nodes_Node_VpdnMirroring

    // Show VPDN Redundancy information.
    VpdnRedundancy Vpdn_Nodes_Node_VpdnRedundancy

    // VPDN history failure list.
    HistoryFailures Vpdn_Nodes_Node_HistoryFailures
}

func (node *Vpdn_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeName, "node-name")
    node.EntityData.AbsolutePath = "Cisco-IOS-XR-tunnel-vpdn-oper:vpdn/nodes/" + node.EntityData.SegmentPath
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("sessions", types.YChild{"Sessions", &node.Sessions})
    node.EntityData.Children.Append("tunnel-destinations", types.YChild{"TunnelDestinations", &node.TunnelDestinations})
    node.EntityData.Children.Append("vpdn-mirroring", types.YChild{"VpdnMirroring", &node.VpdnMirroring})
    node.EntityData.Children.Append("vpdn-redundancy", types.YChild{"VpdnRedundancy", &node.VpdnRedundancy})
    node.EntityData.Children.Append("history-failures", types.YChild{"HistoryFailures", &node.HistoryFailures})
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", node.NodeName})

    node.EntityData.YListKeys = []string {"NodeName"}

    return &(node.EntityData)
}

// Vpdn_Nodes_Node_Sessions
// VPDN session list
type Vpdn_Nodes_Node_Sessions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VPDN session information. The type is slice of
    // Vpdn_Nodes_Node_Sessions_Session.
    Session []*Vpdn_Nodes_Node_Sessions_Session
}

func (sessions *Vpdn_Nodes_Node_Sessions) GetEntityData() *types.CommonEntityData {
    sessions.EntityData.YFilter = sessions.YFilter
    sessions.EntityData.YangName = "sessions"
    sessions.EntityData.BundleName = "cisco_ios_xr"
    sessions.EntityData.ParentYangName = "node"
    sessions.EntityData.SegmentPath = "sessions"
    sessions.EntityData.AbsolutePath = "Cisco-IOS-XR-tunnel-vpdn-oper:vpdn/nodes/node/" + sessions.EntityData.SegmentPath
    sessions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessions.EntityData.Children = types.NewOrderedMap()
    sessions.EntityData.Children.Append("session", types.YChild{"Session", nil})
    for i := range sessions.Session {
        sessions.EntityData.Children.Append(types.GetSegmentPath(sessions.Session[i]), types.YChild{"Session", sessions.Session[i]})
    }
    sessions.EntityData.Leafs = types.NewOrderedMap()

    sessions.EntityData.YListKeys = []string {}

    return &(sessions.EntityData)
}

// Vpdn_Nodes_Node_Sessions_Session
//  VPDN session information
type Vpdn_Nodes_Node_Sessions_Session struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Session label. The type is string with pattern:
    // b'[0-9a-fA-F]{1,8}'.
    SessionLabel interface{}

    // Time to setup session in sec:msec. The type is interface{} with range:
    // 0..4294967295.
    SetupTime interface{}

    // Parent interface name. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
    ParentInterfaceName interface{}

    // Session data.
    Session Vpdn_Nodes_Node_Sessions_Session_Session

    // L2TP data.
    L2tp Vpdn_Nodes_Node_Sessions_Session_L2tp

    // Subscriber data.
    Subscriber Vpdn_Nodes_Node_Sessions_Session_Subscriber

    // Configuration data.
    Configuration Vpdn_Nodes_Node_Sessions_Session_Configuration
}

func (session *Vpdn_Nodes_Node_Sessions_Session) GetEntityData() *types.CommonEntityData {
    session.EntityData.YFilter = session.YFilter
    session.EntityData.YangName = "session"
    session.EntityData.BundleName = "cisco_ios_xr"
    session.EntityData.ParentYangName = "sessions"
    session.EntityData.SegmentPath = "session" + types.AddKeyToken(session.SessionLabel, "session-label")
    session.EntityData.AbsolutePath = "Cisco-IOS-XR-tunnel-vpdn-oper:vpdn/nodes/node/sessions/" + session.EntityData.SegmentPath
    session.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    session.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    session.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    session.EntityData.Children = types.NewOrderedMap()
    session.EntityData.Children.Append("session", types.YChild{"Session", &session.Session})
    session.EntityData.Children.Append("l2tp", types.YChild{"L2tp", &session.L2tp})
    session.EntityData.Children.Append("subscriber", types.YChild{"Subscriber", &session.Subscriber})
    session.EntityData.Children.Append("configuration", types.YChild{"Configuration", &session.Configuration})
    session.EntityData.Leafs = types.NewOrderedMap()
    session.EntityData.Leafs.Append("session-label", types.YLeaf{"SessionLabel", session.SessionLabel})
    session.EntityData.Leafs.Append("setup-time", types.YLeaf{"SetupTime", session.SetupTime})
    session.EntityData.Leafs.Append("parent-interface-name", types.YLeaf{"ParentInterfaceName", session.ParentInterfaceName})

    session.EntityData.YListKeys = []string {"SessionLabel"}

    return &(session.EntityData)
}

// Vpdn_Nodes_Node_Sessions_Session_Session
// Session data
type Vpdn_Nodes_Node_Sessions_Session_Session struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Elapsed time since last change in hh:mm:ss format. The type is string.
    LastChange interface{}

    // Session interface name. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}

    // Authentication username. The type is string.
    Username interface{}

    // Domain name. The type is string.
    DomainName interface{}

    // Session state. The type is SessionState.
    State interface{}

    // L2TP session ID. The type is interface{} with range: 0..65535.
    L2tpSessionId interface{}

    // L2TP tunnel ID. The type is interface{} with range: 0..65535.
    L2tpTunnelId interface{}

    // Session SRG Slave. The type is bool.
    SrgSlave interface{}
}

func (session *Vpdn_Nodes_Node_Sessions_Session_Session) GetEntityData() *types.CommonEntityData {
    session.EntityData.YFilter = session.YFilter
    session.EntityData.YangName = "session"
    session.EntityData.BundleName = "cisco_ios_xr"
    session.EntityData.ParentYangName = "session"
    session.EntityData.SegmentPath = "session"
    session.EntityData.AbsolutePath = "Cisco-IOS-XR-tunnel-vpdn-oper:vpdn/nodes/node/sessions/session/" + session.EntityData.SegmentPath
    session.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    session.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    session.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    session.EntityData.Children = types.NewOrderedMap()
    session.EntityData.Leafs = types.NewOrderedMap()
    session.EntityData.Leafs.Append("last-change", types.YLeaf{"LastChange", session.LastChange})
    session.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", session.InterfaceName})
    session.EntityData.Leafs.Append("username", types.YLeaf{"Username", session.Username})
    session.EntityData.Leafs.Append("domain-name", types.YLeaf{"DomainName", session.DomainName})
    session.EntityData.Leafs.Append("state", types.YLeaf{"State", session.State})
    session.EntityData.Leafs.Append("l2tp-session-id", types.YLeaf{"L2tpSessionId", session.L2tpSessionId})
    session.EntityData.Leafs.Append("l2tp-tunnel-id", types.YLeaf{"L2tpTunnelId", session.L2tpTunnelId})
    session.EntityData.Leafs.Append("srg-slave", types.YLeaf{"SrgSlave", session.SrgSlave})

    session.EntityData.YListKeys = []string {}

    return &(session.EntityData)
}

// Vpdn_Nodes_Node_Sessions_Session_L2tp
// L2TP data
type Vpdn_Nodes_Node_Sessions_Session_L2tp struct {
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
    IsL2tpClassAttributeMaskSet interface{}

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

func (l2tp *Vpdn_Nodes_Node_Sessions_Session_L2tp) GetEntityData() *types.CommonEntityData {
    l2tp.EntityData.YFilter = l2tp.YFilter
    l2tp.EntityData.YangName = "l2tp"
    l2tp.EntityData.BundleName = "cisco_ios_xr"
    l2tp.EntityData.ParentYangName = "session"
    l2tp.EntityData.SegmentPath = "l2tp"
    l2tp.EntityData.AbsolutePath = "Cisco-IOS-XR-tunnel-vpdn-oper:vpdn/nodes/node/sessions/session/" + l2tp.EntityData.SegmentPath
    l2tp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    l2tp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    l2tp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    l2tp.EntityData.Children = types.NewOrderedMap()
    l2tp.EntityData.Leafs = types.NewOrderedMap()
    l2tp.EntityData.Leafs.Append("local-endpoint", types.YLeaf{"LocalEndpoint", l2tp.LocalEndpoint})
    l2tp.EntityData.Leafs.Append("remote-endpoint", types.YLeaf{"RemoteEndpoint", l2tp.RemoteEndpoint})
    l2tp.EntityData.Leafs.Append("call-serial-number", types.YLeaf{"CallSerialNumber", l2tp.CallSerialNumber})
    l2tp.EntityData.Leafs.Append("is-l2tp-class-attribute-mask-set", types.YLeaf{"IsL2tpClassAttributeMaskSet", l2tp.IsL2tpClassAttributeMaskSet})
    l2tp.EntityData.Leafs.Append("local-tunnel-id", types.YLeaf{"LocalTunnelId", l2tp.LocalTunnelId})
    l2tp.EntityData.Leafs.Append("remote-tunnel-id", types.YLeaf{"RemoteTunnelId", l2tp.RemoteTunnelId})
    l2tp.EntityData.Leafs.Append("local-session-id", types.YLeaf{"LocalSessionId", l2tp.LocalSessionId})
    l2tp.EntityData.Leafs.Append("remote-session-id", types.YLeaf{"RemoteSessionId", l2tp.RemoteSessionId})
    l2tp.EntityData.Leafs.Append("remote-port", types.YLeaf{"RemotePort", l2tp.RemotePort})
    l2tp.EntityData.Leafs.Append("tunnel-client-authentication-id", types.YLeaf{"TunnelClientAuthenticationId", l2tp.TunnelClientAuthenticationId})
    l2tp.EntityData.Leafs.Append("tunnel-server-authentication-id", types.YLeaf{"TunnelServerAuthenticationId", l2tp.TunnelServerAuthenticationId})
    l2tp.EntityData.Leafs.Append("tunnel-assignment-id", types.YLeaf{"TunnelAssignmentId", l2tp.TunnelAssignmentId})
    l2tp.EntityData.Leafs.Append("is-tunnel-authentication-enabled", types.YLeaf{"IsTunnelAuthenticationEnabled", l2tp.IsTunnelAuthenticationEnabled})

    l2tp.EntityData.YListKeys = []string {}

    return &(l2tp.EntityData)
}

// Vpdn_Nodes_Node_Sessions_Session_Subscriber
// Subscriber data
type Vpdn_Nodes_Node_Sessions_Session_Subscriber struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // NAS port ID Val. The type is string.
    NasPortIdVal interface{}

    // NAS port type. The type is string.
    NasPortType interface{}

    // Physical channel ID. The type is interface{} with range: 0..4294967295.
    PhysicalChannelId interface{}

    // Receive connect speed in nanoseconds. The type is interface{} with range:
    // 0..18446744073709551615. Units are nanosecond.
    ReceiveConnectSpeed interface{}

    // Transmit connect speed in nanoseconds. The type is interface{} with range:
    // 0..18446744073709551615. Units are nanosecond.
    TransmitConnectSpeed interface{}
}

func (subscriber *Vpdn_Nodes_Node_Sessions_Session_Subscriber) GetEntityData() *types.CommonEntityData {
    subscriber.EntityData.YFilter = subscriber.YFilter
    subscriber.EntityData.YangName = "subscriber"
    subscriber.EntityData.BundleName = "cisco_ios_xr"
    subscriber.EntityData.ParentYangName = "session"
    subscriber.EntityData.SegmentPath = "subscriber"
    subscriber.EntityData.AbsolutePath = "Cisco-IOS-XR-tunnel-vpdn-oper:vpdn/nodes/node/sessions/session/" + subscriber.EntityData.SegmentPath
    subscriber.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    subscriber.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    subscriber.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    subscriber.EntityData.Children = types.NewOrderedMap()
    subscriber.EntityData.Leafs = types.NewOrderedMap()
    subscriber.EntityData.Leafs.Append("nas-port-id-val", types.YLeaf{"NasPortIdVal", subscriber.NasPortIdVal})
    subscriber.EntityData.Leafs.Append("nas-port-type", types.YLeaf{"NasPortType", subscriber.NasPortType})
    subscriber.EntityData.Leafs.Append("physical-channel-id", types.YLeaf{"PhysicalChannelId", subscriber.PhysicalChannelId})
    subscriber.EntityData.Leafs.Append("receive-connect-speed", types.YLeaf{"ReceiveConnectSpeed", subscriber.ReceiveConnectSpeed})
    subscriber.EntityData.Leafs.Append("transmit-connect-speed", types.YLeaf{"TransmitConnectSpeed", subscriber.TransmitConnectSpeed})

    subscriber.EntityData.YListKeys = []string {}

    return &(subscriber.EntityData)
}

// Vpdn_Nodes_Node_Sessions_Session_Configuration
// Configuration data
type Vpdn_Nodes_Node_Sessions_Session_Configuration struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Template name. The type is string.
    TemplateName interface{}

    // VRF name. The type is string.
    VrfName interface{}

    // L2TP busy timeout in seconds. The type is interface{} with range: 0..65535.
    // Units are second.
    L2tpBusyTimeout interface{}

    // TOS mode. The type is TosMode.
    TosMode interface{}

    // TOS setting value. The type is interface{} with range: 0..255.
    Tos interface{}

    // True if DSL line info forwarding is enabled. The type is bool.
    DslLineForwarding interface{}

    // VPN ID.
    VpnId Vpdn_Nodes_Node_Sessions_Session_Configuration_VpnId
}

func (configuration *Vpdn_Nodes_Node_Sessions_Session_Configuration) GetEntityData() *types.CommonEntityData {
    configuration.EntityData.YFilter = configuration.YFilter
    configuration.EntityData.YangName = "configuration"
    configuration.EntityData.BundleName = "cisco_ios_xr"
    configuration.EntityData.ParentYangName = "session"
    configuration.EntityData.SegmentPath = "configuration"
    configuration.EntityData.AbsolutePath = "Cisco-IOS-XR-tunnel-vpdn-oper:vpdn/nodes/node/sessions/session/" + configuration.EntityData.SegmentPath
    configuration.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configuration.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configuration.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configuration.EntityData.Children = types.NewOrderedMap()
    configuration.EntityData.Children.Append("vpn-id", types.YChild{"VpnId", &configuration.VpnId})
    configuration.EntityData.Leafs = types.NewOrderedMap()
    configuration.EntityData.Leafs.Append("template-name", types.YLeaf{"TemplateName", configuration.TemplateName})
    configuration.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", configuration.VrfName})
    configuration.EntityData.Leafs.Append("l2tp-busy-timeout", types.YLeaf{"L2tpBusyTimeout", configuration.L2tpBusyTimeout})
    configuration.EntityData.Leafs.Append("tos-mode", types.YLeaf{"TosMode", configuration.TosMode})
    configuration.EntityData.Leafs.Append("tos", types.YLeaf{"Tos", configuration.Tos})
    configuration.EntityData.Leafs.Append("dsl-line-forwarding", types.YLeaf{"DslLineForwarding", configuration.DslLineForwarding})

    configuration.EntityData.YListKeys = []string {}

    return &(configuration.EntityData)
}

// Vpdn_Nodes_Node_Sessions_Session_Configuration_VpnId
// VPN ID
type Vpdn_Nodes_Node_Sessions_Session_Configuration_VpnId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OUI. The type is interface{} with range: 0..4294967295.
    Oui interface{}

    // Index. The type is interface{} with range: 0..4294967295.
    Index interface{}
}

func (vpnId *Vpdn_Nodes_Node_Sessions_Session_Configuration_VpnId) GetEntityData() *types.CommonEntityData {
    vpnId.EntityData.YFilter = vpnId.YFilter
    vpnId.EntityData.YangName = "vpn-id"
    vpnId.EntityData.BundleName = "cisco_ios_xr"
    vpnId.EntityData.ParentYangName = "configuration"
    vpnId.EntityData.SegmentPath = "vpn-id"
    vpnId.EntityData.AbsolutePath = "Cisco-IOS-XR-tunnel-vpdn-oper:vpdn/nodes/node/sessions/session/configuration/" + vpnId.EntityData.SegmentPath
    vpnId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vpnId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vpnId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vpnId.EntityData.Children = types.NewOrderedMap()
    vpnId.EntityData.Leafs = types.NewOrderedMap()
    vpnId.EntityData.Leafs.Append("oui", types.YLeaf{"Oui", vpnId.Oui})
    vpnId.EntityData.Leafs.Append("index", types.YLeaf{"Index", vpnId.Index})

    vpnId.EntityData.YListKeys = []string {}

    return &(vpnId.EntityData)
}

// Vpdn_Nodes_Node_TunnelDestinations
// VPDN tunnel Destinations
type Vpdn_Nodes_Node_TunnelDestinations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VPDN tunnel destination information. The type is slice of
    // Vpdn_Nodes_Node_TunnelDestinations_TunnelDestination.
    TunnelDestination []*Vpdn_Nodes_Node_TunnelDestinations_TunnelDestination
}

func (tunnelDestinations *Vpdn_Nodes_Node_TunnelDestinations) GetEntityData() *types.CommonEntityData {
    tunnelDestinations.EntityData.YFilter = tunnelDestinations.YFilter
    tunnelDestinations.EntityData.YangName = "tunnel-destinations"
    tunnelDestinations.EntityData.BundleName = "cisco_ios_xr"
    tunnelDestinations.EntityData.ParentYangName = "node"
    tunnelDestinations.EntityData.SegmentPath = "tunnel-destinations"
    tunnelDestinations.EntityData.AbsolutePath = "Cisco-IOS-XR-tunnel-vpdn-oper:vpdn/nodes/node/" + tunnelDestinations.EntityData.SegmentPath
    tunnelDestinations.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnelDestinations.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnelDestinations.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnelDestinations.EntityData.Children = types.NewOrderedMap()
    tunnelDestinations.EntityData.Children.Append("tunnel-destination", types.YChild{"TunnelDestination", nil})
    for i := range tunnelDestinations.TunnelDestination {
        types.SetYListKey(tunnelDestinations.TunnelDestination[i], i)
        tunnelDestinations.EntityData.Children.Append(types.GetSegmentPath(tunnelDestinations.TunnelDestination[i]), types.YChild{"TunnelDestination", tunnelDestinations.TunnelDestination[i]})
    }
    tunnelDestinations.EntityData.Leafs = types.NewOrderedMap()

    tunnelDestinations.EntityData.YListKeys = []string {}

    return &(tunnelDestinations.EntityData)
}

// Vpdn_Nodes_Node_TunnelDestinations_TunnelDestination
// VPDN tunnel destination information
type Vpdn_Nodes_Node_TunnelDestinations_TunnelDestination struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

func (tunnelDestination *Vpdn_Nodes_Node_TunnelDestinations_TunnelDestination) GetEntityData() *types.CommonEntityData {
    tunnelDestination.EntityData.YFilter = tunnelDestination.YFilter
    tunnelDestination.EntityData.YangName = "tunnel-destination"
    tunnelDestination.EntityData.BundleName = "cisco_ios_xr"
    tunnelDestination.EntityData.ParentYangName = "tunnel-destinations"
    tunnelDestination.EntityData.SegmentPath = "tunnel-destination" + types.AddNoKeyToken(tunnelDestination)
    tunnelDestination.EntityData.AbsolutePath = "Cisco-IOS-XR-tunnel-vpdn-oper:vpdn/nodes/node/tunnel-destinations/" + tunnelDestination.EntityData.SegmentPath
    tunnelDestination.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnelDestination.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnelDestination.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnelDestination.EntityData.Children = types.NewOrderedMap()
    tunnelDestination.EntityData.Leafs = types.NewOrderedMap()
    tunnelDestination.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", tunnelDestination.VrfName})
    tunnelDestination.EntityData.Leafs.Append("address", types.YLeaf{"Address", tunnelDestination.Address})
    tunnelDestination.EntityData.Leafs.Append("vrf-name-xr", types.YLeaf{"VrfNameXr", tunnelDestination.VrfNameXr})
    tunnelDestination.EntityData.Leafs.Append("load", types.YLeaf{"Load", tunnelDestination.Load})
    tunnelDestination.EntityData.Leafs.Append("status", types.YLeaf{"Status", tunnelDestination.Status})
    tunnelDestination.EntityData.Leafs.Append("connects", types.YLeaf{"Connects", tunnelDestination.Connects})
    tunnelDestination.EntityData.Leafs.Append("disconnects", types.YLeaf{"Disconnects", tunnelDestination.Disconnects})
    tunnelDestination.EntityData.Leafs.Append("retry", types.YLeaf{"Retry", tunnelDestination.Retry})
    tunnelDestination.EntityData.Leafs.Append("status-change-time", types.YLeaf{"StatusChangeTime", tunnelDestination.StatusChangeTime})

    tunnelDestination.EntityData.YListKeys = []string {}

    return &(tunnelDestination.EntityData)
}

// Vpdn_Nodes_Node_VpdnMirroring
// VPDN Mirroring Statistics
type Vpdn_Nodes_Node_VpdnMirroring struct {
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
    QadSendStats Vpdn_Nodes_Node_VpdnMirroring_QadSendStats

    // qad recv stats.
    QadRecvStats Vpdn_Nodes_Node_VpdnMirroring_QadRecvStats

    // qad send stats last clear.
    QadSendStatsLastClear Vpdn_Nodes_Node_VpdnMirroring_QadSendStatsLastClear

    // qad recv stats last clear.
    QadRecvStatsLastClear Vpdn_Nodes_Node_VpdnMirroring_QadRecvStatsLastClear
}

func (vpdnMirroring *Vpdn_Nodes_Node_VpdnMirroring) GetEntityData() *types.CommonEntityData {
    vpdnMirroring.EntityData.YFilter = vpdnMirroring.YFilter
    vpdnMirroring.EntityData.YangName = "vpdn-mirroring"
    vpdnMirroring.EntityData.BundleName = "cisco_ios_xr"
    vpdnMirroring.EntityData.ParentYangName = "node"
    vpdnMirroring.EntityData.SegmentPath = "vpdn-mirroring"
    vpdnMirroring.EntityData.AbsolutePath = "Cisco-IOS-XR-tunnel-vpdn-oper:vpdn/nodes/node/" + vpdnMirroring.EntityData.SegmentPath
    vpdnMirroring.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vpdnMirroring.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vpdnMirroring.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vpdnMirroring.EntityData.Children = types.NewOrderedMap()
    vpdnMirroring.EntityData.Children.Append("qad-send-stats", types.YChild{"QadSendStats", &vpdnMirroring.QadSendStats})
    vpdnMirroring.EntityData.Children.Append("qad-recv-stats", types.YChild{"QadRecvStats", &vpdnMirroring.QadRecvStats})
    vpdnMirroring.EntityData.Children.Append("qad-send-stats-last-clear", types.YChild{"QadSendStatsLastClear", &vpdnMirroring.QadSendStatsLastClear})
    vpdnMirroring.EntityData.Children.Append("qad-recv-stats-last-clear", types.YChild{"QadRecvStatsLastClear", &vpdnMirroring.QadRecvStatsLastClear})
    vpdnMirroring.EntityData.Leafs = types.NewOrderedMap()
    vpdnMirroring.EntityData.Leafs.Append("sync-not-conn-cnt", types.YLeaf{"SyncNotConnCnt", vpdnMirroring.SyncNotConnCnt})
    vpdnMirroring.EntityData.Leafs.Append("sso-err-cnt", types.YLeaf{"SsoErrCnt", vpdnMirroring.SsoErrCnt})
    vpdnMirroring.EntityData.Leafs.Append("sso-batch-err-cnt", types.YLeaf{"SsoBatchErrCnt", vpdnMirroring.SsoBatchErrCnt})
    vpdnMirroring.EntityData.Leafs.Append("alloc-err-cnt", types.YLeaf{"AllocErrCnt", vpdnMirroring.AllocErrCnt})
    vpdnMirroring.EntityData.Leafs.Append("alloc-cnt", types.YLeaf{"AllocCnt", vpdnMirroring.AllocCnt})

    vpdnMirroring.EntityData.YListKeys = []string {}

    return &(vpdnMirroring.EntityData)
}

// Vpdn_Nodes_Node_VpdnMirroring_QadSendStats
// qad send stats
type Vpdn_Nodes_Node_VpdnMirroring_QadSendStats struct {
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

func (qadSendStats *Vpdn_Nodes_Node_VpdnMirroring_QadSendStats) GetEntityData() *types.CommonEntityData {
    qadSendStats.EntityData.YFilter = qadSendStats.YFilter
    qadSendStats.EntityData.YangName = "qad-send-stats"
    qadSendStats.EntityData.BundleName = "cisco_ios_xr"
    qadSendStats.EntityData.ParentYangName = "vpdn-mirroring"
    qadSendStats.EntityData.SegmentPath = "qad-send-stats"
    qadSendStats.EntityData.AbsolutePath = "Cisco-IOS-XR-tunnel-vpdn-oper:vpdn/nodes/node/vpdn-mirroring/" + qadSendStats.EntityData.SegmentPath
    qadSendStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    qadSendStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    qadSendStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    qadSendStats.EntityData.Children = types.NewOrderedMap()
    qadSendStats.EntityData.Leafs = types.NewOrderedMap()
    qadSendStats.EntityData.Leafs.Append("msgs-sent", types.YLeaf{"MsgsSent", qadSendStats.MsgsSent})
    qadSendStats.EntityData.Leafs.Append("acks-sent", types.YLeaf{"AcksSent", qadSendStats.AcksSent})
    qadSendStats.EntityData.Leafs.Append("no-partner", types.YLeaf{"NoPartner", qadSendStats.NoPartner})
    qadSendStats.EntityData.Leafs.Append("sends-failed", types.YLeaf{"SendsFailed", qadSendStats.SendsFailed})
    qadSendStats.EntityData.Leafs.Append("acks-failed", types.YLeaf{"AcksFailed", qadSendStats.AcksFailed})
    qadSendStats.EntityData.Leafs.Append("pending-acks", types.YLeaf{"PendingAcks", qadSendStats.PendingAcks})
    qadSendStats.EntityData.Leafs.Append("timeouts", types.YLeaf{"Timeouts", qadSendStats.Timeouts})
    qadSendStats.EntityData.Leafs.Append("suspends", types.YLeaf{"Suspends", qadSendStats.Suspends})
    qadSendStats.EntityData.Leafs.Append("resumes", types.YLeaf{"Resumes", qadSendStats.Resumes})
    qadSendStats.EntityData.Leafs.Append("sends-fragment", types.YLeaf{"SendsFragment", qadSendStats.SendsFragment})
    qadSendStats.EntityData.Leafs.Append("qad-last-seq-number", types.YLeaf{"QadLastSeqNumber", qadSendStats.QadLastSeqNumber})
    qadSendStats.EntityData.Leafs.Append("qad-frag-count", types.YLeaf{"QadFragCount", qadSendStats.QadFragCount})
    qadSendStats.EntityData.Leafs.Append("qad-ack-count", types.YLeaf{"QadAckCount", qadSendStats.QadAckCount})
    qadSendStats.EntityData.Leafs.Append("qad-unknown-acks", types.YLeaf{"QadUnknownAcks", qadSendStats.QadUnknownAcks})
    qadSendStats.EntityData.Leafs.Append("qad-timeouts", types.YLeaf{"QadTimeouts", qadSendStats.QadTimeouts})
    qadSendStats.EntityData.Leafs.Append("qad-rx-count", types.YLeaf{"QadRxCount", qadSendStats.QadRxCount})
    qadSendStats.EntityData.Leafs.Append("qad-rx-list-count", types.YLeaf{"QadRxListCount", qadSendStats.QadRxListCount})
    qadSendStats.EntityData.Leafs.Append("qad-rx-list-q-size", types.YLeaf{"QadRxListQSize", qadSendStats.QadRxListQSize})
    qadSendStats.EntityData.Leafs.Append("qad-rx-first-seq-number", types.YLeaf{"QadRxFirstSeqNumber", qadSendStats.QadRxFirstSeqNumber})

    qadSendStats.EntityData.YListKeys = []string {}

    return &(qadSendStats.EntityData)
}

// Vpdn_Nodes_Node_VpdnMirroring_QadRecvStats
// qad recv stats
type Vpdn_Nodes_Node_VpdnMirroring_QadRecvStats struct {
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

func (qadRecvStats *Vpdn_Nodes_Node_VpdnMirroring_QadRecvStats) GetEntityData() *types.CommonEntityData {
    qadRecvStats.EntityData.YFilter = qadRecvStats.YFilter
    qadRecvStats.EntityData.YangName = "qad-recv-stats"
    qadRecvStats.EntityData.BundleName = "cisco_ios_xr"
    qadRecvStats.EntityData.ParentYangName = "vpdn-mirroring"
    qadRecvStats.EntityData.SegmentPath = "qad-recv-stats"
    qadRecvStats.EntityData.AbsolutePath = "Cisco-IOS-XR-tunnel-vpdn-oper:vpdn/nodes/node/vpdn-mirroring/" + qadRecvStats.EntityData.SegmentPath
    qadRecvStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    qadRecvStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    qadRecvStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    qadRecvStats.EntityData.Children = types.NewOrderedMap()
    qadRecvStats.EntityData.Leafs = types.NewOrderedMap()
    qadRecvStats.EntityData.Leafs.Append("msgs-recvd", types.YLeaf{"MsgsRecvd", qadRecvStats.MsgsRecvd})
    qadRecvStats.EntityData.Leafs.Append("acks-recvd", types.YLeaf{"AcksRecvd", qadRecvStats.AcksRecvd})
    qadRecvStats.EntityData.Leafs.Append("recvd-acks-failed", types.YLeaf{"RecvdAcksFailed", qadRecvStats.RecvdAcksFailed})
    qadRecvStats.EntityData.Leafs.Append("init-drops", types.YLeaf{"InitDrops", qadRecvStats.InitDrops})
    qadRecvStats.EntityData.Leafs.Append("msg-drops", types.YLeaf{"MsgDrops", qadRecvStats.MsgDrops})
    qadRecvStats.EntityData.Leafs.Append("ooo-drops", types.YLeaf{"OooDrops", qadRecvStats.OooDrops})
    qadRecvStats.EntityData.Leafs.Append("stale-msgs", types.YLeaf{"StaleMsgs", qadRecvStats.StaleMsgs})

    qadRecvStats.EntityData.YListKeys = []string {}

    return &(qadRecvStats.EntityData)
}

// Vpdn_Nodes_Node_VpdnMirroring_QadSendStatsLastClear
// qad send stats last clear
type Vpdn_Nodes_Node_VpdnMirroring_QadSendStatsLastClear struct {
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

func (qadSendStatsLastClear *Vpdn_Nodes_Node_VpdnMirroring_QadSendStatsLastClear) GetEntityData() *types.CommonEntityData {
    qadSendStatsLastClear.EntityData.YFilter = qadSendStatsLastClear.YFilter
    qadSendStatsLastClear.EntityData.YangName = "qad-send-stats-last-clear"
    qadSendStatsLastClear.EntityData.BundleName = "cisco_ios_xr"
    qadSendStatsLastClear.EntityData.ParentYangName = "vpdn-mirroring"
    qadSendStatsLastClear.EntityData.SegmentPath = "qad-send-stats-last-clear"
    qadSendStatsLastClear.EntityData.AbsolutePath = "Cisco-IOS-XR-tunnel-vpdn-oper:vpdn/nodes/node/vpdn-mirroring/" + qadSendStatsLastClear.EntityData.SegmentPath
    qadSendStatsLastClear.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    qadSendStatsLastClear.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    qadSendStatsLastClear.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    qadSendStatsLastClear.EntityData.Children = types.NewOrderedMap()
    qadSendStatsLastClear.EntityData.Leafs = types.NewOrderedMap()
    qadSendStatsLastClear.EntityData.Leafs.Append("msgs-sent", types.YLeaf{"MsgsSent", qadSendStatsLastClear.MsgsSent})
    qadSendStatsLastClear.EntityData.Leafs.Append("acks-sent", types.YLeaf{"AcksSent", qadSendStatsLastClear.AcksSent})
    qadSendStatsLastClear.EntityData.Leafs.Append("no-partner", types.YLeaf{"NoPartner", qadSendStatsLastClear.NoPartner})
    qadSendStatsLastClear.EntityData.Leafs.Append("sends-failed", types.YLeaf{"SendsFailed", qadSendStatsLastClear.SendsFailed})
    qadSendStatsLastClear.EntityData.Leafs.Append("acks-failed", types.YLeaf{"AcksFailed", qadSendStatsLastClear.AcksFailed})
    qadSendStatsLastClear.EntityData.Leafs.Append("pending-acks", types.YLeaf{"PendingAcks", qadSendStatsLastClear.PendingAcks})
    qadSendStatsLastClear.EntityData.Leafs.Append("timeouts", types.YLeaf{"Timeouts", qadSendStatsLastClear.Timeouts})
    qadSendStatsLastClear.EntityData.Leafs.Append("suspends", types.YLeaf{"Suspends", qadSendStatsLastClear.Suspends})
    qadSendStatsLastClear.EntityData.Leafs.Append("resumes", types.YLeaf{"Resumes", qadSendStatsLastClear.Resumes})
    qadSendStatsLastClear.EntityData.Leafs.Append("sends-fragment", types.YLeaf{"SendsFragment", qadSendStatsLastClear.SendsFragment})
    qadSendStatsLastClear.EntityData.Leafs.Append("qad-last-seq-number", types.YLeaf{"QadLastSeqNumber", qadSendStatsLastClear.QadLastSeqNumber})
    qadSendStatsLastClear.EntityData.Leafs.Append("qad-frag-count", types.YLeaf{"QadFragCount", qadSendStatsLastClear.QadFragCount})
    qadSendStatsLastClear.EntityData.Leafs.Append("qad-ack-count", types.YLeaf{"QadAckCount", qadSendStatsLastClear.QadAckCount})
    qadSendStatsLastClear.EntityData.Leafs.Append("qad-unknown-acks", types.YLeaf{"QadUnknownAcks", qadSendStatsLastClear.QadUnknownAcks})
    qadSendStatsLastClear.EntityData.Leafs.Append("qad-timeouts", types.YLeaf{"QadTimeouts", qadSendStatsLastClear.QadTimeouts})
    qadSendStatsLastClear.EntityData.Leafs.Append("qad-rx-count", types.YLeaf{"QadRxCount", qadSendStatsLastClear.QadRxCount})
    qadSendStatsLastClear.EntityData.Leafs.Append("qad-rx-list-count", types.YLeaf{"QadRxListCount", qadSendStatsLastClear.QadRxListCount})
    qadSendStatsLastClear.EntityData.Leafs.Append("qad-rx-list-q-size", types.YLeaf{"QadRxListQSize", qadSendStatsLastClear.QadRxListQSize})
    qadSendStatsLastClear.EntityData.Leafs.Append("qad-rx-first-seq-number", types.YLeaf{"QadRxFirstSeqNumber", qadSendStatsLastClear.QadRxFirstSeqNumber})

    qadSendStatsLastClear.EntityData.YListKeys = []string {}

    return &(qadSendStatsLastClear.EntityData)
}

// Vpdn_Nodes_Node_VpdnMirroring_QadRecvStatsLastClear
// qad recv stats last clear
type Vpdn_Nodes_Node_VpdnMirroring_QadRecvStatsLastClear struct {
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

func (qadRecvStatsLastClear *Vpdn_Nodes_Node_VpdnMirroring_QadRecvStatsLastClear) GetEntityData() *types.CommonEntityData {
    qadRecvStatsLastClear.EntityData.YFilter = qadRecvStatsLastClear.YFilter
    qadRecvStatsLastClear.EntityData.YangName = "qad-recv-stats-last-clear"
    qadRecvStatsLastClear.EntityData.BundleName = "cisco_ios_xr"
    qadRecvStatsLastClear.EntityData.ParentYangName = "vpdn-mirroring"
    qadRecvStatsLastClear.EntityData.SegmentPath = "qad-recv-stats-last-clear"
    qadRecvStatsLastClear.EntityData.AbsolutePath = "Cisco-IOS-XR-tunnel-vpdn-oper:vpdn/nodes/node/vpdn-mirroring/" + qadRecvStatsLastClear.EntityData.SegmentPath
    qadRecvStatsLastClear.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    qadRecvStatsLastClear.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    qadRecvStatsLastClear.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    qadRecvStatsLastClear.EntityData.Children = types.NewOrderedMap()
    qadRecvStatsLastClear.EntityData.Leafs = types.NewOrderedMap()
    qadRecvStatsLastClear.EntityData.Leafs.Append("msgs-recvd", types.YLeaf{"MsgsRecvd", qadRecvStatsLastClear.MsgsRecvd})
    qadRecvStatsLastClear.EntityData.Leafs.Append("acks-recvd", types.YLeaf{"AcksRecvd", qadRecvStatsLastClear.AcksRecvd})
    qadRecvStatsLastClear.EntityData.Leafs.Append("recvd-acks-failed", types.YLeaf{"RecvdAcksFailed", qadRecvStatsLastClear.RecvdAcksFailed})
    qadRecvStatsLastClear.EntityData.Leafs.Append("init-drops", types.YLeaf{"InitDrops", qadRecvStatsLastClear.InitDrops})
    qadRecvStatsLastClear.EntityData.Leafs.Append("msg-drops", types.YLeaf{"MsgDrops", qadRecvStatsLastClear.MsgDrops})
    qadRecvStatsLastClear.EntityData.Leafs.Append("ooo-drops", types.YLeaf{"OooDrops", qadRecvStatsLastClear.OooDrops})
    qadRecvStatsLastClear.EntityData.Leafs.Append("stale-msgs", types.YLeaf{"StaleMsgs", qadRecvStatsLastClear.StaleMsgs})

    qadRecvStatsLastClear.EntityData.YListKeys = []string {}

    return &(qadRecvStatsLastClear.EntityData)
}

// Vpdn_Nodes_Node_VpdnRedundancy
// Show VPDN Redundancy information
type Vpdn_Nodes_Node_VpdnRedundancy struct {
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

func (vpdnRedundancy *Vpdn_Nodes_Node_VpdnRedundancy) GetEntityData() *types.CommonEntityData {
    vpdnRedundancy.EntityData.YFilter = vpdnRedundancy.YFilter
    vpdnRedundancy.EntityData.YangName = "vpdn-redundancy"
    vpdnRedundancy.EntityData.BundleName = "cisco_ios_xr"
    vpdnRedundancy.EntityData.ParentYangName = "node"
    vpdnRedundancy.EntityData.SegmentPath = "vpdn-redundancy"
    vpdnRedundancy.EntityData.AbsolutePath = "Cisco-IOS-XR-tunnel-vpdn-oper:vpdn/nodes/node/" + vpdnRedundancy.EntityData.SegmentPath
    vpdnRedundancy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vpdnRedundancy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vpdnRedundancy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vpdnRedundancy.EntityData.Children = types.NewOrderedMap()
    vpdnRedundancy.EntityData.Leafs = types.NewOrderedMap()
    vpdnRedundancy.EntityData.Leafs.Append("session-total", types.YLeaf{"SessionTotal", vpdnRedundancy.SessionTotal})
    vpdnRedundancy.EntityData.Leafs.Append("session-synced", types.YLeaf{"SessionSynced", vpdnRedundancy.SessionSynced})
    vpdnRedundancy.EntityData.Leafs.Append("state", types.YLeaf{"State", vpdnRedundancy.State})
    vpdnRedundancy.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", vpdnRedundancy.StartTime})
    vpdnRedundancy.EntityData.Leafs.Append("finish-time", types.YLeaf{"FinishTime", vpdnRedundancy.FinishTime})
    vpdnRedundancy.EntityData.Leafs.Append("abort-time", types.YLeaf{"AbortTime", vpdnRedundancy.AbortTime})

    vpdnRedundancy.EntityData.YListKeys = []string {}

    return &(vpdnRedundancy.EntityData)
}

// Vpdn_Nodes_Node_HistoryFailures
// VPDN history failure list
type Vpdn_Nodes_Node_HistoryFailures struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VPDN history failure information. The type is slice of
    // Vpdn_Nodes_Node_HistoryFailures_HistoryFailure.
    HistoryFailure []*Vpdn_Nodes_Node_HistoryFailures_HistoryFailure
}

func (historyFailures *Vpdn_Nodes_Node_HistoryFailures) GetEntityData() *types.CommonEntityData {
    historyFailures.EntityData.YFilter = historyFailures.YFilter
    historyFailures.EntityData.YangName = "history-failures"
    historyFailures.EntityData.BundleName = "cisco_ios_xr"
    historyFailures.EntityData.ParentYangName = "node"
    historyFailures.EntityData.SegmentPath = "history-failures"
    historyFailures.EntityData.AbsolutePath = "Cisco-IOS-XR-tunnel-vpdn-oper:vpdn/nodes/node/" + historyFailures.EntityData.SegmentPath
    historyFailures.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    historyFailures.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    historyFailures.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    historyFailures.EntityData.Children = types.NewOrderedMap()
    historyFailures.EntityData.Children.Append("history-failure", types.YChild{"HistoryFailure", nil})
    for i := range historyFailures.HistoryFailure {
        types.SetYListKey(historyFailures.HistoryFailure[i], i)
        historyFailures.EntityData.Children.Append(types.GetSegmentPath(historyFailures.HistoryFailure[i]), types.YChild{"HistoryFailure", historyFailures.HistoryFailure[i]})
    }
    historyFailures.EntityData.Leafs = types.NewOrderedMap()

    historyFailures.EntityData.YListKeys = []string {}

    return &(historyFailures.EntityData)
}

// Vpdn_Nodes_Node_HistoryFailures_HistoryFailure
// VPDN history failure information
type Vpdn_Nodes_Node_HistoryFailures_HistoryFailure struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

func (historyFailure *Vpdn_Nodes_Node_HistoryFailures_HistoryFailure) GetEntityData() *types.CommonEntityData {
    historyFailure.EntityData.YFilter = historyFailure.YFilter
    historyFailure.EntityData.YangName = "history-failure"
    historyFailure.EntityData.BundleName = "cisco_ios_xr"
    historyFailure.EntityData.ParentYangName = "history-failures"
    historyFailure.EntityData.SegmentPath = "history-failure" + types.AddNoKeyToken(historyFailure)
    historyFailure.EntityData.AbsolutePath = "Cisco-IOS-XR-tunnel-vpdn-oper:vpdn/nodes/node/history-failures/" + historyFailure.EntityData.SegmentPath
    historyFailure.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    historyFailure.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    historyFailure.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    historyFailure.EntityData.Children = types.NewOrderedMap()
    historyFailure.EntityData.Leafs = types.NewOrderedMap()
    historyFailure.EntityData.Leafs.Append("username", types.YLeaf{"Username", historyFailure.Username})
    historyFailure.EntityData.Leafs.Append("remote-name", types.YLeaf{"RemoteName", historyFailure.RemoteName})
    historyFailure.EntityData.Leafs.Append("username-xr", types.YLeaf{"UsernameXr", historyFailure.UsernameXr})
    historyFailure.EntityData.Leafs.Append("domain-name", types.YLeaf{"DomainName", historyFailure.DomainName})
    historyFailure.EntityData.Leafs.Append("mid", types.YLeaf{"Mid", historyFailure.Mid})
    historyFailure.EntityData.Leafs.Append("nas", types.YLeaf{"Nas", historyFailure.Nas})
    historyFailure.EntityData.Leafs.Append("destination-address", types.YLeaf{"DestinationAddress", historyFailure.DestinationAddress})
    historyFailure.EntityData.Leafs.Append("remote-client-id", types.YLeaf{"RemoteClientId", historyFailure.RemoteClientId})
    historyFailure.EntityData.Leafs.Append("home-gateway", types.YLeaf{"HomeGateway", historyFailure.HomeGateway})
    historyFailure.EntityData.Leafs.Append("source-address", types.YLeaf{"SourceAddress", historyFailure.SourceAddress})
    historyFailure.EntityData.Leafs.Append("local-client-id", types.YLeaf{"LocalClientId", historyFailure.LocalClientId})
    historyFailure.EntityData.Leafs.Append("event-time", types.YLeaf{"EventTime", historyFailure.EventTime})
    historyFailure.EntityData.Leafs.Append("error-repeat-count", types.YLeaf{"ErrorRepeatCount", historyFailure.ErrorRepeatCount})
    historyFailure.EntityData.Leafs.Append("failure-type", types.YLeaf{"FailureType", historyFailure.FailureType})

    historyFailure.EntityData.YListKeys = []string {}

    return &(historyFailure.EntityData)
}

