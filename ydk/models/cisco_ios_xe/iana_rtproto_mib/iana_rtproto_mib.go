// This MIB module defines the IANAipRouteProtocol and
// IANAipMRouteProtocol textual conventions for use in MIBs
// which need to identify unicast or multicast routing
// mechanisms.
// 
// Any additions or changes to the contents of this MIB module
// require either publication of an RFC, or Designated Expert
// Review as defined in RFC 2434, Guidelines for Writing an
// IANA Considerations Section in RFCs.  The Designated Expert 
// will be selected by the IESG Area Director(s) of the Routing
// Area.
package iana_rtproto_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package iana_rtproto_mib"))
}

// IANAipRouteProtocol represents protocols need be supported.
type IANAipRouteProtocol string

const (
    IANAipRouteProtocol_other IANAipRouteProtocol = "other"

    IANAipRouteProtocol_local IANAipRouteProtocol = "local"

    IANAipRouteProtocol_netmgmt IANAipRouteProtocol = "netmgmt"

    IANAipRouteProtocol_icmp IANAipRouteProtocol = "icmp"

    IANAipRouteProtocol_egp IANAipRouteProtocol = "egp"

    IANAipRouteProtocol_ggp IANAipRouteProtocol = "ggp"

    IANAipRouteProtocol_hello IANAipRouteProtocol = "hello"

    IANAipRouteProtocol_rip IANAipRouteProtocol = "rip"

    IANAipRouteProtocol_isIs IANAipRouteProtocol = "isIs"

    IANAipRouteProtocol_esIs IANAipRouteProtocol = "esIs"

    IANAipRouteProtocol_ciscoIgrp IANAipRouteProtocol = "ciscoIgrp"

    IANAipRouteProtocol_bbnSpfIgp IANAipRouteProtocol = "bbnSpfIgp"

    IANAipRouteProtocol_ospf IANAipRouteProtocol = "ospf"

    IANAipRouteProtocol_bgp IANAipRouteProtocol = "bgp"

    IANAipRouteProtocol_idpr IANAipRouteProtocol = "idpr"

    IANAipRouteProtocol_ciscoEigrp IANAipRouteProtocol = "ciscoEigrp"

    IANAipRouteProtocol_dvmrp IANAipRouteProtocol = "dvmrp"
)

// IANAipMRouteProtocol represents those protocols need be supported.
type IANAipMRouteProtocol string

const (
    IANAipMRouteProtocol_other IANAipMRouteProtocol = "other"

    IANAipMRouteProtocol_local IANAipMRouteProtocol = "local"

    IANAipMRouteProtocol_netmgmt IANAipMRouteProtocol = "netmgmt"

    IANAipMRouteProtocol_dvmrp IANAipMRouteProtocol = "dvmrp"

    IANAipMRouteProtocol_mospf IANAipMRouteProtocol = "mospf"

    IANAipMRouteProtocol_pimSparseDense IANAipMRouteProtocol = "pimSparseDense"

    IANAipMRouteProtocol_cbt IANAipMRouteProtocol = "cbt"

    IANAipMRouteProtocol_pimSparseMode IANAipMRouteProtocol = "pimSparseMode"

    IANAipMRouteProtocol_pimDenseMode IANAipMRouteProtocol = "pimDenseMode"

    IANAipMRouteProtocol_igmpOnly IANAipMRouteProtocol = "igmpOnly"

    IANAipMRouteProtocol_bgmp IANAipMRouteProtocol = "bgmp"

    IANAipMRouteProtocol_msdp IANAipMRouteProtocol = "msdp"
)

