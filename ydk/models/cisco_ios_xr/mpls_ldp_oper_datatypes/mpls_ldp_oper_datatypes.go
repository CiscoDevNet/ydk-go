// This module contains a collection of YANG definitions
// for Cisco IOS-XR mpls-ldp-oper-data package operational data.
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package mpls_ldp_oper_datatypes

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package mpls_ldp_oper_datatypes"))
}

// MplsLdpOperAfName represents Mpls ldp oper af name
type MplsLdpOperAfName string

const (
    // IPv4
    MplsLdpOperAfName_ipv4 MplsLdpOperAfName = "ipv4"

    // IPv6
    MplsLdpOperAfName_ipv6 MplsLdpOperAfName = "ipv6"
)

