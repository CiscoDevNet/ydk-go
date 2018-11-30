// This module contains a collection of YANG definitions for
// configuring IP implementations.
// Copyright (c) 2014 IETF Trust and the persons identified as
// authors of the code.  All rights reserved.
// Redistribution and use in source and binary forms, with or
// without modification, is permitted pursuant to, and subject
// to the license terms contained in, the Simplified BSD License
// set forth in Section 4.c of the IETF Trust's Legal Provisions
// Relating to IETF Documents
// (http://trustee.ietf.org/license-info).
// This version of this YANG module is part of RFC 7277; see
// the RFC itself for full legal notices.
package ip

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ip"))
}

// NeighborOrigin represents The origin of a neighbor entry.
type NeighborOrigin string

const (
    // None of the following.
    NeighborOrigin_other NeighborOrigin = "other"

    // Indicates that the mapping has been statically
    // configured - for example, using NETCONF or a Command Line
    // Interface.
    NeighborOrigin_static NeighborOrigin = "static"

    // Indicates that the mapping has been dynamically resolved
    // using, e.g., IPv4 ARP or the IPv6 Neighbor Discovery
    // protocol.
    NeighborOrigin_dynamic NeighborOrigin = "dynamic"
)

// IpAddressOrigin represents The origin of an address.
type IpAddressOrigin string

const (
    // None of the following.
    IpAddressOrigin_other IpAddressOrigin = "other"

    // Indicates that the address has been statically
    // configured - for example, using NETCONF or a Command Line
    // Interface.
    IpAddressOrigin_static IpAddressOrigin = "static"

    // Indicates an address that has been assigned to this
    // system by a DHCP server.
    IpAddressOrigin_dhcp IpAddressOrigin = "dhcp"

    // Indicates an address created by IPv6 stateless
    // autoconfiguration that embeds a link-layer address in its
    // interface identifier.
    IpAddressOrigin_link_layer IpAddressOrigin = "link-layer"

    // Indicates an address chosen by the system at
    // random, e.g., an IPv4 address within 169.254/16, an
    // RFC 4941 temporary address, or an RFC 7217 semantically
    // opaque address.
    IpAddressOrigin_random IpAddressOrigin = "random"
)

