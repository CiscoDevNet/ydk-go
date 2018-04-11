// Defines identity and type defintions associated with
// the OpenConfig BGP RIB modules
package rib_bgp_types

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package rib_bgp_types"))
}

type INVALIDROUTEREASON struct {
}

func (id INVALIDROUTEREASON) String() string {
	return "openconfig-rib-bgp-types:INVALID_ROUTE_REASON"
}

type INVALIDCLUSTERLOOP struct {
}

func (id INVALIDCLUSTERLOOP) String() string {
	return "openconfig-rib-bgp-types:INVALID_CLUSTER_LOOP"
}

type INVALIDASLOOP struct {
}

func (id INVALIDASLOOP) String() string {
	return "openconfig-rib-bgp-types:INVALID_AS_LOOP"
}

type INVALIDORIGINATOR struct {
}

func (id INVALIDORIGINATOR) String() string {
	return "openconfig-rib-bgp-types:INVALID_ORIGINATOR"
}

type INVALIDCONFED struct {
}

func (id INVALIDCONFED) String() string {
	return "openconfig-rib-bgp-types:INVALID_CONFED"
}

type BGPNOTSELECTEDBESTPATH struct {
}

func (id BGPNOTSELECTEDBESTPATH) String() string {
	return "openconfig-rib-bgp-types:BGP_NOT_SELECTED_BESTPATH"
}

type LOCALPREFLOWER struct {
}

func (id LOCALPREFLOWER) String() string {
	return "openconfig-rib-bgp-types:LOCAL_PREF_LOWER"
}

type ASPATHLONGER struct {
}

func (id ASPATHLONGER) String() string {
	return "openconfig-rib-bgp-types:AS_PATH_LONGER"
}

type ORIGINTYPEHIGHER struct {
}

func (id ORIGINTYPEHIGHER) String() string {
	return "openconfig-rib-bgp-types:ORIGIN_TYPE_HIGHER"
}

type MEDHIGHER struct {
}

func (id MEDHIGHER) String() string {
	return "openconfig-rib-bgp-types:MED_HIGHER"
}

type PREFEREXTERNAL struct {
}

func (id PREFEREXTERNAL) String() string {
	return "openconfig-rib-bgp-types:PREFER_EXTERNAL"
}

type NEXTHOPCOSTHIGHER struct {
}

func (id NEXTHOPCOSTHIGHER) String() string {
	return "openconfig-rib-bgp-types:NEXTHOP_COST_HIGHER"
}

type HIGHERROUTERID struct {
}

func (id HIGHERROUTERID) String() string {
	return "openconfig-rib-bgp-types:HIGHER_ROUTER_ID"
}

type HIGHERPEERADDRESS struct {
}

func (id HIGHERPEERADDRESS) String() string {
	return "openconfig-rib-bgp-types:HIGHER_PEER_ADDRESS"
}

type BGPNOTSELECTEDPOLICY struct {
}

func (id BGPNOTSELECTEDPOLICY) String() string {
	return "openconfig-rib-bgp-types:BGP_NOT_SELECTED_POLICY"
}

type REJECTEDIMPORTPOLICY struct {
}

func (id REJECTEDIMPORTPOLICY) String() string {
	return "openconfig-rib-bgp-types:REJECTED_IMPORT_POLICY"
}

