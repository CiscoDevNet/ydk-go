// This module contains a collection of YANG definitions
// for Cisco IOS-XR aaa-locald package configuration.
// 
// This YANG module augments the
//   Cisco-IOS-XR-aaa-lib-cfg
// module with configuration data.
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package aaa_locald_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package aaa_locald_cfg"))
}

// AaaLocaldTaskClass represents Aaa locald task class
type AaaLocaldTaskClass string

const (
    // Permits read operation for a Task ID
    AaaLocaldTaskClass_read AaaLocaldTaskClass = "read"

    // Permits write operation for a Task ID
    AaaLocaldTaskClass_write AaaLocaldTaskClass = "write"

    // Permits execute operation for a Task ID
    AaaLocaldTaskClass_execute AaaLocaldTaskClass = "execute"

    // Permits debug operation for a Task ID
    AaaLocaldTaskClass_debug AaaLocaldTaskClass = "debug"
)

