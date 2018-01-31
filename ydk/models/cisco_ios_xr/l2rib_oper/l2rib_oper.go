// This module contains a collection of YANG definitions
// for Cisco IOS-XR l2rib package operational data.
// 
// This module contains definitions
// for the following management objects:
//   l2rib: L2RIB operational information 
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package l2rib_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package l2rib_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-l2rib-oper l2rib}", reflect.TypeOf(L2Rib{}))
    ydk.RegisterEntity("Cisco-IOS-XR-l2rib-oper:l2rib", reflect.TypeOf(L2Rib{}))
}

// L2ribMacRoute represents L2rib mac route
type L2ribMacRoute string

const (
    // l2rib mac route type invalid
    L2ribMacRoute_l2rib_mac_route_type_invalid L2ribMacRoute = "l2rib-mac-route-type-invalid"

    // l2rib mac route type regular
    L2ribMacRoute_l2rib_mac_route_type_regular L2ribMacRoute = "l2rib-mac-route-type-regular"

    // l2rib mac route type evpn esi
    L2ribMacRoute_l2rib_mac_route_type_evpn_esi L2ribMacRoute = "l2rib-mac-route-type-evpn-esi"

    // l2rib mac route type bmac
    L2ribMacRoute_l2rib_mac_route_type_bmac L2ribMacRoute = "l2rib-mac-route-type-bmac"
)

// L2ribBagProducerId represents L2RIB Producer Types
type L2ribBagProducerId string

const (
    // None
    L2ribBagProducerId_l2rib_bag_prod_none L2ribBagProducerId = "l2rib-bag-prod-none"

    // Best Route
    L2ribBagProducerId_l2rib_bag_prod_best_route L2ribBagProducerId = "l2rib-bag-prod-best-route"

    // Static
    L2ribBagProducerId_l2rib_bag_prod_static L2ribBagProducerId = "l2rib-bag-prod-static"

    // Local
    L2ribBagProducerId_l2rib_bag_prod_local L2ribBagProducerId = "l2rib-bag-prod-local"

    // IS IS
    L2ribBagProducerId_l2rib_bag_prod_isis L2ribBagProducerId = "l2rib-bag-prod-isis"

    // BGP
    L2ribBagProducerId_l2rib_bag_prod_bgp L2ribBagProducerId = "l2rib-bag-prod-bgp"

    // IGMP
    L2ribBagProducerId_l2rib_bag_prod_igmp L2ribBagProducerId = "l2rib-bag-prod-igmp"

    // MLD
    L2ribBagProducerId_l2rib_bag_prod_prod_mld L2ribBagProducerId = "l2rib-bag-prod-prod-mld"

    // OTV
    L2ribBagProducerId_l2rib_bag_prod_prod_otv L2ribBagProducerId = "l2rib-bag-prod-prod-otv"

    // L2VPN
    L2ribBagProducerId_l2rib_bag_prod_prod_l2vpn L2ribBagProducerId = "l2rib-bag-prod-prod-l2vpn"

    // MAC MGR
    L2ribBagProducerId_l2rib_bag_prod_prod_mac_mgr L2ribBagProducerId = "l2rib-bag-prod-prod-mac-mgr"

    // VXLAN
    L2ribBagProducerId_l2rib_bag_prod_prod_vxlan L2ribBagProducerId = "l2rib-bag-prod-prod-vxlan"

    // HMM
    L2ribBagProducerId_l2rib_bag_prod_prod_hmm L2ribBagProducerId = "l2rib-bag-prod-prod-hmm"

    // ARP
    L2ribBagProducerId_l2rib_bag_prod_prod_arp L2ribBagProducerId = "l2rib-bag-prod-prod-arp"

    // All
    L2ribBagProducerId_l2rib_bag_prod_prod_all L2ribBagProducerId = "l2rib-bag-prod-prod-all"
)

// L2ribAfi represents L2rib afi
type L2ribAfi string

const (
    // l2rib address family ipv4
    L2ribAfi_l2rib_address_family_ipv4 L2ribAfi = "l2rib-address-family-ipv4"

    // l2rib address family ipv6
    L2ribAfi_l2rib_address_family_ipv6 L2ribAfi = "l2rib-address-family-ipv6"

    // l2rib address family invalid
    L2ribAfi_l2rib_address_family_invalid L2ribAfi = "l2rib-address-family-invalid"
)

// L2ribBagProducerState represents L2RIB Producer States
type L2ribBagProducerState string

const (
    // Initial
    L2ribBagProducerState_l2rib_bag_prod_state_initial L2ribBagProducerState = "l2rib-bag-prod-state-initial"

    // Stale
    L2ribBagProducerState_l2rib_bag_prod_state_staled L2ribBagProducerState = "l2rib-bag-prod-state-staled"

    // Reconnected
    L2ribBagProducerState_l2rib_bag_prod_state_re_connected L2ribBagProducerState = "l2rib-bag-prod-state-re-connected"

    // Converged
    L2ribBagProducerState_l2rib_bag_prod_state_converged L2ribBagProducerState = "l2rib-bag-prod-state-converged"

    // Delete Pending
    L2ribBagProducerState_l2rib_bag_prod_state_delete_p_end L2ribBagProducerState = "l2rib-bag-prod-state-delete-p-end"
)

// L2ribNextHop represents L2rib next hop
type L2ribNextHop string

const (
    // l2rib next hop invalid
    L2ribNextHop_l2rib_next_hop_invalid L2ribNextHop = "l2rib-next-hop-invalid"

    // l2rib next hop interface ordinal
    L2ribNextHop_l2rib_next_hop_interface_ordinal L2ribNextHop = "l2rib-next-hop-interface-ordinal"

    // l2rib next hop interface index
    L2ribNextHop_l2rib_next_hop_interface_index L2ribNextHop = "l2rib-next-hop-interface-index"

    // l2rib next hop mac
    L2ribNextHop_l2rib_next_hop_mac L2ribNextHop = "l2rib-next-hop-mac"

    // l2rib next hop ipv4
    L2ribNextHop_l2rib_next_hop_ipv4 L2ribNextHop = "l2rib-next-hop-ipv4"

    // l2rib next hop ipv6
    L2ribNextHop_l2rib_next_hop_ipv6 L2ribNextHop = "l2rib-next-hop-ipv6"

    // l2rib next hop overlay
    L2ribNextHop_l2rib_next_hop_overlay L2ribNextHop = "l2rib-next-hop-overlay"

    // l2rib next hop site index
    L2ribNextHop_l2rib_next_hop_site_index L2ribNextHop = "l2rib-next-hop-site-index"

    // l2rib next hop label ed
    L2ribNextHop_l2rib_next_hop_label_ed L2ribNextHop = "l2rib-next-hop-label-ed"

    // l2rib next hop xid
    L2ribNextHop_l2rib_next_hop_xid L2ribNextHop = "l2rib-next-hop-xid"
)

// L2ribBagObj represents L2RIB Object Types
type L2ribBagObj string

const (
    // Invalid Object Type
    L2ribBagObj_l2rib_bag_obj_type_min L2ribBagObj = "l2rib-bag-obj-type-min"

    // All
    L2ribBagObj_l2rib_bag_obj_type_all L2ribBagObj = "l2rib-bag-obj-type-all"

    // Mac
    L2ribBagObj_l2rib_bag_obj_type_mac L2ribBagObj = "l2rib-bag-obj-type-mac"

    // IPv4 Multicast
    L2ribBagObj_l2rib_bag_obj_type_ipv4_mcast L2ribBagObj = "l2rib-bag-obj-type-ipv4-mcast"

    // IPv6 Multicast
    L2ribBagObj_l2rib_bag_obj_type_ipv6_mcast L2ribBagObj = "l2rib-bag-obj-type-ipv6-mcast"

    // Topology
    L2ribBagObj_l2rib_bag_obj_type_topology L2ribBagObj = "l2rib-bag-obj-type-topology"

    // Ethernet AD
    L2ribBagObj_l2rib_bag_obj_type_ead L2ribBagObj = "l2rib-bag-obj-type-ead"

    // EVPN Path List
    L2ribBagObj_l2rib_bag_obj_type_evpn_pl L2ribBagObj = "l2rib-bag-obj-type-evpn-pl"

    // Topology Attribute
    L2ribBagObj_l2rib_bag_obj_type_topo_attribute L2ribBagObj = "l2rib-bag-obj-type-topo-attribute"

    // IMET
    L2ribBagObj_l2rib_bag_obj_type_imet_route L2ribBagObj = "l2rib-bag-obj-type-imet-route"

    // Mac IP
    L2ribBagObj_l2rib_bag_obj_type_mac_ip L2ribBagObj = "l2rib-bag-obj-type-mac-ip"
)

// L2Rib
// L2RIB operational information 
type L2Rib struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // L2RIB detailed producer table.
    ProducersDetails L2Rib_ProducersDetails

    // L2RIB EVPN Summary.
    Summary L2Rib_Summary

    // L2RIB producer table.
    Producers L2Rib_Producers

    // L2RIB client table.
    Clients L2Rib_Clients

    // L2RIB EVPN EVI Detail Table.
    EvisXr L2Rib_EvisXr

    // L2RIB detailed client table.
    ClientsDetails L2Rib_ClientsDetails

    // Container for all EVI Child Tables.
    EviChildTables L2Rib_EviChildTables

    // L2RIB EVPN EVI Table.
    Evis L2Rib_Evis
}

func (l2Rib *L2Rib) GetFilter() yfilter.YFilter { return l2Rib.YFilter }

func (l2Rib *L2Rib) SetFilter(yf yfilter.YFilter) { l2Rib.YFilter = yf }

func (l2Rib *L2Rib) GetGoName(yname string) string {
    if yname == "producers-details" { return "ProducersDetails" }
    if yname == "summary" { return "Summary" }
    if yname == "producers" { return "Producers" }
    if yname == "clients" { return "Clients" }
    if yname == "evis-xr" { return "EvisXr" }
    if yname == "clients-details" { return "ClientsDetails" }
    if yname == "evi-child-tables" { return "EviChildTables" }
    if yname == "evis" { return "Evis" }
    return ""
}

func (l2Rib *L2Rib) GetSegmentPath() string {
    return "Cisco-IOS-XR-l2rib-oper:l2rib"
}

func (l2Rib *L2Rib) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "producers-details" {
        return &l2Rib.ProducersDetails
    }
    if childYangName == "summary" {
        return &l2Rib.Summary
    }
    if childYangName == "producers" {
        return &l2Rib.Producers
    }
    if childYangName == "clients" {
        return &l2Rib.Clients
    }
    if childYangName == "evis-xr" {
        return &l2Rib.EvisXr
    }
    if childYangName == "clients-details" {
        return &l2Rib.ClientsDetails
    }
    if childYangName == "evi-child-tables" {
        return &l2Rib.EviChildTables
    }
    if childYangName == "evis" {
        return &l2Rib.Evis
    }
    return nil
}

func (l2Rib *L2Rib) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["producers-details"] = &l2Rib.ProducersDetails
    children["summary"] = &l2Rib.Summary
    children["producers"] = &l2Rib.Producers
    children["clients"] = &l2Rib.Clients
    children["evis-xr"] = &l2Rib.EvisXr
    children["clients-details"] = &l2Rib.ClientsDetails
    children["evi-child-tables"] = &l2Rib.EviChildTables
    children["evis"] = &l2Rib.Evis
    return children
}

func (l2Rib *L2Rib) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (l2Rib *L2Rib) GetBundleName() string { return "cisco_ios_xr" }

func (l2Rib *L2Rib) GetYangName() string { return "l2rib" }

func (l2Rib *L2Rib) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (l2Rib *L2Rib) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (l2Rib *L2Rib) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (l2Rib *L2Rib) SetParent(parent types.Entity) { l2Rib.parent = parent }

func (l2Rib *L2Rib) GetParent() types.Entity { return l2Rib.parent }

func (l2Rib *L2Rib) GetParentYangName() string { return "Cisco-IOS-XR-l2rib-oper" }

// L2Rib_ProducersDetails
// L2RIB detailed producer table
type L2Rib_ProducersDetails struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // L2RIB producers detail information. The type is slice of
    // L2Rib_ProducersDetails_ProducersDetail.
    ProducersDetail []L2Rib_ProducersDetails_ProducersDetail
}

func (producersDetails *L2Rib_ProducersDetails) GetFilter() yfilter.YFilter { return producersDetails.YFilter }

func (producersDetails *L2Rib_ProducersDetails) SetFilter(yf yfilter.YFilter) { producersDetails.YFilter = yf }

func (producersDetails *L2Rib_ProducersDetails) GetGoName(yname string) string {
    if yname == "producers-detail" { return "ProducersDetail" }
    return ""
}

func (producersDetails *L2Rib_ProducersDetails) GetSegmentPath() string {
    return "producers-details"
}

func (producersDetails *L2Rib_ProducersDetails) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "producers-detail" {
        for _, c := range producersDetails.ProducersDetail {
            if producersDetails.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := L2Rib_ProducersDetails_ProducersDetail{}
        producersDetails.ProducersDetail = append(producersDetails.ProducersDetail, child)
        return &producersDetails.ProducersDetail[len(producersDetails.ProducersDetail)-1]
    }
    return nil
}

func (producersDetails *L2Rib_ProducersDetails) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range producersDetails.ProducersDetail {
        children[producersDetails.ProducersDetail[i].GetSegmentPath()] = &producersDetails.ProducersDetail[i]
    }
    return children
}

func (producersDetails *L2Rib_ProducersDetails) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (producersDetails *L2Rib_ProducersDetails) GetBundleName() string { return "cisco_ios_xr" }

func (producersDetails *L2Rib_ProducersDetails) GetYangName() string { return "producers-details" }

func (producersDetails *L2Rib_ProducersDetails) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (producersDetails *L2Rib_ProducersDetails) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (producersDetails *L2Rib_ProducersDetails) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (producersDetails *L2Rib_ProducersDetails) SetParent(parent types.Entity) { producersDetails.parent = parent }

func (producersDetails *L2Rib_ProducersDetails) GetParent() types.Entity { return producersDetails.parent }

func (producersDetails *L2Rib_ProducersDetails) GetParentYangName() string { return "l2rib" }

// L2Rib_ProducersDetails_ProducersDetail
// L2RIB producers detail information
type L2Rib_ProducersDetails_ProducersDetail struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Object ID. The type is interface{} with range: -2147483648..2147483647.
    ObjectId interface{}

    // Product ID. The type is interface{} with range: -2147483648..2147483647.
    ProductId interface{}

    // Last Update Timestamp. The type is interface{} with range:
    // 0..18446744073709551615.
    LastUpdateTimestamp interface{}

    // Non-detail Producer Bag.
    Producer L2Rib_ProducersDetails_ProducersDetail_Producer

    // Producer Statistics.
    Statistics L2Rib_ProducersDetails_ProducersDetail_Statistics
}

func (producersDetail *L2Rib_ProducersDetails_ProducersDetail) GetFilter() yfilter.YFilter { return producersDetail.YFilter }

func (producersDetail *L2Rib_ProducersDetails_ProducersDetail) SetFilter(yf yfilter.YFilter) { producersDetail.YFilter = yf }

func (producersDetail *L2Rib_ProducersDetails_ProducersDetail) GetGoName(yname string) string {
    if yname == "object-id" { return "ObjectId" }
    if yname == "product-id" { return "ProductId" }
    if yname == "last-update-timestamp" { return "LastUpdateTimestamp" }
    if yname == "producer" { return "Producer" }
    if yname == "statistics" { return "Statistics" }
    return ""
}

func (producersDetail *L2Rib_ProducersDetails_ProducersDetail) GetSegmentPath() string {
    return "producers-detail"
}

func (producersDetail *L2Rib_ProducersDetails_ProducersDetail) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "producer" {
        return &producersDetail.Producer
    }
    if childYangName == "statistics" {
        return &producersDetail.Statistics
    }
    return nil
}

func (producersDetail *L2Rib_ProducersDetails_ProducersDetail) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["producer"] = &producersDetail.Producer
    children["statistics"] = &producersDetail.Statistics
    return children
}

func (producersDetail *L2Rib_ProducersDetails_ProducersDetail) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object-id"] = producersDetail.ObjectId
    leafs["product-id"] = producersDetail.ProductId
    leafs["last-update-timestamp"] = producersDetail.LastUpdateTimestamp
    return leafs
}

func (producersDetail *L2Rib_ProducersDetails_ProducersDetail) GetBundleName() string { return "cisco_ios_xr" }

func (producersDetail *L2Rib_ProducersDetails_ProducersDetail) GetYangName() string { return "producers-detail" }

func (producersDetail *L2Rib_ProducersDetails_ProducersDetail) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (producersDetail *L2Rib_ProducersDetails_ProducersDetail) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (producersDetail *L2Rib_ProducersDetails_ProducersDetail) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (producersDetail *L2Rib_ProducersDetails_ProducersDetail) SetParent(parent types.Entity) { producersDetail.parent = parent }

func (producersDetail *L2Rib_ProducersDetails_ProducersDetail) GetParent() types.Entity { return producersDetail.parent }

func (producersDetail *L2Rib_ProducersDetails_ProducersDetail) GetParentYangName() string { return "producers-details" }

// L2Rib_ProducersDetails_ProducersDetail_Producer
// Non-detail Producer Bag
type L2Rib_ProducersDetails_ProducersDetail_Producer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Client ID. The type is interface{} with range: 0..4294967295.
    ClientId interface{}

    // Object Type. The type is L2ribBagObj.
    ObjectType interface{}

    // Producer ID. The type is L2ribBagProducerId.
    ProducerId interface{}

    // Producer Name. The type is string.
    ProducerName interface{}

    // Admin Distance. The type is interface{} with range: 0..4294967295.
    AdminDistance interface{}

    // Purge Time. The type is interface{} with range: 0..4294967295.
    PurgeTime interface{}

    // Producer State. The type is L2ribBagProducerState.
    State interface{}
}

func (producer *L2Rib_ProducersDetails_ProducersDetail_Producer) GetFilter() yfilter.YFilter { return producer.YFilter }

func (producer *L2Rib_ProducersDetails_ProducersDetail_Producer) SetFilter(yf yfilter.YFilter) { producer.YFilter = yf }

func (producer *L2Rib_ProducersDetails_ProducersDetail_Producer) GetGoName(yname string) string {
    if yname == "client-id" { return "ClientId" }
    if yname == "object-type" { return "ObjectType" }
    if yname == "producer-id" { return "ProducerId" }
    if yname == "producer-name" { return "ProducerName" }
    if yname == "admin-distance" { return "AdminDistance" }
    if yname == "purge-time" { return "PurgeTime" }
    if yname == "state" { return "State" }
    return ""
}

func (producer *L2Rib_ProducersDetails_ProducersDetail_Producer) GetSegmentPath() string {
    return "producer"
}

func (producer *L2Rib_ProducersDetails_ProducersDetail_Producer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (producer *L2Rib_ProducersDetails_ProducersDetail_Producer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (producer *L2Rib_ProducersDetails_ProducersDetail_Producer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["client-id"] = producer.ClientId
    leafs["object-type"] = producer.ObjectType
    leafs["producer-id"] = producer.ProducerId
    leafs["producer-name"] = producer.ProducerName
    leafs["admin-distance"] = producer.AdminDistance
    leafs["purge-time"] = producer.PurgeTime
    leafs["state"] = producer.State
    return leafs
}

func (producer *L2Rib_ProducersDetails_ProducersDetail_Producer) GetBundleName() string { return "cisco_ios_xr" }

func (producer *L2Rib_ProducersDetails_ProducersDetail_Producer) GetYangName() string { return "producer" }

func (producer *L2Rib_ProducersDetails_ProducersDetail_Producer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (producer *L2Rib_ProducersDetails_ProducersDetail_Producer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (producer *L2Rib_ProducersDetails_ProducersDetail_Producer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (producer *L2Rib_ProducersDetails_ProducersDetail_Producer) SetParent(parent types.Entity) { producer.parent = parent }

func (producer *L2Rib_ProducersDetails_ProducersDetail_Producer) GetParent() types.Entity { return producer.parent }

func (producer *L2Rib_ProducersDetails_ProducersDetail_Producer) GetParentYangName() string { return "producers-detail" }

// L2Rib_ProducersDetails_ProducersDetail_Statistics
// Producer Statistics
type L2Rib_ProducersDetails_ProducersDetail_Statistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Producer ID. The type is L2ribBagProducerId.
    ProducerId interface{}

    // Producer Name. The type is string.
    ProducerName interface{}

    // Statistics.
    Statistics L2Rib_ProducersDetails_ProducersDetail_Statistics_Statistics
}

func (statistics *L2Rib_ProducersDetails_ProducersDetail_Statistics) GetFilter() yfilter.YFilter { return statistics.YFilter }

func (statistics *L2Rib_ProducersDetails_ProducersDetail_Statistics) SetFilter(yf yfilter.YFilter) { statistics.YFilter = yf }

func (statistics *L2Rib_ProducersDetails_ProducersDetail_Statistics) GetGoName(yname string) string {
    if yname == "producer-id" { return "ProducerId" }
    if yname == "producer-name" { return "ProducerName" }
    if yname == "statistics" { return "Statistics" }
    return ""
}

func (statistics *L2Rib_ProducersDetails_ProducersDetail_Statistics) GetSegmentPath() string {
    return "statistics"
}

func (statistics *L2Rib_ProducersDetails_ProducersDetail_Statistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "statistics" {
        return &statistics.Statistics
    }
    return nil
}

func (statistics *L2Rib_ProducersDetails_ProducersDetail_Statistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["statistics"] = &statistics.Statistics
    return children
}

func (statistics *L2Rib_ProducersDetails_ProducersDetail_Statistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["producer-id"] = statistics.ProducerId
    leafs["producer-name"] = statistics.ProducerName
    return leafs
}

func (statistics *L2Rib_ProducersDetails_ProducersDetail_Statistics) GetBundleName() string { return "cisco_ios_xr" }

func (statistics *L2Rib_ProducersDetails_ProducersDetail_Statistics) GetYangName() string { return "statistics" }

func (statistics *L2Rib_ProducersDetails_ProducersDetail_Statistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistics *L2Rib_ProducersDetails_ProducersDetail_Statistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistics *L2Rib_ProducersDetails_ProducersDetail_Statistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistics *L2Rib_ProducersDetails_ProducersDetail_Statistics) SetParent(parent types.Entity) { statistics.parent = parent }

func (statistics *L2Rib_ProducersDetails_ProducersDetail_Statistics) GetParent() types.Entity { return statistics.parent }

func (statistics *L2Rib_ProducersDetails_ProducersDetail_Statistics) GetParentYangName() string { return "producers-detail" }

// L2Rib_ProducersDetails_ProducersDetail_Statistics_Statistics
// Statistics
type L2Rib_ProducersDetails_ProducersDetail_Statistics_Statistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Memory Size. The type is interface{} with range: 0..4294967295.
    MemorySize interface{}

    // Number of Objects. The type is interface{} with range: 0..4294967295.
    ObjectCount interface{}

    // End of Interval Timestamp. The type is interface{} with range:
    // 0..18446744073709551615.
    EndofIntervalTs interface{}

    // Extended Counters. The type is slice of
    // L2Rib_ProducersDetails_ProducersDetail_Statistics_Statistics_ExtendedCounter.
    ExtendedCounter []L2Rib_ProducersDetails_ProducersDetail_Statistics_Statistics_ExtendedCounter
}

func (statistics *L2Rib_ProducersDetails_ProducersDetail_Statistics_Statistics) GetFilter() yfilter.YFilter { return statistics.YFilter }

func (statistics *L2Rib_ProducersDetails_ProducersDetail_Statistics_Statistics) SetFilter(yf yfilter.YFilter) { statistics.YFilter = yf }

func (statistics *L2Rib_ProducersDetails_ProducersDetail_Statistics_Statistics) GetGoName(yname string) string {
    if yname == "memory-size" { return "MemorySize" }
    if yname == "object-count" { return "ObjectCount" }
    if yname == "endof-interval-ts" { return "EndofIntervalTs" }
    if yname == "extended-counter" { return "ExtendedCounter" }
    return ""
}

func (statistics *L2Rib_ProducersDetails_ProducersDetail_Statistics_Statistics) GetSegmentPath() string {
    return "statistics"
}

func (statistics *L2Rib_ProducersDetails_ProducersDetail_Statistics_Statistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "extended-counter" {
        for _, c := range statistics.ExtendedCounter {
            if statistics.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := L2Rib_ProducersDetails_ProducersDetail_Statistics_Statistics_ExtendedCounter{}
        statistics.ExtendedCounter = append(statistics.ExtendedCounter, child)
        return &statistics.ExtendedCounter[len(statistics.ExtendedCounter)-1]
    }
    return nil
}

func (statistics *L2Rib_ProducersDetails_ProducersDetail_Statistics_Statistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range statistics.ExtendedCounter {
        children[statistics.ExtendedCounter[i].GetSegmentPath()] = &statistics.ExtendedCounter[i]
    }
    return children
}

func (statistics *L2Rib_ProducersDetails_ProducersDetail_Statistics_Statistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["memory-size"] = statistics.MemorySize
    leafs["object-count"] = statistics.ObjectCount
    leafs["endof-interval-ts"] = statistics.EndofIntervalTs
    return leafs
}

func (statistics *L2Rib_ProducersDetails_ProducersDetail_Statistics_Statistics) GetBundleName() string { return "cisco_ios_xr" }

func (statistics *L2Rib_ProducersDetails_ProducersDetail_Statistics_Statistics) GetYangName() string { return "statistics" }

func (statistics *L2Rib_ProducersDetails_ProducersDetail_Statistics_Statistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistics *L2Rib_ProducersDetails_ProducersDetail_Statistics_Statistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistics *L2Rib_ProducersDetails_ProducersDetail_Statistics_Statistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistics *L2Rib_ProducersDetails_ProducersDetail_Statistics_Statistics) SetParent(parent types.Entity) { statistics.parent = parent }

func (statistics *L2Rib_ProducersDetails_ProducersDetail_Statistics_Statistics) GetParent() types.Entity { return statistics.parent }

func (statistics *L2Rib_ProducersDetails_ProducersDetail_Statistics_Statistics) GetParentYangName() string { return "statistics" }

// L2Rib_ProducersDetails_ProducersDetail_Statistics_Statistics_ExtendedCounter
// Extended Counters
type L2Rib_ProducersDetails_ProducersDetail_Statistics_Statistics_ExtendedCounter struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // CounterType. The type is interface{} with range: 0..255.
    CounterType interface{}

    // CounterName. The type is string.
    CounterName interface{}

    // Real-clock timestamp in msec of first event. The type is interface{} with
    // range: 0..18446744073709551615.
    L2RbFirstEventTs interface{}

    // Real-clock timestamp in msec of last event. The type is interface{} with
    // range: 0..18446744073709551615.
    L2RbLastEventTs interface{}

    // number of events in interval. The type is interface{} with range:
    // 0..4294967295.
    L2RbIntervalEventCount interface{}

    // total number of events. The type is interface{} with range: 0..4294967295.
    L2RbTotalEventCount interface{}
}

func (extendedCounter *L2Rib_ProducersDetails_ProducersDetail_Statistics_Statistics_ExtendedCounter) GetFilter() yfilter.YFilter { return extendedCounter.YFilter }

func (extendedCounter *L2Rib_ProducersDetails_ProducersDetail_Statistics_Statistics_ExtendedCounter) SetFilter(yf yfilter.YFilter) { extendedCounter.YFilter = yf }

func (extendedCounter *L2Rib_ProducersDetails_ProducersDetail_Statistics_Statistics_ExtendedCounter) GetGoName(yname string) string {
    if yname == "counter-type" { return "CounterType" }
    if yname == "counter-name" { return "CounterName" }
    if yname == "l2rb-first-event-ts" { return "L2RbFirstEventTs" }
    if yname == "l2rb-last-event-ts" { return "L2RbLastEventTs" }
    if yname == "l2rb-interval-event-count" { return "L2RbIntervalEventCount" }
    if yname == "l2rb-total-event-count" { return "L2RbTotalEventCount" }
    return ""
}

func (extendedCounter *L2Rib_ProducersDetails_ProducersDetail_Statistics_Statistics_ExtendedCounter) GetSegmentPath() string {
    return "extended-counter"
}

func (extendedCounter *L2Rib_ProducersDetails_ProducersDetail_Statistics_Statistics_ExtendedCounter) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (extendedCounter *L2Rib_ProducersDetails_ProducersDetail_Statistics_Statistics_ExtendedCounter) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (extendedCounter *L2Rib_ProducersDetails_ProducersDetail_Statistics_Statistics_ExtendedCounter) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["counter-type"] = extendedCounter.CounterType
    leafs["counter-name"] = extendedCounter.CounterName
    leafs["l2rb-first-event-ts"] = extendedCounter.L2RbFirstEventTs
    leafs["l2rb-last-event-ts"] = extendedCounter.L2RbLastEventTs
    leafs["l2rb-interval-event-count"] = extendedCounter.L2RbIntervalEventCount
    leafs["l2rb-total-event-count"] = extendedCounter.L2RbTotalEventCount
    return leafs
}

func (extendedCounter *L2Rib_ProducersDetails_ProducersDetail_Statistics_Statistics_ExtendedCounter) GetBundleName() string { return "cisco_ios_xr" }

func (extendedCounter *L2Rib_ProducersDetails_ProducersDetail_Statistics_Statistics_ExtendedCounter) GetYangName() string { return "extended-counter" }

func (extendedCounter *L2Rib_ProducersDetails_ProducersDetail_Statistics_Statistics_ExtendedCounter) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (extendedCounter *L2Rib_ProducersDetails_ProducersDetail_Statistics_Statistics_ExtendedCounter) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (extendedCounter *L2Rib_ProducersDetails_ProducersDetail_Statistics_Statistics_ExtendedCounter) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (extendedCounter *L2Rib_ProducersDetails_ProducersDetail_Statistics_Statistics_ExtendedCounter) SetParent(parent types.Entity) { extendedCounter.parent = parent }

func (extendedCounter *L2Rib_ProducersDetails_ProducersDetail_Statistics_Statistics_ExtendedCounter) GetParent() types.Entity { return extendedCounter.parent }

func (extendedCounter *L2Rib_ProducersDetails_ProducersDetail_Statistics_Statistics_ExtendedCounter) GetParentYangName() string { return "statistics" }

// L2Rib_Summary
// L2RIB EVPN Summary
type L2Rib_Summary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of Converged Tables. The type is interface{} with range:
    // 0..4294967295.
    ConvergedTablesCount interface{}

    // Total Allocated Memory. The type is interface{} with range: 0..4294967295.
    TotalMemory interface{}

    // Per Object Table summary. The type is slice of L2Rib_Summary_TableSummary.
    TableSummary []L2Rib_Summary_TableSummary
}

func (summary *L2Rib_Summary) GetFilter() yfilter.YFilter { return summary.YFilter }

func (summary *L2Rib_Summary) SetFilter(yf yfilter.YFilter) { summary.YFilter = yf }

func (summary *L2Rib_Summary) GetGoName(yname string) string {
    if yname == "converged-tables-count" { return "ConvergedTablesCount" }
    if yname == "total-memory" { return "TotalMemory" }
    if yname == "table-summary" { return "TableSummary" }
    return ""
}

func (summary *L2Rib_Summary) GetSegmentPath() string {
    return "summary"
}

func (summary *L2Rib_Summary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "table-summary" {
        for _, c := range summary.TableSummary {
            if summary.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := L2Rib_Summary_TableSummary{}
        summary.TableSummary = append(summary.TableSummary, child)
        return &summary.TableSummary[len(summary.TableSummary)-1]
    }
    return nil
}

func (summary *L2Rib_Summary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range summary.TableSummary {
        children[summary.TableSummary[i].GetSegmentPath()] = &summary.TableSummary[i]
    }
    return children
}

func (summary *L2Rib_Summary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["converged-tables-count"] = summary.ConvergedTablesCount
    leafs["total-memory"] = summary.TotalMemory
    return leafs
}

func (summary *L2Rib_Summary) GetBundleName() string { return "cisco_ios_xr" }

func (summary *L2Rib_Summary) GetYangName() string { return "summary" }

func (summary *L2Rib_Summary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (summary *L2Rib_Summary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (summary *L2Rib_Summary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (summary *L2Rib_Summary) SetParent(parent types.Entity) { summary.parent = parent }

func (summary *L2Rib_Summary) GetParent() types.Entity { return summary.parent }

func (summary *L2Rib_Summary) GetParentYangName() string { return "l2rib" }

// L2Rib_Summary_TableSummary
// Per Object Table summary
type L2Rib_Summary_TableSummary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Object Type. The type is L2ribBagObj.
    ObjectType interface{}

    // Number of Objects. The type is interface{} with range: 0..4294967295.
    ObjectCount interface{}

    // Allocated Memory. The type is interface{} with range: 0..4294967295.
    TableMemory interface{}

    // Statistics per producer. The type is slice of
    // L2Rib_Summary_TableSummary_ProducerStat.
    ProducerStat []L2Rib_Summary_TableSummary_ProducerStat
}

func (tableSummary *L2Rib_Summary_TableSummary) GetFilter() yfilter.YFilter { return tableSummary.YFilter }

func (tableSummary *L2Rib_Summary_TableSummary) SetFilter(yf yfilter.YFilter) { tableSummary.YFilter = yf }

func (tableSummary *L2Rib_Summary_TableSummary) GetGoName(yname string) string {
    if yname == "object-type" { return "ObjectType" }
    if yname == "object-count" { return "ObjectCount" }
    if yname == "table-memory" { return "TableMemory" }
    if yname == "producer-stat" { return "ProducerStat" }
    return ""
}

func (tableSummary *L2Rib_Summary_TableSummary) GetSegmentPath() string {
    return "table-summary"
}

func (tableSummary *L2Rib_Summary_TableSummary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "producer-stat" {
        for _, c := range tableSummary.ProducerStat {
            if tableSummary.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := L2Rib_Summary_TableSummary_ProducerStat{}
        tableSummary.ProducerStat = append(tableSummary.ProducerStat, child)
        return &tableSummary.ProducerStat[len(tableSummary.ProducerStat)-1]
    }
    return nil
}

func (tableSummary *L2Rib_Summary_TableSummary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range tableSummary.ProducerStat {
        children[tableSummary.ProducerStat[i].GetSegmentPath()] = &tableSummary.ProducerStat[i]
    }
    return children
}

func (tableSummary *L2Rib_Summary_TableSummary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object-type"] = tableSummary.ObjectType
    leafs["object-count"] = tableSummary.ObjectCount
    leafs["table-memory"] = tableSummary.TableMemory
    return leafs
}

func (tableSummary *L2Rib_Summary_TableSummary) GetBundleName() string { return "cisco_ios_xr" }

func (tableSummary *L2Rib_Summary_TableSummary) GetYangName() string { return "table-summary" }

func (tableSummary *L2Rib_Summary_TableSummary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tableSummary *L2Rib_Summary_TableSummary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tableSummary *L2Rib_Summary_TableSummary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tableSummary *L2Rib_Summary_TableSummary) SetParent(parent types.Entity) { tableSummary.parent = parent }

func (tableSummary *L2Rib_Summary_TableSummary) GetParent() types.Entity { return tableSummary.parent }

func (tableSummary *L2Rib_Summary_TableSummary) GetParentYangName() string { return "summary" }

// L2Rib_Summary_TableSummary_ProducerStat
// Statistics per producer
type L2Rib_Summary_TableSummary_ProducerStat struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Producer ID. The type is L2ribBagProducerId.
    ProducerId interface{}

    // Producer Name. The type is string.
    ProducerName interface{}

    // Statistics.
    Statistics L2Rib_Summary_TableSummary_ProducerStat_Statistics
}

func (producerStat *L2Rib_Summary_TableSummary_ProducerStat) GetFilter() yfilter.YFilter { return producerStat.YFilter }

func (producerStat *L2Rib_Summary_TableSummary_ProducerStat) SetFilter(yf yfilter.YFilter) { producerStat.YFilter = yf }

func (producerStat *L2Rib_Summary_TableSummary_ProducerStat) GetGoName(yname string) string {
    if yname == "producer-id" { return "ProducerId" }
    if yname == "producer-name" { return "ProducerName" }
    if yname == "statistics" { return "Statistics" }
    return ""
}

func (producerStat *L2Rib_Summary_TableSummary_ProducerStat) GetSegmentPath() string {
    return "producer-stat"
}

func (producerStat *L2Rib_Summary_TableSummary_ProducerStat) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "statistics" {
        return &producerStat.Statistics
    }
    return nil
}

func (producerStat *L2Rib_Summary_TableSummary_ProducerStat) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["statistics"] = &producerStat.Statistics
    return children
}

func (producerStat *L2Rib_Summary_TableSummary_ProducerStat) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["producer-id"] = producerStat.ProducerId
    leafs["producer-name"] = producerStat.ProducerName
    return leafs
}

func (producerStat *L2Rib_Summary_TableSummary_ProducerStat) GetBundleName() string { return "cisco_ios_xr" }

func (producerStat *L2Rib_Summary_TableSummary_ProducerStat) GetYangName() string { return "producer-stat" }

func (producerStat *L2Rib_Summary_TableSummary_ProducerStat) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (producerStat *L2Rib_Summary_TableSummary_ProducerStat) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (producerStat *L2Rib_Summary_TableSummary_ProducerStat) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (producerStat *L2Rib_Summary_TableSummary_ProducerStat) SetParent(parent types.Entity) { producerStat.parent = parent }

func (producerStat *L2Rib_Summary_TableSummary_ProducerStat) GetParent() types.Entity { return producerStat.parent }

func (producerStat *L2Rib_Summary_TableSummary_ProducerStat) GetParentYangName() string { return "table-summary" }

// L2Rib_Summary_TableSummary_ProducerStat_Statistics
// Statistics
type L2Rib_Summary_TableSummary_ProducerStat_Statistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Memory Size. The type is interface{} with range: 0..4294967295.
    MemorySize interface{}

    // Number of Objects. The type is interface{} with range: 0..4294967295.
    ObjectCount interface{}

    // End of Interval Timestamp. The type is interface{} with range:
    // 0..18446744073709551615.
    EndofIntervalTs interface{}

    // Extended Counters. The type is slice of
    // L2Rib_Summary_TableSummary_ProducerStat_Statistics_ExtendedCounter.
    ExtendedCounter []L2Rib_Summary_TableSummary_ProducerStat_Statistics_ExtendedCounter
}

