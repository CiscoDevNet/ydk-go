// This module contains definitions
// for the Calvados model objects.
// 
// Copyright (c) 2012-2018 by Cisco Systems, Inc.
// All rights reserved.
package inet_address_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package inet_address_mib"))
}

// InetAddressType
type InetAddressType string

const (
    InetAddressType_unknown InetAddressType = "unknown"

    InetAddressType_ipv4 InetAddressType = "ipv4"

    InetAddressType_ipv6 InetAddressType = "ipv6"

    InetAddressType_ipv4z InetAddressType = "ipv4z"

    InetAddressType_ipv6z InetAddressType = "ipv6z"

    InetAddressType_dns InetAddressType = "dns"
)

// InetScopeType
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

// InetVersion
type InetVersion string

const (
    InetVersion_unknown InetVersion = "unknown"

    InetVersion_ipv4 InetVersion = "ipv4"

    InetVersion_ipv6 InetVersion = "ipv6"
)

