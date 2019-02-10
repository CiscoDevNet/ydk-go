// This module contains a collection of YANG definitions
// for Cisco IOS-XR l2rib package operational data.
// 
// This module contains definitions
// for the following management objects:
//   l2rib: L2RIB operational information 
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-l2rib-oper l2rib}", reflect.TypeOf(L2rib{}))
    ydk.RegisterEntity("Cisco-IOS-XR-l2rib-oper:l2rib", reflect.TypeOf(L2rib{}))
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

    // Local Proxy
    L2ribBagProducerId_l2rib_bag_prod_prod_local_proxy L2ribBagProducerId = "l2rib-bag-prod-prod-local-proxy"

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

// L2rib
// L2RIB operational information 
type L2rib struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2RIB detailed producer table.
    ProducersDetails L2rib_ProducersDetails

    // L2RIB EVPN Summary.
    Summary L2rib_Summary

    // L2RIB producer table.
    Producers L2rib_Producers

    // L2RIB client table.
    Clients L2rib_Clients

    // L2RIB EVPN EVI Detail Table.
    EvisXr L2rib_EvisXr

    // L2RIB detailed client table.
    ClientsDetails L2rib_ClientsDetails

    // Container for all EVI Child Tables.
    EviChildTables L2rib_EviChildTables

    // L2RIB EVPN EVI Table.
    Evis L2rib_Evis
}

func (l2rib *L2rib) GetEntityData() *types.CommonEntityData {
    l2rib.EntityData.YFilter = l2rib.YFilter
    l2rib.EntityData.YangName = "l2rib"
    l2rib.EntityData.BundleName = "cisco_ios_xr"
    l2rib.EntityData.ParentYangName = "Cisco-IOS-XR-l2rib-oper"
    l2rib.EntityData.SegmentPath = "Cisco-IOS-XR-l2rib-oper:l2rib"
    l2rib.EntityData.AbsolutePath = l2rib.EntityData.SegmentPath
    l2rib.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    l2rib.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    l2rib.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    l2rib.EntityData.Children = types.NewOrderedMap()
    l2rib.EntityData.Children.Append("producers-details", types.YChild{"ProducersDetails", &l2rib.ProducersDetails})
    l2rib.EntityData.Children.Append("summary", types.YChild{"Summary", &l2rib.Summary})
    l2rib.EntityData.Children.Append("producers", types.YChild{"Producers", &l2rib.Producers})
    l2rib.EntityData.Children.Append("clients", types.YChild{"Clients", &l2rib.Clients})
    l2rib.EntityData.Children.Append("evis-xr", types.YChild{"EvisXr", &l2rib.EvisXr})
    l2rib.EntityData.Children.Append("clients-details", types.YChild{"ClientsDetails", &l2rib.ClientsDetails})
    l2rib.EntityData.Children.Append("evi-child-tables", types.YChild{"EviChildTables", &l2rib.EviChildTables})
    l2rib.EntityData.Children.Append("evis", types.YChild{"Evis", &l2rib.Evis})
    l2rib.EntityData.Leafs = types.NewOrderedMap()

    l2rib.EntityData.YListKeys = []string {}

    return &(l2rib.EntityData)
}

// L2rib_ProducersDetails
// L2RIB detailed producer table
type L2rib_ProducersDetails struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2RIB producers detail information. The type is slice of
    // L2rib_ProducersDetails_ProducersDetail.
    ProducersDetail []*L2rib_ProducersDetails_ProducersDetail
}

func (producersDetails *L2rib_ProducersDetails) GetEntityData() *types.CommonEntityData {
    producersDetails.EntityData.YFilter = producersDetails.YFilter
    producersDetails.EntityData.YangName = "producers-details"
    producersDetails.EntityData.BundleName = "cisco_ios_xr"
    producersDetails.EntityData.ParentYangName = "l2rib"
    producersDetails.EntityData.SegmentPath = "producers-details"
    producersDetails.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/" + producersDetails.EntityData.SegmentPath
    producersDetails.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    producersDetails.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    producersDetails.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    producersDetails.EntityData.Children = types.NewOrderedMap()
    producersDetails.EntityData.Children.Append("producers-detail", types.YChild{"ProducersDetail", nil})
    for i := range producersDetails.ProducersDetail {
        types.SetYListKey(producersDetails.ProducersDetail[i], i)
        producersDetails.EntityData.Children.Append(types.GetSegmentPath(producersDetails.ProducersDetail[i]), types.YChild{"ProducersDetail", producersDetails.ProducersDetail[i]})
    }
    producersDetails.EntityData.Leafs = types.NewOrderedMap()

    producersDetails.EntityData.YListKeys = []string {}

    return &(producersDetails.EntityData)
}

// L2rib_ProducersDetails_ProducersDetail
// L2RIB producers detail information
type L2rib_ProducersDetails_ProducersDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Object ID. The type is interface{} with range: 0..4294967295.
    ObjectId interface{}

    // Product ID. The type is interface{} with range: 0..4294967295.
    ProductId interface{}

    // Last Update Timestamp. The type is interface{} with range:
    // 0..18446744073709551615.
    LastUpdateTimestamp interface{}

    // Non-detail Producer Bag.
    Producer L2rib_ProducersDetails_ProducersDetail_Producer

    // Producer Statistics.
    Statistics L2rib_ProducersDetails_ProducersDetail_Statistics
}

func (producersDetail *L2rib_ProducersDetails_ProducersDetail) GetEntityData() *types.CommonEntityData {
    producersDetail.EntityData.YFilter = producersDetail.YFilter
    producersDetail.EntityData.YangName = "producers-detail"
    producersDetail.EntityData.BundleName = "cisco_ios_xr"
    producersDetail.EntityData.ParentYangName = "producers-details"
    producersDetail.EntityData.SegmentPath = "producers-detail" + types.AddNoKeyToken(producersDetail)
    producersDetail.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/producers-details/" + producersDetail.EntityData.SegmentPath
    producersDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    producersDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    producersDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    producersDetail.EntityData.Children = types.NewOrderedMap()
    producersDetail.EntityData.Children.Append("producer", types.YChild{"Producer", &producersDetail.Producer})
    producersDetail.EntityData.Children.Append("statistics", types.YChild{"Statistics", &producersDetail.Statistics})
    producersDetail.EntityData.Leafs = types.NewOrderedMap()
    producersDetail.EntityData.Leafs.Append("object-id", types.YLeaf{"ObjectId", producersDetail.ObjectId})
    producersDetail.EntityData.Leafs.Append("product-id", types.YLeaf{"ProductId", producersDetail.ProductId})
    producersDetail.EntityData.Leafs.Append("last-update-timestamp", types.YLeaf{"LastUpdateTimestamp", producersDetail.LastUpdateTimestamp})

    producersDetail.EntityData.YListKeys = []string {}

    return &(producersDetail.EntityData)
}

// L2rib_ProducersDetails_ProducersDetail_Producer
// Non-detail Producer Bag
type L2rib_ProducersDetails_ProducersDetail_Producer struct {
    EntityData types.CommonEntityData
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

func (producer *L2rib_ProducersDetails_ProducersDetail_Producer) GetEntityData() *types.CommonEntityData {
    producer.EntityData.YFilter = producer.YFilter
    producer.EntityData.YangName = "producer"
    producer.EntityData.BundleName = "cisco_ios_xr"
    producer.EntityData.ParentYangName = "producers-detail"
    producer.EntityData.SegmentPath = "producer"
    producer.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/producers-details/producers-detail/" + producer.EntityData.SegmentPath
    producer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    producer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    producer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    producer.EntityData.Children = types.NewOrderedMap()
    producer.EntityData.Leafs = types.NewOrderedMap()
    producer.EntityData.Leafs.Append("client-id", types.YLeaf{"ClientId", producer.ClientId})
    producer.EntityData.Leafs.Append("object-type", types.YLeaf{"ObjectType", producer.ObjectType})
    producer.EntityData.Leafs.Append("producer-id", types.YLeaf{"ProducerId", producer.ProducerId})
    producer.EntityData.Leafs.Append("producer-name", types.YLeaf{"ProducerName", producer.ProducerName})
    producer.EntityData.Leafs.Append("admin-distance", types.YLeaf{"AdminDistance", producer.AdminDistance})
    producer.EntityData.Leafs.Append("purge-time", types.YLeaf{"PurgeTime", producer.PurgeTime})
    producer.EntityData.Leafs.Append("state", types.YLeaf{"State", producer.State})

    producer.EntityData.YListKeys = []string {}

    return &(producer.EntityData)
}

// L2rib_ProducersDetails_ProducersDetail_Statistics
// Producer Statistics
type L2rib_ProducersDetails_ProducersDetail_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Producer ID. The type is L2ribBagProducerId.
    ProducerId interface{}

    // Producer Name. The type is string.
    ProducerName interface{}

    // Statistics.
    Statistics L2rib_ProducersDetails_ProducersDetail_Statistics_Statistics
}

func (statistics *L2rib_ProducersDetails_ProducersDetail_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "producers-detail"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/producers-details/producers-detail/" + statistics.EntityData.SegmentPath
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = types.NewOrderedMap()
    statistics.EntityData.Children.Append("statistics", types.YChild{"Statistics", &statistics.Statistics})
    statistics.EntityData.Leafs = types.NewOrderedMap()
    statistics.EntityData.Leafs.Append("producer-id", types.YLeaf{"ProducerId", statistics.ProducerId})
    statistics.EntityData.Leafs.Append("producer-name", types.YLeaf{"ProducerName", statistics.ProducerName})

    statistics.EntityData.YListKeys = []string {}

    return &(statistics.EntityData)
}

// L2rib_ProducersDetails_ProducersDetail_Statistics_Statistics
// Statistics
type L2rib_ProducersDetails_ProducersDetail_Statistics_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Memory Size. The type is interface{} with range: 0..4294967295.
    MemorySize interface{}

    // Number of Objects. The type is interface{} with range: 0..4294967295.
    ObjectCount interface{}

    // End of Interval Timestamp. The type is interface{} with range:
    // 0..18446744073709551615.
    EndofIntervalTs interface{}

    // Extended Counters. The type is slice of
    // L2rib_ProducersDetails_ProducersDetail_Statistics_Statistics_ExtendedCounter.
    ExtendedCounter []*L2rib_ProducersDetails_ProducersDetail_Statistics_Statistics_ExtendedCounter
}

func (statistics *L2rib_ProducersDetails_ProducersDetail_Statistics_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "statistics"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/producers-details/producers-detail/statistics/" + statistics.EntityData.SegmentPath
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = types.NewOrderedMap()
    statistics.EntityData.Children.Append("extended-counter", types.YChild{"ExtendedCounter", nil})
    for i := range statistics.ExtendedCounter {
        types.SetYListKey(statistics.ExtendedCounter[i], i)
        statistics.EntityData.Children.Append(types.GetSegmentPath(statistics.ExtendedCounter[i]), types.YChild{"ExtendedCounter", statistics.ExtendedCounter[i]})
    }
    statistics.EntityData.Leafs = types.NewOrderedMap()
    statistics.EntityData.Leafs.Append("memory-size", types.YLeaf{"MemorySize", statistics.MemorySize})
    statistics.EntityData.Leafs.Append("object-count", types.YLeaf{"ObjectCount", statistics.ObjectCount})
    statistics.EntityData.Leafs.Append("endof-interval-ts", types.YLeaf{"EndofIntervalTs", statistics.EndofIntervalTs})

    statistics.EntityData.YListKeys = []string {}

    return &(statistics.EntityData)
}

// L2rib_ProducersDetails_ProducersDetail_Statistics_Statistics_ExtendedCounter
// Extended Counters
type L2rib_ProducersDetails_ProducersDetail_Statistics_Statistics_ExtendedCounter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // CounterType. The type is interface{} with range: 0..255.
    CounterType interface{}

    // CounterName. The type is string.
    CounterName interface{}

    // Real-clock timestamp in msec of first event. The type is interface{} with
    // range: 0..18446744073709551615.
    L2rbFirstEventTs interface{}

    // Real-clock timestamp in msec of last event. The type is interface{} with
    // range: 0..18446744073709551615.
    L2rbLastEventTs interface{}

    // number of events in interval. The type is interface{} with range:
    // 0..4294967295.
    L2rbIntervalEventCount interface{}

    // total number of events. The type is interface{} with range: 0..4294967295.
    L2rbTotalEventCount interface{}
}

func (extendedCounter *L2rib_ProducersDetails_ProducersDetail_Statistics_Statistics_ExtendedCounter) GetEntityData() *types.CommonEntityData {
    extendedCounter.EntityData.YFilter = extendedCounter.YFilter
    extendedCounter.EntityData.YangName = "extended-counter"
    extendedCounter.EntityData.BundleName = "cisco_ios_xr"
    extendedCounter.EntityData.ParentYangName = "statistics"
    extendedCounter.EntityData.SegmentPath = "extended-counter" + types.AddNoKeyToken(extendedCounter)
    extendedCounter.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/producers-details/producers-detail/statistics/statistics/" + extendedCounter.EntityData.SegmentPath
    extendedCounter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    extendedCounter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    extendedCounter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    extendedCounter.EntityData.Children = types.NewOrderedMap()
    extendedCounter.EntityData.Leafs = types.NewOrderedMap()
    extendedCounter.EntityData.Leafs.Append("counter-type", types.YLeaf{"CounterType", extendedCounter.CounterType})
    extendedCounter.EntityData.Leafs.Append("counter-name", types.YLeaf{"CounterName", extendedCounter.CounterName})
    extendedCounter.EntityData.Leafs.Append("l2rb-first-event-ts", types.YLeaf{"L2rbFirstEventTs", extendedCounter.L2rbFirstEventTs})
    extendedCounter.EntityData.Leafs.Append("l2rb-last-event-ts", types.YLeaf{"L2rbLastEventTs", extendedCounter.L2rbLastEventTs})
    extendedCounter.EntityData.Leafs.Append("l2rb-interval-event-count", types.YLeaf{"L2rbIntervalEventCount", extendedCounter.L2rbIntervalEventCount})
    extendedCounter.EntityData.Leafs.Append("l2rb-total-event-count", types.YLeaf{"L2rbTotalEventCount", extendedCounter.L2rbTotalEventCount})

    extendedCounter.EntityData.YListKeys = []string {}

    return &(extendedCounter.EntityData)
}

// L2rib_Summary
// L2RIB EVPN Summary
type L2rib_Summary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of Converged Tables. The type is interface{} with range:
    // 0..4294967295.
    ConvergedTablesCount interface{}

    // Total Allocated Memory. The type is interface{} with range: 0..4294967295.
    TotalMemory interface{}

    // Per Object Table summary. The type is slice of L2rib_Summary_TableSummary.
    TableSummary []*L2rib_Summary_TableSummary
}

func (summary *L2rib_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "cisco_ios_xr"
    summary.EntityData.ParentYangName = "l2rib"
    summary.EntityData.SegmentPath = "summary"
    summary.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/" + summary.EntityData.SegmentPath
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = types.NewOrderedMap()
    summary.EntityData.Children.Append("table-summary", types.YChild{"TableSummary", nil})
    for i := range summary.TableSummary {
        types.SetYListKey(summary.TableSummary[i], i)
        summary.EntityData.Children.Append(types.GetSegmentPath(summary.TableSummary[i]), types.YChild{"TableSummary", summary.TableSummary[i]})
    }
    summary.EntityData.Leafs = types.NewOrderedMap()
    summary.EntityData.Leafs.Append("converged-tables-count", types.YLeaf{"ConvergedTablesCount", summary.ConvergedTablesCount})
    summary.EntityData.Leafs.Append("total-memory", types.YLeaf{"TotalMemory", summary.TotalMemory})

    summary.EntityData.YListKeys = []string {}

    return &(summary.EntityData)
}

// L2rib_Summary_TableSummary
// Per Object Table summary
type L2rib_Summary_TableSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Object Type. The type is L2ribBagObj.
    ObjectType interface{}

    // Number of Objects. The type is interface{} with range: 0..4294967295.
    ObjectCount interface{}

    // Allocated Memory. The type is interface{} with range: 0..4294967295.
    TableMemory interface{}

    // Statistics per producer. The type is slice of
    // L2rib_Summary_TableSummary_ProducerStat.
    ProducerStat []*L2rib_Summary_TableSummary_ProducerStat
}

func (tableSummary *L2rib_Summary_TableSummary) GetEntityData() *types.CommonEntityData {
    tableSummary.EntityData.YFilter = tableSummary.YFilter
    tableSummary.EntityData.YangName = "table-summary"
    tableSummary.EntityData.BundleName = "cisco_ios_xr"
    tableSummary.EntityData.ParentYangName = "summary"
    tableSummary.EntityData.SegmentPath = "table-summary" + types.AddNoKeyToken(tableSummary)
    tableSummary.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/summary/" + tableSummary.EntityData.SegmentPath
    tableSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tableSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tableSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tableSummary.EntityData.Children = types.NewOrderedMap()
    tableSummary.EntityData.Children.Append("producer-stat", types.YChild{"ProducerStat", nil})
    for i := range tableSummary.ProducerStat {
        types.SetYListKey(tableSummary.ProducerStat[i], i)
        tableSummary.EntityData.Children.Append(types.GetSegmentPath(tableSummary.ProducerStat[i]), types.YChild{"ProducerStat", tableSummary.ProducerStat[i]})
    }
    tableSummary.EntityData.Leafs = types.NewOrderedMap()
    tableSummary.EntityData.Leafs.Append("object-type", types.YLeaf{"ObjectType", tableSummary.ObjectType})
    tableSummary.EntityData.Leafs.Append("object-count", types.YLeaf{"ObjectCount", tableSummary.ObjectCount})
    tableSummary.EntityData.Leafs.Append("table-memory", types.YLeaf{"TableMemory", tableSummary.TableMemory})

    tableSummary.EntityData.YListKeys = []string {}

    return &(tableSummary.EntityData)
}

// L2rib_Summary_TableSummary_ProducerStat
// Statistics per producer
type L2rib_Summary_TableSummary_ProducerStat struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Producer ID. The type is L2ribBagProducerId.
    ProducerId interface{}

    // Producer Name. The type is string.
    ProducerName interface{}

    // Statistics.
    Statistics L2rib_Summary_TableSummary_ProducerStat_Statistics
}

func (producerStat *L2rib_Summary_TableSummary_ProducerStat) GetEntityData() *types.CommonEntityData {
    producerStat.EntityData.YFilter = producerStat.YFilter
    producerStat.EntityData.YangName = "producer-stat"
    producerStat.EntityData.BundleName = "cisco_ios_xr"
    producerStat.EntityData.ParentYangName = "table-summary"
    producerStat.EntityData.SegmentPath = "producer-stat" + types.AddNoKeyToken(producerStat)
    producerStat.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/summary/table-summary/" + producerStat.EntityData.SegmentPath
    producerStat.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    producerStat.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    producerStat.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    producerStat.EntityData.Children = types.NewOrderedMap()
    producerStat.EntityData.Children.Append("statistics", types.YChild{"Statistics", &producerStat.Statistics})
    producerStat.EntityData.Leafs = types.NewOrderedMap()
    producerStat.EntityData.Leafs.Append("producer-id", types.YLeaf{"ProducerId", producerStat.ProducerId})
    producerStat.EntityData.Leafs.Append("producer-name", types.YLeaf{"ProducerName", producerStat.ProducerName})

    producerStat.EntityData.YListKeys = []string {}

    return &(producerStat.EntityData)
}

// L2rib_Summary_TableSummary_ProducerStat_Statistics
// Statistics
type L2rib_Summary_TableSummary_ProducerStat_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Memory Size. The type is interface{} with range: 0..4294967295.
    MemorySize interface{}

    // Number of Objects. The type is interface{} with range: 0..4294967295.
    ObjectCount interface{}

    // End of Interval Timestamp. The type is interface{} with range:
    // 0..18446744073709551615.
    EndofIntervalTs interface{}

    // Extended Counters. The type is slice of
    // L2rib_Summary_TableSummary_ProducerStat_Statistics_ExtendedCounter.
    ExtendedCounter []*L2rib_Summary_TableSummary_ProducerStat_Statistics_ExtendedCounter
}

func (statistics *L2rib_Summary_TableSummary_ProducerStat_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "producer-stat"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/summary/table-summary/producer-stat/" + statistics.EntityData.SegmentPath
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = types.NewOrderedMap()
    statistics.EntityData.Children.Append("extended-counter", types.YChild{"ExtendedCounter", nil})
    for i := range statistics.ExtendedCounter {
        types.SetYListKey(statistics.ExtendedCounter[i], i)
        statistics.EntityData.Children.Append(types.GetSegmentPath(statistics.ExtendedCounter[i]), types.YChild{"ExtendedCounter", statistics.ExtendedCounter[i]})
    }
    statistics.EntityData.Leafs = types.NewOrderedMap()
    statistics.EntityData.Leafs.Append("memory-size", types.YLeaf{"MemorySize", statistics.MemorySize})
    statistics.EntityData.Leafs.Append("object-count", types.YLeaf{"ObjectCount", statistics.ObjectCount})
    statistics.EntityData.Leafs.Append("endof-interval-ts", types.YLeaf{"EndofIntervalTs", statistics.EndofIntervalTs})

    statistics.EntityData.YListKeys = []string {}

    return &(statistics.EntityData)
}

// L2rib_Summary_TableSummary_ProducerStat_Statistics_ExtendedCounter
// Extended Counters
type L2rib_Summary_TableSummary_ProducerStat_Statistics_ExtendedCounter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // CounterType. The type is interface{} with range: 0..255.
    CounterType interface{}

    // CounterName. The type is string.
    CounterName interface{}

    // Real-clock timestamp in msec of first event. The type is interface{} with
    // range: 0..18446744073709551615.
    L2rbFirstEventTs interface{}

    // Real-clock timestamp in msec of last event. The type is interface{} with
    // range: 0..18446744073709551615.
    L2rbLastEventTs interface{}

    // number of events in interval. The type is interface{} with range:
    // 0..4294967295.
    L2rbIntervalEventCount interface{}

    // total number of events. The type is interface{} with range: 0..4294967295.
    L2rbTotalEventCount interface{}
}

func (extendedCounter *L2rib_Summary_TableSummary_ProducerStat_Statistics_ExtendedCounter) GetEntityData() *types.CommonEntityData {
    extendedCounter.EntityData.YFilter = extendedCounter.YFilter
    extendedCounter.EntityData.YangName = "extended-counter"
    extendedCounter.EntityData.BundleName = "cisco_ios_xr"
    extendedCounter.EntityData.ParentYangName = "statistics"
    extendedCounter.EntityData.SegmentPath = "extended-counter" + types.AddNoKeyToken(extendedCounter)
    extendedCounter.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/summary/table-summary/producer-stat/statistics/" + extendedCounter.EntityData.SegmentPath
    extendedCounter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    extendedCounter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    extendedCounter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    extendedCounter.EntityData.Children = types.NewOrderedMap()
    extendedCounter.EntityData.Leafs = types.NewOrderedMap()
    extendedCounter.EntityData.Leafs.Append("counter-type", types.YLeaf{"CounterType", extendedCounter.CounterType})
    extendedCounter.EntityData.Leafs.Append("counter-name", types.YLeaf{"CounterName", extendedCounter.CounterName})
    extendedCounter.EntityData.Leafs.Append("l2rb-first-event-ts", types.YLeaf{"L2rbFirstEventTs", extendedCounter.L2rbFirstEventTs})
    extendedCounter.EntityData.Leafs.Append("l2rb-last-event-ts", types.YLeaf{"L2rbLastEventTs", extendedCounter.L2rbLastEventTs})
    extendedCounter.EntityData.Leafs.Append("l2rb-interval-event-count", types.YLeaf{"L2rbIntervalEventCount", extendedCounter.L2rbIntervalEventCount})
    extendedCounter.EntityData.Leafs.Append("l2rb-total-event-count", types.YLeaf{"L2rbTotalEventCount", extendedCounter.L2rbTotalEventCount})

    extendedCounter.EntityData.YListKeys = []string {}

    return &(extendedCounter.EntityData)
}

// L2rib_Producers
// L2RIB producer table
type L2rib_Producers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2RIB producers. The type is slice of L2rib_Producers_Producer.
    Producer []*L2rib_Producers_Producer
}

func (producers *L2rib_Producers) GetEntityData() *types.CommonEntityData {
    producers.EntityData.YFilter = producers.YFilter
    producers.EntityData.YangName = "producers"
    producers.EntityData.BundleName = "cisco_ios_xr"
    producers.EntityData.ParentYangName = "l2rib"
    producers.EntityData.SegmentPath = "producers"
    producers.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/" + producers.EntityData.SegmentPath
    producers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    producers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    producers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    producers.EntityData.Children = types.NewOrderedMap()
    producers.EntityData.Children.Append("producer", types.YChild{"Producer", nil})
    for i := range producers.Producer {
        types.SetYListKey(producers.Producer[i], i)
        producers.EntityData.Children.Append(types.GetSegmentPath(producers.Producer[i]), types.YChild{"Producer", producers.Producer[i]})
    }
    producers.EntityData.Leafs = types.NewOrderedMap()

    producers.EntityData.YListKeys = []string {}

    return &(producers.EntityData)
}

// L2rib_Producers_Producer
// L2RIB producers
type L2rib_Producers_Producer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Object ID. The type is interface{} with range: 0..4294967295.
    ObjectId interface{}

    // Product ID. The type is interface{} with range: 0..4294967295.
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

func (producer *L2rib_Producers_Producer) GetEntityData() *types.CommonEntityData {
    producer.EntityData.YFilter = producer.YFilter
    producer.EntityData.YangName = "producer"
    producer.EntityData.BundleName = "cisco_ios_xr"
    producer.EntityData.ParentYangName = "producers"
    producer.EntityData.SegmentPath = "producer" + types.AddNoKeyToken(producer)
    producer.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/producers/" + producer.EntityData.SegmentPath
    producer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    producer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    producer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    producer.EntityData.Children = types.NewOrderedMap()
    producer.EntityData.Leafs = types.NewOrderedMap()
    producer.EntityData.Leafs.Append("object-id", types.YLeaf{"ObjectId", producer.ObjectId})
    producer.EntityData.Leafs.Append("product-id", types.YLeaf{"ProductId", producer.ProductId})
    producer.EntityData.Leafs.Append("client-id", types.YLeaf{"ClientId", producer.ClientId})
    producer.EntityData.Leafs.Append("object-type", types.YLeaf{"ObjectType", producer.ObjectType})
    producer.EntityData.Leafs.Append("producer-id", types.YLeaf{"ProducerId", producer.ProducerId})
    producer.EntityData.Leafs.Append("producer-name", types.YLeaf{"ProducerName", producer.ProducerName})
    producer.EntityData.Leafs.Append("admin-distance", types.YLeaf{"AdminDistance", producer.AdminDistance})
    producer.EntityData.Leafs.Append("purge-time", types.YLeaf{"PurgeTime", producer.PurgeTime})
    producer.EntityData.Leafs.Append("state", types.YLeaf{"State", producer.State})

    producer.EntityData.YListKeys = []string {}

    return &(producer.EntityData)
}