func (statistics *L2Rib_Summary_TableSummary_ProducerStat_Statistics) GetFilter() yfilter.YFilter { return statistics.YFilter }

func (statistics *L2Rib_Summary_TableSummary_ProducerStat_Statistics) SetFilter(yf yfilter.YFilter) { statistics.YFilter = yf }

func (statistics *L2Rib_Summary_TableSummary_ProducerStat_Statistics) GetGoName(yname string) string {
    if yname == "memory-size" { return "MemorySize" }
    if yname == "object-count" { return "ObjectCount" }
    if yname == "endof-interval-ts" { return "EndofIntervalTs" }
    if yname == "extended-counter" { return "ExtendedCounter" }
    return ""
}

func (statistics *L2Rib_Summary_TableSummary_ProducerStat_Statistics) GetSegmentPath() string {
    return "statistics"
}

func (statistics *L2Rib_Summary_TableSummary_ProducerStat_Statistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "extended-counter" {
        for _, c := range statistics.ExtendedCounter {
            if statistics.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := L2Rib_Summary_TableSummary_ProducerStat_Statistics_ExtendedCounter{}
        statistics.ExtendedCounter = append(statistics.ExtendedCounter, child)
        return &statistics.ExtendedCounter[len(statistics.ExtendedCounter)-1]
    }
    return nil
}

func (statistics *L2Rib_Summary_TableSummary_ProducerStat_Statistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range statistics.ExtendedCounter {
        children[statistics.ExtendedCounter[i].GetSegmentPath()] = &statistics.ExtendedCounter[i]
    }
    return children
}

func (statistics *L2Rib_Summary_TableSummary_ProducerStat_Statistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["memory-size"] = statistics.MemorySize
    leafs["object-count"] = statistics.ObjectCount
    leafs["endof-interval-ts"] = statistics.EndofIntervalTs
    return leafs
}

func (statistics *L2Rib_Summary_TableSummary_ProducerStat_Statistics) GetBundleName() string { return "cisco_ios_xr" }

func (statistics *L2Rib_Summary_TableSummary_ProducerStat_Statistics) GetYangName() string { return "statistics" }

func (statistics *L2Rib_Summary_TableSummary_ProducerStat_Statistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistics *L2Rib_Summary_TableSummary_ProducerStat_Statistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistics *L2Rib_Summary_TableSummary_ProducerStat_Statistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistics *L2Rib_Summary_TableSummary_ProducerStat_Statistics) SetParent(parent types.Entity) { statistics.parent = parent }

func (statistics *L2Rib_Summary_TableSummary_ProducerStat_Statistics) GetParent() types.Entity { return statistics.parent }

func (statistics *L2Rib_Summary_TableSummary_ProducerStat_Statistics) GetParentYangName() string { return "producer-stat" }

// L2Rib_Summary_TableSummary_ProducerStat_Statistics_ExtendedCounter
// Extended Counters
type L2Rib_Summary_TableSummary_ProducerStat_Statistics_ExtendedCounter struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // CounterType. The type is interface{} with range: 0..255.
    CounterType interface{}

    // CounterName. The type is string.
    CounterName interface{}

    // Real-clock timestamp in msec of first event. The type is interface{} with
    // range: 0..18446744073709551615.
    L2RbFirstEventTs interface{}

    // Real-clock timestamp in msec of last event. The type is interface{} with
    // range: 0..18446744073709551615.
    L2RbLastEventTs interface{}

    // number of events in interval. The type is interface{} with range:
    // 0..4294967295.
    L2RbIntervalEventCount interface{}

    // total number of events. The type is interface{} with range: 0..4294967295.
    L2RbTotalEventCount interface{}
}

func (extendedCounter *L2Rib_Summary_TableSummary_ProducerStat_Statistics_ExtendedCounter) GetFilter() yfilter.YFilter { return extendedCounter.YFilter }

func (extendedCounter *L2Rib_Summary_TableSummary_ProducerStat_Statistics_ExtendedCounter) SetFilter(yf yfilter.YFilter) { extendedCounter.YFilter = yf }

func (extendedCounter *L2Rib_Summary_TableSummary_ProducerStat_Statistics_ExtendedCounter) GetGoName(yname string) string {
    if yname == "counter-type" { return "CounterType" }
    if yname == "counter-name" { return "CounterName" }
    if yname == "l2rb-first-event-ts" { return "L2RbFirstEventTs" }
    if yname == "l2rb-last-event-ts" { return "L2RbLastEventTs" }
    if yname == "l2rb-interval-event-count" { return "L2RbIntervalEventCount" }
    if yname == "l2rb-total-event-count" { return "L2RbTotalEventCount" }
    return ""
}

func (extendedCounter *L2Rib_Summary_TableSummary_ProducerStat_Statistics_ExtendedCounter) GetSegmentPath() string {
    return "extended-counter"
}

func (extendedCounter *L2Rib_Summary_TableSummary_ProducerStat_Statistics_ExtendedCounter) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (extendedCounter *L2Rib_Summary_TableSummary_ProducerStat_Statistics_ExtendedCounter) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (extendedCounter *L2Rib_Summary_TableSummary_ProducerStat_Statistics_ExtendedCounter) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["counter-type"] = extendedCounter.CounterType
    leafs["counter-name"] = extendedCounter.CounterName
    leafs["l2rb-first-event-ts"] = extendedCounter.L2RbFirstEventTs
    leafs["l2rb-last-event-ts"] = extendedCounter.L2RbLastEventTs
    leafs["l2rb-interval-event-count"] = extendedCounter.L2RbIntervalEventCount
    leafs["l2rb-total-event-count"] = extendedCounter.L2RbTotalEventCount
    return leafs
}

func (extendedCounter *L2Rib_Summary_TableSummary_ProducerStat_Statistics_ExtendedCounter) GetBundleName() string { return "cisco_ios_xr" }

func (extendedCounter *L2Rib_Summary_TableSummary_ProducerStat_Statistics_ExtendedCounter) GetYangName() string { return "extended-counter" }

func (extendedCounter *L2Rib_Summary_TableSummary_ProducerStat_Statistics_ExtendedCounter) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (extendedCounter *L2Rib_Summary_TableSummary_ProducerStat_Statistics_ExtendedCounter) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (extendedCounter *L2Rib_Summary_TableSummary_ProducerStat_Statistics_ExtendedCounter) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (extendedCounter *L2Rib_Summary_TableSummary_ProducerStat_Statistics_ExtendedCounter) SetParent(parent types.Entity) { extendedCounter.parent = parent }

func (extendedCounter *L2Rib_Summary_TableSummary_ProducerStat_Statistics_ExtendedCounter) GetParent() types.Entity { return extendedCounter.parent }

func (extendedCounter *L2Rib_Summary_TableSummary_ProducerStat_Statistics_ExtendedCounter) GetParentYangName() string { return "statistics" }

// L2Rib_Producers
// L2RIB producer table
type L2Rib_Producers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // L2RIB producers. The type is slice of L2Rib_Producers_Producer.
    Producer []L2Rib_Producers_Producer
}

func (producers *L2Rib_Producers) GetFilter() yfilter.YFilter { return producers.YFilter }

func (producers *L2Rib_Producers) SetFilter(yf yfilter.YFilter) { producers.YFilter = yf }

func (producers *L2Rib_Producers) GetGoName(yname string) string {
    if yname == "producer" { return "Producer" }
    return ""
}

func (producers *L2Rib_Producers) GetSegmentPath() string {
    return "producers"
}

func (producers *L2Rib_Producers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "producer" {
        for _, c := range producers.Producer {
            if producers.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := L2Rib_Producers_Producer{}
        producers.Producer = append(producers.Producer, child)
        return &producers.Producer[len(producers.Producer)-1]
    }
    return nil
}

func (producers *L2Rib_Producers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range producers.Producer {
        children[producers.Producer[i].GetSegmentPath()] = &producers.Producer[i]
    }
    return children
}

func (producers *L2Rib_Producers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (producers *L2Rib_Producers) GetBundleName() string { return "cisco_ios_xr" }

func (producers *L2Rib_Producers) GetYangName() string { return "producers" }

func (producers *L2Rib_Producers) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (producers *L2Rib_Producers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (producers *L2Rib_Producers) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (producers *L2Rib_Producers) SetParent(parent types.Entity) { producers.parent = parent }

func (producers *L2Rib_Producers) GetParent() types.Entity { return producers.parent }

func (producers *L2Rib_Producers) GetParentYangName() string { return "l2rib" }

// L2Rib_Producers_Producer
// L2RIB producers
type L2Rib_Producers_Producer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Object ID. The type is interface{} with range: -2147483648..2147483647.
    ObjectId interface{}

    // Product ID. The type is interface{} with range: -2147483648..2147483647.
    ProductId interface{}

    // Client ID. The type is interface{} with range: 0..4294967295.
    ClientId interface{}

    // Object Type. The type is L2ribBagObj.
    ObjectType interface{}

    // Producer ID. The type is L2ribBagProducerId.
    ProducerId interface{}

    // Producer Name. The type is string.
    ProducerName interface{}

    // Admin Distance. The type is interface{} with range: 0..4294967295.
    AdminDistance interface{}

    // Purge Time. The type is interface{} with range: 0..4294967295.
    PurgeTime interface{}

    // Producer State. The type is L2ribBagProducerState.
    State interface{}
}

func (producer *L2Rib_Producers_Producer) GetFilter() yfilter.YFilter { return producer.YFilter }

func (producer *L2Rib_Producers_Producer) SetFilter(yf yfilter.YFilter) { producer.YFilter = yf }

func (producer *L2Rib_Producers_Producer) GetGoName(yname string) string {
    if yname == "object-id" { return "ObjectId" }
    if yname == "product-id" { return "ProductId" }
    if yname == "client-id" { return "ClientId" }
    if yname == "object-type" { return "ObjectType" }
    if yname == "producer-id" { return "ProducerId" }
    if yname == "producer-name" { return "ProducerName" }
    if yname == "admin-distance" { return "AdminDistance" }
    if yname == "purge-time" { return "PurgeTime" }
    if yname == "state" { return "State" }
    return ""
}

func (producer *L2Rib_Producers_Producer) GetSegmentPath() string {
    return "producer"
}

func (producer *L2Rib_Producers_Producer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (producer *L2Rib_Producers_Producer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (producer *L2Rib_Producers_Producer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object-id"] = producer.ObjectId
    leafs["product-id"] = producer.ProductId
    leafs["client-id"] = producer.ClientId
    leafs["object-type"] = producer.ObjectType
    leafs["producer-id"] = producer.ProducerId
    leafs["producer-name"] = producer.ProducerName
    leafs["admin-distance"] = producer.AdminDistance
    leafs["purge-time"] = producer.PurgeTime
    leafs["state"] = producer.State
    return leafs
}

func (producer *L2Rib_Producers_Producer) GetBundleName() string { return "cisco_ios_xr" }

func (producer *L2Rib_Producers_Producer) GetYangName() string { return "producer" }

func (producer *L2Rib_Producers_Producer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (producer *L2Rib_Producers_Producer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (producer *L2Rib_Producers_Producer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (producer *L2Rib_Producers_Producer) SetParent(parent types.Entity) { producer.parent = parent }

func (producer *L2Rib_Producers_Producer) GetParent() types.Entity { return producer.parent }

func (producer *L2Rib_Producers_Producer) GetParentYangName() string { return "producers" }

// L2Rib_Clients
// L2RIB client table
type L2Rib_Clients struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // L2RIB clients. The type is slice of L2Rib_Clients_Client.
    Client []L2Rib_Clients_Client
}

func (clients *L2Rib_Clients) GetFilter() yfilter.YFilter { return clients.YFilter }

func (clients *L2Rib_Clients) SetFilter(yf yfilter.YFilter) { clients.YFilter = yf }

func (clients *L2Rib_Clients) GetGoName(yname string) string {
    if yname == "client" { return "Client" }
    return ""
}

func (clients *L2Rib_Clients) GetSegmentPath() string {
    return "clients"
}

func (clients *L2Rib_Clients) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "client" {
        for _, c := range clients.Client {
            if clients.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := L2Rib_Clients_Client{}
        clients.Client = append(clients.Client, child)
        return &clients.Client[len(clients.Client)-1]
    }
    return nil
}

func (clients *L2Rib_Clients) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range clients.Client {
        children[clients.Client[i].GetSegmentPath()] = &clients.Client[i]
    }
    return children
}

func (clients *L2Rib_Clients) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (clients *L2Rib_Clients) GetBundleName() string { return "cisco_ios_xr" }

func (clients *L2Rib_Clients) GetYangName() string { return "clients" }

func (clients *L2Rib_Clients) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clients *L2Rib_Clients) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clients *L2Rib_Clients) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clients *L2Rib_Clients) SetParent(parent types.Entity) { clients.parent = parent }

func (clients *L2Rib_Clients) GetParent() types.Entity { return clients.parent }

func (clients *L2Rib_Clients) GetParentYangName() string { return "l2rib" }

// L2Rib_Clients_Client
// L2RIB clients
type L2Rib_Clients_Client struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Client ID. The type is interface{} with range:
    // -2147483648..2147483647.
    ClientId interface{}

    // Client ID. The type is interface{} with range: 0..4294967295.
    ClientIdXr interface{}

    // Process ID. The type is interface{} with range: 0..4294967295.
    ProcessId interface{}

    // Node ID. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeId interface{}

    // Process Name. The type is string.
    ProcName interface{}

    // Process Suffix. The type is string.
    ProcSuffix interface{}
}

func (client *L2Rib_Clients_Client) GetFilter() yfilter.YFilter { return client.YFilter }

func (client *L2Rib_Clients_Client) SetFilter(yf yfilter.YFilter) { client.YFilter = yf }

func (client *L2Rib_Clients_Client) GetGoName(yname string) string {
    if yname == "client-id" { return "ClientId" }
    if yname == "client-id-xr" { return "ClientIdXr" }
    if yname == "process-id" { return "ProcessId" }
    if yname == "node-id" { return "NodeId" }
    if yname == "proc-name" { return "ProcName" }
    if yname == "proc-suffix" { return "ProcSuffix" }
    return ""
}

func (client *L2Rib_Clients_Client) GetSegmentPath() string {
    return "client" + "[client-id='" + fmt.Sprintf("%v", client.ClientId) + "']"
}

func (client *L2Rib_Clients_Client) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (client *L2Rib_Clients_Client) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (client *L2Rib_Clients_Client) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["client-id"] = client.ClientId
    leafs["client-id-xr"] = client.ClientIdXr
    leafs["process-id"] = client.ProcessId
    leafs["node-id"] = client.NodeId
    leafs["proc-name"] = client.ProcName
    leafs["proc-suffix"] = client.ProcSuffix
    return leafs
}

func (client *L2Rib_Clients_Client) GetBundleName() string { return "cisco_ios_xr" }

func (client *L2Rib_Clients_Client) GetYangName() string { return "client" }

func (client *L2Rib_Clients_Client) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (client *L2Rib_Clients_Client) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (client *L2Rib_Clients_Client) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (client *L2Rib_Clients_Client) SetParent(parent types.Entity) { client.parent = parent }

func (client *L2Rib_Clients_Client) GetParent() types.Entity { return client.parent }

func (client *L2Rib_Clients_Client) GetParentYangName() string { return "clients" }

// L2Rib_EvisXr
// L2RIB EVPN EVI Detail Table
type L2Rib_EvisXr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // L2RIB EVPN EVI Entry. The type is slice of L2Rib_EvisXr_Evi.
    Evi []L2Rib_EvisXr_Evi
}

func (evisXr *L2Rib_EvisXr) GetFilter() yfilter.YFilter { return evisXr.YFilter }

func (evisXr *L2Rib_EvisXr) SetFilter(yf yfilter.YFilter) { evisXr.YFilter = yf }

func (evisXr *L2Rib_EvisXr) GetGoName(yname string) string {
    if yname == "evi" { return "Evi" }
    return ""
}

func (evisXr *L2Rib_EvisXr) GetSegmentPath() string {
    return "evis-xr"
}

func (evisXr *L2Rib_EvisXr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "evi" {
        for _, c := range evisXr.Evi {
            if evisXr.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := L2Rib_EvisXr_Evi{}
        evisXr.Evi = append(evisXr.Evi, child)
        return &evisXr.Evi[len(evisXr.Evi)-1]
    }
    return nil
}

func (evisXr *L2Rib_EvisXr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range evisXr.Evi {
        children[evisXr.Evi[i].GetSegmentPath()] = &evisXr.Evi[i]
    }
    return children
}

func (evisXr *L2Rib_EvisXr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (evisXr *L2Rib_EvisXr) GetBundleName() string { return "cisco_ios_xr" }

func (evisXr *L2Rib_EvisXr) GetYangName() string { return "evis-xr" }

func (evisXr *L2Rib_EvisXr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (evisXr *L2Rib_EvisXr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (evisXr *L2Rib_EvisXr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (evisXr *L2Rib_EvisXr) SetParent(parent types.Entity) { evisXr.parent = parent }

func (evisXr *L2Rib_EvisXr) GetParent() types.Entity { return evisXr.parent }

func (evisXr *L2Rib_EvisXr) GetParentYangName() string { return "l2rib" }

// L2Rib_EvisXr_Evi
// L2RIB EVPN EVI Entry
type L2Rib_EvisXr_Evi struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. EVI ID. The type is interface{} with range:
    // -2147483648..2147483647.
    Evi interface{}

    // l2r vni. The type is interface{} with range: 0..4294967295.
    L2RVni interface{}

    // l2r encap type. The type is interface{} with range: 0..65535.
    L2REncapType interface{}

    // l2r nve iod. The type is interface{} with range: 0..4294967295.
    L2RNveIod interface{}

    // l2r nve ifhandle. The type is interface{} with range: 0..4294967295.
    L2RNveIfhandle interface{}

    // VTEP IP. The type is string.
    VtepIp interface{}

    // l2r topo txid. The type is interface{} with range: 0..4294967295.
    L2RTopoTxid interface{}

    // Topology Flags. The type is interface{} with range: 0..4294967295.
    L2RTopoFlags interface{}

    // Topology.
    Topology L2Rib_EvisXr_Evi_Topology
}

func (evi *L2Rib_EvisXr_Evi) GetFilter() yfilter.YFilter { return evi.YFilter }

func (evi *L2Rib_EvisXr_Evi) SetFilter(yf yfilter.YFilter) { evi.YFilter = yf }

func (evi *L2Rib_EvisXr_Evi) GetGoName(yname string) string {
    if yname == "evi" { return "Evi" }
    if yname == "l2r-vni" { return "L2RVni" }
    if yname == "l2r-encap-type" { return "L2REncapType" }
    if yname == "l2r-nve-iod" { return "L2RNveIod" }
    if yname == "l2r-nve-ifhandle" { return "L2RNveIfhandle" }
    if yname == "vtep-ip" { return "VtepIp" }
    if yname == "l2r-topo-txid" { return "L2RTopoTxid" }
    if yname == "l2r-topo-flags" { return "L2RTopoFlags" }
    if yname == "topology" { return "Topology" }
    return ""
}

func (evi *L2Rib_EvisXr_Evi) GetSegmentPath() string {
    return "evi" + "[evi='" + fmt.Sprintf("%v", evi.Evi) + "']"
}

func (evi *L2Rib_EvisXr_Evi) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "topology" {
        return &evi.Topology
    }
    return nil
}

func (evi *L2Rib_EvisXr_Evi) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["topology"] = &evi.Topology
    return children
}

func (evi *L2Rib_EvisXr_Evi) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["evi"] = evi.Evi
    leafs["l2r-vni"] = evi.L2RVni
    leafs["l2r-encap-type"] = evi.L2REncapType
    leafs["l2r-nve-iod"] = evi.L2RNveIod
    leafs["l2r-nve-ifhandle"] = evi.L2RNveIfhandle
    leafs["vtep-ip"] = evi.VtepIp
    leafs["l2r-topo-txid"] = evi.L2RTopoTxid
    leafs["l2r-topo-flags"] = evi.L2RTopoFlags
    return leafs
}

func (evi *L2Rib_EvisXr_Evi) GetBundleName() string { return "cisco_ios_xr" }

func (evi *L2Rib_EvisXr_Evi) GetYangName() string { return "evi" }

func (evi *L2Rib_EvisXr_Evi) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (evi *L2Rib_EvisXr_Evi) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (evi *L2Rib_EvisXr_Evi) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (evi *L2Rib_EvisXr_Evi) SetParent(parent types.Entity) { evi.parent = parent }

func (evi *L2Rib_EvisXr_Evi) GetParent() types.Entity { return evi.parent }

func (evi *L2Rib_EvisXr_Evi) GetParentYangName() string { return "evis-xr" }

// L2Rib_EvisXr_Evi_Topology
// Topology
type L2Rib_EvisXr_Evi_Topology struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Topology ID. The type is interface{} with range: 0..4294967295.
    TopologyId interface{}

    // Topology Name. The type is string.
    TopologyName interface{}

    // Topology Type. The type is interface{} with range: 0..4294967295.
    TopologyType interface{}
}

func (topology *L2Rib_EvisXr_Evi_Topology) GetFilter() yfilter.YFilter { return topology.YFilter }

func (topology *L2Rib_EvisXr_Evi_Topology) SetFilter(yf yfilter.YFilter) { topology.YFilter = yf }

func (topology *L2Rib_EvisXr_Evi_Topology) GetGoName(yname string) string {
    if yname == "topology-id" { return "TopologyId" }
    if yname == "topology-name" { return "TopologyName" }
    if yname == "topology-type" { return "TopologyType" }
    return ""
}

func (topology *L2Rib_EvisXr_Evi_Topology) GetSegmentPath() string {
    return "topology"
}

func (topology *L2Rib_EvisXr_Evi_Topology) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (topology *L2Rib_EvisXr_Evi_Topology) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (topology *L2Rib_EvisXr_Evi_Topology) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["topology-id"] = topology.TopologyId
    leafs["topology-name"] = topology.TopologyName
    leafs["topology-type"] = topology.TopologyType
    return leafs
}

func (topology *L2Rib_EvisXr_Evi_Topology) GetBundleName() string { return "cisco_ios_xr" }

func (topology *L2Rib_EvisXr_Evi_Topology) GetYangName() string { return "topology" }

func (topology *L2Rib_EvisXr_Evi_Topology) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (topology *L2Rib_EvisXr_Evi_Topology) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (topology *L2Rib_EvisXr_Evi_Topology) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (topology *L2Rib_EvisXr_Evi_Topology) SetParent(parent types.Entity) { topology.parent = parent }

func (topology *L2Rib_EvisXr_Evi_Topology) GetParent() types.Entity { return topology.parent }

func (topology *L2Rib_EvisXr_Evi_Topology) GetParentYangName() string { return "evi" }

// L2Rib_ClientsDetails
// L2RIB detailed client table
type L2Rib_ClientsDetails struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // L2RIB clients detail information. The type is slice of
    // L2Rib_ClientsDetails_ClientsDetail.
    ClientsDetail []L2Rib_ClientsDetails_ClientsDetail
}

func (clientsDetails *L2Rib_ClientsDetails) GetFilter() yfilter.YFilter { return clientsDetails.YFilter }

func (clientsDetails *L2Rib_ClientsDetails) SetFilter(yf yfilter.YFilter) { clientsDetails.YFilter = yf }

func (clientsDetails *L2Rib_ClientsDetails) GetGoName(yname string) string {
    if yname == "clients-detail" { return "ClientsDetail" }
    return ""
}

func (clientsDetails *L2Rib_ClientsDetails) GetSegmentPath() string {
    return "clients-details"
}

func (clientsDetails *L2Rib_ClientsDetails) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "clients-detail" {
        for _, c := range clientsDetails.ClientsDetail {
            if clientsDetails.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := L2Rib_ClientsDetails_ClientsDetail{}
        clientsDetails.ClientsDetail = append(clientsDetails.ClientsDetail, child)
        return &clientsDetails.ClientsDetail[len(clientsDetails.ClientsDetail)-1]
    }
    return nil
}

func (clientsDetails *L2Rib_ClientsDetails) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range clientsDetails.ClientsDetail {
        children[clientsDetails.ClientsDetail[i].GetSegmentPath()] = &clientsDetails.ClientsDetail[i]
    }
    return children
}

func (clientsDetails *L2Rib_ClientsDetails) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (clientsDetails *L2Rib_ClientsDetails) GetBundleName() string { return "cisco_ios_xr" }

func (clientsDetails *L2Rib_ClientsDetails) GetYangName() string { return "clients-details" }

func (clientsDetails *L2Rib_ClientsDetails) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clientsDetails *L2Rib_ClientsDetails) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clientsDetails *L2Rib_ClientsDetails) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clientsDetails *L2Rib_ClientsDetails) SetParent(parent types.Entity) { clientsDetails.parent = parent }

func (clientsDetails *L2Rib_ClientsDetails) GetParent() types.Entity { return clientsDetails.parent }

func (clientsDetails *L2Rib_ClientsDetails) GetParentYangName() string { return "l2rib" }

// L2Rib_ClientsDetails_ClientsDetail
// L2RIB clients detail information
type L2Rib_ClientsDetails_ClientsDetail struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Client ID. The type is interface{} with range:
    // -2147483648..2147483647.
    ClientId interface{}

    // Number of Producers. The type is interface{} with range: 0..255.
    ProducerCount interface{}

    // Last Update Timestamp. The type is interface{} with range:
    // 0..18446744073709551615.
    LastUpdateTimestamp interface{}

    // Non-detail Client bag.
    Client L2Rib_ClientsDetails_ClientsDetail_Client

    // Registration table statistics.
    RegistrationTableStatistics L2Rib_ClientsDetails_ClientsDetail_RegistrationTableStatistics

    // List of Producers. The type is slice of
    // L2Rib_ClientsDetails_ClientsDetail_ProducerArray.
    ProducerArray []L2Rib_ClientsDetails_ClientsDetail_ProducerArray
}

func (clientsDetail *L2Rib_ClientsDetails_ClientsDetail) GetFilter() yfilter.YFilter { return clientsDetail.YFilter }

func (clientsDetail *L2Rib_ClientsDetails_ClientsDetail) SetFilter(yf yfilter.YFilter) { clientsDetail.YFilter = yf }

func (clientsDetail *L2Rib_ClientsDetails_ClientsDetail) GetGoName(yname string) string {
    if yname == "client-id" { return "ClientId" }
    if yname == "producer-count" { return "ProducerCount" }
    if yname == "last-update-timestamp" { return "LastUpdateTimestamp" }
    if yname == "client" { return "Client" }
    if yname == "registration-table-statistics" { return "RegistrationTableStatistics" }
    if yname == "producer-array" { return "ProducerArray" }
    return ""
}

func (clientsDetail *L2Rib_ClientsDetails_ClientsDetail) GetSegmentPath() string {
    return "clients-detail" + "[client-id='" + fmt.Sprintf("%v", clientsDetail.ClientId) + "']"
}

