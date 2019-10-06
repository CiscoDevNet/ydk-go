// This module contains definitions
// for the Calvados model objects.
// 
// This module contains a collection of YANG
// definitions for Cisco IOS-XR SysAdmin configuration.
// 
// The System Admin Manager (CM)
// 
// Copyright(c) 2010-2017 by Cisco Systems, Inc.
// All rights reserved.
// 
// Copyright (c) 2012-2018 by Cisco Systems, Inc.
// All rights reserved.
package sysadmin_cm

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package sysadmin_cm"))
    ydk.RegisterEntity("{http://www.cisco.com/ns/yang/Cisco-IOS-XR-sysadmin-cm node-inventory}", reflect.TypeOf(NodeInventory{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sysadmin-cm:node-inventory", reflect.TypeOf(NodeInventory{}))
    ydk.RegisterEntity("{http://www.cisco.com/ns/yang/Cisco-IOS-XR-sysadmin-cm card-inventory}", reflect.TypeOf(CardInventory{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sysadmin-cm:card-inventory", reflect.TypeOf(CardInventory{}))
    ydk.RegisterEntity("{http://www.cisco.com/ns/yang/Cisco-IOS-XR-sysadmin-cm rack-inventory}", reflect.TypeOf(RackInventory{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sysadmin-cm:rack-inventory", reflect.TypeOf(RackInventory{}))
    ydk.RegisterEntity("{http://www.cisco.com/ns/yang/Cisco-IOS-XR-sysadmin-cm system-service-inventory}", reflect.TypeOf(SystemServiceInventory{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sysadmin-cm:system-service-inventory", reflect.TypeOf(SystemServiceInventory{}))
    ydk.RegisterEntity("{http://www.cisco.com/ns/yang/Cisco-IOS-XR-sysadmin-cm rack-service-inventory}", reflect.TypeOf(RackServiceInventory{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sysadmin-cm:rack-service-inventory", reflect.TypeOf(RackServiceInventory{}))
    ydk.RegisterEntity("{http://www.cisco.com/ns/yang/Cisco-IOS-XR-sysadmin-cm sdr-inventory}", reflect.TypeOf(SdrInventory{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sysadmin-cm:sdr-inventory", reflect.TypeOf(SdrInventory{}))
    ydk.RegisterEntity("{http://www.cisco.com/ns/yang/Cisco-IOS-XR-sysadmin-cm leader-statistics}", reflect.TypeOf(LeaderStatistics{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sysadmin-cm:leader-statistics", reflect.TypeOf(LeaderStatistics{}))
    ydk.RegisterEntity("{http://www.cisco.com/ns/yang/Cisco-IOS-XR-sysadmin-cm topology-neighbors}", reflect.TypeOf(TopologyNeighbors{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sysadmin-cm:topology-neighbors", reflect.TypeOf(TopologyNeighbors{}))
    ydk.RegisterEntity("{http://www.cisco.com/ns/yang/Cisco-IOS-XR-sysadmin-cm placement}", reflect.TypeOf(Placement{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sysadmin-cm:placement", reflect.TypeOf(Placement{}))
}

// AreaType
type AreaType string

const (
    AreaType_SYSTEM AreaType = "SYSTEM"

    AreaType_RACK AreaType = "RACK"

    AreaType_UNKNOWN AreaType = "UNKNOWN"
)

// NodeInventory
// System Admin Manager Node Inventory. All accesses are
// read-only. CLI show command looks like show node-inventory
// location <location>
type NodeInventory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // System Admin Manager Node Inventory. All accesses are read-only. CLI show
    // command looks like show node-inventory location <location>.
    Summary NodeInventory_Summary

    // System Admin Manager Node Inventory. All accesses are read-only. CLI show
    // command looks like show node-inventory location <location>.
    Detail NodeInventory_Detail
}

func (nodeInventory *NodeInventory) GetEntityData() *types.CommonEntityData {
    nodeInventory.EntityData.YFilter = nodeInventory.YFilter
    nodeInventory.EntityData.YangName = "node-inventory"
    nodeInventory.EntityData.BundleName = "cisco_ios_xr"
    nodeInventory.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-cm"
    nodeInventory.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-cm:node-inventory"
    nodeInventory.EntityData.AbsolutePath = nodeInventory.EntityData.SegmentPath
    nodeInventory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeInventory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeInventory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeInventory.EntityData.Children = types.NewOrderedMap()
    nodeInventory.EntityData.Children.Append("summary", types.YChild{"Summary", &nodeInventory.Summary})
    nodeInventory.EntityData.Children.Append("detail", types.YChild{"Detail", &nodeInventory.Detail})
    nodeInventory.EntityData.Leafs = types.NewOrderedMap()

    nodeInventory.EntityData.YListKeys = []string {}

    return &(nodeInventory.EntityData)
}

// NodeInventory_Summary
// System Admin Manager Node Inventory. All accesses are
// read-only. CLI show command looks like show node-inventory
// location <location>
type NodeInventory_Summary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of NodeInventory_Summary_NodeLocations.
    NodeLocations []*NodeInventory_Summary_NodeLocations
}

func (summary *NodeInventory_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "cisco_ios_xr"
    summary.EntityData.ParentYangName = "node-inventory"
    summary.EntityData.SegmentPath = "summary"
    summary.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-cm:node-inventory/" + summary.EntityData.SegmentPath
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = types.NewOrderedMap()
    summary.EntityData.Children.Append("node_locations", types.YChild{"NodeLocations", nil})
    for i := range summary.NodeLocations {
        summary.EntityData.Children.Append(types.GetSegmentPath(summary.NodeLocations[i]), types.YChild{"NodeLocations", summary.NodeLocations[i]})
    }
    summary.EntityData.Leafs = types.NewOrderedMap()

    summary.EntityData.YListKeys = []string {}

    return &(summary.EntityData)
}

// NodeInventory_Summary_NodeLocations
type NodeInventory_Summary_NodeLocations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    NodeLocation interface{}

    // The type is slice of NodeInventory_Summary_NodeLocations_Nodei.
    Nodei []*NodeInventory_Summary_NodeLocations_Nodei
}

func (nodeLocations *NodeInventory_Summary_NodeLocations) GetEntityData() *types.CommonEntityData {
    nodeLocations.EntityData.YFilter = nodeLocations.YFilter
    nodeLocations.EntityData.YangName = "node_locations"
    nodeLocations.EntityData.BundleName = "cisco_ios_xr"
    nodeLocations.EntityData.ParentYangName = "summary"
    nodeLocations.EntityData.SegmentPath = "node_locations" + types.AddKeyToken(nodeLocations.NodeLocation, "node_location")
    nodeLocations.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-cm:node-inventory/summary/" + nodeLocations.EntityData.SegmentPath
    nodeLocations.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeLocations.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeLocations.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeLocations.EntityData.Children = types.NewOrderedMap()
    nodeLocations.EntityData.Children.Append("nodei", types.YChild{"Nodei", nil})
    for i := range nodeLocations.Nodei {
        nodeLocations.EntityData.Children.Append(types.GetSegmentPath(nodeLocations.Nodei[i]), types.YChild{"Nodei", nodeLocations.Nodei[i]})
    }
    nodeLocations.EntityData.Leafs = types.NewOrderedMap()
    nodeLocations.EntityData.Leafs.Append("node_location", types.YLeaf{"NodeLocation", nodeLocations.NodeLocation})

    nodeLocations.EntityData.YListKeys = []string {"NodeLocation"}

    return &(nodeLocations.EntityData)
}

// NodeInventory_Summary_NodeLocations_Nodei
type NodeInventory_Summary_NodeLocations_Nodei struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. IP address of the node. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IpAddress interface{}

    // Node Type. The type is string.
    Type interface{}

    // Node MAC. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddress interface{}

    // Card serial# the node belongs to. The type is string.
    CardSerial interface{}

    // Node NTI. The type is interface{} with range: 0..4294967295.
    Nti interface{}
}

func (nodei *NodeInventory_Summary_NodeLocations_Nodei) GetEntityData() *types.CommonEntityData {
    nodei.EntityData.YFilter = nodei.YFilter
    nodei.EntityData.YangName = "nodei"
    nodei.EntityData.BundleName = "cisco_ios_xr"
    nodei.EntityData.ParentYangName = "node_locations"
    nodei.EntityData.SegmentPath = "nodei" + types.AddKeyToken(nodei.IpAddress, "ip_address")
    nodei.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-cm:node-inventory/summary/node_locations/" + nodei.EntityData.SegmentPath
    nodei.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodei.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodei.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodei.EntityData.Children = types.NewOrderedMap()
    nodei.EntityData.Leafs = types.NewOrderedMap()
    nodei.EntityData.Leafs.Append("ip_address", types.YLeaf{"IpAddress", nodei.IpAddress})
    nodei.EntityData.Leafs.Append("type", types.YLeaf{"Type", nodei.Type})
    nodei.EntityData.Leafs.Append("mac_address", types.YLeaf{"MacAddress", nodei.MacAddress})
    nodei.EntityData.Leafs.Append("card_serial", types.YLeaf{"CardSerial", nodei.CardSerial})
    nodei.EntityData.Leafs.Append("nti", types.YLeaf{"Nti", nodei.Nti})

    nodei.EntityData.YListKeys = []string {"IpAddress"}

    return &(nodei.EntityData)
}

// NodeInventory_Detail
// System Admin Manager Node Inventory. All accesses are
// read-only. CLI show command looks like show node-inventory
// location <location>
type NodeInventory_Detail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of NodeInventory_Detail_NodeLocations.
    NodeLocations []*NodeInventory_Detail_NodeLocations
}

func (detail *NodeInventory_Detail) GetEntityData() *types.CommonEntityData {
    detail.EntityData.YFilter = detail.YFilter
    detail.EntityData.YangName = "detail"
    detail.EntityData.BundleName = "cisco_ios_xr"
    detail.EntityData.ParentYangName = "node-inventory"
    detail.EntityData.SegmentPath = "detail"
    detail.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-cm:node-inventory/" + detail.EntityData.SegmentPath
    detail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detail.EntityData.Children = types.NewOrderedMap()
    detail.EntityData.Children.Append("node_locations", types.YChild{"NodeLocations", nil})
    for i := range detail.NodeLocations {
        detail.EntityData.Children.Append(types.GetSegmentPath(detail.NodeLocations[i]), types.YChild{"NodeLocations", detail.NodeLocations[i]})
    }
    detail.EntityData.Leafs = types.NewOrderedMap()

    detail.EntityData.YListKeys = []string {}

    return &(detail.EntityData)
}

// NodeInventory_Detail_NodeLocations
type NodeInventory_Detail_NodeLocations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    NodeLocation interface{}

    // The type is slice of NodeInventory_Detail_NodeLocations_Nodei.
    Nodei []*NodeInventory_Detail_NodeLocations_Nodei
}

func (nodeLocations *NodeInventory_Detail_NodeLocations) GetEntityData() *types.CommonEntityData {
    nodeLocations.EntityData.YFilter = nodeLocations.YFilter
    nodeLocations.EntityData.YangName = "node_locations"
    nodeLocations.EntityData.BundleName = "cisco_ios_xr"
    nodeLocations.EntityData.ParentYangName = "detail"
    nodeLocations.EntityData.SegmentPath = "node_locations" + types.AddKeyToken(nodeLocations.NodeLocation, "node_location")
    nodeLocations.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-cm:node-inventory/detail/" + nodeLocations.EntityData.SegmentPath
    nodeLocations.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeLocations.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeLocations.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeLocations.EntityData.Children = types.NewOrderedMap()
    nodeLocations.EntityData.Children.Append("nodei", types.YChild{"Nodei", nil})
    for i := range nodeLocations.Nodei {
        nodeLocations.EntityData.Children.Append(types.GetSegmentPath(nodeLocations.Nodei[i]), types.YChild{"Nodei", nodeLocations.Nodei[i]})
    }
    nodeLocations.EntityData.Leafs = types.NewOrderedMap()
    nodeLocations.EntityData.Leafs.Append("node_location", types.YLeaf{"NodeLocation", nodeLocations.NodeLocation})

    nodeLocations.EntityData.YListKeys = []string {"NodeLocation"}

    return &(nodeLocations.EntityData)
}

// NodeInventory_Detail_NodeLocations_Nodei
type NodeInventory_Detail_NodeLocations_Nodei struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. IP address of the node. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IpAddress interface{}

    // Node Type. The type is string.
    Type interface{}

    // Node MAC. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddress interface{}

    // Card serial# the node belongs to. The type is string.
    CardSerial interface{}

    // Node NTI. The type is interface{} with range: 0..4294967295.
    Nti interface{}

    // Node in Restart. The type is bool.
    Restart interface{}

    // Node Data. The type is string.
    Data interface{}

    // SDR Name the node belongs to. The type is string.
    Sdr interface{}
}

func (nodei *NodeInventory_Detail_NodeLocations_Nodei) GetEntityData() *types.CommonEntityData {
    nodei.EntityData.YFilter = nodei.YFilter
    nodei.EntityData.YangName = "nodei"
    nodei.EntityData.BundleName = "cisco_ios_xr"
    nodei.EntityData.ParentYangName = "node_locations"
    nodei.EntityData.SegmentPath = "nodei" + types.AddKeyToken(nodei.IpAddress, "ip_address")
    nodei.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-cm:node-inventory/detail/node_locations/" + nodei.EntityData.SegmentPath
    nodei.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodei.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodei.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodei.EntityData.Children = types.NewOrderedMap()
    nodei.EntityData.Leafs = types.NewOrderedMap()
    nodei.EntityData.Leafs.Append("ip_address", types.YLeaf{"IpAddress", nodei.IpAddress})
    nodei.EntityData.Leafs.Append("type", types.YLeaf{"Type", nodei.Type})
    nodei.EntityData.Leafs.Append("mac_address", types.YLeaf{"MacAddress", nodei.MacAddress})
    nodei.EntityData.Leafs.Append("card_serial", types.YLeaf{"CardSerial", nodei.CardSerial})
    nodei.EntityData.Leafs.Append("nti", types.YLeaf{"Nti", nodei.Nti})
    nodei.EntityData.Leafs.Append("restart", types.YLeaf{"Restart", nodei.Restart})
    nodei.EntityData.Leafs.Append("data", types.YLeaf{"Data", nodei.Data})
    nodei.EntityData.Leafs.Append("sdr", types.YLeaf{"Sdr", nodei.Sdr})

    nodei.EntityData.YListKeys = []string {"IpAddress"}

    return &(nodei.EntityData)
}

// CardInventory
// System Admin Manager Card Inventory. All accesses are
// read-only. CLI show command looks like show card-inventory
// location <location>
type CardInventory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of CardInventory_CardLocations.
    CardLocations []*CardInventory_CardLocations
}

func (cardInventory *CardInventory) GetEntityData() *types.CommonEntityData {
    cardInventory.EntityData.YFilter = cardInventory.YFilter
    cardInventory.EntityData.YangName = "card-inventory"
    cardInventory.EntityData.BundleName = "cisco_ios_xr"
    cardInventory.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-cm"
    cardInventory.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-cm:card-inventory"
    cardInventory.EntityData.AbsolutePath = cardInventory.EntityData.SegmentPath
    cardInventory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cardInventory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cardInventory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cardInventory.EntityData.Children = types.NewOrderedMap()
    cardInventory.EntityData.Children.Append("card_locations", types.YChild{"CardLocations", nil})
    for i := range cardInventory.CardLocations {
        cardInventory.EntityData.Children.Append(types.GetSegmentPath(cardInventory.CardLocations[i]), types.YChild{"CardLocations", cardInventory.CardLocations[i]})
    }
    cardInventory.EntityData.Leafs = types.NewOrderedMap()

    cardInventory.EntityData.YListKeys = []string {}

    return &(cardInventory.EntityData)
}

// CardInventory_CardLocations
type CardInventory_CardLocations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    CardLocation interface{}

    // The type is slice of CardInventory_CardLocations_Cardi.
    Cardi []*CardInventory_CardLocations_Cardi
}

func (cardLocations *CardInventory_CardLocations) GetEntityData() *types.CommonEntityData {
    cardLocations.EntityData.YFilter = cardLocations.YFilter
    cardLocations.EntityData.YangName = "card_locations"
    cardLocations.EntityData.BundleName = "cisco_ios_xr"
    cardLocations.EntityData.ParentYangName = "card-inventory"
    cardLocations.EntityData.SegmentPath = "card_locations" + types.AddKeyToken(cardLocations.CardLocation, "card_location")
    cardLocations.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-cm:card-inventory/" + cardLocations.EntityData.SegmentPath
    cardLocations.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cardLocations.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cardLocations.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cardLocations.EntityData.Children = types.NewOrderedMap()
    cardLocations.EntityData.Children.Append("cardi", types.YChild{"Cardi", nil})
    for i := range cardLocations.Cardi {
        cardLocations.EntityData.Children.Append(types.GetSegmentPath(cardLocations.Cardi[i]), types.YChild{"Cardi", cardLocations.Cardi[i]})
    }
    cardLocations.EntityData.Leafs = types.NewOrderedMap()
    cardLocations.EntityData.Leafs.Append("card_location", types.YLeaf{"CardLocation", cardLocations.CardLocation})

    cardLocations.EntityData.YListKeys = []string {"CardLocation"}

    return &(cardLocations.EntityData)
}

// CardInventory_CardLocations_Cardi
type CardInventory_CardLocations_Cardi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Serial Number of the Card. The type is string.
    CardSerial interface{}

    // Node name. The type is string.
    NodeId interface{}

    // Card Type. The type is string.
    CardType interface{}

    // Card State. The type is string.
    HwState interface{}

    // Card Software State. The type is string.
    SwState interface{}

    // Card Slot Number. The type is interface{} with range: 0..4294967295.
    Slot interface{}

    // Card CTI. The type is interface{} with range: 0..4294967295.
    Cti interface{}
}

func (cardi *CardInventory_CardLocations_Cardi) GetEntityData() *types.CommonEntityData {
    cardi.EntityData.YFilter = cardi.YFilter
    cardi.EntityData.YangName = "cardi"
    cardi.EntityData.BundleName = "cisco_ios_xr"
    cardi.EntityData.ParentYangName = "card_locations"
    cardi.EntityData.SegmentPath = "cardi" + types.AddKeyToken(cardi.CardSerial, "card_serial")
    cardi.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-cm:card-inventory/card_locations/" + cardi.EntityData.SegmentPath
    cardi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cardi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cardi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cardi.EntityData.Children = types.NewOrderedMap()
    cardi.EntityData.Leafs = types.NewOrderedMap()
    cardi.EntityData.Leafs.Append("card_serial", types.YLeaf{"CardSerial", cardi.CardSerial})
    cardi.EntityData.Leafs.Append("node_id", types.YLeaf{"NodeId", cardi.NodeId})
    cardi.EntityData.Leafs.Append("card_type", types.YLeaf{"CardType", cardi.CardType})
    cardi.EntityData.Leafs.Append("hw_state", types.YLeaf{"HwState", cardi.HwState})
    cardi.EntityData.Leafs.Append("sw_state", types.YLeaf{"SwState", cardi.SwState})
    cardi.EntityData.Leafs.Append("slot", types.YLeaf{"Slot", cardi.Slot})
    cardi.EntityData.Leafs.Append("cti", types.YLeaf{"Cti", cardi.Cti})

    cardi.EntityData.YListKeys = []string {"CardSerial"}

    return &(cardi.EntityData)
}

// RackInventory
// System Admin Manager Rack Inventory
type RackInventory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of RackInventory_RackLocations.
    RackLocations []*RackInventory_RackLocations
}

func (rackInventory *RackInventory) GetEntityData() *types.CommonEntityData {
    rackInventory.EntityData.YFilter = rackInventory.YFilter
    rackInventory.EntityData.YangName = "rack-inventory"
    rackInventory.EntityData.BundleName = "cisco_ios_xr"
    rackInventory.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-cm"
    rackInventory.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-cm:rack-inventory"
    rackInventory.EntityData.AbsolutePath = rackInventory.EntityData.SegmentPath
    rackInventory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rackInventory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rackInventory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rackInventory.EntityData.Children = types.NewOrderedMap()
    rackInventory.EntityData.Children.Append("rack_locations", types.YChild{"RackLocations", nil})
    for i := range rackInventory.RackLocations {
        rackInventory.EntityData.Children.Append(types.GetSegmentPath(rackInventory.RackLocations[i]), types.YChild{"RackLocations", rackInventory.RackLocations[i]})
    }
    rackInventory.EntityData.Leafs = types.NewOrderedMap()

    rackInventory.EntityData.YListKeys = []string {}

    return &(rackInventory.EntityData)
}

// RackInventory_RackLocations
type RackInventory_RackLocations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    RackLocation interface{}

    // The type is slice of RackInventory_RackLocations_Racki.
    Racki []*RackInventory_RackLocations_Racki
}

func (rackLocations *RackInventory_RackLocations) GetEntityData() *types.CommonEntityData {
    rackLocations.EntityData.YFilter = rackLocations.YFilter
    rackLocations.EntityData.YangName = "rack_locations"
    rackLocations.EntityData.BundleName = "cisco_ios_xr"
    rackLocations.EntityData.ParentYangName = "rack-inventory"
    rackLocations.EntityData.SegmentPath = "rack_locations" + types.AddKeyToken(rackLocations.RackLocation, "rack_location")
    rackLocations.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-cm:rack-inventory/" + rackLocations.EntityData.SegmentPath
    rackLocations.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rackLocations.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rackLocations.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rackLocations.EntityData.Children = types.NewOrderedMap()
    rackLocations.EntityData.Children.Append("racki", types.YChild{"Racki", nil})
    for i := range rackLocations.Racki {
        rackLocations.EntityData.Children.Append(types.GetSegmentPath(rackLocations.Racki[i]), types.YChild{"Racki", rackLocations.Racki[i]})
    }
    rackLocations.EntityData.Leafs = types.NewOrderedMap()
    rackLocations.EntityData.Leafs.Append("rack_location", types.YLeaf{"RackLocation", rackLocations.RackLocation})

    rackLocations.EntityData.YListKeys = []string {"RackLocation"}

    return &(rackLocations.EntityData)
}

// RackInventory_RackLocations_Racki
type RackInventory_RackLocations_Racki struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Serial Number of the Rack. The type is string.
    RackSerial interface{}

    // Rack Number. The type is interface{} with range: -2147483648..2147483647.
    RackNumber interface{}

    // Rack State. The type is interface{} with range: -2147483648..2147483647.
    RackState interface{}
}

func (racki *RackInventory_RackLocations_Racki) GetEntityData() *types.CommonEntityData {
    racki.EntityData.YFilter = racki.YFilter
    racki.EntityData.YangName = "racki"
    racki.EntityData.BundleName = "cisco_ios_xr"
    racki.EntityData.ParentYangName = "rack_locations"
    racki.EntityData.SegmentPath = "racki" + types.AddKeyToken(racki.RackSerial, "rack_serial")
    racki.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-cm:rack-inventory/rack_locations/" + racki.EntityData.SegmentPath
    racki.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    racki.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    racki.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    racki.EntityData.Children = types.NewOrderedMap()
    racki.EntityData.Leafs = types.NewOrderedMap()
    racki.EntityData.Leafs.Append("rack_serial", types.YLeaf{"RackSerial", racki.RackSerial})
    racki.EntityData.Leafs.Append("rack_number", types.YLeaf{"RackNumber", racki.RackNumber})
    racki.EntityData.Leafs.Append("rack_state", types.YLeaf{"RackState", racki.RackState})

    racki.EntityData.YListKeys = []string {"RackSerial"}

    return &(racki.EntityData)
}

// SystemServiceInventory
// System Admin Manager System Services Inventory
type SystemServiceInventory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of SystemServiceInventory_SsvcLocations.
    SsvcLocations []*SystemServiceInventory_SsvcLocations
}

func (systemServiceInventory *SystemServiceInventory) GetEntityData() *types.CommonEntityData {
    systemServiceInventory.EntityData.YFilter = systemServiceInventory.YFilter
    systemServiceInventory.EntityData.YangName = "system-service-inventory"
    systemServiceInventory.EntityData.BundleName = "cisco_ios_xr"
    systemServiceInventory.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-cm"
    systemServiceInventory.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-cm:system-service-inventory"
    systemServiceInventory.EntityData.AbsolutePath = systemServiceInventory.EntityData.SegmentPath
    systemServiceInventory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    systemServiceInventory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    systemServiceInventory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    systemServiceInventory.EntityData.Children = types.NewOrderedMap()
    systemServiceInventory.EntityData.Children.Append("ssvc_locations", types.YChild{"SsvcLocations", nil})
    for i := range systemServiceInventory.SsvcLocations {
        systemServiceInventory.EntityData.Children.Append(types.GetSegmentPath(systemServiceInventory.SsvcLocations[i]), types.YChild{"SsvcLocations", systemServiceInventory.SsvcLocations[i]})
    }
    systemServiceInventory.EntityData.Leafs = types.NewOrderedMap()

    systemServiceInventory.EntityData.YListKeys = []string {}

    return &(systemServiceInventory.EntityData)
}

// SystemServiceInventory_SsvcLocations
type SystemServiceInventory_SsvcLocations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    SsvcLocation interface{}

    // The type is slice of SystemServiceInventory_SsvcLocations_Ssvci.
    Ssvci []*SystemServiceInventory_SsvcLocations_Ssvci
}

func (ssvcLocations *SystemServiceInventory_SsvcLocations) GetEntityData() *types.CommonEntityData {
    ssvcLocations.EntityData.YFilter = ssvcLocations.YFilter
    ssvcLocations.EntityData.YangName = "ssvc_locations"
    ssvcLocations.EntityData.BundleName = "cisco_ios_xr"
    ssvcLocations.EntityData.ParentYangName = "system-service-inventory"
    ssvcLocations.EntityData.SegmentPath = "ssvc_locations" + types.AddKeyToken(ssvcLocations.SsvcLocation, "ssvc_location")
    ssvcLocations.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-cm:system-service-inventory/" + ssvcLocations.EntityData.SegmentPath
    ssvcLocations.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ssvcLocations.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ssvcLocations.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ssvcLocations.EntityData.Children = types.NewOrderedMap()
    ssvcLocations.EntityData.Children.Append("ssvci", types.YChild{"Ssvci", nil})
    for i := range ssvcLocations.Ssvci {
        ssvcLocations.EntityData.Children.Append(types.GetSegmentPath(ssvcLocations.Ssvci[i]), types.YChild{"Ssvci", ssvcLocations.Ssvci[i]})
    }
    ssvcLocations.EntityData.Leafs = types.NewOrderedMap()
    ssvcLocations.EntityData.Leafs.Append("ssvc_location", types.YLeaf{"SsvcLocation", ssvcLocations.SsvcLocation})

    ssvcLocations.EntityData.YListKeys = []string {"SsvcLocation"}

    return &(ssvcLocations.EntityData)
}

// SystemServiceInventory_SsvcLocations_Ssvci
type SystemServiceInventory_SsvcLocations_Ssvci struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Service Name. The type is string.
    SvcName interface{}

    // Serial Number of the first card selected. The type is string.
    PlacementFirst interface{}

    // Node id of the first card selected. The type is string.
    NodeidFirst interface{}

    // Serial Number of the second card selected. The type is string.
    PlacementSecond interface{}

    // Node id of the second card selected. The type is string.
    NodeidSecond interface{}

    // Service Load. The type is interface{} with range: 0..255.
    SvcLoad interface{}
}

func (ssvci *SystemServiceInventory_SsvcLocations_Ssvci) GetEntityData() *types.CommonEntityData {
    ssvci.EntityData.YFilter = ssvci.YFilter
    ssvci.EntityData.YangName = "ssvci"
    ssvci.EntityData.BundleName = "cisco_ios_xr"
    ssvci.EntityData.ParentYangName = "ssvc_locations"
    ssvci.EntityData.SegmentPath = "ssvci" + types.AddKeyToken(ssvci.SvcName, "svc_name")
    ssvci.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-cm:system-service-inventory/ssvc_locations/" + ssvci.EntityData.SegmentPath
    ssvci.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ssvci.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ssvci.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ssvci.EntityData.Children = types.NewOrderedMap()
    ssvci.EntityData.Leafs = types.NewOrderedMap()
    ssvci.EntityData.Leafs.Append("svc_name", types.YLeaf{"SvcName", ssvci.SvcName})
    ssvci.EntityData.Leafs.Append("placement_first", types.YLeaf{"PlacementFirst", ssvci.PlacementFirst})
    ssvci.EntityData.Leafs.Append("nodeid_first", types.YLeaf{"NodeidFirst", ssvci.NodeidFirst})
    ssvci.EntityData.Leafs.Append("placement_second", types.YLeaf{"PlacementSecond", ssvci.PlacementSecond})
    ssvci.EntityData.Leafs.Append("nodeid_second", types.YLeaf{"NodeidSecond", ssvci.NodeidSecond})
    ssvci.EntityData.Leafs.Append("svc_load", types.YLeaf{"SvcLoad", ssvci.SvcLoad})

    ssvci.EntityData.YListKeys = []string {"SvcName"}

    return &(ssvci.EntityData)
}

// RackServiceInventory
// System Admin Manager Rack Services Inventory
type RackServiceInventory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of RackServiceInventory_RsvcLocations.
    RsvcLocations []*RackServiceInventory_RsvcLocations
}

func (rackServiceInventory *RackServiceInventory) GetEntityData() *types.CommonEntityData {
    rackServiceInventory.EntityData.YFilter = rackServiceInventory.YFilter
    rackServiceInventory.EntityData.YangName = "rack-service-inventory"
    rackServiceInventory.EntityData.BundleName = "cisco_ios_xr"
    rackServiceInventory.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-cm"
    rackServiceInventory.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-cm:rack-service-inventory"
    rackServiceInventory.EntityData.AbsolutePath = rackServiceInventory.EntityData.SegmentPath
    rackServiceInventory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rackServiceInventory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rackServiceInventory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rackServiceInventory.EntityData.Children = types.NewOrderedMap()
    rackServiceInventory.EntityData.Children.Append("rsvc_locations", types.YChild{"RsvcLocations", nil})
    for i := range rackServiceInventory.RsvcLocations {
        rackServiceInventory.EntityData.Children.Append(types.GetSegmentPath(rackServiceInventory.RsvcLocations[i]), types.YChild{"RsvcLocations", rackServiceInventory.RsvcLocations[i]})
    }
    rackServiceInventory.EntityData.Leafs = types.NewOrderedMap()

    rackServiceInventory.EntityData.YListKeys = []string {}

    return &(rackServiceInventory.EntityData)
}

// RackServiceInventory_RsvcLocations
type RackServiceInventory_RsvcLocations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    RsvcLocation interface{}

    // The type is slice of RackServiceInventory_RsvcLocations_Rsvci.
    Rsvci []*RackServiceInventory_RsvcLocations_Rsvci
}

func (rsvcLocations *RackServiceInventory_RsvcLocations) GetEntityData() *types.CommonEntityData {
    rsvcLocations.EntityData.YFilter = rsvcLocations.YFilter
    rsvcLocations.EntityData.YangName = "rsvc_locations"
    rsvcLocations.EntityData.BundleName = "cisco_ios_xr"
    rsvcLocations.EntityData.ParentYangName = "rack-service-inventory"
    rsvcLocations.EntityData.SegmentPath = "rsvc_locations" + types.AddKeyToken(rsvcLocations.RsvcLocation, "rsvc_location")
    rsvcLocations.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-cm:rack-service-inventory/" + rsvcLocations.EntityData.SegmentPath
    rsvcLocations.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rsvcLocations.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rsvcLocations.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rsvcLocations.EntityData.Children = types.NewOrderedMap()
    rsvcLocations.EntityData.Children.Append("rsvci", types.YChild{"Rsvci", nil})
    for i := range rsvcLocations.Rsvci {
        rsvcLocations.EntityData.Children.Append(types.GetSegmentPath(rsvcLocations.Rsvci[i]), types.YChild{"Rsvci", rsvcLocations.Rsvci[i]})
    }
    rsvcLocations.EntityData.Leafs = types.NewOrderedMap()
    rsvcLocations.EntityData.Leafs.Append("rsvc_location", types.YLeaf{"RsvcLocation", rsvcLocations.RsvcLocation})

    rsvcLocations.EntityData.YListKeys = []string {"RsvcLocation"}

    return &(rsvcLocations.EntityData)
}

// RackServiceInventory_RsvcLocations_Rsvci
type RackServiceInventory_RsvcLocations_Rsvci struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Service Name. The type is string.
    SvcName interface{}

    // Serial Number of the first card selected. The type is string.
    PlacementFirst interface{}

    // Node id of the first card selected. The type is string.
    NodeidFirst interface{}

    // Serial Number of the second card selected. The type is string.
    PlacementSecond interface{}

    // Node id of the second card selected. The type is string.
    NodeidSecond interface{}

    // Service Load. The type is interface{} with range: 0..255.
    SvcLoad interface{}
}

func (rsvci *RackServiceInventory_RsvcLocations_Rsvci) GetEntityData() *types.CommonEntityData {
    rsvci.EntityData.YFilter = rsvci.YFilter
    rsvci.EntityData.YangName = "rsvci"
    rsvci.EntityData.BundleName = "cisco_ios_xr"
    rsvci.EntityData.ParentYangName = "rsvc_locations"
    rsvci.EntityData.SegmentPath = "rsvci" + types.AddKeyToken(rsvci.SvcName, "svc_name")
    rsvci.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-cm:rack-service-inventory/rsvc_locations/" + rsvci.EntityData.SegmentPath
    rsvci.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rsvci.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rsvci.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rsvci.EntityData.Children = types.NewOrderedMap()
    rsvci.EntityData.Leafs = types.NewOrderedMap()
    rsvci.EntityData.Leafs.Append("svc_name", types.YLeaf{"SvcName", rsvci.SvcName})
    rsvci.EntityData.Leafs.Append("placement_first", types.YLeaf{"PlacementFirst", rsvci.PlacementFirst})
    rsvci.EntityData.Leafs.Append("nodeid_first", types.YLeaf{"NodeidFirst", rsvci.NodeidFirst})
    rsvci.EntityData.Leafs.Append("placement_second", types.YLeaf{"PlacementSecond", rsvci.PlacementSecond})
    rsvci.EntityData.Leafs.Append("nodeid_second", types.YLeaf{"NodeidSecond", rsvci.NodeidSecond})
    rsvci.EntityData.Leafs.Append("svc_load", types.YLeaf{"SvcLoad", rsvci.SvcLoad})

    rsvci.EntityData.YListKeys = []string {"SvcName"}

    return &(rsvci.EntityData)
}

// SdrInventory
// System Admin Manager SDR Inventory
type SdrInventory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of SdrInventory_SdrLocations.
    SdrLocations []*SdrInventory_SdrLocations
}

func (sdrInventory *SdrInventory) GetEntityData() *types.CommonEntityData {
    sdrInventory.EntityData.YFilter = sdrInventory.YFilter
    sdrInventory.EntityData.YangName = "sdr-inventory"
    sdrInventory.EntityData.BundleName = "cisco_ios_xr"
    sdrInventory.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-cm"
    sdrInventory.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-cm:sdr-inventory"
    sdrInventory.EntityData.AbsolutePath = sdrInventory.EntityData.SegmentPath
    sdrInventory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sdrInventory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sdrInventory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sdrInventory.EntityData.Children = types.NewOrderedMap()
    sdrInventory.EntityData.Children.Append("sdr_locations", types.YChild{"SdrLocations", nil})
    for i := range sdrInventory.SdrLocations {
        sdrInventory.EntityData.Children.Append(types.GetSegmentPath(sdrInventory.SdrLocations[i]), types.YChild{"SdrLocations", sdrInventory.SdrLocations[i]})
    }
    sdrInventory.EntityData.Leafs = types.NewOrderedMap()

    sdrInventory.EntityData.YListKeys = []string {}

    return &(sdrInventory.EntityData)
}

// SdrInventory_SdrLocations
type SdrInventory_SdrLocations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    SdrLocation interface{}

    // The type is slice of SdrInventory_SdrLocations_Sdri.
    Sdri []*SdrInventory_SdrLocations_Sdri
}

func (sdrLocations *SdrInventory_SdrLocations) GetEntityData() *types.CommonEntityData {
    sdrLocations.EntityData.YFilter = sdrLocations.YFilter
    sdrLocations.EntityData.YangName = "sdr_locations"
    sdrLocations.EntityData.BundleName = "cisco_ios_xr"
    sdrLocations.EntityData.ParentYangName = "sdr-inventory"
    sdrLocations.EntityData.SegmentPath = "sdr_locations" + types.AddKeyToken(sdrLocations.SdrLocation, "sdr_location")
    sdrLocations.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-cm:sdr-inventory/" + sdrLocations.EntityData.SegmentPath
    sdrLocations.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sdrLocations.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sdrLocations.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sdrLocations.EntityData.Children = types.NewOrderedMap()
    sdrLocations.EntityData.Children.Append("sdri", types.YChild{"Sdri", nil})
    for i := range sdrLocations.Sdri {
        sdrLocations.EntityData.Children.Append(types.GetSegmentPath(sdrLocations.Sdri[i]), types.YChild{"Sdri", sdrLocations.Sdri[i]})
    }
    sdrLocations.EntityData.Leafs = types.NewOrderedMap()
    sdrLocations.EntityData.Leafs.Append("sdr_location", types.YLeaf{"SdrLocation", sdrLocations.SdrLocation})

    sdrLocations.EntityData.YListKeys = []string {"SdrLocation"}

    return &(sdrLocations.EntityData)
}

// SdrInventory_SdrLocations_Sdri
type SdrInventory_SdrLocations_Sdri struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. SDR NAME. The type is string.
    SdrName interface{}

    // SDR ID. The type is interface{} with range: 0..4294967295.
    SdrId interface{}

    // SDR VLAN ID. The type is interface{} with range: 0..255.
    SdrVlanBaseid interface{}

    // SDR Image Version. The type is interface{} with range:
    // 0..18446744073709551615.
    SdrVersion interface{}
}

func (sdri *SdrInventory_SdrLocations_Sdri) GetEntityData() *types.CommonEntityData {
    sdri.EntityData.YFilter = sdri.YFilter
    sdri.EntityData.YangName = "sdri"
    sdri.EntityData.BundleName = "cisco_ios_xr"
    sdri.EntityData.ParentYangName = "sdr_locations"
    sdri.EntityData.SegmentPath = "sdri" + types.AddKeyToken(sdri.SdrName, "sdr_name")
    sdri.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-cm:sdr-inventory/sdr_locations/" + sdri.EntityData.SegmentPath
    sdri.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sdri.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sdri.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sdri.EntityData.Children = types.NewOrderedMap()
    sdri.EntityData.Leafs = types.NewOrderedMap()
    sdri.EntityData.Leafs.Append("sdr_name", types.YLeaf{"SdrName", sdri.SdrName})
    sdri.EntityData.Leafs.Append("sdr_id", types.YLeaf{"SdrId", sdri.SdrId})
    sdri.EntityData.Leafs.Append("sdr_vlan_baseid", types.YLeaf{"SdrVlanBaseid", sdri.SdrVlanBaseid})
    sdri.EntityData.Leafs.Append("sdr_version", types.YLeaf{"SdrVersion", sdri.SdrVersion})

    sdri.EntityData.YListKeys = []string {"SdrName"}

    return &(sdri.EntityData)
}

// LeaderStatistics
// System Admin Manager Leader Statistics
type LeaderStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of LeaderStatistics_LdrLocations.
    LdrLocations []*LeaderStatistics_LdrLocations
}

func (leaderStatistics *LeaderStatistics) GetEntityData() *types.CommonEntityData {
    leaderStatistics.EntityData.YFilter = leaderStatistics.YFilter
    leaderStatistics.EntityData.YangName = "leader-statistics"
    leaderStatistics.EntityData.BundleName = "cisco_ios_xr"
    leaderStatistics.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-cm"
    leaderStatistics.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-cm:leader-statistics"
    leaderStatistics.EntityData.AbsolutePath = leaderStatistics.EntityData.SegmentPath
    leaderStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    leaderStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    leaderStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    leaderStatistics.EntityData.Children = types.NewOrderedMap()
    leaderStatistics.EntityData.Children.Append("ldr_locations", types.YChild{"LdrLocations", nil})
    for i := range leaderStatistics.LdrLocations {
        leaderStatistics.EntityData.Children.Append(types.GetSegmentPath(leaderStatistics.LdrLocations[i]), types.YChild{"LdrLocations", leaderStatistics.LdrLocations[i]})
    }
    leaderStatistics.EntityData.Leafs = types.NewOrderedMap()

    leaderStatistics.EntityData.YListKeys = []string {}

    return &(leaderStatistics.EntityData)
}

// LeaderStatistics_LdrLocations
type LeaderStatistics_LdrLocations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LdrLocation interface{}

    // Primary System Leader. The type is string.
    Syslead interface{}

    // Backup System Leader. The type is string.
    BkupSyslead interface{}

    // Primary Rack Leader. The type is string.
    Racklead interface{}

    // Backup Rack Leader. The type is string.
    BkupRacklead interface{}

    // L1 DIS. The type is string.
    L1Dis interface{}

    // L2 DIS. The type is string.
    L2Dis interface{}
}

func (ldrLocations *LeaderStatistics_LdrLocations) GetEntityData() *types.CommonEntityData {
    ldrLocations.EntityData.YFilter = ldrLocations.YFilter
    ldrLocations.EntityData.YangName = "ldr_locations"
    ldrLocations.EntityData.BundleName = "cisco_ios_xr"
    ldrLocations.EntityData.ParentYangName = "leader-statistics"
    ldrLocations.EntityData.SegmentPath = "ldr_locations" + types.AddKeyToken(ldrLocations.LdrLocation, "ldr_location")
    ldrLocations.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-cm:leader-statistics/" + ldrLocations.EntityData.SegmentPath
    ldrLocations.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ldrLocations.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ldrLocations.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ldrLocations.EntityData.Children = types.NewOrderedMap()
    ldrLocations.EntityData.Leafs = types.NewOrderedMap()
    ldrLocations.EntityData.Leafs.Append("ldr_location", types.YLeaf{"LdrLocation", ldrLocations.LdrLocation})
    ldrLocations.EntityData.Leafs.Append("syslead", types.YLeaf{"Syslead", ldrLocations.Syslead})
    ldrLocations.EntityData.Leafs.Append("bkup_syslead", types.YLeaf{"BkupSyslead", ldrLocations.BkupSyslead})
    ldrLocations.EntityData.Leafs.Append("racklead", types.YLeaf{"Racklead", ldrLocations.Racklead})
    ldrLocations.EntityData.Leafs.Append("bkup_racklead", types.YLeaf{"BkupRacklead", ldrLocations.BkupRacklead})
    ldrLocations.EntityData.Leafs.Append("l1_dis", types.YLeaf{"L1Dis", ldrLocations.L1Dis})
    ldrLocations.EntityData.Leafs.Append("l2_dis", types.YLeaf{"L2Dis", ldrLocations.L2Dis})

    ldrLocations.EntityData.YListKeys = []string {"LdrLocation"}

    return &(ldrLocations.EntityData)
}

// TopologyNeighbors
// System Admin Manager Neighbors of a location
type TopologyNeighbors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of TopologyNeighbors_NbrLocations.
    NbrLocations []*TopologyNeighbors_NbrLocations
}

func (topologyNeighbors *TopologyNeighbors) GetEntityData() *types.CommonEntityData {
    topologyNeighbors.EntityData.YFilter = topologyNeighbors.YFilter
    topologyNeighbors.EntityData.YangName = "topology-neighbors"
    topologyNeighbors.EntityData.BundleName = "cisco_ios_xr"
    topologyNeighbors.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-cm"
    topologyNeighbors.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-cm:topology-neighbors"
    topologyNeighbors.EntityData.AbsolutePath = topologyNeighbors.EntityData.SegmentPath
    topologyNeighbors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    topologyNeighbors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    topologyNeighbors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    topologyNeighbors.EntityData.Children = types.NewOrderedMap()
    topologyNeighbors.EntityData.Children.Append("nbr_locations", types.YChild{"NbrLocations", nil})
    for i := range topologyNeighbors.NbrLocations {
        topologyNeighbors.EntityData.Children.Append(types.GetSegmentPath(topologyNeighbors.NbrLocations[i]), types.YChild{"NbrLocations", topologyNeighbors.NbrLocations[i]})
    }
    topologyNeighbors.EntityData.Leafs = types.NewOrderedMap()

    topologyNeighbors.EntityData.YListKeys = []string {}

    return &(topologyNeighbors.EntityData)
}

// TopologyNeighbors_NbrLocations
type TopologyNeighbors_NbrLocations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    NbrLocation interface{}

    // The type is slice of TopologyNeighbors_NbrLocations_Nbri.
    Nbri []*TopologyNeighbors_NbrLocations_Nbri
}

