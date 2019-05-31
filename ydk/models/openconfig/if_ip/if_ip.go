// Model for managing IP interfaces.
// 
// This model reuses most of the IETF YANG model for IP management
// described by RFC 7277.  The primary differences are in the
// structure of configuration and state data.
package if_ip

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package if_ip"))
}

// IpAddressOrigin represents The origin of an address.
type IpAddressOrigin string

const (
    // None of the following.
    IpAddressOrigin_OTHER IpAddressOrigin = "OTHER"

    // Indicates that the address has been statically
    // configured - for example, using NETCONF or a Command Line
    // Interface.
    IpAddressOrigin_STATIC IpAddressOrigin = "STATIC"

    // Indicates an address that has been assigned to this
    // system by a DHCP server.
    IpAddressOrigin_DHCP IpAddressOrigin = "DHCP"

    // Indicates an address created by IPv6 stateless
    // autoconfiguration that embeds a link-layer address in its
    // interface identifier.
    IpAddressOrigin_LINK_LAYER IpAddressOrigin = "LINK_LAYER"

    // [adapted from RFC 7277]
    // 
    // Indicates an address chosen by the system at
    // random, e.g., an IPv4 address within 169.254/16, an
    // RFC 4941 temporary address, or an RFC 7217 semantically
    // opaque address.
    IpAddressOrigin_RANDOM IpAddressOrigin = "RANDOM"
)

// NeighborOrigin represents The origin of a neighbor entry.
type NeighborOrigin string

const (
    // None of the following.
    NeighborOrigin_OTHER NeighborOrigin = "OTHER"

    // Indicates that the mapping has been statically
    // configured - for example, using NETCONF or a Command Line
    // Interface.
    NeighborOrigin_STATIC NeighborOrigin = "STATIC"

    // [adapted from RFC 7277]
    // 
    // Indicates that the mapping has been dynamically resolved
    // using, e.g., IPv4 ARP or the IPv6 Neighbor Discovery
    // protocol.
    NeighborOrigin_DYNAMIC NeighborOrigin = "DYNAMIC"
)

