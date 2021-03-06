// This module contains general LSDB type definitions for use in ISIS YANG
// model. 
package isis_lsdb_types

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package isis_lsdb_types"))
}

type AREAADDRESSES struct {
}

func (id AREAADDRESSES) String() string {
	return "openconfig-isis-lsdb-types:AREA_ADDRESSES"
}

type ISREACHABILITYLINKPROTECTIONTYPE struct {
}

func (id ISREACHABILITYLINKPROTECTIONTYPE) String() string {
	return "openconfig-isis-lsdb-types:IS_REACHABILITY_LINK_PROTECTION_TYPE"
}

type ISREACHABILITYSUBTLVSTYPE struct {
}

func (id ISREACHABILITYSUBTLVSTYPE) String() string {
	return "openconfig-isis-lsdb-types:IS_REACHABILITY_SUBTLVS_TYPE"
}

type ISREACHABILITYLINKLOSS struct {
}

func (id ISREACHABILITYLINKLOSS) String() string {
	return "openconfig-isis-lsdb-types:IS_REACHABILITY_LINK_LOSS"
}

type MTISNEIGHBORATTRIBUTE struct {
}

func (id MTISNEIGHBORATTRIBUTE) String() string {
	return "openconfig-isis-lsdb-types:MT_IS_NEIGHBOR_ATTRIBUTE"
}

type ISREACHABILITYIPV4INTERFACEADDRESS struct {
}

func (id ISREACHABILITYIPV4INTERFACEADDRESS) String() string {
	return "openconfig-isis-lsdb-types:IS_REACHABILITY_IPV4_INTERFACE_ADDRESS"
}

type ISREACHABILITYIPV4NEIGHBORADDRESS struct {
}

func (id ISREACHABILITYIPV4NEIGHBORADDRESS) String() string {
	return "openconfig-isis-lsdb-types:IS_REACHABILITY_IPV4_NEIGHBOR_ADDRESS"
}

type ISREACHABILITYEXTENDEDADMINGROUP struct {
}

func (id ISREACHABILITYEXTENDEDADMINGROUP) String() string {
	return "openconfig-isis-lsdb-types:IS_REACHABILITY_EXTENDED_ADMIN_GROUP"
}

type NLPID struct {
}

func (id NLPID) String() string {
	return "openconfig-isis-lsdb-types:NLPID"
}

type MTIPV6REACHABILITY struct {
}

func (id MTIPV6REACHABILITY) String() string {
	return "openconfig-isis-lsdb-types:MT_IPV6_REACHABILITY"
}

type MTISN struct {
}

func (id MTISN) String() string {
	return "openconfig-isis-lsdb-types:MT_ISN"
}

type ISNEIGHBORATTRIBUTE struct {
}

func (id ISNEIGHBORATTRIBUTE) String() string {
	return "openconfig-isis-lsdb-types:IS_NEIGHBOR_ATTRIBUTE"
}

type ISISALIASID struct {
}

func (id ISISALIASID) String() string {
	return "openconfig-isis-lsdb-types:ISIS_ALIAS_ID"
}

type IPV6SRLG struct {
}

func (id IPV6SRLG) String() string {
	return "openconfig-isis-lsdb-types:IPV6_SRLG"
}

type ISREACHABILITYBANDWIDTHCONSTRAINTS struct {
}

func (id ISREACHABILITYBANDWIDTHCONSTRAINTS) String() string {
	return "openconfig-isis-lsdb-types:IS_REACHABILITY_BANDWIDTH_CONSTRAINTS"
}

type IPV4TEROUTERID struct {
}

func (id IPV4TEROUTERID) String() string {
	return "openconfig-isis-lsdb-types:IPV4_TE_ROUTER_ID"
}

type ISREACHABILITYIPV6INTERFACEADDRESS struct {
}

func (id ISREACHABILITYIPV6INTERFACEADDRESS) String() string {
	return "openconfig-isis-lsdb-types:IS_REACHABILITY_IPV6_INTERFACE_ADDRESS"
}

type ROUTERCAPABILITYSUBTLVSTYPE struct {
}

func (id ROUTERCAPABILITYSUBTLVSTYPE) String() string {
	return "openconfig-isis-lsdb-types:ROUTER_CAPABILITY_SUBTLVS_TYPE"
}

type ISREACHABILITYADJSID struct {
}

func (id ISREACHABILITYADJSID) String() string {
	return "openconfig-isis-lsdb-types:IS_REACHABILITY_ADJ_SID"
}

type MULTITOPOLOGY struct {
}

func (id MULTITOPOLOGY) String() string {
	return "openconfig-isis-lsdb-types:MULTI_TOPOLOGY"
}

type ISREACHABILITYLINKATTRIBUTES struct {
}

