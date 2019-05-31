// This module contains a collection of YANG definitions
// for Cisco IOS-XR sdr-invmgr package operational data.
// 
// This module contains definitions
// for the following management objects:
//   sdr-inventory: SDR information
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package sdr_invmgr_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package sdr_invmgr_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-sdr-invmgr-oper sdr-inventory}", reflect.TypeOf(SdrInventory{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sdr-invmgr-oper:sdr-inventory", reflect.TypeOf(SdrInventory{}))
}

// SdrInventory
// SDR information
type SdrInventory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RackTable.
    Racks SdrInventory_Racks

    // Memory.
    Memory SdrInventory_Memory
}

func (sdrInventory *SdrInventory) GetEntityData() *types.CommonEntityData {
    sdrInventory.EntityData.YFilter = sdrInventory.YFilter
    sdrInventory.EntityData.YangName = "sdr-inventory"
    sdrInventory.EntityData.BundleName = "cisco_ios_xr"
    sdrInventory.EntityData.ParentYangName = "Cisco-IOS-XR-sdr-invmgr-oper"
    sdrInventory.EntityData.SegmentPath = "Cisco-IOS-XR-sdr-invmgr-oper:sdr-inventory"
    sdrInventory.EntityData.AbsolutePath = sdrInventory.EntityData.SegmentPath
    sdrInventory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sdrInventory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sdrInventory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sdrInventory.EntityData.Children = types.NewOrderedMap()
    sdrInventory.EntityData.Children.Append("racks", types.YChild{"Racks", &sdrInventory.Racks})
    sdrInventory.EntityData.Children.Append("memory", types.YChild{"Memory", &sdrInventory.Memory})
    sdrInventory.EntityData.Leafs = types.NewOrderedMap()

    sdrInventory.EntityData.YListKeys = []string {}

    return &(sdrInventory.EntityData)
}

// SdrInventory_Racks
// RackTable
type SdrInventory_Racks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Rack name. The type is slice of SdrInventory_Racks_Rack.
    Rack []*SdrInventory_Racks_Rack
}

func (racks *SdrInventory_Racks) GetEntityData() *types.CommonEntityData {
    racks.EntityData.YFilter = racks.YFilter
    racks.EntityData.YangName = "racks"
    racks.EntityData.BundleName = "cisco_ios_xr"
    racks.EntityData.ParentYangName = "sdr-inventory"
    racks.EntityData.SegmentPath = "racks"
    racks.EntityData.AbsolutePath = "Cisco-IOS-XR-sdr-invmgr-oper:sdr-inventory/" + racks.EntityData.SegmentPath
    racks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    racks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    racks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    racks.EntityData.Children = types.NewOrderedMap()
    racks.EntityData.Children.Append("rack", types.YChild{"Rack", nil})
    for i := range racks.Rack {
        racks.EntityData.Children.Append(types.GetSegmentPath(racks.Rack[i]), types.YChild{"Rack", racks.Rack[i]})
    }
    racks.EntityData.Leafs = types.NewOrderedMap()

    racks.EntityData.YListKeys = []string {}

    return &(racks.EntityData)
}

// SdrInventory_Racks_Rack
// Rack name
type SdrInventory_Racks_Rack struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Rack name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Name interface{}

    // Slot name. The type is slice of SdrInventory_Racks_Rack_Slot.
    Slot []*SdrInventory_Racks_Rack_Slot
}

func (rack *SdrInventory_Racks_Rack) GetEntityData() *types.CommonEntityData {
    rack.EntityData.YFilter = rack.YFilter
    rack.EntityData.YangName = "rack"
    rack.EntityData.BundleName = "cisco_ios_xr"
    rack.EntityData.ParentYangName = "racks"
    rack.EntityData.SegmentPath = "rack" + types.AddKeyToken(rack.Name, "name")
    rack.EntityData.AbsolutePath = "Cisco-IOS-XR-sdr-invmgr-oper:sdr-inventory/racks/" + rack.EntityData.SegmentPath
    rack.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rack.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rack.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rack.EntityData.Children = types.NewOrderedMap()
    rack.EntityData.Children.Append("slot", types.YChild{"Slot", nil})
    for i := range rack.Slot {
        rack.EntityData.Children.Append(types.GetSegmentPath(rack.Slot[i]), types.YChild{"Slot", rack.Slot[i]})
    }
    rack.EntityData.Leafs = types.NewOrderedMap()
    rack.EntityData.Leafs.Append("name", types.YLeaf{"Name", rack.Name})

    rack.EntityData.YListKeys = []string {"Name"}

    return &(rack.EntityData)
}

