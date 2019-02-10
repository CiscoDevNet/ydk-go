// This module contains a collection of YANG
// definitions for Cisco IOS-XR SysAdmin configuration.
// 
// Copyright(c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package sysadmin_rvm_mgr

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package sysadmin_rvm_mgr"))
    ydk.RegisterEntity("{http://www.cisco.com/ns/yang/Cisco-IOS-XR-sysadmin-rvm-mgr RVM}", reflect.TypeOf(RVM{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sysadmin-rvm-mgr:RVM", reflect.TypeOf(RVM{}))
}

// NodetypeTd
type NodetypeTd string

const (
    NodetypeTd_sysadmin NodetypeTd = "sysadmin"

    NodetypeTd_service NodetypeTd = "service"
)

// FlagtypeTd
type FlagtypeTd string

const (
    FlagtypeTd_clear FlagtypeTd = "clear"

    FlagtypeTd_set FlagtypeTd = "set"
)

// RVM
// RVM Manager Info
type RVM struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of RVM_AllLocations.
    AllLocations []*RVM_AllLocations
}

func (rVM *RVM) GetEntityData() *types.CommonEntityData {
    rVM.EntityData.YFilter = rVM.YFilter
    rVM.EntityData.YangName = "RVM"
    rVM.EntityData.BundleName = "cisco_ios_xr"
    rVM.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-rvm-mgr"
    rVM.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-rvm-mgr:RVM"
    rVM.EntityData.AbsolutePath = rVM.EntityData.SegmentPath
    rVM.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rVM.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rVM.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rVM.EntityData.Children = types.NewOrderedMap()
    rVM.EntityData.Children.Append("all-locations", types.YChild{"AllLocations", nil})
    for i := range rVM.AllLocations {
        rVM.EntityData.Children.Append(types.GetSegmentPath(rVM.AllLocations[i]), types.YChild{"AllLocations", rVM.AllLocations[i]})
    }
    rVM.EntityData.Leafs = types.NewOrderedMap()

    rVM.EntityData.YListKeys = []string {}

    return &(rVM.EntityData)
}

// RVM_AllLocations
type RVM_AllLocations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    Location interface{}

    // RVM Manager Info.
    Objects RVM_AllLocations_Objects

    // RVM Manager Node Info.
    Node RVM_AllLocations_Node

    // RVM Manager Card Info.
    Card RVM_AllLocations_Card
}

func (allLocations *RVM_AllLocations) GetEntityData() *types.CommonEntityData {
    allLocations.EntityData.YFilter = allLocations.YFilter
    allLocations.EntityData.YangName = "all-locations"
    allLocations.EntityData.BundleName = "cisco_ios_xr"
    allLocations.EntityData.ParentYangName = "RVM"
    allLocations.EntityData.SegmentPath = "all-locations" + types.AddKeyToken(allLocations.Location, "location")
    allLocations.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-rvm-mgr:RVM/" + allLocations.EntityData.SegmentPath
    allLocations.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allLocations.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allLocations.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allLocations.EntityData.Children = types.NewOrderedMap()
    allLocations.EntityData.Children.Append("objects", types.YChild{"Objects", &allLocations.Objects})
    allLocations.EntityData.Children.Append("node", types.YChild{"Node", &allLocations.Node})
    allLocations.EntityData.Children.Append("card", types.YChild{"Card", &allLocations.Card})
    allLocations.EntityData.Leafs = types.NewOrderedMap()
    allLocations.EntityData.Leafs.Append("location", types.YLeaf{"Location", allLocations.Location})

    allLocations.EntityData.YListKeys = []string {"Location"}

    return &(allLocations.EntityData)
}

// RVM_AllLocations_Objects
// RVM Manager Info
type RVM_AllLocations_Objects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of RVM_AllLocations_Objects_AllObjs.
    AllObjs []*RVM_AllLocations_Objects_AllObjs
}

