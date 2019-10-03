// This module contains a collection of YANG definitions
// for Cisco IOS-XR ipv4-ma package operational data.
// 
// This YANG module augments the
//   Cisco-IOS-XR-ipv4-io-oper
// module with state data.
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package ipv4_ma_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ipv4_ma_oper"))
}

// ImStateEnum represents Im state enum
type ImStateEnum string

const (
    // im state not ready
    ImStateEnum_im_state_not_ready ImStateEnum = "im-state-not-ready"

    // im state admin down
    ImStateEnum_im_state_admin_down ImStateEnum = "im-state-admin-down"

    // im state down
    ImStateEnum_im_state_down ImStateEnum = "im-state-down"

    // im state up
    ImStateEnum_im_state_up ImStateEnum = "im-state-up"

    // im state shutdown
    ImStateEnum_im_state_shutdown ImStateEnum = "im-state-shutdown"

    // im state err disable
    ImStateEnum_im_state_err_disable ImStateEnum = "im-state-err-disable"

    // im state down immediate
    ImStateEnum_im_state_down_immediate ImStateEnum = "im-state-down-immediate"

    // im state down immediate admin
    ImStateEnum_im_state_down_immediate_admin ImStateEnum = "im-state-down-immediate-admin"

    // im state down graceful
    ImStateEnum_im_state_down_graceful ImStateEnum = "im-state-down-graceful"

    // im state begin shutdown
    ImStateEnum_im_state_begin_shutdown ImStateEnum = "im-state-begin-shutdown"

    // im state end shutdown
    ImStateEnum_im_state_end_shutdown ImStateEnum = "im-state-end-shutdown"

    // im state begin error disable
    ImStateEnum_im_state_begin_error_disable ImStateEnum = "im-state-begin-error-disable"

    // im state end error disable
    ImStateEnum_im_state_end_error_disable ImStateEnum = "im-state-end-error-disable"

    // im state begin down graceful
    ImStateEnum_im_state_begin_down_graceful ImStateEnum = "im-state-begin-down-graceful"

    // im state reset
    ImStateEnum_im_state_reset ImStateEnum = "im-state-reset"

    // im state operational
    ImStateEnum_im_state_operational ImStateEnum = "im-state-operational"

    // im state not operational
    ImStateEnum_im_state_not_operational ImStateEnum = "im-state-not-operational"

    // im state unknown
    ImStateEnum_im_state_unknown ImStateEnum = "im-state-unknown"

    // im state last
    ImStateEnum_im_state_last ImStateEnum = "im-state-last"
)

// Ipv4MaOperConfig represents ipv4 client type
type Ipv4MaOperConfig string

const (
    // ipv4 ma oper client none
    Ipv4MaOperConfig_ipv4_ma_oper_client_none Ipv4MaOperConfig = "ipv4-ma-oper-client-none"

    // ipv4 ma oper non oc client
    Ipv4MaOperConfig_ipv4_ma_oper_non_oc_client Ipv4MaOperConfig = "ipv4-ma-oper-non-oc-client"

    // ipv4 ma oper oc client
    Ipv4MaOperConfig_ipv4_ma_oper_oc_client Ipv4MaOperConfig = "ipv4-ma-oper-oc-client"
)

// RpfMode represents Interface line states
type RpfMode string

const (
    // Strict RPF
    RpfMode_strict RpfMode = "strict"

    // Loose RPF
    RpfMode_loose RpfMode = "loose"
)

// Ipv4MaOperLineState represents Interface line states
type Ipv4MaOperLineState string

const (
    // Interface state is unknown
    Ipv4MaOperLineState_unknown Ipv4MaOperLineState = "unknown"

    // Interface has been shutdown
    Ipv4MaOperLineState_shutdown Ipv4MaOperLineState = "shutdown"

    // Interface state is down
    Ipv4MaOperLineState_down Ipv4MaOperLineState = "down"

    // Interface state is up
    Ipv4MaOperLineState_up Ipv4MaOperLineState = "up"
)

