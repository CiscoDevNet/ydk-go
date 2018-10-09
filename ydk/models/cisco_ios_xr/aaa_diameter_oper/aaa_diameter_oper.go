// This module contains a collection of YANG definitions
// for Cisco IOS-XR aaa-diameter package operational data.
// 
// This YANG module augments the
//   Cisco-IOS-XR-aaa-locald-oper
// module with state data.
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package aaa_diameter_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package aaa_diameter_oper"))
}

// DisconnectCause represents Disconnect cause values
type DisconnectCause string

const (
    // Disconnect caused by reboot
    DisconnectCause_reboot DisconnectCause = "reboot"

    // Disconnect due to server busy
    DisconnectCause_busy DisconnectCause = "busy"

    // Disconnect as server does not want to talk
    DisconnectCause_do_not_wait_to_talk DisconnectCause = "do-not-wait-to-talk"
)

// SecurityTypeValue represents Security type values
type SecurityTypeValue string

const (
    // No security type
    SecurityTypeValue_security_type_none SecurityTypeValue = "security-type-none"

    // TLS security
    SecurityTypeValue_type_ SecurityTypeValue = "type"

    // IPSEC security
    SecurityTypeValue_ipsec SecurityTypeValue = "ipsec"
)

// PeerStateValue represents Peer State Values
type PeerStateValue string

const (
    // No Peer states
    PeerStateValue_state_none PeerStateValue = "state-none"

    // Peer closed
    PeerStateValue_closed PeerStateValue = "closed"

    // Waiting for ACK
    PeerStateValue_wait_connection_ack PeerStateValue = "wait-connection-ack"

    // Waiting for CEA
    PeerStateValue_wait_cea PeerStateValue = "wait-cea"

    // Peer open
    PeerStateValue_state_open PeerStateValue = "state-open"

    // Peer closed
    PeerStateValue_closing PeerStateValue = "closing"

    // Peer in suspect state
    PeerStateValue_suspect PeerStateValue = "suspect"
)

// ProtocolTypeValue represents Protocol type values
type ProtocolTypeValue string

const (
    // No protocol used
    ProtocolTypeValue_protocol_none ProtocolTypeValue = "protocol-none"

    // TCP protocol
    ProtocolTypeValue_tcp ProtocolTypeValue = "tcp"
)

// Peer represents  Peer type values
type Peer string

const (
    // Peer not defined
    Peer_undefined Peer = "undefined"

    // Server type
    Peer_server Peer = "server"
)

// WhoInitiatedDisconnect represents Who initiated to disconnect
type WhoInitiatedDisconnect string

const (
    // None
    WhoInitiatedDisconnect_none WhoInitiatedDisconnect = "none"

    // Disconnected by host
    WhoInitiatedDisconnect_host WhoInitiatedDisconnect = "host"

    // Disconnected by peer
    WhoInitiatedDisconnect_peer WhoInitiatedDisconnect = "peer"
)