func (nbrLocations *TopologyNeighbors_NbrLocations) GetEntityData() *types.CommonEntityData {
    nbrLocations.EntityData.YFilter = nbrLocations.YFilter
    nbrLocations.EntityData.YangName = "nbr_locations"
    nbrLocations.EntityData.BundleName = "cisco_ios_xr"
    nbrLocations.EntityData.ParentYangName = "topology-neighbors"
    nbrLocations.EntityData.SegmentPath = "nbr_locations" + types.AddKeyToken(nbrLocations.NbrLocation, "nbr_location")
    nbrLocations.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-cm:topology-neighbors/" + nbrLocations.EntityData.SegmentPath
    nbrLocations.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nbrLocations.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nbrLocations.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nbrLocations.EntityData.Children = types.NewOrderedMap()
    nbrLocations.EntityData.Children.Append("nbri", types.YChild{"Nbri", nil})
    for i := range nbrLocations.Nbri {
        nbrLocations.EntityData.Children.Append(types.GetSegmentPath(nbrLocations.Nbri[i]), types.YChild{"Nbri", nbrLocations.Nbri[i]})
    }
    nbrLocations.EntityData.Leafs = types.NewOrderedMap()
    nbrLocations.EntityData.Leafs.Append("nbr_location", types.YLeaf{"NbrLocation", nbrLocations.NbrLocation})

    nbrLocations.EntityData.YListKeys = []string {"NbrLocation"}

    return &(nbrLocations.EntityData)
}

