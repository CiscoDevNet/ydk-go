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
    parent types.Entity
    YFilter yfilter.YFilter

    // RackTable.
    Racks SdrInventory_Racks
}

func (sdrInventory *SdrInventory) GetFilter() yfilter.YFilter { return sdrInventory.YFilter }

func (sdrInventory *SdrInventory) SetFilter(yf yfilter.YFilter) { sdrInventory.YFilter = yf }

func (sdrInventory *SdrInventory) GetGoName(yname string) string {
    if yname == "racks" { return "Racks" }
    return ""
}

func (sdrInventory *SdrInventory) GetSegmentPath() string {
    return "Cisco-IOS-XR-sdr-invmgr-oper:sdr-inventory"
}

func (sdrInventory *SdrInventory) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "racks" {
        return &sdrInventory.Racks
    }
    return nil
}

func (sdrInventory *SdrInventory) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["racks"] = &sdrInventory.Racks
    return children
}

func (sdrInventory *SdrInventory) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sdrInventory *SdrInventory) GetBundleName() string { return "cisco_ios_xr" }

func (sdrInventory *SdrInventory) GetYangName() string { return "sdr-inventory" }

func (sdrInventory *SdrInventory) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sdrInventory *SdrInventory) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sdrInventory *SdrInventory) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sdrInventory *SdrInventory) SetParent(parent types.Entity) { sdrInventory.parent = parent }

func (sdrInventory *SdrInventory) GetParent() types.Entity { return sdrInventory.parent }

func (sdrInventory *SdrInventory) GetParentYangName() string { return "Cisco-IOS-XR-sdr-invmgr-oper" }

// SdrInventory_Racks
// RackTable
type SdrInventory_Racks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Rack name. The type is slice of SdrInventory_Racks_Rack.
    Rack []SdrInventory_Racks_Rack
}

func (racks *SdrInventory_Racks) GetFilter() yfilter.YFilter { return racks.YFilter }

func (racks *SdrInventory_Racks) SetFilter(yf yfilter.YFilter) { racks.YFilter = yf }

func (racks *SdrInventory_Racks) GetGoName(yname string) string {
    if yname == "rack" { return "Rack" }
    return ""
}

func (racks *SdrInventory_Racks) GetSegmentPath() string {
    return "racks"
}

func (racks *SdrInventory_Racks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rack" {
        for _, c := range racks.Rack {
            if racks.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SdrInventory_Racks_Rack{}
        racks.Rack = append(racks.Rack, child)
        return &racks.Rack[len(racks.Rack)-1]
    }
    return nil
}

func (racks *SdrInventory_Racks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range racks.Rack {
        children[racks.Rack[i].GetSegmentPath()] = &racks.Rack[i]
    }
    return children
}

func (racks *SdrInventory_Racks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (racks *SdrInventory_Racks) GetBundleName() string { return "cisco_ios_xr" }

func (racks *SdrInventory_Racks) GetYangName() string { return "racks" }

func (racks *SdrInventory_Racks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (racks *SdrInventory_Racks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (racks *SdrInventory_Racks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (racks *SdrInventory_Racks) SetParent(parent types.Entity) { racks.parent = parent }

func (racks *SdrInventory_Racks) GetParent() types.Entity { return racks.parent }

func (racks *SdrInventory_Racks) GetParentYangName() string { return "sdr-inventory" }

// SdrInventory_Racks_Rack
// Rack name
type SdrInventory_Racks_Rack struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Rack name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // Slot name. The type is slice of SdrInventory_Racks_Rack_Slot.
    Slot []SdrInventory_Racks_Rack_Slot
}

func (rack *SdrInventory_Racks_Rack) GetFilter() yfilter.YFilter { return rack.YFilter }

func (rack *SdrInventory_Racks_Rack) SetFilter(yf yfilter.YFilter) { rack.YFilter = yf }

func (rack *SdrInventory_Racks_Rack) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "slot" { return "Slot" }
    return ""
}

func (rack *SdrInventory_Racks_Rack) GetSegmentPath() string {
    return "rack" + "[name='" + fmt.Sprintf("%v", rack.Name) + "']"
}

func (rack *SdrInventory_Racks_Rack) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "slot" {
        for _, c := range rack.Slot {
            if rack.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SdrInventory_Racks_Rack_Slot{}
        rack.Slot = append(rack.Slot, child)
        return &rack.Slot[len(rack.Slot)-1]
    }
    return nil
}

func (rack *SdrInventory_Racks_Rack) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range rack.Slot {
        children[rack.Slot[i].GetSegmentPath()] = &rack.Slot[i]
    }
    return children
}

func (rack *SdrInventory_Racks_Rack) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = rack.Name
    return leafs
}

