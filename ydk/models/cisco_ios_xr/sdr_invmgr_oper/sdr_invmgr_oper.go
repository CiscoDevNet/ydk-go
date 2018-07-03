// This module contains a collection of YANG definitions
// for Cisco IOS-XR sdr-invmgr package operational data.
// 
// This module contains definitions
// for the following management objects:
//   sdr-inventory: SDR information
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
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
}

func (sdrInventory *SdrInventory) GetEntityData() *types.CommonEntityData {
    sdrInventory.EntityData.YFilter = sdrInventory.YFilter
    sdrInventory.EntityData.YangName = "sdr-inventory"
    sdrInventory.EntityData.BundleName = "cisco_ios_xr"
    sdrInventory.EntityData.ParentYangName = "Cisco-IOS-XR-sdr-invmgr-oper"
    sdrInventory.EntityData.SegmentPath = "Cisco-IOS-XR-sdr-invmgr-oper:sdr-inventory"
    sdrInventory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sdrInventory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sdrInventory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sdrInventory.EntityData.Children = types.NewOrderedMap()
    sdrInventory.EntityData.Children.Append("racks", types.YChild{"Racks", &sdrInventory.Racks})
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

    // This attribute is a key. Rack name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
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

    // This attribute is a key. Slot name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
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

    // This attribute is a key. Card. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
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
    attributes.EntityData.Leafs.Append("pi-slot-number", types.YLeaf{"PiSlotNumber", attributes.PiSlotNumber})
    attributes.EntityData.Leafs.Append("shutdown", types.YLeaf{"Shutdown", attributes.Shutdown})
    attributes.EntityData.Leafs.Append("ctype", types.YLeaf{"Ctype", attributes.Ctype})
    attributes.EntityData.Leafs.Append("card-state-string", types.YLeaf{"CardStateString", attributes.CardStateString})
    attributes.EntityData.Leafs.Append("monitor", types.YLeaf{"Monitor", attributes.Monitor})

    attributes.EntityData.YListKeys = []string {}

    return &(attributes.EntityData)
}