func (clientsDetail *L2Rib_ClientsDetails_ClientsDetail) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "client" {
        return &clientsDetail.Client
    }
    if childYangName == "registration-table-statistics" {
        return &clientsDetail.RegistrationTableStatistics
    }
    if childYangName == "producer-array" {
        for _, c := range clientsDetail.ProducerArray {
            if clientsDetail.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := L2Rib_ClientsDetails_ClientsDetail_ProducerArray{}
        clientsDetail.ProducerArray = append(clientsDetail.ProducerArray, child)
        return &clientsDetail.ProducerArray[len(clientsDetail.ProducerArray)-1]
    }
    return nil
}

func (clientsDetail *L2Rib_ClientsDetails_ClientsDetail) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["client"] = &clientsDetail.Client
    children["registration-table-statistics"] = &clientsDetail.RegistrationTableStatistics
    for i := range clientsDetail.ProducerArray {
        children[clientsDetail.ProducerArray[i].GetSegmentPath()] = &clientsDetail.ProducerArray[i]
    }
    return children
}

func (clientsDetail *L2Rib_ClientsDetails_ClientsDetail) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["client-id"] = clientsDetail.ClientId
    leafs["producer-count"] = clientsDetail.ProducerCount
    leafs["last-update-timestamp"] = clientsDetail.LastUpdateTimestamp
    return leafs
}

func (clientsDetail *L2Rib_ClientsDetails_ClientsDetail) GetBundleName() string { return "cisco_ios_xr" }

func (clientsDetail *L2Rib_ClientsDetails_ClientsDetail) GetYangName() string { return "clients-detail" }

func (clientsDetail *L2Rib_ClientsDetails_ClientsDetail) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clientsDetail *L2Rib_ClientsDetails_ClientsDetail) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clientsDetail *L2Rib_ClientsDetails_ClientsDetail) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clientsDetail *L2Rib_ClientsDetails_ClientsDetail) SetParent(parent types.Entity) { clientsDetail.parent = parent }

func (clientsDetail *L2Rib_ClientsDetails_ClientsDetail) GetParent() types.Entity { return clientsDetail.parent }

func (clientsDetail *L2Rib_ClientsDetails_ClientsDetail) GetParentYangName() string { return "clients-details" }

// L2Rib_ClientsDetails_ClientsDetail_Client
// Non-detail Client bag
type L2Rib_ClientsDetails_ClientsDetail_Client struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Client ID. The type is interface{} with range: 0..4294967295.
    ClientIdXr interface{}

    // Process ID. The type is interface{} with range: 0..4294967295.
    ProcessId interface{}

    // Node ID. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeId interface{}

    // Process Name. The type is string.
    ProcName interface{}

    // Process Suffix. The type is string.
    ProcSuffix interface{}
}

func (client *L2Rib_ClientsDetails_ClientsDetail_Client) GetFilter() yfilter.YFilter { return client.YFilter }

func (client *L2Rib_ClientsDetails_ClientsDetail_Client) SetFilter(yf yfilter.YFilter) { client.YFilter = yf }

func (client *L2Rib_ClientsDetails_ClientsDetail_Client) GetGoName(yname string) string {
    if yname == "client-id-xr" { return "ClientIdXr" }
    if yname == "process-id" { return "ProcessId" }
    if yname == "node-id" { return "NodeId" }
    if yname == "proc-name" { return "ProcName" }
    if yname == "proc-suffix" { return "ProcSuffix" }
    return ""
}

func (client *L2Rib_ClientsDetails_ClientsDetail_Client) GetSegmentPath() string {
    return "client"
}

func (client *L2Rib_ClientsDetails_ClientsDetail_Client) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (client *L2Rib_ClientsDetails_ClientsDetail_Client) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (client *L2Rib_ClientsDetails_ClientsDetail_Client) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["client-id-xr"] = client.ClientIdXr
    leafs["process-id"] = client.ProcessId
    leafs["node-id"] = client.NodeId
    leafs["proc-name"] = client.ProcName
    leafs["proc-suffix"] = client.ProcSuffix
    return leafs
}

func (client *L2Rib_ClientsDetails_ClientsDetail_Client) GetBundleName() string { return "cisco_ios_xr" }

func (client *L2Rib_ClientsDetails_ClientsDetail_Client) GetYangName() string { return "client" }

func (client *L2Rib_ClientsDetails_ClientsDetail_Client) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (client *L2Rib_ClientsDetails_ClientsDetail_Client) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (client *L2Rib_ClientsDetails_ClientsDetail_Client) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (client *L2Rib_ClientsDetails_ClientsDetail_Client) SetParent(parent types.Entity) { client.parent = parent }

func (client *L2Rib_ClientsDetails_ClientsDetail_Client) GetParent() types.Entity { return client.parent }

func (client *L2Rib_ClientsDetails_ClientsDetail_Client) GetParentYangName() string { return "clients-detail" }

// L2Rib_ClientsDetails_ClientsDetail_RegistrationTableStatistics
// Registration table statistics
type L2Rib_ClientsDetails_ClientsDetail_RegistrationTableStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Producer ID. The type is L2ribBagProducerId.
    ProducerId interface{}

    // Producer Name. The type is string.
    ProducerName interface{}

    // Statistics.
    Statistics L2Rib_ClientsDetails_ClientsDetail_RegistrationTableStatistics_Statistics
}

func (registrationTableStatistics *L2Rib_ClientsDetails_ClientsDetail_RegistrationTableStatistics) GetFilter() yfilter.YFilter { return registrationTableStatistics.YFilter }

func (registrationTableStatistics *L2Rib_ClientsDetails_ClientsDetail_RegistrationTableStatistics) SetFilter(yf yfilter.YFilter) { registrationTableStatistics.YFilter = yf }

func (registrationTableStatistics *L2Rib_ClientsDetails_ClientsDetail_RegistrationTableStatistics) GetGoName(yname string) string {
    if yname == "producer-id" { return "ProducerId" }
    if yname == "producer-name" { return "ProducerName" }
    if yname == "statistics" { return "Statistics" }
    return ""
}

func (registrationTableStatistics *L2Rib_ClientsDetails_ClientsDetail_RegistrationTableStatistics) GetSegmentPath() string {
    return "registration-table-statistics"
}

func (registrationTableStatistics *L2Rib_ClientsDetails_ClientsDetail_RegistrationTableStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "statistics" {
        return &registrationTableStatistics.Statistics
    }
    return nil
}

func (registrationTableStatistics *L2Rib_ClientsDetails_ClientsDetail_RegistrationTableStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["statistics"] = &registrationTableStatistics.Statistics
    return children
}

func (registrationTableStatistics *L2Rib_ClientsDetails_ClientsDetail_RegistrationTableStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["producer-id"] = registrationTableStatistics.ProducerId
    leafs["producer-name"] = registrationTableStatistics.ProducerName
    return leafs
}

func (registrationTableStatistics *L2Rib_ClientsDetails_ClientsDetail_RegistrationTableStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (registrationTableStatistics *L2Rib_ClientsDetails_ClientsDetail_RegistrationTableStatistics) GetYangName() string { return "registration-table-statistics" }

func (registrationTableStatistics *L2Rib_ClientsDetails_ClientsDetail_RegistrationTableStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (registrationTableStatistics *L2Rib_ClientsDetails_ClientsDetail_RegistrationTableStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (registrationTableStatistics *L2Rib_ClientsDetails_ClientsDetail_RegistrationTableStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (registrationTableStatistics *L2Rib_ClientsDetails_ClientsDetail_RegistrationTableStatistics) SetParent(parent types.Entity) { registrationTableStatistics.parent = parent }

func (registrationTableStatistics *L2Rib_ClientsDetails_ClientsDetail_RegistrationTableStatistics) GetParent() types.Entity { return registrationTableStatistics.parent }

func (registrationTableStatistics *L2Rib_ClientsDetails_ClientsDetail_RegistrationTableStatistics) GetParentYangName() string { return "clients-detail" }

// L2Rib_ClientsDetails_ClientsDetail_RegistrationTableStatistics_Statistics
// Statistics
type L2Rib_ClientsDetails_ClientsDetail_RegistrationTableStatistics_Statistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Memory Size. The type is interface{} with range: 0..4294967295.
    MemorySize interface{}

    // Number of Objects. The type is interface{} with range: 0..4294967295.
    ObjectCount interface{}

    // End of Interval Timestamp. The type is interface{} with range:
    // 0..18446744073709551615.
    EndofIntervalTs interface{}

    // Extended Counters. The type is slice of
    // L2Rib_ClientsDetails_ClientsDetail_RegistrationTableStatistics_Statistics_ExtendedCounter.
    ExtendedCounter []L2Rib_ClientsDetails_ClientsDetail_RegistrationTableStatistics_Statistics_ExtendedCounter
}

func (statistics *L2Rib_ClientsDetails_ClientsDetail_RegistrationTableStatistics_Statistics) GetFilter() yfilter.YFilter { return statistics.YFilter }

func (statistics *L2Rib_ClientsDetails_ClientsDetail_RegistrationTableStatistics_Statistics) SetFilter(yf yfilter.YFilter) { statistics.YFilter = yf }

func (statistics *L2Rib_ClientsDetails_ClientsDetail_RegistrationTableStatistics_Statistics) GetGoName(yname string) string {
    if yname == "memory-size" { return "MemorySize" }
    if yname == "object-count" { return "ObjectCount" }
    if yname == "endof-interval-ts" { return "EndofIntervalTs" }
    if yname == "extended-counter" { return "ExtendedCounter" }
    return ""
}

func (statistics *L2Rib_ClientsDetails_ClientsDetail_RegistrationTableStatistics_Statistics) GetSegmentPath() string {
    return "statistics"
}

func (statistics *L2Rib_ClientsDetails_ClientsDetail_RegistrationTableStatistics_Statistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "extended-counter" {
        for _, c := range statistics.ExtendedCounter {
            if statistics.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := L2Rib_ClientsDetails_ClientsDetail_RegistrationTableStatistics_Statistics_ExtendedCounter{}
        statistics.ExtendedCounter = append(statistics.ExtendedCounter, child)
        return &statistics.ExtendedCounter[len(statistics.ExtendedCounter)-1]
    }
    return nil
}

func (statistics *L2Rib_ClientsDetails_ClientsDetail_RegistrationTableStatistics_Statistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range statistics.ExtendedCounter {
        children[statistics.ExtendedCounter[i].GetSegmentPath()] = &statistics.ExtendedCounter[i]
    }
    return children
}

func (statistics *L2Rib_ClientsDetails_ClientsDetail_RegistrationTableStatistics_Statistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["memory-size"] = statistics.MemorySize
    leafs["object-count"] = statistics.ObjectCount
    leafs["endof-interval-ts"] = statistics.EndofIntervalTs
    return leafs
}

func (statistics *L2Rib_ClientsDetails_ClientsDetail_RegistrationTableStatistics_Statistics) GetBundleName() string { return "cisco_ios_xr" }

func (statistics *L2Rib_ClientsDetails_ClientsDetail_RegistrationTableStatistics_Statistics) GetYangName() string { return "statistics" }

func (statistics *L2Rib_ClientsDetails_ClientsDetail_RegistrationTableStatistics_Statistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistics *L2Rib_ClientsDetails_ClientsDetail_RegistrationTableStatistics_Statistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistics *L2Rib_ClientsDetails_ClientsDetail_RegistrationTableStatistics_Statistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistics *L2Rib_ClientsDetails_ClientsDetail_RegistrationTableStatistics_Statistics) SetParent(parent types.Entity) { statistics.parent = parent }

func (statistics *L2Rib_ClientsDetails_ClientsDetail_RegistrationTableStatistics_Statistics) GetParent() types.Entity { return statistics.parent }

func (statistics *L2Rib_ClientsDetails_ClientsDetail_RegistrationTableStatistics_Statistics) GetParentYangName() string { return "registration-table-statistics" }

// L2Rib_ClientsDetails_ClientsDetail_RegistrationTableStatistics_Statistics_ExtendedCounter
// Extended Counters
type L2Rib_ClientsDetails_ClientsDetail_RegistrationTableStatistics_Statistics_ExtendedCounter struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // CounterType. The type is interface{} with range: 0..255.
    CounterType interface{}

    // CounterName. The type is string.
    CounterName interface{}

    // Real-clock timestamp in msec of first event. The type is interface{} with
    // range: 0..18446744073709551615.
    L2RbFirstEventTs interface{}

    // Real-clock timestamp in msec of last event. The type is interface{} with
    // range: 0..18446744073709551615.
    L2RbLastEventTs interface{}

    // number of events in interval. The type is interface{} with range:
    // 0..4294967295.
    L2RbIntervalEventCount interface{}

    // total number of events. The type is interface{} with range: 0..4294967295.
    L2RbTotalEventCount interface{}
}

func (extendedCounter *L2Rib_ClientsDetails_ClientsDetail_RegistrationTableStatistics_Statistics_ExtendedCounter) GetFilter() yfilter.YFilter { return extendedCounter.YFilter }

func (extendedCounter *L2Rib_ClientsDetails_ClientsDetail_RegistrationTableStatistics_Statistics_ExtendedCounter) SetFilter(yf yfilter.YFilter) { extendedCounter.YFilter = yf }

func (extendedCounter *L2Rib_ClientsDetails_ClientsDetail_RegistrationTableStatistics_Statistics_ExtendedCounter) GetGoName(yname string) string {
    if yname == "counter-type" { return "CounterType" }
    if yname == "counter-name" { return "CounterName" }
    if yname == "l2rb-first-event-ts" { return "L2RbFirstEventTs" }
    if yname == "l2rb-last-event-ts" { return "L2RbLastEventTs" }
    if yname == "l2rb-interval-event-count" { return "L2RbIntervalEventCount" }
    if yname == "l2rb-total-event-count" { return "L2RbTotalEventCount" }
    return ""
}

func (extendedCounter *L2Rib_ClientsDetails_ClientsDetail_RegistrationTableStatistics_Statistics_ExtendedCounter) GetSegmentPath() string {
    return "extended-counter"
}

func (extendedCounter *L2Rib_ClientsDetails_ClientsDetail_RegistrationTableStatistics_Statistics_ExtendedCounter) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (extendedCounter *L2Rib_ClientsDetails_ClientsDetail_RegistrationTableStatistics_Statistics_ExtendedCounter) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (extendedCounter *L2Rib_ClientsDetails_ClientsDetail_RegistrationTableStatistics_Statistics_ExtendedCounter) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["counter-type"] = extendedCounter.CounterType
    leafs["counter-name"] = extendedCounter.CounterName
    leafs["l2rb-first-event-ts"] = extendedCounter.L2RbFirstEventTs
    leafs["l2rb-last-event-ts"] = extendedCounter.L2RbLastEventTs
    leafs["l2rb-interval-event-count"] = extendedCounter.L2RbIntervalEventCount
    leafs["l2rb-total-event-count"] = extendedCounter.L2RbTotalEventCount
    return leafs
}

func (extendedCounter *L2Rib_ClientsDetails_ClientsDetail_RegistrationTableStatistics_Statistics_ExtendedCounter) GetBundleName() string { return "cisco_ios_xr" }

func (extendedCounter *L2Rib_ClientsDetails_ClientsDetail_RegistrationTableStatistics_Statistics_ExtendedCounter) GetYangName() string { return "extended-counter" }

func (extendedCounter *L2Rib_ClientsDetails_ClientsDetail_RegistrationTableStatistics_Statistics_ExtendedCounter) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (extendedCounter *L2Rib_ClientsDetails_ClientsDetail_RegistrationTableStatistics_Statistics_ExtendedCounter) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (extendedCounter *L2Rib_ClientsDetails_ClientsDetail_RegistrationTableStatistics_Statistics_ExtendedCounter) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (extendedCounter *L2Rib_ClientsDetails_ClientsDetail_RegistrationTableStatistics_Statistics_ExtendedCounter) SetParent(parent types.Entity) { extendedCounter.parent = parent }

func (extendedCounter *L2Rib_ClientsDetails_ClientsDetail_RegistrationTableStatistics_Statistics_ExtendedCounter) GetParent() types.Entity { return extendedCounter.parent }

func (extendedCounter *L2Rib_ClientsDetails_ClientsDetail_RegistrationTableStatistics_Statistics_ExtendedCounter) GetParentYangName() string { return "statistics" }

// L2Rib_ClientsDetails_ClientsDetail_ProducerArray
// List of Producers
type L2Rib_ClientsDetails_ClientsDetail_ProducerArray struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Object Type. The type is L2ribBagObj.
    ObjectType interface{}

    // Producer ID. The type is L2ribBagProducerId.
    ProducerId interface{}

    // Producer Name. The type is string.
    ProducerName interface{}

    // Admin Distance. The type is interface{} with range: 0..4294967295.
    AdminDistance interface{}

    // Purge Time. The type is interface{} with range: 0..4294967295.
    PurgeTime interface{}
}

func (producerArray *L2Rib_ClientsDetails_ClientsDetail_ProducerArray) GetFilter() yfilter.YFilter { return producerArray.YFilter }

func (producerArray *L2Rib_ClientsDetails_ClientsDetail_ProducerArray) SetFilter(yf yfilter.YFilter) { producerArray.YFilter = yf }

func (producerArray *L2Rib_ClientsDetails_ClientsDetail_ProducerArray) GetGoName(yname string) string {
    if yname == "object-type" { return "ObjectType" }
    if yname == "producer-id" { return "ProducerId" }
    if yname == "producer-name" { return "ProducerName" }
    if yname == "admin-distance" { return "AdminDistance" }
    if yname == "purge-time" { return "PurgeTime" }
    return ""
}

func (producerArray *L2Rib_ClientsDetails_ClientsDetail_ProducerArray) GetSegmentPath() string {
    return "producer-array"
}

func (producerArray *L2Rib_ClientsDetails_ClientsDetail_ProducerArray) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (producerArray *L2Rib_ClientsDetails_ClientsDetail_ProducerArray) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (producerArray *L2Rib_ClientsDetails_ClientsDetail_ProducerArray) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object-type"] = producerArray.ObjectType
    leafs["producer-id"] = producerArray.ProducerId
    leafs["producer-name"] = producerArray.ProducerName
    leafs["admin-distance"] = producerArray.AdminDistance
    leafs["purge-time"] = producerArray.PurgeTime
    return leafs
}

func (producerArray *L2Rib_ClientsDetails_ClientsDetail_ProducerArray) GetBundleName() string { return "cisco_ios_xr" }

func (producerArray *L2Rib_ClientsDetails_ClientsDetail_ProducerArray) GetYangName() string { return "producer-array" }

func (producerArray *L2Rib_ClientsDetails_ClientsDetail_ProducerArray) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (producerArray *L2Rib_ClientsDetails_ClientsDetail_ProducerArray) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (producerArray *L2Rib_ClientsDetails_ClientsDetail_ProducerArray) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (producerArray *L2Rib_ClientsDetails_ClientsDetail_ProducerArray) SetParent(parent types.Entity) { producerArray.parent = parent }

func (producerArray *L2Rib_ClientsDetails_ClientsDetail_ProducerArray) GetParent() types.Entity { return producerArray.parent }

func (producerArray *L2Rib_ClientsDetails_ClientsDetail_ProducerArray) GetParentYangName() string { return "clients-detail" }

// L2Rib_EviChildTables
// Container for all EVI Child Tables
type L2Rib_EviChildTables struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // L2RIB EVPN EVI MAC IP Detail table.
    MacipDetails L2Rib_EviChildTables_MacipDetails

    // L2RIB EVPN EVI MAC IP table.
    MacIps L2Rib_EviChildTables_MacIps

    // L2RIB EVPN EVI MAC table.
    Macs L2Rib_EviChildTables_Macs

    // L2RIB EVPN EVI MAC Detail table.
    MacDetails L2Rib_EviChildTables_MacDetails
}

func (eviChildTables *L2Rib_EviChildTables) GetFilter() yfilter.YFilter { return eviChildTables.YFilter }

func (eviChildTables *L2Rib_EviChildTables) SetFilter(yf yfilter.YFilter) { eviChildTables.YFilter = yf }

func (eviChildTables *L2Rib_EviChildTables) GetGoName(yname string) string {
    if yname == "macip-details" { return "MacipDetails" }
    if yname == "mac-ips" { return "MacIps" }
    if yname == "macs" { return "Macs" }
    if yname == "mac-details" { return "MacDetails" }
    return ""
}

func (eviChildTables *L2Rib_EviChildTables) GetSegmentPath() string {
    return "evi-child-tables"
}

func (eviChildTables *L2Rib_EviChildTables) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "macip-details" {
        return &eviChildTables.MacipDetails
    }
    if childYangName == "mac-ips" {
        return &eviChildTables.MacIps
    }
    if childYangName == "macs" {
        return &eviChildTables.Macs
    }
    if childYangName == "mac-details" {
        return &eviChildTables.MacDetails
    }
    return nil
}

func (eviChildTables *L2Rib_EviChildTables) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["macip-details"] = &eviChildTables.MacipDetails
    children["mac-ips"] = &eviChildTables.MacIps
    children["macs"] = &eviChildTables.Macs
    children["mac-details"] = &eviChildTables.MacDetails
    return children
}

func (eviChildTables *L2Rib_EviChildTables) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (eviChildTables *L2Rib_EviChildTables) GetBundleName() string { return "cisco_ios_xr" }

func (eviChildTables *L2Rib_EviChildTables) GetYangName() string { return "evi-child-tables" }

func (eviChildTables *L2Rib_EviChildTables) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (eviChildTables *L2Rib_EviChildTables) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (eviChildTables *L2Rib_EviChildTables) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (eviChildTables *L2Rib_EviChildTables) SetParent(parent types.Entity) { eviChildTables.parent = parent }

func (eviChildTables *L2Rib_EviChildTables) GetParent() types.Entity { return eviChildTables.parent }

func (eviChildTables *L2Rib_EviChildTables) GetParentYangName() string { return "l2rib" }

// L2Rib_EviChildTables_MacipDetails
// L2RIB EVPN EVI MAC IP Detail table
type L2Rib_EviChildTables_MacipDetails struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // L2RIB EVPN MAC IP Detail table. The type is slice of
    // L2Rib_EviChildTables_MacipDetails_MacipDetail.
    MacipDetail []L2Rib_EviChildTables_MacipDetails_MacipDetail
}

func (macipDetails *L2Rib_EviChildTables_MacipDetails) GetFilter() yfilter.YFilter { return macipDetails.YFilter }

func (macipDetails *L2Rib_EviChildTables_MacipDetails) SetFilter(yf yfilter.YFilter) { macipDetails.YFilter = yf }

func (macipDetails *L2Rib_EviChildTables_MacipDetails) GetGoName(yname string) string {
    if yname == "macip-detail" { return "MacipDetail" }
    return ""
}

func (macipDetails *L2Rib_EviChildTables_MacipDetails) GetSegmentPath() string {
    return "macip-details"
}

func (macipDetails *L2Rib_EviChildTables_MacipDetails) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "macip-detail" {
        for _, c := range macipDetails.MacipDetail {
            if macipDetails.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := L2Rib_EviChildTables_MacipDetails_MacipDetail{}
        macipDetails.MacipDetail = append(macipDetails.MacipDetail, child)
        return &macipDetails.MacipDetail[len(macipDetails.MacipDetail)-1]
    }
    return nil
}

func (macipDetails *L2Rib_EviChildTables_MacipDetails) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range macipDetails.MacipDetail {
        children[macipDetails.MacipDetail[i].GetSegmentPath()] = &macipDetails.MacipDetail[i]
    }
    return children
}

func (macipDetails *L2Rib_EviChildTables_MacipDetails) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (macipDetails *L2Rib_EviChildTables_MacipDetails) GetBundleName() string { return "cisco_ios_xr" }

func (macipDetails *L2Rib_EviChildTables_MacipDetails) GetYangName() string { return "macip-details" }

func (macipDetails *L2Rib_EviChildTables_MacipDetails) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (macipDetails *L2Rib_EviChildTables_MacipDetails) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (macipDetails *L2Rib_EviChildTables_MacipDetails) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (macipDetails *L2Rib_EviChildTables_MacipDetails) SetParent(parent types.Entity) { macipDetails.parent = parent }

func (macipDetails *L2Rib_EviChildTables_MacipDetails) GetParent() types.Entity { return macipDetails.parent }

func (macipDetails *L2Rib_EviChildTables_MacipDetails) GetParentYangName() string { return "evi-child-tables" }

// L2Rib_EviChildTables_MacipDetails_MacipDetail
// L2RIB EVPN MAC IP Detail table
type L2Rib_EviChildTables_MacipDetails_MacipDetail struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // EVPN ID. The type is interface{} with range: -2147483648..2147483647.
    Evi interface{}

    // Tag ID. The type is interface{} with range: -2147483648..2147483647.
    TagId interface{}

    // MAC IP Address. The type is string with length: 1..15.
    MacAddr interface{}

    // IP Address. The type is string with length: 1..15.
    IpAddr interface{}

    // Admin distance. The type is interface{} with range:
    // -2147483648..2147483647.
    AdminDist interface{}

    // Producer ID. The type is interface{} with range: -2147483648..2147483647.
    ProdId interface{}

    // MAC-IP route sequence Number. The type is interface{} with range:
    // 0..4294967295.
    SequenceNumber interface{}

    // MAC-IP route flags. The type is interface{} with range: 0..4294967295.
    Flags interface{}

    // SOO. The type is interface{} with range: 0..4294967295.
    Soo interface{}

    // Last Update Timestamp. The type is interface{} with range:
    // 0..18446744073709551615.
    LastUpdateTimestamp interface{}

    // MAC-IP Route.
    MacIpRoute L2Rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute

    // Mac-IP Route Opaque Data TLV.
    RtTlv L2Rib_EviChildTables_MacipDetails_MacipDetail_RtTlv

    // Mac-IP Route Opaque NH TLV.
    NhTlv L2Rib_EviChildTables_MacipDetails_MacipDetail_NhTlv
}

func (macipDetail *L2Rib_EviChildTables_MacipDetails_MacipDetail) GetFilter() yfilter.YFilter { return macipDetail.YFilter }

func (macipDetail *L2Rib_EviChildTables_MacipDetails_MacipDetail) SetFilter(yf yfilter.YFilter) { macipDetail.YFilter = yf }

func (macipDetail *L2Rib_EviChildTables_MacipDetails_MacipDetail) GetGoName(yname string) string {
    if yname == "evi" { return "Evi" }
    if yname == "tag-id" { return "TagId" }
    if yname == "mac-addr" { return "MacAddr" }
    if yname == "ip-addr" { return "IpAddr" }
    if yname == "admin-dist" { return "AdminDist" }
    if yname == "prod-id" { return "ProdId" }
    if yname == "sequence-number" { return "SequenceNumber" }
    if yname == "flags" { return "Flags" }
    if yname == "soo" { return "Soo" }
    if yname == "last-update-timestamp" { return "LastUpdateTimestamp" }
    if yname == "mac-ip-route" { return "MacIpRoute" }
    if yname == "rt-tlv" { return "RtTlv" }
    if yname == "nh-tlv" { return "NhTlv" }
    return ""
}

func (macipDetail *L2Rib_EviChildTables_MacipDetails_MacipDetail) GetSegmentPath() string {
    return "macip-detail"
}

func (macipDetail *L2Rib_EviChildTables_MacipDetails_MacipDetail) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mac-ip-route" {
        return &macipDetail.MacIpRoute
    }
    if childYangName == "rt-tlv" {
        return &macipDetail.RtTlv
    }
    if childYangName == "nh-tlv" {
        return &macipDetail.NhTlv
    }
    return nil
}

func (macipDetail *L2Rib_EviChildTables_MacipDetails_MacipDetail) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["mac-ip-route"] = &macipDetail.MacIpRoute
    children["rt-tlv"] = &macipDetail.RtTlv
    children["nh-tlv"] = &macipDetail.NhTlv
    return children
}

func (macipDetail *L2Rib_EviChildTables_MacipDetails_MacipDetail) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["evi"] = macipDetail.Evi
    leafs["tag-id"] = macipDetail.TagId
    leafs["mac-addr"] = macipDetail.MacAddr
    leafs["ip-addr"] = macipDetail.IpAddr
    leafs["admin-dist"] = macipDetail.AdminDist
    leafs["prod-id"] = macipDetail.ProdId
    leafs["sequence-number"] = macipDetail.SequenceNumber
    leafs["flags"] = macipDetail.Flags
    leafs["soo"] = macipDetail.Soo
    leafs["last-update-timestamp"] = macipDetail.LastUpdateTimestamp
    return leafs
}

func (macipDetail *L2Rib_EviChildTables_MacipDetails_MacipDetail) GetBundleName() string { return "cisco_ios_xr" }

func (macipDetail *L2Rib_EviChildTables_MacipDetails_MacipDetail) GetYangName() string { return "macip-detail" }

func (macipDetail *L2Rib_EviChildTables_MacipDetails_MacipDetail) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (macipDetail *L2Rib_EviChildTables_MacipDetails_MacipDetail) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (macipDetail *L2Rib_EviChildTables_MacipDetails_MacipDetail) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (macipDetail *L2Rib_EviChildTables_MacipDetails_MacipDetail) SetParent(parent types.Entity) { macipDetail.parent = parent }

func (macipDetail *L2Rib_EviChildTables_MacipDetails_MacipDetail) GetParent() types.Entity { return macipDetail.parent }

func (macipDetail *L2Rib_EviChildTables_MacipDetails_MacipDetail) GetParentYangName() string { return "macip-details" }

// L2Rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute
// MAC-IP Route
type L2Rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MAC Address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddress interface{}

    // IP Address. The type is string.
    IpAddress interface{}

    // Admin Distance. The type is interface{} with range: 0..255.
    AdminDistance interface{}

    // Producer ID. The type is interface{} with range: 0..255.
    ProducerId interface{}

    // Topology ID. The type is interface{} with range: 0..4294967295.
    TopologyId interface{}

    // Next Hop.
    NextHop L2Rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute_NextHop
}

func (macIpRoute *L2Rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute) GetFilter() yfilter.YFilter { return macIpRoute.YFilter }

func (macIpRoute *L2Rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute) SetFilter(yf yfilter.YFilter) { macIpRoute.YFilter = yf }

func (macIpRoute *L2Rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute) GetGoName(yname string) string {
    if yname == "mac-address" { return "MacAddress" }
    if yname == "ip-address" { return "IpAddress" }
    if yname == "admin-distance" { return "AdminDistance" }
    if yname == "producer-id" { return "ProducerId" }
    if yname == "topology-id" { return "TopologyId" }
    if yname == "next-hop" { return "NextHop" }
    return ""
}

func (macIpRoute *L2Rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute) GetSegmentPath() string {
    return "mac-ip-route"
}

func (macIpRoute *L2Rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "next-hop" {
        return &macIpRoute.NextHop
    }
    return nil
}

func (macIpRoute *L2Rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["next-hop"] = &macIpRoute.NextHop
    return children
}

func (macIpRoute *L2Rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mac-address"] = macIpRoute.MacAddress
    leafs["ip-address"] = macIpRoute.IpAddress
    leafs["admin-distance"] = macIpRoute.AdminDistance
    leafs["producer-id"] = macIpRoute.ProducerId
    leafs["topology-id"] = macIpRoute.TopologyId
    return leafs
}

func (macIpRoute *L2Rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute) GetBundleName() string { return "cisco_ios_xr" }

func (macIpRoute *L2Rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute) GetYangName() string { return "mac-ip-route" }

func (macIpRoute *L2Rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (macIpRoute *L2Rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (macIpRoute *L2Rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (macIpRoute *L2Rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute) SetParent(parent types.Entity) { macIpRoute.parent = parent }

func (macIpRoute *L2Rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute) GetParent() types.Entity { return macIpRoute.parent }

func (macIpRoute *L2Rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute) GetParentYangName() string { return "macip-detail" }

// L2Rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute_NextHop
// Next Hop
type L2Rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute_NextHop struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Next-hop TOPOLOGY ID. The type is interface{} with range: 0..4294967295.
    TopologyId interface{}

    // Next-hop flags. The type is interface{} with range: 0..65535.
    Flags interface{}

    // Next hop.
    NextHop L2Rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute_NextHop_NextHop
}

func (nextHop *L2Rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute_NextHop) GetFilter() yfilter.YFilter { return nextHop.YFilter }

func (nextHop *L2Rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute_NextHop) SetFilter(yf yfilter.YFilter) { nextHop.YFilter = yf }

func (nextHop *L2Rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute_NextHop) GetGoName(yname string) string {
    if yname == "topology-id" { return "TopologyId" }
    if yname == "flags" { return "Flags" }
    if yname == "next-hop" { return "NextHop" }
    return ""
}

func (nextHop *L2Rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute_NextHop) GetSegmentPath() string {
    return "next-hop"
}

func (nextHop *L2Rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute_NextHop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "next-hop" {
        return &nextHop.NextHop
    }
    return nil
}

func (nextHop *L2Rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute_NextHop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["next-hop"] = &nextHop.NextHop
    return children
}

func (nextHop *L2Rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute_NextHop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["topology-id"] = nextHop.TopologyId
    leafs["flags"] = nextHop.Flags
    return leafs
}

func (nextHop *L2Rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute_NextHop) GetBundleName() string { return "cisco_ios_xr" }

func (nextHop *L2Rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute_NextHop) GetYangName() string { return "next-hop" }

func (nextHop *L2Rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute_NextHop) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nextHop *L2Rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute_NextHop) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nextHop *L2Rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute_NextHop) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nextHop *L2Rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute_NextHop) SetParent(parent types.Entity) { nextHop.parent = parent }

