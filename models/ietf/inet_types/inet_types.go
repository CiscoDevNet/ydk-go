// This module contains a collection of generally useful derived
// YANG data types for Internet addresses and related things.
// 
// Copyright (c) 2013 IETF Trust and the persons identified as
// authors of the code.  All rights reserved.
// 
// Redistribution and use in source and binary forms, with or
// without modification, is permitted pursuant to, and subject
// to the license terms contained in, the Simplified BSD License
// set forth in Section 4.c of the IETF Trust's Legal Provisions
// Relating to IETF Documents
// (http://trustee.ietf.org/license-info).
// 
// This version of this YANG module is part of RFC 6991; see
// the RFC itself for full legal notices.
package inet_types

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package inet_types"))
}

// IpVersion represents to the InetVersion textual convention of the SMIv2.
type IpVersion string

const (
    // An unknown or unspecified version of the Internet
    // protocol.
    IpVersion_unknown IpVersion = "unknown"

    // The IPv4 protocol as defined in RFC 791.
    IpVersion_ipv4 IpVersion = "ipv4"

    // The IPv6 protocol as defined in RFC 2460.
    IpVersion_ipv6 IpVersion = "ipv6"
)

