// This module contains a collection of generally useful
// derived YANG data types.
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package aaa_lib_datatypes

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package aaa_lib_datatypes"))
}

// AaaAccountingUpdate represents Aaa accounting update
type AaaAccountingUpdate string

const (
    // Not Set
    AaaAccountingUpdate_none AaaAccountingUpdate = "none"

    // Update records for new accountable information
    // only
    AaaAccountingUpdate_newinfo AaaAccountingUpdate = "newinfo"

    // Update records at periodic intervals
    AaaAccountingUpdate_periodic AaaAccountingUpdate = "periodic"
)

// AaaAccounting represents Aaa accounting
type AaaAccounting string

const (
    // Not Set
    AaaAccounting_not_set AaaAccounting = "not-set"

    // Start Stop
    AaaAccounting_start_stop AaaAccounting = "start-stop"

    // Stop Only
    AaaAccounting_stop_only AaaAccounting = "stop-only"
)

// AaaMethod represents Aaa method
type AaaMethod string

const (
    // Not Set
    AaaMethod_not_set AaaMethod = "not-set"

    // None
    AaaMethod_none AaaMethod = "none"

    // Local
    AaaMethod_local AaaMethod = "local"

    // Radius
    AaaMethod_radius AaaMethod = "radius"

    // TACACS+
    AaaMethod_tacacs_plus AaaMethod = "tacacs-plus"

    // DSMD
    AaaMethod_dsmd AaaMethod = "dsmd"

    // SGBP
    AaaMethod_sgbp AaaMethod = "sgbp"

    // AcctD
    AaaMethod_acct_d AaaMethod = "acct-d"

    // Error
    AaaMethod_error_ AaaMethod = "error"

    // If Authenticated
    AaaMethod_if_authenticated AaaMethod = "if-authenticated"

    // Server Group
    AaaMethod_server_group AaaMethod = "server-group"

    // Server Group Not Defined
    AaaMethod_server_group_not_defined AaaMethod = "server-group-not-defined"

    // Line
    AaaMethod_line AaaMethod = "line"

    // Enable
    AaaMethod_enable AaaMethod = "enable"

    // Kerberos
    AaaMethod_kerberos AaaMethod = "kerberos"

    // Diameter
    AaaMethod_diameter AaaMethod = "diameter"

    // Last
    AaaMethod_last AaaMethod = "last"
)

// AaaAccountingBroadcast represents Aaa accounting broadcast
type AaaAccountingBroadcast string

const (
    // Disable Broadcast
    AaaAccountingBroadcast_disable AaaAccountingBroadcast = "disable"

    // Enable Broadcast
    AaaAccountingBroadcast_enable AaaAccountingBroadcast = "enable"
)

// AaaAccountingRpFailover represents Aaa accounting rp failover
type AaaAccountingRpFailover string

const (
    // Disable rpfailover
    AaaAccountingRpFailover_disable AaaAccountingRpFailover = "disable"

    // Enable rpfailover
    AaaAccountingRpFailover_enable AaaAccountingRpFailover = "enable"
)