func (nextHop *L2Rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute_NextHop) GetParent() types.Entity { return nextHop.parent }

func (nextHop *L2Rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute_NextHop) GetParentYangName() string { return "mac-ip-route" }

// L2Rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute_NextHop_NextHop
// Next hop
type L2Rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute_NextHop_NextHop struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Type. The type is L2ribNextHop.
    Type interface{}

    // IPV4 address Next Hop. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPV6 address Next Hop. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}

    // MAC address Next Hop. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    Mac interface{}

    // Intefrace Handle Next Hop. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceHandle interface{}

    // XID Next Hop. The type is interface{} with range: 0..4294967295.
    Xid interface{}

    // Labeled Next Hop.
    Labeled L2Rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute_NextHop_NextHop_Labeled
}

func (nextHop *L2Rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute_NextHop_NextHop) GetFilter() yfilter.YFilter { return nextHop.YFilter }

func (nextHop *L2Rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute_NextHop_NextHop) SetFilter(yf yfilter.YFilter) { nextHop.YFilter = yf }

func (nextHop *L2Rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute_NextHop_NextHop) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "ipv4" { return "Ipv4" }
    if yname == "ipv6" { return "Ipv6" }
    if yname == "mac" { return "Mac" }
    if yname == "interface-handle" { return "InterfaceHandle" }
    if yname == "xid" { return "Xid" }
    if yname == "labeled" { return "Labeled" }
    return ""
}

func (nextHop *L2Rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute_NextHop_NextHop) GetSegmentPath() string {
    return "next-hop"
}

func (nextHop *L2Rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute_NextHop_NextHop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "labeled" {
        return &nextHop.Labeled
    }
    return nil
}

func (nextHop *L2Rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute_NextHop_NextHop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["labeled"] = &nextHop.Labeled
    return children
}

func (nextHop *L2Rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute_NextHop_NextHop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = nextHop.Type
    leafs["ipv4"] = nextHop.Ipv4
    leafs["ipv6"] = nextHop.Ipv6
    leafs["mac"] = nextHop.Mac
    leafs["interface-handle"] = nextHop.InterfaceHandle
    leafs["xid"] = nextHop.Xid
    return leafs
}

func (nextHop *L2Rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute_NextHop_NextHop) GetBundleName() string { return "cisco_ios_xr" }

func (nextHop *L2Rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute_NextHop_NextHop) GetYangName() string { return "next-hop" }

func (nextHop *L2Rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute_NextHop_NextHop) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nextHop *L2Rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute_NextHop_NextHop) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nextHop *L2Rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute_NextHop_NextHop) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nextHop *L2Rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute_NextHop_NextHop) SetParent(parent types.Entity) { nextHop.parent = parent }

func (nextHop *L2Rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute_NextHop_NextHop) GetParent() types.Entity { return nextHop.parent }

func (nextHop *L2Rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute_NextHop_NextHop) GetParentYangName() string { return "next-hop" }

// L2Rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute_NextHop_NextHop_Labeled
// Labeled Next Hop
type L2Rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute_NextHop_NextHop_Labeled struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // L2RIB Address Family. The type is L2ribAfi.
    AddressFamily interface{}

    // IP Address (V6 Format). The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IpAddress interface{}

    // Label. The type is interface{} with range: 0..4294967295.
    Label interface{}

    // Internal Label. The type is bool.
    Internal interface{}
}

func (labeled *L2Rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute_NextHop_NextHop_Labeled) GetFilter() yfilter.YFilter { return labeled.YFilter }

func (labeled *L2Rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute_NextHop_NextHop_Labeled) SetFilter(yf yfilter.YFilter) { labeled.YFilter = yf }

func (labeled *L2Rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute_NextHop_NextHop_Labeled) GetGoName(yname string) string {
    if yname == "address-family" { return "AddressFamily" }
    if yname == "ip-address" { return "IpAddress" }
    if yname == "label" { return "Label" }
    if yname == "internal" { return "Internal" }
    return ""
}

func (labeled *L2Rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute_NextHop_NextHop_Labeled) GetSegmentPath() string {
    return "labeled"
}

func (labeled *L2Rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute_NextHop_NextHop_Labeled) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (labeled *L2Rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute_NextHop_NextHop_Labeled) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (labeled *L2Rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute_NextHop_NextHop_Labeled) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address-family"] = labeled.AddressFamily
    leafs["ip-address"] = labeled.IpAddress
    leafs["label"] = labeled.Label
    leafs["internal"] = labeled.Internal
    return leafs
}

func (labeled *L2Rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute_NextHop_NextHop_Labeled) GetBundleName() string { return "cisco_ios_xr" }

func (labeled *L2Rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute_NextHop_NextHop_Labeled) GetYangName() string { return "labeled" }

func (labeled *L2Rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute_NextHop_NextHop_Labeled) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (labeled *L2Rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute_NextHop_NextHop_Labeled) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (labeled *L2Rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute_NextHop_NextHop_Labeled) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (labeled *L2Rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute_NextHop_NextHop_Labeled) SetParent(parent types.Entity) { labeled.parent = parent }

func (labeled *L2Rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute_NextHop_NextHop_Labeled) GetParent() types.Entity { return labeled.parent }

func (labeled *L2Rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute_NextHop_NextHop_Labeled) GetParentYangName() string { return "next-hop" }

// L2Rib_EviChildTables_MacipDetails_MacipDetail_RtTlv
// Mac-IP Route Opaque Data TLV
type L2Rib_EviChildTables_MacipDetails_MacipDetail_RtTlv struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TLV Type. The type is interface{} with range: 0..65535.
    TlvType interface{}

    // TLV Length. The type is interface{} with range: 0..65535.
    TlvLen interface{}

    // TLV Value. The type is slice of interface{} with range: 0..255.
    TlvVal []interface{}
}

func (rtTlv *L2Rib_EviChildTables_MacipDetails_MacipDetail_RtTlv) GetFilter() yfilter.YFilter { return rtTlv.YFilter }

func (rtTlv *L2Rib_EviChildTables_MacipDetails_MacipDetail_RtTlv) SetFilter(yf yfilter.YFilter) { rtTlv.YFilter = yf }

func (rtTlv *L2Rib_EviChildTables_MacipDetails_MacipDetail_RtTlv) GetGoName(yname string) string {
    if yname == "tlv-type" { return "TlvType" }
    if yname == "tlv-len" { return "TlvLen" }
    if yname == "tlv-val" { return "TlvVal" }
    return ""
}

func (rtTlv *L2Rib_EviChildTables_MacipDetails_MacipDetail_RtTlv) GetSegmentPath() string {
    return "rt-tlv"
}

func (rtTlv *L2Rib_EviChildTables_MacipDetails_MacipDetail_RtTlv) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rtTlv *L2Rib_EviChildTables_MacipDetails_MacipDetail_RtTlv) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rtTlv *L2Rib_EviChildTables_MacipDetails_MacipDetail_RtTlv) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tlv-type"] = rtTlv.TlvType
    leafs["tlv-len"] = rtTlv.TlvLen
    leafs["tlv-val"] = rtTlv.TlvVal
    return leafs
}

func (rtTlv *L2Rib_EviChildTables_MacipDetails_MacipDetail_RtTlv) GetBundleName() string { return "cisco_ios_xr" }

func (rtTlv *L2Rib_EviChildTables_MacipDetails_MacipDetail_RtTlv) GetYangName() string { return "rt-tlv" }

func (rtTlv *L2Rib_EviChildTables_MacipDetails_MacipDetail_RtTlv) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rtTlv *L2Rib_EviChildTables_MacipDetails_MacipDetail_RtTlv) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rtTlv *L2Rib_EviChildTables_MacipDetails_MacipDetail_RtTlv) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rtTlv *L2Rib_EviChildTables_MacipDetails_MacipDetail_RtTlv) SetParent(parent types.Entity) { rtTlv.parent = parent }

func (rtTlv *L2Rib_EviChildTables_MacipDetails_MacipDetail_RtTlv) GetParent() types.Entity { return rtTlv.parent }

func (rtTlv *L2Rib_EviChildTables_MacipDetails_MacipDetail_RtTlv) GetParentYangName() string { return "macip-detail" }

// L2Rib_EviChildTables_MacipDetails_MacipDetail_NhTlv
// Mac-IP Route Opaque NH TLV
type L2Rib_EviChildTables_MacipDetails_MacipDetail_NhTlv struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TLV Type. The type is interface{} with range: 0..65535.
    TlvType interface{}

    // TLV Length. The type is interface{} with range: 0..65535.
    TlvLen interface{}

    // TLV Value. The type is slice of interface{} with range: 0..255.
    TlvVal []interface{}
}

func (nhTlv *L2Rib_EviChildTables_MacipDetails_MacipDetail_NhTlv) GetFilter() yfilter.YFilter { return nhTlv.YFilter }

func (nhTlv *L2Rib_EviChildTables_MacipDetails_MacipDetail_NhTlv) SetFilter(yf yfilter.YFilter) { nhTlv.YFilter = yf }

func (nhTlv *L2Rib_EviChildTables_MacipDetails_MacipDetail_NhTlv) GetGoName(yname string) string {
    if yname == "tlv-type" { return "TlvType" }
    if yname == "tlv-len" { return "TlvLen" }
    if yname == "tlv-val" { return "TlvVal" }
    return ""
}

func (nhTlv *L2Rib_EviChildTables_MacipDetails_MacipDetail_NhTlv) GetSegmentPath() string {
    return "nh-tlv"
}

func (nhTlv *L2Rib_EviChildTables_MacipDetails_MacipDetail_NhTlv) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nhTlv *L2Rib_EviChildTables_MacipDetails_MacipDetail_NhTlv) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nhTlv *L2Rib_EviChildTables_MacipDetails_MacipDetail_NhTlv) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tlv-type"] = nhTlv.TlvType
    leafs["tlv-len"] = nhTlv.TlvLen
    leafs["tlv-val"] = nhTlv.TlvVal
    return leafs
}

func (nhTlv *L2Rib_EviChildTables_MacipDetails_MacipDetail_NhTlv) GetBundleName() string { return "cisco_ios_xr" }

func (nhTlv *L2Rib_EviChildTables_MacipDetails_MacipDetail_NhTlv) GetYangName() string { return "nh-tlv" }

func (nhTlv *L2Rib_EviChildTables_MacipDetails_MacipDetail_NhTlv) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nhTlv *L2Rib_EviChildTables_MacipDetails_MacipDetail_NhTlv) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nhTlv *L2Rib_EviChildTables_MacipDetails_MacipDetail_NhTlv) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nhTlv *L2Rib_EviChildTables_MacipDetails_MacipDetail_NhTlv) SetParent(parent types.Entity) { nhTlv.parent = parent }

func (nhTlv *L2Rib_EviChildTables_MacipDetails_MacipDetail_NhTlv) GetParent() types.Entity { return nhTlv.parent }

func (nhTlv *L2Rib_EviChildTables_MacipDetails_MacipDetail_NhTlv) GetParentYangName() string { return "macip-detail" }

// L2Rib_EviChildTables_MacIps
// L2RIB EVPN EVI MAC IP table
type L2Rib_EviChildTables_MacIps struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // L2RIB EVPN MAC IP table. The type is slice of
    // L2Rib_EviChildTables_MacIps_MacIp.
    MacIp []L2Rib_EviChildTables_MacIps_MacIp
}

func (macIps *L2Rib_EviChildTables_MacIps) GetFilter() yfilter.YFilter { return macIps.YFilter }

func (macIps *L2Rib_EviChildTables_MacIps) SetFilter(yf yfilter.YFilter) { macIps.YFilter = yf }

func (macIps *L2Rib_EviChildTables_MacIps) GetGoName(yname string) string {
    if yname == "mac-ip" { return "MacIp" }
    return ""
}

func (macIps *L2Rib_EviChildTables_MacIps) GetSegmentPath() string {
    return "mac-ips"
}

func (macIps *L2Rib_EviChildTables_MacIps) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mac-ip" {
        for _, c := range macIps.MacIp {
            if macIps.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := L2Rib_EviChildTables_MacIps_MacIp{}
        macIps.MacIp = append(macIps.MacIp, child)
        return &macIps.MacIp[len(macIps.MacIp)-1]
    }
    return nil
}

func (macIps *L2Rib_EviChildTables_MacIps) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range macIps.MacIp {
        children[macIps.MacIp[i].GetSegmentPath()] = &macIps.MacIp[i]
    }
    return children
}

func (macIps *L2Rib_EviChildTables_MacIps) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (macIps *L2Rib_EviChildTables_MacIps) GetBundleName() string { return "cisco_ios_xr" }

func (macIps *L2Rib_EviChildTables_MacIps) GetYangName() string { return "mac-ips" }

func (macIps *L2Rib_EviChildTables_MacIps) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (macIps *L2Rib_EviChildTables_MacIps) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (macIps *L2Rib_EviChildTables_MacIps) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (macIps *L2Rib_EviChildTables_MacIps) SetParent(parent types.Entity) { macIps.parent = parent }

func (macIps *L2Rib_EviChildTables_MacIps) GetParent() types.Entity { return macIps.parent }

func (macIps *L2Rib_EviChildTables_MacIps) GetParentYangName() string { return "evi-child-tables" }

// L2Rib_EviChildTables_MacIps_MacIp
// L2RIB EVPN MAC IP table
type L2Rib_EviChildTables_MacIps_MacIp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // EVPN ID. The type is interface{} with range: -2147483648..2147483647.
    Evi interface{}

    // Tag ID. The type is interface{} with range: -2147483648..2147483647.
    TagId interface{}

    // MAC-IP Address. The type is string with length: 1..15.
    MacAddr interface{}

    // IP Address. The type is string with length: 1..15.
    IpAddr interface{}

    // Admin distance. The type is interface{} with range:
    // -2147483648..2147483647.
    AdminDist interface{}

    // Producer ID. The type is interface{} with range: -2147483648..2147483647.
    ProdId interface{}

    // MAC Address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddress interface{}

    // IP Address. The type is string.
    IpAddress interface{}

    // Admin Distance. The type is interface{} with range: 0..255.
    AdminDistance interface{}

    // Producer ID. The type is interface{} with range: 0..255.
    ProducerId interface{}

    // Topology ID. The type is interface{} with range: 0..4294967295.
    TopologyId interface{}

    // Next Hop.
    NextHop L2Rib_EviChildTables_MacIps_MacIp_NextHop
}

func (macIp *L2Rib_EviChildTables_MacIps_MacIp) GetFilter() yfilter.YFilter { return macIp.YFilter }

func (macIp *L2Rib_EviChildTables_MacIps_MacIp) SetFilter(yf yfilter.YFilter) { macIp.YFilter = yf }

func (macIp *L2Rib_EviChildTables_MacIps_MacIp) GetGoName(yname string) string {
    if yname == "evi" { return "Evi" }
    if yname == "tag-id" { return "TagId" }
    if yname == "mac-addr" { return "MacAddr" }
    if yname == "ip-addr" { return "IpAddr" }
    if yname == "admin-dist" { return "AdminDist" }
    if yname == "prod-id" { return "ProdId" }
    if yname == "mac-address" { return "MacAddress" }
    if yname == "ip-address" { return "IpAddress" }
    if yname == "admin-distance" { return "AdminDistance" }
    if yname == "producer-id" { return "ProducerId" }
    if yname == "topology-id" { return "TopologyId" }
    if yname == "next-hop" { return "NextHop" }
    return ""
}

func (macIp *L2Rib_EviChildTables_MacIps_MacIp) GetSegmentPath() string {
    return "mac-ip"
}

func (macIp *L2Rib_EviChildTables_MacIps_MacIp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "next-hop" {
        return &macIp.NextHop
    }
    return nil
}

func (macIp *L2Rib_EviChildTables_MacIps_MacIp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["next-hop"] = &macIp.NextHop
    return children
}

func (macIp *L2Rib_EviChildTables_MacIps_MacIp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["evi"] = macIp.Evi
    leafs["tag-id"] = macIp.TagId
    leafs["mac-addr"] = macIp.MacAddr
    leafs["ip-addr"] = macIp.IpAddr
    leafs["admin-dist"] = macIp.AdminDist
    leafs["prod-id"] = macIp.ProdId
    leafs["mac-address"] = macIp.MacAddress
    leafs["ip-address"] = macIp.IpAddress
    leafs["admin-distance"] = macIp.AdminDistance
    leafs["producer-id"] = macIp.ProducerId
    leafs["topology-id"] = macIp.TopologyId
    return leafs
}

func (macIp *L2Rib_EviChildTables_MacIps_MacIp) GetBundleName() string { return "cisco_ios_xr" }

func (macIp *L2Rib_EviChildTables_MacIps_MacIp) GetYangName() string { return "mac-ip" }

func (macIp *L2Rib_EviChildTables_MacIps_MacIp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (macIp *L2Rib_EviChildTables_MacIps_MacIp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (macIp *L2Rib_EviChildTables_MacIps_MacIp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (macIp *L2Rib_EviChildTables_MacIps_MacIp) SetParent(parent types.Entity) { macIp.parent = parent }

func (macIp *L2Rib_EviChildTables_MacIps_MacIp) GetParent() types.Entity { return macIp.parent }

func (macIp *L2Rib_EviChildTables_MacIps_MacIp) GetParentYangName() string { return "mac-ips" }

// L2Rib_EviChildTables_MacIps_MacIp_NextHop
// Next Hop
type L2Rib_EviChildTables_MacIps_MacIp_NextHop struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Next-hop TOPOLOGY ID. The type is interface{} with range: 0..4294967295.
    TopologyId interface{}

    // Next-hop flags. The type is interface{} with range: 0..65535.
    Flags interface{}

    // Next hop.
    NextHop L2Rib_EviChildTables_MacIps_MacIp_NextHop_NextHop
}

func (nextHop *L2Rib_EviChildTables_MacIps_MacIp_NextHop) GetFilter() yfilter.YFilter { return nextHop.YFilter }

func (nextHop *L2Rib_EviChildTables_MacIps_MacIp_NextHop) SetFilter(yf yfilter.YFilter) { nextHop.YFilter = yf }

func (nextHop *L2Rib_EviChildTables_MacIps_MacIp_NextHop) GetGoName(yname string) string {
    if yname == "topology-id" { return "TopologyId" }
    if yname == "flags" { return "Flags" }
    if yname == "next-hop" { return "NextHop" }
    return ""
}

func (nextHop *L2Rib_EviChildTables_MacIps_MacIp_NextHop) GetSegmentPath() string {
    return "next-hop"
}

func (nextHop *L2Rib_EviChildTables_MacIps_MacIp_NextHop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "next-hop" {
        return &nextHop.NextHop
    }
    return nil
}

func (nextHop *L2Rib_EviChildTables_MacIps_MacIp_NextHop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["next-hop"] = &nextHop.NextHop
    return children
}

func (nextHop *L2Rib_EviChildTables_MacIps_MacIp_NextHop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["topology-id"] = nextHop.TopologyId
    leafs["flags"] = nextHop.Flags
    return leafs
}

func (nextHop *L2Rib_EviChildTables_MacIps_MacIp_NextHop) GetBundleName() string { return "cisco_ios_xr" }

func (nextHop *L2Rib_EviChildTables_MacIps_MacIp_NextHop) GetYangName() string { return "next-hop" }

func (nextHop *L2Rib_EviChildTables_MacIps_MacIp_NextHop) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nextHop *L2Rib_EviChildTables_MacIps_MacIp_NextHop) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nextHop *L2Rib_EviChildTables_MacIps_MacIp_NextHop) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nextHop *L2Rib_EviChildTables_MacIps_MacIp_NextHop) SetParent(parent types.Entity) { nextHop.parent = parent }

func (nextHop *L2Rib_EviChildTables_MacIps_MacIp_NextHop) GetParent() types.Entity { return nextHop.parent }

func (nextHop *L2Rib_EviChildTables_MacIps_MacIp_NextHop) GetParentYangName() string { return "mac-ip" }

// L2Rib_EviChildTables_MacIps_MacIp_NextHop_NextHop
// Next hop
type L2Rib_EviChildTables_MacIps_MacIp_NextHop_NextHop struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Type. The type is L2ribNextHop.
    Type interface{}

    // IPV4 address Next Hop. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPV6 address Next Hop. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}

    // MAC address Next Hop. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    Mac interface{}

    // Intefrace Handle Next Hop. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceHandle interface{}

    // XID Next Hop. The type is interface{} with range: 0..4294967295.
    Xid interface{}

    // Labeled Next Hop.
    Labeled L2Rib_EviChildTables_MacIps_MacIp_NextHop_NextHop_Labeled
}

func (nextHop *L2Rib_EviChildTables_MacIps_MacIp_NextHop_NextHop) GetFilter() yfilter.YFilter { return nextHop.YFilter }

func (nextHop *L2Rib_EviChildTables_MacIps_MacIp_NextHop_NextHop) SetFilter(yf yfilter.YFilter) { nextHop.YFilter = yf }

func (nextHop *L2Rib_EviChildTables_MacIps_MacIp_NextHop_NextHop) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "ipv4" { return "Ipv4" }
    if yname == "ipv6" { return "Ipv6" }
    if yname == "mac" { return "Mac" }
    if yname == "interface-handle" { return "InterfaceHandle" }
    if yname == "xid" { return "Xid" }
    if yname == "labeled" { return "Labeled" }
    return ""
}

func (nextHop *L2Rib_EviChildTables_MacIps_MacIp_NextHop_NextHop) GetSegmentPath() string {
    return "next-hop"
}

func (nextHop *L2Rib_EviChildTables_MacIps_MacIp_NextHop_NextHop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "labeled" {
        return &nextHop.Labeled
    }
    return nil
}

func (nextHop *L2Rib_EviChildTables_MacIps_MacIp_NextHop_NextHop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["labeled"] = &nextHop.Labeled
    return children
}

func (nextHop *L2Rib_EviChildTables_MacIps_MacIp_NextHop_NextHop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = nextHop.Type
    leafs["ipv4"] = nextHop.Ipv4
    leafs["ipv6"] = nextHop.Ipv6
    leafs["mac"] = nextHop.Mac
    leafs["interface-handle"] = nextHop.InterfaceHandle
    leafs["xid"] = nextHop.Xid
    return leafs
}

func (nextHop *L2Rib_EviChildTables_MacIps_MacIp_NextHop_NextHop) GetBundleName() string { return "cisco_ios_xr" }

func (nextHop *L2Rib_EviChildTables_MacIps_MacIp_NextHop_NextHop) GetYangName() string { return "next-hop" }

func (nextHop *L2Rib_EviChildTables_MacIps_MacIp_NextHop_NextHop) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nextHop *L2Rib_EviChildTables_MacIps_MacIp_NextHop_NextHop) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nextHop *L2Rib_EviChildTables_MacIps_MacIp_NextHop_NextHop) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nextHop *L2Rib_EviChildTables_MacIps_MacIp_NextHop_NextHop) SetParent(parent types.Entity) { nextHop.parent = parent }

func (nextHop *L2Rib_EviChildTables_MacIps_MacIp_NextHop_NextHop) GetParent() types.Entity { return nextHop.parent }

func (nextHop *L2Rib_EviChildTables_MacIps_MacIp_NextHop_NextHop) GetParentYangName() string { return "next-hop" }

// L2Rib_EviChildTables_MacIps_MacIp_NextHop_NextHop_Labeled
// Labeled Next Hop
type L2Rib_EviChildTables_MacIps_MacIp_NextHop_NextHop_Labeled struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // L2RIB Address Family. The type is L2ribAfi.
    AddressFamily interface{}

    // IP Address (V6 Format). The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IpAddress interface{}

    // Label. The type is interface{} with range: 0..4294967295.
    Label interface{}

    // Internal Label. The type is bool.
    Internal interface{}
}

func (labeled *L2Rib_EviChildTables_MacIps_MacIp_NextHop_NextHop_Labeled) GetFilter() yfilter.YFilter { return labeled.YFilter }

func (labeled *L2Rib_EviChildTables_MacIps_MacIp_NextHop_NextHop_Labeled) SetFilter(yf yfilter.YFilter) { labeled.YFilter = yf }

func (labeled *L2Rib_EviChildTables_MacIps_MacIp_NextHop_NextHop_Labeled) GetGoName(yname string) string {
    if yname == "address-family" { return "AddressFamily" }
    if yname == "ip-address" { return "IpAddress" }
    if yname == "label" { return "Label" }
    if yname == "internal" { return "Internal" }
    return ""
}

func (labeled *L2Rib_EviChildTables_MacIps_MacIp_NextHop_NextHop_Labeled) GetSegmentPath() string {
    return "labeled"
}

func (labeled *L2Rib_EviChildTables_MacIps_MacIp_NextHop_NextHop_Labeled) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (labeled *L2Rib_EviChildTables_MacIps_MacIp_NextHop_NextHop_Labeled) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (labeled *L2Rib_EviChildTables_MacIps_MacIp_NextHop_NextHop_Labeled) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address-family"] = labeled.AddressFamily
    leafs["ip-address"] = labeled.IpAddress
    leafs["label"] = labeled.Label
    leafs["internal"] = labeled.Internal
    return leafs
}

func (labeled *L2Rib_EviChildTables_MacIps_MacIp_NextHop_NextHop_Labeled) GetBundleName() string { return "cisco_ios_xr" }

func (labeled *L2Rib_EviChildTables_MacIps_MacIp_NextHop_NextHop_Labeled) GetYangName() string { return "labeled" }

func (labeled *L2Rib_EviChildTables_MacIps_MacIp_NextHop_NextHop_Labeled) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (labeled *L2Rib_EviChildTables_MacIps_MacIp_NextHop_NextHop_Labeled) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (labeled *L2Rib_EviChildTables_MacIps_MacIp_NextHop_NextHop_Labeled) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (labeled *L2Rib_EviChildTables_MacIps_MacIp_NextHop_NextHop_Labeled) SetParent(parent types.Entity) { labeled.parent = parent }

func (labeled *L2Rib_EviChildTables_MacIps_MacIp_NextHop_NextHop_Labeled) GetParent() types.Entity { return labeled.parent }

func (labeled *L2Rib_EviChildTables_MacIps_MacIp_NextHop_NextHop_Labeled) GetParentYangName() string { return "next-hop" }

// L2Rib_EviChildTables_Macs
// L2RIB EVPN EVI MAC table
type L2Rib_EviChildTables_Macs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // L2RIB EVPN MAC table. The type is slice of L2Rib_EviChildTables_Macs_Mac.
    Mac []L2Rib_EviChildTables_Macs_Mac
}

func (macs *L2Rib_EviChildTables_Macs) GetFilter() yfilter.YFilter { return macs.YFilter }

func (macs *L2Rib_EviChildTables_Macs) SetFilter(yf yfilter.YFilter) { macs.YFilter = yf }

func (macs *L2Rib_EviChildTables_Macs) GetGoName(yname string) string {
    if yname == "mac" { return "Mac" }
    return ""
}

func (macs *L2Rib_EviChildTables_Macs) GetSegmentPath() string {
    return "macs"
}

func (macs *L2Rib_EviChildTables_Macs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mac" {
        for _, c := range macs.Mac {
            if macs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := L2Rib_EviChildTables_Macs_Mac{}
        macs.Mac = append(macs.Mac, child)
        return &macs.Mac[len(macs.Mac)-1]
    }
    return nil
}

func (macs *L2Rib_EviChildTables_Macs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range macs.Mac {
        children[macs.Mac[i].GetSegmentPath()] = &macs.Mac[i]
    }
    return children
}

func (macs *L2Rib_EviChildTables_Macs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (macs *L2Rib_EviChildTables_Macs) GetBundleName() string { return "cisco_ios_xr" }

func (macs *L2Rib_EviChildTables_Macs) GetYangName() string { return "macs" }

func (macs *L2Rib_EviChildTables_Macs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (macs *L2Rib_EviChildTables_Macs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (macs *L2Rib_EviChildTables_Macs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (macs *L2Rib_EviChildTables_Macs) SetParent(parent types.Entity) { macs.parent = parent }

func (macs *L2Rib_EviChildTables_Macs) GetParent() types.Entity { return macs.parent }

func (macs *L2Rib_EviChildTables_Macs) GetParentYangName() string { return "evi-child-tables" }

// L2Rib_EviChildTables_Macs_Mac
// L2RIB EVPN MAC table
type L2Rib_EviChildTables_Macs_Mac struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // EVPN ID. The type is interface{} with range: -2147483648..2147483647.
    Evi interface{}

    // Tag ID. The type is interface{} with range: -2147483648..2147483647.
    TagId interface{}

    // MAC Address. The type is string with length: 1..15.
    MacAddr interface{}

    // Admin distance. The type is interface{} with range:
    // -2147483648..2147483647.
    AdminDist interface{}

    // Producer ID. The type is interface{} with range: -2147483648..2147483647.
    ProdId interface{}

    // MAC Address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddress interface{}

    // Admin Distance. The type is interface{} with range: 0..255.
    AdminDistance interface{}

    // Producer ID. The type is interface{} with range: 0..255.
    ProducerId interface{}

    // Topology ID. The type is interface{} with range: 0..4294967295.
    TopologyId interface{}

    // MAC route.
    Route L2Rib_EviChildTables_Macs_Mac_Route
}

func (mac *L2Rib_EviChildTables_Macs_Mac) GetFilter() yfilter.YFilter { return mac.YFilter }

func (mac *L2Rib_EviChildTables_Macs_Mac) SetFilter(yf yfilter.YFilter) { mac.YFilter = yf }

func (mac *L2Rib_EviChildTables_Macs_Mac) GetGoName(yname string) string {
    if yname == "evi" { return "Evi" }
    if yname == "tag-id" { return "TagId" }
    if yname == "mac-addr" { return "MacAddr" }
    if yname == "admin-dist" { return "AdminDist" }
    if yname == "prod-id" { return "ProdId" }
    if yname == "mac-address" { return "MacAddress" }
    if yname == "admin-distance" { return "AdminDistance" }
    if yname == "producer-id" { return "ProducerId" }
    if yname == "topology-id" { return "TopologyId" }
    if yname == "route" { return "Route" }
    return ""
}

func (mac *L2Rib_EviChildTables_Macs_Mac) GetSegmentPath() string {
    return "mac"
}

func (mac *L2Rib_EviChildTables_Macs_Mac) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "route" {
        return &mac.Route
    }
    return nil
}

func (mac *L2Rib_EviChildTables_Macs_Mac) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["route"] = &mac.Route
    return children
}

func (mac *L2Rib_EviChildTables_Macs_Mac) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["evi"] = mac.Evi
    leafs["tag-id"] = mac.TagId
    leafs["mac-addr"] = mac.MacAddr
    leafs["admin-dist"] = mac.AdminDist
    leafs["prod-id"] = mac.ProdId
    leafs["mac-address"] = mac.MacAddress
    leafs["admin-distance"] = mac.AdminDistance
    leafs["producer-id"] = mac.ProducerId
    leafs["topology-id"] = mac.TopologyId
    return leafs
}

func (mac *L2Rib_EviChildTables_Macs_Mac) GetBundleName() string { return "cisco_ios_xr" }

func (mac *L2Rib_EviChildTables_Macs_Mac) GetYangName() string { return "mac" }

func (mac *L2Rib_EviChildTables_Macs_Mac) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mac *L2Rib_EviChildTables_Macs_Mac) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mac *L2Rib_EviChildTables_Macs_Mac) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mac *L2Rib_EviChildTables_Macs_Mac) SetParent(parent types.Entity) { mac.parent = parent }

func (mac *L2Rib_EviChildTables_Macs_Mac) GetParent() types.Entity { return mac.parent }

func (mac *L2Rib_EviChildTables_Macs_Mac) GetParentYangName() string { return "macs" }

