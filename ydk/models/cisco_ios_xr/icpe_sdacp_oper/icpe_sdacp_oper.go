// This module contains a collection of YANG definitions
// for Cisco IOS-XR icpe-sdacp package operational data.
// 
// This YANG module augments the
//   Cisco-IOS-XR-icpe-infra-oper
// module with state data.
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package icpe_sdacp_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package icpe_sdacp_oper"))
}

// DpmProtoHostState represents Dpm proto host state
type DpmProtoHostState string

const (
    // Idle
    DpmProtoHostState_dpm_proto_host_state_idle DpmProtoHostState = "dpm-proto-host-state-idle"

    // Discovered
    DpmProtoHostState_dpm_proto_host_state_discovered DpmProtoHostState = "dpm-proto-host-state-discovered"

    // Rejecting
    DpmProtoHostState_dpm_proto_host_state_rejecting DpmProtoHostState = "dpm-proto-host-state-rejecting"
)

// DpmProtoState represents Dpm proto state
type DpmProtoState string

const (
    // Idle
    DpmProtoState_dpm_proto_state_idle DpmProtoState = "dpm-proto-state-idle"

    // Probing
    DpmProtoState_dpm_proto_state_probing DpmProtoState = "dpm-proto-state-probing"

    // Legacy
    DpmProtoState_dpm_proto_state_legacy DpmProtoState = "dpm-proto-state-legacy"

    // Configuring
    DpmProtoState_dpm_proto_state_configuring DpmProtoState = "dpm-proto-state-configuring"

    // Discovered
    DpmProtoState_dpm_proto_state_discovered DpmProtoState = "dpm-proto-state-discovered"

    // Rejecting
    DpmProtoState_dpm_proto_state_rejecting DpmProtoState = "dpm-proto-state-rejecting"

    // Seen
    DpmProtoState_dpm_proto_state_seen DpmProtoState = "dpm-proto-state-seen"
)

// IcpeCpmChanFsmState represents Icpe cpm chan fsm state
type IcpeCpmChanFsmState string

const (
    // Down
    IcpeCpmChanFsmState_icpe_cpm_chan_fsm_state_down IcpeCpmChanFsmState = "icpe-cpm-chan-fsm-state-down"

    // Not supported
    IcpeCpmChanFsmState_icpe_cpm_chan_fsm_state_not_supported IcpeCpmChanFsmState = "icpe-cpm-chan-fsm-state-not-supported"

    // Closed
    IcpeCpmChanFsmState_icpe_cpm_chan_fsm_state_closed IcpeCpmChanFsmState = "icpe-cpm-chan-fsm-state-closed"

    // Opening
    IcpeCpmChanFsmState_icpe_cpm_chan_fsm_state_opening IcpeCpmChanFsmState = "icpe-cpm-chan-fsm-state-opening"

    // Opened
    IcpeCpmChanFsmState_icpe_cpm_chan_fsm_state_opened IcpeCpmChanFsmState = "icpe-cpm-chan-fsm-state-opened"

    // Open
    IcpeCpmChanFsmState_icpe_cpm_chan_fsm_state_open IcpeCpmChanFsmState = "icpe-cpm-chan-fsm-state-open"
)

// IcpeCpmChannelResyncState represents Icpe cpm channel resync state
type IcpeCpmChannelResyncState string

const (
    // Unknown
    IcpeCpmChannelResyncState_icpe_cpm_channel_resync_state_unknown IcpeCpmChannelResyncState = "icpe-cpm-channel-resync-state-unknown"

    // Not in resync
    IcpeCpmChannelResyncState_icpe_cpm_channel_resync_state_not_in_resync IcpeCpmChannelResyncState = "icpe-cpm-channel-resync-state-not-in-resync"

    // In client resync
    IcpeCpmChannelResyncState_icpe_cpm_channel_resync_state_in_client_resync IcpeCpmChannelResyncState = "icpe-cpm-channel-resync-state-in-client-resync"

    // In satellite resync
    IcpeCpmChannelResyncState_icpe_cpm_channel_resync_state_in_satellite_resync IcpeCpmChannelResyncState = "icpe-cpm-channel-resync-state-in-satellite-resync"
)

// IcpeCpmControlFsmState represents Icpe cpm control fsm state
type IcpeCpmControlFsmState string

const (
    // Disconnected
    IcpeCpmControlFsmState_icpe_cpm_control_fsm_state_disconnected IcpeCpmControlFsmState = "icpe-cpm-control-fsm-state-disconnected"

    // Connecting
    IcpeCpmControlFsmState_icpe_cpm_control_fsm_state_connecting IcpeCpmControlFsmState = "icpe-cpm-control-fsm-state-connecting"

    // Authenticating
    IcpeCpmControlFsmState_icpe_cpm_control_fsm_state_authenticating IcpeCpmControlFsmState = "icpe-cpm-control-fsm-state-authenticating"

    // Checking Version
    IcpeCpmControlFsmState_icpe_cpm_control_fsm_state_check_ing_ver IcpeCpmControlFsmState = "icpe-cpm-control-fsm-state-check-ing-ver"

    // Connected
    IcpeCpmControlFsmState_icpe_cpm_control_fsm_state_connected IcpeCpmControlFsmState = "icpe-cpm-control-fsm-state-connected"

    // ISSU In Progress
    IcpeCpmControlFsmState_icpe_cpm_control_fsm_state_issu IcpeCpmControlFsmState = "icpe-cpm-control-fsm-state-issu"
)

