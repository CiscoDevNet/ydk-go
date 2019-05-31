// This module contains a collection of YANG definitions for
// configuring diffserv specification implementations.
// 
// Copyright (c) 2014 IETF Trust and the persons identified as
// authors of the code.  All rights reserved.
// 
// Redistribution and use in source and binary forms, with or
// without modification, is permitted pursuant to, and subject
// to the license terms contained in, the Simplified BSD License
// set forth in Section 4.c of the IETF Trust's Legal Provisions
// Relating to IETF Documents
// (http://trustee.ietf.org/license-info).
// 
// This version of this YANG module is part of RFC XXXX; see
// the RFC itself for full legal notices.
package diffserv_target

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package diffserv_target"))
}

type Direction struct {
}

func (id Direction) String() string {
	return "ietf-diffserv-target:direction"
}

type Inbound struct {
}

func (id Inbound) String() string {
	return "ietf-diffserv-target:inbound"
}

type Outbound struct {
}

func (id Outbound) String() string {
	return "ietf-diffserv-target:outbound"
}