// SdrInventory_Racks_Rack_Slot
// Slot name
type SdrInventory_Racks_Rack_Slot struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Slot name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Name interface{}

    // Card. The type is slice of SdrInventory_Racks_Rack_Slot_Card.
    Card []*SdrInventory_Racks_Rack_Slot_Card
}

func (slot *SdrInventory_Racks_Rack_Slot) GetEntityData() *types.CommonEntityData {
    slot.EntityData.YFilter = slot.YFilter
    slot.EntityData.YangName = "slot"
    slot.EntityData.BundleName = "cisco_ios_xr"
    slot.EntityData.ParentYangName = "rack"
    slot.EntityData.SegmentPath = "slot" + types.AddKeyToken(slot.Name, "name")
    slot.EntityData.AbsolutePath = "Cisco-IOS-XR-sdr-invmgr-oper:sdr-inventory/racks/rack/" + slot.EntityData.SegmentPath
    slot.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    slot.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    slot.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    slot.EntityData.Children = types.NewOrderedMap()
    slot.EntityData.Children.Append("card", types.YChild{"Card", nil})
    for i := range slot.Card {
        slot.EntityData.Children.Append(types.GetSegmentPath(slot.Card[i]), types.YChild{"Card", slot.Card[i]})
    }
    slot.EntityData.Leafs = types.NewOrderedMap()
    slot.EntityData.Leafs.Append("name", types.YLeaf{"Name", slot.Name})

    slot.EntityData.YListKeys = []string {"Name"}

    return &(slot.EntityData)
}

// SdrInventory_Racks_Rack_Slot_Card
// Card
type SdrInventory_Racks_Rack_Slot_Card struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Card. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Name interface{}

    // Attributes.
    Attributes SdrInventory_Racks_Rack_Slot_Card_Attributes
}

func (card *SdrInventory_Racks_Rack_Slot_Card) GetEntityData() *types.CommonEntityData {
    card.EntityData.YFilter = card.YFilter
    card.EntityData.YangName = "card"
    card.EntityData.BundleName = "cisco_ios_xr"
    card.EntityData.ParentYangName = "slot"
    card.EntityData.SegmentPath = "card" + types.AddKeyToken(card.Name, "name")
    card.EntityData.AbsolutePath = "Cisco-IOS-XR-sdr-invmgr-oper:sdr-inventory/racks/rack/slot/" + card.EntityData.SegmentPath
    card.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    card.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    card.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    card.EntityData.Children = types.NewOrderedMap()
    card.EntityData.Children.Append("attributes", types.YChild{"Attributes", &card.Attributes})
    card.EntityData.Leafs = types.NewOrderedMap()
    card.EntityData.Leafs.Append("name", types.YLeaf{"Name", card.Name})

    card.EntityData.YListKeys = []string {"Name"}

    return &(card.EntityData)
}

// SdrInventory_Racks_Rack_Slot_Card_Attributes
// Attributes
type SdrInventory_Racks_Rack_Slot_Card_Attributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config State String. The type is string.
    ConfigStateString interface{}

    // Power. The type is interface{} with range: 0..4294967295. The default value
    // is 0.
    Power interface{}

    // ConfigState. The type is interface{} with range: 0..4294967295. The default
    // value is 0.
    ConfigState interface{}

    // CardState. The type is interface{} with range: 0..4294967295. The default
    // value is 0.
    CardState interface{}

    // VM State information. The type is interface{} with range: 0..4294967295.
    // The default value is 0.
    VmState interface{}

    // Card Admin State. The type is interface{} with range: 0..4294967295. The
    // default value is 0.
    CardAdminState interface{}

    // CardType. The type is interface{} with range: 0..4294967295. The default
    // value is 0.
    CardType interface{}

    // Card Type String. The type is string.
    CardTypeString interface{}

    // Node Name String. The type is string.
    NodeNameString interface{}

    // Card Valid. The type is interface{} with range: 0..4294967295. The default
    // value is 0.
    CardValid interface{}

    // Pi Slot Number. The type is interface{} with range: 0..4294967295. The
    // default value is 0.
    PiSlotNumber interface{}

    // Shutdown. The type is interface{} with range: 0..4294967295. The default
    // value is 0.
    Shutdown interface{}

    // CType. The type is interface{} with range: 0..4294967295. The default value
    // is 0.
    Ctype interface{}

    // Card State String. The type is string.
    CardStateString interface{}

    // Monitor. The type is interface{} with range: 0..4294967295. The default
    // value is 0.
    Monitor interface{}
}

