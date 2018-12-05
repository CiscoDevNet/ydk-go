// This YANG module augments the 'ietf-routing' module with basic
// configuration and state data for IPv4 unicast routing.
// 
// Copyright (c) 2014 IETF Trust and the persons identified as
// authors of the code. All rights reserved.
// 
// Redistribution and use in source and binary forms, with or
// without modification, is permitted pursuant to, and subject to
// the license terms contained in, the Simplified BSD License set
// forth in Section 4.c of the IETF Trust's Legal Provisions
// Relating to IETF Documents
// (http://trustee.ietf.org/license-info).
// 
// The key words 'MUST', 'MUST NOT', 'REQUIRED', 'SHALL', 'SHALL
// NOT', 'SHOULD', 'SHOULD NOT', 'RECOMMENDED', 'MAY', and
// 'OPTIONAL' in the module text are to be interpreted as described
// in RFC 2119 (http://tools.ietf.org/html/rfc2119).
// 
// This version of this YANG module is part of RFC XXXX
// (http://tools.ietf.org/html/rfcXXXX); see the RFC itself for
// full legal notices.
package ipv4_unicast_routing

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ipv4_unicast_routing"))
}

type Ipv4Unicast struct {
}

func (id Ipv4Unicast) String() string {
	return "ietf-ipv4-unicast-routing:ipv4-unicast"
}

