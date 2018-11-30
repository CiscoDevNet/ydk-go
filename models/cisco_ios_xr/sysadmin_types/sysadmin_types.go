// This module contains definitions
// for the Calvados model objects.
// 
// This module contains a collection of YANG
// definitions for Cisco IOS-XR SysAdmin configuration.
// 
// This module contains a collection of derived 
// YANG data types specific to Sysadmin.
// 
// Copyright(c) 2011-2016 by Cisco Systems, Inc.
// All rights reserved.
// 
// Copyright (c) 2012-2018 by Cisco Systems, Inc.
// All rights reserved.
package sysadmin_types

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package sysadmin_types"))
}

// GenericOperStatus
type GenericOperStatus string

const (
    GenericOperStatus_up GenericOperStatus = "up"

    GenericOperStatus_down GenericOperStatus = "down"
)

// GenericHaRole
type GenericHaRole string

const (
    GenericHaRole_no_ha_role GenericHaRole = "no-ha-role"

    GenericHaRole_Active GenericHaRole = "Active"

    GenericHaRole_Standby GenericHaRole = "Standby"
)

// Adminstate
type Adminstate string

const (
    Adminstate_disable Adminstate = "disable"
)

// RackId represents BSC racks are identified by numbers 128..129
type RackId string

const (
    RackId_L0 RackId = "L0"

    RackId_L1 RackId = "L1"

    RackId_L2 RackId = "L2"

    RackId_L3 RackId = "L3"

    RackId_L4 RackId = "L4"

    RackId_L5 RackId = "L5"

    RackId_L6 RackId = "L6"

    RackId_L7 RackId = "L7"

    RackId_L8 RackId = "L8"

    RackId_L9 RackId = "L9"

    RackId_L10 RackId = "L10"

    RackId_L11 RackId = "L11"

    RackId_L12 RackId = "L12"

    RackId_L13 RackId = "L13"

    RackId_L14 RackId = "L14"

    RackId_L15 RackId = "L15"

    RackId_F0 RackId = "F0"

    RackId_F1 RackId = "F1"

    RackId_F2 RackId = "F2"

    RackId_F3 RackId = "F3"

    RackId_B0 RackId = "B0"

    RackId_B1 RackId = "B1"
)

// FabricLinkTypes
type FabricLinkTypes string

const (
    FabricLinkTypes_S1 FabricLinkTypes = "S1"

    FabricLinkTypes_S2 FabricLinkTypes = "S2"

    FabricLinkTypes_S3 FabricLinkTypes = "S3"
)

// GenericOperStatus_
type GenericOperStatus_ string

const (
    GenericOperStatus__up GenericOperStatus_ = "up"

    GenericOperStatus__down GenericOperStatus_ = "down"
)

// GenericOperStatus_
type GenericOperStatus_ string

const (
    GenericOperStatus__up GenericOperStatus_ = "up"

    GenericOperStatus__down GenericOperStatus_ = "down"
)