func (rack *SdrInventory_Racks_Rack) GetBundleName() string { return "cisco_ios_xr" }

func (rack *SdrInventory_Racks_Rack) GetYangName() string { return "rack" }

func (rack *SdrInventory_Racks_Rack) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rack *SdrInventory_Racks_Rack) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rack *SdrInventory_Racks_Rack) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rack *SdrInventory_Racks_Rack) SetParent(parent types.Entity) { rack.parent = parent }

func (rack *SdrInventory_Racks_Rack) GetParent() types.Entity { return rack.parent }

func (rack *SdrInventory_Racks_Rack) GetParentYangName() string { return "racks" }

// SdrInventory_Racks_Rack_Slot
// Slot name
type SdrInventory_Racks_Rack_Slot struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Slot name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // Card. The type is slice of SdrInventory_Racks_Rack_Slot_Card.
    Card []SdrInventory_Racks_Rack_Slot_Card
}

func (slot *SdrInventory_Racks_Rack_Slot) GetFilter() yfilter.YFilter { return slot.YFilter }

func (slot *SdrInventory_Racks_Rack_Slot) SetFilter(yf yfilter.YFilter) { slot.YFilter = yf }

func (slot *SdrInventory_Racks_Rack_Slot) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "card" { return "Card" }
    return ""
}

func (slot *SdrInventory_Racks_Rack_Slot) GetSegmentPath() string {
    return "slot" + "[name='" + fmt.Sprintf("%v", slot.Name) + "']"
}

func (slot *SdrInventory_Racks_Rack_Slot) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "card" {
        for _, c := range slot.Card {
            if slot.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SdrInventory_Racks_Rack_Slot_Card{}
        slot.Card = append(slot.Card, child)
        return &slot.Card[len(slot.Card)-1]
    }
    return nil
}

func (slot *SdrInventory_Racks_Rack_Slot) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range slot.Card {
        children[slot.Card[i].GetSegmentPath()] = &slot.Card[i]
    }
    return children
}

func (slot *SdrInventory_Racks_Rack_Slot) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = slot.Name
    return leafs
}

func (slot *SdrInventory_Racks_Rack_Slot) GetBundleName() string { return "cisco_ios_xr" }

func (slot *SdrInventory_Racks_Rack_Slot) GetYangName() string { return "slot" }

func (slot *SdrInventory_Racks_Rack_Slot) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (slot *SdrInventory_Racks_Rack_Slot) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (slot *SdrInventory_Racks_Rack_Slot) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (slot *SdrInventory_Racks_Rack_Slot) SetParent(parent types.Entity) { slot.parent = parent }

func (slot *SdrInventory_Racks_Rack_Slot) GetParent() types.Entity { return slot.parent }

func (slot *SdrInventory_Racks_Rack_Slot) GetParentYangName() string { return "rack" }

// SdrInventory_Racks_Rack_Slot_Card
// Card
type SdrInventory_Racks_Rack_Slot_Card struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Card. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // Attributes.
    Attributes SdrInventory_Racks_Rack_Slot_Card_Attributes
}

func (card *SdrInventory_Racks_Rack_Slot_Card) GetFilter() yfilter.YFilter { return card.YFilter }

func (card *SdrInventory_Racks_Rack_Slot_Card) SetFilter(yf yfilter.YFilter) { card.YFilter = yf }

func (card *SdrInventory_Racks_Rack_Slot_Card) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "attributes" { return "Attributes" }
    return ""
}

func (card *SdrInventory_Racks_Rack_Slot_Card) GetSegmentPath() string {
    return "card" + "[name='" + fmt.Sprintf("%v", card.Name) + "']"
}

func (card *SdrInventory_Racks_Rack_Slot_Card) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "attributes" {
        return &card.Attributes
    }
    return nil
}

func (card *SdrInventory_Racks_Rack_Slot_Card) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["attributes"] = &card.Attributes
    return children
}