// L2rib_Clients
// L2RIB client table
type L2rib_Clients struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2RIB clients. The type is slice of L2rib_Clients_Client.
    Client []*L2rib_Clients_Client
}

func (clients *L2rib_Clients) GetEntityData() *types.CommonEntityData {
    clients.EntityData.YFilter = clients.YFilter
    clients.EntityData.YangName = "clients"
    clients.EntityData.BundleName = "cisco_ios_xr"
    clients.EntityData.ParentYangName = "l2rib"
    clients.EntityData.SegmentPath = "clients"
    clients.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/" + clients.EntityData.SegmentPath
    clients.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clients.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clients.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clients.EntityData.Children = types.NewOrderedMap()
    clients.EntityData.Children.Append("client", types.YChild{"Client", nil})
    for i := range clients.Client {
        clients.EntityData.Children.Append(types.GetSegmentPath(clients.Client[i]), types.YChild{"Client", clients.Client[i]})
    }
    clients.EntityData.Leafs = types.NewOrderedMap()

    clients.EntityData.YListKeys = []string {}

    return &(clients.EntityData)
}

// L2rib_Clients_Client
// L2RIB clients
type L2rib_Clients_Client struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Client ID. The type is interface{} with range:
    // 0..4294967295.
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

func (client *L2rib_Clients_Client) GetEntityData() *types.CommonEntityData {
    client.EntityData.YFilter = client.YFilter
    client.EntityData.YangName = "client"
    client.EntityData.BundleName = "cisco_ios_xr"
    client.EntityData.ParentYangName = "clients"
    client.EntityData.SegmentPath = "client" + types.AddKeyToken(client.ClientId, "client-id")
    client.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/clients/" + client.EntityData.SegmentPath
    client.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    client.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    client.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    client.EntityData.Children = types.NewOrderedMap()
    client.EntityData.Leafs = types.NewOrderedMap()
    client.EntityData.Leafs.Append("client-id", types.YLeaf{"ClientId", client.ClientId})
    client.EntityData.Leafs.Append("client-id-xr", types.YLeaf{"ClientIdXr", client.ClientIdXr})
    client.EntityData.Leafs.Append("process-id", types.YLeaf{"ProcessId", client.ProcessId})
    client.EntityData.Leafs.Append("node-id", types.YLeaf{"NodeId", client.NodeId})
    client.EntityData.Leafs.Append("proc-name", types.YLeaf{"ProcName", client.ProcName})
    client.EntityData.Leafs.Append("proc-suffix", types.YLeaf{"ProcSuffix", client.ProcSuffix})

    client.EntityData.YListKeys = []string {"ClientId"}

    return &(client.EntityData)
}

// L2rib_EvisXr
// L2RIB EVPN EVI Detail Table
type L2rib_EvisXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2RIB EVPN EVI Entry. The type is slice of L2rib_EvisXr_Evi.
    Evi []*L2rib_EvisXr_Evi
}

func (evisXr *L2rib_EvisXr) GetEntityData() *types.CommonEntityData {
    evisXr.EntityData.YFilter = evisXr.YFilter
    evisXr.EntityData.YangName = "evis-xr"
    evisXr.EntityData.BundleName = "cisco_ios_xr"
    evisXr.EntityData.ParentYangName = "l2rib"
    evisXr.EntityData.SegmentPath = "evis-xr"
    evisXr.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/" + evisXr.EntityData.SegmentPath
    evisXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evisXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evisXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evisXr.EntityData.Children = types.NewOrderedMap()
    evisXr.EntityData.Children.Append("evi", types.YChild{"Evi", nil})
    for i := range evisXr.Evi {
        evisXr.EntityData.Children.Append(types.GetSegmentPath(evisXr.Evi[i]), types.YChild{"Evi", evisXr.Evi[i]})
    }
    evisXr.EntityData.Leafs = types.NewOrderedMap()

    evisXr.EntityData.YListKeys = []string {}

    return &(evisXr.EntityData)
}

// L2rib_EvisXr_Evi
// L2RIB EVPN EVI Entry
type L2rib_EvisXr_Evi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. EVI ID. The type is interface{} with range:
    // 0..4294967295.
    Evi interface{}

    // l2r vni. The type is interface{} with range: 0..4294967295.
    L2rVni interface{}

    // l2r encap type. The type is interface{} with range: 0..65535.
    L2rEncapType interface{}

    // l2r nve iod. The type is interface{} with range: 0..4294967295.
    L2rNveIod interface{}

    // l2r nve ifhandle. The type is interface{} with range: 0..4294967295.
    L2rNveIfhandle interface{}

    // VTEP IP. The type is string.
    VtepIp interface{}

    // l2r topo txid. The type is interface{} with range: 0..4294967295.
    L2rTopoTxid interface{}

    // Topology Flags. The type is interface{} with range: 0..4294967295.
    L2rTopoFlags interface{}

    // Topology.
    Topology L2rib_EvisXr_Evi_Topology
}

func (evi *L2rib_EvisXr_Evi) GetEntityData() *types.CommonEntityData {
    evi.EntityData.YFilter = evi.YFilter
    evi.EntityData.YangName = "evi"
    evi.EntityData.BundleName = "cisco_ios_xr"
    evi.EntityData.ParentYangName = "evis-xr"
    evi.EntityData.SegmentPath = "evi" + types.AddKeyToken(evi.Evi, "evi")
    evi.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evis-xr/" + evi.EntityData.SegmentPath
    evi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evi.EntityData.Children = types.NewOrderedMap()
    evi.EntityData.Children.Append("topology", types.YChild{"Topology", &evi.Topology})
    evi.EntityData.Leafs = types.NewOrderedMap()
    evi.EntityData.Leafs.Append("evi", types.YLeaf{"Evi", evi.Evi})
    evi.EntityData.Leafs.Append("l2r-vni", types.YLeaf{"L2rVni", evi.L2rVni})
    evi.EntityData.Leafs.Append("l2r-encap-type", types.YLeaf{"L2rEncapType", evi.L2rEncapType})
    evi.EntityData.Leafs.Append("l2r-nve-iod", types.YLeaf{"L2rNveIod", evi.L2rNveIod})
    evi.EntityData.Leafs.Append("l2r-nve-ifhandle", types.YLeaf{"L2rNveIfhandle", evi.L2rNveIfhandle})
    evi.EntityData.Leafs.Append("vtep-ip", types.YLeaf{"VtepIp", evi.VtepIp})
    evi.EntityData.Leafs.Append("l2r-topo-txid", types.YLeaf{"L2rTopoTxid", evi.L2rTopoTxid})
    evi.EntityData.Leafs.Append("l2r-topo-flags", types.YLeaf{"L2rTopoFlags", evi.L2rTopoFlags})

    evi.EntityData.YListKeys = []string {"Evi"}

    return &(evi.EntityData)
}

// L2rib_EvisXr_Evi_Topology
// Topology
type L2rib_EvisXr_Evi_Topology struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Topology ID. The type is interface{} with range: 0..4294967295.
    TopologyId interface{}

    // Topology Name. The type is string.
    TopologyName interface{}

    // Topology Type. The type is interface{} with range: 0..4294967295.
    TopologyType interface{}
}

func (topology *L2rib_EvisXr_Evi_Topology) GetEntityData() *types.CommonEntityData {
    topology.EntityData.YFilter = topology.YFilter
    topology.EntityData.YangName = "topology"
    topology.EntityData.BundleName = "cisco_ios_xr"
    topology.EntityData.ParentYangName = "evi"
    topology.EntityData.SegmentPath = "topology"
    topology.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evis-xr/evi/" + topology.EntityData.SegmentPath
    topology.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    topology.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    topology.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    topology.EntityData.Children = types.NewOrderedMap()
    topology.EntityData.Leafs = types.NewOrderedMap()
    topology.EntityData.Leafs.Append("topology-id", types.YLeaf{"TopologyId", topology.TopologyId})
    topology.EntityData.Leafs.Append("topology-name", types.YLeaf{"TopologyName", topology.TopologyName})
    topology.EntityData.Leafs.Append("topology-type", types.YLeaf{"TopologyType", topology.TopologyType})

    topology.EntityData.YListKeys = []string {}

    return &(topology.EntityData)
}

// L2rib_ClientsDetails
// L2RIB detailed client table
type L2rib_ClientsDetails struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2RIB clients detail information. The type is slice of
    // L2rib_ClientsDetails_ClientsDetail.
    ClientsDetail []*L2rib_ClientsDetails_ClientsDetail
}

func (clientsDetails *L2rib_ClientsDetails) GetEntityData() *types.CommonEntityData {
    clientsDetails.EntityData.YFilter = clientsDetails.YFilter
    clientsDetails.EntityData.YangName = "clients-details"
    clientsDetails.EntityData.BundleName = "cisco_ios_xr"
    clientsDetails.EntityData.ParentYangName = "l2rib"
    clientsDetails.EntityData.SegmentPath = "clients-details"
    clientsDetails.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/" + clientsDetails.EntityData.SegmentPath
    clientsDetails.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clientsDetails.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clientsDetails.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clientsDetails.EntityData.Children = types.NewOrderedMap()
    clientsDetails.EntityData.Children.Append("clients-detail", types.YChild{"ClientsDetail", nil})
    for i := range clientsDetails.ClientsDetail {
        clientsDetails.EntityData.Children.Append(types.GetSegmentPath(clientsDetails.ClientsDetail[i]), types.YChild{"ClientsDetail", clientsDetails.ClientsDetail[i]})
    }
    clientsDetails.EntityData.Leafs = types.NewOrderedMap()

    clientsDetails.EntityData.YListKeys = []string {}

    return &(clientsDetails.EntityData)
}

// L2rib_ClientsDetails_ClientsDetail
// L2RIB clients detail information
type L2rib_ClientsDetails_ClientsDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Client ID. The type is interface{} with range:
    // 0..4294967295.
    ClientId interface{}

    // Number of Producers. The type is interface{} with range: 0..255.
    ProducerCount interface{}

    // Last Update Timestamp. The type is interface{} with range:
    // 0..18446744073709551615.
    LastUpdateTimestamp interface{}

    // Non-detail Client bag.
    Client L2rib_ClientsDetails_ClientsDetail_Client

    // Registration table statistics.
    RegistrationTableStatistics L2rib_ClientsDetails_ClientsDetail_RegistrationTableStatistics

    // List of Producers. The type is slice of
    // L2rib_ClientsDetails_ClientsDetail_ProducerArray.
    ProducerArray []*L2rib_ClientsDetails_ClientsDetail_ProducerArray
}

func (clientsDetail *L2rib_ClientsDetails_ClientsDetail) GetEntityData() *types.CommonEntityData {
    clientsDetail.EntityData.YFilter = clientsDetail.YFilter
    clientsDetail.EntityData.YangName = "clients-detail"
    clientsDetail.EntityData.BundleName = "cisco_ios_xr"
    clientsDetail.EntityData.ParentYangName = "clients-details"
    clientsDetail.EntityData.SegmentPath = "clients-detail" + types.AddKeyToken(clientsDetail.ClientId, "client-id")
    clientsDetail.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/clients-details/" + clientsDetail.EntityData.SegmentPath
    clientsDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clientsDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clientsDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clientsDetail.EntityData.Children = types.NewOrderedMap()
    clientsDetail.EntityData.Children.Append("client", types.YChild{"Client", &clientsDetail.Client})
    clientsDetail.EntityData.Children.Append("registration-table-statistics", types.YChild{"RegistrationTableStatistics", &clientsDetail.RegistrationTableStatistics})
    clientsDetail.EntityData.Children.Append("producer-array", types.YChild{"ProducerArray", nil})
    for i := range clientsDetail.ProducerArray {
        types.SetYListKey(clientsDetail.ProducerArray[i], i)
        clientsDetail.EntityData.Children.Append(types.GetSegmentPath(clientsDetail.ProducerArray[i]), types.YChild{"ProducerArray", clientsDetail.ProducerArray[i]})
    }
    clientsDetail.EntityData.Leafs = types.NewOrderedMap()
    clientsDetail.EntityData.Leafs.Append("client-id", types.YLeaf{"ClientId", clientsDetail.ClientId})
    clientsDetail.EntityData.Leafs.Append("producer-count", types.YLeaf{"ProducerCount", clientsDetail.ProducerCount})
    clientsDetail.EntityData.Leafs.Append("last-update-timestamp", types.YLeaf{"LastUpdateTimestamp", clientsDetail.LastUpdateTimestamp})

    clientsDetail.EntityData.YListKeys = []string {"ClientId"}

    return &(clientsDetail.EntityData)
}

// L2rib_ClientsDetails_ClientsDetail_Client
// Non-detail Client bag
type L2rib_ClientsDetails_ClientsDetail_Client struct {
    EntityData types.CommonEntityData
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

func (client *L2rib_ClientsDetails_ClientsDetail_Client) GetEntityData() *types.CommonEntityData {
    client.EntityData.YFilter = client.YFilter
    client.EntityData.YangName = "client"
    client.EntityData.BundleName = "cisco_ios_xr"
    client.EntityData.ParentYangName = "clients-detail"
    client.EntityData.SegmentPath = "client"
    client.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/clients-details/clients-detail/" + client.EntityData.SegmentPath
    client.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    client.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    client.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    client.EntityData.Children = types.NewOrderedMap()
    client.EntityData.Leafs = types.NewOrderedMap()
    client.EntityData.Leafs.Append("client-id-xr", types.YLeaf{"ClientIdXr", client.ClientIdXr})
    client.EntityData.Leafs.Append("process-id", types.YLeaf{"ProcessId", client.ProcessId})
    client.EntityData.Leafs.Append("node-id", types.YLeaf{"NodeId", client.NodeId})
    client.EntityData.Leafs.Append("proc-name", types.YLeaf{"ProcName", client.ProcName})
    client.EntityData.Leafs.Append("proc-suffix", types.YLeaf{"ProcSuffix", client.ProcSuffix})

    client.EntityData.YListKeys = []string {}

    return &(client.EntityData)
}

// L2rib_ClientsDetails_ClientsDetail_RegistrationTableStatistics
// Registration table statistics
type L2rib_ClientsDetails_ClientsDetail_RegistrationTableStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Producer ID. The type is L2ribBagProducerId.
    ProducerId interface{}

    // Producer Name. The type is string.
    ProducerName interface{}

    // Statistics.
    Statistics L2rib_ClientsDetails_ClientsDetail_RegistrationTableStatistics_Statistics
}

func (registrationTableStatistics *L2rib_ClientsDetails_ClientsDetail_RegistrationTableStatistics) GetEntityData() *types.CommonEntityData {
    registrationTableStatistics.EntityData.YFilter = registrationTableStatistics.YFilter
    registrationTableStatistics.EntityData.YangName = "registration-table-statistics"
    registrationTableStatistics.EntityData.BundleName = "cisco_ios_xr"
    registrationTableStatistics.EntityData.ParentYangName = "clients-detail"
    registrationTableStatistics.EntityData.SegmentPath = "registration-table-statistics"
    registrationTableStatistics.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/clients-details/clients-detail/" + registrationTableStatistics.EntityData.SegmentPath
    registrationTableStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    registrationTableStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    registrationTableStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    registrationTableStatistics.EntityData.Children = types.NewOrderedMap()
    registrationTableStatistics.EntityData.Children.Append("statistics", types.YChild{"Statistics", &registrationTableStatistics.Statistics})
    registrationTableStatistics.EntityData.Leafs = types.NewOrderedMap()
    registrationTableStatistics.EntityData.Leafs.Append("producer-id", types.YLeaf{"ProducerId", registrationTableStatistics.ProducerId})
    registrationTableStatistics.EntityData.Leafs.Append("producer-name", types.YLeaf{"ProducerName", registrationTableStatistics.ProducerName})

    registrationTableStatistics.EntityData.YListKeys = []string {}

    return &(registrationTableStatistics.EntityData)
}

// L2rib_ClientsDetails_ClientsDetail_RegistrationTableStatistics_Statistics
// Statistics
type L2rib_ClientsDetails_ClientsDetail_RegistrationTableStatistics_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Memory Size. The type is interface{} with range: 0..4294967295.
    MemorySize interface{}

    // Number of Objects. The type is interface{} with range: 0..4294967295.
    ObjectCount interface{}

    // End of Interval Timestamp. The type is interface{} with range:
    // 0..18446744073709551615.
    EndofIntervalTs interface{}

    // Extended Counters. The type is slice of
    // L2rib_ClientsDetails_ClientsDetail_RegistrationTableStatistics_Statistics_ExtendedCounter.
    ExtendedCounter []*L2rib_ClientsDetails_ClientsDetail_RegistrationTableStatistics_Statistics_ExtendedCounter
}

func (statistics *L2rib_ClientsDetails_ClientsDetail_RegistrationTableStatistics_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "registration-table-statistics"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/clients-details/clients-detail/registration-table-statistics/" + statistics.EntityData.SegmentPath
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = types.NewOrderedMap()
    statistics.EntityData.Children.Append("extended-counter", types.YChild{"ExtendedCounter", nil})
    for i := range statistics.ExtendedCounter {
        types.SetYListKey(statistics.ExtendedCounter[i], i)
        statistics.EntityData.Children.Append(types.GetSegmentPath(statistics.ExtendedCounter[i]), types.YChild{"ExtendedCounter", statistics.ExtendedCounter[i]})
    }
    statistics.EntityData.Leafs = types.NewOrderedMap()
    statistics.EntityData.Leafs.Append("memory-size", types.YLeaf{"MemorySize", statistics.MemorySize})
    statistics.EntityData.Leafs.Append("object-count", types.YLeaf{"ObjectCount", statistics.ObjectCount})
    statistics.EntityData.Leafs.Append("endof-interval-ts", types.YLeaf{"EndofIntervalTs", statistics.EndofIntervalTs})

    statistics.EntityData.YListKeys = []string {}

    return &(statistics.EntityData)
}

// L2rib_ClientsDetails_ClientsDetail_RegistrationTableStatistics_Statistics_ExtendedCounter
// Extended Counters
type L2rib_ClientsDetails_ClientsDetail_RegistrationTableStatistics_Statistics_ExtendedCounter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // CounterType. The type is interface{} with range: 0..255.
    CounterType interface{}

    // CounterName. The type is string.
    CounterName interface{}

    // Real-clock timestamp in msec of first event. The type is interface{} with
    // range: 0..18446744073709551615.
    L2rbFirstEventTs interface{}

    // Real-clock timestamp in msec of last event. The type is interface{} with
    // range: 0..18446744073709551615.
    L2rbLastEventTs interface{}

    // number of events in interval. The type is interface{} with range:
    // 0..4294967295.
    L2rbIntervalEventCount interface{}

    // total number of events. The type is interface{} with range: 0..4294967295.
    L2rbTotalEventCount interface{}
}

func (extendedCounter *L2rib_ClientsDetails_ClientsDetail_RegistrationTableStatistics_Statistics_ExtendedCounter) GetEntityData() *types.CommonEntityData {
    extendedCounter.EntityData.YFilter = extendedCounter.YFilter
    extendedCounter.EntityData.YangName = "extended-counter"
    extendedCounter.EntityData.BundleName = "cisco_ios_xr"
    extendedCounter.EntityData.ParentYangName = "statistics"
    extendedCounter.EntityData.SegmentPath = "extended-counter" + types.AddNoKeyToken(extendedCounter)
    extendedCounter.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/clients-details/clients-detail/registration-table-statistics/statistics/" + extendedCounter.EntityData.SegmentPath
    extendedCounter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    extendedCounter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    extendedCounter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    extendedCounter.EntityData.Children = types.NewOrderedMap()
    extendedCounter.EntityData.Leafs = types.NewOrderedMap()
    extendedCounter.EntityData.Leafs.Append("counter-type", types.YLeaf{"CounterType", extendedCounter.CounterType})
    extendedCounter.EntityData.Leafs.Append("counter-name", types.YLeaf{"CounterName", extendedCounter.CounterName})
    extendedCounter.EntityData.Leafs.Append("l2rb-first-event-ts", types.YLeaf{"L2rbFirstEventTs", extendedCounter.L2rbFirstEventTs})
    extendedCounter.EntityData.Leafs.Append("l2rb-last-event-ts", types.YLeaf{"L2rbLastEventTs", extendedCounter.L2rbLastEventTs})
    extendedCounter.EntityData.Leafs.Append("l2rb-interval-event-count", types.YLeaf{"L2rbIntervalEventCount", extendedCounter.L2rbIntervalEventCount})
    extendedCounter.EntityData.Leafs.Append("l2rb-total-event-count", types.YLeaf{"L2rbTotalEventCount", extendedCounter.L2rbTotalEventCount})

    extendedCounter.EntityData.YListKeys = []string {}

    return &(extendedCounter.EntityData)
}

// L2rib_ClientsDetails_ClientsDetail_ProducerArray
// List of Producers
type L2rib_ClientsDetails_ClientsDetail_ProducerArray struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

func (producerArray *L2rib_ClientsDetails_ClientsDetail_ProducerArray) GetEntityData() *types.CommonEntityData {
    producerArray.EntityData.YFilter = producerArray.YFilter
    producerArray.EntityData.YangName = "producer-array"
    producerArray.EntityData.BundleName = "cisco_ios_xr"
    producerArray.EntityData.ParentYangName = "clients-detail"
    producerArray.EntityData.SegmentPath = "producer-array" + types.AddNoKeyToken(producerArray)
    producerArray.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/clients-details/clients-detail/" + producerArray.EntityData.SegmentPath
    producerArray.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    producerArray.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    producerArray.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    producerArray.EntityData.Children = types.NewOrderedMap()
    producerArray.EntityData.Leafs = types.NewOrderedMap()
    producerArray.EntityData.Leafs.Append("object-type", types.YLeaf{"ObjectType", producerArray.ObjectType})
    producerArray.EntityData.Leafs.Append("producer-id", types.YLeaf{"ProducerId", producerArray.ProducerId})
    producerArray.EntityData.Leafs.Append("producer-name", types.YLeaf{"ProducerName", producerArray.ProducerName})
    producerArray.EntityData.Leafs.Append("admin-distance", types.YLeaf{"AdminDistance", producerArray.AdminDistance})
    producerArray.EntityData.Leafs.Append("purge-time", types.YLeaf{"PurgeTime", producerArray.PurgeTime})

    producerArray.EntityData.YListKeys = []string {}

    return &(producerArray.EntityData)
}

// L2rib_EviChildTables
// Container for all EVI Child Tables
type L2rib_EviChildTables struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2RIB EVPN EVI MAC IP Detail table.
    MacipDetails L2rib_EviChildTables_MacipDetails

    // L2RIB EVPN EVI MAC IP table.
    MacIps L2rib_EviChildTables_MacIps

    // L2RIB EVPN EVI MAC table.
    Macs L2rib_EviChildTables_Macs

    // L2RIB EVPN EVI IMET table.
    Imets L2rib_EviChildTables_Imets

    // L2RIB EVPN EVI MAC Detail table.
    MacDetails L2rib_EviChildTables_MacDetails

    // L2RIB EVPN EVI IMET Detail table.
    ImetDetails L2rib_EviChildTables_ImetDetails
}

func (eviChildTables *L2rib_EviChildTables) GetEntityData() *types.CommonEntityData {
    eviChildTables.EntityData.YFilter = eviChildTables.YFilter
    eviChildTables.EntityData.YangName = "evi-child-tables"
    eviChildTables.EntityData.BundleName = "cisco_ios_xr"
    eviChildTables.EntityData.ParentYangName = "l2rib"
    eviChildTables.EntityData.SegmentPath = "evi-child-tables"
    eviChildTables.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/" + eviChildTables.EntityData.SegmentPath
    eviChildTables.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    eviChildTables.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    eviChildTables.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    eviChildTables.EntityData.Children = types.NewOrderedMap()
    eviChildTables.EntityData.Children.Append("macip-details", types.YChild{"MacipDetails", &eviChildTables.MacipDetails})
    eviChildTables.EntityData.Children.Append("mac-ips", types.YChild{"MacIps", &eviChildTables.MacIps})
    eviChildTables.EntityData.Children.Append("macs", types.YChild{"Macs", &eviChildTables.Macs})
    eviChildTables.EntityData.Children.Append("imets", types.YChild{"Imets", &eviChildTables.Imets})
    eviChildTables.EntityData.Children.Append("mac-details", types.YChild{"MacDetails", &eviChildTables.MacDetails})
    eviChildTables.EntityData.Children.Append("imet-details", types.YChild{"ImetDetails", &eviChildTables.ImetDetails})
    eviChildTables.EntityData.Leafs = types.NewOrderedMap()

    eviChildTables.EntityData.YListKeys = []string {}

    return &(eviChildTables.EntityData)
}

// L2rib_EviChildTables_MacipDetails
// L2RIB EVPN EVI MAC IP Detail table
type L2rib_EviChildTables_MacipDetails struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2RIB EVPN MAC IP Detail table. The type is slice of
    // L2rib_EviChildTables_MacipDetails_MacipDetail.
    MacipDetail []*L2rib_EviChildTables_MacipDetails_MacipDetail
}

func (macipDetails *L2rib_EviChildTables_MacipDetails) GetEntityData() *types.CommonEntityData {
    macipDetails.EntityData.YFilter = macipDetails.YFilter
    macipDetails.EntityData.YangName = "macip-details"
    macipDetails.EntityData.BundleName = "cisco_ios_xr"
    macipDetails.EntityData.ParentYangName = "evi-child-tables"
    macipDetails.EntityData.SegmentPath = "macip-details"
    macipDetails.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/" + macipDetails.EntityData.SegmentPath
    macipDetails.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macipDetails.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macipDetails.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macipDetails.EntityData.Children = types.NewOrderedMap()
    macipDetails.EntityData.Children.Append("macip-detail", types.YChild{"MacipDetail", nil})
    for i := range macipDetails.MacipDetail {
        types.SetYListKey(macipDetails.MacipDetail[i], i)
        macipDetails.EntityData.Children.Append(types.GetSegmentPath(macipDetails.MacipDetail[i]), types.YChild{"MacipDetail", macipDetails.MacipDetail[i]})
    }
    macipDetails.EntityData.Leafs = types.NewOrderedMap()

    macipDetails.EntityData.YListKeys = []string {}

    return &(macipDetails.EntityData)
}