// L2Rib_EviChildTables_Macs_Mac_Route
// MAC route
type L2Rib_EviChildTables_Macs_Mac_Route struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Type. The type is L2ribMacRoute.
    Type interface{}

    // Regular MAC route.
    Regular L2Rib_EviChildTables_Macs_Mac_Route_Regular

    // EVPN ESI MAC route.
    EvpnEsi L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi

    // BMAC route.
    Bmac L2Rib_EviChildTables_Macs_Mac_Route_Bmac
}

func (route *L2Rib_EviChildTables_Macs_Mac_Route) GetFilter() yfilter.YFilter { return route.YFilter }

func (route *L2Rib_EviChildTables_Macs_Mac_Route) SetFilter(yf yfilter.YFilter) { route.YFilter = yf }

func (route *L2Rib_EviChildTables_Macs_Mac_Route) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "regular" { return "Regular" }
    if yname == "evpn-esi" { return "EvpnEsi" }
    if yname == "bmac" { return "Bmac" }
    return ""
}

func (route *L2Rib_EviChildTables_Macs_Mac_Route) GetSegmentPath() string {
    return "route"
}

func (route *L2Rib_EviChildTables_Macs_Mac_Route) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "regular" {
        return &route.Regular
    }
    if childYangName == "evpn-esi" {
        return &route.EvpnEsi
    }
    if childYangName == "bmac" {
        return &route.Bmac
    }
    return nil
}

func (route *L2Rib_EviChildTables_Macs_Mac_Route) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["regular"] = &route.Regular
    children["evpn-esi"] = &route.EvpnEsi
    children["bmac"] = &route.Bmac
    return children
}

func (route *L2Rib_EviChildTables_Macs_Mac_Route) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = route.Type
    return leafs
}

func (route *L2Rib_EviChildTables_Macs_Mac_Route) GetBundleName() string { return "cisco_ios_xr" }

func (route *L2Rib_EviChildTables_Macs_Mac_Route) GetYangName() string { return "route" }

func (route *L2Rib_EviChildTables_Macs_Mac_Route) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (route *L2Rib_EviChildTables_Macs_Mac_Route) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (route *L2Rib_EviChildTables_Macs_Mac_Route) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (route *L2Rib_EviChildTables_Macs_Mac_Route) SetParent(parent types.Entity) { route.parent = parent }

func (route *L2Rib_EviChildTables_Macs_Mac_Route) GetParent() types.Entity { return route.parent }

func (route *L2Rib_EviChildTables_Macs_Mac_Route) GetParentYangName() string { return "mac" }

// L2Rib_EviChildTables_Macs_Mac_Route_Regular
// Regular MAC route
type L2Rib_EviChildTables_Macs_Mac_Route_Regular struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Next Hop.
    NextHop L2Rib_EviChildTables_Macs_Mac_Route_Regular_NextHop
}

func (regular *L2Rib_EviChildTables_Macs_Mac_Route_Regular) GetFilter() yfilter.YFilter { return regular.YFilter }

func (regular *L2Rib_EviChildTables_Macs_Mac_Route_Regular) SetFilter(yf yfilter.YFilter) { regular.YFilter = yf }

func (regular *L2Rib_EviChildTables_Macs_Mac_Route_Regular) GetGoName(yname string) string {
    if yname == "next-hop" { return "NextHop" }
    return ""
}

func (regular *L2Rib_EviChildTables_Macs_Mac_Route_Regular) GetSegmentPath() string {
    return "regular"
}

func (regular *L2Rib_EviChildTables_Macs_Mac_Route_Regular) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "next-hop" {
        return &regular.NextHop
    }
    return nil
}

func (regular *L2Rib_EviChildTables_Macs_Mac_Route_Regular) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["next-hop"] = &regular.NextHop
    return children
}

func (regular *L2Rib_EviChildTables_Macs_Mac_Route_Regular) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (regular *L2Rib_EviChildTables_Macs_Mac_Route_Regular) GetBundleName() string { return "cisco_ios_xr" }

func (regular *L2Rib_EviChildTables_Macs_Mac_Route_Regular) GetYangName() string { return "regular" }

func (regular *L2Rib_EviChildTables_Macs_Mac_Route_Regular) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (regular *L2Rib_EviChildTables_Macs_Mac_Route_Regular) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (regular *L2Rib_EviChildTables_Macs_Mac_Route_Regular) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (regular *L2Rib_EviChildTables_Macs_Mac_Route_Regular) SetParent(parent types.Entity) { regular.parent = parent }

func (regular *L2Rib_EviChildTables_Macs_Mac_Route_Regular) GetParent() types.Entity { return regular.parent }

func (regular *L2Rib_EviChildTables_Macs_Mac_Route_Regular) GetParentYangName() string { return "route" }

// L2Rib_EviChildTables_Macs_Mac_Route_Regular_NextHop
// Next Hop
type L2Rib_EviChildTables_Macs_Mac_Route_Regular_NextHop struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Next-hop TOPOLOGY ID. The type is interface{} with range: 0..4294967295.
    TopologyId interface{}

    // Next-hop flags. The type is interface{} with range: 0..65535.
    Flags interface{}

    // Next hop.
    NextHop L2Rib_EviChildTables_Macs_Mac_Route_Regular_NextHop_NextHop
}

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_Regular_NextHop) GetFilter() yfilter.YFilter { return nextHop.YFilter }

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_Regular_NextHop) SetFilter(yf yfilter.YFilter) { nextHop.YFilter = yf }

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_Regular_NextHop) GetGoName(yname string) string {
    if yname == "topology-id" { return "TopologyId" }
    if yname == "flags" { return "Flags" }
    if yname == "next-hop" { return "NextHop" }
    return ""
}

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_Regular_NextHop) GetSegmentPath() string {
    return "next-hop"
}

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_Regular_NextHop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "next-hop" {
        return &nextHop.NextHop
    }
    return nil
}

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_Regular_NextHop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["next-hop"] = &nextHop.NextHop
    return children
}

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_Regular_NextHop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["topology-id"] = nextHop.TopologyId
    leafs["flags"] = nextHop.Flags
    return leafs
}

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_Regular_NextHop) GetBundleName() string { return "cisco_ios_xr" }

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_Regular_NextHop) GetYangName() string { return "next-hop" }

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_Regular_NextHop) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_Regular_NextHop) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_Regular_NextHop) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_Regular_NextHop) SetParent(parent types.Entity) { nextHop.parent = parent }

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_Regular_NextHop) GetParent() types.Entity { return nextHop.parent }

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_Regular_NextHop) GetParentYangName() string { return "regular" }

// L2Rib_EviChildTables_Macs_Mac_Route_Regular_NextHop_NextHop
// Next hop
type L2Rib_EviChildTables_Macs_Mac_Route_Regular_NextHop_NextHop struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Type. The type is L2ribNextHop.
    Type interface{}

    // IPV4 address Next Hop. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPV6 address Next Hop. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}

    // MAC address Next Hop. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    Mac interface{}

    // Intefrace Handle Next Hop. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceHandle interface{}

    // XID Next Hop. The type is interface{} with range: 0..4294967295.
    Xid interface{}

    // Labeled Next Hop.
    Labeled L2Rib_EviChildTables_Macs_Mac_Route_Regular_NextHop_NextHop_Labeled
}

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_Regular_NextHop_NextHop) GetFilter() yfilter.YFilter { return nextHop.YFilter }

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_Regular_NextHop_NextHop) SetFilter(yf yfilter.YFilter) { nextHop.YFilter = yf }

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_Regular_NextHop_NextHop) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "ipv4" { return "Ipv4" }
    if yname == "ipv6" { return "Ipv6" }
    if yname == "mac" { return "Mac" }
    if yname == "interface-handle" { return "InterfaceHandle" }
    if yname == "xid" { return "Xid" }
    if yname == "labeled" { return "Labeled" }
    return ""
}

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_Regular_NextHop_NextHop) GetSegmentPath() string {
    return "next-hop"
}

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_Regular_NextHop_NextHop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "labeled" {
        return &nextHop.Labeled
    }
    return nil
}

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_Regular_NextHop_NextHop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["labeled"] = &nextHop.Labeled
    return children
}

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_Regular_NextHop_NextHop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = nextHop.Type
    leafs["ipv4"] = nextHop.Ipv4
    leafs["ipv6"] = nextHop.Ipv6
    leafs["mac"] = nextHop.Mac
    leafs["interface-handle"] = nextHop.InterfaceHandle
    leafs["xid"] = nextHop.Xid
    return leafs
}

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_Regular_NextHop_NextHop) GetBundleName() string { return "cisco_ios_xr" }

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_Regular_NextHop_NextHop) GetYangName() string { return "next-hop" }

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_Regular_NextHop_NextHop) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_Regular_NextHop_NextHop) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_Regular_NextHop_NextHop) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_Regular_NextHop_NextHop) SetParent(parent types.Entity) { nextHop.parent = parent }

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_Regular_NextHop_NextHop) GetParent() types.Entity { return nextHop.parent }

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_Regular_NextHop_NextHop) GetParentYangName() string { return "next-hop" }

// L2Rib_EviChildTables_Macs_Mac_Route_Regular_NextHop_NextHop_Labeled
// Labeled Next Hop
type L2Rib_EviChildTables_Macs_Mac_Route_Regular_NextHop_NextHop_Labeled struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // L2RIB Address Family. The type is L2ribAfi.
    AddressFamily interface{}

    // IP Address (V6 Format). The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IpAddress interface{}

    // Label. The type is interface{} with range: 0..4294967295.
    Label interface{}

    // Internal Label. The type is bool.
    Internal interface{}
}

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_Regular_NextHop_NextHop_Labeled) GetFilter() yfilter.YFilter { return labeled.YFilter }

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_Regular_NextHop_NextHop_Labeled) SetFilter(yf yfilter.YFilter) { labeled.YFilter = yf }

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_Regular_NextHop_NextHop_Labeled) GetGoName(yname string) string {
    if yname == "address-family" { return "AddressFamily" }
    if yname == "ip-address" { return "IpAddress" }
    if yname == "label" { return "Label" }
    if yname == "internal" { return "Internal" }
    return ""
}

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_Regular_NextHop_NextHop_Labeled) GetSegmentPath() string {
    return "labeled"
}

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_Regular_NextHop_NextHop_Labeled) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_Regular_NextHop_NextHop_Labeled) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_Regular_NextHop_NextHop_Labeled) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address-family"] = labeled.AddressFamily
    leafs["ip-address"] = labeled.IpAddress
    leafs["label"] = labeled.Label
    leafs["internal"] = labeled.Internal
    return leafs
}

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_Regular_NextHop_NextHop_Labeled) GetBundleName() string { return "cisco_ios_xr" }

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_Regular_NextHop_NextHop_Labeled) GetYangName() string { return "labeled" }

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_Regular_NextHop_NextHop_Labeled) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_Regular_NextHop_NextHop_Labeled) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_Regular_NextHop_NextHop_Labeled) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_Regular_NextHop_NextHop_Labeled) SetParent(parent types.Entity) { labeled.parent = parent }

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_Regular_NextHop_NextHop_Labeled) GetParent() types.Entity { return labeled.parent }

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_Regular_NextHop_NextHop_Labeled) GetParentYangName() string { return "next-hop" }

// L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi
// EVPN ESI MAC route
type L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MAC route sequence number. The type is interface{} with range:
    // 0..4294967295.
    SequenceNumber interface{}

    // Forwarding State. True means forward, False means drop. The type is bool.
    ForwardState interface{}

    // Ethernet Segment Identifier.
    EthernetSegmentId L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_EthernetSegmentId

    // Path list information. Set for detailed MAC route information.
    PathList L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList
}

func (evpnEsi *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi) GetFilter() yfilter.YFilter { return evpnEsi.YFilter }

func (evpnEsi *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi) SetFilter(yf yfilter.YFilter) { evpnEsi.YFilter = yf }

func (evpnEsi *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi) GetGoName(yname string) string {
    if yname == "sequence-number" { return "SequenceNumber" }
    if yname == "forward-state" { return "ForwardState" }
    if yname == "ethernet-segment-id" { return "EthernetSegmentId" }
    if yname == "path-list" { return "PathList" }
    return ""
}

func (evpnEsi *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi) GetSegmentPath() string {
    return "evpn-esi"
}

func (evpnEsi *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ethernet-segment-id" {
        return &evpnEsi.EthernetSegmentId
    }
    if childYangName == "path-list" {
        return &evpnEsi.PathList
    }
    return nil
}

func (evpnEsi *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ethernet-segment-id"] = &evpnEsi.EthernetSegmentId
    children["path-list"] = &evpnEsi.PathList
    return children
}

func (evpnEsi *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sequence-number"] = evpnEsi.SequenceNumber
    leafs["forward-state"] = evpnEsi.ForwardState
    return leafs
}

func (evpnEsi *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi) GetBundleName() string { return "cisco_ios_xr" }

func (evpnEsi *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi) GetYangName() string { return "evpn-esi" }

func (evpnEsi *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (evpnEsi *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (evpnEsi *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (evpnEsi *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi) SetParent(parent types.Entity) { evpnEsi.parent = parent }

func (evpnEsi *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi) GetParent() types.Entity { return evpnEsi.parent }

func (evpnEsi *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi) GetParentYangName() string { return "route" }

// L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_EthernetSegmentId
// Ethernet Segment Identifier
type L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_EthernetSegmentId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // LACP System Priority. The type is interface{} with range: 0..65535.
    SystemPriority interface{}

    // LACP System Id. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    SystemId interface{}

    // LACP Port Key. The type is interface{} with range: 0..65535.
    PortKey interface{}
}

func (ethernetSegmentId *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_EthernetSegmentId) GetFilter() yfilter.YFilter { return ethernetSegmentId.YFilter }

func (ethernetSegmentId *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_EthernetSegmentId) SetFilter(yf yfilter.YFilter) { ethernetSegmentId.YFilter = yf }

func (ethernetSegmentId *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_EthernetSegmentId) GetGoName(yname string) string {
    if yname == "system-priority" { return "SystemPriority" }
    if yname == "system-id" { return "SystemId" }
    if yname == "port-key" { return "PortKey" }
    return ""
}

func (ethernetSegmentId *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_EthernetSegmentId) GetSegmentPath() string {
    return "ethernet-segment-id"
}

func (ethernetSegmentId *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_EthernetSegmentId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ethernetSegmentId *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_EthernetSegmentId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ethernetSegmentId *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_EthernetSegmentId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["system-priority"] = ethernetSegmentId.SystemPriority
    leafs["system-id"] = ethernetSegmentId.SystemId
    leafs["port-key"] = ethernetSegmentId.PortKey
    return leafs
}

func (ethernetSegmentId *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_EthernetSegmentId) GetBundleName() string { return "cisco_ios_xr" }

func (ethernetSegmentId *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_EthernetSegmentId) GetYangName() string { return "ethernet-segment-id" }

func (ethernetSegmentId *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_EthernetSegmentId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ethernetSegmentId *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_EthernetSegmentId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ethernetSegmentId *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_EthernetSegmentId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ethernetSegmentId *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_EthernetSegmentId) SetParent(parent types.Entity) { ethernetSegmentId.parent = parent }

func (ethernetSegmentId *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_EthernetSegmentId) GetParent() types.Entity { return ethernetSegmentId.parent }

func (ethernetSegmentId *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_EthernetSegmentId) GetParentYangName() string { return "evpn-esi" }

// L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList
// Path list information. Set for detailed MAC
// route information
type L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ID of EAD route producer. The type is interface{} with range: 0..255.
    ProducerId interface{}

    // Number of MAC routes bound to this Path list. The type is interface{} with
    // range: 0..4294967295.
    MacCount interface{}

    // Path list local Label. The type is interface{} with range: 0..4294967295.
    LocalLabel interface{}

    // Type-specific Path List info.
    PathListInfo L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo

    // Array of Next Hops for MAC Path List. The type is slice of
    // L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_NextHopArray.
    NextHopArray []L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_NextHopArray
}

func (pathList *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList) GetFilter() yfilter.YFilter { return pathList.YFilter }

func (pathList *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList) SetFilter(yf yfilter.YFilter) { pathList.YFilter = yf }

func (pathList *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList) GetGoName(yname string) string {
    if yname == "producer-id" { return "ProducerId" }
    if yname == "mac-count" { return "MacCount" }
    if yname == "local-label" { return "LocalLabel" }
    if yname == "path-list-info" { return "PathListInfo" }
    if yname == "next-hop-array" { return "NextHopArray" }
    return ""
}

func (pathList *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList) GetSegmentPath() string {
    return "path-list"
}

func (pathList *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "path-list-info" {
        return &pathList.PathListInfo
    }
    if childYangName == "next-hop-array" {
        for _, c := range pathList.NextHopArray {
            if pathList.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_NextHopArray{}
        pathList.NextHopArray = append(pathList.NextHopArray, child)
        return &pathList.NextHopArray[len(pathList.NextHopArray)-1]
    }
    return nil
}

func (pathList *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["path-list-info"] = &pathList.PathListInfo
    for i := range pathList.NextHopArray {
        children[pathList.NextHopArray[i].GetSegmentPath()] = &pathList.NextHopArray[i]
    }
    return children
}

func (pathList *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["producer-id"] = pathList.ProducerId
    leafs["mac-count"] = pathList.MacCount
    leafs["local-label"] = pathList.LocalLabel
    return leafs
}

func (pathList *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList) GetBundleName() string { return "cisco_ios_xr" }

func (pathList *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList) GetYangName() string { return "path-list" }

func (pathList *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pathList *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pathList *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pathList *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList) SetParent(parent types.Entity) { pathList.parent = parent }

func (pathList *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList) GetParent() types.Entity { return pathList.parent }

func (pathList *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList) GetParentYangName() string { return "evpn-esi" }

// L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo
// Type-specific Path List info
type L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Type. The type is L2ribMacRoute.
    Type interface{}

    // ESI Path List.
    PathListEsi L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi

    // MAC Path List.
    PathListMac L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListMac
}

func (pathListInfo *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo) GetFilter() yfilter.YFilter { return pathListInfo.YFilter }

func (pathListInfo *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo) SetFilter(yf yfilter.YFilter) { pathListInfo.YFilter = yf }

func (pathListInfo *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "path-list-esi" { return "PathListEsi" }
    if yname == "path-list-mac" { return "PathListMac" }
    return ""
}

func (pathListInfo *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo) GetSegmentPath() string {
    return "path-list-info"
}

func (pathListInfo *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "path-list-esi" {
        return &pathListInfo.PathListEsi
    }
    if childYangName == "path-list-mac" {
        return &pathListInfo.PathListMac
    }
    return nil
}

func (pathListInfo *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["path-list-esi"] = &pathListInfo.PathListEsi
    children["path-list-mac"] = &pathListInfo.PathListMac
    return children
}

func (pathListInfo *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = pathListInfo.Type
    return leafs
}

func (pathListInfo *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo) GetBundleName() string { return "cisco_ios_xr" }

func (pathListInfo *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo) GetYangName() string { return "path-list-info" }

func (pathListInfo *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pathListInfo *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pathListInfo *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pathListInfo *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo) SetParent(parent types.Entity) { pathListInfo.parent = parent }

func (pathListInfo *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo) GetParent() types.Entity { return pathListInfo.parent }

func (pathListInfo *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo) GetParentYangName() string { return "path-list" }

// L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi
// ESI Path List
type L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Path list Resolved. The type is bool.
    Resolved interface{}

    // Ethernet Segment Identifier.
    EthernetSegmentId L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_EthernetSegmentId

    // Array of Next Hops from MAC Update. The type is slice of
    // L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray.
    MacUpdateNextHopArray []L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray
}

func (pathListEsi *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi) GetFilter() yfilter.YFilter { return pathListEsi.YFilter }

func (pathListEsi *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi) SetFilter(yf yfilter.YFilter) { pathListEsi.YFilter = yf }

func (pathListEsi *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi) GetGoName(yname string) string {
    if yname == "resolved" { return "Resolved" }
    if yname == "ethernet-segment-id" { return "EthernetSegmentId" }
    if yname == "mac-update-next-hop-array" { return "MacUpdateNextHopArray" }
    return ""
}

func (pathListEsi *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi) GetSegmentPath() string {
    return "path-list-esi"
}

func (pathListEsi *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ethernet-segment-id" {
        return &pathListEsi.EthernetSegmentId
    }
    if childYangName == "mac-update-next-hop-array" {
        for _, c := range pathListEsi.MacUpdateNextHopArray {
            if pathListEsi.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray{}
        pathListEsi.MacUpdateNextHopArray = append(pathListEsi.MacUpdateNextHopArray, child)
        return &pathListEsi.MacUpdateNextHopArray[len(pathListEsi.MacUpdateNextHopArray)-1]
    }
    return nil
}

func (pathListEsi *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ethernet-segment-id"] = &pathListEsi.EthernetSegmentId
    for i := range pathListEsi.MacUpdateNextHopArray {
        children[pathListEsi.MacUpdateNextHopArray[i].GetSegmentPath()] = &pathListEsi.MacUpdateNextHopArray[i]
    }
    return children
}

func (pathListEsi *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["resolved"] = pathListEsi.Resolved
    return leafs
}

func (pathListEsi *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi) GetBundleName() string { return "cisco_ios_xr" }

func (pathListEsi *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi) GetYangName() string { return "path-list-esi" }

func (pathListEsi *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pathListEsi *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pathListEsi *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pathListEsi *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi) SetParent(parent types.Entity) { pathListEsi.parent = parent }

func (pathListEsi *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi) GetParent() types.Entity { return pathListEsi.parent }

func (pathListEsi *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi) GetParentYangName() string { return "path-list-info" }

// L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_EthernetSegmentId
// Ethernet Segment Identifier
type L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_EthernetSegmentId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // LACP System Priority. The type is interface{} with range: 0..65535.
    SystemPriority interface{}

    // LACP System Id. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    SystemId interface{}

    // LACP Port Key. The type is interface{} with range: 0..65535.
    PortKey interface{}
}

func (ethernetSegmentId *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_EthernetSegmentId) GetFilter() yfilter.YFilter { return ethernetSegmentId.YFilter }

func (ethernetSegmentId *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_EthernetSegmentId) SetFilter(yf yfilter.YFilter) { ethernetSegmentId.YFilter = yf }

func (ethernetSegmentId *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_EthernetSegmentId) GetGoName(yname string) string {
    if yname == "system-priority" { return "SystemPriority" }
    if yname == "system-id" { return "SystemId" }
    if yname == "port-key" { return "PortKey" }
    return ""
}

func (ethernetSegmentId *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_EthernetSegmentId) GetSegmentPath() string {
    return "ethernet-segment-id"
}

func (ethernetSegmentId *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_EthernetSegmentId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ethernetSegmentId *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_EthernetSegmentId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ethernetSegmentId *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_EthernetSegmentId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["system-priority"] = ethernetSegmentId.SystemPriority
    leafs["system-id"] = ethernetSegmentId.SystemId
    leafs["port-key"] = ethernetSegmentId.PortKey
    return leafs
}

func (ethernetSegmentId *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_EthernetSegmentId) GetBundleName() string { return "cisco_ios_xr" }

func (ethernetSegmentId *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_EthernetSegmentId) GetYangName() string { return "ethernet-segment-id" }

func (ethernetSegmentId *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_EthernetSegmentId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ethernetSegmentId *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_EthernetSegmentId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ethernetSegmentId *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_EthernetSegmentId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ethernetSegmentId *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_EthernetSegmentId) SetParent(parent types.Entity) { ethernetSegmentId.parent = parent }

func (ethernetSegmentId *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_EthernetSegmentId) GetParent() types.Entity { return ethernetSegmentId.parent }

func (ethernetSegmentId *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_EthernetSegmentId) GetParentYangName() string { return "path-list-esi" }

// L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray
// Array of Next Hops from MAC Update
type L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Next-hop TOPOLOGY ID. The type is interface{} with range: 0..4294967295.
    TopologyId interface{}

    // Next-hop flags. The type is interface{} with range: 0..65535.
    Flags interface{}

    // Next hop.
    NextHop L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop
}

func (macUpdateNextHopArray *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray) GetFilter() yfilter.YFilter { return macUpdateNextHopArray.YFilter }

func (macUpdateNextHopArray *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray) SetFilter(yf yfilter.YFilter) { macUpdateNextHopArray.YFilter = yf }

func (macUpdateNextHopArray *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray) GetGoName(yname string) string {
    if yname == "topology-id" { return "TopologyId" }
    if yname == "flags" { return "Flags" }
    if yname == "next-hop" { return "NextHop" }
    return ""
}

func (macUpdateNextHopArray *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray) GetSegmentPath() string {
    return "mac-update-next-hop-array"
}

func (macUpdateNextHopArray *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "next-hop" {
        return &macUpdateNextHopArray.NextHop
    }
    return nil
}

func (macUpdateNextHopArray *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["next-hop"] = &macUpdateNextHopArray.NextHop
    return children
}

func (macUpdateNextHopArray *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["topology-id"] = macUpdateNextHopArray.TopologyId
    leafs["flags"] = macUpdateNextHopArray.Flags
    return leafs
}

func (macUpdateNextHopArray *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray) GetBundleName() string { return "cisco_ios_xr" }

func (macUpdateNextHopArray *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray) GetYangName() string { return "mac-update-next-hop-array" }

func (macUpdateNextHopArray *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (macUpdateNextHopArray *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (macUpdateNextHopArray *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (macUpdateNextHopArray *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray) SetParent(parent types.Entity) { macUpdateNextHopArray.parent = parent }

func (macUpdateNextHopArray *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray) GetParent() types.Entity { return macUpdateNextHopArray.parent }

func (macUpdateNextHopArray *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray) GetParentYangName() string { return "path-list-esi" }

// L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop
// Next hop
type L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Type. The type is L2ribNextHop.
    Type interface{}

    // IPV4 address Next Hop. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPV6 address Next Hop. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}

    // MAC address Next Hop. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    Mac interface{}

    // Intefrace Handle Next Hop. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceHandle interface{}

    // XID Next Hop. The type is interface{} with range: 0..4294967295.
    Xid interface{}

    // Labeled Next Hop.
    Labeled L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled
}

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop) GetFilter() yfilter.YFilter { return nextHop.YFilter }

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop) SetFilter(yf yfilter.YFilter) { nextHop.YFilter = yf }

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "ipv4" { return "Ipv4" }
    if yname == "ipv6" { return "Ipv6" }
    if yname == "mac" { return "Mac" }
    if yname == "interface-handle" { return "InterfaceHandle" }
    if yname == "xid" { return "Xid" }
    if yname == "labeled" { return "Labeled" }
    return ""
}

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop) GetSegmentPath() string {
    return "next-hop"
}

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "labeled" {
        return &nextHop.Labeled
    }
    return nil
}

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["labeled"] = &nextHop.Labeled
    return children
}

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = nextHop.Type
    leafs["ipv4"] = nextHop.Ipv4
    leafs["ipv6"] = nextHop.Ipv6
    leafs["mac"] = nextHop.Mac
    leafs["interface-handle"] = nextHop.InterfaceHandle
    leafs["xid"] = nextHop.Xid
    return leafs
}

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop) GetBundleName() string { return "cisco_ios_xr" }

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop) GetYangName() string { return "next-hop" }

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop) SetParent(parent types.Entity) { nextHop.parent = parent }

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop) GetParent() types.Entity { return nextHop.parent }

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop) GetParentYangName() string { return "mac-update-next-hop-array" }

// L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled
// Labeled Next Hop
type L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // L2RIB Address Family. The type is L2ribAfi.
    AddressFamily interface{}

    // IP Address (V6 Format). The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IpAddress interface{}

    // Label. The type is interface{} with range: 0..4294967295.
    Label interface{}

    // Internal Label. The type is bool.
    Internal interface{}
}

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled) GetFilter() yfilter.YFilter { return labeled.YFilter }

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled) SetFilter(yf yfilter.YFilter) { labeled.YFilter = yf }

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled) GetGoName(yname string) string {
    if yname == "address-family" { return "AddressFamily" }
    if yname == "ip-address" { return "IpAddress" }
    if yname == "label" { return "Label" }
    if yname == "internal" { return "Internal" }
    return ""
}

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled) GetSegmentPath() string {
    return "labeled"
}

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address-family"] = labeled.AddressFamily
    leafs["ip-address"] = labeled.IpAddress
    leafs["label"] = labeled.Label
    leafs["internal"] = labeled.Internal
    return leafs
}

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled) GetBundleName() string { return "cisco_ios_xr" }

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled) GetYangName() string { return "labeled" }

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled) SetParent(parent types.Entity) { labeled.parent = parent }

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled) GetParent() types.Entity { return labeled.parent }

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled) GetParentYangName() string { return "next-hop" }

// L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListMac
// MAC Path List
type L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListMac struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MAC Address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddress interface{}
}

func (pathListMac *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListMac) GetFilter() yfilter.YFilter { return pathListMac.YFilter }

func (pathListMac *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListMac) SetFilter(yf yfilter.YFilter) { pathListMac.YFilter = yf }

func (pathListMac *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListMac) GetGoName(yname string) string {
    if yname == "mac-address" { return "MacAddress" }
    return ""
}

func (pathListMac *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListMac) GetSegmentPath() string {
    return "path-list-mac"
}

func (pathListMac *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListMac) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pathListMac *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListMac) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pathListMac *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListMac) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mac-address"] = pathListMac.MacAddress
    return leafs
}

func (pathListMac *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListMac) GetBundleName() string { return "cisco_ios_xr" }

func (pathListMac *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListMac) GetYangName() string { return "path-list-mac" }

func (pathListMac *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListMac) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pathListMac *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListMac) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pathListMac *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListMac) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pathListMac *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListMac) SetParent(parent types.Entity) { pathListMac.parent = parent }

func (pathListMac *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListMac) GetParent() types.Entity { return pathListMac.parent }

func (pathListMac *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListMac) GetParentYangName() string { return "path-list-info" }

// L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_NextHopArray
// Array of Next Hops for MAC Path List
type L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_NextHopArray struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Next-hop TOPOLOGY ID. The type is interface{} with range: 0..4294967295.
    TopologyId interface{}

    // Next-hop flags. The type is interface{} with range: 0..65535.
    Flags interface{}

    // Next hop.
    NextHop L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_NextHopArray_NextHop
}

func (nextHopArray *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_NextHopArray) GetFilter() yfilter.YFilter { return nextHopArray.YFilter }

func (nextHopArray *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_NextHopArray) SetFilter(yf yfilter.YFilter) { nextHopArray.YFilter = yf }

func (nextHopArray *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_NextHopArray) GetGoName(yname string) string {
    if yname == "topology-id" { return "TopologyId" }
    if yname == "flags" { return "Flags" }
    if yname == "next-hop" { return "NextHop" }
    return ""
}

func (nextHopArray *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_NextHopArray) GetSegmentPath() string {
    return "next-hop-array"
}

func (nextHopArray *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_NextHopArray) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "next-hop" {
        return &nextHopArray.NextHop
    }
    return nil
}

func (nextHopArray *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_NextHopArray) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["next-hop"] = &nextHopArray.NextHop
    return children
}

func (nextHopArray *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_NextHopArray) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["topology-id"] = nextHopArray.TopologyId
    leafs["flags"] = nextHopArray.Flags
    return leafs
}

func (nextHopArray *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_NextHopArray) GetBundleName() string { return "cisco_ios_xr" }

func (nextHopArray *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_NextHopArray) GetYangName() string { return "next-hop-array" }

func (nextHopArray *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_NextHopArray) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nextHopArray *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_NextHopArray) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nextHopArray *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_NextHopArray) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nextHopArray *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_NextHopArray) SetParent(parent types.Entity) { nextHopArray.parent = parent }

func (nextHopArray *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_NextHopArray) GetParent() types.Entity { return nextHopArray.parent }

func (nextHopArray *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_NextHopArray) GetParentYangName() string { return "path-list" }