func (card *SdrInventory_Racks_Rack_Slot_Card) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = card.Name
    return leafs
}

func (card *SdrInventory_Racks_Rack_Slot_Card) GetBundleName() string { return "cisco_ios_xr" }

func (card *SdrInventory_Racks_Rack_Slot_Card) GetYangName() string { return "card" }

func (card *SdrInventory_Racks_Rack_Slot_Card) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (card *SdrInventory_Racks_Rack_Slot_Card) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (card *SdrInventory_Racks_Rack_Slot_Card) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (card *SdrInventory_Racks_Rack_Slot_Card) SetParent(parent types.Entity) { card.parent = parent }

func (card *SdrInventory_Racks_Rack_Slot_Card) GetParent() types.Entity { return card.parent }

func (card *SdrInventory_Racks_Rack_Slot_Card) GetParentYangName() string { return "slot" }

// SdrInventory_Racks_Rack_Slot_Card_Attributes
// Attributes
type SdrInventory_Racks_Rack_Slot_Card_Attributes struct {
    parent types.Entity
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

func (attributes *SdrInventory_Racks_Rack_Slot_Card_Attributes) GetFilter() yfilter.YFilter { return attributes.YFilter }

func (attributes *SdrInventory_Racks_Rack_Slot_Card_Attributes) SetFilter(yf yfilter.YFilter) { attributes.YFilter = yf }

func (attributes *SdrInventory_Racks_Rack_Slot_Card_Attributes) GetGoName(yname string) string {
    if yname == "config-state-string" { return "ConfigStateString" }
    if yname == "power" { return "Power" }
    if yname == "config-state" { return "ConfigState" }
    if yname == "card-state" { return "CardState" }
    if yname == "vm-state" { return "VmState" }
    if yname == "card-admin-state" { return "CardAdminState" }
    if yname == "card-type" { return "CardType" }
    if yname == "card-type-string" { return "CardTypeString" }
    if yname == "node-name-string" { return "NodeNameString" }
    if yname == "pi-slot-number" { return "PiSlotNumber" }
    if yname == "shutdown" { return "Shutdown" }
    if yname == "ctype" { return "Ctype" }
    if yname == "card-state-string" { return "CardStateString" }
    if yname == "monitor" { return "Monitor" }
    return ""
}

func (attributes *SdrInventory_Racks_Rack_Slot_Card_Attributes) GetSegmentPath() string {
    return "attributes"
}

func (attributes *SdrInventory_Racks_Rack_Slot_Card_Attributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (attributes *SdrInventory_Racks_Rack_Slot_Card_Attributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (attributes *SdrInventory_Racks_Rack_Slot_Card_Attributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["config-state-string"] = attributes.ConfigStateString
    leafs["power"] = attributes.Power
    leafs["config-state"] = attributes.ConfigState
    leafs["card-state"] = attributes.CardState
    leafs["vm-state"] = attributes.VmState
    leafs["card-admin-state"] = attributes.CardAdminState
    leafs["card-type"] = attributes.CardType
    leafs["card-type-string"] = attributes.CardTypeString
    leafs["node-name-string"] = attributes.NodeNameString
    leafs["pi-slot-number"] = attributes.PiSlotNumber
    leafs["shutdown"] = attributes.Shutdown
    leafs["ctype"] = attributes.Ctype
    leafs["card-state-string"] = attributes.CardStateString
    leafs["monitor"] = attributes.Monitor
    return leafs
}

func (attributes *SdrInventory_Racks_Rack_Slot_Card_Attributes) GetBundleName() string { return "cisco_ios_xr" }

func (attributes *SdrInventory_Racks_Rack_Slot_Card_Attributes) GetYangName() string { return "attributes" }

func (attributes *SdrInventory_Racks_Rack_Slot_Card_Attributes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (attributes *SdrInventory_Racks_Rack_Slot_Card_Attributes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (attributes *SdrInventory_Racks_Rack_Slot_Card_Attributes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (attributes *SdrInventory_Racks_Rack_Slot_Card_Attributes) SetParent(parent types.Entity) { attributes.parent = parent }

func (attributes *SdrInventory_Racks_Rack_Slot_Card_Attributes) GetParent() types.Entity { return attributes.parent }

func (attributes *SdrInventory_Racks_Rack_Slot_Card_Attributes) GetParentYangName() string { return "card" }