// L2rib_EviChildTables_MacipDetails_MacipDetail
// L2RIB EVPN MAC IP Detail table
type L2rib_EviChildTables_MacipDetails_MacipDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // EVPN ID. The type is interface{} with range: 0..4294967295.
    Evi interface{}

    // Tag ID. The type is interface{} with range: 0..4294967295.
    TagId interface{}

    // MAC IP Address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddr interface{}

    // IP Address. The type is one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IpAddr interface{}

    // Admin distance. The type is interface{} with range: 0..4294967295.
    AdminDist interface{}

    // Producer ID. The type is interface{} with range: 0..4294967295.
    ProdId interface{}

    // MAC-IP route sequence Number. The type is interface{} with range:
    // 0..4294967295.
    SequenceNumber interface{}

    // MAC-IP route flags. The type is string with pattern: [0-9a-fA-F]{1,8}.
    Flags interface{}

    // SOO. The type is interface{} with range: 0..4294967295.
    Soo interface{}

    // Last Update Timestamp. The type is interface{} with range:
    // 0..18446744073709551615.
    LastUpdateTimestamp interface{}

    // MAC-IP Route.
    MacIpRoute L2rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute

    // Mac-IP Route Opaque Data TLV.
    RtTlv L2rib_EviChildTables_MacipDetails_MacipDetail_RtTlv

    // Mac-IP Route Opaque NH TLV.
    NhTlv L2rib_EviChildTables_MacipDetails_MacipDetail_NhTlv
}

func (macipDetail *L2rib_EviChildTables_MacipDetails_MacipDetail) GetEntityData() *types.CommonEntityData {
    macipDetail.EntityData.YFilter = macipDetail.YFilter
    macipDetail.EntityData.YangName = "macip-detail"
    macipDetail.EntityData.BundleName = "cisco_ios_xr"
    macipDetail.EntityData.ParentYangName = "macip-details"
    macipDetail.EntityData.SegmentPath = "macip-detail" + types.AddNoKeyToken(macipDetail)
    macipDetail.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/macip-details/" + macipDetail.EntityData.SegmentPath
    macipDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macipDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macipDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macipDetail.EntityData.Children = types.NewOrderedMap()
    macipDetail.EntityData.Children.Append("mac-ip-route", types.YChild{"MacIpRoute", &macipDetail.MacIpRoute})
    macipDetail.EntityData.Children.Append("rt-tlv", types.YChild{"RtTlv", &macipDetail.RtTlv})
    macipDetail.EntityData.Children.Append("nh-tlv", types.YChild{"NhTlv", &macipDetail.NhTlv})
    macipDetail.EntityData.Leafs = types.NewOrderedMap()
    macipDetail.EntityData.Leafs.Append("evi", types.YLeaf{"Evi", macipDetail.Evi})
    macipDetail.EntityData.Leafs.Append("tag-id", types.YLeaf{"TagId", macipDetail.TagId})
    macipDetail.EntityData.Leafs.Append("mac-addr", types.YLeaf{"MacAddr", macipDetail.MacAddr})
    macipDetail.EntityData.Leafs.Append("ip-addr", types.YLeaf{"IpAddr", macipDetail.IpAddr})
    macipDetail.EntityData.Leafs.Append("admin-dist", types.YLeaf{"AdminDist", macipDetail.AdminDist})
    macipDetail.EntityData.Leafs.Append("prod-id", types.YLeaf{"ProdId", macipDetail.ProdId})
    macipDetail.EntityData.Leafs.Append("sequence-number", types.YLeaf{"SequenceNumber", macipDetail.SequenceNumber})
    macipDetail.EntityData.Leafs.Append("flags", types.YLeaf{"Flags", macipDetail.Flags})
    macipDetail.EntityData.Leafs.Append("soo", types.YLeaf{"Soo", macipDetail.Soo})
    macipDetail.EntityData.Leafs.Append("last-update-timestamp", types.YLeaf{"LastUpdateTimestamp", macipDetail.LastUpdateTimestamp})

    macipDetail.EntityData.YListKeys = []string {}

    return &(macipDetail.EntityData)
}

// L2rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute
// MAC-IP Route
type L2rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute struct {
    EntityData types.CommonEntityData
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
    NextHop L2rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute_NextHop
}

func (macIpRoute *L2rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute) GetEntityData() *types.CommonEntityData {
    macIpRoute.EntityData.YFilter = macIpRoute.YFilter
    macIpRoute.EntityData.YangName = "mac-ip-route"
    macIpRoute.EntityData.BundleName = "cisco_ios_xr"
    macIpRoute.EntityData.ParentYangName = "macip-detail"
    macIpRoute.EntityData.SegmentPath = "mac-ip-route"
    macIpRoute.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/macip-details/macip-detail/" + macIpRoute.EntityData.SegmentPath
    macIpRoute.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macIpRoute.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macIpRoute.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macIpRoute.EntityData.Children = types.NewOrderedMap()
    macIpRoute.EntityData.Children.Append("next-hop", types.YChild{"NextHop", &macIpRoute.NextHop})
    macIpRoute.EntityData.Leafs = types.NewOrderedMap()
    macIpRoute.EntityData.Leafs.Append("mac-address", types.YLeaf{"MacAddress", macIpRoute.MacAddress})
    macIpRoute.EntityData.Leafs.Append("ip-address", types.YLeaf{"IpAddress", macIpRoute.IpAddress})
    macIpRoute.EntityData.Leafs.Append("admin-distance", types.YLeaf{"AdminDistance", macIpRoute.AdminDistance})
    macIpRoute.EntityData.Leafs.Append("producer-id", types.YLeaf{"ProducerId", macIpRoute.ProducerId})
    macIpRoute.EntityData.Leafs.Append("topology-id", types.YLeaf{"TopologyId", macIpRoute.TopologyId})

    macIpRoute.EntityData.YListKeys = []string {}

    return &(macIpRoute.EntityData)
}

// L2rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute_NextHop
// Next Hop
type L2rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute_NextHop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Next-hop TOPOLOGY ID. The type is interface{} with range: 0..4294967295.
    TopologyId interface{}

    // Next-hop flags. The type is interface{} with range: 0..65535.
    Flags interface{}

    // Next hop.
    NextHop L2rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute_NextHop_NextHop
}

func (nextHop *L2rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute_NextHop) GetEntityData() *types.CommonEntityData {
    nextHop.EntityData.YFilter = nextHop.YFilter
    nextHop.EntityData.YangName = "next-hop"
    nextHop.EntityData.BundleName = "cisco_ios_xr"
    nextHop.EntityData.ParentYangName = "mac-ip-route"
    nextHop.EntityData.SegmentPath = "next-hop"
    nextHop.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/macip-details/macip-detail/mac-ip-route/" + nextHop.EntityData.SegmentPath
    nextHop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nextHop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nextHop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nextHop.EntityData.Children = types.NewOrderedMap()
    nextHop.EntityData.Children.Append("next-hop", types.YChild{"NextHop", &nextHop.NextHop})
    nextHop.EntityData.Leafs = types.NewOrderedMap()
    nextHop.EntityData.Leafs.Append("topology-id", types.YLeaf{"TopologyId", nextHop.TopologyId})
    nextHop.EntityData.Leafs.Append("flags", types.YLeaf{"Flags", nextHop.Flags})

    nextHop.EntityData.YListKeys = []string {}

    return &(nextHop.EntityData)
}

// L2rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute_NextHop_NextHop
// Next hop
type L2rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute_NextHop_NextHop struct {
    EntityData types.CommonEntityData
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
    // [a-zA-Z0-9._/-]+.
    InterfaceHandle interface{}

    // XID Next Hop. The type is interface{} with range: 0..4294967295.
    Xid interface{}

    // Labeled Next Hop.
    Labeled L2rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute_NextHop_NextHop_Labeled
}

func (nextHop *L2rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute_NextHop_NextHop) GetEntityData() *types.CommonEntityData {
    nextHop.EntityData.YFilter = nextHop.YFilter
    nextHop.EntityData.YangName = "next-hop"
    nextHop.EntityData.BundleName = "cisco_ios_xr"
    nextHop.EntityData.ParentYangName = "next-hop"
    nextHop.EntityData.SegmentPath = "next-hop"
    nextHop.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/macip-details/macip-detail/mac-ip-route/next-hop/" + nextHop.EntityData.SegmentPath
    nextHop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nextHop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nextHop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nextHop.EntityData.Children = types.NewOrderedMap()
    nextHop.EntityData.Children.Append("labeled", types.YChild{"Labeled", &nextHop.Labeled})
    nextHop.EntityData.Leafs = types.NewOrderedMap()
    nextHop.EntityData.Leafs.Append("type", types.YLeaf{"Type", nextHop.Type})
    nextHop.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", nextHop.Ipv4})
    nextHop.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", nextHop.Ipv6})
    nextHop.EntityData.Leafs.Append("mac", types.YLeaf{"Mac", nextHop.Mac})
    nextHop.EntityData.Leafs.Append("interface-handle", types.YLeaf{"InterfaceHandle", nextHop.InterfaceHandle})
    nextHop.EntityData.Leafs.Append("xid", types.YLeaf{"Xid", nextHop.Xid})

    nextHop.EntityData.YListKeys = []string {}

    return &(nextHop.EntityData)
}

// L2rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute_NextHop_NextHop_Labeled
// Labeled Next Hop
type L2rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute_NextHop_NextHop_Labeled struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2RIB Address Family. The type is L2ribAfi.
    AddressFamily interface{}

    // IP Address. The type is string.
    IpAddress interface{}

    // Label. The type is interface{} with range: 0..4294967295.
    Label interface{}

    // Internal Label. The type is bool.
    Internal interface{}
}

func (labeled *L2rib_EviChildTables_MacipDetails_MacipDetail_MacIpRoute_NextHop_NextHop_Labeled) GetEntityData() *types.CommonEntityData {
    labeled.EntityData.YFilter = labeled.YFilter
    labeled.EntityData.YangName = "labeled"
    labeled.EntityData.BundleName = "cisco_ios_xr"
    labeled.EntityData.ParentYangName = "next-hop"
    labeled.EntityData.SegmentPath = "labeled"
    labeled.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/macip-details/macip-detail/mac-ip-route/next-hop/next-hop/" + labeled.EntityData.SegmentPath
    labeled.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    labeled.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    labeled.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    labeled.EntityData.Children = types.NewOrderedMap()
    labeled.EntityData.Leafs = types.NewOrderedMap()
    labeled.EntityData.Leafs.Append("address-family", types.YLeaf{"AddressFamily", labeled.AddressFamily})
    labeled.EntityData.Leafs.Append("ip-address", types.YLeaf{"IpAddress", labeled.IpAddress})
    labeled.EntityData.Leafs.Append("label", types.YLeaf{"Label", labeled.Label})
    labeled.EntityData.Leafs.Append("internal", types.YLeaf{"Internal", labeled.Internal})

    labeled.EntityData.YListKeys = []string {}

    return &(labeled.EntityData)
}

// L2rib_EviChildTables_MacipDetails_MacipDetail_RtTlv
// Mac-IP Route Opaque Data TLV
type L2rib_EviChildTables_MacipDetails_MacipDetail_RtTlv struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TLV Type. The type is interface{} with range: 0..65535.
    TlvType interface{}

    // TLV Length. The type is interface{} with range: 0..65535.
    TlvLen interface{}

    // TLV Value. The type is slice of
    // L2rib_EviChildTables_MacipDetails_MacipDetail_RtTlv_TlvVal.
    TlvVal []*L2rib_EviChildTables_MacipDetails_MacipDetail_RtTlv_TlvVal
}

func (rtTlv *L2rib_EviChildTables_MacipDetails_MacipDetail_RtTlv) GetEntityData() *types.CommonEntityData {
    rtTlv.EntityData.YFilter = rtTlv.YFilter
    rtTlv.EntityData.YangName = "rt-tlv"
    rtTlv.EntityData.BundleName = "cisco_ios_xr"
    rtTlv.EntityData.ParentYangName = "macip-detail"
    rtTlv.EntityData.SegmentPath = "rt-tlv"
    rtTlv.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/macip-details/macip-detail/" + rtTlv.EntityData.SegmentPath
    rtTlv.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rtTlv.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rtTlv.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rtTlv.EntityData.Children = types.NewOrderedMap()
    rtTlv.EntityData.Children.Append("tlv-val", types.YChild{"TlvVal", nil})
    for i := range rtTlv.TlvVal {
        types.SetYListKey(rtTlv.TlvVal[i], i)
        rtTlv.EntityData.Children.Append(types.GetSegmentPath(rtTlv.TlvVal[i]), types.YChild{"TlvVal", rtTlv.TlvVal[i]})
    }
    rtTlv.EntityData.Leafs = types.NewOrderedMap()
    rtTlv.EntityData.Leafs.Append("tlv-type", types.YLeaf{"TlvType", rtTlv.TlvType})
    rtTlv.EntityData.Leafs.Append("tlv-len", types.YLeaf{"TlvLen", rtTlv.TlvLen})

    rtTlv.EntityData.YListKeys = []string {}

    return &(rtTlv.EntityData)
}

// L2rib_EviChildTables_MacipDetails_MacipDetail_RtTlv_TlvVal
// TLV Value
type L2rib_EviChildTables_MacipDetails_MacipDetail_RtTlv_TlvVal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // The type is interface{} with range: 0..255.
    Entry interface{}
}

func (tlvVal *L2rib_EviChildTables_MacipDetails_MacipDetail_RtTlv_TlvVal) GetEntityData() *types.CommonEntityData {
    tlvVal.EntityData.YFilter = tlvVal.YFilter
    tlvVal.EntityData.YangName = "tlv-val"
    tlvVal.EntityData.BundleName = "cisco_ios_xr"
    tlvVal.EntityData.ParentYangName = "rt-tlv"
    tlvVal.EntityData.SegmentPath = "tlv-val" + types.AddNoKeyToken(tlvVal)
    tlvVal.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/macip-details/macip-detail/rt-tlv/" + tlvVal.EntityData.SegmentPath
    tlvVal.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tlvVal.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tlvVal.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tlvVal.EntityData.Children = types.NewOrderedMap()
    tlvVal.EntityData.Leafs = types.NewOrderedMap()
    tlvVal.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", tlvVal.Entry})

    tlvVal.EntityData.YListKeys = []string {}

    return &(tlvVal.EntityData)
}

// L2rib_EviChildTables_MacipDetails_MacipDetail_NhTlv
// Mac-IP Route Opaque NH TLV
type L2rib_EviChildTables_MacipDetails_MacipDetail_NhTlv struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TLV Type. The type is interface{} with range: 0..65535.
    TlvType interface{}

    // TLV Length. The type is interface{} with range: 0..65535.
    TlvLen interface{}

    // TLV Value. The type is slice of
    // L2rib_EviChildTables_MacipDetails_MacipDetail_NhTlv_TlvVal.
    TlvVal []*L2rib_EviChildTables_MacipDetails_MacipDetail_NhTlv_TlvVal
}

func (nhTlv *L2rib_EviChildTables_MacipDetails_MacipDetail_NhTlv) GetEntityData() *types.CommonEntityData {
    nhTlv.EntityData.YFilter = nhTlv.YFilter
    nhTlv.EntityData.YangName = "nh-tlv"
    nhTlv.EntityData.BundleName = "cisco_ios_xr"
    nhTlv.EntityData.ParentYangName = "macip-detail"
    nhTlv.EntityData.SegmentPath = "nh-tlv"
    nhTlv.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/macip-details/macip-detail/" + nhTlv.EntityData.SegmentPath
    nhTlv.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nhTlv.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nhTlv.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nhTlv.EntityData.Children = types.NewOrderedMap()
    nhTlv.EntityData.Children.Append("tlv-val", types.YChild{"TlvVal", nil})
    for i := range nhTlv.TlvVal {
        types.SetYListKey(nhTlv.TlvVal[i], i)
        nhTlv.EntityData.Children.Append(types.GetSegmentPath(nhTlv.TlvVal[i]), types.YChild{"TlvVal", nhTlv.TlvVal[i]})
    }
    nhTlv.EntityData.Leafs = types.NewOrderedMap()
    nhTlv.EntityData.Leafs.Append("tlv-type", types.YLeaf{"TlvType", nhTlv.TlvType})
    nhTlv.EntityData.Leafs.Append("tlv-len", types.YLeaf{"TlvLen", nhTlv.TlvLen})

    nhTlv.EntityData.YListKeys = []string {}

    return &(nhTlv.EntityData)
}

// L2rib_EviChildTables_MacipDetails_MacipDetail_NhTlv_TlvVal
// TLV Value
type L2rib_EviChildTables_MacipDetails_MacipDetail_NhTlv_TlvVal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // The type is interface{} with range: 0..255.
    Entry interface{}
}

func (tlvVal *L2rib_EviChildTables_MacipDetails_MacipDetail_NhTlv_TlvVal) GetEntityData() *types.CommonEntityData {
    tlvVal.EntityData.YFilter = tlvVal.YFilter
    tlvVal.EntityData.YangName = "tlv-val"
    tlvVal.EntityData.BundleName = "cisco_ios_xr"
    tlvVal.EntityData.ParentYangName = "nh-tlv"
    tlvVal.EntityData.SegmentPath = "tlv-val" + types.AddNoKeyToken(tlvVal)
    tlvVal.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/macip-details/macip-detail/nh-tlv/" + tlvVal.EntityData.SegmentPath
    tlvVal.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tlvVal.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tlvVal.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tlvVal.EntityData.Children = types.NewOrderedMap()
    tlvVal.EntityData.Leafs = types.NewOrderedMap()
    tlvVal.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", tlvVal.Entry})

    tlvVal.EntityData.YListKeys = []string {}

    return &(tlvVal.EntityData)
}

// L2rib_EviChildTables_MacIps
// L2RIB EVPN EVI MAC IP table
type L2rib_EviChildTables_MacIps struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2RIB EVPN MAC IP table. The type is slice of
    // L2rib_EviChildTables_MacIps_MacIp.
    MacIp []*L2rib_EviChildTables_MacIps_MacIp
}

func (macIps *L2rib_EviChildTables_MacIps) GetEntityData() *types.CommonEntityData {
    macIps.EntityData.YFilter = macIps.YFilter
    macIps.EntityData.YangName = "mac-ips"
    macIps.EntityData.BundleName = "cisco_ios_xr"
    macIps.EntityData.ParentYangName = "evi-child-tables"
    macIps.EntityData.SegmentPath = "mac-ips"
    macIps.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/" + macIps.EntityData.SegmentPath
    macIps.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macIps.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macIps.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macIps.EntityData.Children = types.NewOrderedMap()
    macIps.EntityData.Children.Append("mac-ip", types.YChild{"MacIp", nil})
    for i := range macIps.MacIp {
        types.SetYListKey(macIps.MacIp[i], i)
        macIps.EntityData.Children.Append(types.GetSegmentPath(macIps.MacIp[i]), types.YChild{"MacIp", macIps.MacIp[i]})
    }
    macIps.EntityData.Leafs = types.NewOrderedMap()

    macIps.EntityData.YListKeys = []string {}

    return &(macIps.EntityData)
}

// L2rib_EviChildTables_MacIps_MacIp
// L2RIB EVPN MAC IP table
type L2rib_EviChildTables_MacIps_MacIp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // EVPN ID. The type is interface{} with range: 0..4294967295.
    Evi interface{}

    // Tag ID. The type is interface{} with range: 0..4294967295.
    TagId interface{}

    // MAC-IP Address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddr interface{}

    // IP Address. The type is one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IpAddr interface{}

    // Admin distance. The type is interface{} with range: 0..4294967295.
    AdminDist interface{}

    // Producer ID. The type is interface{} with range: 0..4294967295.
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
    NextHop L2rib_EviChildTables_MacIps_MacIp_NextHop
}

func (macIp *L2rib_EviChildTables_MacIps_MacIp) GetEntityData() *types.CommonEntityData {
    macIp.EntityData.YFilter = macIp.YFilter
    macIp.EntityData.YangName = "mac-ip"
    macIp.EntityData.BundleName = "cisco_ios_xr"
    macIp.EntityData.ParentYangName = "mac-ips"
    macIp.EntityData.SegmentPath = "mac-ip" + types.AddNoKeyToken(macIp)
    macIp.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/mac-ips/" + macIp.EntityData.SegmentPath
    macIp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macIp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macIp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macIp.EntityData.Children = types.NewOrderedMap()
    macIp.EntityData.Children.Append("next-hop", types.YChild{"NextHop", &macIp.NextHop})
    macIp.EntityData.Leafs = types.NewOrderedMap()
    macIp.EntityData.Leafs.Append("evi", types.YLeaf{"Evi", macIp.Evi})
    macIp.EntityData.Leafs.Append("tag-id", types.YLeaf{"TagId", macIp.TagId})
    macIp.EntityData.Leafs.Append("mac-addr", types.YLeaf{"MacAddr", macIp.MacAddr})
    macIp.EntityData.Leafs.Append("ip-addr", types.YLeaf{"IpAddr", macIp.IpAddr})
    macIp.EntityData.Leafs.Append("admin-dist", types.YLeaf{"AdminDist", macIp.AdminDist})
    macIp.EntityData.Leafs.Append("prod-id", types.YLeaf{"ProdId", macIp.ProdId})
    macIp.EntityData.Leafs.Append("mac-address", types.YLeaf{"MacAddress", macIp.MacAddress})
    macIp.EntityData.Leafs.Append("ip-address", types.YLeaf{"IpAddress", macIp.IpAddress})
    macIp.EntityData.Leafs.Append("admin-distance", types.YLeaf{"AdminDistance", macIp.AdminDistance})
    macIp.EntityData.Leafs.Append("producer-id", types.YLeaf{"ProducerId", macIp.ProducerId})
    macIp.EntityData.Leafs.Append("topology-id", types.YLeaf{"TopologyId", macIp.TopologyId})

    macIp.EntityData.YListKeys = []string {}

    return &(macIp.EntityData)
}

// L2rib_EviChildTables_MacIps_MacIp_NextHop
// Next Hop
type L2rib_EviChildTables_MacIps_MacIp_NextHop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Next-hop TOPOLOGY ID. The type is interface{} with range: 0..4294967295.
    TopologyId interface{}

    // Next-hop flags. The type is interface{} with range: 0..65535.
    Flags interface{}

    // Next hop.
    NextHop L2rib_EviChildTables_MacIps_MacIp_NextHop_NextHop
}

func (nextHop *L2rib_EviChildTables_MacIps_MacIp_NextHop) GetEntityData() *types.CommonEntityData {
    nextHop.EntityData.YFilter = nextHop.YFilter
    nextHop.EntityData.YangName = "next-hop"
    nextHop.EntityData.BundleName = "cisco_ios_xr"
    nextHop.EntityData.ParentYangName = "mac-ip"
    nextHop.EntityData.SegmentPath = "next-hop"
    nextHop.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/mac-ips/mac-ip/" + nextHop.EntityData.SegmentPath
    nextHop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nextHop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nextHop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nextHop.EntityData.Children = types.NewOrderedMap()
    nextHop.EntityData.Children.Append("next-hop", types.YChild{"NextHop", &nextHop.NextHop})
    nextHop.EntityData.Leafs = types.NewOrderedMap()
    nextHop.EntityData.Leafs.Append("topology-id", types.YLeaf{"TopologyId", nextHop.TopologyId})
    nextHop.EntityData.Leafs.Append("flags", types.YLeaf{"Flags", nextHop.Flags})

    nextHop.EntityData.YListKeys = []string {}

    return &(nextHop.EntityData)
}

// L2rib_EviChildTables_MacIps_MacIp_NextHop_NextHop
// Next hop
type L2rib_EviChildTables_MacIps_MacIp_NextHop_NextHop struct {
    EntityData types.CommonEntityData
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
    // [a-zA-Z0-9._/-]+.
    InterfaceHandle interface{}

    // XID Next Hop. The type is interface{} with range: 0..4294967295.
    Xid interface{}

    // Labeled Next Hop.
    Labeled L2rib_EviChildTables_MacIps_MacIp_NextHop_NextHop_Labeled
}

func (nextHop *L2rib_EviChildTables_MacIps_MacIp_NextHop_NextHop) GetEntityData() *types.CommonEntityData {
    nextHop.EntityData.YFilter = nextHop.YFilter
    nextHop.EntityData.YangName = "next-hop"
    nextHop.EntityData.BundleName = "cisco_ios_xr"
    nextHop.EntityData.ParentYangName = "next-hop"
    nextHop.EntityData.SegmentPath = "next-hop"
    nextHop.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/mac-ips/mac-ip/next-hop/" + nextHop.EntityData.SegmentPath
    nextHop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nextHop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nextHop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nextHop.EntityData.Children = types.NewOrderedMap()
    nextHop.EntityData.Children.Append("labeled", types.YChild{"Labeled", &nextHop.Labeled})
    nextHop.EntityData.Leafs = types.NewOrderedMap()
    nextHop.EntityData.Leafs.Append("type", types.YLeaf{"Type", nextHop.Type})
    nextHop.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", nextHop.Ipv4})
    nextHop.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", nextHop.Ipv6})
    nextHop.EntityData.Leafs.Append("mac", types.YLeaf{"Mac", nextHop.Mac})
    nextHop.EntityData.Leafs.Append("interface-handle", types.YLeaf{"InterfaceHandle", nextHop.InterfaceHandle})
    nextHop.EntityData.Leafs.Append("xid", types.YLeaf{"Xid", nextHop.Xid})

    nextHop.EntityData.YListKeys = []string {}

    return &(nextHop.EntityData)
}

// L2rib_EviChildTables_MacIps_MacIp_NextHop_NextHop_Labeled
// Labeled Next Hop
type L2rib_EviChildTables_MacIps_MacIp_NextHop_NextHop_Labeled struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2RIB Address Family. The type is L2ribAfi.
    AddressFamily interface{}

    // IP Address. The type is string.
    IpAddress interface{}

    // Label. The type is interface{} with range: 0..4294967295.
    Label interface{}

    // Internal Label. The type is bool.
    Internal interface{}
}

