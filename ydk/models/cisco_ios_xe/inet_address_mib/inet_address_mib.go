// This MIB module defines textual conventions for
// representing Internet addresses.  An Internet
// address can be an IPv4 address, an IPv6 address,
// or a DNS domain name.  This module also defines
// textual conventions for Internet port numbers,
// autonomous system numbers, and the length of an
// Internet address prefix.
// 
// Copyright (C) The Internet Society (2005).  This version
// of this MIB module is part of RFC 4001, see the RFC
// itself for full legal notices.
package inet_address_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package inet_address_mib"))
}

// InetAddressType represents ipv6(2) to ipv4(1)).
type InetAddressType string

const (
    InetAddressType_unknown InetAddressType = "unknown"

    InetAddressType_ipv4 InetAddressType = "ipv4"

    InetAddressType_ipv6 InetAddressType = "ipv6"

    InetAddressType_ipv4z InetAddressType = "ipv4z"

    InetAddressType_ipv6z InetAddressType = "ipv6z"

    InetAddressType_dns InetAddressType = "dns"
)

// InetScopeType represents not yet assigned.
type InetScopeType string

const (
    InetScopeType_interfaceLocal InetScopeType = "interfaceLocal"

    InetScopeType_linkLocal InetScopeType = "linkLocal"

    InetScopeType_subnetLocal InetScopeType = "subnetLocal"

    InetScopeType_adminLocal InetScopeType = "adminLocal"

    InetScopeType_siteLocal InetScopeType = "siteLocal"

    InetScopeType_organizationLocal InetScopeType = "organizationLocal"

    InetScopeType_global InetScopeType = "global"
)

// InetVersion represents purpose.
type InetVersion string

const (
    InetVersion_unknown InetVersion = "unknown"

    InetVersion_ipv4 InetVersion = "ipv4"

    InetVersion_ipv6 InetVersion = "ipv6"
)

