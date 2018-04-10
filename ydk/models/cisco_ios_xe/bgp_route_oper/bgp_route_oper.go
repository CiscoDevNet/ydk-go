// This module contains a collection of YANG definitions
// route for bgp route operational data.
// Copyright (c) 2016-2017 by Cisco Systems, Inc.
// All rights reserved.
package bgp_route_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package bgp_route_oper"))
}

// BgpOriginCode represents BGP origin code
type BgpOriginCode string

const (
    // BGP origin code IGP
    BgpOriginCode_origin_igp BgpOriginCode = "origin-igp"

    // BGP origin code EGP
    BgpOriginCode_origin_egp BgpOriginCode = "origin-egp"

    // BGP origin code incomplete
    BgpOriginCode_origin_incomplete BgpOriginCode = "origin-incomplete"
)

// BgpRpkiStatus represents BGP RPKI status
type BgpRpkiStatus string

const (
    BgpRpkiStatus_rpki_valid BgpRpkiStatus = "rpki-valid"

    BgpRpkiStatus_rpki_invalid BgpRpkiStatus = "rpki-invalid"

    BgpRpkiStatus_rpki_not_found BgpRpkiStatus = "rpki-not-found"

    BgpRpkiStatus_rpki_not_enabled BgpRpkiStatus = "rpki-not-enabled"

    BgpRpkiStatus_rpki_illegal BgpRpkiStatus = "rpki-illegal"
)

// BgpRouteFilters represents BGP route filters
type BgpRouteFilters string

const (
    // BGP all route filter
    BgpRouteFilters_bgp_rf_all BgpRouteFilters = "bgp-rf-all"

    // BGP CIDR only route filter
    BgpRouteFilters_bgp_rf_cidr_only BgpRouteFilters = "bgp-rf-cidr-only"

    // BGP label route filter
    BgpRouteFilters_bgp_rf_label BgpRouteFilters = "bgp-rf-label"

    // BGP RIB failure route filter
    BgpRouteFilters_bgp_rf_rib_failure BgpRouteFilters = "bgp-rf-rib-failure"

    // BGP injected route filter
    BgpRouteFilters_bgp_rf_injected BgpRouteFilters = "bgp-rf-injected"

    // BGP inconsistent route filter
    BgpRouteFilters_bgp_rf_inconsistent BgpRouteFilters = "bgp-rf-inconsistent"

    // BGP community route filter
    BgpRouteFilters_bgp_rf_community BgpRouteFilters = "bgp-rf-community"

    // BGP extcommunity route filter
    BgpRouteFilters_bgp_rf_extcommunity BgpRouteFilters = "bgp-rf-extcommunity"

    // BGP OER controlled route filter
    BgpRouteFilters_bgp_rf_oer_controlled BgpRouteFilters = "bgp-rf-oer-controlled"

    // BGP pending route filter
    BgpRouteFilters_bgp_rf_pending BgpRouteFilters = "bgp-rf-pending"
)

// BgpNeighborRouteFilters represents BGP neighbor route filters
type BgpNeighborRouteFilters string

const (
    // BGP received routes post policy
    BgpNeighborRouteFilters_bgp_nrf_post_received BgpNeighborRouteFilters = "bgp-nrf-post-received"

    // BGP received routes pre policy
    BgpNeighborRouteFilters_bgp_nrf_pre_received BgpNeighborRouteFilters = "bgp-nrf-pre-received"

    // BGP pre advertised pre policy
    BgpNeighborRouteFilters_bgp_nrf_pre_advertised BgpNeighborRouteFilters = "bgp-nrf-pre-advertised"

    // BGP post advertised post policy
    BgpNeighborRouteFilters_bgp_nrf_post_advertised BgpNeighborRouteFilters = "bgp-nrf-post-advertised"
)