func (labeled *L2rib_EviChildTables_MacIps_MacIp_NextHop_NextHop_Labeled) GetEntityData() *types.CommonEntityData {
    labeled.EntityData.YFilter = labeled.YFilter
    labeled.EntityData.YangName = "labeled"
    labeled.EntityData.BundleName = "cisco_ios_xr"
    labeled.EntityData.ParentYangName = "next-hop"
    labeled.EntityData.SegmentPath = "labeled"
    labeled.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/mac-ips/mac-ip/next-hop/next-hop/" + labeled.EntityData.SegmentPath
    labeled.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    labeled.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    labeled.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    labeled.EntityData.Children = types.NewOrderedMap()
    labeled.EntityData.Leafs = types.NewOrderedMap()
    labeled.EntityData.Leafs.Append("address-family", types.YLeaf{"AddressFamily", labeled.AddressFamily})
    labeled.EntityData.Leafs.Append("ip-address", types.YLeaf{"IpAddress", labeled.IpAddress})
    labeled.EntityData.Leafs.Append("label", types.YLeaf{"Label", labeled.Label})
    labeled.EntityData.Leafs.Append("internal", types.YLeaf{"Internal", labeled.Internal})

    labeled.EntityData.YListKeys = []string {}

    return &(labeled.EntityData)
}

// L2rib_EviChildTables_Macs
// L2RIB EVPN EVI MAC table
type L2rib_EviChildTables_Macs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2RIB EVPN MAC table. The type is slice of L2rib_EviChildTables_Macs_Mac.
    Mac []*L2rib_EviChildTables_Macs_Mac
}

func (macs *L2rib_EviChildTables_Macs) GetEntityData() *types.CommonEntityData {
    macs.EntityData.YFilter = macs.YFilter
    macs.EntityData.YangName = "macs"
    macs.EntityData.BundleName = "cisco_ios_xr"
    macs.EntityData.ParentYangName = "evi-child-tables"
    macs.EntityData.SegmentPath = "macs"
    macs.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/" + macs.EntityData.SegmentPath
    macs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macs.EntityData.Children = types.NewOrderedMap()
    macs.EntityData.Children.Append("mac", types.YChild{"Mac", nil})
    for i := range macs.Mac {
        types.SetYListKey(macs.Mac[i], i)
        macs.EntityData.Children.Append(types.GetSegmentPath(macs.Mac[i]), types.YChild{"Mac", macs.Mac[i]})
    }
    macs.EntityData.Leafs = types.NewOrderedMap()

    macs.EntityData.YListKeys = []string {}

    return &(macs.EntityData)
}

// L2rib_EviChildTables_Macs_Mac
// L2RIB EVPN MAC table
type L2rib_EviChildTables_Macs_Mac struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // EVPN ID. The type is interface{} with range: 0..4294967295.
    Evi interface{}

    // Tag ID. The type is interface{} with range: 0..4294967295.
    TagId interface{}

    // MAC Address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddr interface{}

    // Admin distance. The type is interface{} with range: 0..4294967295.
    AdminDist interface{}

    // Producer ID. The type is interface{} with range: 0..4294967295.
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
    Route L2rib_EviChildTables_Macs_Mac_Route
}

func (mac *L2rib_EviChildTables_Macs_Mac) GetEntityData() *types.CommonEntityData {
    mac.EntityData.YFilter = mac.YFilter
    mac.EntityData.YangName = "mac"
    mac.EntityData.BundleName = "cisco_ios_xr"
    mac.EntityData.ParentYangName = "macs"
    mac.EntityData.SegmentPath = "mac" + types.AddNoKeyToken(mac)
    mac.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/macs/" + mac.EntityData.SegmentPath
    mac.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mac.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mac.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mac.EntityData.Children = types.NewOrderedMap()
    mac.EntityData.Children.Append("route", types.YChild{"Route", &mac.Route})
    mac.EntityData.Leafs = types.NewOrderedMap()
    mac.EntityData.Leafs.Append("evi", types.YLeaf{"Evi", mac.Evi})
    mac.EntityData.Leafs.Append("tag-id", types.YLeaf{"TagId", mac.TagId})
    mac.EntityData.Leafs.Append("mac-addr", types.YLeaf{"MacAddr", mac.MacAddr})
    mac.EntityData.Leafs.Append("admin-dist", types.YLeaf{"AdminDist", mac.AdminDist})
    mac.EntityData.Leafs.Append("prod-id", types.YLeaf{"ProdId", mac.ProdId})
    mac.EntityData.Leafs.Append("mac-address", types.YLeaf{"MacAddress", mac.MacAddress})
    mac.EntityData.Leafs.Append("admin-distance", types.YLeaf{"AdminDistance", mac.AdminDistance})
    mac.EntityData.Leafs.Append("producer-id", types.YLeaf{"ProducerId", mac.ProducerId})
    mac.EntityData.Leafs.Append("topology-id", types.YLeaf{"TopologyId", mac.TopologyId})

    mac.EntityData.YListKeys = []string {}

    return &(mac.EntityData)
}

// L2rib_EviChildTables_Macs_Mac_Route
// MAC route
type L2rib_EviChildTables_Macs_Mac_Route struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type. The type is L2ribMacRoute.
    Type interface{}

    // Regular MAC route.
    Regular L2rib_EviChildTables_Macs_Mac_Route_Regular

    // EVPN ESI MAC route.
    EvpnEsi L2rib_EviChildTables_Macs_Mac_Route_EvpnEsi

    // BMAC route.
    Bmac L2rib_EviChildTables_Macs_Mac_Route_Bmac
}

func (route *L2rib_EviChildTables_Macs_Mac_Route) GetEntityData() *types.CommonEntityData {
    route.EntityData.YFilter = route.YFilter
    route.EntityData.YangName = "route"
    route.EntityData.BundleName = "cisco_ios_xr"
    route.EntityData.ParentYangName = "mac"
    route.EntityData.SegmentPath = "route"
    route.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/macs/mac/" + route.EntityData.SegmentPath
    route.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    route.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    route.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    route.EntityData.Children = types.NewOrderedMap()
    route.EntityData.Children.Append("regular", types.YChild{"Regular", &route.Regular})
    route.EntityData.Children.Append("evpn-esi", types.YChild{"EvpnEsi", &route.EvpnEsi})
    route.EntityData.Children.Append("bmac", types.YChild{"Bmac", &route.Bmac})
    route.EntityData.Leafs = types.NewOrderedMap()
    route.EntityData.Leafs.Append("type", types.YLeaf{"Type", route.Type})

    route.EntityData.YListKeys = []string {}

    return &(route.EntityData)
}

// L2rib_EviChildTables_Macs_Mac_Route_Regular
// Regular MAC route
type L2rib_EviChildTables_Macs_Mac_Route_Regular struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Next Hop.
    NextHop L2rib_EviChildTables_Macs_Mac_Route_Regular_NextHop
}

func (regular *L2rib_EviChildTables_Macs_Mac_Route_Regular) GetEntityData() *types.CommonEntityData {
    regular.EntityData.YFilter = regular.YFilter
    regular.EntityData.YangName = "regular"
    regular.EntityData.BundleName = "cisco_ios_xr"
    regular.EntityData.ParentYangName = "route"
    regular.EntityData.SegmentPath = "regular"
    regular.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/macs/mac/route/" + regular.EntityData.SegmentPath
    regular.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    regular.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    regular.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    regular.EntityData.Children = types.NewOrderedMap()
    regular.EntityData.Children.Append("next-hop", types.YChild{"NextHop", &regular.NextHop})
    regular.EntityData.Leafs = types.NewOrderedMap()

    regular.EntityData.YListKeys = []string {}

    return &(regular.EntityData)
}

// L2rib_EviChildTables_Macs_Mac_Route_Regular_NextHop
// Next Hop
type L2rib_EviChildTables_Macs_Mac_Route_Regular_NextHop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Next-hop TOPOLOGY ID. The type is interface{} with range: 0..4294967295.
    TopologyId interface{}

    // Next-hop flags. The type is interface{} with range: 0..65535.
    Flags interface{}

    // Next hop.
    NextHop L2rib_EviChildTables_Macs_Mac_Route_Regular_NextHop_NextHop
}

func (nextHop *L2rib_EviChildTables_Macs_Mac_Route_Regular_NextHop) GetEntityData() *types.CommonEntityData {
    nextHop.EntityData.YFilter = nextHop.YFilter
    nextHop.EntityData.YangName = "next-hop"
    nextHop.EntityData.BundleName = "cisco_ios_xr"
    nextHop.EntityData.ParentYangName = "regular"
    nextHop.EntityData.SegmentPath = "next-hop"
    nextHop.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/macs/mac/route/regular/" + nextHop.EntityData.SegmentPath
    nextHop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nextHop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nextHop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nextHop.EntityData.Children = types.NewOrderedMap()
    nextHop.EntityData.Children.Append("next-hop", types.YChild{"NextHop", &nextHop.NextHop})
    nextHop.EntityData.Leafs = types.NewOrderedMap()
    nextHop.EntityData.Leafs.Append("topology-id", types.YLeaf{"TopologyId", nextHop.TopologyId})
    nextHop.EntityData.Leafs.Append("flags", types.YLeaf{"Flags", nextHop.Flags})

    nextHop.EntityData.YListKeys = []string {}

    return &(nextHop.EntityData)
}

// L2rib_EviChildTables_Macs_Mac_Route_Regular_NextHop_NextHop
// Next hop
type L2rib_EviChildTables_Macs_Mac_Route_Regular_NextHop_NextHop struct {
    EntityData types.CommonEntityData
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
    // [a-zA-Z0-9._/-]+.
    InterfaceHandle interface{}

    // XID Next Hop. The type is interface{} with range: 0..4294967295.
    Xid interface{}

    // Labeled Next Hop.
    Labeled L2rib_EviChildTables_Macs_Mac_Route_Regular_NextHop_NextHop_Labeled
}

func (nextHop *L2rib_EviChildTables_Macs_Mac_Route_Regular_NextHop_NextHop) GetEntityData() *types.CommonEntityData {
    nextHop.EntityData.YFilter = nextHop.YFilter
    nextHop.EntityData.YangName = "next-hop"
    nextHop.EntityData.BundleName = "cisco_ios_xr"
    nextHop.EntityData.ParentYangName = "next-hop"
    nextHop.EntityData.SegmentPath = "next-hop"
    nextHop.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/macs/mac/route/regular/next-hop/" + nextHop.EntityData.SegmentPath
    nextHop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nextHop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nextHop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nextHop.EntityData.Children = types.NewOrderedMap()
    nextHop.EntityData.Children.Append("labeled", types.YChild{"Labeled", &nextHop.Labeled})
    nextHop.EntityData.Leafs = types.NewOrderedMap()
    nextHop.EntityData.Leafs.Append("type", types.YLeaf{"Type", nextHop.Type})
    nextHop.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", nextHop.Ipv4})
    nextHop.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", nextHop.Ipv6})
    nextHop.EntityData.Leafs.Append("mac", types.YLeaf{"Mac", nextHop.Mac})
    nextHop.EntityData.Leafs.Append("interface-handle", types.YLeaf{"InterfaceHandle", nextHop.InterfaceHandle})
    nextHop.EntityData.Leafs.Append("xid", types.YLeaf{"Xid", nextHop.Xid})

    nextHop.EntityData.YListKeys = []string {}

    return &(nextHop.EntityData)
}

// L2rib_EviChildTables_Macs_Mac_Route_Regular_NextHop_NextHop_Labeled
// Labeled Next Hop
type L2rib_EviChildTables_Macs_Mac_Route_Regular_NextHop_NextHop_Labeled struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2RIB Address Family. The type is L2ribAfi.
    AddressFamily interface{}

    // IP Address. The type is string.
    IpAddress interface{}

    // Label. The type is interface{} with range: 0..4294967295.
    Label interface{}

    // Internal Label. The type is bool.
    Internal interface{}
}

func (labeled *L2rib_EviChildTables_Macs_Mac_Route_Regular_NextHop_NextHop_Labeled) GetEntityData() *types.CommonEntityData {
    labeled.EntityData.YFilter = labeled.YFilter
    labeled.EntityData.YangName = "labeled"
    labeled.EntityData.BundleName = "cisco_ios_xr"
    labeled.EntityData.ParentYangName = "next-hop"
    labeled.EntityData.SegmentPath = "labeled"
    labeled.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/macs/mac/route/regular/next-hop/next-hop/" + labeled.EntityData.SegmentPath
    labeled.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    labeled.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    labeled.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    labeled.EntityData.Children = types.NewOrderedMap()
    labeled.EntityData.Leafs = types.NewOrderedMap()
    labeled.EntityData.Leafs.Append("address-family", types.YLeaf{"AddressFamily", labeled.AddressFamily})
    labeled.EntityData.Leafs.Append("ip-address", types.YLeaf{"IpAddress", labeled.IpAddress})
    labeled.EntityData.Leafs.Append("label", types.YLeaf{"Label", labeled.Label})
    labeled.EntityData.Leafs.Append("internal", types.YLeaf{"Internal", labeled.Internal})

    labeled.EntityData.YListKeys = []string {}

    return &(labeled.EntityData)
}

// L2rib_EviChildTables_Macs_Mac_Route_EvpnEsi
// EVPN ESI MAC route
type L2rib_EviChildTables_Macs_Mac_Route_EvpnEsi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MAC route sequence number. The type is interface{} with range:
    // 0..4294967295.
    SequenceNumber interface{}

    // Forwarding State. True means forward, False means drop. The type is bool.
    ForwardState interface{}

    // Ethernet Segment Identifier.
    EthernetSegmentId L2rib_EviChildTables_Macs_Mac_Route_EvpnEsi_EthernetSegmentId

    // Path list information. Set for detailed MAC route information.
    PathList L2rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList
}

func (evpnEsi *L2rib_EviChildTables_Macs_Mac_Route_EvpnEsi) GetEntityData() *types.CommonEntityData {
    evpnEsi.EntityData.YFilter = evpnEsi.YFilter
    evpnEsi.EntityData.YangName = "evpn-esi"
    evpnEsi.EntityData.BundleName = "cisco_ios_xr"
    evpnEsi.EntityData.ParentYangName = "route"
    evpnEsi.EntityData.SegmentPath = "evpn-esi"
    evpnEsi.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/macs/mac/route/" + evpnEsi.EntityData.SegmentPath
    evpnEsi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnEsi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnEsi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnEsi.EntityData.Children = types.NewOrderedMap()
    evpnEsi.EntityData.Children.Append("ethernet-segment-id", types.YChild{"EthernetSegmentId", &evpnEsi.EthernetSegmentId})
    evpnEsi.EntityData.Children.Append("path-list", types.YChild{"PathList", &evpnEsi.PathList})
    evpnEsi.EntityData.Leafs = types.NewOrderedMap()
    evpnEsi.EntityData.Leafs.Append("sequence-number", types.YLeaf{"SequenceNumber", evpnEsi.SequenceNumber})
    evpnEsi.EntityData.Leafs.Append("forward-state", types.YLeaf{"ForwardState", evpnEsi.ForwardState})

    evpnEsi.EntityData.YListKeys = []string {}

    return &(evpnEsi.EntityData)
}

// L2rib_EviChildTables_Macs_Mac_Route_EvpnEsi_EthernetSegmentId
// Ethernet Segment Identifier
type L2rib_EviChildTables_Macs_Mac_Route_EvpnEsi_EthernetSegmentId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LACP System Priority. The type is interface{} with range: 0..65535.
    SystemPriority interface{}

    // LACP System Id. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    SystemId interface{}

    // LACP Port Key. The type is interface{} with range: 0..65535.
    PortKey interface{}
}

func (ethernetSegmentId *L2rib_EviChildTables_Macs_Mac_Route_EvpnEsi_EthernetSegmentId) GetEntityData() *types.CommonEntityData {
    ethernetSegmentId.EntityData.YFilter = ethernetSegmentId.YFilter
    ethernetSegmentId.EntityData.YangName = "ethernet-segment-id"
    ethernetSegmentId.EntityData.BundleName = "cisco_ios_xr"
    ethernetSegmentId.EntityData.ParentYangName = "evpn-esi"
    ethernetSegmentId.EntityData.SegmentPath = "ethernet-segment-id"
    ethernetSegmentId.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/macs/mac/route/evpn-esi/" + ethernetSegmentId.EntityData.SegmentPath
    ethernetSegmentId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ethernetSegmentId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ethernetSegmentId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ethernetSegmentId.EntityData.Children = types.NewOrderedMap()
    ethernetSegmentId.EntityData.Leafs = types.NewOrderedMap()
    ethernetSegmentId.EntityData.Leafs.Append("system-priority", types.YLeaf{"SystemPriority", ethernetSegmentId.SystemPriority})
    ethernetSegmentId.EntityData.Leafs.Append("system-id", types.YLeaf{"SystemId", ethernetSegmentId.SystemId})
    ethernetSegmentId.EntityData.Leafs.Append("port-key", types.YLeaf{"PortKey", ethernetSegmentId.PortKey})

    ethernetSegmentId.EntityData.YListKeys = []string {}

    return &(ethernetSegmentId.EntityData)
}

// L2rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList
// Path list information. Set for detailed MAC
// route information
type L2rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ID of EAD route producer. The type is interface{} with range: 0..255.
    ProducerId interface{}

    // Number of MAC routes bound to this Path list. The type is interface{} with
    // range: 0..4294967295.
    MacCount interface{}

    // Path list local Label. The type is interface{} with range: 0..4294967295.
    LocalLabel interface{}

    // Type-specific Path List info.
    PathListInfo L2rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo

    // Array of Next Hops for MAC Path List. The type is slice of
    // L2rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_NextHopArray.
    NextHopArray []*L2rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_NextHopArray
}

func (pathList *L2rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList) GetEntityData() *types.CommonEntityData {
    pathList.EntityData.YFilter = pathList.YFilter
    pathList.EntityData.YangName = "path-list"
    pathList.EntityData.BundleName = "cisco_ios_xr"
    pathList.EntityData.ParentYangName = "evpn-esi"
    pathList.EntityData.SegmentPath = "path-list"
    pathList.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/macs/mac/route/evpn-esi/" + pathList.EntityData.SegmentPath
    pathList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pathList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pathList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pathList.EntityData.Children = types.NewOrderedMap()
    pathList.EntityData.Children.Append("path-list-info", types.YChild{"PathListInfo", &pathList.PathListInfo})
    pathList.EntityData.Children.Append("next-hop-array", types.YChild{"NextHopArray", nil})
    for i := range pathList.NextHopArray {
        types.SetYListKey(pathList.NextHopArray[i], i)
        pathList.EntityData.Children.Append(types.GetSegmentPath(pathList.NextHopArray[i]), types.YChild{"NextHopArray", pathList.NextHopArray[i]})
    }
    pathList.EntityData.Leafs = types.NewOrderedMap()
    pathList.EntityData.Leafs.Append("producer-id", types.YLeaf{"ProducerId", pathList.ProducerId})
    pathList.EntityData.Leafs.Append("mac-count", types.YLeaf{"MacCount", pathList.MacCount})
    pathList.EntityData.Leafs.Append("local-label", types.YLeaf{"LocalLabel", pathList.LocalLabel})

    pathList.EntityData.YListKeys = []string {}

    return &(pathList.EntityData)
}

// L2rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo
// Type-specific Path List info
type L2rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type. The type is L2ribMacRoute.
    Type interface{}

    // ESI Path List.
    PathListEsi L2rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi

    // MAC Path List.
    PathListMac L2rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListMac
}

func (pathListInfo *L2rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo) GetEntityData() *types.CommonEntityData {
    pathListInfo.EntityData.YFilter = pathListInfo.YFilter
    pathListInfo.EntityData.YangName = "path-list-info"
    pathListInfo.EntityData.BundleName = "cisco_ios_xr"
    pathListInfo.EntityData.ParentYangName = "path-list"
    pathListInfo.EntityData.SegmentPath = "path-list-info"
    pathListInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/macs/mac/route/evpn-esi/path-list/" + pathListInfo.EntityData.SegmentPath
    pathListInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pathListInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pathListInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pathListInfo.EntityData.Children = types.NewOrderedMap()
    pathListInfo.EntityData.Children.Append("path-list-esi", types.YChild{"PathListEsi", &pathListInfo.PathListEsi})
    pathListInfo.EntityData.Children.Append("path-list-mac", types.YChild{"PathListMac", &pathListInfo.PathListMac})
    pathListInfo.EntityData.Leafs = types.NewOrderedMap()
    pathListInfo.EntityData.Leafs.Append("type", types.YLeaf{"Type", pathListInfo.Type})

    pathListInfo.EntityData.YListKeys = []string {}

    return &(pathListInfo.EntityData)
}

// L2rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi
// ESI Path List
type L2rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Path list Resolved. The type is bool.
    Resolved interface{}

    // Ethernet Segment Identifier.
    EthernetSegmentId L2rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_EthernetSegmentId

    // Array of Next Hops from MAC Update. The type is slice of
    // L2rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray.
    MacUpdateNextHopArray []*L2rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray
}

func (pathListEsi *L2rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi) GetEntityData() *types.CommonEntityData {
    pathListEsi.EntityData.YFilter = pathListEsi.YFilter
    pathListEsi.EntityData.YangName = "path-list-esi"
    pathListEsi.EntityData.BundleName = "cisco_ios_xr"
    pathListEsi.EntityData.ParentYangName = "path-list-info"
    pathListEsi.EntityData.SegmentPath = "path-list-esi"
    pathListEsi.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/macs/mac/route/evpn-esi/path-list/path-list-info/" + pathListEsi.EntityData.SegmentPath
    pathListEsi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pathListEsi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pathListEsi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pathListEsi.EntityData.Children = types.NewOrderedMap()
    pathListEsi.EntityData.Children.Append("ethernet-segment-id", types.YChild{"EthernetSegmentId", &pathListEsi.EthernetSegmentId})
    pathListEsi.EntityData.Children.Append("mac-update-next-hop-array", types.YChild{"MacUpdateNextHopArray", nil})
    for i := range pathListEsi.MacUpdateNextHopArray {
        types.SetYListKey(pathListEsi.MacUpdateNextHopArray[i], i)
        pathListEsi.EntityData.Children.Append(types.GetSegmentPath(pathListEsi.MacUpdateNextHopArray[i]), types.YChild{"MacUpdateNextHopArray", pathListEsi.MacUpdateNextHopArray[i]})
    }
    pathListEsi.EntityData.Leafs = types.NewOrderedMap()
    pathListEsi.EntityData.Leafs.Append("resolved", types.YLeaf{"Resolved", pathListEsi.Resolved})

    pathListEsi.EntityData.YListKeys = []string {}

    return &(pathListEsi.EntityData)
}

// L2rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_EthernetSegmentId
// Ethernet Segment Identifier
type L2rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_EthernetSegmentId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LACP System Priority. The type is interface{} with range: 0..65535.
    SystemPriority interface{}

    // LACP System Id. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    SystemId interface{}

    // LACP Port Key. The type is interface{} with range: 0..65535.
    PortKey interface{}
}

func (ethernetSegmentId *L2rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_EthernetSegmentId) GetEntityData() *types.CommonEntityData {
    ethernetSegmentId.EntityData.YFilter = ethernetSegmentId.YFilter
    ethernetSegmentId.EntityData.YangName = "ethernet-segment-id"
    ethernetSegmentId.EntityData.BundleName = "cisco_ios_xr"
    ethernetSegmentId.EntityData.ParentYangName = "path-list-esi"
    ethernetSegmentId.EntityData.SegmentPath = "ethernet-segment-id"
    ethernetSegmentId.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/macs/mac/route/evpn-esi/path-list/path-list-info/path-list-esi/" + ethernetSegmentId.EntityData.SegmentPath
    ethernetSegmentId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ethernetSegmentId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ethernetSegmentId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ethernetSegmentId.EntityData.Children = types.NewOrderedMap()
    ethernetSegmentId.EntityData.Leafs = types.NewOrderedMap()
    ethernetSegmentId.EntityData.Leafs.Append("system-priority", types.YLeaf{"SystemPriority", ethernetSegmentId.SystemPriority})
    ethernetSegmentId.EntityData.Leafs.Append("system-id", types.YLeaf{"SystemId", ethernetSegmentId.SystemId})
    ethernetSegmentId.EntityData.Leafs.Append("port-key", types.YLeaf{"PortKey", ethernetSegmentId.PortKey})

    ethernetSegmentId.EntityData.YListKeys = []string {}

    return &(ethernetSegmentId.EntityData)
}

// L2rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray
// Array of Next Hops from MAC Update
type L2rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Next-hop TOPOLOGY ID. The type is interface{} with range: 0..4294967295.
    TopologyId interface{}

    // Next-hop flags. The type is interface{} with range: 0..65535.
    Flags interface{}

    // Next hop.
    NextHop L2rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop
}

func (macUpdateNextHopArray *L2rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray) GetEntityData() *types.CommonEntityData {
    macUpdateNextHopArray.EntityData.YFilter = macUpdateNextHopArray.YFilter
    macUpdateNextHopArray.EntityData.YangName = "mac-update-next-hop-array"
    macUpdateNextHopArray.EntityData.BundleName = "cisco_ios_xr"
    macUpdateNextHopArray.EntityData.ParentYangName = "path-list-esi"
    macUpdateNextHopArray.EntityData.SegmentPath = "mac-update-next-hop-array" + types.AddNoKeyToken(macUpdateNextHopArray)
    macUpdateNextHopArray.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/macs/mac/route/evpn-esi/path-list/path-list-info/path-list-esi/" + macUpdateNextHopArray.EntityData.SegmentPath
    macUpdateNextHopArray.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macUpdateNextHopArray.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macUpdateNextHopArray.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macUpdateNextHopArray.EntityData.Children = types.NewOrderedMap()
    macUpdateNextHopArray.EntityData.Children.Append("next-hop", types.YChild{"NextHop", &macUpdateNextHopArray.NextHop})
    macUpdateNextHopArray.EntityData.Leafs = types.NewOrderedMap()
    macUpdateNextHopArray.EntityData.Leafs.Append("topology-id", types.YLeaf{"TopologyId", macUpdateNextHopArray.TopologyId})
    macUpdateNextHopArray.EntityData.Leafs.Append("flags", types.YLeaf{"Flags", macUpdateNextHopArray.Flags})

    macUpdateNextHopArray.EntityData.YListKeys = []string {}

    return &(macUpdateNextHopArray.EntityData)
}