func (attributes *SdrInventory_Racks_Rack_Slot_Card_Attributes) GetEntityData() *types.CommonEntityData {
    attributes.EntityData.YFilter = attributes.YFilter
    attributes.EntityData.YangName = "attributes"
    attributes.EntityData.BundleName = "cisco_ios_xr"
    attributes.EntityData.ParentYangName = "card"
    attributes.EntityData.SegmentPath = "attributes"
    attributes.EntityData.AbsolutePath = "Cisco-IOS-XR-sdr-invmgr-oper:sdr-inventory/racks/rack/slot/card/" + attributes.EntityData.SegmentPath
    attributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attributes.EntityData.Children = types.NewOrderedMap()
    attributes.EntityData.Leafs = types.NewOrderedMap()
    attributes.EntityData.Leafs.Append("config-state-string", types.YLeaf{"ConfigStateString", attributes.ConfigStateString})
    attributes.EntityData.Leafs.Append("power", types.YLeaf{"Power", attributes.Power})
    attributes.EntityData.Leafs.Append("config-state", types.YLeaf{"ConfigState", attributes.ConfigState})
    attributes.EntityData.Leafs.Append("card-state", types.YLeaf{"CardState", attributes.CardState})
    attributes.EntityData.Leafs.Append("vm-state", types.YLeaf{"VmState", attributes.VmState})
    attributes.EntityData.Leafs.Append("card-admin-state", types.YLeaf{"CardAdminState", attributes.CardAdminState})
    attributes.EntityData.Leafs.Append("card-type", types.YLeaf{"CardType", attributes.CardType})
    attributes.EntityData.Leafs.Append("card-type-string", types.YLeaf{"CardTypeString", attributes.CardTypeString})
    attributes.EntityData.Leafs.Append("node-name-string", types.YLeaf{"NodeNameString", attributes.NodeNameString})
    attributes.EntityData.Leafs.Append("card-valid", types.YLeaf{"CardValid", attributes.CardValid})
    attributes.EntityData.Leafs.Append("pi-slot-number", types.YLeaf{"PiSlotNumber", attributes.PiSlotNumber})
    attributes.EntityData.Leafs.Append("shutdown", types.YLeaf{"Shutdown", attributes.Shutdown})
    attributes.EntityData.Leafs.Append("ctype", types.YLeaf{"Ctype", attributes.Ctype})
    attributes.EntityData.Leafs.Append("card-state-string", types.YLeaf{"CardStateString", attributes.CardStateString})
    attributes.EntityData.Leafs.Append("monitor", types.YLeaf{"Monitor", attributes.Monitor})

    attributes.EntityData.YListKeys = []string {}

    return &(attributes.EntityData)
}

// SdrInventory_Memory
// Memory
type SdrInventory_Memory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RackTable.
    Racks SdrInventory_Memory_Racks
}

func (memory *SdrInventory_Memory) GetEntityData() *types.CommonEntityData {
    memory.EntityData.YFilter = memory.YFilter
    memory.EntityData.YangName = "memory"
    memory.EntityData.BundleName = "cisco_ios_xr"
    memory.EntityData.ParentYangName = "sdr-inventory"
    memory.EntityData.SegmentPath = "memory"
    memory.EntityData.AbsolutePath = "Cisco-IOS-XR-sdr-invmgr-oper:sdr-inventory/" + memory.EntityData.SegmentPath
    memory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    memory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    memory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    memory.EntityData.Children = types.NewOrderedMap()
    memory.EntityData.Children.Append("racks", types.YChild{"Racks", &memory.Racks})
    memory.EntityData.Leafs = types.NewOrderedMap()

    memory.EntityData.YListKeys = []string {}

    return &(memory.EntityData)
}

// SdrInventory_Memory_Racks
// RackTable
type SdrInventory_Memory_Racks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Rack name. The type is slice of SdrInventory_Memory_Racks_Rack.
    Rack []*SdrInventory_Memory_Racks_Rack
}

func (racks *SdrInventory_Memory_Racks) GetEntityData() *types.CommonEntityData {
    racks.EntityData.YFilter = racks.YFilter
    racks.EntityData.YangName = "racks"
    racks.EntityData.BundleName = "cisco_ios_xr"
    racks.EntityData.ParentYangName = "memory"
    racks.EntityData.SegmentPath = "racks"
    racks.EntityData.AbsolutePath = "Cisco-IOS-XR-sdr-invmgr-oper:sdr-inventory/memory/" + racks.EntityData.SegmentPath
    racks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    racks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    racks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    racks.EntityData.Children = types.NewOrderedMap()
    racks.EntityData.Children.Append("rack", types.YChild{"Rack", nil})
    for i := range racks.Rack {
        racks.EntityData.Children.Append(types.GetSegmentPath(racks.Rack[i]), types.YChild{"Rack", racks.Rack[i]})
    }
    racks.EntityData.Leafs = types.NewOrderedMap()

    racks.EntityData.YListKeys = []string {}

    return &(racks.EntityData)
}