// L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_NextHopArray_NextHop
// Next hop
type L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_NextHopArray_NextHop struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Type. The type is L2ribNextHop.
    Type interface{}

    // IPV4 address Next Hop. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPV6 address Next Hop. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}

    // MAC address Next Hop. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    Mac interface{}

    // Intefrace Handle Next Hop. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceHandle interface{}

    // XID Next Hop. The type is interface{} with range: 0..4294967295.
    Xid interface{}

    // Labeled Next Hop.
    Labeled L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_NextHopArray_NextHop_Labeled
}

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_NextHopArray_NextHop) GetFilter() yfilter.YFilter { return nextHop.YFilter }

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_NextHopArray_NextHop) SetFilter(yf yfilter.YFilter) { nextHop.YFilter = yf }

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_NextHopArray_NextHop) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "ipv4" { return "Ipv4" }
    if yname == "ipv6" { return "Ipv6" }
    if yname == "mac" { return "Mac" }
    if yname == "interface-handle" { return "InterfaceHandle" }
    if yname == "xid" { return "Xid" }
    if yname == "labeled" { return "Labeled" }
    return ""
}

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_NextHopArray_NextHop) GetSegmentPath() string {
    return "next-hop"
}

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_NextHopArray_NextHop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "labeled" {
        return &nextHop.Labeled
    }
    return nil
}

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_NextHopArray_NextHop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["labeled"] = &nextHop.Labeled
    return children
}

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_NextHopArray_NextHop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = nextHop.Type
    leafs["ipv4"] = nextHop.Ipv4
    leafs["ipv6"] = nextHop.Ipv6
    leafs["mac"] = nextHop.Mac
    leafs["interface-handle"] = nextHop.InterfaceHandle
    leafs["xid"] = nextHop.Xid
    return leafs
}

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_NextHopArray_NextHop) GetBundleName() string { return "cisco_ios_xr" }

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_NextHopArray_NextHop) GetYangName() string { return "next-hop" }

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_NextHopArray_NextHop) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_NextHopArray_NextHop) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_NextHopArray_NextHop) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_NextHopArray_NextHop) SetParent(parent types.Entity) { nextHop.parent = parent }

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_NextHopArray_NextHop) GetParent() types.Entity { return nextHop.parent }

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_NextHopArray_NextHop) GetParentYangName() string { return "next-hop-array" }

// L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_NextHopArray_NextHop_Labeled
// Labeled Next Hop
type L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_NextHopArray_NextHop_Labeled struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // L2RIB Address Family. The type is L2ribAfi.
    AddressFamily interface{}

    // IP Address (V6 Format). The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IpAddress interface{}

    // Label. The type is interface{} with range: 0..4294967295.
    Label interface{}

    // Internal Label. The type is bool.
    Internal interface{}
}

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_NextHopArray_NextHop_Labeled) GetFilter() yfilter.YFilter { return labeled.YFilter }

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_NextHopArray_NextHop_Labeled) SetFilter(yf yfilter.YFilter) { labeled.YFilter = yf }

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_NextHopArray_NextHop_Labeled) GetGoName(yname string) string {
    if yname == "address-family" { return "AddressFamily" }
    if yname == "ip-address" { return "IpAddress" }
    if yname == "label" { return "Label" }
    if yname == "internal" { return "Internal" }
    return ""
}

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_NextHopArray_NextHop_Labeled) GetSegmentPath() string {
    return "labeled"
}

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_NextHopArray_NextHop_Labeled) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_NextHopArray_NextHop_Labeled) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_NextHopArray_NextHop_Labeled) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address-family"] = labeled.AddressFamily
    leafs["ip-address"] = labeled.IpAddress
    leafs["label"] = labeled.Label
    leafs["internal"] = labeled.Internal
    return leafs
}

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_NextHopArray_NextHop_Labeled) GetBundleName() string { return "cisco_ios_xr" }

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_NextHopArray_NextHop_Labeled) GetYangName() string { return "labeled" }

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_NextHopArray_NextHop_Labeled) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_NextHopArray_NextHop_Labeled) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_NextHopArray_NextHop_Labeled) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_NextHopArray_NextHop_Labeled) SetParent(parent types.Entity) { labeled.parent = parent }

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_NextHopArray_NextHop_Labeled) GetParent() types.Entity { return labeled.parent }

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_NextHopArray_NextHop_Labeled) GetParentYangName() string { return "next-hop" }

// L2Rib_EviChildTables_Macs_Mac_Route_Bmac
// BMAC route
type L2Rib_EviChildTables_Macs_Mac_Route_Bmac struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BMAC Address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    BmacAddress interface{}

    // Forwarding State. True means forward, False means drop. The type is bool.
    ForwardState interface{}

    // Path list information.
    PathList L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList
}

func (bmac *L2Rib_EviChildTables_Macs_Mac_Route_Bmac) GetFilter() yfilter.YFilter { return bmac.YFilter }

func (bmac *L2Rib_EviChildTables_Macs_Mac_Route_Bmac) SetFilter(yf yfilter.YFilter) { bmac.YFilter = yf }

func (bmac *L2Rib_EviChildTables_Macs_Mac_Route_Bmac) GetGoName(yname string) string {
    if yname == "bmac-address" { return "BmacAddress" }
    if yname == "forward-state" { return "ForwardState" }
    if yname == "path-list" { return "PathList" }
    return ""
}

func (bmac *L2Rib_EviChildTables_Macs_Mac_Route_Bmac) GetSegmentPath() string {
    return "bmac"
}

func (bmac *L2Rib_EviChildTables_Macs_Mac_Route_Bmac) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "path-list" {
        return &bmac.PathList
    }
    return nil
}

func (bmac *L2Rib_EviChildTables_Macs_Mac_Route_Bmac) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["path-list"] = &bmac.PathList
    return children
}

func (bmac *L2Rib_EviChildTables_Macs_Mac_Route_Bmac) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["bmac-address"] = bmac.BmacAddress
    leafs["forward-state"] = bmac.ForwardState
    return leafs
}

func (bmac *L2Rib_EviChildTables_Macs_Mac_Route_Bmac) GetBundleName() string { return "cisco_ios_xr" }

func (bmac *L2Rib_EviChildTables_Macs_Mac_Route_Bmac) GetYangName() string { return "bmac" }

func (bmac *L2Rib_EviChildTables_Macs_Mac_Route_Bmac) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bmac *L2Rib_EviChildTables_Macs_Mac_Route_Bmac) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bmac *L2Rib_EviChildTables_Macs_Mac_Route_Bmac) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bmac *L2Rib_EviChildTables_Macs_Mac_Route_Bmac) SetParent(parent types.Entity) { bmac.parent = parent }

func (bmac *L2Rib_EviChildTables_Macs_Mac_Route_Bmac) GetParent() types.Entity { return bmac.parent }

func (bmac *L2Rib_EviChildTables_Macs_Mac_Route_Bmac) GetParentYangName() string { return "route" }

// L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList
// Path list information
type L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ID of EAD route producer. The type is interface{} with range: 0..255.
    ProducerId interface{}

    // Number of MAC routes bound to this Path list. The type is interface{} with
    // range: 0..4294967295.
    MacCount interface{}

    // Path list local Label. The type is interface{} with range: 0..4294967295.
    LocalLabel interface{}

    // Type-specific Path List info.
    PathListInfo L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo

    // Array of Next Hops for MAC Path List. The type is slice of
    // L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_NextHopArray.
    NextHopArray []L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_NextHopArray
}

func (pathList *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList) GetFilter() yfilter.YFilter { return pathList.YFilter }

func (pathList *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList) SetFilter(yf yfilter.YFilter) { pathList.YFilter = yf }

func (pathList *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList) GetGoName(yname string) string {
    if yname == "producer-id" { return "ProducerId" }
    if yname == "mac-count" { return "MacCount" }
    if yname == "local-label" { return "LocalLabel" }
    if yname == "path-list-info" { return "PathListInfo" }
    if yname == "next-hop-array" { return "NextHopArray" }
    return ""
}

func (pathList *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList) GetSegmentPath() string {
    return "path-list"
}

func (pathList *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "path-list-info" {
        return &pathList.PathListInfo
    }
    if childYangName == "next-hop-array" {
        for _, c := range pathList.NextHopArray {
            if pathList.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_NextHopArray{}
        pathList.NextHopArray = append(pathList.NextHopArray, child)
        return &pathList.NextHopArray[len(pathList.NextHopArray)-1]
    }
    return nil
}

func (pathList *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["path-list-info"] = &pathList.PathListInfo
    for i := range pathList.NextHopArray {
        children[pathList.NextHopArray[i].GetSegmentPath()] = &pathList.NextHopArray[i]
    }
    return children
}

func (pathList *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["producer-id"] = pathList.ProducerId
    leafs["mac-count"] = pathList.MacCount
    leafs["local-label"] = pathList.LocalLabel
    return leafs
}

func (pathList *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList) GetBundleName() string { return "cisco_ios_xr" }

func (pathList *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList) GetYangName() string { return "path-list" }

func (pathList *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pathList *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pathList *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pathList *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList) SetParent(parent types.Entity) { pathList.parent = parent }

func (pathList *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList) GetParent() types.Entity { return pathList.parent }

func (pathList *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList) GetParentYangName() string { return "bmac" }

// L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo
// Type-specific Path List info
type L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Type. The type is L2ribMacRoute.
    Type interface{}

    // ESI Path List.
    PathListEsi L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi

    // MAC Path List.
    PathListMac L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListMac
}

func (pathListInfo *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo) GetFilter() yfilter.YFilter { return pathListInfo.YFilter }

func (pathListInfo *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo) SetFilter(yf yfilter.YFilter) { pathListInfo.YFilter = yf }

func (pathListInfo *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "path-list-esi" { return "PathListEsi" }
    if yname == "path-list-mac" { return "PathListMac" }
    return ""
}

func (pathListInfo *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo) GetSegmentPath() string {
    return "path-list-info"
}

func (pathListInfo *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "path-list-esi" {
        return &pathListInfo.PathListEsi
    }
    if childYangName == "path-list-mac" {
        return &pathListInfo.PathListMac
    }
    return nil
}

func (pathListInfo *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["path-list-esi"] = &pathListInfo.PathListEsi
    children["path-list-mac"] = &pathListInfo.PathListMac
    return children
}

func (pathListInfo *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = pathListInfo.Type
    return leafs
}

func (pathListInfo *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo) GetBundleName() string { return "cisco_ios_xr" }

func (pathListInfo *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo) GetYangName() string { return "path-list-info" }

func (pathListInfo *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pathListInfo *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pathListInfo *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pathListInfo *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo) SetParent(parent types.Entity) { pathListInfo.parent = parent }

func (pathListInfo *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo) GetParent() types.Entity { return pathListInfo.parent }

func (pathListInfo *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo) GetParentYangName() string { return "path-list" }

// L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi
// ESI Path List
type L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Path list Resolved. The type is bool.
    Resolved interface{}

    // Ethernet Segment Identifier.
    EthernetSegmentId L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_EthernetSegmentId

    // Array of Next Hops from MAC Update. The type is slice of
    // L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray.
    MacUpdateNextHopArray []L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray
}

func (pathListEsi *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi) GetFilter() yfilter.YFilter { return pathListEsi.YFilter }

func (pathListEsi *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi) SetFilter(yf yfilter.YFilter) { pathListEsi.YFilter = yf }

func (pathListEsi *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi) GetGoName(yname string) string {
    if yname == "resolved" { return "Resolved" }
    if yname == "ethernet-segment-id" { return "EthernetSegmentId" }
    if yname == "mac-update-next-hop-array" { return "MacUpdateNextHopArray" }
    return ""
}

func (pathListEsi *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi) GetSegmentPath() string {
    return "path-list-esi"
}

func (pathListEsi *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ethernet-segment-id" {
        return &pathListEsi.EthernetSegmentId
    }
    if childYangName == "mac-update-next-hop-array" {
        for _, c := range pathListEsi.MacUpdateNextHopArray {
            if pathListEsi.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray{}
        pathListEsi.MacUpdateNextHopArray = append(pathListEsi.MacUpdateNextHopArray, child)
        return &pathListEsi.MacUpdateNextHopArray[len(pathListEsi.MacUpdateNextHopArray)-1]
    }
    return nil
}

func (pathListEsi *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ethernet-segment-id"] = &pathListEsi.EthernetSegmentId
    for i := range pathListEsi.MacUpdateNextHopArray {
        children[pathListEsi.MacUpdateNextHopArray[i].GetSegmentPath()] = &pathListEsi.MacUpdateNextHopArray[i]
    }
    return children
}

func (pathListEsi *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["resolved"] = pathListEsi.Resolved
    return leafs
}

func (pathListEsi *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi) GetBundleName() string { return "cisco_ios_xr" }

func (pathListEsi *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi) GetYangName() string { return "path-list-esi" }

func (pathListEsi *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pathListEsi *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pathListEsi *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pathListEsi *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi) SetParent(parent types.Entity) { pathListEsi.parent = parent }

func (pathListEsi *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi) GetParent() types.Entity { return pathListEsi.parent }

func (pathListEsi *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi) GetParentYangName() string { return "path-list-info" }

// L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_EthernetSegmentId
// Ethernet Segment Identifier
type L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_EthernetSegmentId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // LACP System Priority. The type is interface{} with range: 0..65535.
    SystemPriority interface{}

    // LACP System Id. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    SystemId interface{}

    // LACP Port Key. The type is interface{} with range: 0..65535.
    PortKey interface{}
}

func (ethernetSegmentId *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_EthernetSegmentId) GetFilter() yfilter.YFilter { return ethernetSegmentId.YFilter }

func (ethernetSegmentId *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_EthernetSegmentId) SetFilter(yf yfilter.YFilter) { ethernetSegmentId.YFilter = yf }

func (ethernetSegmentId *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_EthernetSegmentId) GetGoName(yname string) string {
    if yname == "system-priority" { return "SystemPriority" }
    if yname == "system-id" { return "SystemId" }
    if yname == "port-key" { return "PortKey" }
    return ""
}

func (ethernetSegmentId *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_EthernetSegmentId) GetSegmentPath() string {
    return "ethernet-segment-id"
}

func (ethernetSegmentId *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_EthernetSegmentId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ethernetSegmentId *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_EthernetSegmentId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ethernetSegmentId *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_EthernetSegmentId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["system-priority"] = ethernetSegmentId.SystemPriority
    leafs["system-id"] = ethernetSegmentId.SystemId
    leafs["port-key"] = ethernetSegmentId.PortKey
    return leafs
}

func (ethernetSegmentId *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_EthernetSegmentId) GetBundleName() string { return "cisco_ios_xr" }

func (ethernetSegmentId *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_EthernetSegmentId) GetYangName() string { return "ethernet-segment-id" }

func (ethernetSegmentId *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_EthernetSegmentId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ethernetSegmentId *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_EthernetSegmentId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ethernetSegmentId *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_EthernetSegmentId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ethernetSegmentId *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_EthernetSegmentId) SetParent(parent types.Entity) { ethernetSegmentId.parent = parent }

func (ethernetSegmentId *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_EthernetSegmentId) GetParent() types.Entity { return ethernetSegmentId.parent }

func (ethernetSegmentId *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_EthernetSegmentId) GetParentYangName() string { return "path-list-esi" }

// L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray
// Array of Next Hops from MAC Update
type L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Next-hop TOPOLOGY ID. The type is interface{} with range: 0..4294967295.
    TopologyId interface{}

    // Next-hop flags. The type is interface{} with range: 0..65535.
    Flags interface{}

    // Next hop.
    NextHop L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop
}

func (macUpdateNextHopArray *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray) GetFilter() yfilter.YFilter { return macUpdateNextHopArray.YFilter }

func (macUpdateNextHopArray *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray) SetFilter(yf yfilter.YFilter) { macUpdateNextHopArray.YFilter = yf }

func (macUpdateNextHopArray *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray) GetGoName(yname string) string {
    if yname == "topology-id" { return "TopologyId" }
    if yname == "flags" { return "Flags" }
    if yname == "next-hop" { return "NextHop" }
    return ""
}

func (macUpdateNextHopArray *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray) GetSegmentPath() string {
    return "mac-update-next-hop-array"
}

func (macUpdateNextHopArray *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "next-hop" {
        return &macUpdateNextHopArray.NextHop
    }
    return nil
}

func (macUpdateNextHopArray *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["next-hop"] = &macUpdateNextHopArray.NextHop
    return children
}

func (macUpdateNextHopArray *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["topology-id"] = macUpdateNextHopArray.TopologyId
    leafs["flags"] = macUpdateNextHopArray.Flags
    return leafs
}

func (macUpdateNextHopArray *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray) GetBundleName() string { return "cisco_ios_xr" }

func (macUpdateNextHopArray *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray) GetYangName() string { return "mac-update-next-hop-array" }

func (macUpdateNextHopArray *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (macUpdateNextHopArray *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (macUpdateNextHopArray *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (macUpdateNextHopArray *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray) SetParent(parent types.Entity) { macUpdateNextHopArray.parent = parent }

func (macUpdateNextHopArray *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray) GetParent() types.Entity { return macUpdateNextHopArray.parent }

func (macUpdateNextHopArray *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray) GetParentYangName() string { return "path-list-esi" }

// L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop
// Next hop
type L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Type. The type is L2ribNextHop.
    Type interface{}

    // IPV4 address Next Hop. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPV6 address Next Hop. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}

    // MAC address Next Hop. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    Mac interface{}

    // Intefrace Handle Next Hop. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceHandle interface{}

    // XID Next Hop. The type is interface{} with range: 0..4294967295.
    Xid interface{}

    // Labeled Next Hop.
    Labeled L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled
}

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop) GetFilter() yfilter.YFilter { return nextHop.YFilter }

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop) SetFilter(yf yfilter.YFilter) { nextHop.YFilter = yf }

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "ipv4" { return "Ipv4" }
    if yname == "ipv6" { return "Ipv6" }
    if yname == "mac" { return "Mac" }
    if yname == "interface-handle" { return "InterfaceHandle" }
    if yname == "xid" { return "Xid" }
    if yname == "labeled" { return "Labeled" }
    return ""
}

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop) GetSegmentPath() string {
    return "next-hop"
}

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "labeled" {
        return &nextHop.Labeled
    }
    return nil
}

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["labeled"] = &nextHop.Labeled
    return children
}

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = nextHop.Type
    leafs["ipv4"] = nextHop.Ipv4
    leafs["ipv6"] = nextHop.Ipv6
    leafs["mac"] = nextHop.Mac
    leafs["interface-handle"] = nextHop.InterfaceHandle
    leafs["xid"] = nextHop.Xid
    return leafs
}

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop) GetBundleName() string { return "cisco_ios_xr" }

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop) GetYangName() string { return "next-hop" }

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop) SetParent(parent types.Entity) { nextHop.parent = parent }

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop) GetParent() types.Entity { return nextHop.parent }

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop) GetParentYangName() string { return "mac-update-next-hop-array" }

// L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled
// Labeled Next Hop
type L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // L2RIB Address Family. The type is L2ribAfi.
    AddressFamily interface{}

    // IP Address (V6 Format). The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IpAddress interface{}

    // Label. The type is interface{} with range: 0..4294967295.
    Label interface{}

    // Internal Label. The type is bool.
    Internal interface{}
}

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled) GetFilter() yfilter.YFilter { return labeled.YFilter }

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled) SetFilter(yf yfilter.YFilter) { labeled.YFilter = yf }

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled) GetGoName(yname string) string {
    if yname == "address-family" { return "AddressFamily" }
    if yname == "ip-address" { return "IpAddress" }
    if yname == "label" { return "Label" }
    if yname == "internal" { return "Internal" }
    return ""
}

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled) GetSegmentPath() string {
    return "labeled"
}

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address-family"] = labeled.AddressFamily
    leafs["ip-address"] = labeled.IpAddress
    leafs["label"] = labeled.Label
    leafs["internal"] = labeled.Internal
    return leafs
}

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled) GetBundleName() string { return "cisco_ios_xr" }

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled) GetYangName() string { return "labeled" }

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled) SetParent(parent types.Entity) { labeled.parent = parent }

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled) GetParent() types.Entity { return labeled.parent }

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled) GetParentYangName() string { return "next-hop" }

// L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListMac
// MAC Path List
type L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListMac struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MAC Address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddress interface{}
}

func (pathListMac *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListMac) GetFilter() yfilter.YFilter { return pathListMac.YFilter }

func (pathListMac *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListMac) SetFilter(yf yfilter.YFilter) { pathListMac.YFilter = yf }

func (pathListMac *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListMac) GetGoName(yname string) string {
    if yname == "mac-address" { return "MacAddress" }
    return ""
}

func (pathListMac *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListMac) GetSegmentPath() string {
    return "path-list-mac"
}

func (pathListMac *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListMac) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pathListMac *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListMac) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pathListMac *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListMac) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mac-address"] = pathListMac.MacAddress
    return leafs
}

func (pathListMac *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListMac) GetBundleName() string { return "cisco_ios_xr" }

func (pathListMac *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListMac) GetYangName() string { return "path-list-mac" }

func (pathListMac *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListMac) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pathListMac *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListMac) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pathListMac *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListMac) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pathListMac *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListMac) SetParent(parent types.Entity) { pathListMac.parent = parent }

func (pathListMac *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListMac) GetParent() types.Entity { return pathListMac.parent }

func (pathListMac *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListMac) GetParentYangName() string { return "path-list-info" }

// L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_NextHopArray
// Array of Next Hops for MAC Path List
type L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_NextHopArray struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Next-hop TOPOLOGY ID. The type is interface{} with range: 0..4294967295.
    TopologyId interface{}

    // Next-hop flags. The type is interface{} with range: 0..65535.
    Flags interface{}

    // Next hop.
    NextHop L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_NextHopArray_NextHop
}

func (nextHopArray *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_NextHopArray) GetFilter() yfilter.YFilter { return nextHopArray.YFilter }

func (nextHopArray *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_NextHopArray) SetFilter(yf yfilter.YFilter) { nextHopArray.YFilter = yf }

func (nextHopArray *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_NextHopArray) GetGoName(yname string) string {
    if yname == "topology-id" { return "TopologyId" }
    if yname == "flags" { return "Flags" }
    if yname == "next-hop" { return "NextHop" }
    return ""
}

func (nextHopArray *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_NextHopArray) GetSegmentPath() string {
    return "next-hop-array"
}

func (nextHopArray *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_NextHopArray) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "next-hop" {
        return &nextHopArray.NextHop
    }
    return nil
}

func (nextHopArray *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_NextHopArray) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["next-hop"] = &nextHopArray.NextHop
    return children
}

func (nextHopArray *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_NextHopArray) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["topology-id"] = nextHopArray.TopologyId
    leafs["flags"] = nextHopArray.Flags
    return leafs
}

func (nextHopArray *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_NextHopArray) GetBundleName() string { return "cisco_ios_xr" }

func (nextHopArray *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_NextHopArray) GetYangName() string { return "next-hop-array" }

func (nextHopArray *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_NextHopArray) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nextHopArray *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_NextHopArray) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nextHopArray *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_NextHopArray) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nextHopArray *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_NextHopArray) SetParent(parent types.Entity) { nextHopArray.parent = parent }

func (nextHopArray *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_NextHopArray) GetParent() types.Entity { return nextHopArray.parent }

func (nextHopArray *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_NextHopArray) GetParentYangName() string { return "path-list" }

// L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_NextHopArray_NextHop
// Next hop
type L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_NextHopArray_NextHop struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Type. The type is L2ribNextHop.
    Type interface{}

    // IPV4 address Next Hop. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPV6 address Next Hop. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}

    // MAC address Next Hop. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    Mac interface{}

    // Intefrace Handle Next Hop. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceHandle interface{}

    // XID Next Hop. The type is interface{} with range: 0..4294967295.
    Xid interface{}

    // Labeled Next Hop.
    Labeled L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_NextHopArray_NextHop_Labeled
}

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_NextHopArray_NextHop) GetFilter() yfilter.YFilter { return nextHop.YFilter }

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_NextHopArray_NextHop) SetFilter(yf yfilter.YFilter) { nextHop.YFilter = yf }

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_NextHopArray_NextHop) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "ipv4" { return "Ipv4" }
    if yname == "ipv6" { return "Ipv6" }
    if yname == "mac" { return "Mac" }
    if yname == "interface-handle" { return "InterfaceHandle" }
    if yname == "xid" { return "Xid" }
    if yname == "labeled" { return "Labeled" }
    return ""
}

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_NextHopArray_NextHop) GetSegmentPath() string {
    return "next-hop"
}

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_NextHopArray_NextHop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "labeled" {
        return &nextHop.Labeled
    }
    return nil
}

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_NextHopArray_NextHop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["labeled"] = &nextHop.Labeled
    return children
}

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_NextHopArray_NextHop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = nextHop.Type
    leafs["ipv4"] = nextHop.Ipv4
    leafs["ipv6"] = nextHop.Ipv6
    leafs["mac"] = nextHop.Mac
    leafs["interface-handle"] = nextHop.InterfaceHandle
    leafs["xid"] = nextHop.Xid
    return leafs
}

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_NextHopArray_NextHop) GetBundleName() string { return "cisco_ios_xr" }

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_NextHopArray_NextHop) GetYangName() string { return "next-hop" }

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_NextHopArray_NextHop) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_NextHopArray_NextHop) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_NextHopArray_NextHop) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_NextHopArray_NextHop) SetParent(parent types.Entity) { nextHop.parent = parent }

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_NextHopArray_NextHop) GetParent() types.Entity { return nextHop.parent }

func (nextHop *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_NextHopArray_NextHop) GetParentYangName() string { return "next-hop-array" }

// L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_NextHopArray_NextHop_Labeled
// Labeled Next Hop
type L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_NextHopArray_NextHop_Labeled struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // L2RIB Address Family. The type is L2ribAfi.
    AddressFamily interface{}

    // IP Address (V6 Format). The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IpAddress interface{}

    // Label. The type is interface{} with range: 0..4294967295.
    Label interface{}

    // Internal Label. The type is bool.
    Internal interface{}
}

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_NextHopArray_NextHop_Labeled) GetFilter() yfilter.YFilter { return labeled.YFilter }

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_NextHopArray_NextHop_Labeled) SetFilter(yf yfilter.YFilter) { labeled.YFilter = yf }

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_NextHopArray_NextHop_Labeled) GetGoName(yname string) string {
    if yname == "address-family" { return "AddressFamily" }
    if yname == "ip-address" { return "IpAddress" }
    if yname == "label" { return "Label" }
    if yname == "internal" { return "Internal" }
    return ""
}

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_NextHopArray_NextHop_Labeled) GetSegmentPath() string {
    return "labeled"
}

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_NextHopArray_NextHop_Labeled) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_NextHopArray_NextHop_Labeled) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_NextHopArray_NextHop_Labeled) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address-family"] = labeled.AddressFamily
    leafs["ip-address"] = labeled.IpAddress
    leafs["label"] = labeled.Label
    leafs["internal"] = labeled.Internal
    return leafs
}

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_NextHopArray_NextHop_Labeled) GetBundleName() string { return "cisco_ios_xr" }

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_NextHopArray_NextHop_Labeled) GetYangName() string { return "labeled" }

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_NextHopArray_NextHop_Labeled) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_NextHopArray_NextHop_Labeled) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_NextHopArray_NextHop_Labeled) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_NextHopArray_NextHop_Labeled) SetParent(parent types.Entity) { labeled.parent = parent }

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_NextHopArray_NextHop_Labeled) GetParent() types.Entity { return labeled.parent }

func (labeled *L2Rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_NextHopArray_NextHop_Labeled) GetParentYangName() string { return "next-hop" }

// L2Rib_EviChildTables_MacDetails
// L2RIB EVPN EVI MAC Detail table
type L2Rib_EviChildTables_MacDetails struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // L2RIB EVPN MAC Detail table. The type is slice of
    // L2Rib_EviChildTables_MacDetails_MacDetail.
    MacDetail []L2Rib_EviChildTables_MacDetails_MacDetail
}

func (macDetails *L2Rib_EviChildTables_MacDetails) GetFilter() yfilter.YFilter { return macDetails.YFilter }

func (macDetails *L2Rib_EviChildTables_MacDetails) SetFilter(yf yfilter.YFilter) { macDetails.YFilter = yf }

func (macDetails *L2Rib_EviChildTables_MacDetails) GetGoName(yname string) string {
    if yname == "mac-detail" { return "MacDetail" }
    return ""
}

func (macDetails *L2Rib_EviChildTables_MacDetails) GetSegmentPath() string {
    return "mac-details"
}

func (macDetails *L2Rib_EviChildTables_MacDetails) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mac-detail" {
        for _, c := range macDetails.MacDetail {
            if macDetails.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := L2Rib_EviChildTables_MacDetails_MacDetail{}
        macDetails.MacDetail = append(macDetails.MacDetail, child)
        return &macDetails.MacDetail[len(macDetails.MacDetail)-1]
    }
    return nil
}

func (macDetails *L2Rib_EviChildTables_MacDetails) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range macDetails.MacDetail {
        children[macDetails.MacDetail[i].GetSegmentPath()] = &macDetails.MacDetail[i]
    }
    return children
}

func (macDetails *L2Rib_EviChildTables_MacDetails) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (macDetails *L2Rib_EviChildTables_MacDetails) GetBundleName() string { return "cisco_ios_xr" }

func (macDetails *L2Rib_EviChildTables_MacDetails) GetYangName() string { return "mac-details" }

func (macDetails *L2Rib_EviChildTables_MacDetails) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (macDetails *L2Rib_EviChildTables_MacDetails) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (macDetails *L2Rib_EviChildTables_MacDetails) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (macDetails *L2Rib_EviChildTables_MacDetails) SetParent(parent types.Entity) { macDetails.parent = parent }

func (macDetails *L2Rib_EviChildTables_MacDetails) GetParent() types.Entity { return macDetails.parent }

func (macDetails *L2Rib_EviChildTables_MacDetails) GetParentYangName() string { return "evi-child-tables" }

// L2Rib_EviChildTables_MacDetails_MacDetail
// L2RIB EVPN MAC Detail table
type L2Rib_EviChildTables_MacDetails_MacDetail struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // EVPN ID. The type is interface{} with range: -2147483648..2147483647.
    Evi interface{}

    // Tag ID. The type is interface{} with range: -2147483648..2147483647.
    TagId interface{}

    // MAC Address. The type is string with length: 1..15.
    MacAddr interface{}

    // Admin distance. The type is interface{} with range:
    // -2147483648..2147483647.
    AdminDist interface{}

    // Producer ID. The type is interface{} with range: -2147483648..2147483647.
    ProdId interface{}

    // MAC route sequence Number. The type is interface{} with range:
    // 0..4294967295.
    SequenceNumber interface{}

    // MAC route flags. The type is interface{} with range: 0..4294967295.
    Flags interface{}

    // BASE flags. The type is interface{} with range: 0..4294967295.
    Baseflags interface{}

    // SOO. The type is interface{} with range: 0..4294967295.
    Soo interface{}

    // Slot ID. The type is interface{} with range: 0..4294967295.
    SlotId interface{}

    // ESI. The type is string.
    Esi interface{}

    // Last Update Timestamp. The type is interface{} with range:
    // 0..18446744073709551615.
    LastUpdateTimestamp interface{}

    // MAC Route.
    MacRoute L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute

    // Mac Route Opaque Data TLV.
    RtTlv L2Rib_EviChildTables_MacDetails_MacDetail_RtTlv
}

func (macDetail *L2Rib_EviChildTables_MacDetails_MacDetail) GetFilter() yfilter.YFilter { return macDetail.YFilter }

func (macDetail *L2Rib_EviChildTables_MacDetails_MacDetail) SetFilter(yf yfilter.YFilter) { macDetail.YFilter = yf }

func (macDetail *L2Rib_EviChildTables_MacDetails_MacDetail) GetGoName(yname string) string {
    if yname == "evi" { return "Evi" }
    if yname == "tag-id" { return "TagId" }
    if yname == "mac-addr" { return "MacAddr" }
    if yname == "admin-dist" { return "AdminDist" }
    if yname == "prod-id" { return "ProdId" }
    if yname == "sequence-number" { return "SequenceNumber" }
    if yname == "flags" { return "Flags" }
    if yname == "baseflags" { return "Baseflags" }
    if yname == "soo" { return "Soo" }
    if yname == "slot-id" { return "SlotId" }
    if yname == "esi" { return "Esi" }
    if yname == "last-update-timestamp" { return "LastUpdateTimestamp" }
    if yname == "mac-route" { return "MacRoute" }
    if yname == "rt-tlv" { return "RtTlv" }
    return ""
}

func (macDetail *L2Rib_EviChildTables_MacDetails_MacDetail) GetSegmentPath() string {
    return "mac-detail"
}

func (macDetail *L2Rib_EviChildTables_MacDetails_MacDetail) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mac-route" {
        return &macDetail.MacRoute
    }
    if childYangName == "rt-tlv" {
        return &macDetail.RtTlv
    }
    return nil
}

func (macDetail *L2Rib_EviChildTables_MacDetails_MacDetail) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["mac-route"] = &macDetail.MacRoute
    children["rt-tlv"] = &macDetail.RtTlv
    return children
}