// L2rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop
// Next hop
type L2rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop struct {
    EntityData types.CommonEntityData
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
    // [a-zA-Z0-9._/-]+.
    InterfaceHandle interface{}

    // XID Next Hop. The type is interface{} with range: 0..4294967295.
    Xid interface{}

    // Labeled Next Hop.
    Labeled L2rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled
}

func (nextHop *L2rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop) GetEntityData() *types.CommonEntityData {
    nextHop.EntityData.YFilter = nextHop.YFilter
    nextHop.EntityData.YangName = "next-hop"
    nextHop.EntityData.BundleName = "cisco_ios_xr"
    nextHop.EntityData.ParentYangName = "mac-update-next-hop-array"
    nextHop.EntityData.SegmentPath = "next-hop"
    nextHop.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/macs/mac/route/evpn-esi/path-list/path-list-info/path-list-esi/mac-update-next-hop-array/" + nextHop.EntityData.SegmentPath
    nextHop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nextHop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nextHop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nextHop.EntityData.Children = types.NewOrderedMap()
    nextHop.EntityData.Children.Append("labeled", types.YChild{"Labeled", &nextHop.Labeled})
    nextHop.EntityData.Leafs = types.NewOrderedMap()
    nextHop.EntityData.Leafs.Append("type", types.YLeaf{"Type", nextHop.Type})
    nextHop.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", nextHop.Ipv4})
    nextHop.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", nextHop.Ipv6})
    nextHop.EntityData.Leafs.Append("mac", types.YLeaf{"Mac", nextHop.Mac})
    nextHop.EntityData.Leafs.Append("interface-handle", types.YLeaf{"InterfaceHandle", nextHop.InterfaceHandle})
    nextHop.EntityData.Leafs.Append("xid", types.YLeaf{"Xid", nextHop.Xid})

    nextHop.EntityData.YListKeys = []string {}

    return &(nextHop.EntityData)
}

// L2rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled
// Labeled Next Hop
type L2rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2RIB Address Family. The type is L2ribAfi.
    AddressFamily interface{}

    // IP Address. The type is string.
    IpAddress interface{}

    // Label. The type is interface{} with range: 0..4294967295.
    Label interface{}

    // Internal Label. The type is bool.
    Internal interface{}
}

func (labeled *L2rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled) GetEntityData() *types.CommonEntityData {
    labeled.EntityData.YFilter = labeled.YFilter
    labeled.EntityData.YangName = "labeled"
    labeled.EntityData.BundleName = "cisco_ios_xr"
    labeled.EntityData.ParentYangName = "next-hop"
    labeled.EntityData.SegmentPath = "labeled"
    labeled.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/macs/mac/route/evpn-esi/path-list/path-list-info/path-list-esi/mac-update-next-hop-array/next-hop/" + labeled.EntityData.SegmentPath
    labeled.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    labeled.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    labeled.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    labeled.EntityData.Children = types.NewOrderedMap()
    labeled.EntityData.Leafs = types.NewOrderedMap()
    labeled.EntityData.Leafs.Append("address-family", types.YLeaf{"AddressFamily", labeled.AddressFamily})
    labeled.EntityData.Leafs.Append("ip-address", types.YLeaf{"IpAddress", labeled.IpAddress})
    labeled.EntityData.Leafs.Append("label", types.YLeaf{"Label", labeled.Label})
    labeled.EntityData.Leafs.Append("internal", types.YLeaf{"Internal", labeled.Internal})

    labeled.EntityData.YListKeys = []string {}

    return &(labeled.EntityData)
}

// L2rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListMac
// MAC Path List
type L2rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListMac struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MAC Address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddress interface{}
}

func (pathListMac *L2rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_PathListInfo_PathListMac) GetEntityData() *types.CommonEntityData {
    pathListMac.EntityData.YFilter = pathListMac.YFilter
    pathListMac.EntityData.YangName = "path-list-mac"
    pathListMac.EntityData.BundleName = "cisco_ios_xr"
    pathListMac.EntityData.ParentYangName = "path-list-info"
    pathListMac.EntityData.SegmentPath = "path-list-mac"
    pathListMac.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/macs/mac/route/evpn-esi/path-list/path-list-info/" + pathListMac.EntityData.SegmentPath
    pathListMac.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pathListMac.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pathListMac.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pathListMac.EntityData.Children = types.NewOrderedMap()
    pathListMac.EntityData.Leafs = types.NewOrderedMap()
    pathListMac.EntityData.Leafs.Append("mac-address", types.YLeaf{"MacAddress", pathListMac.MacAddress})

    pathListMac.EntityData.YListKeys = []string {}

    return &(pathListMac.EntityData)
}

// L2rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_NextHopArray
// Array of Next Hops for MAC Path List
type L2rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_NextHopArray struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Next-hop TOPOLOGY ID. The type is interface{} with range: 0..4294967295.
    TopologyId interface{}

    // Next-hop flags. The type is interface{} with range: 0..65535.
    Flags interface{}

    // Next hop.
    NextHop L2rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_NextHopArray_NextHop
}

func (nextHopArray *L2rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_NextHopArray) GetEntityData() *types.CommonEntityData {
    nextHopArray.EntityData.YFilter = nextHopArray.YFilter
    nextHopArray.EntityData.YangName = "next-hop-array"
    nextHopArray.EntityData.BundleName = "cisco_ios_xr"
    nextHopArray.EntityData.ParentYangName = "path-list"
    nextHopArray.EntityData.SegmentPath = "next-hop-array" + types.AddNoKeyToken(nextHopArray)
    nextHopArray.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/macs/mac/route/evpn-esi/path-list/" + nextHopArray.EntityData.SegmentPath
    nextHopArray.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nextHopArray.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nextHopArray.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nextHopArray.EntityData.Children = types.NewOrderedMap()
    nextHopArray.EntityData.Children.Append("next-hop", types.YChild{"NextHop", &nextHopArray.NextHop})
    nextHopArray.EntityData.Leafs = types.NewOrderedMap()
    nextHopArray.EntityData.Leafs.Append("topology-id", types.YLeaf{"TopologyId", nextHopArray.TopologyId})
    nextHopArray.EntityData.Leafs.Append("flags", types.YLeaf{"Flags", nextHopArray.Flags})

    nextHopArray.EntityData.YListKeys = []string {}

    return &(nextHopArray.EntityData)
}

// L2rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_NextHopArray_NextHop
// Next hop
type L2rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_NextHopArray_NextHop struct {
    EntityData types.CommonEntityData
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
    // [a-zA-Z0-9._/-]+.
    InterfaceHandle interface{}

    // XID Next Hop. The type is interface{} with range: 0..4294967295.
    Xid interface{}

    // Labeled Next Hop.
    Labeled L2rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_NextHopArray_NextHop_Labeled
}

func (nextHop *L2rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_NextHopArray_NextHop) GetEntityData() *types.CommonEntityData {
    nextHop.EntityData.YFilter = nextHop.YFilter
    nextHop.EntityData.YangName = "next-hop"
    nextHop.EntityData.BundleName = "cisco_ios_xr"
    nextHop.EntityData.ParentYangName = "next-hop-array"
    nextHop.EntityData.SegmentPath = "next-hop"
    nextHop.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/macs/mac/route/evpn-esi/path-list/next-hop-array/" + nextHop.EntityData.SegmentPath
    nextHop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nextHop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nextHop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nextHop.EntityData.Children = types.NewOrderedMap()
    nextHop.EntityData.Children.Append("labeled", types.YChild{"Labeled", &nextHop.Labeled})
    nextHop.EntityData.Leafs = types.NewOrderedMap()
    nextHop.EntityData.Leafs.Append("type", types.YLeaf{"Type", nextHop.Type})
    nextHop.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", nextHop.Ipv4})
    nextHop.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", nextHop.Ipv6})
    nextHop.EntityData.Leafs.Append("mac", types.YLeaf{"Mac", nextHop.Mac})
    nextHop.EntityData.Leafs.Append("interface-handle", types.YLeaf{"InterfaceHandle", nextHop.InterfaceHandle})
    nextHop.EntityData.Leafs.Append("xid", types.YLeaf{"Xid", nextHop.Xid})

    nextHop.EntityData.YListKeys = []string {}

    return &(nextHop.EntityData)
}

// L2rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_NextHopArray_NextHop_Labeled
// Labeled Next Hop
type L2rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_NextHopArray_NextHop_Labeled struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2RIB Address Family. The type is L2ribAfi.
    AddressFamily interface{}

    // IP Address. The type is string.
    IpAddress interface{}

    // Label. The type is interface{} with range: 0..4294967295.
    Label interface{}

    // Internal Label. The type is bool.
    Internal interface{}
}

func (labeled *L2rib_EviChildTables_Macs_Mac_Route_EvpnEsi_PathList_NextHopArray_NextHop_Labeled) GetEntityData() *types.CommonEntityData {
    labeled.EntityData.YFilter = labeled.YFilter
    labeled.EntityData.YangName = "labeled"
    labeled.EntityData.BundleName = "cisco_ios_xr"
    labeled.EntityData.ParentYangName = "next-hop"
    labeled.EntityData.SegmentPath = "labeled"
    labeled.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/macs/mac/route/evpn-esi/path-list/next-hop-array/next-hop/" + labeled.EntityData.SegmentPath
    labeled.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    labeled.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    labeled.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    labeled.EntityData.Children = types.NewOrderedMap()
    labeled.EntityData.Leafs = types.NewOrderedMap()
    labeled.EntityData.Leafs.Append("address-family", types.YLeaf{"AddressFamily", labeled.AddressFamily})
    labeled.EntityData.Leafs.Append("ip-address", types.YLeaf{"IpAddress", labeled.IpAddress})
    labeled.EntityData.Leafs.Append("label", types.YLeaf{"Label", labeled.Label})
    labeled.EntityData.Leafs.Append("internal", types.YLeaf{"Internal", labeled.Internal})

    labeled.EntityData.YListKeys = []string {}

    return &(labeled.EntityData)
}

// L2rib_EviChildTables_Macs_Mac_Route_Bmac
// BMAC route
type L2rib_EviChildTables_Macs_Mac_Route_Bmac struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BMAC Address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    BmacAddress interface{}

    // Forwarding State. True means forward, False means drop. The type is bool.
    ForwardState interface{}

    // Path list information.
    PathList L2rib_EviChildTables_Macs_Mac_Route_Bmac_PathList
}

func (bmac *L2rib_EviChildTables_Macs_Mac_Route_Bmac) GetEntityData() *types.CommonEntityData {
    bmac.EntityData.YFilter = bmac.YFilter
    bmac.EntityData.YangName = "bmac"
    bmac.EntityData.BundleName = "cisco_ios_xr"
    bmac.EntityData.ParentYangName = "route"
    bmac.EntityData.SegmentPath = "bmac"
    bmac.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/macs/mac/route/" + bmac.EntityData.SegmentPath
    bmac.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bmac.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bmac.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bmac.EntityData.Children = types.NewOrderedMap()
    bmac.EntityData.Children.Append("path-list", types.YChild{"PathList", &bmac.PathList})
    bmac.EntityData.Leafs = types.NewOrderedMap()
    bmac.EntityData.Leafs.Append("bmac-address", types.YLeaf{"BmacAddress", bmac.BmacAddress})
    bmac.EntityData.Leafs.Append("forward-state", types.YLeaf{"ForwardState", bmac.ForwardState})

    bmac.EntityData.YListKeys = []string {}

    return &(bmac.EntityData)
}

// L2rib_EviChildTables_Macs_Mac_Route_Bmac_PathList
// Path list information
type L2rib_EviChildTables_Macs_Mac_Route_Bmac_PathList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ID of EAD route producer. The type is interface{} with range: 0..255.
    ProducerId interface{}

    // Number of MAC routes bound to this Path list. The type is interface{} with
    // range: 0..4294967295.
    MacCount interface{}

    // Path list local Label. The type is interface{} with range: 0..4294967295.
    LocalLabel interface{}

    // Type-specific Path List info.
    PathListInfo L2rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo

    // Array of Next Hops for MAC Path List. The type is slice of
    // L2rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_NextHopArray.
    NextHopArray []*L2rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_NextHopArray
}

func (pathList *L2rib_EviChildTables_Macs_Mac_Route_Bmac_PathList) GetEntityData() *types.CommonEntityData {
    pathList.EntityData.YFilter = pathList.YFilter
    pathList.EntityData.YangName = "path-list"
    pathList.EntityData.BundleName = "cisco_ios_xr"
    pathList.EntityData.ParentYangName = "bmac"
    pathList.EntityData.SegmentPath = "path-list"
    pathList.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/macs/mac/route/bmac/" + pathList.EntityData.SegmentPath
    pathList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pathList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pathList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pathList.EntityData.Children = types.NewOrderedMap()
    pathList.EntityData.Children.Append("path-list-info", types.YChild{"PathListInfo", &pathList.PathListInfo})
    pathList.EntityData.Children.Append("next-hop-array", types.YChild{"NextHopArray", nil})
    for i := range pathList.NextHopArray {
        types.SetYListKey(pathList.NextHopArray[i], i)
        pathList.EntityData.Children.Append(types.GetSegmentPath(pathList.NextHopArray[i]), types.YChild{"NextHopArray", pathList.NextHopArray[i]})
    }
    pathList.EntityData.Leafs = types.NewOrderedMap()
    pathList.EntityData.Leafs.Append("producer-id", types.YLeaf{"ProducerId", pathList.ProducerId})
    pathList.EntityData.Leafs.Append("mac-count", types.YLeaf{"MacCount", pathList.MacCount})
    pathList.EntityData.Leafs.Append("local-label", types.YLeaf{"LocalLabel", pathList.LocalLabel})

    pathList.EntityData.YListKeys = []string {}

    return &(pathList.EntityData)
}

// L2rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo
// Type-specific Path List info
type L2rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type. The type is L2ribMacRoute.
    Type interface{}

    // ESI Path List.
    PathListEsi L2rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi

    // MAC Path List.
    PathListMac L2rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListMac
}

func (pathListInfo *L2rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo) GetEntityData() *types.CommonEntityData {
    pathListInfo.EntityData.YFilter = pathListInfo.YFilter
    pathListInfo.EntityData.YangName = "path-list-info"
    pathListInfo.EntityData.BundleName = "cisco_ios_xr"
    pathListInfo.EntityData.ParentYangName = "path-list"
    pathListInfo.EntityData.SegmentPath = "path-list-info"
    pathListInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/macs/mac/route/bmac/path-list/" + pathListInfo.EntityData.SegmentPath
    pathListInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pathListInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pathListInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pathListInfo.EntityData.Children = types.NewOrderedMap()
    pathListInfo.EntityData.Children.Append("path-list-esi", types.YChild{"PathListEsi", &pathListInfo.PathListEsi})
    pathListInfo.EntityData.Children.Append("path-list-mac", types.YChild{"PathListMac", &pathListInfo.PathListMac})
    pathListInfo.EntityData.Leafs = types.NewOrderedMap()
    pathListInfo.EntityData.Leafs.Append("type", types.YLeaf{"Type", pathListInfo.Type})

    pathListInfo.EntityData.YListKeys = []string {}

    return &(pathListInfo.EntityData)
}

// L2rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi
// ESI Path List
type L2rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Path list Resolved. The type is bool.
    Resolved interface{}

    // Ethernet Segment Identifier.
    EthernetSegmentId L2rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_EthernetSegmentId

    // Array of Next Hops from MAC Update. The type is slice of
    // L2rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray.
    MacUpdateNextHopArray []*L2rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray
}

func (pathListEsi *L2rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi) GetEntityData() *types.CommonEntityData {
    pathListEsi.EntityData.YFilter = pathListEsi.YFilter
    pathListEsi.EntityData.YangName = "path-list-esi"
    pathListEsi.EntityData.BundleName = "cisco_ios_xr"
    pathListEsi.EntityData.ParentYangName = "path-list-info"
    pathListEsi.EntityData.SegmentPath = "path-list-esi"
    pathListEsi.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/macs/mac/route/bmac/path-list/path-list-info/" + pathListEsi.EntityData.SegmentPath
    pathListEsi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pathListEsi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pathListEsi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pathListEsi.EntityData.Children = types.NewOrderedMap()
    pathListEsi.EntityData.Children.Append("ethernet-segment-id", types.YChild{"EthernetSegmentId", &pathListEsi.EthernetSegmentId})
    pathListEsi.EntityData.Children.Append("mac-update-next-hop-array", types.YChild{"MacUpdateNextHopArray", nil})
    for i := range pathListEsi.MacUpdateNextHopArray {
        types.SetYListKey(pathListEsi.MacUpdateNextHopArray[i], i)
        pathListEsi.EntityData.Children.Append(types.GetSegmentPath(pathListEsi.MacUpdateNextHopArray[i]), types.YChild{"MacUpdateNextHopArray", pathListEsi.MacUpdateNextHopArray[i]})
    }
    pathListEsi.EntityData.Leafs = types.NewOrderedMap()
    pathListEsi.EntityData.Leafs.Append("resolved", types.YLeaf{"Resolved", pathListEsi.Resolved})

    pathListEsi.EntityData.YListKeys = []string {}

    return &(pathListEsi.EntityData)
}

// L2rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_EthernetSegmentId
// Ethernet Segment Identifier
type L2rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_EthernetSegmentId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LACP System Priority. The type is interface{} with range: 0..65535.
    SystemPriority interface{}

    // LACP System Id. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    SystemId interface{}

    // LACP Port Key. The type is interface{} with range: 0..65535.
    PortKey interface{}
}

func (ethernetSegmentId *L2rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_EthernetSegmentId) GetEntityData() *types.CommonEntityData {
    ethernetSegmentId.EntityData.YFilter = ethernetSegmentId.YFilter
    ethernetSegmentId.EntityData.YangName = "ethernet-segment-id"
    ethernetSegmentId.EntityData.BundleName = "cisco_ios_xr"
    ethernetSegmentId.EntityData.ParentYangName = "path-list-esi"
    ethernetSegmentId.EntityData.SegmentPath = "ethernet-segment-id"
    ethernetSegmentId.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/macs/mac/route/bmac/path-list/path-list-info/path-list-esi/" + ethernetSegmentId.EntityData.SegmentPath
    ethernetSegmentId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ethernetSegmentId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ethernetSegmentId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ethernetSegmentId.EntityData.Children = types.NewOrderedMap()
    ethernetSegmentId.EntityData.Leafs = types.NewOrderedMap()
    ethernetSegmentId.EntityData.Leafs.Append("system-priority", types.YLeaf{"SystemPriority", ethernetSegmentId.SystemPriority})
    ethernetSegmentId.EntityData.Leafs.Append("system-id", types.YLeaf{"SystemId", ethernetSegmentId.SystemId})
    ethernetSegmentId.EntityData.Leafs.Append("port-key", types.YLeaf{"PortKey", ethernetSegmentId.PortKey})

    ethernetSegmentId.EntityData.YListKeys = []string {}

    return &(ethernetSegmentId.EntityData)
}

// L2rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray
// Array of Next Hops from MAC Update
type L2rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Next-hop TOPOLOGY ID. The type is interface{} with range: 0..4294967295.
    TopologyId interface{}

    // Next-hop flags. The type is interface{} with range: 0..65535.
    Flags interface{}

    // Next hop.
    NextHop L2rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop
}

func (macUpdateNextHopArray *L2rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray) GetEntityData() *types.CommonEntityData {
    macUpdateNextHopArray.EntityData.YFilter = macUpdateNextHopArray.YFilter
    macUpdateNextHopArray.EntityData.YangName = "mac-update-next-hop-array"
    macUpdateNextHopArray.EntityData.BundleName = "cisco_ios_xr"
    macUpdateNextHopArray.EntityData.ParentYangName = "path-list-esi"
    macUpdateNextHopArray.EntityData.SegmentPath = "mac-update-next-hop-array" + types.AddNoKeyToken(macUpdateNextHopArray)
    macUpdateNextHopArray.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/macs/mac/route/bmac/path-list/path-list-info/path-list-esi/" + macUpdateNextHopArray.EntityData.SegmentPath
    macUpdateNextHopArray.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macUpdateNextHopArray.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macUpdateNextHopArray.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macUpdateNextHopArray.EntityData.Children = types.NewOrderedMap()
    macUpdateNextHopArray.EntityData.Children.Append("next-hop", types.YChild{"NextHop", &macUpdateNextHopArray.NextHop})
    macUpdateNextHopArray.EntityData.Leafs = types.NewOrderedMap()
    macUpdateNextHopArray.EntityData.Leafs.Append("topology-id", types.YLeaf{"TopologyId", macUpdateNextHopArray.TopologyId})
    macUpdateNextHopArray.EntityData.Leafs.Append("flags", types.YLeaf{"Flags", macUpdateNextHopArray.Flags})

    macUpdateNextHopArray.EntityData.YListKeys = []string {}

    return &(macUpdateNextHopArray.EntityData)
}

// L2rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop
// Next hop
type L2rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop struct {
    EntityData types.CommonEntityData
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
    // [a-zA-Z0-9._/-]+.
    InterfaceHandle interface{}

    // XID Next Hop. The type is interface{} with range: 0..4294967295.
    Xid interface{}

    // Labeled Next Hop.
    Labeled L2rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled
}

func (nextHop *L2rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop) GetEntityData() *types.CommonEntityData {
    nextHop.EntityData.YFilter = nextHop.YFilter
    nextHop.EntityData.YangName = "next-hop"
    nextHop.EntityData.BundleName = "cisco_ios_xr"
    nextHop.EntityData.ParentYangName = "mac-update-next-hop-array"
    nextHop.EntityData.SegmentPath = "next-hop"
    nextHop.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/macs/mac/route/bmac/path-list/path-list-info/path-list-esi/mac-update-next-hop-array/" + nextHop.EntityData.SegmentPath
    nextHop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nextHop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nextHop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nextHop.EntityData.Children = types.NewOrderedMap()
    nextHop.EntityData.Children.Append("labeled", types.YChild{"Labeled", &nextHop.Labeled})
    nextHop.EntityData.Leafs = types.NewOrderedMap()
    nextHop.EntityData.Leafs.Append("type", types.YLeaf{"Type", nextHop.Type})
    nextHop.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", nextHop.Ipv4})
    nextHop.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", nextHop.Ipv6})
    nextHop.EntityData.Leafs.Append("mac", types.YLeaf{"Mac", nextHop.Mac})
    nextHop.EntityData.Leafs.Append("interface-handle", types.YLeaf{"InterfaceHandle", nextHop.InterfaceHandle})
    nextHop.EntityData.Leafs.Append("xid", types.YLeaf{"Xid", nextHop.Xid})

    nextHop.EntityData.YListKeys = []string {}

    return &(nextHop.EntityData)
}

// L2rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled
// Labeled Next Hop
type L2rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2RIB Address Family. The type is L2ribAfi.
    AddressFamily interface{}

    // IP Address. The type is string.
    IpAddress interface{}

    // Label. The type is interface{} with range: 0..4294967295.
    Label interface{}

    // Internal Label. The type is bool.
    Internal interface{}
}

func (labeled *L2rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled) GetEntityData() *types.CommonEntityData {
    labeled.EntityData.YFilter = labeled.YFilter
    labeled.EntityData.YangName = "labeled"
    labeled.EntityData.BundleName = "cisco_ios_xr"
    labeled.EntityData.ParentYangName = "next-hop"
    labeled.EntityData.SegmentPath = "labeled"
    labeled.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/macs/mac/route/bmac/path-list/path-list-info/path-list-esi/mac-update-next-hop-array/next-hop/" + labeled.EntityData.SegmentPath
    labeled.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    labeled.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    labeled.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    labeled.EntityData.Children = types.NewOrderedMap()
    labeled.EntityData.Leafs = types.NewOrderedMap()
    labeled.EntityData.Leafs.Append("address-family", types.YLeaf{"AddressFamily", labeled.AddressFamily})
    labeled.EntityData.Leafs.Append("ip-address", types.YLeaf{"IpAddress", labeled.IpAddress})
    labeled.EntityData.Leafs.Append("label", types.YLeaf{"Label", labeled.Label})
    labeled.EntityData.Leafs.Append("internal", types.YLeaf{"Internal", labeled.Internal})

    labeled.EntityData.YListKeys = []string {}

    return &(labeled.EntityData)
}

// L2rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListMac
// MAC Path List
type L2rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListMac struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MAC Address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddress interface{}
}

func (pathListMac *L2rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_PathListInfo_PathListMac) GetEntityData() *types.CommonEntityData {
    pathListMac.EntityData.YFilter = pathListMac.YFilter
    pathListMac.EntityData.YangName = "path-list-mac"
    pathListMac.EntityData.BundleName = "cisco_ios_xr"
    pathListMac.EntityData.ParentYangName = "path-list-info"
    pathListMac.EntityData.SegmentPath = "path-list-mac"
    pathListMac.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/macs/mac/route/bmac/path-list/path-list-info/" + pathListMac.EntityData.SegmentPath
    pathListMac.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pathListMac.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pathListMac.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pathListMac.EntityData.Children = types.NewOrderedMap()
    pathListMac.EntityData.Leafs = types.NewOrderedMap()
    pathListMac.EntityData.Leafs.Append("mac-address", types.YLeaf{"MacAddress", pathListMac.MacAddress})

    pathListMac.EntityData.YListKeys = []string {}

    return &(pathListMac.EntityData)
}

// L2rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_NextHopArray
// Array of Next Hops for MAC Path List
type L2rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_NextHopArray struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Next-hop TOPOLOGY ID. The type is interface{} with range: 0..4294967295.
    TopologyId interface{}

    // Next-hop flags. The type is interface{} with range: 0..65535.
    Flags interface{}

    // Next hop.
    NextHop L2rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_NextHopArray_NextHop
}

func (nextHopArray *L2rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_NextHopArray) GetEntityData() *types.CommonEntityData {
    nextHopArray.EntityData.YFilter = nextHopArray.YFilter
    nextHopArray.EntityData.YangName = "next-hop-array"
    nextHopArray.EntityData.BundleName = "cisco_ios_xr"
    nextHopArray.EntityData.ParentYangName = "path-list"
    nextHopArray.EntityData.SegmentPath = "next-hop-array" + types.AddNoKeyToken(nextHopArray)
    nextHopArray.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/macs/mac/route/bmac/path-list/" + nextHopArray.EntityData.SegmentPath
    nextHopArray.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nextHopArray.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nextHopArray.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nextHopArray.EntityData.Children = types.NewOrderedMap()
    nextHopArray.EntityData.Children.Append("next-hop", types.YChild{"NextHop", &nextHopArray.NextHop})
    nextHopArray.EntityData.Leafs = types.NewOrderedMap()
    nextHopArray.EntityData.Leafs.Append("topology-id", types.YLeaf{"TopologyId", nextHopArray.TopologyId})
    nextHopArray.EntityData.Leafs.Append("flags", types.YLeaf{"Flags", nextHopArray.Flags})

    nextHopArray.EntityData.YListKeys = []string {}

    return &(nextHopArray.EntityData)
}