func (id ISREACHABILITYLINKATTRIBUTES) String() string {
	return "openconfig-isis-lsdb-types:IS_REACHABILITY_LINK_ATTRIBUTES"
}

type IPV6INTERFACEADDRESSES struct {
}

func (id IPV6INTERFACEADDRESSES) String() string {
	return "openconfig-isis-lsdb-types:IPV6_INTERFACE_ADDRESSES"
}

type ISISSUBTLVTYPE struct {
}

func (id ISISSUBTLVTYPE) String() string {
	return "openconfig-isis-lsdb-types:ISIS_SUBTLV_TYPE"
}

type DYNAMICNAME struct {
}

func (id DYNAMICNAME) String() string {
	return "openconfig-isis-lsdb-types:DYNAMIC_NAME"
}

type IPREACHABILITYTAG struct {
}

func (id IPREACHABILITYTAG) String() string {
	return "openconfig-isis-lsdb-types:IP_REACHABILITY_TAG"
}

type ISREACHABILITYAVAILABLEBANDWIDTH struct {
}

func (id ISREACHABILITYAVAILABLEBANDWIDTH) String() string {
	return "openconfig-isis-lsdb-types:IS_REACHABILITY_AVAILABLE_BANDWIDTH"
}

type IPV4INTERFACEADDRESSES struct {
}

func (id IPV4INTERFACEADDRESSES) String() string {
	return "openconfig-isis-lsdb-types:IPV4_INTERFACE_ADDRESSES"
}

type IPV6TEROUTERID struct {
}

func (id IPV6TEROUTERID) String() string {
	return "openconfig-isis-lsdb-types:IPV6_TE_ROUTER_ID"
}

type INSTANCEID struct {
}

func (id INSTANCEID) String() string {
	return "openconfig-isis-lsdb-types:INSTANCE_ID"
}

type ISREACHABILITYADJLANSID struct {
}

func (id ISREACHABILITYADJLANSID) String() string {
	return "openconfig-isis-lsdb-types:IS_REACHABILITY_ADJ_LAN_SID"
}

type MTIPV4REACHABILITY struct {
}

func (id MTIPV4REACHABILITY) String() string {
	return "openconfig-isis-lsdb-types:MT_IPV4_REACHABILITY"
}

type IPREACHABILITYIPV4ROUTERID struct {
}

func (id IPREACHABILITYIPV4ROUTERID) String() string {
	return "openconfig-isis-lsdb-types:IP_REACHABILITY_IPV4_ROUTER_ID"
}

type ISREACHABILITYUTILIZEDBANDWIDTH struct {
}

func (id ISREACHABILITYUTILIZEDBANDWIDTH) String() string {
	return "openconfig-isis-lsdb-types:IS_REACHABILITY_UTILIZED_BANDWIDTH"
}

type IPREACHABILITYSUBTLVSTYPE struct {
}

func (id IPREACHABILITYSUBTLVSTYPE) String() string {
	return "openconfig-isis-lsdb-types:IP_REACHABILITY_SUBTLVS_TYPE"
}

type ISREACHABILITYMINMAXLINKDELAY struct {
}

func (id ISREACHABILITYMINMAXLINKDELAY) String() string {
	return "openconfig-isis-lsdb-types:IS_REACHABILITY_MIN_MAX_LINK_DELAY"
}

type IPREACHABILITYPREFIXSID struct {
}

func (id IPREACHABILITYPREFIXSID) String() string {
	return "openconfig-isis-lsdb-types:IP_REACHABILITY_PREFIX_SID"
}

type EXTENDEDIPV4REACHABILITY struct {
}

func (id EXTENDEDIPV4REACHABILITY) String() string {
	return "openconfig-isis-lsdb-types:EXTENDED_IPV4_REACHABILITY"
}

type IPREACHABILITYPREFIXFLAGS struct {
}

func (id IPREACHABILITYPREFIXFLAGS) String() string {
	return "openconfig-isis-lsdb-types:IP_REACHABILITY_PREFIX_FLAGS"
}

type ISREACHABILITYMAXRESERVABLEBANDWIDTH struct {
}

func (id ISREACHABILITYMAXRESERVABLEBANDWIDTH) String() string {
	return "openconfig-isis-lsdb-types:IS_REACHABILITY_MAX_RESERVABLE_BANDWIDTH"
}

type PURGEOI struct {
}

func (id PURGEOI) String() string {
	return "openconfig-isis-lsdb-types:PURGE_OI"
}

type IPV4EXTERNALREACHABILITY struct {
}

func (id IPV4EXTERNALREACHABILITY) String() string {
	return "openconfig-isis-lsdb-types:IPV4_EXTERNAL_REACHABILITY"
}

type IPV6REACHABILITY struct {
}