func (objects *RVM_AllLocations_Objects) GetEntityData() *types.CommonEntityData {
    objects.EntityData.YFilter = objects.YFilter
    objects.EntityData.YangName = "objects"
    objects.EntityData.BundleName = "cisco_ios_xr"
    objects.EntityData.ParentYangName = "all-locations"
    objects.EntityData.SegmentPath = "objects"
    objects.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-rvm-mgr:RVM/all-locations/" + objects.EntityData.SegmentPath
    objects.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    objects.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    objects.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    objects.EntityData.Children = types.NewOrderedMap()
    objects.EntityData.Children.Append("all-objs", types.YChild{"AllObjs", nil})
    for i := range objects.AllObjs {
        objects.EntityData.Children.Append(types.GetSegmentPath(objects.AllObjs[i]), types.YChild{"AllObjs", objects.AllObjs[i]})
    }
    objects.EntityData.Leafs = types.NewOrderedMap()

    objects.EntityData.YListKeys = []string {}

    return &(objects.EntityData)
}

// RVM_AllLocations_Objects_AllObjs
type RVM_AllLocations_Objects_AllObjs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Index into obj_db array. The type is interface{}
    // with range: 0..4294967295.
    Index interface{}

    // Number of allocated nodes. The type is interface{} with range:
    // 0..4294967295.
    NumAllocated interface{}

    // Number of freed nodes. The type is interface{} with range: 0..4294967295.
    NumFreed interface{}

    // Number of current object nodes. The type is interface{} with range:
    // 0..4294967295.
    NumObjects interface{}
}

func (allObjs *RVM_AllLocations_Objects_AllObjs) GetEntityData() *types.CommonEntityData {
    allObjs.EntityData.YFilter = allObjs.YFilter
    allObjs.EntityData.YangName = "all-objs"
    allObjs.EntityData.BundleName = "cisco_ios_xr"
    allObjs.EntityData.ParentYangName = "objects"
    allObjs.EntityData.SegmentPath = "all-objs" + types.AddKeyToken(allObjs.Index, "index")
    allObjs.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-rvm-mgr:RVM/all-locations/objects/" + allObjs.EntityData.SegmentPath
    allObjs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allObjs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allObjs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allObjs.EntityData.Children = types.NewOrderedMap()
    allObjs.EntityData.Leafs = types.NewOrderedMap()
    allObjs.EntityData.Leafs.Append("index", types.YLeaf{"Index", allObjs.Index})
    allObjs.EntityData.Leafs.Append("num_allocated", types.YLeaf{"NumAllocated", allObjs.NumAllocated})
    allObjs.EntityData.Leafs.Append("num_freed", types.YLeaf{"NumFreed", allObjs.NumFreed})
    allObjs.EntityData.Leafs.Append("num_objects", types.YLeaf{"NumObjects", allObjs.NumObjects})

    allObjs.EntityData.YListKeys = []string {"Index"}

    return &(allObjs.EntityData)
}

// RVM_AllLocations_Node
// RVM Manager Node Info
type RVM_AllLocations_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of RVM_AllLocations_Node_AllNodes.
    AllNodes []*RVM_AllLocations_Node_AllNodes
}

func (node *RVM_AllLocations_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "all-locations"
    node.EntityData.SegmentPath = "node"
    node.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-rvm-mgr:RVM/all-locations/" + node.EntityData.SegmentPath
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("all-nodes", types.YChild{"AllNodes", nil})
    for i := range node.AllNodes {
        node.EntityData.Children.Append(types.GetSegmentPath(node.AllNodes[i]), types.YChild{"AllNodes", node.AllNodes[i]})
    }
    node.EntityData.Leafs = types.NewOrderedMap()

    node.EntityData.YListKeys = []string {}

    return &(node.EntityData)
}

// RVM_AllLocations_Node_AllNodes
type RVM_AllLocations_Node_AllNodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. IP address of the node. The type is interface{}
    // with range: 0..4294967295.
    Ipv4Addr interface{}

    // IP address of the node in string format. The type is string.
    Ipv4AddrStr interface{}

    // Node type state and flag. The type is string.
    NodeInfo interface{}

    // Node heartbeat related info. The type is string.
    NodeHbInfo interface{}

    // Card info the node belongs to. The type is string.
    NodeCardInfo interface{}
}