// L2rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_NextHopArray_NextHop
// Next hop
type L2rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_NextHopArray_NextHop struct {
    EntityData types.CommonEntityData
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
    // [a-zA-Z0-9._/-]+.
    InterfaceHandle interface{}

    // XID Next Hop. The type is interface{} with range: 0..4294967295.
    Xid interface{}

    // Labeled Next Hop.
    Labeled L2rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_NextHopArray_NextHop_Labeled
}

func (nextHop *L2rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_NextHopArray_NextHop) GetEntityData() *types.CommonEntityData {
    nextHop.EntityData.YFilter = nextHop.YFilter
    nextHop.EntityData.YangName = "next-hop"
    nextHop.EntityData.BundleName = "cisco_ios_xr"
    nextHop.EntityData.ParentYangName = "next-hop-array"
    nextHop.EntityData.SegmentPath = "next-hop"
    nextHop.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/macs/mac/route/bmac/path-list/next-hop-array/" + nextHop.EntityData.SegmentPath
    nextHop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nextHop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nextHop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nextHop.EntityData.Children = types.NewOrderedMap()
    nextHop.EntityData.Children.Append("labeled", types.YChild{"Labeled", &nextHop.Labeled})
    nextHop.EntityData.Leafs = types.NewOrderedMap()
    nextHop.EntityData.Leafs.Append("type", types.YLeaf{"Type", nextHop.Type})
    nextHop.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", nextHop.Ipv4})
    nextHop.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", nextHop.Ipv6})
    nextHop.EntityData.Leafs.Append("mac", types.YLeaf{"Mac", nextHop.Mac})
    nextHop.EntityData.Leafs.Append("interface-handle", types.YLeaf{"InterfaceHandle", nextHop.InterfaceHandle})
    nextHop.EntityData.Leafs.Append("xid", types.YLeaf{"Xid", nextHop.Xid})

    nextHop.EntityData.YListKeys = []string {}

    return &(nextHop.EntityData)
}

// L2rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_NextHopArray_NextHop_Labeled
// Labeled Next Hop
type L2rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_NextHopArray_NextHop_Labeled struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2RIB Address Family. The type is L2ribAfi.
    AddressFamily interface{}

    // IP Address. The type is string.
    IpAddress interface{}

    // Label. The type is interface{} with range: 0..4294967295.
    Label interface{}

    // Internal Label. The type is bool.
    Internal interface{}
}

func (labeled *L2rib_EviChildTables_Macs_Mac_Route_Bmac_PathList_NextHopArray_NextHop_Labeled) GetEntityData() *types.CommonEntityData {
    labeled.EntityData.YFilter = labeled.YFilter
    labeled.EntityData.YangName = "labeled"
    labeled.EntityData.BundleName = "cisco_ios_xr"
    labeled.EntityData.ParentYangName = "next-hop"
    labeled.EntityData.SegmentPath = "labeled"
    labeled.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/macs/mac/route/bmac/path-list/next-hop-array/next-hop/" + labeled.EntityData.SegmentPath
    labeled.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    labeled.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    labeled.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    labeled.EntityData.Children = types.NewOrderedMap()
    labeled.EntityData.Leafs = types.NewOrderedMap()
    labeled.EntityData.Leafs.Append("address-family", types.YLeaf{"AddressFamily", labeled.AddressFamily})
    labeled.EntityData.Leafs.Append("ip-address", types.YLeaf{"IpAddress", labeled.IpAddress})
    labeled.EntityData.Leafs.Append("label", types.YLeaf{"Label", labeled.Label})
    labeled.EntityData.Leafs.Append("internal", types.YLeaf{"Internal", labeled.Internal})

    labeled.EntityData.YListKeys = []string {}

    return &(labeled.EntityData)
}

// L2rib_EviChildTables_Imets
// L2RIB EVPN EVI IMET table
type L2rib_EviChildTables_Imets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2RIB EVPN IMET table. The type is slice of
    // L2rib_EviChildTables_Imets_Imet.
    Imet []*L2rib_EviChildTables_Imets_Imet
}

func (imets *L2rib_EviChildTables_Imets) GetEntityData() *types.CommonEntityData {
    imets.EntityData.YFilter = imets.YFilter
    imets.EntityData.YangName = "imets"
    imets.EntityData.BundleName = "cisco_ios_xr"
    imets.EntityData.ParentYangName = "evi-child-tables"
    imets.EntityData.SegmentPath = "imets"
    imets.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/" + imets.EntityData.SegmentPath
    imets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    imets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    imets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    imets.EntityData.Children = types.NewOrderedMap()
    imets.EntityData.Children.Append("imet", types.YChild{"Imet", nil})
    for i := range imets.Imet {
        types.SetYListKey(imets.Imet[i], i)
        imets.EntityData.Children.Append(types.GetSegmentPath(imets.Imet[i]), types.YChild{"Imet", imets.Imet[i]})
    }
    imets.EntityData.Leafs = types.NewOrderedMap()

    imets.EntityData.YListKeys = []string {}

    return &(imets.EntityData)
}

// L2rib_EviChildTables_Imets_Imet
// L2RIB EVPN IMET table
type L2rib_EviChildTables_Imets_Imet struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // EVPN ID. The type is interface{} with range: 0..4294967295.
    Evi interface{}

    // Tag ID. The type is interface{} with range: 0..4294967295.
    TagId interface{}

    // Originating Router IP Address. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IpAddr interface{}

    // Admin distance. The type is interface{} with range: 0..4294967295.
    AdminDist interface{}

    // Producer ID. The type is interface{} with range: 0..4294967295.
    ProdId interface{}

    // Originating Router IP Address. The type is string.
    VtepiPaddr interface{}

    // Admin Distance. The type is interface{} with range: 0..255.
    AdminDistance interface{}

    // Producer of Imet route. The type is interface{} with range: 0..255.
    ProducerId interface{}

    // Topo ID. The type is interface{} with range: 0..4294967295.
    TopoId interface{}

    // Ethernet Tag ID. The type is interface{} with range: 0..4294967295.
    EthernetTagId interface{}
}

func (imet *L2rib_EviChildTables_Imets_Imet) GetEntityData() *types.CommonEntityData {
    imet.EntityData.YFilter = imet.YFilter
    imet.EntityData.YangName = "imet"
    imet.EntityData.BundleName = "cisco_ios_xr"
    imet.EntityData.ParentYangName = "imets"
    imet.EntityData.SegmentPath = "imet" + types.AddNoKeyToken(imet)
    imet.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/imets/" + imet.EntityData.SegmentPath
    imet.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    imet.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    imet.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    imet.EntityData.Children = types.NewOrderedMap()
    imet.EntityData.Leafs = types.NewOrderedMap()
    imet.EntityData.Leafs.Append("evi", types.YLeaf{"Evi", imet.Evi})
    imet.EntityData.Leafs.Append("tag-id", types.YLeaf{"TagId", imet.TagId})
    imet.EntityData.Leafs.Append("ip-addr", types.YLeaf{"IpAddr", imet.IpAddr})
    imet.EntityData.Leafs.Append("admin-dist", types.YLeaf{"AdminDist", imet.AdminDist})
    imet.EntityData.Leafs.Append("prod-id", types.YLeaf{"ProdId", imet.ProdId})
    imet.EntityData.Leafs.Append("vtepi-paddr", types.YLeaf{"VtepiPaddr", imet.VtepiPaddr})
    imet.EntityData.Leafs.Append("admin-distance", types.YLeaf{"AdminDistance", imet.AdminDistance})
    imet.EntityData.Leafs.Append("producer-id", types.YLeaf{"ProducerId", imet.ProducerId})
    imet.EntityData.Leafs.Append("topo-id", types.YLeaf{"TopoId", imet.TopoId})
    imet.EntityData.Leafs.Append("ethernet-tag-id", types.YLeaf{"EthernetTagId", imet.EthernetTagId})

    imet.EntityData.YListKeys = []string {}

    return &(imet.EntityData)
}

// L2rib_EviChildTables_MacDetails
// L2RIB EVPN EVI MAC Detail table
type L2rib_EviChildTables_MacDetails struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2RIB EVPN MAC Detail table. The type is slice of
    // L2rib_EviChildTables_MacDetails_MacDetail.
    MacDetail []*L2rib_EviChildTables_MacDetails_MacDetail
}

func (macDetails *L2rib_EviChildTables_MacDetails) GetEntityData() *types.CommonEntityData {
    macDetails.EntityData.YFilter = macDetails.YFilter
    macDetails.EntityData.YangName = "mac-details"
    macDetails.EntityData.BundleName = "cisco_ios_xr"
    macDetails.EntityData.ParentYangName = "evi-child-tables"
    macDetails.EntityData.SegmentPath = "mac-details"
    macDetails.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/" + macDetails.EntityData.SegmentPath
    macDetails.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macDetails.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macDetails.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macDetails.EntityData.Children = types.NewOrderedMap()
    macDetails.EntityData.Children.Append("mac-detail", types.YChild{"MacDetail", nil})
    for i := range macDetails.MacDetail {
        types.SetYListKey(macDetails.MacDetail[i], i)
        macDetails.EntityData.Children.Append(types.GetSegmentPath(macDetails.MacDetail[i]), types.YChild{"MacDetail", macDetails.MacDetail[i]})
    }
    macDetails.EntityData.Leafs = types.NewOrderedMap()

    macDetails.EntityData.YListKeys = []string {}

    return &(macDetails.EntityData)
}

// L2rib_EviChildTables_MacDetails_MacDetail
// L2RIB EVPN MAC Detail table
type L2rib_EviChildTables_MacDetails_MacDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // EVPN ID. The type is interface{} with range: 0..4294967295.
    Evi interface{}

    // Tag ID. The type is interface{} with range: 0..4294967295.
    TagId interface{}

    // MAC Address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddr interface{}

    // Admin distance. The type is interface{} with range: 0..4294967295.
    AdminDist interface{}

    // Producer ID. The type is interface{} with range: 0..4294967295.
    ProdId interface{}

    // MAC route sequence Number. The type is interface{} with range:
    // 0..4294967295.
    SequenceNumber interface{}

    // MAC route flags. The type is string with pattern: [0-9a-fA-F]{1,8}.
    Flags interface{}

    // BASE flags. The type is string with pattern: [0-9a-fA-F]{1,8}.
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
    MacRoute L2rib_EviChildTables_MacDetails_MacDetail_MacRoute

    // Mac Route Opaque Data TLV.
    RtTlv L2rib_EviChildTables_MacDetails_MacDetail_RtTlv
}

func (macDetail *L2rib_EviChildTables_MacDetails_MacDetail) GetEntityData() *types.CommonEntityData {
    macDetail.EntityData.YFilter = macDetail.YFilter
    macDetail.EntityData.YangName = "mac-detail"
    macDetail.EntityData.BundleName = "cisco_ios_xr"
    macDetail.EntityData.ParentYangName = "mac-details"
    macDetail.EntityData.SegmentPath = "mac-detail" + types.AddNoKeyToken(macDetail)
    macDetail.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/mac-details/" + macDetail.EntityData.SegmentPath
    macDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macDetail.EntityData.Children = types.NewOrderedMap()
    macDetail.EntityData.Children.Append("mac-route", types.YChild{"MacRoute", &macDetail.MacRoute})
    macDetail.EntityData.Children.Append("rt-tlv", types.YChild{"RtTlv", &macDetail.RtTlv})
    macDetail.EntityData.Leafs = types.NewOrderedMap()
    macDetail.EntityData.Leafs.Append("evi", types.YLeaf{"Evi", macDetail.Evi})
    macDetail.EntityData.Leafs.Append("tag-id", types.YLeaf{"TagId", macDetail.TagId})
    macDetail.EntityData.Leafs.Append("mac-addr", types.YLeaf{"MacAddr", macDetail.MacAddr})
    macDetail.EntityData.Leafs.Append("admin-dist", types.YLeaf{"AdminDist", macDetail.AdminDist})
    macDetail.EntityData.Leafs.Append("prod-id", types.YLeaf{"ProdId", macDetail.ProdId})
    macDetail.EntityData.Leafs.Append("sequence-number", types.YLeaf{"SequenceNumber", macDetail.SequenceNumber})
    macDetail.EntityData.Leafs.Append("flags", types.YLeaf{"Flags", macDetail.Flags})
    macDetail.EntityData.Leafs.Append("baseflags", types.YLeaf{"Baseflags", macDetail.Baseflags})
    macDetail.EntityData.Leafs.Append("soo", types.YLeaf{"Soo", macDetail.Soo})
    macDetail.EntityData.Leafs.Append("slot-id", types.YLeaf{"SlotId", macDetail.SlotId})
    macDetail.EntityData.Leafs.Append("esi", types.YLeaf{"Esi", macDetail.Esi})
    macDetail.EntityData.Leafs.Append("last-update-timestamp", types.YLeaf{"LastUpdateTimestamp", macDetail.LastUpdateTimestamp})

    macDetail.EntityData.YListKeys = []string {}

    return &(macDetail.EntityData)
}

// L2rib_EviChildTables_MacDetails_MacDetail_MacRoute
// MAC Route
type L2rib_EviChildTables_MacDetails_MacDetail_MacRoute struct {
    EntityData types.CommonEntityData
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
    Route L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route
}

func (macRoute *L2rib_EviChildTables_MacDetails_MacDetail_MacRoute) GetEntityData() *types.CommonEntityData {
    macRoute.EntityData.YFilter = macRoute.YFilter
    macRoute.EntityData.YangName = "mac-route"
    macRoute.EntityData.BundleName = "cisco_ios_xr"
    macRoute.EntityData.ParentYangName = "mac-detail"
    macRoute.EntityData.SegmentPath = "mac-route"
    macRoute.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/mac-details/mac-detail/" + macRoute.EntityData.SegmentPath
    macRoute.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macRoute.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macRoute.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macRoute.EntityData.Children = types.NewOrderedMap()
    macRoute.EntityData.Children.Append("route", types.YChild{"Route", &macRoute.Route})
    macRoute.EntityData.Leafs = types.NewOrderedMap()
    macRoute.EntityData.Leafs.Append("mac-address", types.YLeaf{"MacAddress", macRoute.MacAddress})
    macRoute.EntityData.Leafs.Append("admin-distance", types.YLeaf{"AdminDistance", macRoute.AdminDistance})
    macRoute.EntityData.Leafs.Append("producer-id", types.YLeaf{"ProducerId", macRoute.ProducerId})
    macRoute.EntityData.Leafs.Append("topology-id", types.YLeaf{"TopologyId", macRoute.TopologyId})

    macRoute.EntityData.YListKeys = []string {}

    return &(macRoute.EntityData)
}

// L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route
// MAC route
type L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type. The type is L2ribMacRoute.
    Type interface{}

    // Regular MAC route.
    Regular L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular

    // EVPN ESI MAC route.
    EvpnEsi L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi

    // BMAC route.
    Bmac L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac
}

func (route *L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route) GetEntityData() *types.CommonEntityData {
    route.EntityData.YFilter = route.YFilter
    route.EntityData.YangName = "route"
    route.EntityData.BundleName = "cisco_ios_xr"
    route.EntityData.ParentYangName = "mac-route"
    route.EntityData.SegmentPath = "route"
    route.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/mac-details/mac-detail/mac-route/" + route.EntityData.SegmentPath
    route.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    route.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    route.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    route.EntityData.Children = types.NewOrderedMap()
    route.EntityData.Children.Append("regular", types.YChild{"Regular", &route.Regular})
    route.EntityData.Children.Append("evpn-esi", types.YChild{"EvpnEsi", &route.EvpnEsi})
    route.EntityData.Children.Append("bmac", types.YChild{"Bmac", &route.Bmac})
    route.EntityData.Leafs = types.NewOrderedMap()
    route.EntityData.Leafs.Append("type", types.YLeaf{"Type", route.Type})

    route.EntityData.YListKeys = []string {}

    return &(route.EntityData)
}

// L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular
// Regular MAC route
type L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Next Hop.
    NextHop L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular_NextHop
}

func (regular *L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular) GetEntityData() *types.CommonEntityData {
    regular.EntityData.YFilter = regular.YFilter
    regular.EntityData.YangName = "regular"
    regular.EntityData.BundleName = "cisco_ios_xr"
    regular.EntityData.ParentYangName = "route"
    regular.EntityData.SegmentPath = "regular"
    regular.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/mac-details/mac-detail/mac-route/route/" + regular.EntityData.SegmentPath
    regular.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    regular.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    regular.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    regular.EntityData.Children = types.NewOrderedMap()
    regular.EntityData.Children.Append("next-hop", types.YChild{"NextHop", &regular.NextHop})
    regular.EntityData.Leafs = types.NewOrderedMap()

    regular.EntityData.YListKeys = []string {}

    return &(regular.EntityData)
}

// L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular_NextHop
// Next Hop
type L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular_NextHop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Next-hop TOPOLOGY ID. The type is interface{} with range: 0..4294967295.
    TopologyId interface{}

    // Next-hop flags. The type is interface{} with range: 0..65535.
    Flags interface{}

    // Next hop.
    NextHop L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular_NextHop_NextHop
}

func (nextHop *L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular_NextHop) GetEntityData() *types.CommonEntityData {
    nextHop.EntityData.YFilter = nextHop.YFilter
    nextHop.EntityData.YangName = "next-hop"
    nextHop.EntityData.BundleName = "cisco_ios_xr"
    nextHop.EntityData.ParentYangName = "regular"
    nextHop.EntityData.SegmentPath = "next-hop"
    nextHop.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/mac-details/mac-detail/mac-route/route/regular/" + nextHop.EntityData.SegmentPath
    nextHop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nextHop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nextHop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nextHop.EntityData.Children = types.NewOrderedMap()
    nextHop.EntityData.Children.Append("next-hop", types.YChild{"NextHop", &nextHop.NextHop})
    nextHop.EntityData.Leafs = types.NewOrderedMap()
    nextHop.EntityData.Leafs.Append("topology-id", types.YLeaf{"TopologyId", nextHop.TopologyId})
    nextHop.EntityData.Leafs.Append("flags", types.YLeaf{"Flags", nextHop.Flags})

    nextHop.EntityData.YListKeys = []string {}

    return &(nextHop.EntityData)
}

// L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular_NextHop_NextHop
// Next hop
type L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular_NextHop_NextHop struct {
    EntityData types.CommonEntityData
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
    // [a-zA-Z0-9._/-]+.
    InterfaceHandle interface{}

    // XID Next Hop. The type is interface{} with range: 0..4294967295.
    Xid interface{}

    // Labeled Next Hop.
    Labeled L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular_NextHop_NextHop_Labeled
}

func (nextHop *L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular_NextHop_NextHop) GetEntityData() *types.CommonEntityData {
    nextHop.EntityData.YFilter = nextHop.YFilter
    nextHop.EntityData.YangName = "next-hop"
    nextHop.EntityData.BundleName = "cisco_ios_xr"
    nextHop.EntityData.ParentYangName = "next-hop"
    nextHop.EntityData.SegmentPath = "next-hop"
    nextHop.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/mac-details/mac-detail/mac-route/route/regular/next-hop/" + nextHop.EntityData.SegmentPath
    nextHop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nextHop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nextHop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nextHop.EntityData.Children = types.NewOrderedMap()
    nextHop.EntityData.Children.Append("labeled", types.YChild{"Labeled", &nextHop.Labeled})
    nextHop.EntityData.Leafs = types.NewOrderedMap()
    nextHop.EntityData.Leafs.Append("type", types.YLeaf{"Type", nextHop.Type})
    nextHop.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", nextHop.Ipv4})
    nextHop.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", nextHop.Ipv6})
    nextHop.EntityData.Leafs.Append("mac", types.YLeaf{"Mac", nextHop.Mac})
    nextHop.EntityData.Leafs.Append("interface-handle", types.YLeaf{"InterfaceHandle", nextHop.InterfaceHandle})
    nextHop.EntityData.Leafs.Append("xid", types.YLeaf{"Xid", nextHop.Xid})

    nextHop.EntityData.YListKeys = []string {}

    return &(nextHop.EntityData)
}

// L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular_NextHop_NextHop_Labeled
// Labeled Next Hop
type L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular_NextHop_NextHop_Labeled struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2RIB Address Family. The type is L2ribAfi.
    AddressFamily interface{}

    // IP Address. The type is string.
    IpAddress interface{}

    // Label. The type is interface{} with range: 0..4294967295.
    Label interface{}

    // Internal Label. The type is bool.
    Internal interface{}
}

func (labeled *L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Regular_NextHop_NextHop_Labeled) GetEntityData() *types.CommonEntityData {
    labeled.EntityData.YFilter = labeled.YFilter
    labeled.EntityData.YangName = "labeled"
    labeled.EntityData.BundleName = "cisco_ios_xr"
    labeled.EntityData.ParentYangName = "next-hop"
    labeled.EntityData.SegmentPath = "labeled"
    labeled.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/mac-details/mac-detail/mac-route/route/regular/next-hop/next-hop/" + labeled.EntityData.SegmentPath
    labeled.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    labeled.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    labeled.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    labeled.EntityData.Children = types.NewOrderedMap()
    labeled.EntityData.Leafs = types.NewOrderedMap()
    labeled.EntityData.Leafs.Append("address-family", types.YLeaf{"AddressFamily", labeled.AddressFamily})
    labeled.EntityData.Leafs.Append("ip-address", types.YLeaf{"IpAddress", labeled.IpAddress})
    labeled.EntityData.Leafs.Append("label", types.YLeaf{"Label", labeled.Label})
    labeled.EntityData.Leafs.Append("internal", types.YLeaf{"Internal", labeled.Internal})

    labeled.EntityData.YListKeys = []string {}

    return &(labeled.EntityData)
}

// L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi
// EVPN ESI MAC route
type L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MAC route sequence number. The type is interface{} with range:
    // 0..4294967295.
    SequenceNumber interface{}

    // Forwarding State. True means forward, False means drop. The type is bool.
    ForwardState interface{}

    // Ethernet Segment Identifier.
    EthernetSegmentId L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_EthernetSegmentId

    // Path list information. Set for detailed MAC route information.
    PathList L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList
}

func (evpnEsi *L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi) GetEntityData() *types.CommonEntityData {
    evpnEsi.EntityData.YFilter = evpnEsi.YFilter
    evpnEsi.EntityData.YangName = "evpn-esi"
    evpnEsi.EntityData.BundleName = "cisco_ios_xr"
    evpnEsi.EntityData.ParentYangName = "route"
    evpnEsi.EntityData.SegmentPath = "evpn-esi"
    evpnEsi.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/mac-details/mac-detail/mac-route/route/" + evpnEsi.EntityData.SegmentPath
    evpnEsi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnEsi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnEsi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnEsi.EntityData.Children = types.NewOrderedMap()
    evpnEsi.EntityData.Children.Append("ethernet-segment-id", types.YChild{"EthernetSegmentId", &evpnEsi.EthernetSegmentId})
    evpnEsi.EntityData.Children.Append("path-list", types.YChild{"PathList", &evpnEsi.PathList})
    evpnEsi.EntityData.Leafs = types.NewOrderedMap()
    evpnEsi.EntityData.Leafs.Append("sequence-number", types.YLeaf{"SequenceNumber", evpnEsi.SequenceNumber})
    evpnEsi.EntityData.Leafs.Append("forward-state", types.YLeaf{"ForwardState", evpnEsi.ForwardState})

    evpnEsi.EntityData.YListKeys = []string {}

    return &(evpnEsi.EntityData)
}

// L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_EthernetSegmentId
// Ethernet Segment Identifier
type L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_EthernetSegmentId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LACP System Priority. The type is interface{} with range: 0..65535.
    SystemPriority interface{}

    // LACP System Id. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    SystemId interface{}

    // LACP Port Key. The type is interface{} with range: 0..65535.
    PortKey interface{}
}

func (ethernetSegmentId *L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_EthernetSegmentId) GetEntityData() *types.CommonEntityData {
    ethernetSegmentId.EntityData.YFilter = ethernetSegmentId.YFilter
    ethernetSegmentId.EntityData.YangName = "ethernet-segment-id"
    ethernetSegmentId.EntityData.BundleName = "cisco_ios_xr"
    ethernetSegmentId.EntityData.ParentYangName = "evpn-esi"
    ethernetSegmentId.EntityData.SegmentPath = "ethernet-segment-id"
    ethernetSegmentId.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/mac-details/mac-detail/mac-route/route/evpn-esi/" + ethernetSegmentId.EntityData.SegmentPath
    ethernetSegmentId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ethernetSegmentId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ethernetSegmentId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ethernetSegmentId.EntityData.Children = types.NewOrderedMap()
    ethernetSegmentId.EntityData.Leafs = types.NewOrderedMap()
    ethernetSegmentId.EntityData.Leafs.Append("system-priority", types.YLeaf{"SystemPriority", ethernetSegmentId.SystemPriority})
    ethernetSegmentId.EntityData.Leafs.Append("system-id", types.YLeaf{"SystemId", ethernetSegmentId.SystemId})
    ethernetSegmentId.EntityData.Leafs.Append("port-key", types.YLeaf{"PortKey", ethernetSegmentId.PortKey})

    ethernetSegmentId.EntityData.YListKeys = []string {}

    return &(ethernetSegmentId.EntityData)
}

// L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList
// Path list information. Set for detailed MAC
// route information
type L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ID of EAD route producer. The type is interface{} with range: 0..255.
    ProducerId interface{}

    // Number of MAC routes bound to this Path list. The type is interface{} with
    // range: 0..4294967295.
    MacCount interface{}

    // Path list local Label. The type is interface{} with range: 0..4294967295.
    LocalLabel interface{}

    // Type-specific Path List info.
    PathListInfo L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo

    // Array of Next Hops for MAC Path List. The type is slice of
    // L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_NextHopArray.
    NextHopArray []*L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_NextHopArray
}