// TopologyNeighbors_NbrLocations_Nbri
type TopologyNeighbors_NbrLocations_Nbri struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Neighbor System ID. The type is string.
    NbrSystemId interface{}

    // This attribute is a key. Neighbor Area Type. The type is string.
    NbrAreaType interface{}

    // Adjacency Interface. The type is string.
    NbrInterface interface{}

    // Neighbor State. The type is string.
    NbrState interface{}

    // Neighbor Hold Time. The type is interface{} with range:
    // 0..18446744073709551615.
    NbrHoldtime interface{}

    // Neighbor Up Time. The type is string.
    NbrUptime interface{}
}

func (nbri *TopologyNeighbors_NbrLocations_Nbri) GetEntityData() *types.CommonEntityData {
    nbri.EntityData.YFilter = nbri.YFilter
    nbri.EntityData.YangName = "nbri"
    nbri.EntityData.BundleName = "cisco_ios_xr"
    nbri.EntityData.ParentYangName = "nbr_locations"
    nbri.EntityData.SegmentPath = "nbri" + types.AddKeyToken(nbri.NbrSystemId, "nbr_system_id") + types.AddKeyToken(nbri.NbrAreaType, "nbr_area_type")
    nbri.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-cm:topology-neighbors/nbr_locations/" + nbri.EntityData.SegmentPath
    nbri.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nbri.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nbri.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nbri.EntityData.Children = types.NewOrderedMap()
    nbri.EntityData.Leafs = types.NewOrderedMap()
    nbri.EntityData.Leafs.Append("nbr_system_id", types.YLeaf{"NbrSystemId", nbri.NbrSystemId})
    nbri.EntityData.Leafs.Append("nbr_area_type", types.YLeaf{"NbrAreaType", nbri.NbrAreaType})
    nbri.EntityData.Leafs.Append("nbr_interface", types.YLeaf{"NbrInterface", nbri.NbrInterface})
    nbri.EntityData.Leafs.Append("nbr_state", types.YLeaf{"NbrState", nbri.NbrState})
    nbri.EntityData.Leafs.Append("nbr_holdtime", types.YLeaf{"NbrHoldtime", nbri.NbrHoldtime})
    nbri.EntityData.Leafs.Append("nbr_uptime", types.YLeaf{"NbrUptime", nbri.NbrUptime})

    nbri.EntityData.YListKeys = []string {"NbrSystemId", "NbrAreaType"}

    return &(nbri.EntityData)
}

// Placement
type Placement struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (placement *Placement) GetEntityData() *types.CommonEntityData {
    placement.EntityData.YFilter = placement.YFilter
    placement.EntityData.YangName = "placement"
    placement.EntityData.BundleName = "cisco_ios_xr"
    placement.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-cm"
    placement.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-cm:placement"
    placement.EntityData.AbsolutePath = placement.EntityData.SegmentPath
    placement.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    placement.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    placement.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    placement.EntityData.Children = types.NewOrderedMap()
    placement.EntityData.Leafs = types.NewOrderedMap()

    placement.EntityData.YListKeys = []string {}

    return &(placement.EntityData)
}