func (allNodes *RVM_AllLocations_Node_AllNodes) GetEntityData() *types.CommonEntityData {
    allNodes.EntityData.YFilter = allNodes.YFilter
    allNodes.EntityData.YangName = "all-nodes"
    allNodes.EntityData.BundleName = "cisco_ios_xr"
    allNodes.EntityData.ParentYangName = "node"
    allNodes.EntityData.SegmentPath = "all-nodes" + types.AddKeyToken(allNodes.Ipv4Addr, "ipv4_addr")
    allNodes.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-rvm-mgr:RVM/all-locations/node/" + allNodes.EntityData.SegmentPath
    allNodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allNodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allNodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allNodes.EntityData.Children = types.NewOrderedMap()
    allNodes.EntityData.Leafs = types.NewOrderedMap()
    allNodes.EntityData.Leafs.Append("ipv4_addr", types.YLeaf{"Ipv4Addr", allNodes.Ipv4Addr})
    allNodes.EntityData.Leafs.Append("ipv4_addr_str", types.YLeaf{"Ipv4AddrStr", allNodes.Ipv4AddrStr})
    allNodes.EntityData.Leafs.Append("node_info", types.YLeaf{"NodeInfo", allNodes.NodeInfo})
    allNodes.EntityData.Leafs.Append("node_hb_info", types.YLeaf{"NodeHbInfo", allNodes.NodeHbInfo})
    allNodes.EntityData.Leafs.Append("node_card_info", types.YLeaf{"NodeCardInfo", allNodes.NodeCardInfo})

    allNodes.EntityData.YListKeys = []string {"Ipv4Addr"}

    return &(allNodes.EntityData)
}

// RVM_AllLocations_Card
// RVM Manager Card Info
type RVM_AllLocations_Card struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of RVM_AllLocations_Card_AllCards.
    AllCards []*RVM_AllLocations_Card_AllCards
}

func (card *RVM_AllLocations_Card) GetEntityData() *types.CommonEntityData {
    card.EntityData.YFilter = card.YFilter
    card.EntityData.YangName = "card"
    card.EntityData.BundleName = "cisco_ios_xr"
    card.EntityData.ParentYangName = "all-locations"
    card.EntityData.SegmentPath = "card"
    card.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-rvm-mgr:RVM/all-locations/" + card.EntityData.SegmentPath
    card.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    card.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    card.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    card.EntityData.Children = types.NewOrderedMap()
    card.EntityData.Children.Append("all-cards", types.YChild{"AllCards", nil})
    for i := range card.AllCards {
        card.EntityData.Children.Append(types.GetSegmentPath(card.AllCards[i]), types.YChild{"AllCards", card.AllCards[i]})
    }
    card.EntityData.Leafs = types.NewOrderedMap()

    card.EntityData.YListKeys = []string {}

    return &(card.EntityData)
}

// RVM_AllLocations_Card_AllCards
type RVM_AllLocations_Card_AllCards struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Serial number of the card. The type is string.
    SerialNum interface{}

    // Card flags. The type is string.
    CardFlag interface{}

    // Card rack and slot num. The type is string.
    CardInfo interface{}

    // Sysadmin nodes on this card. The type is string.
    SysadminNodes interface{}

    // XR nodes on this card. The type is string.
    XrNodes interface{}
}

func (allCards *RVM_AllLocations_Card_AllCards) GetEntityData() *types.CommonEntityData {
    allCards.EntityData.YFilter = allCards.YFilter
    allCards.EntityData.YangName = "all-cards"
    allCards.EntityData.BundleName = "cisco_ios_xr"
    allCards.EntityData.ParentYangName = "card"
    allCards.EntityData.SegmentPath = "all-cards" + types.AddKeyToken(allCards.SerialNum, "serial_num")
    allCards.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-rvm-mgr:RVM/all-locations/card/" + allCards.EntityData.SegmentPath
    allCards.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allCards.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allCards.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allCards.EntityData.Children = types.NewOrderedMap()
    allCards.EntityData.Leafs = types.NewOrderedMap()
    allCards.EntityData.Leafs.Append("serial_num", types.YLeaf{"SerialNum", allCards.SerialNum})
    allCards.EntityData.Leafs.Append("card_flag", types.YLeaf{"CardFlag", allCards.CardFlag})
    allCards.EntityData.Leafs.Append("card_info", types.YLeaf{"CardInfo", allCards.CardInfo})
    allCards.EntityData.Leafs.Append("sysadmin_nodes", types.YLeaf{"SysadminNodes", allCards.SysadminNodes})
    allCards.EntityData.Leafs.Append("xr_nodes", types.YLeaf{"XrNodes", allCards.XrNodes})

    allCards.EntityData.YListKeys = []string {"SerialNum"}

    return &(allCards.EntityData)
}