func (pathList *L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList) GetEntityData() *types.CommonEntityData {
    pathList.EntityData.YFilter = pathList.YFilter
    pathList.EntityData.YangName = "path-list"
    pathList.EntityData.BundleName = "cisco_ios_xr"
    pathList.EntityData.ParentYangName = "evpn-esi"
    pathList.EntityData.SegmentPath = "path-list"
    pathList.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/mac-details/mac-detail/mac-route/route/evpn-esi/" + pathList.EntityData.SegmentPath
    pathList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pathList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pathList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pathList.EntityData.Children = types.NewOrderedMap()
    pathList.EntityData.Children.Append("path-list-info", types.YChild{"PathListInfo", &pathList.PathListInfo})
    pathList.EntityData.Children.Append("next-hop-array", types.YChild{"NextHopArray", nil})
    for i := range pathList.NextHopArray {
        types.SetYListKey(pathList.NextHopArray[i], i)
        pathList.EntityData.Children.Append(types.GetSegmentPath(pathList.NextHopArray[i]), types.YChild{"NextHopArray", pathList.NextHopArray[i]})
    }
    pathList.EntityData.Leafs = types.NewOrderedMap()
    pathList.EntityData.Leafs.Append("producer-id", types.YLeaf{"ProducerId", pathList.ProducerId})
    pathList.EntityData.Leafs.Append("mac-count", types.YLeaf{"MacCount", pathList.MacCount})
    pathList.EntityData.Leafs.Append("local-label", types.YLeaf{"LocalLabel", pathList.LocalLabel})

    pathList.EntityData.YListKeys = []string {}

    return &(pathList.EntityData)
}

// L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo
// Type-specific Path List info
type L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type. The type is L2ribMacRoute.
    Type interface{}

    // ESI Path List.
    PathListEsi L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi

    // MAC Path List.
    PathListMac L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListMac
}

func (pathListInfo *L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo) GetEntityData() *types.CommonEntityData {
    pathListInfo.EntityData.YFilter = pathListInfo.YFilter
    pathListInfo.EntityData.YangName = "path-list-info"
    pathListInfo.EntityData.BundleName = "cisco_ios_xr"
    pathListInfo.EntityData.ParentYangName = "path-list"
    pathListInfo.EntityData.SegmentPath = "path-list-info"
    pathListInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/mac-details/mac-detail/mac-route/route/evpn-esi/path-list/" + pathListInfo.EntityData.SegmentPath
    pathListInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pathListInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pathListInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pathListInfo.EntityData.Children = types.NewOrderedMap()
    pathListInfo.EntityData.Children.Append("path-list-esi", types.YChild{"PathListEsi", &pathListInfo.PathListEsi})
    pathListInfo.EntityData.Children.Append("path-list-mac", types.YChild{"PathListMac", &pathListInfo.PathListMac})
    pathListInfo.EntityData.Leafs = types.NewOrderedMap()
    pathListInfo.EntityData.Leafs.Append("type", types.YLeaf{"Type", pathListInfo.Type})

    pathListInfo.EntityData.YListKeys = []string {}

    return &(pathListInfo.EntityData)
}

// L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi
// ESI Path List
type L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Path list Resolved. The type is bool.
    Resolved interface{}

    // Ethernet Segment Identifier.
    EthernetSegmentId L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_EthernetSegmentId

    // Array of Next Hops from MAC Update. The type is slice of
    // L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray.
    MacUpdateNextHopArray []*L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray
}

func (pathListEsi *L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi) GetEntityData() *types.CommonEntityData {
    pathListEsi.EntityData.YFilter = pathListEsi.YFilter
    pathListEsi.EntityData.YangName = "path-list-esi"
    pathListEsi.EntityData.BundleName = "cisco_ios_xr"
    pathListEsi.EntityData.ParentYangName = "path-list-info"
    pathListEsi.EntityData.SegmentPath = "path-list-esi"
    pathListEsi.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/mac-details/mac-detail/mac-route/route/evpn-esi/path-list/path-list-info/" + pathListEsi.EntityData.SegmentPath
    pathListEsi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pathListEsi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pathListEsi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pathListEsi.EntityData.Children = types.NewOrderedMap()
    pathListEsi.EntityData.Children.Append("ethernet-segment-id", types.YChild{"EthernetSegmentId", &pathListEsi.EthernetSegmentId})
    pathListEsi.EntityData.Children.Append("mac-update-next-hop-array", types.YChild{"MacUpdateNextHopArray", nil})
    for i := range pathListEsi.MacUpdateNextHopArray {
        types.SetYListKey(pathListEsi.MacUpdateNextHopArray[i], i)
        pathListEsi.EntityData.Children.Append(types.GetSegmentPath(pathListEsi.MacUpdateNextHopArray[i]), types.YChild{"MacUpdateNextHopArray", pathListEsi.MacUpdateNextHopArray[i]})
    }
    pathListEsi.EntityData.Leafs = types.NewOrderedMap()
    pathListEsi.EntityData.Leafs.Append("resolved", types.YLeaf{"Resolved", pathListEsi.Resolved})

    pathListEsi.EntityData.YListKeys = []string {}

    return &(pathListEsi.EntityData)
}

// L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_EthernetSegmentId
// Ethernet Segment Identifier
type L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_EthernetSegmentId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LACP System Priority. The type is interface{} with range: 0..65535.
    SystemPriority interface{}

    // LACP System Id. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    SystemId interface{}

    // LACP Port Key. The type is interface{} with range: 0..65535.
    PortKey interface{}
}

func (ethernetSegmentId *L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_EthernetSegmentId) GetEntityData() *types.CommonEntityData {
    ethernetSegmentId.EntityData.YFilter = ethernetSegmentId.YFilter
    ethernetSegmentId.EntityData.YangName = "ethernet-segment-id"
    ethernetSegmentId.EntityData.BundleName = "cisco_ios_xr"
    ethernetSegmentId.EntityData.ParentYangName = "path-list-esi"
    ethernetSegmentId.EntityData.SegmentPath = "ethernet-segment-id"
    ethernetSegmentId.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/mac-details/mac-detail/mac-route/route/evpn-esi/path-list/path-list-info/path-list-esi/" + ethernetSegmentId.EntityData.SegmentPath
    ethernetSegmentId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ethernetSegmentId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ethernetSegmentId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ethernetSegmentId.EntityData.Children = types.NewOrderedMap()
    ethernetSegmentId.EntityData.Leafs = types.NewOrderedMap()
    ethernetSegmentId.EntityData.Leafs.Append("system-priority", types.YLeaf{"SystemPriority", ethernetSegmentId.SystemPriority})
    ethernetSegmentId.EntityData.Leafs.Append("system-id", types.YLeaf{"SystemId", ethernetSegmentId.SystemId})
    ethernetSegmentId.EntityData.Leafs.Append("port-key", types.YLeaf{"PortKey", ethernetSegmentId.PortKey})

    ethernetSegmentId.EntityData.YListKeys = []string {}

    return &(ethernetSegmentId.EntityData)
}

// L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray
// Array of Next Hops from MAC Update
type L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Next-hop TOPOLOGY ID. The type is interface{} with range: 0..4294967295.
    TopologyId interface{}

    // Next-hop flags. The type is interface{} with range: 0..65535.
    Flags interface{}

    // Next hop.
    NextHop L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop
}

func (macUpdateNextHopArray *L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray) GetEntityData() *types.CommonEntityData {
    macUpdateNextHopArray.EntityData.YFilter = macUpdateNextHopArray.YFilter
    macUpdateNextHopArray.EntityData.YangName = "mac-update-next-hop-array"
    macUpdateNextHopArray.EntityData.BundleName = "cisco_ios_xr"
    macUpdateNextHopArray.EntityData.ParentYangName = "path-list-esi"
    macUpdateNextHopArray.EntityData.SegmentPath = "mac-update-next-hop-array" + types.AddNoKeyToken(macUpdateNextHopArray)
    macUpdateNextHopArray.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/mac-details/mac-detail/mac-route/route/evpn-esi/path-list/path-list-info/path-list-esi/" + macUpdateNextHopArray.EntityData.SegmentPath
    macUpdateNextHopArray.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macUpdateNextHopArray.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macUpdateNextHopArray.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macUpdateNextHopArray.EntityData.Children = types.NewOrderedMap()
    macUpdateNextHopArray.EntityData.Children.Append("next-hop", types.YChild{"NextHop", &macUpdateNextHopArray.NextHop})
    macUpdateNextHopArray.EntityData.Leafs = types.NewOrderedMap()
    macUpdateNextHopArray.EntityData.Leafs.Append("topology-id", types.YLeaf{"TopologyId", macUpdateNextHopArray.TopologyId})
    macUpdateNextHopArray.EntityData.Leafs.Append("flags", types.YLeaf{"Flags", macUpdateNextHopArray.Flags})

    macUpdateNextHopArray.EntityData.YListKeys = []string {}

    return &(macUpdateNextHopArray.EntityData)
}

// L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop
// Next hop
type L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop struct {
    EntityData types.CommonEntityData
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
    // [a-zA-Z0-9._/-]+.
    InterfaceHandle interface{}

    // XID Next Hop. The type is interface{} with range: 0..4294967295.
    Xid interface{}

    // Labeled Next Hop.
    Labeled L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled
}

func (nextHop *L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop) GetEntityData() *types.CommonEntityData {
    nextHop.EntityData.YFilter = nextHop.YFilter
    nextHop.EntityData.YangName = "next-hop"
    nextHop.EntityData.BundleName = "cisco_ios_xr"
    nextHop.EntityData.ParentYangName = "mac-update-next-hop-array"
    nextHop.EntityData.SegmentPath = "next-hop"
    nextHop.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/mac-details/mac-detail/mac-route/route/evpn-esi/path-list/path-list-info/path-list-esi/mac-update-next-hop-array/" + nextHop.EntityData.SegmentPath
    nextHop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nextHop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nextHop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nextHop.EntityData.Children = types.NewOrderedMap()
    nextHop.EntityData.Children.Append("labeled", types.YChild{"Labeled", &nextHop.Labeled})
    nextHop.EntityData.Leafs = types.NewOrderedMap()
    nextHop.EntityData.Leafs.Append("type", types.YLeaf{"Type", nextHop.Type})
    nextHop.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", nextHop.Ipv4})
    nextHop.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", nextHop.Ipv6})
    nextHop.EntityData.Leafs.Append("mac", types.YLeaf{"Mac", nextHop.Mac})
    nextHop.EntityData.Leafs.Append("interface-handle", types.YLeaf{"InterfaceHandle", nextHop.InterfaceHandle})
    nextHop.EntityData.Leafs.Append("xid", types.YLeaf{"Xid", nextHop.Xid})

    nextHop.EntityData.YListKeys = []string {}

    return &(nextHop.EntityData)
}

// L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled
// Labeled Next Hop
type L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2RIB Address Family. The type is L2ribAfi.
    AddressFamily interface{}

    // IP Address. The type is string.
    IpAddress interface{}

    // Label. The type is interface{} with range: 0..4294967295.
    Label interface{}

    // Internal Label. The type is bool.
    Internal interface{}
}

func (labeled *L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled) GetEntityData() *types.CommonEntityData {
    labeled.EntityData.YFilter = labeled.YFilter
    labeled.EntityData.YangName = "labeled"
    labeled.EntityData.BundleName = "cisco_ios_xr"
    labeled.EntityData.ParentYangName = "next-hop"
    labeled.EntityData.SegmentPath = "labeled"
    labeled.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/mac-details/mac-detail/mac-route/route/evpn-esi/path-list/path-list-info/path-list-esi/mac-update-next-hop-array/next-hop/" + labeled.EntityData.SegmentPath
    labeled.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    labeled.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    labeled.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    labeled.EntityData.Children = types.NewOrderedMap()
    labeled.EntityData.Leafs = types.NewOrderedMap()
    labeled.EntityData.Leafs.Append("address-family", types.YLeaf{"AddressFamily", labeled.AddressFamily})
    labeled.EntityData.Leafs.Append("ip-address", types.YLeaf{"IpAddress", labeled.IpAddress})
    labeled.EntityData.Leafs.Append("label", types.YLeaf{"Label", labeled.Label})
    labeled.EntityData.Leafs.Append("internal", types.YLeaf{"Internal", labeled.Internal})

    labeled.EntityData.YListKeys = []string {}

    return &(labeled.EntityData)
}

// L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListMac
// MAC Path List
type L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListMac struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MAC Address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddress interface{}
}

func (pathListMac *L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_PathListInfo_PathListMac) GetEntityData() *types.CommonEntityData {
    pathListMac.EntityData.YFilter = pathListMac.YFilter
    pathListMac.EntityData.YangName = "path-list-mac"
    pathListMac.EntityData.BundleName = "cisco_ios_xr"
    pathListMac.EntityData.ParentYangName = "path-list-info"
    pathListMac.EntityData.SegmentPath = "path-list-mac"
    pathListMac.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/mac-details/mac-detail/mac-route/route/evpn-esi/path-list/path-list-info/" + pathListMac.EntityData.SegmentPath
    pathListMac.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pathListMac.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pathListMac.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pathListMac.EntityData.Children = types.NewOrderedMap()
    pathListMac.EntityData.Leafs = types.NewOrderedMap()
    pathListMac.EntityData.Leafs.Append("mac-address", types.YLeaf{"MacAddress", pathListMac.MacAddress})

    pathListMac.EntityData.YListKeys = []string {}

    return &(pathListMac.EntityData)
}

// L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_NextHopArray
// Array of Next Hops for MAC Path List
type L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_NextHopArray struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Next-hop TOPOLOGY ID. The type is interface{} with range: 0..4294967295.
    TopologyId interface{}

    // Next-hop flags. The type is interface{} with range: 0..65535.
    Flags interface{}

    // Next hop.
    NextHop L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_NextHopArray_NextHop
}

func (nextHopArray *L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_NextHopArray) GetEntityData() *types.CommonEntityData {
    nextHopArray.EntityData.YFilter = nextHopArray.YFilter
    nextHopArray.EntityData.YangName = "next-hop-array"
    nextHopArray.EntityData.BundleName = "cisco_ios_xr"
    nextHopArray.EntityData.ParentYangName = "path-list"
    nextHopArray.EntityData.SegmentPath = "next-hop-array" + types.AddNoKeyToken(nextHopArray)
    nextHopArray.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/mac-details/mac-detail/mac-route/route/evpn-esi/path-list/" + nextHopArray.EntityData.SegmentPath
    nextHopArray.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nextHopArray.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nextHopArray.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nextHopArray.EntityData.Children = types.NewOrderedMap()
    nextHopArray.EntityData.Children.Append("next-hop", types.YChild{"NextHop", &nextHopArray.NextHop})
    nextHopArray.EntityData.Leafs = types.NewOrderedMap()
    nextHopArray.EntityData.Leafs.Append("topology-id", types.YLeaf{"TopologyId", nextHopArray.TopologyId})
    nextHopArray.EntityData.Leafs.Append("flags", types.YLeaf{"Flags", nextHopArray.Flags})

    nextHopArray.EntityData.YListKeys = []string {}

    return &(nextHopArray.EntityData)
}

// L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_NextHopArray_NextHop
// Next hop
type L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_NextHopArray_NextHop struct {
    EntityData types.CommonEntityData
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
    // [a-zA-Z0-9._/-]+.
    InterfaceHandle interface{}

    // XID Next Hop. The type is interface{} with range: 0..4294967295.
    Xid interface{}

    // Labeled Next Hop.
    Labeled L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_NextHopArray_NextHop_Labeled
}

func (nextHop *L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_NextHopArray_NextHop) GetEntityData() *types.CommonEntityData {
    nextHop.EntityData.YFilter = nextHop.YFilter
    nextHop.EntityData.YangName = "next-hop"
    nextHop.EntityData.BundleName = "cisco_ios_xr"
    nextHop.EntityData.ParentYangName = "next-hop-array"
    nextHop.EntityData.SegmentPath = "next-hop"
    nextHop.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/mac-details/mac-detail/mac-route/route/evpn-esi/path-list/next-hop-array/" + nextHop.EntityData.SegmentPath
    nextHop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nextHop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nextHop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nextHop.EntityData.Children = types.NewOrderedMap()
    nextHop.EntityData.Children.Append("labeled", types.YChild{"Labeled", &nextHop.Labeled})
    nextHop.EntityData.Leafs = types.NewOrderedMap()
    nextHop.EntityData.Leafs.Append("type", types.YLeaf{"Type", nextHop.Type})
    nextHop.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", nextHop.Ipv4})
    nextHop.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", nextHop.Ipv6})
    nextHop.EntityData.Leafs.Append("mac", types.YLeaf{"Mac", nextHop.Mac})
    nextHop.EntityData.Leafs.Append("interface-handle", types.YLeaf{"InterfaceHandle", nextHop.InterfaceHandle})
    nextHop.EntityData.Leafs.Append("xid", types.YLeaf{"Xid", nextHop.Xid})

    nextHop.EntityData.YListKeys = []string {}

    return &(nextHop.EntityData)
}

// L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_NextHopArray_NextHop_Labeled
// Labeled Next Hop
type L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_NextHopArray_NextHop_Labeled struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2RIB Address Family. The type is L2ribAfi.
    AddressFamily interface{}

    // IP Address. The type is string.
    IpAddress interface{}

    // Label. The type is interface{} with range: 0..4294967295.
    Label interface{}

    // Internal Label. The type is bool.
    Internal interface{}
}

func (labeled *L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_EvpnEsi_PathList_NextHopArray_NextHop_Labeled) GetEntityData() *types.CommonEntityData {
    labeled.EntityData.YFilter = labeled.YFilter
    labeled.EntityData.YangName = "labeled"
    labeled.EntityData.BundleName = "cisco_ios_xr"
    labeled.EntityData.ParentYangName = "next-hop"
    labeled.EntityData.SegmentPath = "labeled"
    labeled.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/mac-details/mac-detail/mac-route/route/evpn-esi/path-list/next-hop-array/next-hop/" + labeled.EntityData.SegmentPath
    labeled.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    labeled.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    labeled.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    labeled.EntityData.Children = types.NewOrderedMap()
    labeled.EntityData.Leafs = types.NewOrderedMap()
    labeled.EntityData.Leafs.Append("address-family", types.YLeaf{"AddressFamily", labeled.AddressFamily})
    labeled.EntityData.Leafs.Append("ip-address", types.YLeaf{"IpAddress", labeled.IpAddress})
    labeled.EntityData.Leafs.Append("label", types.YLeaf{"Label", labeled.Label})
    labeled.EntityData.Leafs.Append("internal", types.YLeaf{"Internal", labeled.Internal})

    labeled.EntityData.YListKeys = []string {}

    return &(labeled.EntityData)
}

// L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac
// BMAC route
type L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BMAC Address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    BmacAddress interface{}

    // Forwarding State. True means forward, False means drop. The type is bool.
    ForwardState interface{}

    // Path list information.
    PathList L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList
}

func (bmac *L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac) GetEntityData() *types.CommonEntityData {
    bmac.EntityData.YFilter = bmac.YFilter
    bmac.EntityData.YangName = "bmac"
    bmac.EntityData.BundleName = "cisco_ios_xr"
    bmac.EntityData.ParentYangName = "route"
    bmac.EntityData.SegmentPath = "bmac"
    bmac.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/mac-details/mac-detail/mac-route/route/" + bmac.EntityData.SegmentPath
    bmac.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bmac.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bmac.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bmac.EntityData.Children = types.NewOrderedMap()
    bmac.EntityData.Children.Append("path-list", types.YChild{"PathList", &bmac.PathList})
    bmac.EntityData.Leafs = types.NewOrderedMap()
    bmac.EntityData.Leafs.Append("bmac-address", types.YLeaf{"BmacAddress", bmac.BmacAddress})
    bmac.EntityData.Leafs.Append("forward-state", types.YLeaf{"ForwardState", bmac.ForwardState})

    bmac.EntityData.YListKeys = []string {}

    return &(bmac.EntityData)
}

// L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList
// Path list information
type L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ID of EAD route producer. The type is interface{} with range: 0..255.
    ProducerId interface{}

    // Number of MAC routes bound to this Path list. The type is interface{} with
    // range: 0..4294967295.
    MacCount interface{}

    // Path list local Label. The type is interface{} with range: 0..4294967295.
    LocalLabel interface{}

    // Type-specific Path List info.
    PathListInfo L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo

    // Array of Next Hops for MAC Path List. The type is slice of
    // L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_NextHopArray.
    NextHopArray []*L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_NextHopArray
}

func (pathList *L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList) GetEntityData() *types.CommonEntityData {
    pathList.EntityData.YFilter = pathList.YFilter
    pathList.EntityData.YangName = "path-list"
    pathList.EntityData.BundleName = "cisco_ios_xr"
    pathList.EntityData.ParentYangName = "bmac"
    pathList.EntityData.SegmentPath = "path-list"
    pathList.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/mac-details/mac-detail/mac-route/route/bmac/" + pathList.EntityData.SegmentPath
    pathList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pathList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pathList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pathList.EntityData.Children = types.NewOrderedMap()
    pathList.EntityData.Children.Append("path-list-info", types.YChild{"PathListInfo", &pathList.PathListInfo})
    pathList.EntityData.Children.Append("next-hop-array", types.YChild{"NextHopArray", nil})
    for i := range pathList.NextHopArray {
        types.SetYListKey(pathList.NextHopArray[i], i)
        pathList.EntityData.Children.Append(types.GetSegmentPath(pathList.NextHopArray[i]), types.YChild{"NextHopArray", pathList.NextHopArray[i]})
    }
    pathList.EntityData.Leafs = types.NewOrderedMap()
    pathList.EntityData.Leafs.Append("producer-id", types.YLeaf{"ProducerId", pathList.ProducerId})
    pathList.EntityData.Leafs.Append("mac-count", types.YLeaf{"MacCount", pathList.MacCount})
    pathList.EntityData.Leafs.Append("local-label", types.YLeaf{"LocalLabel", pathList.LocalLabel})

    pathList.EntityData.YListKeys = []string {}

    return &(pathList.EntityData)
}

// L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo
// Type-specific Path List info
type L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type. The type is L2ribMacRoute.
    Type interface{}

    // ESI Path List.
    PathListEsi L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi

    // MAC Path List.
    PathListMac L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListMac
}

func (pathListInfo *L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo) GetEntityData() *types.CommonEntityData {
    pathListInfo.EntityData.YFilter = pathListInfo.YFilter
    pathListInfo.EntityData.YangName = "path-list-info"
    pathListInfo.EntityData.BundleName = "cisco_ios_xr"
    pathListInfo.EntityData.ParentYangName = "path-list"
    pathListInfo.EntityData.SegmentPath = "path-list-info"
    pathListInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/mac-details/mac-detail/mac-route/route/bmac/path-list/" + pathListInfo.EntityData.SegmentPath
    pathListInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pathListInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pathListInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pathListInfo.EntityData.Children = types.NewOrderedMap()
    pathListInfo.EntityData.Children.Append("path-list-esi", types.YChild{"PathListEsi", &pathListInfo.PathListEsi})
    pathListInfo.EntityData.Children.Append("path-list-mac", types.YChild{"PathListMac", &pathListInfo.PathListMac})
    pathListInfo.EntityData.Leafs = types.NewOrderedMap()
    pathListInfo.EntityData.Leafs.Append("type", types.YLeaf{"Type", pathListInfo.Type})

    pathListInfo.EntityData.YListKeys = []string {}

    return &(pathListInfo.EntityData)
}

// L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi
// ESI Path List
type L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Path list Resolved. The type is bool.
    Resolved interface{}

    // Ethernet Segment Identifier.
    EthernetSegmentId L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_EthernetSegmentId

    // Array of Next Hops from MAC Update. The type is slice of
    // L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray.
    MacUpdateNextHopArray []*L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray
}

func (pathListEsi *L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi) GetEntityData() *types.CommonEntityData {
    pathListEsi.EntityData.YFilter = pathListEsi.YFilter
    pathListEsi.EntityData.YangName = "path-list-esi"
    pathListEsi.EntityData.BundleName = "cisco_ios_xr"
    pathListEsi.EntityData.ParentYangName = "path-list-info"
    pathListEsi.EntityData.SegmentPath = "path-list-esi"
    pathListEsi.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/mac-details/mac-detail/mac-route/route/bmac/path-list/path-list-info/" + pathListEsi.EntityData.SegmentPath
    pathListEsi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pathListEsi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pathListEsi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pathListEsi.EntityData.Children = types.NewOrderedMap()
    pathListEsi.EntityData.Children.Append("ethernet-segment-id", types.YChild{"EthernetSegmentId", &pathListEsi.EthernetSegmentId})
    pathListEsi.EntityData.Children.Append("mac-update-next-hop-array", types.YChild{"MacUpdateNextHopArray", nil})
    for i := range pathListEsi.MacUpdateNextHopArray {
        types.SetYListKey(pathListEsi.MacUpdateNextHopArray[i], i)
        pathListEsi.EntityData.Children.Append(types.GetSegmentPath(pathListEsi.MacUpdateNextHopArray[i]), types.YChild{"MacUpdateNextHopArray", pathListEsi.MacUpdateNextHopArray[i]})
    }
    pathListEsi.EntityData.Leafs = types.NewOrderedMap()
    pathListEsi.EntityData.Leafs.Append("resolved", types.YLeaf{"Resolved", pathListEsi.Resolved})

    pathListEsi.EntityData.YListKeys = []string {}

    return &(pathListEsi.EntityData)
}

// L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_EthernetSegmentId
// Ethernet Segment Identifier
type L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_EthernetSegmentId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LACP System Priority. The type is interface{} with range: 0..65535.
    SystemPriority interface{}

    // LACP System Id. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    SystemId interface{}

    // LACP Port Key. The type is interface{} with range: 0..65535.
    PortKey interface{}
}

