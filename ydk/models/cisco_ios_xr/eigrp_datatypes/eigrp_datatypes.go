// This module contains a collection of generally useful
// derived YANG data types.
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package eigrp_datatypes

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package eigrp_datatypes"))
}

// EigrpAf represents Eigrp af
type EigrpAf string

const (
    // IPv4 address family
    EigrpAf_ipv4 EigrpAf = "ipv4"

    // IPv6 address family
    EigrpAf_ipv6 EigrpAf = "ipv6"
)

