// This module contains a collection of YANG definitions
// for Cisco IOS-XR tty-management package operational data.
// 
// This YANG module augments the
//   Cisco-IOS-XR-tty-server-oper
// module with state data.
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package tty_management_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package tty_management_oper"))
}

type Ipv4 struct {
}

func (id Ipv4) String() string {
	return "Cisco-IOS-XR-tty-management-oper-sub1:ipv4"
}

type HostAfIdBase struct {
}

func (id HostAfIdBase) String() string {
	return "Cisco-IOS-XR-tty-management-oper-sub1:Host-af-id-base"
}

type Ipv6 struct {
}

func (id Ipv6) String() string {
	return "Cisco-IOS-XR-tty-management-oper-sub1:ipv6"
}

// TransportService represents Transport service protocol
type TransportService string

const (
    // Unknown service
    TransportService_unknown TransportService = "unknown"

    // Telnet
    TransportService_telnet TransportService = "telnet"

    // Remote login
    TransportService_rlogin TransportService = "rlogin"

    // SSH
    TransportService_ssh TransportService = "ssh"
)

