// Cisco XE Native Open Shortest Path First (OSPF) Yang model.
// Copyright (c) 2016-2017 by Cisco Systems, Inc.
// All rights reserved.
package ospf

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ospf"))
}

// RedistOspfExternalType
type RedistOspfExternalType string

const (
    RedistOspfExternalType_Y_1 RedistOspfExternalType = "1"

    RedistOspfExternalType_Y_2 RedistOspfExternalType = "2"
)