// SdrInventory_Memory_Racks_Rack
// Rack name
type SdrInventory_Memory_Racks_Rack struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Rack name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Name interface{}

    // NodeIDTable.
    NodeIds SdrInventory_Memory_Racks_Rack_NodeIds
}

func (rack *SdrInventory_Memory_Racks_Rack) GetEntityData() *types.CommonEntityData {
    rack.EntityData.YFilter = rack.YFilter
    rack.EntityData.YangName = "rack"
    rack.EntityData.BundleName = "cisco_ios_xr"
    rack.EntityData.ParentYangName = "racks"
    rack.EntityData.SegmentPath = "rack" + types.AddKeyToken(rack.Name, "name")
    rack.EntityData.AbsolutePath = "Cisco-IOS-XR-sdr-invmgr-oper:sdr-inventory/memory/racks/" + rack.EntityData.SegmentPath
    rack.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rack.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rack.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rack.EntityData.Children = types.NewOrderedMap()
    rack.EntityData.Children.Append("node-ids", types.YChild{"NodeIds", &rack.NodeIds})
    rack.EntityData.Leafs = types.NewOrderedMap()
    rack.EntityData.Leafs.Append("name", types.YLeaf{"Name", rack.Name})

    rack.EntityData.YListKeys = []string {"Name"}

    return &(rack.EntityData)
}

// SdrInventory_Memory_Racks_Rack_NodeIds
// NodeIDTable
type SdrInventory_Memory_Racks_Rack_NodeIds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Inventory Admin Mempool Data  Bag. The type is slice of
    // SdrInventory_Memory_Racks_Rack_NodeIds_NodeId.
    NodeId []*SdrInventory_Memory_Racks_Rack_NodeIds_NodeId
}

func (nodeIds *SdrInventory_Memory_Racks_Rack_NodeIds) GetEntityData() *types.CommonEntityData {
    nodeIds.EntityData.YFilter = nodeIds.YFilter
    nodeIds.EntityData.YangName = "node-ids"
    nodeIds.EntityData.BundleName = "cisco_ios_xr"
    nodeIds.EntityData.ParentYangName = "rack"
    nodeIds.EntityData.SegmentPath = "node-ids"
    nodeIds.EntityData.AbsolutePath = "Cisco-IOS-XR-sdr-invmgr-oper:sdr-inventory/memory/racks/rack/" + nodeIds.EntityData.SegmentPath
    nodeIds.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeIds.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeIds.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeIds.EntityData.Children = types.NewOrderedMap()
    nodeIds.EntityData.Children.Append("node-id", types.YChild{"NodeId", nil})
    for i := range nodeIds.NodeId {
        nodeIds.EntityData.Children.Append(types.GetSegmentPath(nodeIds.NodeId[i]), types.YChild{"NodeId", nodeIds.NodeId[i]})
    }
    nodeIds.EntityData.Leafs = types.NewOrderedMap()

    nodeIds.EntityData.YListKeys = []string {}

    return &(nodeIds.EntityData)
}

// SdrInventory_Memory_Racks_Rack_NodeIds_NodeId
// Inventory Admin Mempool Data  Bag
type SdrInventory_Memory_Racks_Rack_NodeIds_NodeId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. nodeid. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Name interface{}

    // Total Memory. The type is interface{} with range: 0..4294967295.
    TotalMemory interface{}

    // Avaialble Memory. The type is interface{} with range: 0..4294967295.
    AvailableMemory interface{}
}

func (nodeId *SdrInventory_Memory_Racks_Rack_NodeIds_NodeId) GetEntityData() *types.CommonEntityData {
    nodeId.EntityData.YFilter = nodeId.YFilter
    nodeId.EntityData.YangName = "node-id"
    nodeId.EntityData.BundleName = "cisco_ios_xr"
    nodeId.EntityData.ParentYangName = "node-ids"
    nodeId.EntityData.SegmentPath = "node-id" + types.AddKeyToken(nodeId.Name, "name")
    nodeId.EntityData.AbsolutePath = "Cisco-IOS-XR-sdr-invmgr-oper:sdr-inventory/memory/racks/rack/node-ids/" + nodeId.EntityData.SegmentPath
    nodeId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeId.EntityData.Children = types.NewOrderedMap()
    nodeId.EntityData.Leafs = types.NewOrderedMap()
    nodeId.EntityData.Leafs.Append("name", types.YLeaf{"Name", nodeId.Name})
    nodeId.EntityData.Leafs.Append("total-memory", types.YLeaf{"TotalMemory", nodeId.TotalMemory})
    nodeId.EntityData.Leafs.Append("available-memory", types.YLeaf{"AvailableMemory", nodeId.AvailableMemory})

    nodeId.EntityData.YListKeys = []string {"Name"}

    return &(nodeId.EntityData)
}