func (ethernetSegmentId *L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_EthernetSegmentId) GetEntityData() *types.CommonEntityData {
    ethernetSegmentId.EntityData.YFilter = ethernetSegmentId.YFilter
    ethernetSegmentId.EntityData.YangName = "ethernet-segment-id"
    ethernetSegmentId.EntityData.BundleName = "cisco_ios_xr"
    ethernetSegmentId.EntityData.ParentYangName = "path-list-esi"
    ethernetSegmentId.EntityData.SegmentPath = "ethernet-segment-id"
    ethernetSegmentId.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/mac-details/mac-detail/mac-route/route/bmac/path-list/path-list-info/path-list-esi/" + ethernetSegmentId.EntityData.SegmentPath
    ethernetSegmentId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ethernetSegmentId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ethernetSegmentId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ethernetSegmentId.EntityData.Children = types.NewOrderedMap()
    ethernetSegmentId.EntityData.Leafs = types.NewOrderedMap()
    ethernetSegmentId.EntityData.Leafs.Append("system-priority", types.YLeaf{"SystemPriority", ethernetSegmentId.SystemPriority})
    ethernetSegmentId.EntityData.Leafs.Append("system-id", types.YLeaf{"SystemId", ethernetSegmentId.SystemId})
    ethernetSegmentId.EntityData.Leafs.Append("port-key", types.YLeaf{"PortKey", ethernetSegmentId.PortKey})

    ethernetSegmentId.EntityData.YListKeys = []string {}

    return &(ethernetSegmentId.EntityData)
}

// L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray
// Array of Next Hops from MAC Update
type L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Next-hop TOPOLOGY ID. The type is interface{} with range: 0..4294967295.
    TopologyId interface{}

    // Next-hop flags. The type is interface{} with range: 0..65535.
    Flags interface{}

    // Next hop.
    NextHop L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop
}

func (macUpdateNextHopArray *L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray) GetEntityData() *types.CommonEntityData {
    macUpdateNextHopArray.EntityData.YFilter = macUpdateNextHopArray.YFilter
    macUpdateNextHopArray.EntityData.YangName = "mac-update-next-hop-array"
    macUpdateNextHopArray.EntityData.BundleName = "cisco_ios_xr"
    macUpdateNextHopArray.EntityData.ParentYangName = "path-list-esi"
    macUpdateNextHopArray.EntityData.SegmentPath = "mac-update-next-hop-array" + types.AddNoKeyToken(macUpdateNextHopArray)
    macUpdateNextHopArray.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/mac-details/mac-detail/mac-route/route/bmac/path-list/path-list-info/path-list-esi/" + macUpdateNextHopArray.EntityData.SegmentPath
    macUpdateNextHopArray.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macUpdateNextHopArray.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macUpdateNextHopArray.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macUpdateNextHopArray.EntityData.Children = types.NewOrderedMap()
    macUpdateNextHopArray.EntityData.Children.Append("next-hop", types.YChild{"NextHop", &macUpdateNextHopArray.NextHop})
    macUpdateNextHopArray.EntityData.Leafs = types.NewOrderedMap()
    macUpdateNextHopArray.EntityData.Leafs.Append("topology-id", types.YLeaf{"TopologyId", macUpdateNextHopArray.TopologyId})
    macUpdateNextHopArray.EntityData.Leafs.Append("flags", types.YLeaf{"Flags", macUpdateNextHopArray.Flags})

    macUpdateNextHopArray.EntityData.YListKeys = []string {}

    return &(macUpdateNextHopArray.EntityData)
}

// L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop
// Next hop
type L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop struct {
    EntityData types.CommonEntityData
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
    // [a-zA-Z0-9._/-]+.
    InterfaceHandle interface{}

    // XID Next Hop. The type is interface{} with range: 0..4294967295.
    Xid interface{}

    // Labeled Next Hop.
    Labeled L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled
}

func (nextHop *L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop) GetEntityData() *types.CommonEntityData {
    nextHop.EntityData.YFilter = nextHop.YFilter
    nextHop.EntityData.YangName = "next-hop"
    nextHop.EntityData.BundleName = "cisco_ios_xr"
    nextHop.EntityData.ParentYangName = "mac-update-next-hop-array"
    nextHop.EntityData.SegmentPath = "next-hop"
    nextHop.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/mac-details/mac-detail/mac-route/route/bmac/path-list/path-list-info/path-list-esi/mac-update-next-hop-array/" + nextHop.EntityData.SegmentPath
    nextHop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nextHop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nextHop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nextHop.EntityData.Children = types.NewOrderedMap()
    nextHop.EntityData.Children.Append("labeled", types.YChild{"Labeled", &nextHop.Labeled})
    nextHop.EntityData.Leafs = types.NewOrderedMap()
    nextHop.EntityData.Leafs.Append("type", types.YLeaf{"Type", nextHop.Type})
    nextHop.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", nextHop.Ipv4})
    nextHop.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", nextHop.Ipv6})
    nextHop.EntityData.Leafs.Append("mac", types.YLeaf{"Mac", nextHop.Mac})
    nextHop.EntityData.Leafs.Append("interface-handle", types.YLeaf{"InterfaceHandle", nextHop.InterfaceHandle})
    nextHop.EntityData.Leafs.Append("xid", types.YLeaf{"Xid", nextHop.Xid})

    nextHop.EntityData.YListKeys = []string {}

    return &(nextHop.EntityData)
}

// L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled
// Labeled Next Hop
type L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2RIB Address Family. The type is L2ribAfi.
    AddressFamily interface{}

    // IP Address. The type is string.
    IpAddress interface{}

    // Label. The type is interface{} with range: 0..4294967295.
    Label interface{}

    // Internal Label. The type is bool.
    Internal interface{}
}

func (labeled *L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListEsi_MacUpdateNextHopArray_NextHop_Labeled) GetEntityData() *types.CommonEntityData {
    labeled.EntityData.YFilter = labeled.YFilter
    labeled.EntityData.YangName = "labeled"
    labeled.EntityData.BundleName = "cisco_ios_xr"
    labeled.EntityData.ParentYangName = "next-hop"
    labeled.EntityData.SegmentPath = "labeled"
    labeled.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/mac-details/mac-detail/mac-route/route/bmac/path-list/path-list-info/path-list-esi/mac-update-next-hop-array/next-hop/" + labeled.EntityData.SegmentPath
    labeled.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    labeled.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    labeled.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    labeled.EntityData.Children = types.NewOrderedMap()
    labeled.EntityData.Leafs = types.NewOrderedMap()
    labeled.EntityData.Leafs.Append("address-family", types.YLeaf{"AddressFamily", labeled.AddressFamily})
    labeled.EntityData.Leafs.Append("ip-address", types.YLeaf{"IpAddress", labeled.IpAddress})
    labeled.EntityData.Leafs.Append("label", types.YLeaf{"Label", labeled.Label})
    labeled.EntityData.Leafs.Append("internal", types.YLeaf{"Internal", labeled.Internal})

    labeled.EntityData.YListKeys = []string {}

    return &(labeled.EntityData)
}

// L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListMac
// MAC Path List
type L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListMac struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MAC Address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddress interface{}
}

func (pathListMac *L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_PathListInfo_PathListMac) GetEntityData() *types.CommonEntityData {
    pathListMac.EntityData.YFilter = pathListMac.YFilter
    pathListMac.EntityData.YangName = "path-list-mac"
    pathListMac.EntityData.BundleName = "cisco_ios_xr"
    pathListMac.EntityData.ParentYangName = "path-list-info"
    pathListMac.EntityData.SegmentPath = "path-list-mac"
    pathListMac.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/mac-details/mac-detail/mac-route/route/bmac/path-list/path-list-info/" + pathListMac.EntityData.SegmentPath
    pathListMac.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pathListMac.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pathListMac.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pathListMac.EntityData.Children = types.NewOrderedMap()
    pathListMac.EntityData.Leafs = types.NewOrderedMap()
    pathListMac.EntityData.Leafs.Append("mac-address", types.YLeaf{"MacAddress", pathListMac.MacAddress})

    pathListMac.EntityData.YListKeys = []string {}

    return &(pathListMac.EntityData)
}

// L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_NextHopArray
// Array of Next Hops for MAC Path List
type L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_NextHopArray struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Next-hop TOPOLOGY ID. The type is interface{} with range: 0..4294967295.
    TopologyId interface{}

    // Next-hop flags. The type is interface{} with range: 0..65535.
    Flags interface{}

    // Next hop.
    NextHop L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_NextHopArray_NextHop
}

func (nextHopArray *L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_NextHopArray) GetEntityData() *types.CommonEntityData {
    nextHopArray.EntityData.YFilter = nextHopArray.YFilter
    nextHopArray.EntityData.YangName = "next-hop-array"
    nextHopArray.EntityData.BundleName = "cisco_ios_xr"
    nextHopArray.EntityData.ParentYangName = "path-list"
    nextHopArray.EntityData.SegmentPath = "next-hop-array" + types.AddNoKeyToken(nextHopArray)
    nextHopArray.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/mac-details/mac-detail/mac-route/route/bmac/path-list/" + nextHopArray.EntityData.SegmentPath
    nextHopArray.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nextHopArray.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nextHopArray.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nextHopArray.EntityData.Children = types.NewOrderedMap()
    nextHopArray.EntityData.Children.Append("next-hop", types.YChild{"NextHop", &nextHopArray.NextHop})
    nextHopArray.EntityData.Leafs = types.NewOrderedMap()
    nextHopArray.EntityData.Leafs.Append("topology-id", types.YLeaf{"TopologyId", nextHopArray.TopologyId})
    nextHopArray.EntityData.Leafs.Append("flags", types.YLeaf{"Flags", nextHopArray.Flags})

    nextHopArray.EntityData.YListKeys = []string {}

    return &(nextHopArray.EntityData)
}

// L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_NextHopArray_NextHop
// Next hop
type L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_NextHopArray_NextHop struct {
    EntityData types.CommonEntityData
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
    // [a-zA-Z0-9._/-]+.
    InterfaceHandle interface{}

    // XID Next Hop. The type is interface{} with range: 0..4294967295.
    Xid interface{}

    // Labeled Next Hop.
    Labeled L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_NextHopArray_NextHop_Labeled
}

func (nextHop *L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_NextHopArray_NextHop) GetEntityData() *types.CommonEntityData {
    nextHop.EntityData.YFilter = nextHop.YFilter
    nextHop.EntityData.YangName = "next-hop"
    nextHop.EntityData.BundleName = "cisco_ios_xr"
    nextHop.EntityData.ParentYangName = "next-hop-array"
    nextHop.EntityData.SegmentPath = "next-hop"
    nextHop.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/mac-details/mac-detail/mac-route/route/bmac/path-list/next-hop-array/" + nextHop.EntityData.SegmentPath
    nextHop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nextHop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nextHop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nextHop.EntityData.Children = types.NewOrderedMap()
    nextHop.EntityData.Children.Append("labeled", types.YChild{"Labeled", &nextHop.Labeled})
    nextHop.EntityData.Leafs = types.NewOrderedMap()
    nextHop.EntityData.Leafs.Append("type", types.YLeaf{"Type", nextHop.Type})
    nextHop.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", nextHop.Ipv4})
    nextHop.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", nextHop.Ipv6})
    nextHop.EntityData.Leafs.Append("mac", types.YLeaf{"Mac", nextHop.Mac})
    nextHop.EntityData.Leafs.Append("interface-handle", types.YLeaf{"InterfaceHandle", nextHop.InterfaceHandle})
    nextHop.EntityData.Leafs.Append("xid", types.YLeaf{"Xid", nextHop.Xid})

    nextHop.EntityData.YListKeys = []string {}

    return &(nextHop.EntityData)
}

// L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_NextHopArray_NextHop_Labeled
// Labeled Next Hop
type L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_NextHopArray_NextHop_Labeled struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2RIB Address Family. The type is L2ribAfi.
    AddressFamily interface{}

    // IP Address. The type is string.
    IpAddress interface{}

    // Label. The type is interface{} with range: 0..4294967295.
    Label interface{}

    // Internal Label. The type is bool.
    Internal interface{}
}

func (labeled *L2rib_EviChildTables_MacDetails_MacDetail_MacRoute_Route_Bmac_PathList_NextHopArray_NextHop_Labeled) GetEntityData() *types.CommonEntityData {
    labeled.EntityData.YFilter = labeled.YFilter
    labeled.EntityData.YangName = "labeled"
    labeled.EntityData.BundleName = "cisco_ios_xr"
    labeled.EntityData.ParentYangName = "next-hop"
    labeled.EntityData.SegmentPath = "labeled"
    labeled.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/mac-details/mac-detail/mac-route/route/bmac/path-list/next-hop-array/next-hop/" + labeled.EntityData.SegmentPath
    labeled.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    labeled.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    labeled.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    labeled.EntityData.Children = types.NewOrderedMap()
    labeled.EntityData.Leafs = types.NewOrderedMap()
    labeled.EntityData.Leafs.Append("address-family", types.YLeaf{"AddressFamily", labeled.AddressFamily})
    labeled.EntityData.Leafs.Append("ip-address", types.YLeaf{"IpAddress", labeled.IpAddress})
    labeled.EntityData.Leafs.Append("label", types.YLeaf{"Label", labeled.Label})
    labeled.EntityData.Leafs.Append("internal", types.YLeaf{"Internal", labeled.Internal})

    labeled.EntityData.YListKeys = []string {}

    return &(labeled.EntityData)
}

// L2rib_EviChildTables_MacDetails_MacDetail_RtTlv
// Mac Route Opaque Data TLV
type L2rib_EviChildTables_MacDetails_MacDetail_RtTlv struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TLV Type. The type is interface{} with range: 0..65535.
    TlvType interface{}

    // TLV Length. The type is interface{} with range: 0..65535.
    TlvLen interface{}

    // TLV Value. The type is slice of
    // L2rib_EviChildTables_MacDetails_MacDetail_RtTlv_TlvVal.
    TlvVal []*L2rib_EviChildTables_MacDetails_MacDetail_RtTlv_TlvVal
}

func (rtTlv *L2rib_EviChildTables_MacDetails_MacDetail_RtTlv) GetEntityData() *types.CommonEntityData {
    rtTlv.EntityData.YFilter = rtTlv.YFilter
    rtTlv.EntityData.YangName = "rt-tlv"
    rtTlv.EntityData.BundleName = "cisco_ios_xr"
    rtTlv.EntityData.ParentYangName = "mac-detail"
    rtTlv.EntityData.SegmentPath = "rt-tlv"
    rtTlv.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/mac-details/mac-detail/" + rtTlv.EntityData.SegmentPath
    rtTlv.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rtTlv.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rtTlv.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rtTlv.EntityData.Children = types.NewOrderedMap()
    rtTlv.EntityData.Children.Append("tlv-val", types.YChild{"TlvVal", nil})
    for i := range rtTlv.TlvVal {
        types.SetYListKey(rtTlv.TlvVal[i], i)
        rtTlv.EntityData.Children.Append(types.GetSegmentPath(rtTlv.TlvVal[i]), types.YChild{"TlvVal", rtTlv.TlvVal[i]})
    }
    rtTlv.EntityData.Leafs = types.NewOrderedMap()
    rtTlv.EntityData.Leafs.Append("tlv-type", types.YLeaf{"TlvType", rtTlv.TlvType})
    rtTlv.EntityData.Leafs.Append("tlv-len", types.YLeaf{"TlvLen", rtTlv.TlvLen})

    rtTlv.EntityData.YListKeys = []string {}

    return &(rtTlv.EntityData)
}

// L2rib_EviChildTables_MacDetails_MacDetail_RtTlv_TlvVal
// TLV Value
type L2rib_EviChildTables_MacDetails_MacDetail_RtTlv_TlvVal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // The type is interface{} with range: 0..255.
    Entry interface{}
}

func (tlvVal *L2rib_EviChildTables_MacDetails_MacDetail_RtTlv_TlvVal) GetEntityData() *types.CommonEntityData {
    tlvVal.EntityData.YFilter = tlvVal.YFilter
    tlvVal.EntityData.YangName = "tlv-val"
    tlvVal.EntityData.BundleName = "cisco_ios_xr"
    tlvVal.EntityData.ParentYangName = "rt-tlv"
    tlvVal.EntityData.SegmentPath = "tlv-val" + types.AddNoKeyToken(tlvVal)
    tlvVal.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/mac-details/mac-detail/rt-tlv/" + tlvVal.EntityData.SegmentPath
    tlvVal.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tlvVal.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tlvVal.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tlvVal.EntityData.Children = types.NewOrderedMap()
    tlvVal.EntityData.Leafs = types.NewOrderedMap()
    tlvVal.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", tlvVal.Entry})

    tlvVal.EntityData.YListKeys = []string {}

    return &(tlvVal.EntityData)
}

// L2rib_EviChildTables_ImetDetails
// L2RIB EVPN EVI IMET Detail table
type L2rib_EviChildTables_ImetDetails struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2RIB EVPN IMET Detail table. The type is slice of
    // L2rib_EviChildTables_ImetDetails_ImetDetail.
    ImetDetail []*L2rib_EviChildTables_ImetDetails_ImetDetail
}

func (imetDetails *L2rib_EviChildTables_ImetDetails) GetEntityData() *types.CommonEntityData {
    imetDetails.EntityData.YFilter = imetDetails.YFilter
    imetDetails.EntityData.YangName = "imet-details"
    imetDetails.EntityData.BundleName = "cisco_ios_xr"
    imetDetails.EntityData.ParentYangName = "evi-child-tables"
    imetDetails.EntityData.SegmentPath = "imet-details"
    imetDetails.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/" + imetDetails.EntityData.SegmentPath
    imetDetails.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    imetDetails.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    imetDetails.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    imetDetails.EntityData.Children = types.NewOrderedMap()
    imetDetails.EntityData.Children.Append("imet-detail", types.YChild{"ImetDetail", nil})
    for i := range imetDetails.ImetDetail {
        types.SetYListKey(imetDetails.ImetDetail[i], i)
        imetDetails.EntityData.Children.Append(types.GetSegmentPath(imetDetails.ImetDetail[i]), types.YChild{"ImetDetail", imetDetails.ImetDetail[i]})
    }
    imetDetails.EntityData.Leafs = types.NewOrderedMap()

    imetDetails.EntityData.YListKeys = []string {}

    return &(imetDetails.EntityData)
}

// L2rib_EviChildTables_ImetDetails_ImetDetail
// L2RIB EVPN IMET Detail table
type L2rib_EviChildTables_ImetDetails_ImetDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // EVPN ID. The type is interface{} with range: 0..4294967295.
    Evi interface{}

    // Tag ID. The type is interface{} with range: 0..4294967295.
    TagId interface{}

    // Originating Router IP Address. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IpAddr interface{}

    // Admin distance. The type is interface{} with range: 0..4294967295.
    AdminDist interface{}

    // Producer ID. The type is interface{} with range: 0..4294967295.
    ProdId interface{}

    // Tunnel Id. The type is string.
    TunnelId interface{}

    // IMET route flags. The type is interface{} with range: 0..4294967295.
    Flags interface{}

    // IMET route tunnel type. The type is interface{} with range: 0..4294967295.
    TunnelType interface{}

    // Label. The type is interface{} with range: 0..4294967295.
    L2rLabel interface{}

    // Encap Type. The type is interface{} with range: 0..4294967295.
    EncapType interface{}

    // Last Update Timestamp. The type is interface{} with range:
    // 0..18446744073709551615.
    LastUpdateTimestamp interface{}

    // Imet Route base information.
    ImetRouteBase L2rib_EviChildTables_ImetDetails_ImetDetail_ImetRouteBase
}

func (imetDetail *L2rib_EviChildTables_ImetDetails_ImetDetail) GetEntityData() *types.CommonEntityData {
    imetDetail.EntityData.YFilter = imetDetail.YFilter
    imetDetail.EntityData.YangName = "imet-detail"
    imetDetail.EntityData.BundleName = "cisco_ios_xr"
    imetDetail.EntityData.ParentYangName = "imet-details"
    imetDetail.EntityData.SegmentPath = "imet-detail" + types.AddNoKeyToken(imetDetail)
    imetDetail.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/imet-details/" + imetDetail.EntityData.SegmentPath
    imetDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    imetDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    imetDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    imetDetail.EntityData.Children = types.NewOrderedMap()
    imetDetail.EntityData.Children.Append("imet-route-base", types.YChild{"ImetRouteBase", &imetDetail.ImetRouteBase})
    imetDetail.EntityData.Leafs = types.NewOrderedMap()
    imetDetail.EntityData.Leafs.Append("evi", types.YLeaf{"Evi", imetDetail.Evi})
    imetDetail.EntityData.Leafs.Append("tag-id", types.YLeaf{"TagId", imetDetail.TagId})
    imetDetail.EntityData.Leafs.Append("ip-addr", types.YLeaf{"IpAddr", imetDetail.IpAddr})
    imetDetail.EntityData.Leafs.Append("admin-dist", types.YLeaf{"AdminDist", imetDetail.AdminDist})
    imetDetail.EntityData.Leafs.Append("prod-id", types.YLeaf{"ProdId", imetDetail.ProdId})
    imetDetail.EntityData.Leafs.Append("tunnel-id", types.YLeaf{"TunnelId", imetDetail.TunnelId})
    imetDetail.EntityData.Leafs.Append("flags", types.YLeaf{"Flags", imetDetail.Flags})
    imetDetail.EntityData.Leafs.Append("tunnel-type", types.YLeaf{"TunnelType", imetDetail.TunnelType})
    imetDetail.EntityData.Leafs.Append("l2r-label", types.YLeaf{"L2rLabel", imetDetail.L2rLabel})
    imetDetail.EntityData.Leafs.Append("encap-type", types.YLeaf{"EncapType", imetDetail.EncapType})
    imetDetail.EntityData.Leafs.Append("last-update-timestamp", types.YLeaf{"LastUpdateTimestamp", imetDetail.LastUpdateTimestamp})

    imetDetail.EntityData.YListKeys = []string {}

    return &(imetDetail.EntityData)
}

// L2rib_EviChildTables_ImetDetails_ImetDetail_ImetRouteBase
// Imet Route base information
type L2rib_EviChildTables_ImetDetails_ImetDetail_ImetRouteBase struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Originating Router IP Address. The type is string.
    VtepiPaddr interface{}

    // Admin Distance. The type is interface{} with range: 0..255.
    AdminDistance interface{}

    // Producer of Imet route. The type is interface{} with range: 0..255.
    ProducerId interface{}

    // Topo ID. The type is interface{} with range: 0..4294967295.
    TopoId interface{}

    // Ethernet Tag ID. The type is interface{} with range: 0..4294967295.
    EthernetTagId interface{}
}

func (imetRouteBase *L2rib_EviChildTables_ImetDetails_ImetDetail_ImetRouteBase) GetEntityData() *types.CommonEntityData {
    imetRouteBase.EntityData.YFilter = imetRouteBase.YFilter
    imetRouteBase.EntityData.YangName = "imet-route-base"
    imetRouteBase.EntityData.BundleName = "cisco_ios_xr"
    imetRouteBase.EntityData.ParentYangName = "imet-detail"
    imetRouteBase.EntityData.SegmentPath = "imet-route-base"
    imetRouteBase.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evi-child-tables/imet-details/imet-detail/" + imetRouteBase.EntityData.SegmentPath
    imetRouteBase.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    imetRouteBase.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    imetRouteBase.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    imetRouteBase.EntityData.Children = types.NewOrderedMap()
    imetRouteBase.EntityData.Leafs = types.NewOrderedMap()
    imetRouteBase.EntityData.Leafs.Append("vtepi-paddr", types.YLeaf{"VtepiPaddr", imetRouteBase.VtepiPaddr})
    imetRouteBase.EntityData.Leafs.Append("admin-distance", types.YLeaf{"AdminDistance", imetRouteBase.AdminDistance})
    imetRouteBase.EntityData.Leafs.Append("producer-id", types.YLeaf{"ProducerId", imetRouteBase.ProducerId})
    imetRouteBase.EntityData.Leafs.Append("topo-id", types.YLeaf{"TopoId", imetRouteBase.TopoId})
    imetRouteBase.EntityData.Leafs.Append("ethernet-tag-id", types.YLeaf{"EthernetTagId", imetRouteBase.EthernetTagId})

    imetRouteBase.EntityData.YListKeys = []string {}

    return &(imetRouteBase.EntityData)
}

// L2rib_Evis
// L2RIB EVPN EVI Table
type L2rib_Evis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2RIB EVPN EVI Entry. The type is slice of L2rib_Evis_Evi.
    Evi []*L2rib_Evis_Evi
}

func (evis *L2rib_Evis) GetEntityData() *types.CommonEntityData {
    evis.EntityData.YFilter = evis.YFilter
    evis.EntityData.YangName = "evis"
    evis.EntityData.BundleName = "cisco_ios_xr"
    evis.EntityData.ParentYangName = "l2rib"
    evis.EntityData.SegmentPath = "evis"
    evis.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/" + evis.EntityData.SegmentPath
    evis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evis.EntityData.Children = types.NewOrderedMap()
    evis.EntityData.Children.Append("evi", types.YChild{"Evi", nil})
    for i := range evis.Evi {
        evis.EntityData.Children.Append(types.GetSegmentPath(evis.Evi[i]), types.YChild{"Evi", evis.Evi[i]})
    }
    evis.EntityData.Leafs = types.NewOrderedMap()

    evis.EntityData.YListKeys = []string {}

    return &(evis.EntityData)
}

// L2rib_Evis_Evi
// L2RIB EVPN EVI Entry
type L2rib_Evis_Evi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. EVI ID. The type is interface{} with range:
    // 0..4294967295.
    Evi interface{}

    // Topology ID. The type is interface{} with range: 0..4294967295.
    TopologyId interface{}

    // Topology Name. The type is string.
    TopologyName interface{}

    // Topology Type. The type is interface{} with range: 0..4294967295.
    TopologyType interface{}
}

func (evi *L2rib_Evis_Evi) GetEntityData() *types.CommonEntityData {
    evi.EntityData.YFilter = evi.YFilter
    evi.EntityData.YangName = "evi"
    evi.EntityData.BundleName = "cisco_ios_xr"
    evi.EntityData.ParentYangName = "evis"
    evi.EntityData.SegmentPath = "evi" + types.AddKeyToken(evi.Evi, "evi")
    evi.EntityData.AbsolutePath = "Cisco-IOS-XR-l2rib-oper:l2rib/evis/" + evi.EntityData.SegmentPath
    evi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evi.EntityData.Children = types.NewOrderedMap()
    evi.EntityData.Leafs = types.NewOrderedMap()
    evi.EntityData.Leafs.Append("evi", types.YLeaf{"Evi", evi.Evi})
    evi.EntityData.Leafs.Append("topology-id", types.YLeaf{"TopologyId", evi.TopologyId})
    evi.EntityData.Leafs.Append("topology-name", types.YLeaf{"TopologyName", evi.TopologyName})
    evi.EntityData.Leafs.Append("topology-type", types.YLeaf{"TopologyType", evi.TopologyType})

    evi.EntityData.YListKeys = []string {"Evi"}

    return &(evi.EntityData)
}

