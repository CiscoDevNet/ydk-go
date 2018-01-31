// Cisco XE Native Embedded Event Manager (EEM) Yang model.
// Copyright (c) 2016 by Cisco Systems, Inc.
// All rights reserved.
package eem

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package eem"))
}

// OperatorType
type OperatorType string

const (
    OperatorType_eq OperatorType = "eq"

    OperatorType_ge OperatorType = "ge"

    OperatorType_gt OperatorType = "gt"

    OperatorType_le OperatorType = "le"

    OperatorType_lt OperatorType = "lt"

    OperatorType_ne OperatorType = "ne"
)

