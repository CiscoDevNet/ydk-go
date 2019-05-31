// This module contains a collection of generally useful
// derived YANG data types.
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package segment_routing_srv6_datatypes

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package segment_routing_srv6_datatypes"))
}

// Srv6EncapsulationHopLimitOption represents Srv6 encapsulation hop limit option
type Srv6EncapsulationHopLimitOption string

const (
    // Set Value
    Srv6EncapsulationHopLimitOption_count Srv6EncapsulationHopLimitOption = "count"
)