func (macDetail *L2Rib_EviChildTables_MacDetails_MacDetail) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["evi"] = macDetail.Evi
    leafs["tag-id"] = macDetail.TagId
    leafs["mac-addr"] = macDetail.MacAddr
    leafs["admin-dist"] = macDetail.AdminDist
    leafs["prod-id"] = macDetail.ProdId
    leafs["sequence-number"] = macDetail.SequenceNumber
    leafs["flags"] = macDetail.Flags
    leafs["baseflags"] = macDetail.Baseflags
    leafs["soo"] = macDetail.Soo
    leafs["slot-id"] = macDetail.SlotId
    leafs["esi"] = macDetail.Esi
    leafs["last-update-timestamp"] = macDetail.LastUpdateTimestamp
    return leafs
}

func (macDetail *L2Rib_EviChildTables_MacDetails_MacDetail) GetBundleName() string { return "cisco_ios_xr" }

func (macDetail *L2Rib_EviChildTables_MacDetails_MacDetail) GetYangName() string { return "mac-detail" }

func (macDetail *L2Rib_EviChildTables_MacDetails_MacDetail) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (macDetail *L2Rib_EviChildTables_MacDetails_MacDetail) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (macDetail *L2Rib_EviChildTables_MacDetails_MacDetail) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (macDetail *L2Rib_EviChildTables_MacDetails_MacDetail) SetParent(parent types.Entity) { macDetail.parent = parent }

func (macDetail *L2Rib_EviChildTables_MacDetails_MacDetail) GetParent() types.Entity { return macDetail.parent }

func (macDetail *L2Rib_EviChildTables_MacDetails_MacDetail) GetParentYangName() string { return "mac-details" }

// L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute
// MAC Route
type L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MAC Address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddress interface{}

    // Admin Distance. The type is interface{} with range: 0..255.
    AdminDistance interface{}

    // Producer ID. The type is interface{} with range: 0..255.
    ProducerId interface{}

    // Topology ID. The type is interface{} with range: 0..4294967295.
    TopologyId interface{}

    // MAC route.
    Route L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route
}

func (macRoute *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute) GetFilter() yfilter.YFilter { return macRoute.YFilter }

func (macRoute *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute) SetFilter(yf yfilter.YFilter) { macRoute.YFilter = yf }

func (macRoute *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute) GetGoName(yname string) string {
    if yname == "mac-address" { return "MacAddress" }
    if yname == "admin-distance" { return "AdminDistance" }
    if yname == "producer-id" { return "ProducerId" }
    if yname == "topology-id" { return "TopologyId" }
    if yname == "route" { return "Route" }
    return ""
}

func (macRoute *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute) GetSegmentPath() string {
    return "mac-route"
}

func (macRoute *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "route" {
        return &macRoute.Route
    }
    return nil
}

func (macRoute *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["route"] = &macRoute.Route
    return children
}

func (macRoute *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mac-address"] = macRoute.MacAddress
    leafs["admin-distance"] = macRoute.AdminDistance
    leafs["producer-id"] = macRoute.ProducerId
    leafs["topology-id"] = macRoute.TopologyId
    return leafs
}

func (macRoute *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute) GetBundleName() string { return "cisco_ios_xr" }

func (macRoute *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute) GetYangName() string { return "mac-route" }

func (macRoute *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (macRoute *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (macRoute *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (macRoute *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute) SetParent(parent types.Entity) { macRoute.parent = parent }

func (macRoute *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute) GetParent() types.Entity { return macRoute.parent }

func (macRoute *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute) GetParentYangName() string { return "mac-detail" }

// L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route
// MAC route
type L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Type. The type is L2ribMacRoute.
    Type interface{}

    // Regular MAC route.
    Regular L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular

    // EVPN ESI MAC route.
    EvpnEsi L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi

    // BMAC route.
    Bmac L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac
}

func (route *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route) GetFilter() yfilter.YFilter { return route.YFilter }

func (route *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route) SetFilter(yf yfilter.YFilter) { route.YFilter = yf }

func (route *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "regular" { return "Regular" }
    if yname == "evpn-esi" { return "EvpnEsi" }
    if yname == "bmac" { return "Bmac" }
    return ""
}

func (route *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route) GetSegmentPath() string {
    return "route"
}

func (route *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "regular" {
        return &route.Regular
    }
    if childYangName == "evpn-esi" {
        return &route.EvpnEsi
    }
    if childYangName == "bmac" {
        return &route.Bmac
    }
    return nil
}

func (route *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["regular"] = &route.Regular
    children["evpn-esi"] = &route.EvpnEsi
    children["bmac"] = &route.Bmac
    return children
}

func (route *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = route.Type
    return leafs
}

func (route *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route) GetBundleName() string { return "cisco_ios_xr" }

func (route *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route) GetYangName() string { return "route" }

func (route *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (route *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (route *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (route *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route) SetParent(parent types.Entity) { route.parent = parent }

func (route *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route) GetParent() types.Entity { return route.parent }

func (route *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route) GetParentYangName() string { return "mac-route" }

// L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular
// Regular MAC route
type L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Next Hop.
    NextHop L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular_NextHop
}

func (regular *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular) GetFilter() yfilter.YFilter { return regular.YFilter }

func (regular *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular) SetFilter(yf yfilter.YFilter) { regular.YFilter = yf }

func (regular *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular) GetGoName(yname string) string {
    if yname == "next-hop" { return "NextHop" }
    return ""
}

func (regular *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular) GetSegmentPath() string {
    return "regular"
}

func (regular *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "next-hop" {
        return &regular.NextHop
    }
    return nil
}

func (regular *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["next-hop"] = &regular.NextHop
    return children
}

func (regular *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (regular *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular) GetBundleName() string { return "cisco_ios_xr" }

func (regular *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular) GetYangName() string { return "regular" }

func (regular *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (regular *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (regular *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (regular *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular) SetParent(parent types.Entity) { regular.parent = parent }

func (regular *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular) GetParent() types.Entity { return regular.parent }

func (regular *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular) GetParentYangName() string { return "route" }

// L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular_NextHop
// Next Hop
type L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular_NextHop struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Next-hop TOPOLOGY ID. The type is interface{} with range: 0..4294967295.
    TopologyId interface{}

    // Next-hop flags. The type is interface{} with range: 0..65535.
    Flags interface{}

    // Next hop.
    NextHop L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular_NextHop_NextHop
}

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular_NextHop) GetFilter() yfilter.YFilter { return nextHop.YFilter }

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular_NextHop) SetFilter(yf yfilter.YFilter) { nextHop.YFilter = yf }

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular_NextHop) GetGoName(yname string) string {
    if yname == "topology-id" { return "TopologyId" }
    if yname == "flags" { return "Flags" }
    if yname == "next-hop" { return "NextHop" }
    return ""
}

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular_NextHop) GetSegmentPath() string {
    return "next-hop"
}

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular_NextHop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "next-hop" {
        return &nextHop.NextHop
    }
    return nil
}

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular_NextHop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["next-hop"] = &nextHop.NextHop
    return children
}

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular_NextHop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["topology-id"] = nextHop.TopologyId
    leafs["flags"] = nextHop.Flags
    return leafs
}

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular_NextHop) GetBundleName() string { return "cisco_ios_xr" }

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular_NextHop) GetYangName() string { return "next-hop" }

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular_NextHop) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular_NextHop) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular_NextHop) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular_NextHop) SetParent(parent types.Entity) { nextHop.parent = parent }

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular_NextHop) GetParent() types.Entity { return nextHop.parent }

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular_NextHop) GetParentYangName() string { return "regular" }

// L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular_NextHop_NextHop
// Next hop
type L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular_NextHop_NextHop struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Type. The type is L2ribNextHop.
    Type interface{}

    // IPV4 address Next Hop. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPV6 address Next Hop. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}

    // MAC address Next Hop. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    Mac interface{}

    // Intefrace Handle Next Hop. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceHandle interface{}

    // XID Next Hop. The type is interface{} with range: 0..4294967295.
    Xid interface{}

    // Labeled Next Hop.
    Labeled L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular_NextHop_NextHop_Labeled
}

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular_NextHop_NextHop) GetFilter() yfilter.YFilter { return nextHop.YFilter }

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular_NextHop_NextHop) SetFilter(yf yfilter.YFilter) { nextHop.YFilter = yf }

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular_NextHop_NextHop) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "ipv4" { return "Ipv4" }
    if yname == "ipv6" { return "Ipv6" }
    if yname == "mac" { return "Mac" }
    if yname == "interface-handle" { return "InterfaceHandle" }
    if yname == "xid" { return "Xid" }
    if yname == "labeled" { return "Labeled" }
    return ""
}

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular_NextHop_NextHop) GetSegmentPath() string {
    return "next-hop"
}

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular_NextHop_NextHop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "labeled" {
        return &nextHop.Labeled
    }
    return nil
}

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular_NextHop_NextHop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["labeled"] = &nextHop.Labeled
    return children
}

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular_NextHop_NextHop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = nextHop.Type
    leafs["ipv4"] = nextHop.Ipv4
    leafs["ipv6"] = nextHop.Ipv6
    leafs["mac"] = nextHop.Mac
    leafs["interface-handle"] = nextHop.InterfaceHandle
    leafs["xid"] = nextHop.Xid
    return leafs
}

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular_NextHop_NextHop) GetBundleName() string { return "cisco_ios_xr" }

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular_NextHop_NextHop) GetYangName() string { return "next-hop" }

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular_NextHop_NextHop) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular_NextHop_NextHop) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular_NextHop_NextHop) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular_NextHop_NextHop) SetParent(parent types.Entity) { nextHop.parent = parent }

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular_NextHop_NextHop) GetParent() types.Entity { return nextHop.parent }

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular_NextHop_NextHop) GetParentYangName() string { return "next-hop" }

// L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular_NextHop_NextHop_Labeled
// Labeled Next Hop
type L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular_NextHop_NextHop_Labeled struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // L2RIB Address Family. The type is L2ribAfi.
    AddressFamily interface{}

    // IP Address (V6 Format). The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IpAddress interface{}

    // Label. The type is interface{} with range: 0..4294967295.
    Label interface{}

    // Internal Label. The type is bool.
    Internal interface{}
}

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular_NextHop_NextHop_Labeled) GetFilter() yfilter.YFilter { return labeled.YFilter }

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular_NextHop_NextHop_Labeled) SetFilter(yf yfilter.YFilter) { labeled.YFilter = yf }

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular_NextHop_NextHop_Labeled) GetGoName(yname string) string {
    if yname == "address-family" { return "AddressFamily" }
    if yname == "ip-address" { return "IpAddress" }
    if yname == "label" { return "Label" }
    if yname == "internal" { return "Internal" }
    return ""
}

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular_NextHop_NextHop_Labeled) GetSegmentPath() string {
    return "labeled"
}

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular_NextHop_NextHop_Labeled) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular_NextHop_NextHop_Labeled) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular_NextHop_NextHop_Labeled) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address-family"] = labeled.AddressFamily
    leafs["ip-address"] = labeled.IpAddress
    leafs["label"] = labeled.Label
    leafs["internal"] = labeled.Internal
    return leafs
}

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular_NextHop_NextHop_Labeled) GetBundleName() string { return "cisco_ios_xr" }

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular_NextHop_NextHop_Labeled) GetYangName() string { return "labeled" }

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular_NextHop_NextHop_Labeled) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular_NextHop_NextHop_Labeled) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular_NextHop_NextHop_Labeled) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular_NextHop_NextHop_Labeled) SetParent(parent types.Entity) { labeled.parent = parent }

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular_NextHop_NextHop_Labeled) GetParent() types.Entity { return labeled.parent }

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular_NextHop_NextHop_Labeled) GetParentYangName() string { return "next-hop" }

// L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi
// EVPN ESI MAC route
type L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MAC route sequence number. The type is interface{} with range:
    // 0..4294967295.
    SequenceNumber interface{}

    // Forwarding State. True means forward, False means drop. The type is bool.
    ForwardState interface{}

    // Ethernet Segment Identifier.
    EthernetSegmentId L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_EthernetSegmentId

    // Path list information. Set for detailed MAC route information.
    PathList L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList
}

func (evpnEsi *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi) GetFilter() yfilter.YFilter { return evpnEsi.YFilter }

func (evpnEsi *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi) SetFilter(yf yfilter.YFilter) { evpnEsi.YFilter = yf }

func (evpnEsi *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi) GetGoName(yname string) string {
    if yname == "sequence-number" { return "SequenceNumber" }
    if yname == "forward-state" { return "ForwardState" }
    if yname == "ethernet-segment-id" { return "EthernetSegmentId" }
    if yname == "path-list" { return "PathList" }
    return ""
}

func (evpnEsi *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi) GetSegmentPath() string {
    return "evpn-esi"
}

func (evpnEsi *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ethernet-segment-id" {
        return &evpnEsi.EthernetSegmentId
    }
    if childYangName == "path-list" {
        return &evpnEsi.PathList
    }
    return nil
}

func (evpnEsi *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ethernet-segment-id"] = &evpnEsi.EthernetSegmentId
    children["path-list"] = &evpnEsi.PathList
    return children
}

func (evpnEsi *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sequence-number"] = evpnEsi.SequenceNumber
    leafs["forward-state"] = evpnEsi.ForwardState
    return leafs
}

func (evpnEsi *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi) GetBundleName() string { return "cisco_ios_xr" }

func (evpnEsi *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi) GetYangName() string { return "evpn-esi" }

func (evpnEsi *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (evpnEsi *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (evpnEsi *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (evpnEsi *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi) SetParent(parent types.Entity) { evpnEsi.parent = parent }

func (evpnEsi *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi) GetParent() types.Entity { return evpnEsi.parent }

func (evpnEsi *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi) GetParentYangName() string { return "route" }

// L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_EthernetSegmentId
// Ethernet Segment Identifier
type L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_EthernetSegmentId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // LACP System Priority. The type is interface{} with range: 0..65535.
    SystemPriority interface{}

    // LACP System Id. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    SystemId interface{}

    // LACP Port Key. The type is interface{} with range: 0..65535.
    PortKey interface{}
}

func (ethernetSegmentId *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_EthernetSegmentId) GetFilter() yfilter.YFilter { return ethernetSegmentId.YFilter }

func (ethernetSegmentId *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_EthernetSegmentId) SetFilter(yf yfilter.YFilter) { ethernetSegmentId.YFilter = yf }

func (ethernetSegmentId *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_EthernetSegmentId) GetGoName(yname string) string {
    if yname == "system-priority" { return "SystemPriority" }
    if yname == "system-id" { return "SystemId" }
    if yname == "port-key" { return "PortKey" }
    return ""
}

func (ethernetSegmentId *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_EthernetSegmentId) GetSegmentPath() string {
    return "ethernet-segment-id"
}

func (ethernetSegmentId *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_EthernetSegmentId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ethernetSegmentId *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_EthernetSegmentId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ethernetSegmentId *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_EthernetSegmentId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["system-priority"] = ethernetSegmentId.SystemPriority
    leafs["system-id"] = ethernetSegmentId.SystemId
    leafs["port-key"] = ethernetSegmentId.PortKey
    return leafs
}

func (ethernetSegmentId *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_EthernetSegmentId) GetBundleName() string { return "cisco_ios_xr" }

func (ethernetSegmentId *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_EthernetSegmentId) GetYangName() string { return "ethernet-segment-id" }

func (ethernetSegmentId *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_EthernetSegmentId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ethernetSegmentId *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_EthernetSegmentId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ethernetSegmentId *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_EthernetSegmentId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ethernetSegmentId *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_EthernetSegmentId) SetParent(parent types.Entity) { ethernetSegmentId.parent = parent }

func (ethernetSegmentId *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_EthernetSegmentId) GetParent() types.Entity { return ethernetSegmentId.parent }

func (ethernetSegmentId *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_EthernetSegmentId) GetParentYangName() string { return "evpn-esi" }

// L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList
// Path list information. Set for detailed MAC
// route information
type L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ID of EAD route producer. The type is interface{} with range: 0..255.
    ProducerId interface{}

    // Number of MAC routes bound to this Path list. The type is interface{} with
    // range: 0..4294967295.
    MacCount interface{}

    // Path list local Label. The type is interface{} with range: 0..4294967295.
    LocalLabel interface{}

    // Type-specific Path List info.
    PathListInfo L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo

    // Array of Next Hops for MAC Path List. The type is slice of
    // L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_NextHopArray.
    NextHopArray []L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_NextHopArray
}

func (pathList *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList) GetFilter() yfilter.YFilter { return pathList.YFilter }

func (pathList *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList) SetFilter(yf yfilter.YFilter) { pathList.YFilter = yf }

func (pathList *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList) GetGoName(yname string) string {
    if yname == "producer-id" { return "ProducerId" }
    if yname == "mac-count" { return "MacCount" }
    if yname == "local-label" { return "LocalLabel" }
    if yname == "path-list-info" { return "PathListInfo" }
    if yname == "next-hop-array" { return "NextHopArray" }
    return ""
}

func (pathList *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList) GetSegmentPath() string {
    return "path-list"
}

func (pathList *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "path-list-info" {
        return &pathList.PathListInfo
    }
    if childYangName == "next-hop-array" {
        for _, c := range pathList.NextHopArray {
            if pathList.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_NextHopArray{}
        pathList.NextHopArray = append(pathList.NextHopArray, child)
        return &pathList.NextHopArray[len(pathList.NextHopArray)-1]
    }
    return nil
}

func (pathList *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["path-list-info"] = &pathList.PathListInfo
    for i := range pathList.NextHopArray {
        children[pathList.NextHopArray[i].GetSegmentPath()] = &pathList.NextHopArray[i]
    }
    return children
}

func (pathList *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["producer-id"] = pathList.ProducerId
    leafs["mac-count"] = pathList.MacCount
    leafs["local-label"] = pathList.LocalLabel
    return leafs
}

func (pathList *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList) GetBundleName() string { return "cisco_ios_xr" }

func (pathList *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList) GetYangName() string { return "path-list" }

func (pathList *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pathList *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pathList *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pathList *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList) SetParent(parent types.Entity) { pathList.parent = parent }

func (pathList *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList) GetParent() types.Entity { return pathList.parent }

func (pathList *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList) GetParentYangName() string { return "evpn-esi" }

// L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo
// Type-specific Path List info
type L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Type. The type is L2ribMacRoute.
    Type interface{}

    // ESI Path List.
    PathListEsi L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi

    // MAC Path List.
    PathListMac L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListMac
}

func (pathListInfo *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo) GetFilter() yfilter.YFilter { return pathListInfo.YFilter }

func (pathListInfo *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo) SetFilter(yf yfilter.YFilter) { pathListInfo.YFilter = yf }

func (pathListInfo *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "path-list-esi" { return "PathListEsi" }
    if yname == "path-list-mac" { return "PathListMac" }
    return ""
}

func (pathListInfo *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo) GetSegmentPath() string {
    return "path-list-info"
}

func (pathListInfo *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "path-list-esi" {
        return &pathListInfo.PathListEsi
    }
    if childYangName == "path-list-mac" {
        return &pathListInfo.PathListMac
    }
    return nil
}

func (pathListInfo *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["path-list-esi"] = &pathListInfo.PathListEsi
    children["path-list-mac"] = &pathListInfo.PathListMac
    return children
}

func (pathListInfo *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = pathListInfo.Type
    return leafs
}

func (pathListInfo *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo) GetBundleName() string { return "cisco_ios_xr" }

func (pathListInfo *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo) GetYangName() string { return "path-list-info" }

func (pathListInfo *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pathListInfo *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pathListInfo *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pathListInfo *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo) SetParent(parent types.Entity) { pathListInfo.parent = parent }

func (pathListInfo *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo) GetParent() types.Entity { return pathListInfo.parent }

func (pathListInfo *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo) GetParentYangName() string { return "path-list" }

// L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi
// ESI Path List
type L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Path list Resolved. The type is bool.
    Resolved interface{}

    // Ethernet Segment Identifier.
    EthernetSegmentId L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_EthernetSegmentId

    // Array of Next Hops from MAC Update. The type is slice of
    // L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray.
    MacUpdateNextHopArray []L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray
}

func (pathListEsi *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi) GetFilter() yfilter.YFilter { return pathListEsi.YFilter }

func (pathListEsi *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi) SetFilter(yf yfilter.YFilter) { pathListEsi.YFilter = yf }

func (pathListEsi *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi) GetGoName(yname string) string {
    if yname == "resolved" { return "Resolved" }
    if yname == "ethernet-segment-id" { return "EthernetSegmentId" }
    if yname == "mac-update-next-hop-array" { return "MacUpdateNextHopArray" }
    return ""
}

func (pathListEsi *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi) GetSegmentPath() string {
    return "path-list-esi"
}

func (pathListEsi *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ethernet-segment-id" {
        return &pathListEsi.EthernetSegmentId
    }
    if childYangName == "mac-update-next-hop-array" {
        for _, c := range pathListEsi.MacUpdateNextHopArray {
            if pathListEsi.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray{}
        pathListEsi.MacUpdateNextHopArray = append(pathListEsi.MacUpdateNextHopArray, child)
        return &pathListEsi.MacUpdateNextHopArray[len(pathListEsi.MacUpdateNextHopArray)-1]
    }
    return nil
}

func (pathListEsi *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ethernet-segment-id"] = &pathListEsi.EthernetSegmentId
    for i := range pathListEsi.MacUpdateNextHopArray {
        children[pathListEsi.MacUpdateNextHopArray[i].GetSegmentPath()] = &pathListEsi.MacUpdateNextHopArray[i]
    }
    return children
}

func (pathListEsi *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["resolved"] = pathListEsi.Resolved
    return leafs
}

func (pathListEsi *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi) GetBundleName() string { return "cisco_ios_xr" }

func (pathListEsi *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi) GetYangName() string { return "path-list-esi" }

func (pathListEsi *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pathListEsi *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pathListEsi *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pathListEsi *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi) SetParent(parent types.Entity) { pathListEsi.parent = parent }

func (pathListEsi *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi) GetParent() types.Entity { return pathListEsi.parent }

func (pathListEsi *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi) GetParentYangName() string { return "path-list-info" }

// L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_EthernetSegmentId
// Ethernet Segment Identifier
type L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_EthernetSegmentId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // LACP System Priority. The type is interface{} with range: 0..65535.
    SystemPriority interface{}

    // LACP System Id. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    SystemId interface{}

    // LACP Port Key. The type is interface{} with range: 0..65535.
    PortKey interface{}
}

func (ethernetSegmentId *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_EthernetSegmentId) GetFilter() yfilter.YFilter { return ethernetSegmentId.YFilter }

func (ethernetSegmentId *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_EthernetSegmentId) SetFilter(yf yfilter.YFilter) { ethernetSegmentId.YFilter = yf }

func (ethernetSegmentId *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_EthernetSegmentId) GetGoName(yname string) string {
    if yname == "system-priority" { return "SystemPriority" }
    if yname == "system-id" { return "SystemId" }
    if yname == "port-key" { return "PortKey" }
    return ""
}

func (ethernetSegmentId *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_EthernetSegmentId) GetSegmentPath() string {
    return "ethernet-segment-id"
}

func (ethernetSegmentId *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_EthernetSegmentId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ethernetSegmentId *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_EthernetSegmentId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ethernetSegmentId *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_EthernetSegmentId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["system-priority"] = ethernetSegmentId.SystemPriority
    leafs["system-id"] = ethernetSegmentId.SystemId
    leafs["port-key"] = ethernetSegmentId.PortKey
    return leafs
}

func (ethernetSegmentId *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_EthernetSegmentId) GetBundleName() string { return "cisco_ios_xr" }

func (ethernetSegmentId *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_EthernetSegmentId) GetYangName() string { return "ethernet-segment-id" }

func (ethernetSegmentId *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_EthernetSegmentId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ethernetSegmentId *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_EthernetSegmentId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ethernetSegmentId *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_EthernetSegmentId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ethernetSegmentId *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_EthernetSegmentId) SetParent(parent types.Entity) { ethernetSegmentId.parent = parent }

func (ethernetSegmentId *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_EthernetSegmentId) GetParent() types.Entity { return ethernetSegmentId.parent }

func (ethernetSegmentId *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_EthernetSegmentId) GetParentYangName() string { return "path-list-esi" }

// L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray
// Array of Next Hops from MAC Update
type L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Next-hop TOPOLOGY ID. The type is interface{} with range: 0..4294967295.
    TopologyId interface{}

    // Next-hop flags. The type is interface{} with range: 0..65535.
    Flags interface{}

    // Next hop.
    NextHop L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop
}

func (macUpdateNextHopArray *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray) GetFilter() yfilter.YFilter { return macUpdateNextHopArray.YFilter }

func (macUpdateNextHopArray *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray) SetFilter(yf yfilter.YFilter) { macUpdateNextHopArray.YFilter = yf }

func (macUpdateNextHopArray *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray) GetGoName(yname string) string {
    if yname == "topology-id" { return "TopologyId" }
    if yname == "flags" { return "Flags" }
    if yname == "next-hop" { return "NextHop" }
    return ""
}

func (macUpdateNextHopArray *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray) GetSegmentPath() string {
    return "mac-update-next-hop-array"
}

func (macUpdateNextHopArray *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "next-hop" {
        return &macUpdateNextHopArray.NextHop
    }
    return nil
}

func (macUpdateNextHopArray *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["next-hop"] = &macUpdateNextHopArray.NextHop
    return children
}

func (macUpdateNextHopArray *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["topology-id"] = macUpdateNextHopArray.TopologyId
    leafs["flags"] = macUpdateNextHopArray.Flags
    return leafs
}

func (macUpdateNextHopArray *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray) GetBundleName() string { return "cisco_ios_xr" }

func (macUpdateNextHopArray *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray) GetYangName() string { return "mac-update-next-hop-array" }

func (macUpdateNextHopArray *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (macUpdateNextHopArray *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (macUpdateNextHopArray *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (macUpdateNextHopArray *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray) SetParent(parent types.Entity) { macUpdateNextHopArray.parent = parent }

func (macUpdateNextHopArray *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray) GetParent() types.Entity { return macUpdateNextHopArray.parent }

func (macUpdateNextHopArray *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray) GetParentYangName() string { return "path-list-esi" }

// L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop
// Next hop
type L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Type. The type is L2ribNextHop.
    Type interface{}

    // IPV4 address Next Hop. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPV6 address Next Hop. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}

    // MAC address Next Hop. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    Mac interface{}

    // Intefrace Handle Next Hop. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceHandle interface{}

    // XID Next Hop. The type is interface{} with range: 0..4294967295.
    Xid interface{}

    // Labeled Next Hop.
    Labeled L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled
}

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop) GetFilter() yfilter.YFilter { return nextHop.YFilter }

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop) SetFilter(yf yfilter.YFilter) { nextHop.YFilter = yf }

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "ipv4" { return "Ipv4" }
    if yname == "ipv6" { return "Ipv6" }
    if yname == "mac" { return "Mac" }
    if yname == "interface-handle" { return "InterfaceHandle" }
    if yname == "xid" { return "Xid" }
    if yname == "labeled" { return "Labeled" }
    return ""
}

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop) GetSegmentPath() string {
    return "next-hop"
}

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "labeled" {
        return &nextHop.Labeled
    }
    return nil
}

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["labeled"] = &nextHop.Labeled
    return children
}

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = nextHop.Type
    leafs["ipv4"] = nextHop.Ipv4
    leafs["ipv6"] = nextHop.Ipv6
    leafs["mac"] = nextHop.Mac
    leafs["interface-handle"] = nextHop.InterfaceHandle
    leafs["xid"] = nextHop.Xid
    return leafs
}

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop) GetBundleName() string { return "cisco_ios_xr" }

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop) GetYangName() string { return "next-hop" }

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop) SetParent(parent types.Entity) { nextHop.parent = parent }

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop) GetParent() types.Entity { return nextHop.parent }

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop) GetParentYangName() string { return "mac-update-next-hop-array" }

// L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled
// Labeled Next Hop
type L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // L2RIB Address Family. The type is L2ribAfi.
    AddressFamily interface{}

    // IP Address (V6 Format). The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IpAddress interface{}

    // Label. The type is interface{} with range: 0..4294967295.
    Label interface{}

    // Internal Label. The type is bool.
    Internal interface{}
}

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled) GetFilter() yfilter.YFilter { return labeled.YFilter }

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled) SetFilter(yf yfilter.YFilter) { labeled.YFilter = yf }

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled) GetGoName(yname string) string {
    if yname == "address-family" { return "AddressFamily" }
    if yname == "ip-address" { return "IpAddress" }
    if yname == "label" { return "Label" }
    if yname == "internal" { return "Internal" }
    return ""
}

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled) GetSegmentPath() string {
    return "labeled"
}

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address-family"] = labeled.AddressFamily
    leafs["ip-address"] = labeled.IpAddress
    leafs["label"] = labeled.Label
    leafs["internal"] = labeled.Internal
    return leafs
}

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled) GetBundleName() string { return "cisco_ios_xr" }

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled) GetYangName() string { return "labeled" }

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled) SetParent(parent types.Entity) { labeled.parent = parent }

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled) GetParent() types.Entity { return labeled.parent }

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled) GetParentYangName() string { return "next-hop" }

// L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListMac
// MAC Path List
type L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListMac struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MAC Address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddress interface{}
}

func (pathListMac *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListMac) GetFilter() yfilter.YFilter { return pathListMac.YFilter }

func (pathListMac *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListMac) SetFilter(yf yfilter.YFilter) { pathListMac.YFilter = yf }

func (pathListMac *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListMac) GetGoName(yname string) string {
    if yname == "mac-address" { return "MacAddress" }
    return ""
}

func (pathListMac *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListMac) GetSegmentPath() string {
    return "path-list-mac"
}

func (pathListMac *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListMac) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pathListMac *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListMac) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pathListMac *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListMac) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mac-address"] = pathListMac.MacAddress
    return leafs
}

func (pathListMac *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListMac) GetBundleName() string { return "cisco_ios_xr" }

func (pathListMac *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListMac) GetYangName() string { return "path-list-mac" }

func (pathListMac *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListMac) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pathListMac *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListMac) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pathListMac *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListMac) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pathListMac *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListMac) SetParent(parent types.Entity) { pathListMac.parent = parent }

func (pathListMac *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListMac) GetParent() types.Entity { return pathListMac.parent }

func (pathListMac *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListMac) GetParentYangName() string { return "path-list-info" }

// L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_NextHopArray
// Array of Next Hops for MAC Path List
type L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_NextHopArray struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Next-hop TOPOLOGY ID. The type is interface{} with range: 0..4294967295.
    TopologyId interface{}

    // Next-hop flags. The type is interface{} with range: 0..65535.
    Flags interface{}

    // Next hop.
    NextHop L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_NextHopArray_NextHop
}

func (nextHopArray *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_NextHopArray) GetFilter() yfilter.YFilter { return nextHopArray.YFilter }

func (nextHopArray *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_NextHopArray) SetFilter(yf yfilter.YFilter) { nextHopArray.YFilter = yf }

func (nextHopArray *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_NextHopArray) GetGoName(yname string) string {
    if yname == "topology-id" { return "TopologyId" }
    if yname == "flags" { return "Flags" }
    if yname == "next-hop" { return "NextHop" }
    return ""
}

func (nextHopArray *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_NextHopArray) GetSegmentPath() string {
    return "next-hop-array"
}

func (nextHopArray *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_NextHopArray) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "next-hop" {
        return &nextHopArray.NextHop
    }
    return nil
}

func (nextHopArray *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_NextHopArray) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["next-hop"] = &nextHopArray.NextHop
    return children
}

func (nextHopArray *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_NextHopArray) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["topology-id"] = nextHopArray.TopologyId
    leafs["flags"] = nextHopArray.Flags
    return leafs
}

func (nextHopArray *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_NextHopArray) GetBundleName() string { return "cisco_ios_xr" }

func (nextHopArray *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_NextHopArray) GetYangName() string { return "next-hop-array" }

func (nextHopArray *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_NextHopArray) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nextHopArray *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_NextHopArray) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nextHopArray *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_NextHopArray) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nextHopArray *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_NextHopArray) SetParent(parent types.Entity) { nextHopArray.parent = parent }

func (nextHopArray *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_NextHopArray) GetParent() types.Entity { return nextHopArray.parent }

func (nextHopArray *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_NextHopArray) GetParentYangName() string { return "path-list" }

