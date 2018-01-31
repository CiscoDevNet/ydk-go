// This module contains a collection of generally useful
// derived YANG data types.
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package common_acl_datatypes

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package common_acl_datatypes"))
}

// AclUsageAppIdEnum represents Acl usage app id enum
type AclUsageAppIdEnum string

const (
    // General Usage Statistics
    AclUsageAppIdEnum_pfilter AclUsageAppIdEnum = "pfilter"

    // Usage staistics related to BGP Traffic
    AclUsageAppIdEnum_bgp AclUsageAppIdEnum = "bgp"

    // Usage staistics related to OSPF Traffic
    AclUsageAppIdEnum_ospf AclUsageAppIdEnum = "ospf"
)