func (id IPV6REACHABILITY) String() string {
	return "openconfig-isis-lsdb-types:IPV6_REACHABILITY"
}

type ISREACHABILITYMAXLINKBANDWIDTH struct {
}

func (id ISREACHABILITYMAXLINKBANDWIDTH) String() string {
	return "openconfig-isis-lsdb-types:IS_REACHABILITY_MAX_LINK_BANDWIDTH"
}

type IPV4SRLG struct {
}

func (id IPV4SRLG) String() string {
	return "openconfig-isis-lsdb-types:IPV4_SRLG"
}

type IPV4INTERNALREACHABILITY struct {
}

func (id IPV4INTERNALREACHABILITY) String() string {
	return "openconfig-isis-lsdb-types:IPV4_INTERNAL_REACHABILITY"
}

type ISREACHABILITYUNRESERVEDBANDWIDTH struct {
}

func (id ISREACHABILITYUNRESERVEDBANDWIDTH) String() string {
	return "openconfig-isis-lsdb-types:IS_REACHABILITY_UNRESERVED_BANDWIDTH"
}

type ISISTLVTYPE struct {
}

func (id ISISTLVTYPE) String() string {
	return "openconfig-isis-lsdb-types:ISIS_TLV_TYPE"
}

type ROUTERCAPABILITYSRCAPABILITY struct {
}

func (id ROUTERCAPABILITYSRCAPABILITY) String() string {
	return "openconfig-isis-lsdb-types:ROUTER_CAPABILITY_SR_CAPABILITY"
}

type ROUTERCAPABILITY struct {
}

func (id ROUTERCAPABILITY) String() string {
	return "openconfig-isis-lsdb-types:ROUTER_CAPABILITY"
}

type EXTENDEDISREACHABILITY struct {
}

func (id EXTENDEDISREACHABILITY) String() string {
	return "openconfig-isis-lsdb-types:EXTENDED_IS_REACHABILITY"
}

type ISREACHABILITYADMINGROUP struct {
}

func (id ISREACHABILITYADMINGROUP) String() string {
	return "openconfig-isis-lsdb-types:IS_REACHABILITY_ADMIN_GROUP"
}

type IPREACHABILITYTAG64 struct {
}

func (id IPREACHABILITYTAG64) String() string {
	return "openconfig-isis-lsdb-types:IP_REACHABILITY_TAG64"
}

type ISREACHABILITYLINKDELAY struct {
}

func (id ISREACHABILITYLINKDELAY) String() string {
	return "openconfig-isis-lsdb-types:IS_REACHABILITY_LINK_DELAY"
}

type ISREACHABILITYTEDEFAULTMETRIC struct {
}

func (id ISREACHABILITYTEDEFAULTMETRIC) String() string {
	return "openconfig-isis-lsdb-types:IS_REACHABILITY_TE_DEFAULT_METRIC"
}

type AUTHENTICATION struct {
}

func (id AUTHENTICATION) String() string {
	return "openconfig-isis-lsdb-types:AUTHENTICATION"
}

type IISNEIGHBORS struct {
}

func (id IISNEIGHBORS) String() string {
	return "openconfig-isis-lsdb-types:IIS_NEIGHBORS"
}

type ISREACHABILITYUNCONSTRAINEDLSP struct {
}

func (id ISREACHABILITYUNCONSTRAINEDLSP) String() string {
	return "openconfig-isis-lsdb-types:IS_REACHABILITY_UNCONSTRAINED_LSP"
}

type ISREACHABILITYIPV6NEIGHBORADDRESS struct {
}

func (id ISREACHABILITYIPV6NEIGHBORADDRESS) String() string {
	return "openconfig-isis-lsdb-types:IS_REACHABILITY_IPV6_NEIGHBOR_ADDRESS"
}

type ISREACHABILITYRESIDUALBANDWIDTH struct {
}

func (id ISREACHABILITYRESIDUALBANDWIDTH) String() string {
	return "openconfig-isis-lsdb-types:IS_REACHABILITY_RESIDUAL_BANDWIDTH"
}

type ISREACHABILITYLINKDELAYVARIATION struct {
}

func (id ISREACHABILITYLINKDELAYVARIATION) String() string {
	return "openconfig-isis-lsdb-types:IS_REACHABILITY_LINK_DELAY_VARIATION"
}

type ROUTERCAPABILITYSRALGORITHM struct {
}

func (id ROUTERCAPABILITYSRALGORITHM) String() string {
	return "openconfig-isis-lsdb-types:ROUTER_CAPABILITY_SR_ALGORITHM"
}

type IPREACHABILITYIPV6ROUTERID struct {
}

func (id IPREACHABILITYIPV6ROUTERID) String() string {
	return "openconfig-isis-lsdb-types:IP_REACHABILITY_IPV6_ROUTER_ID"
}

