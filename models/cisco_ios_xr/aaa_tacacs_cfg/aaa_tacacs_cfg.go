// This module contains a collection of YANG definitions
// for Cisco IOS-XR aaa-tacacs package configuration.
// 
// This YANG module augments the
//   Cisco-IOS-XR-aaa-lib-cfg,
//   Cisco-IOS-XR-aaa-locald-cfg
// modules with configuration data.
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package aaa_tacacs_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package aaa_tacacs_cfg"))
}

// TacacsDscpValue represents Tacacs dscp value
type TacacsDscpValue string

const (
    // Match packets with AF11 DSCP
    TacacsDscpValue_af11 TacacsDscpValue = "af11"

    // Match packets with AF12 DSCP
    TacacsDscpValue_af12 TacacsDscpValue = "af12"

    // Match packets with AF13 DSCP
    TacacsDscpValue_af13 TacacsDscpValue = "af13"

    // Match packets with AF21 DSCP
    TacacsDscpValue_af21 TacacsDscpValue = "af21"

    // Match packets with AF22 DSCP
    TacacsDscpValue_af22 TacacsDscpValue = "af22"

    // Match packets with AF23 DSCP
    TacacsDscpValue_af23 TacacsDscpValue = "af23"

    // Match packets with AF31 DSCP
    TacacsDscpValue_af31 TacacsDscpValue = "af31"

    // Match packets with AF32 DSCP
    TacacsDscpValue_af32 TacacsDscpValue = "af32"

    // Match packets with AF33 DSCP
    TacacsDscpValue_af33 TacacsDscpValue = "af33"

    // Match packets with AF41 DSCP
    TacacsDscpValue_af41 TacacsDscpValue = "af41"

    // Match packets with AF42 DSCP
    TacacsDscpValue_af42 TacacsDscpValue = "af42"

    // Match packets with AF43 DSCP
    TacacsDscpValue_af43 TacacsDscpValue = "af43"

    // Match packets with CS1 DSCP
    TacacsDscpValue_cs1 TacacsDscpValue = "cs1"

    // Match packets with CS2 DSCP
    TacacsDscpValue_cs2 TacacsDscpValue = "cs2"

    // Match packets with CS3 DSCP
    TacacsDscpValue_cs3 TacacsDscpValue = "cs3"

    // Match packets with CS4 DSCP
    TacacsDscpValue_cs4 TacacsDscpValue = "cs4"

    // Match packets with CS5 DSCP
    TacacsDscpValue_cs5 TacacsDscpValue = "cs5"

    // Match packets with CS6 DSCP
    TacacsDscpValue_cs6 TacacsDscpValue = "cs6"

    // Match packets with CS7 DSCP
    TacacsDscpValue_cs7 TacacsDscpValue = "cs7"

    // Match packets with 0000 DSCP
    TacacsDscpValue_default_ TacacsDscpValue = "default"

    // Match packets with EF DSCP
    TacacsDscpValue_ef TacacsDscpValue = "ef"
)

