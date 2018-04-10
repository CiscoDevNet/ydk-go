// This module contains a collection of generally useful
// derived YANG data types.
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package tty_management_datatypes

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package tty_management_datatypes"))
}

// TtyPager represents Tty pager
type TtyPager string

const (
    // More paging Utility
    TtyPager_more TtyPager = "more"

    // Less paging Utility
    TtyPager_less TtyPager = "less"

    // No Paging Utility
    TtyPager_none TtyPager = "none"
)

// TtyEscapeChar represents Tty escape char
type TtyEscapeChar string

const (
    // Cause escape on BREAK
    TtyEscapeChar_break_ TtyEscapeChar = "break"

    // Use default escape character
    TtyEscapeChar_default_ TtyEscapeChar = "default"

    // Disable escape entirely
    TtyEscapeChar_none TtyEscapeChar = "none"
)

// TtyTransportProtocolSelect represents Tty transport protocol select
type TtyTransportProtocolSelect string

const (
    // No protocols
    TtyTransportProtocolSelect_none TtyTransportProtocolSelect = "none"

    // All protocols
    TtyTransportProtocolSelect_all TtyTransportProtocolSelect = "all"

    // Some Protocol
    TtyTransportProtocolSelect_some TtyTransportProtocolSelect = "some"
)

// TtySessionTimeoutDirection represents Tty session timeout direction
type TtySessionTimeoutDirection string

const (
    // Input traffic
    TtySessionTimeoutDirection_in TtySessionTimeoutDirection = "in"

    // In & Output traffic
    TtySessionTimeoutDirection_in_out TtySessionTimeoutDirection = "in-out"
)

// TtyTransportProtocol represents Tty transport protocol
type TtyTransportProtocol string

const (
    // No protocols
    TtyTransportProtocol_none TtyTransportProtocol = "none"

    // TCP/IP Telnet protocol
    TtyTransportProtocol_telnet TtyTransportProtocol = "telnet"

    // Unix ssh protocol
    TtyTransportProtocol_ssh TtyTransportProtocol = "ssh"
)