// L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_NextHopArray_NextHop
// Next hop
type L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_NextHopArray_NextHop struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Type. The type is L2ribNextHop.
    Type interface{}

    // IPV4 address Next Hop. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPV6 address Next Hop. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}

    // MAC address Next Hop. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    Mac interface{}

    // Intefrace Handle Next Hop. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceHandle interface{}

    // XID Next Hop. The type is interface{} with range: 0..4294967295.
    Xid interface{}

    // Labeled Next Hop.
    Labeled L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_NextHopArray_NextHop_Labeled
}

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_NextHopArray_NextHop) GetFilter() yfilter.YFilter { return nextHop.YFilter }

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_NextHopArray_NextHop) SetFilter(yf yfilter.YFilter) { nextHop.YFilter = yf }

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_NextHopArray_NextHop) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "ipv4" { return "Ipv4" }
    if yname == "ipv6" { return "Ipv6" }
    if yname == "mac" { return "Mac" }
    if yname == "interface-handle" { return "InterfaceHandle" }
    if yname == "xid" { return "Xid" }
    if yname == "labeled" { return "Labeled" }
    return ""
}

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_NextHopArray_NextHop) GetSegmentPath() string {
    return "next-hop"
}

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_NextHopArray_NextHop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "labeled" {
        return &nextHop.Labeled
    }
    return nil
}

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_NextHopArray_NextHop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["labeled"] = &nextHop.Labeled
    return children
}

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_NextHopArray_NextHop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = nextHop.Type
    leafs["ipv4"] = nextHop.Ipv4
    leafs["ipv6"] = nextHop.Ipv6
    leafs["mac"] = nextHop.Mac
    leafs["interface-handle"] = nextHop.InterfaceHandle
    leafs["xid"] = nextHop.Xid
    return leafs
}

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_NextHopArray_NextHop) GetBundleName() string { return "cisco_ios_xr" }

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_NextHopArray_NextHop) GetYangName() string { return "next-hop" }

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_NextHopArray_NextHop) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_NextHopArray_NextHop) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_NextHopArray_NextHop) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_NextHopArray_NextHop) SetParent(parent types.Entity) { nextHop.parent = parent }

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_NextHopArray_NextHop) GetParent() types.Entity { return nextHop.parent }

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_NextHopArray_NextHop) GetParentYangName() string { return "next-hop-array" }

// L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_NextHopArray_NextHop_Labeled
// Labeled Next Hop
type L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_NextHopArray_NextHop_Labeled struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // L2RIB Address Family. The type is L2ribAfi.
    AddressFamily interface{}

    // IP Address (V6 Format). The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IpAddress interface{}

    // Label. The type is interface{} with range: 0..4294967295.
    Label interface{}

    // Internal Label. The type is bool.
    Internal interface{}
}

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_NextHopArray_NextHop_Labeled) GetFilter() yfilter.YFilter { return labeled.YFilter }

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_NextHopArray_NextHop_Labeled) SetFilter(yf yfilter.YFilter) { labeled.YFilter = yf }

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_NextHopArray_NextHop_Labeled) GetGoName(yname string) string {
    if yname == "address-family" { return "AddressFamily" }
    if yname == "ip-address" { return "IpAddress" }
    if yname == "label" { return "Label" }
    if yname == "internal" { return "Internal" }
    return ""
}

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_NextHopArray_NextHop_Labeled) GetSegmentPath() string {
    return "labeled"
}

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_NextHopArray_NextHop_Labeled) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_NextHopArray_NextHop_Labeled) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_NextHopArray_NextHop_Labeled) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address-family"] = labeled.AddressFamily
    leafs["ip-address"] = labeled.IpAddress
    leafs["label"] = labeled.Label
    leafs["internal"] = labeled.Internal
    return leafs
}

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_NextHopArray_NextHop_Labeled) GetBundleName() string { return "cisco_ios_xr" }

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_NextHopArray_NextHop_Labeled) GetYangName() string { return "labeled" }

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_NextHopArray_NextHop_Labeled) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_NextHopArray_NextHop_Labeled) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_NextHopArray_NextHop_Labeled) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_NextHopArray_NextHop_Labeled) SetParent(parent types.Entity) { labeled.parent = parent }

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_NextHopArray_NextHop_Labeled) GetParent() types.Entity { return labeled.parent }

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_NextHopArray_NextHop_Labeled) GetParentYangName() string { return "next-hop" }

// L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac
// BMAC route
type L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BMAC Address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    BmacAddress interface{}

    // Forwarding State. True means forward, False means drop. The type is bool.
    ForwardState interface{}

    // Path list information.
    PathList L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList
}

func (bmac *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac) GetFilter() yfilter.YFilter { return bmac.YFilter }

func (bmac *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac) SetFilter(yf yfilter.YFilter) { bmac.YFilter = yf }

func (bmac *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac) GetGoName(yname string) string {
    if yname == "bmac-address" { return "BmacAddress" }
    if yname == "forward-state" { return "ForwardState" }
    if yname == "path-list" { return "PathList" }
    return ""
}

func (bmac *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac) GetSegmentPath() string {
    return "bmac"
}

func (bmac *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "path-list" {
        return &bmac.PathList
    }
    return nil
}

func (bmac *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["path-list"] = &bmac.PathList
    return children
}

func (bmac *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["bmac-address"] = bmac.BmacAddress
    leafs["forward-state"] = bmac.ForwardState
    return leafs
}

func (bmac *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac) GetBundleName() string { return "cisco_ios_xr" }

func (bmac *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac) GetYangName() string { return "bmac" }

func (bmac *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bmac *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bmac *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bmac *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac) SetParent(parent types.Entity) { bmac.parent = parent }

func (bmac *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac) GetParent() types.Entity { return bmac.parent }

func (bmac *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac) GetParentYangName() string { return "route" }

// L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList
// Path list information
type L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ID of EAD route producer. The type is interface{} with range: 0..255.
    ProducerId interface{}

    // Number of MAC routes bound to this Path list. The type is interface{} with
    // range: 0..4294967295.
    MacCount interface{}

    // Path list local Label. The type is interface{} with range: 0..4294967295.
    LocalLabel interface{}

    // Type-specific Path List info.
    PathListInfo L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo

    // Array of Next Hops for MAC Path List. The type is slice of
    // L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_NextHopArray.
    NextHopArray []L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_NextHopArray
}

func (pathList *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList) GetFilter() yfilter.YFilter { return pathList.YFilter }

func (pathList *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList) SetFilter(yf yfilter.YFilter) { pathList.YFilter = yf }

func (pathList *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList) GetGoName(yname string) string {
    if yname == "producer-id" { return "ProducerId" }
    if yname == "mac-count" { return "MacCount" }
    if yname == "local-label" { return "LocalLabel" }
    if yname == "path-list-info" { return "PathListInfo" }
    if yname == "next-hop-array" { return "NextHopArray" }
    return ""
}

func (pathList *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList) GetSegmentPath() string {
    return "path-list"
}

func (pathList *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "path-list-info" {
        return &pathList.PathListInfo
    }
    if childYangName == "next-hop-array" {
        for _, c := range pathList.NextHopArray {
            if pathList.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_NextHopArray{}
        pathList.NextHopArray = append(pathList.NextHopArray, child)
        return &pathList.NextHopArray[len(pathList.NextHopArray)-1]
    }
    return nil
}

func (pathList *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["path-list-info"] = &pathList.PathListInfo
    for i := range pathList.NextHopArray {
        children[pathList.NextHopArray[i].GetSegmentPath()] = &pathList.NextHopArray[i]
    }
    return children
}

func (pathList *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["producer-id"] = pathList.ProducerId
    leafs["mac-count"] = pathList.MacCount
    leafs["local-label"] = pathList.LocalLabel
    return leafs
}

func (pathList *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList) GetBundleName() string { return "cisco_ios_xr" }

func (pathList *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList) GetYangName() string { return "path-list" }

func (pathList *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pathList *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pathList *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pathList *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList) SetParent(parent types.Entity) { pathList.parent = parent }

func (pathList *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList) GetParent() types.Entity { return pathList.parent }

func (pathList *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList) GetParentYangName() string { return "bmac" }

// L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo
// Type-specific Path List info
type L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Type. The type is L2ribMacRoute.
    Type interface{}

    // ESI Path List.
    PathListEsi L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi

    // MAC Path List.
    PathListMac L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListMac
}

func (pathListInfo *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo) GetFilter() yfilter.YFilter { return pathListInfo.YFilter }

func (pathListInfo *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo) SetFilter(yf yfilter.YFilter) { pathListInfo.YFilter = yf }

func (pathListInfo *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "path-list-esi" { return "PathListEsi" }
    if yname == "path-list-mac" { return "PathListMac" }
    return ""
}

func (pathListInfo *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo) GetSegmentPath() string {
    return "path-list-info"
}

func (pathListInfo *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "path-list-esi" {
        return &pathListInfo.PathListEsi
    }
    if childYangName == "path-list-mac" {
        return &pathListInfo.PathListMac
    }
    return nil
}

func (pathListInfo *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["path-list-esi"] = &pathListInfo.PathListEsi
    children["path-list-mac"] = &pathListInfo.PathListMac
    return children
}

func (pathListInfo *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = pathListInfo.Type
    return leafs
}

func (pathListInfo *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo) GetBundleName() string { return "cisco_ios_xr" }

func (pathListInfo *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo) GetYangName() string { return "path-list-info" }

func (pathListInfo *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pathListInfo *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pathListInfo *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pathListInfo *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo) SetParent(parent types.Entity) { pathListInfo.parent = parent }

func (pathListInfo *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo) GetParent() types.Entity { return pathListInfo.parent }

func (pathListInfo *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo) GetParentYangName() string { return "path-list" }

// L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi
// ESI Path List
type L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Path list Resolved. The type is bool.
    Resolved interface{}

    // Ethernet Segment Identifier.
    EthernetSegmentId L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_EthernetSegmentId

    // Array of Next Hops from MAC Update. The type is slice of
    // L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray.
    MacUpdateNextHopArray []L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray
}

func (pathListEsi *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi) GetFilter() yfilter.YFilter { return pathListEsi.YFilter }

func (pathListEsi *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi) SetFilter(yf yfilter.YFilter) { pathListEsi.YFilter = yf }

func (pathListEsi *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi) GetGoName(yname string) string {
    if yname == "resolved" { return "Resolved" }
    if yname == "ethernet-segment-id" { return "EthernetSegmentId" }
    if yname == "mac-update-next-hop-array" { return "MacUpdateNextHopArray" }
    return ""
}

func (pathListEsi *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi) GetSegmentPath() string {
    return "path-list-esi"
}

func (pathListEsi *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ethernet-segment-id" {
        return &pathListEsi.EthernetSegmentId
    }
    if childYangName == "mac-update-next-hop-array" {
        for _, c := range pathListEsi.MacUpdateNextHopArray {
            if pathListEsi.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray{}
        pathListEsi.MacUpdateNextHopArray = append(pathListEsi.MacUpdateNextHopArray, child)
        return &pathListEsi.MacUpdateNextHopArray[len(pathListEsi.MacUpdateNextHopArray)-1]
    }
    return nil
}

func (pathListEsi *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ethernet-segment-id"] = &pathListEsi.EthernetSegmentId
    for i := range pathListEsi.MacUpdateNextHopArray {
        children[pathListEsi.MacUpdateNextHopArray[i].GetSegmentPath()] = &pathListEsi.MacUpdateNextHopArray[i]
    }
    return children
}

func (pathListEsi *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["resolved"] = pathListEsi.Resolved
    return leafs
}

func (pathListEsi *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi) GetBundleName() string { return "cisco_ios_xr" }

func (pathListEsi *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi) GetYangName() string { return "path-list-esi" }

func (pathListEsi *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pathListEsi *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pathListEsi *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pathListEsi *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi) SetParent(parent types.Entity) { pathListEsi.parent = parent }

func (pathListEsi *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi) GetParent() types.Entity { return pathListEsi.parent }

func (pathListEsi *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi) GetParentYangName() string { return "path-list-info" }

// L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_EthernetSegmentId
// Ethernet Segment Identifier
type L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_EthernetSegmentId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // LACP System Priority. The type is interface{} with range: 0..65535.
    SystemPriority interface{}

    // LACP System Id. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    SystemId interface{}

    // LACP Port Key. The type is interface{} with range: 0..65535.
    PortKey interface{}
}

func (ethernetSegmentId *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_EthernetSegmentId) GetFilter() yfilter.YFilter { return ethernetSegmentId.YFilter }

func (ethernetSegmentId *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_EthernetSegmentId) SetFilter(yf yfilter.YFilter) { ethernetSegmentId.YFilter = yf }

func (ethernetSegmentId *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_EthernetSegmentId) GetGoName(yname string) string {
    if yname == "system-priority" { return "SystemPriority" }
    if yname == "system-id" { return "SystemId" }
    if yname == "port-key" { return "PortKey" }
    return ""
}

func (ethernetSegmentId *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_EthernetSegmentId) GetSegmentPath() string {
    return "ethernet-segment-id"
}

func (ethernetSegmentId *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_EthernetSegmentId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ethernetSegmentId *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_EthernetSegmentId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ethernetSegmentId *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_EthernetSegmentId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["system-priority"] = ethernetSegmentId.SystemPriority
    leafs["system-id"] = ethernetSegmentId.SystemId
    leafs["port-key"] = ethernetSegmentId.PortKey
    return leafs
}

func (ethernetSegmentId *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_EthernetSegmentId) GetBundleName() string { return "cisco_ios_xr" }

func (ethernetSegmentId *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_EthernetSegmentId) GetYangName() string { return "ethernet-segment-id" }

func (ethernetSegmentId *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_EthernetSegmentId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ethernetSegmentId *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_EthernetSegmentId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ethernetSegmentId *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_EthernetSegmentId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ethernetSegmentId *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_EthernetSegmentId) SetParent(parent types.Entity) { ethernetSegmentId.parent = parent }

func (ethernetSegmentId *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_EthernetSegmentId) GetParent() types.Entity { return ethernetSegmentId.parent }

func (ethernetSegmentId *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_EthernetSegmentId) GetParentYangName() string { return "path-list-esi" }

// L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray
// Array of Next Hops from MAC Update
type L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Next-hop TOPOLOGY ID. The type is interface{} with range: 0..4294967295.
    TopologyId interface{}

    // Next-hop flags. The type is interface{} with range: 0..65535.
    Flags interface{}

    // Next hop.
    NextHop L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop
}

func (macUpdateNextHopArray *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray) GetFilter() yfilter.YFilter { return macUpdateNextHopArray.YFilter }

func (macUpdateNextHopArray *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray) SetFilter(yf yfilter.YFilter) { macUpdateNextHopArray.YFilter = yf }

func (macUpdateNextHopArray *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray) GetGoName(yname string) string {
    if yname == "topology-id" { return "TopologyId" }
    if yname == "flags" { return "Flags" }
    if yname == "next-hop" { return "NextHop" }
    return ""
}

func (macUpdateNextHopArray *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray) GetSegmentPath() string {
    return "mac-update-next-hop-array"
}

func (macUpdateNextHopArray *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "next-hop" {
        return &macUpdateNextHopArray.NextHop
    }
    return nil
}

func (macUpdateNextHopArray *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["next-hop"] = &macUpdateNextHopArray.NextHop
    return children
}

func (macUpdateNextHopArray *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["topology-id"] = macUpdateNextHopArray.TopologyId
    leafs["flags"] = macUpdateNextHopArray.Flags
    return leafs
}

func (macUpdateNextHopArray *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray) GetBundleName() string { return "cisco_ios_xr" }

func (macUpdateNextHopArray *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray) GetYangName() string { return "mac-update-next-hop-array" }

func (macUpdateNextHopArray *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (macUpdateNextHopArray *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (macUpdateNextHopArray *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (macUpdateNextHopArray *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray) SetParent(parent types.Entity) { macUpdateNextHopArray.parent = parent }

func (macUpdateNextHopArray *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray) GetParent() types.Entity { return macUpdateNextHopArray.parent }

func (macUpdateNextHopArray *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray) GetParentYangName() string { return "path-list-esi" }

// L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop
// Next hop
type L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Type. The type is L2ribNextHop.
    Type interface{}

    // IPV4 address Next Hop. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPV6 address Next Hop. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}

    // MAC address Next Hop. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    Mac interface{}

    // Intefrace Handle Next Hop. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceHandle interface{}

    // XID Next Hop. The type is interface{} with range: 0..4294967295.
    Xid interface{}

    // Labeled Next Hop.
    Labeled L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled
}

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop) GetFilter() yfilter.YFilter { return nextHop.YFilter }

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop) SetFilter(yf yfilter.YFilter) { nextHop.YFilter = yf }

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "ipv4" { return "Ipv4" }
    if yname == "ipv6" { return "Ipv6" }
    if yname == "mac" { return "Mac" }
    if yname == "interface-handle" { return "InterfaceHandle" }
    if yname == "xid" { return "Xid" }
    if yname == "labeled" { return "Labeled" }
    return ""
}

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop) GetSegmentPath() string {
    return "next-hop"
}

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "labeled" {
        return &nextHop.Labeled
    }
    return nil
}

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["labeled"] = &nextHop.Labeled
    return children
}

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = nextHop.Type
    leafs["ipv4"] = nextHop.Ipv4
    leafs["ipv6"] = nextHop.Ipv6
    leafs["mac"] = nextHop.Mac
    leafs["interface-handle"] = nextHop.InterfaceHandle
    leafs["xid"] = nextHop.Xid
    return leafs
}

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop) GetBundleName() string { return "cisco_ios_xr" }

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop) GetYangName() string { return "next-hop" }

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop) SetParent(parent types.Entity) { nextHop.parent = parent }

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop) GetParent() types.Entity { return nextHop.parent }

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop) GetParentYangName() string { return "mac-update-next-hop-array" }

// L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled
// Labeled Next Hop
type L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // L2RIB Address Family. The type is L2ribAfi.
    AddressFamily interface{}

    // IP Address (V6 Format). The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IpAddress interface{}

    // Label. The type is interface{} with range: 0..4294967295.
    Label interface{}

    // Internal Label. The type is bool.
    Internal interface{}
}

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled) GetFilter() yfilter.YFilter { return labeled.YFilter }

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled) SetFilter(yf yfilter.YFilter) { labeled.YFilter = yf }

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled) GetGoName(yname string) string {
    if yname == "address-family" { return "AddressFamily" }
    if yname == "ip-address" { return "IpAddress" }
    if yname == "label" { return "Label" }
    if yname == "internal" { return "Internal" }
    return ""
}

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled) GetSegmentPath() string {
    return "labeled"
}

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address-family"] = labeled.AddressFamily
    leafs["ip-address"] = labeled.IpAddress
    leafs["label"] = labeled.Label
    leafs["internal"] = labeled.Internal
    return leafs
}

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled) GetBundleName() string { return "cisco_ios_xr" }

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled) GetYangName() string { return "labeled" }

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled) SetParent(parent types.Entity) { labeled.parent = parent }

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled) GetParent() types.Entity { return labeled.parent }

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled) GetParentYangName() string { return "next-hop" }

// L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListMac
// MAC Path List
type L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListMac struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MAC Address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddress interface{}
}

func (pathListMac *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListMac) GetFilter() yfilter.YFilter { return pathListMac.YFilter }

func (pathListMac *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListMac) SetFilter(yf yfilter.YFilter) { pathListMac.YFilter = yf }

func (pathListMac *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListMac) GetGoName(yname string) string {
    if yname == "mac-address" { return "MacAddress" }
    return ""
}

func (pathListMac *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListMac) GetSegmentPath() string {
    return "path-list-mac"
}

func (pathListMac *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListMac) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pathListMac *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListMac) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pathListMac *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListMac) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mac-address"] = pathListMac.MacAddress
    return leafs
}

func (pathListMac *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListMac) GetBundleName() string { return "cisco_ios_xr" }

func (pathListMac *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListMac) GetYangName() string { return "path-list-mac" }

func (pathListMac *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListMac) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pathListMac *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListMac) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pathListMac *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListMac) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pathListMac *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListMac) SetParent(parent types.Entity) { pathListMac.parent = parent }

func (pathListMac *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListMac) GetParent() types.Entity { return pathListMac.parent }

func (pathListMac *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListMac) GetParentYangName() string { return "path-list-info" }

// L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_NextHopArray
// Array of Next Hops for MAC Path List
type L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_NextHopArray struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Next-hop TOPOLOGY ID. The type is interface{} with range: 0..4294967295.
    TopologyId interface{}

    // Next-hop flags. The type is interface{} with range: 0..65535.
    Flags interface{}

    // Next hop.
    NextHop L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_NextHopArray_NextHop
}

func (nextHopArray *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_NextHopArray) GetFilter() yfilter.YFilter { return nextHopArray.YFilter }

func (nextHopArray *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_NextHopArray) SetFilter(yf yfilter.YFilter) { nextHopArray.YFilter = yf }

func (nextHopArray *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_NextHopArray) GetGoName(yname string) string {
    if yname == "topology-id" { return "TopologyId" }
    if yname == "flags" { return "Flags" }
    if yname == "next-hop" { return "NextHop" }
    return ""
}

func (nextHopArray *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_NextHopArray) GetSegmentPath() string {
    return "next-hop-array"
}

func (nextHopArray *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_NextHopArray) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "next-hop" {
        return &nextHopArray.NextHop
    }
    return nil
}

func (nextHopArray *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_NextHopArray) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["next-hop"] = &nextHopArray.NextHop
    return children
}

func (nextHopArray *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_NextHopArray) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["topology-id"] = nextHopArray.TopologyId
    leafs["flags"] = nextHopArray.Flags
    return leafs
}

func (nextHopArray *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_NextHopArray) GetBundleName() string { return "cisco_ios_xr" }

func (nextHopArray *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_NextHopArray) GetYangName() string { return "next-hop-array" }

func (nextHopArray *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_NextHopArray) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nextHopArray *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_NextHopArray) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nextHopArray *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_NextHopArray) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nextHopArray *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_NextHopArray) SetParent(parent types.Entity) { nextHopArray.parent = parent }

func (nextHopArray *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_NextHopArray) GetParent() types.Entity { return nextHopArray.parent }

func (nextHopArray *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_NextHopArray) GetParentYangName() string { return "path-list" }

// L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_NextHopArray_NextHop
// Next hop
type L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_NextHopArray_NextHop struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Type. The type is L2ribNextHop.
    Type interface{}

    // IPV4 address Next Hop. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPV6 address Next Hop. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}

    // MAC address Next Hop. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    Mac interface{}

    // Intefrace Handle Next Hop. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceHandle interface{}

    // XID Next Hop. The type is interface{} with range: 0..4294967295.
    Xid interface{}

    // Labeled Next Hop.
    Labeled L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_NextHopArray_NextHop_Labeled
}

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_NextHopArray_NextHop) GetFilter() yfilter.YFilter { return nextHop.YFilter }

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_NextHopArray_NextHop) SetFilter(yf yfilter.YFilter) { nextHop.YFilter = yf }

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_NextHopArray_NextHop) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "ipv4" { return "Ipv4" }
    if yname == "ipv6" { return "Ipv6" }
    if yname == "mac" { return "Mac" }
    if yname == "interface-handle" { return "InterfaceHandle" }
    if yname == "xid" { return "Xid" }
    if yname == "labeled" { return "Labeled" }
    return ""
}

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_NextHopArray_NextHop) GetSegmentPath() string {
    return "next-hop"
}

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_NextHopArray_NextHop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "labeled" {
        return &nextHop.Labeled
    }
    return nil
}

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_NextHopArray_NextHop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["labeled"] = &nextHop.Labeled
    return children
}

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_NextHopArray_NextHop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = nextHop.Type
    leafs["ipv4"] = nextHop.Ipv4
    leafs["ipv6"] = nextHop.Ipv6
    leafs["mac"] = nextHop.Mac
    leafs["interface-handle"] = nextHop.InterfaceHandle
    leafs["xid"] = nextHop.Xid
    return leafs
}

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_NextHopArray_NextHop) GetBundleName() string { return "cisco_ios_xr" }

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_NextHopArray_NextHop) GetYangName() string { return "next-hop" }

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_NextHopArray_NextHop) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_NextHopArray_NextHop) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_NextHopArray_NextHop) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_NextHopArray_NextHop) SetParent(parent types.Entity) { nextHop.parent = parent }

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_NextHopArray_NextHop) GetParent() types.Entity { return nextHop.parent }

func (nextHop *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_NextHopArray_NextHop) GetParentYangName() string { return "next-hop-array" }

// L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_NextHopArray_NextHop_Labeled
// Labeled Next Hop
type L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_NextHopArray_NextHop_Labeled struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // L2RIB Address Family. The type is L2ribAfi.
    AddressFamily interface{}

    // IP Address (V6 Format). The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IpAddress interface{}

    // Label. The type is interface{} with range: 0..4294967295.
    Label interface{}

    // Internal Label. The type is bool.
    Internal interface{}
}

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_NextHopArray_NextHop_Labeled) GetFilter() yfilter.YFilter { return labeled.YFilter }

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_NextHopArray_NextHop_Labeled) SetFilter(yf yfilter.YFilter) { labeled.YFilter = yf }

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_NextHopArray_NextHop_Labeled) GetGoName(yname string) string {
    if yname == "address-family" { return "AddressFamily" }
    if yname == "ip-address" { return "IpAddress" }
    if yname == "label" { return "Label" }
    if yname == "internal" { return "Internal" }
    return ""
}

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_NextHopArray_NextHop_Labeled) GetSegmentPath() string {
    return "labeled"
}

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_NextHopArray_NextHop_Labeled) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_NextHopArray_NextHop_Labeled) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_NextHopArray_NextHop_Labeled) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address-family"] = labeled.AddressFamily
    leafs["ip-address"] = labeled.IpAddress
    leafs["label"] = labeled.Label
    leafs["internal"] = labeled.Internal
    return leafs
}

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_NextHopArray_NextHop_Labeled) GetBundleName() string { return "cisco_ios_xr" }

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_NextHopArray_NextHop_Labeled) GetYangName() string { return "labeled" }

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_NextHopArray_NextHop_Labeled) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_NextHopArray_NextHop_Labeled) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_NextHopArray_NextHop_Labeled) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_NextHopArray_NextHop_Labeled) SetParent(parent types.Entity) { labeled.parent = parent }

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_NextHopArray_NextHop_Labeled) GetParent() types.Entity { return labeled.parent }

func (labeled *L2Rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_NextHopArray_NextHop_Labeled) GetParentYangName() string { return "next-hop" }

// L2Rib_EviChildTables_MacDetails_MacDetail_RtTlv
// Mac Route Opaque Data TLV
type L2Rib_EviChildTables_MacDetails_MacDetail_RtTlv struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TLV Type. The type is interface{} with range: 0..65535.
    TlvType interface{}

    // TLV Length. The type is interface{} with range: 0..65535.
    TlvLen interface{}

    // TLV Value. The type is slice of interface{} with range: 0..255.
    TlvVal []interface{}
}

func (rtTlv *L2Rib_EviChildTables_MacDetails_MacDetail_RtTlv) GetFilter() yfilter.YFilter { return rtTlv.YFilter }

func (rtTlv *L2Rib_EviChildTables_MacDetails_MacDetail_RtTlv) SetFilter(yf yfilter.YFilter) { rtTlv.YFilter = yf }

func (rtTlv *L2Rib_EviChildTables_MacDetails_MacDetail_RtTlv) GetGoName(yname string) string {
    if yname == "tlv-type" { return "TlvType" }
    if yname == "tlv-len" { return "TlvLen" }
    if yname == "tlv-val" { return "TlvVal" }
    return ""
}

func (rtTlv *L2Rib_EviChildTables_MacDetails_MacDetail_RtTlv) GetSegmentPath() string {
    return "rt-tlv"
}

func (rtTlv *L2Rib_EviChildTables_MacDetails_MacDetail_RtTlv) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rtTlv *L2Rib_EviChildTables_MacDetails_MacDetail_RtTlv) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rtTlv *L2Rib_EviChildTables_MacDetails_MacDetail_RtTlv) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tlv-type"] = rtTlv.TlvType
    leafs["tlv-len"] = rtTlv.TlvLen
    leafs["tlv-val"] = rtTlv.TlvVal
    return leafs
}

func (rtTlv *L2Rib_EviChildTables_MacDetails_MacDetail_RtTlv) GetBundleName() string { return "cisco_ios_xr" }

func (rtTlv *L2Rib_EviChildTables_MacDetails_MacDetail_RtTlv) GetYangName() string { return "rt-tlv" }

func (rtTlv *L2Rib_EviChildTables_MacDetails_MacDetail_RtTlv) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rtTlv *L2Rib_EviChildTables_MacDetails_MacDetail_RtTlv) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rtTlv *L2Rib_EviChildTables_MacDetails_MacDetail_RtTlv) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rtTlv *L2Rib_EviChildTables_MacDetails_MacDetail_RtTlv) SetParent(parent types.Entity) { rtTlv.parent = parent }

func (rtTlv *L2Rib_EviChildTables_MacDetails_MacDetail_RtTlv) GetParent() types.Entity { return rtTlv.parent }

func (rtTlv *L2Rib_EviChildTables_MacDetails_MacDetail_RtTlv) GetParentYangName() string { return "mac-detail" }

// L2Rib_Evis
// L2RIB EVPN EVI Table
type L2Rib_Evis struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // L2RIB EVPN EVI Entry. The type is slice of L2Rib_Evis_Evi.
    Evi []L2Rib_Evis_Evi
}

func (evis *L2Rib_Evis) GetFilter() yfilter.YFilter { return evis.YFilter }

func (evis *L2Rib_Evis) SetFilter(yf yfilter.YFilter) { evis.YFilter = yf }

func (evis *L2Rib_Evis) GetGoName(yname string) string {
    if yname == "evi" { return "Evi" }
    return ""
}

func (evis *L2Rib_Evis) GetSegmentPath() string {
    return "evis"
}

func (evis *L2Rib_Evis) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "evi" {
        for _, c := range evis.Evi {
            if evis.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := L2Rib_Evis_Evi{}
        evis.Evi = append(evis.Evi, child)
        return &evis.Evi[len(evis.Evi)-1]
    }
    return nil
}

func (evis *L2Rib_Evis) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range evis.Evi {
        children[evis.Evi[i].GetSegmentPath()] = &evis.Evi[i]
    }
    return children
}

func (evis *L2Rib_Evis) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (evis *L2Rib_Evis) GetBundleName() string { return "cisco_ios_xr" }

func (evis *L2Rib_Evis) GetYangName() string { return "evis" }

func (evis *L2Rib_Evis) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (evis *L2Rib_Evis) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (evis *L2Rib_Evis) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (evis *L2Rib_Evis) SetParent(parent types.Entity) { evis.parent = parent }

func (evis *L2Rib_Evis) GetParent() types.Entity { return evis.parent }

func (evis *L2Rib_Evis) GetParentYangName() string { return "l2rib" }

// L2Rib_Evis_Evi
// L2RIB EVPN EVI Entry
type L2Rib_Evis_Evi struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. EVI ID. The type is interface{} with range:
    // -2147483648..2147483647.
    Evi interface{}

    // Topology ID. The type is interface{} with range: 0..4294967295.
    TopologyId interface{}

    // Topology Name. The type is string.
    TopologyName interface{}

    // Topology Type. The type is interface{} with range: 0..4294967295.
    TopologyType interface{}
}

func (evi *L2Rib_Evis_Evi) GetFilter() yfilter.YFilter { return evi.YFilter }

func (evi *L2Rib_Evis_Evi) SetFilter(yf yfilter.YFilter) { evi.YFilter = yf }

func (evi *L2Rib_Evis_Evi) GetGoName(yname string) string {
    if yname == "evi" { return "Evi" }
    if yname == "topology-id" { return "TopologyId" }
    if yname == "topology-name" { return "TopologyName" }
    if yname == "topology-type" { return "TopologyType" }
    return ""
}

func (evi *L2Rib_Evis_Evi) GetSegmentPath() string {
    return "evi" + "[evi='" + fmt.Sprintf("%v", evi.Evi) + "']"
}

func (evi *L2Rib_Evis_Evi) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (evi *L2Rib_Evis_Evi) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (evi *L2Rib_Evis_Evi) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["evi"] = evi.Evi
    leafs["topology-id"] = evi.TopologyId
    leafs["topology-name"] = evi.TopologyName
    leafs["topology-type"] = evi.TopologyType
    return leafs
}

func (evi *L2Rib_Evis_Evi) GetBundleName() string { return "cisco_ios_xr" }

func (evi *L2Rib_Evis_Evi) GetYangName() string { return "evi" }

func (evi *L2Rib_Evis_Evi) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (evi *L2Rib_Evis_Evi) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (evi *L2Rib_Evis_Evi) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (evi *L2Rib_Evis_Evi) SetParent(parent types.Entity) { evi.parent = parent }

func (evi *L2Rib_Evis_Evi) GetParent() types.Entity { return evi.parent }

func (evi *L2Rib_Evis_Evi) GetParentYangName() string { return "evis" }

