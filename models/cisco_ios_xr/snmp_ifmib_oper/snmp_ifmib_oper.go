// This module contains a collection of YANG definitions
// for Cisco IOS-XR snmp-ifmib package operational data.
// 
// This YANG module augments the
//   Cisco-IOS-XR-snmp-agent-oper
// module with state data.
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package snmp_ifmib_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package snmp_ifmib_oper"))
}

// LinkUpDownStatus represents Link up down status
type LinkUpDownStatus string

const (
    // LinkUpDown notification is enabled
    LinkUpDownStatus_enabled LinkUpDownStatus = "enabled"

    // LinkUpDown notification is disabled
    LinkUpDownStatus_disabled LinkUpDownStatus = "disabled"
)

