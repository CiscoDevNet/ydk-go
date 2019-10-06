// Cisco XE Native Intermediate System-to-Intermediate System (IS-IS) Yang model.
// Copyright (c) 2016-2017 by Cisco Systems, Inc.
// All rights reserved.
package isis

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package isis"))
}

// AuthenticationLevelType
type AuthenticationLevelType string

const (
    AuthenticationLevelType_level_1 AuthenticationLevelType = "level-1"

    AuthenticationLevelType_level_2 AuthenticationLevelType = "level-2"
)

// IsisLevelType
type IsisLevelType string

const (
    IsisLevelType_level_1 IsisLevelType = "level-1"

    IsisLevelType_level_1_2 IsisLevelType = "level-1-2"

    IsisLevelType_level_2 IsisLevelType = "level-2"
)

// IsisRoutesLevelType
type IsisRoutesLevelType string

const (
    IsisRoutesLevelType_level_1 IsisRoutesLevelType = "level-1"

    IsisRoutesLevelType_level_1_2 IsisRoutesLevelType = "level-1-2"

    IsisRoutesLevelType_level_2 IsisRoutesLevelType = "level-2"
)

