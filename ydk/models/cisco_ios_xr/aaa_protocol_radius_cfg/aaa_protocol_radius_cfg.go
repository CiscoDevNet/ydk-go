// This module contains a collection of YANG definitions
// for Cisco IOS-XR aaa-protocol-radius package configuration.
// 
// This YANG module augments the
//   Cisco-IOS-XR-aaa-locald-cfg,
//   Cisco-IOS-XR-aaa-lib-cfg
// modules with configuration data.
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package aaa_protocol_radius_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package aaa_protocol_radius_cfg"))
}

// AaaDscpValue represents Aaa dscp value
type AaaDscpValue string

const (
    // Match packets with AF11 DSCP
    AaaDscpValue_af11 AaaDscpValue = "af11"

    // Match packets with AF12 DSCP
    AaaDscpValue_af12 AaaDscpValue = "af12"

    // Match packets with AF13 DSCP
    AaaDscpValue_af13 AaaDscpValue = "af13"

    // Match packets with AF21 DSCP
    AaaDscpValue_af21 AaaDscpValue = "af21"

    // Match packets with AF22 DSCP
    AaaDscpValue_af22 AaaDscpValue = "af22"

    // Match packets with AF23 DSCP
    AaaDscpValue_af23 AaaDscpValue = "af23"

    // Match packets with AF31 DSCP
    AaaDscpValue_af31 AaaDscpValue = "af31"

    // Match packets with AF32 DSCP
    AaaDscpValue_af32 AaaDscpValue = "af32"

    // Match packets with AF33 DSCP
    AaaDscpValue_af33 AaaDscpValue = "af33"

    // Match packets with AF41 DSCP
    AaaDscpValue_af41 AaaDscpValue = "af41"

    // Match packets with AF42 DSCP
    AaaDscpValue_af42 AaaDscpValue = "af42"

    // Match packets with AF43 DSCP
    AaaDscpValue_af43 AaaDscpValue = "af43"

    // Match packets with CS1 DSCP
    AaaDscpValue_cs1 AaaDscpValue = "cs1"

    // Match packets with CS2 DSCP
    AaaDscpValue_cs2 AaaDscpValue = "cs2"

    // Match packets with CS3 DSCP
    AaaDscpValue_cs3 AaaDscpValue = "cs3"

    // Match packets with CS4 DSCP
    AaaDscpValue_cs4 AaaDscpValue = "cs4"

    // Match packets with CS5 DSCP
    AaaDscpValue_cs5 AaaDscpValue = "cs5"

    // Match packets with CS6 DSCP
    AaaDscpValue_cs6 AaaDscpValue = "cs6"

    // Match packets with CS7 DSCP
    AaaDscpValue_cs7 AaaDscpValue = "cs7"

    // Match packets with 0000 DSCP
    AaaDscpValue_default_ AaaDscpValue = "default"

    // Match packets with EF DSCP
    AaaDscpValue_ef AaaDscpValue = "ef"
)

// AaaAction represents Aaa action
type AaaAction string

const (
    // Accept
    AaaAction_accept AaaAction = "accept"

    // Reject
    AaaAction_reject AaaAction = "reject"
)

// AaaAuthentication represents Aaa authentication
type AaaAuthentication string

const (
    // All
    AaaAuthentication_all AaaAuthentication = "all"

    // Any
    AaaAuthentication_any AaaAuthentication = "any"

    // Session key
    AaaAuthentication_session_key AaaAuthentication = "session-key"
)

// AaaSelectKey represents Aaa select key
type AaaSelectKey string

const (
    // Server key
    AaaSelectKey_server_key AaaSelectKey = "server-key"

    // Session  key
    AaaSelectKey_session_key AaaSelectKey = "session-key"
)

// AaaDirection represents Aaa direction
type AaaDirection string

const (
    // Inbound
    AaaDirection_inbound AaaDirection = "inbound"

    // Outbound
    AaaDirection_outbound AaaDirection = "outbound"
)

// AaaConfig represents Aaa config
type AaaConfig string

const (
    // False
    AaaConfig_false AaaConfig = "false"

    // True
    AaaConfig_true AaaConfig = "true"
)

