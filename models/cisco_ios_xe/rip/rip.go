// Cisco XE Native Routing Information Protocol (RIP) Yang model.
// Copyright (c) 2016-2017 by Cisco Systems, Inc.
// All rights reserved.
package rip

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package rip"))
}

// OffsetListInOutType
type OffsetListInOutType string

const (
    OffsetListInOutType_in OffsetListInOutType = "in"

    OffsetListInOutType_out OffsetListInOutType = "out"
)

