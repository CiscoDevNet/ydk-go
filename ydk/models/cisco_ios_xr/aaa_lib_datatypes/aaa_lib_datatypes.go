// This module contains a collection of generally useful
// derived YANG data types.
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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

// AaaMethodAccounting represents Aaa method accounting
type AaaMethodAccounting string

const (
    // Not Set
    AaaMethodAccounting_not_set AaaMethodAccounting = "not-set"

    // None
    AaaMethodAccounting_none AaaMethodAccounting = "none"

    // Radius
    AaaMethodAccounting_radius AaaMethodAccounting = "radius"

    // TACACS+
    AaaMethodAccounting_tacacs_plus AaaMethodAccounting = "tacacs-plus"

    // DSMD
    AaaMethodAccounting_dsmd AaaMethodAccounting = "dsmd"

    // SGBP
    AaaMethodAccounting_sgbp AaaMethodAccounting = "sgbp"

    // AcctD
    AaaMethodAccounting_acct_d AaaMethodAccounting = "acct-d"

    // Error
    AaaMethodAccounting_error_ AaaMethodAccounting = "error"

    // If Authenticated
    AaaMethodAccounting_if_authenticated AaaMethodAccounting = "if-authenticated"

    // Server Group
    AaaMethodAccounting_server_group AaaMethodAccounting = "server-group"

    // Server Group Not Defined
    AaaMethodAccounting_server_group_not_defined AaaMethodAccounting = "server-group-not-defined"

    // Line
    AaaMethodAccounting_line AaaMethodAccounting = "line"

    // Enable
    AaaMethodAccounting_enable AaaMethodAccounting = "enable"

    // Kerberos
    AaaMethodAccounting_kerberos AaaMethodAccounting = "kerberos"

    // Diameter
    AaaMethodAccounting_diameter AaaMethodAccounting = "diameter"

    // Last
    AaaMethodAccounting_last AaaMethodAccounting = "last"

    // Local
    AaaMethodAccounting_local AaaMethodAccounting = "local"
)

// AaaAccountingRpFailover represents Aaa accounting rp failover
type AaaAccountingRpFailover string

const (
    // Disable rpfailover
    AaaAccountingRpFailover_disable AaaAccountingRpFailover = "disable"

    // Enable rpfailover
    AaaAccountingRpFailover_enable AaaAccountingRpFailover = "enable"
)

