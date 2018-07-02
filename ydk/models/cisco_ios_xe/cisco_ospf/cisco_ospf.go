// This YANG module defines the Cisco OSPF common 
// yang model
// Copyright (c) 2016 by Cisco Systems, Inc.
// All rights reserved.
package cisco_ospf

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_ospf"))
}

// OspfLogAdj represents Ospf log adjacency changes
type OspfLogAdj string

const (
    // Enable log adjacency brief
    OspfLogAdj_enable OspfLogAdj = "enable"

    // Enable log adjcency detail
    OspfLogAdj_detail OspfLogAdj = "detail"

    // disable log adjacency
    OspfLogAdj_disable OspfLogAdj = "disable"
)

// PrefixApplicability represents Ospf uloop avoidance
type PrefixApplicability string

const (
    // Protected prefixes only
    PrefixApplicability_protected PrefixApplicability = "protected"

    // All prefixes
    PrefixApplicability_all PrefixApplicability = "all"
)

// AccessListInOutType represents Access list in and out
type AccessListInOutType string

const (
    // Filter incoming routing updates
    AccessListInOutType_in AccessListInOutType = "in"

    // Filter outgoing routing updates
    AccessListInOutType_out AccessListInOutType = "out"
)

// OspfExternalType represents external route types
type OspfExternalType string

const (
    // type 1
    OspfExternalType_Y_1 OspfExternalType = "1"

    // type 2
    OspfExternalType_Y_2 OspfExternalType = "2"
)

