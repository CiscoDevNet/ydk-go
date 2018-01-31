// This module contains a collection of generally useful
// derived YANG data types.
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ipv4_autorp_datatypes

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ipv4_autorp_datatypes"))
}

// AutoRpProtocolMode represents Auto rp protocol mode
type AutoRpProtocolMode string

const (
    // Sparse Mode
    AutoRpProtocolMode_sparse AutoRpProtocolMode = "sparse"

    // Bidirectional Mode
    AutoRpProtocolMode_bidirectional AutoRpProtocolMode = "bidirectional"
)

