// This module contains a collection of YANG definitions
// for Cisco IOS-XR asr9k-sc-invmgr package operational data.
// 
// This module contains definitions
// for the following management objects:
//   inventory: Logical Router Inventory operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package asr9k_sc_invmgr_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package asr9k_sc_invmgr_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-asr9k-sc-invmgr-oper inventory}", reflect.TypeOf(Inventory{}))
    ydk.RegisterEntity("Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory", reflect.TypeOf(Inventory{}))
}

// CardResetReason represents Card reset reason
type CardResetReason string

const (
    // reset unknown
    CardResetReason_reset_unknown CardResetReason = "reset-unknown"

    // power up
    CardResetReason_power_up CardResetReason = "power-up"

    // parity error
    CardResetReason_parity_error CardResetReason = "parity-error"

    // clear config reset
    CardResetReason_clear_config_reset CardResetReason = "clear-config-reset"

    // manual reset
    CardResetReason_manual_reset CardResetReason = "manual-reset"

    // watch dog timeout reset
    CardResetReason_watch_dog_timeout_reset CardResetReason = "watch-dog-timeout-reset"

    // resource overflow reset
    CardResetReason_resource_overflow_reset CardResetReason = "resource-overflow-reset"

    // missing task reset
    CardResetReason_missing_task_reset CardResetReason = "missing-task-reset"

    // low voltage reset
    CardResetReason_low_voltage_reset CardResetReason = "low-voltage-reset"

    // controller reset
    CardResetReason_controller_reset CardResetReason = "controller-reset"

    // system reset
    CardResetReason_system_reset CardResetReason = "system-reset"

    // switchover reset
    CardResetReason_switchover_reset CardResetReason = "switchover-reset"

    // upgrade reset
    CardResetReason_upgrade_reset CardResetReason = "upgrade-reset"

    // downgrade reset
    CardResetReason_downgrade_reset CardResetReason = "downgrade-reset"

    // cache error reset
    CardResetReason_cache_error_reset CardResetReason = "cache-error-reset"

    // device driver reset
    CardResetReason_device_driver_reset CardResetReason = "device-driver-reset"

    // software exception reset
    CardResetReason_software_exception_reset CardResetReason = "software-exception-reset"

    // restore config reset
    CardResetReason_restore_config_reset CardResetReason = "restore-config-reset"

    // abort rev reset
    CardResetReason_abort_rev_reset CardResetReason = "abort-rev-reset"

    // burn boot reset
    CardResetReason_burn_boot_reset CardResetReason = "burn-boot-reset"

    // standby cd healthier reset
    CardResetReason_standby_cd_healthier_reset CardResetReason = "standby-cd-healthier-reset"

    // non native config clear reset
    CardResetReason_non_native_config_clear_reset CardResetReason = "non-native-config-clear-reset"

    // memory protection error reset
    CardResetReason_memory_protection_error_reset CardResetReason = "memory-protection-error-reset"

    // card reset reason max
    CardResetReason_card_reset_reason_max CardResetReason = "card-reset-reason-max"
)

// Inventory
// Logical Router Inventory operational data
type Inventory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of racks.
    Racks Inventory_Racks
}

func (inventory *Inventory) GetEntityData() *types.CommonEntityData {
    inventory.EntityData.YFilter = inventory.YFilter
    inventory.EntityData.YangName = "inventory"
    inventory.EntityData.BundleName = "cisco_ios_xr"
    inventory.EntityData.ParentYangName = "Cisco-IOS-XR-asr9k-sc-invmgr-oper"
    inventory.EntityData.SegmentPath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory"
    inventory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inventory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inventory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inventory.EntityData.Children = make(map[string]types.YChild)
    inventory.EntityData.Children["racks"] = types.YChild{"Racks", &inventory.Racks}
    inventory.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(inventory.EntityData)
}

// Inventory_Racks
// Table of racks
type Inventory_Racks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Rack number. The type is slice of Inventory_Racks_Rack.
    Rack []Inventory_Racks_Rack
}

func (racks *Inventory_Racks) GetEntityData() *types.CommonEntityData {
    racks.EntityData.YFilter = racks.YFilter
    racks.EntityData.YangName = "racks"
    racks.EntityData.BundleName = "cisco_ios_xr"
    racks.EntityData.ParentYangName = "inventory"
    racks.EntityData.SegmentPath = "racks"
    racks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    racks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    racks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    racks.EntityData.Children = make(map[string]types.YChild)
    racks.EntityData.Children["rack"] = types.YChild{"Rack", nil}
    for i := range racks.Rack {
        racks.EntityData.Children[types.GetSegmentPath(&racks.Rack[i])] = types.YChild{"Rack", &racks.Rack[i]}
    }
    racks.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(racks.EntityData)
}

// Inventory_Racks_Rack
// Rack number
type Inventory_Racks_Rack struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Rack number. The type is interface{} with range:
    // -2147483648..2147483647.
    Number interface{}

    // Slot table contains all slots in the rack.
    Slots Inventory_Racks_Rack_Slots
}

func (rack *Inventory_Racks_Rack) GetEntityData() *types.CommonEntityData {
    rack.EntityData.YFilter = rack.YFilter
    rack.EntityData.YangName = "rack"
    rack.EntityData.BundleName = "cisco_ios_xr"
    rack.EntityData.ParentYangName = "racks"
    rack.EntityData.SegmentPath = "rack" + "[number='" + fmt.Sprintf("%v", rack.Number) + "']"
    rack.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rack.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rack.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rack.EntityData.Children = make(map[string]types.YChild)
    rack.EntityData.Children["slots"] = types.YChild{"Slots", &rack.Slots}
    rack.EntityData.Leafs = make(map[string]types.YLeaf)
    rack.EntityData.Leafs["number"] = types.YLeaf{"Number", rack.Number}
    return &(rack.EntityData)
}

// Inventory_Racks_Rack_Slots
// Slot table contains all slots in the rack
type Inventory_Racks_Rack_Slots struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Slot number. The type is slice of Inventory_Racks_Rack_Slots_Slot.
    Slot []Inventory_Racks_Rack_Slots_Slot
}

func (slots *Inventory_Racks_Rack_Slots) GetEntityData() *types.CommonEntityData {
    slots.EntityData.YFilter = slots.YFilter
    slots.EntityData.YangName = "slots"
    slots.EntityData.BundleName = "cisco_ios_xr"
    slots.EntityData.ParentYangName = "rack"
    slots.EntityData.SegmentPath = "slots"
    slots.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    slots.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    slots.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    slots.EntityData.Children = make(map[string]types.YChild)
    slots.EntityData.Children["slot"] = types.YChild{"Slot", nil}
    for i := range slots.Slot {
        slots.EntityData.Children[types.GetSegmentPath(&slots.Slot[i])] = types.YChild{"Slot", &slots.Slot[i]}
    }
    slots.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(slots.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot
// Slot number
type Inventory_Racks_Rack_Slots_Slot struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Slot number. The type is interface{} with range:
    // -2147483648..2147483647.
    Number interface{}

    // Card table contains all cards in the slot.
    Cards Inventory_Racks_Rack_Slots_Slot_Cards

    // Attributes.
    BasicAttributes Inventory_Racks_Rack_Slots_Slot_BasicAttributes
}

func (slot *Inventory_Racks_Rack_Slots_Slot) GetEntityData() *types.CommonEntityData {
    slot.EntityData.YFilter = slot.YFilter
    slot.EntityData.YangName = "slot"
    slot.EntityData.BundleName = "cisco_ios_xr"
    slot.EntityData.ParentYangName = "slots"
    slot.EntityData.SegmentPath = "slot" + "[number='" + fmt.Sprintf("%v", slot.Number) + "']"
    slot.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    slot.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    slot.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    slot.EntityData.Children = make(map[string]types.YChild)
    slot.EntityData.Children["cards"] = types.YChild{"Cards", &slot.Cards}
    slot.EntityData.Children["basic-attributes"] = types.YChild{"BasicAttributes", &slot.BasicAttributes}
    slot.EntityData.Leafs = make(map[string]types.YLeaf)
    slot.EntityData.Leafs["number"] = types.YLeaf{"Number", slot.Number}
    return &(slot.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards
// Card table contains all cards in the slot
type Inventory_Racks_Rack_Slots_Slot_Cards struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Card number. The type is slice of
    // Inventory_Racks_Rack_Slots_Slot_Cards_Card.
    Card []Inventory_Racks_Rack_Slots_Slot_Cards_Card
}

func (cards *Inventory_Racks_Rack_Slots_Slot_Cards) GetEntityData() *types.CommonEntityData {
    cards.EntityData.YFilter = cards.YFilter
    cards.EntityData.YangName = "cards"
    cards.EntityData.BundleName = "cisco_ios_xr"
    cards.EntityData.ParentYangName = "slot"
    cards.EntityData.SegmentPath = "cards"
    cards.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cards.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cards.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cards.EntityData.Children = make(map[string]types.YChild)
    cards.EntityData.Children["card"] = types.YChild{"Card", nil}
    for i := range cards.Card {
        cards.EntityData.Children[types.GetSegmentPath(&cards.Card[i])] = types.YChild{"Card", &cards.Card[i]}
    }
    cards.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cards.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card
// Card number
type Inventory_Racks_Rack_Slots_Slot_Cards_Card struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. card number. The type is interface{} with range:
    // -2147483648..2147483647.
    Number interface{}

    // SubSlotTable contains all subslots in a Slot.
    SubSlots Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots

    // HWComponent table contains all HW modules within the card .
    HwComponents Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents

    // ModuleSensorTable contains all sensors in a Module.
    Sensors Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors

    // PortSlotTable contains all optics ports in a SPA/PLIM.
    PortSlots Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots

    // Attributes.
    BasicAttributes Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes
}

func (card *Inventory_Racks_Rack_Slots_Slot_Cards_Card) GetEntityData() *types.CommonEntityData {
    card.EntityData.YFilter = card.YFilter
    card.EntityData.YangName = "card"
    card.EntityData.BundleName = "cisco_ios_xr"
    card.EntityData.ParentYangName = "cards"
    card.EntityData.SegmentPath = "card" + "[number='" + fmt.Sprintf("%v", card.Number) + "']"
    card.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    card.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    card.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    card.EntityData.Children = make(map[string]types.YChild)
    card.EntityData.Children["sub-slots"] = types.YChild{"SubSlots", &card.SubSlots}
    card.EntityData.Children["hw-components"] = types.YChild{"HwComponents", &card.HwComponents}
    card.EntityData.Children["sensors"] = types.YChild{"Sensors", &card.Sensors}
    card.EntityData.Children["port-slots"] = types.YChild{"PortSlots", &card.PortSlots}
    card.EntityData.Children["basic-attributes"] = types.YChild{"BasicAttributes", &card.BasicAttributes}
    card.EntityData.Leafs = make(map[string]types.YLeaf)
    card.EntityData.Leafs["number"] = types.YLeaf{"Number", card.Number}
    return &(card.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots
// SubSlotTable contains all subslots in a
// Slot
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SubSlot number. The type is slice of
    // Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot.
    SubSlot []Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot
}

func (subSlots *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots) GetEntityData() *types.CommonEntityData {
    subSlots.EntityData.YFilter = subSlots.YFilter
    subSlots.EntityData.YangName = "sub-slots"
    subSlots.EntityData.BundleName = "cisco_ios_xr"
    subSlots.EntityData.ParentYangName = "card"
    subSlots.EntityData.SegmentPath = "sub-slots"
    subSlots.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    subSlots.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    subSlots.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    subSlots.EntityData.Children = make(map[string]types.YChild)
    subSlots.EntityData.Children["sub-slot"] = types.YChild{"SubSlot", nil}
    for i := range subSlots.SubSlot {
        subSlots.EntityData.Children[types.GetSegmentPath(&subSlots.SubSlot[i])] = types.YChild{"SubSlot", &subSlots.SubSlot[i]}
    }
    subSlots.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(subSlots.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot
// SubSlot number
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. subslot number. The type is interface{} with
    // range: -2147483648..2147483647.
    Number interface{}

    // Module string.
    Module Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module

    // Attributes.
    BasicAttributes Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes
}

func (subSlot *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot) GetEntityData() *types.CommonEntityData {
    subSlot.EntityData.YFilter = subSlot.YFilter
    subSlot.EntityData.YangName = "sub-slot"
    subSlot.EntityData.BundleName = "cisco_ios_xr"
    subSlot.EntityData.ParentYangName = "sub-slots"
    subSlot.EntityData.SegmentPath = "sub-slot" + "[number='" + fmt.Sprintf("%v", subSlot.Number) + "']"
    subSlot.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    subSlot.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    subSlot.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    subSlot.EntityData.Children = make(map[string]types.YChild)
    subSlot.EntityData.Children["module"] = types.YChild{"Module", &subSlot.Module}
    subSlot.EntityData.Children["basic-attributes"] = types.YChild{"BasicAttributes", &subSlot.BasicAttributes}
    subSlot.EntityData.Leafs = make(map[string]types.YLeaf)
    subSlot.EntityData.Leafs["number"] = types.YLeaf{"Number", subSlot.Number}
    return &(subSlot.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module
// Module string
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ModuleSensorTable contains all sensors in a Module.
    Sensors Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors

    // PortSlotTable contains all optics ports in a SPA/PLIM.
    PortSlots Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots

    // Attributes.
    BasicAttributes Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes
}

func (module *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module) GetEntityData() *types.CommonEntityData {
    module.EntityData.YFilter = module.YFilter
    module.EntityData.YangName = "module"
    module.EntityData.BundleName = "cisco_ios_xr"
    module.EntityData.ParentYangName = "sub-slot"
    module.EntityData.SegmentPath = "module"
    module.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    module.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    module.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    module.EntityData.Children = make(map[string]types.YChild)
    module.EntityData.Children["sensors"] = types.YChild{"Sensors", &module.Sensors}
    module.EntityData.Children["port-slots"] = types.YChild{"PortSlots", &module.PortSlots}
    module.EntityData.Children["basic-attributes"] = types.YChild{"BasicAttributes", &module.BasicAttributes}
    module.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(module.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors
// ModuleSensorTable contains all sensors in a
// Module.
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sensor number in the Module. The type is slice of
    // Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor.
    Sensor []Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor
}

func (sensors *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors) GetEntityData() *types.CommonEntityData {
    sensors.EntityData.YFilter = sensors.YFilter
    sensors.EntityData.YangName = "sensors"
    sensors.EntityData.BundleName = "cisco_ios_xr"
    sensors.EntityData.ParentYangName = "module"
    sensors.EntityData.SegmentPath = "sensors"
    sensors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sensors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sensors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sensors.EntityData.Children = make(map[string]types.YChild)
    sensors.EntityData.Children["sensor"] = types.YChild{"Sensor", nil}
    for i := range sensors.Sensor {
        sensors.EntityData.Children[types.GetSegmentPath(&sensors.Sensor[i])] = types.YChild{"Sensor", &sensors.Sensor[i]}
    }
    sensors.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sensors.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor
// Sensor number in the Module
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. sensor number. The type is interface{} with range:
    // -2147483648..2147483647.
    Number interface{}

    // Attributes.
    BasicAttributes Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes
}

func (sensor *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor) GetEntityData() *types.CommonEntityData {
    sensor.EntityData.YFilter = sensor.YFilter
    sensor.EntityData.YangName = "sensor"
    sensor.EntityData.BundleName = "cisco_ios_xr"
    sensor.EntityData.ParentYangName = "sensors"
    sensor.EntityData.SegmentPath = "sensor" + "[number='" + fmt.Sprintf("%v", sensor.Number) + "']"
    sensor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sensor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sensor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sensor.EntityData.Children = make(map[string]types.YChild)
    sensor.EntityData.Children["basic-attributes"] = types.YChild{"BasicAttributes", &sensor.BasicAttributes}
    sensor.EntityData.Leafs = make(map[string]types.YLeaf)
    sensor.EntityData.Leafs["number"] = types.YLeaf{"Number", sensor.Number}
    return &(sensor.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes
// Attributes
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Inventory information.
    BasicInfo Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_BasicInfo

    // Field Replaceable Unit (FRU) inventory information.
    FruInfo Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_FruInfo
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes) GetEntityData() *types.CommonEntityData {
    basicAttributes.EntityData.YFilter = basicAttributes.YFilter
    basicAttributes.EntityData.YangName = "basic-attributes"
    basicAttributes.EntityData.BundleName = "cisco_ios_xr"
    basicAttributes.EntityData.ParentYangName = "sensor"
    basicAttributes.EntityData.SegmentPath = "basic-attributes"
    basicAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicAttributes.EntityData.Children = make(map[string]types.YChild)
    basicAttributes.EntityData.Children["basic-info"] = types.YChild{"BasicInfo", &basicAttributes.BasicInfo}
    basicAttributes.EntityData.Children["fru-info"] = types.YChild{"FruInfo", &basicAttributes.FruInfo}
    basicAttributes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(basicAttributes.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_BasicInfo
// Inventory information
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_BasicInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // describes in user-readable terms       what the entity in question does.
    // The type is string with length: 0..255.
    Description interface{}

    // maps to the vendor OID string. The type is string with length: 0..255.
    VendorType interface{}

    // name string for the entity. The type is string with length: 0..255.
    Name interface{}

    // hw revision string. The type is string with length: 0..255.
    HardwareRevision interface{}

    // firmware revision string. The type is string with length: 0..255.
    FirmwareRevision interface{}

    // software revision string. The type is string with length: 0..255.
    SoftwareRevision interface{}

    // chip module hw revision string. The type is string with length: 0..255.
    ChipHardwareRevision interface{}

    // serial number. The type is string with length: 0..255.
    SerialNumber interface{}

    // manufacturer's name. The type is string with length: 0..255.
    ManufacturerName interface{}

    // model name. The type is string with length: 0..255.
    ModelName interface{}

    // asset Identification string. The type is string with length: 0..255.
    AssetIdStr interface{}

    // asset Identification. The type is interface{} with range:
    // -2147483648..2147483647.
    AssetIdentification interface{}

    // 1 if Field Replaceable Unit 0, if not. The type is bool.
    IsFieldReplaceableUnit interface{}

    // Manufacture Asset Tags. The type is interface{} with range:
    // -2147483648..2147483647.
    ManufacturerAssetTags interface{}

    // Major&minor class of the entity. The type is interface{} with range:
    // -2147483648..2147483647.
    CompositeClassCode interface{}

    // Size of memory associated with       the entity where applicable. The type
    // is interface{} with range: -2147483648..2147483647.
    MemorySize interface{}

    // sysdb name of sensor in the envmon EDM. The type is string with length:
    // 0..255.
    EnvironmentalMonitorPath interface{}

    // useful for storing an entity alias . The type is string with length:
    // 0..255.
    Alias interface{}

    // indicates if this entity is group       or not. The type is bool.
    GroupFlag interface{}

    // integer value for New Deviation Number 0x88. The type is interface{} with
    // range: -2147483648..2147483647.
    NewDeviationNumber interface{}

    // integer value for plim type if     applicable to this entity. The type is
    // interface{} with range: -2147483648..2147483647.
    PhysicalLayerInterfaceModuleType interface{}

    // 1 if UnrecognizedFRU and 0 for recognizedFRU. The type is bool.
    UnrecognizedFru interface{}

    // integer value for Redundancy State if     applicable to this entity. The
    // type is interface{} with range: -2147483648..2147483647.
    Redundancystate interface{}

    // 1 if ce port found, 0 if not. The type is bool.
    Ceport interface{}

    // 1 if xr scoped, 0 if not. The type is bool.
    XrScoped interface{}

    // Unique id for an entity. The type is interface{} with range:
    // -2147483648..2147483647.
    UniqueId interface{}
}

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_BasicInfo) GetEntityData() *types.CommonEntityData {
    basicInfo.EntityData.YFilter = basicInfo.YFilter
    basicInfo.EntityData.YangName = "basic-info"
    basicInfo.EntityData.BundleName = "cisco_ios_xr"
    basicInfo.EntityData.ParentYangName = "basic-attributes"
    basicInfo.EntityData.SegmentPath = "basic-info"
    basicInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicInfo.EntityData.Children = make(map[string]types.YChild)
    basicInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    basicInfo.EntityData.Leafs["description"] = types.YLeaf{"Description", basicInfo.Description}
    basicInfo.EntityData.Leafs["vendor-type"] = types.YLeaf{"VendorType", basicInfo.VendorType}
    basicInfo.EntityData.Leafs["name"] = types.YLeaf{"Name", basicInfo.Name}
    basicInfo.EntityData.Leafs["hardware-revision"] = types.YLeaf{"HardwareRevision", basicInfo.HardwareRevision}
    basicInfo.EntityData.Leafs["firmware-revision"] = types.YLeaf{"FirmwareRevision", basicInfo.FirmwareRevision}
    basicInfo.EntityData.Leafs["software-revision"] = types.YLeaf{"SoftwareRevision", basicInfo.SoftwareRevision}
    basicInfo.EntityData.Leafs["chip-hardware-revision"] = types.YLeaf{"ChipHardwareRevision", basicInfo.ChipHardwareRevision}
    basicInfo.EntityData.Leafs["serial-number"] = types.YLeaf{"SerialNumber", basicInfo.SerialNumber}
    basicInfo.EntityData.Leafs["manufacturer-name"] = types.YLeaf{"ManufacturerName", basicInfo.ManufacturerName}
    basicInfo.EntityData.Leafs["model-name"] = types.YLeaf{"ModelName", basicInfo.ModelName}
    basicInfo.EntityData.Leafs["asset-id-str"] = types.YLeaf{"AssetIdStr", basicInfo.AssetIdStr}
    basicInfo.EntityData.Leafs["asset-identification"] = types.YLeaf{"AssetIdentification", basicInfo.AssetIdentification}
    basicInfo.EntityData.Leafs["is-field-replaceable-unit"] = types.YLeaf{"IsFieldReplaceableUnit", basicInfo.IsFieldReplaceableUnit}
    basicInfo.EntityData.Leafs["manufacturer-asset-tags"] = types.YLeaf{"ManufacturerAssetTags", basicInfo.ManufacturerAssetTags}
    basicInfo.EntityData.Leafs["composite-class-code"] = types.YLeaf{"CompositeClassCode", basicInfo.CompositeClassCode}
    basicInfo.EntityData.Leafs["memory-size"] = types.YLeaf{"MemorySize", basicInfo.MemorySize}
    basicInfo.EntityData.Leafs["environmental-monitor-path"] = types.YLeaf{"EnvironmentalMonitorPath", basicInfo.EnvironmentalMonitorPath}
    basicInfo.EntityData.Leafs["alias"] = types.YLeaf{"Alias", basicInfo.Alias}
    basicInfo.EntityData.Leafs["group-flag"] = types.YLeaf{"GroupFlag", basicInfo.GroupFlag}
    basicInfo.EntityData.Leafs["new-deviation-number"] = types.YLeaf{"NewDeviationNumber", basicInfo.NewDeviationNumber}
    basicInfo.EntityData.Leafs["physical-layer-interface-module-type"] = types.YLeaf{"PhysicalLayerInterfaceModuleType", basicInfo.PhysicalLayerInterfaceModuleType}
    basicInfo.EntityData.Leafs["unrecognized-fru"] = types.YLeaf{"UnrecognizedFru", basicInfo.UnrecognizedFru}
    basicInfo.EntityData.Leafs["redundancystate"] = types.YLeaf{"Redundancystate", basicInfo.Redundancystate}
    basicInfo.EntityData.Leafs["ceport"] = types.YLeaf{"Ceport", basicInfo.Ceport}
    basicInfo.EntityData.Leafs["xr-scoped"] = types.YLeaf{"XrScoped", basicInfo.XrScoped}
    basicInfo.EntityData.Leafs["unique-id"] = types.YLeaf{"UniqueId", basicInfo.UniqueId}
    return &(basicInfo.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_FruInfo
// Field Replaceable Unit (FRU) inventory
// information
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_FruInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // card admin state: shutdown or not. The type is interface{} with range:
    // -2147483648..2147483647.
    CardAdministrativeState interface{}

    // power admin state: up or down. The type is interface{} with range:
    // -2147483648..2147483647.
    PowerAdministrativeState interface{}

    // card operation state. The type is interface{} with range:
    // -2147483648..2147483647.
    CardOperationalState interface{}

    // card is monitored by a manager or left unmonitored. The type is interface{}
    // with range: -2147483648..2147483647.
    CardMonitorState interface{}

    // card reset reason enum. The type is CardResetReason.
    CardResetReason interface{}

    // power current: not implemented. The type is interface{} with range:
    // -2147483648..2147483647.
    PowerCurrentMeasurement interface{}

    // Power operation state. The type is interface{} with range:
    // -2147483648..2147483647.
    PowerOperationalState interface{}

    // last card oper change state.
    LastOperationalStateChange Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_FruInfo_LastOperationalStateChange

    // card up time.
    CardUpTime Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_FruInfo_CardUpTime
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_FruInfo) GetEntityData() *types.CommonEntityData {
    fruInfo.EntityData.YFilter = fruInfo.YFilter
    fruInfo.EntityData.YangName = "fru-info"
    fruInfo.EntityData.BundleName = "cisco_ios_xr"
    fruInfo.EntityData.ParentYangName = "basic-attributes"
    fruInfo.EntityData.SegmentPath = "fru-info"
    fruInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fruInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fruInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fruInfo.EntityData.Children = make(map[string]types.YChild)
    fruInfo.EntityData.Children["last-operational-state-change"] = types.YChild{"LastOperationalStateChange", &fruInfo.LastOperationalStateChange}
    fruInfo.EntityData.Children["card-up-time"] = types.YChild{"CardUpTime", &fruInfo.CardUpTime}
    fruInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    fruInfo.EntityData.Leafs["card-administrative-state"] = types.YLeaf{"CardAdministrativeState", fruInfo.CardAdministrativeState}
    fruInfo.EntityData.Leafs["power-administrative-state"] = types.YLeaf{"PowerAdministrativeState", fruInfo.PowerAdministrativeState}
    fruInfo.EntityData.Leafs["card-operational-state"] = types.YLeaf{"CardOperationalState", fruInfo.CardOperationalState}
    fruInfo.EntityData.Leafs["card-monitor-state"] = types.YLeaf{"CardMonitorState", fruInfo.CardMonitorState}
    fruInfo.EntityData.Leafs["card-reset-reason"] = types.YLeaf{"CardResetReason", fruInfo.CardResetReason}
    fruInfo.EntityData.Leafs["power-current-measurement"] = types.YLeaf{"PowerCurrentMeasurement", fruInfo.PowerCurrentMeasurement}
    fruInfo.EntityData.Leafs["power-operational-state"] = types.YLeaf{"PowerOperationalState", fruInfo.PowerOperationalState}
    return &(fruInfo.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_FruInfo_LastOperationalStateChange
// last card oper change state
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_FruInfo_LastOperationalStateChange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_FruInfo_LastOperationalStateChange) GetEntityData() *types.CommonEntityData {
    lastOperationalStateChange.EntityData.YFilter = lastOperationalStateChange.YFilter
    lastOperationalStateChange.EntityData.YangName = "last-operational-state-change"
    lastOperationalStateChange.EntityData.BundleName = "cisco_ios_xr"
    lastOperationalStateChange.EntityData.ParentYangName = "fru-info"
    lastOperationalStateChange.EntityData.SegmentPath = "last-operational-state-change"
    lastOperationalStateChange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastOperationalStateChange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastOperationalStateChange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastOperationalStateChange.EntityData.Children = make(map[string]types.YChild)
    lastOperationalStateChange.EntityData.Leafs = make(map[string]types.YLeaf)
    lastOperationalStateChange.EntityData.Leafs["time-in-seconds"] = types.YLeaf{"TimeInSeconds", lastOperationalStateChange.TimeInSeconds}
    lastOperationalStateChange.EntityData.Leafs["time-in-nano-seconds"] = types.YLeaf{"TimeInNanoSeconds", lastOperationalStateChange.TimeInNanoSeconds}
    return &(lastOperationalStateChange.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_FruInfo_CardUpTime
// card up time
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_FruInfo_CardUpTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_FruInfo_CardUpTime) GetEntityData() *types.CommonEntityData {
    cardUpTime.EntityData.YFilter = cardUpTime.YFilter
    cardUpTime.EntityData.YangName = "card-up-time"
    cardUpTime.EntityData.BundleName = "cisco_ios_xr"
    cardUpTime.EntityData.ParentYangName = "fru-info"
    cardUpTime.EntityData.SegmentPath = "card-up-time"
    cardUpTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cardUpTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cardUpTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cardUpTime.EntityData.Children = make(map[string]types.YChild)
    cardUpTime.EntityData.Leafs = make(map[string]types.YLeaf)
    cardUpTime.EntityData.Leafs["time-in-seconds"] = types.YLeaf{"TimeInSeconds", cardUpTime.TimeInSeconds}
    cardUpTime.EntityData.Leafs["time-in-nano-seconds"] = types.YLeaf{"TimeInNanoSeconds", cardUpTime.TimeInNanoSeconds}
    return &(cardUpTime.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots
// PortSlotTable contains all optics ports in a
// SPA/PLIM.
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PortSlot number in the SPA/PLIM. The type is slice of
    // Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot.
    PortSlot []Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot
}

func (portSlots *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots) GetEntityData() *types.CommonEntityData {
    portSlots.EntityData.YFilter = portSlots.YFilter
    portSlots.EntityData.YangName = "port-slots"
    portSlots.EntityData.BundleName = "cisco_ios_xr"
    portSlots.EntityData.ParentYangName = "module"
    portSlots.EntityData.SegmentPath = "port-slots"
    portSlots.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portSlots.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portSlots.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portSlots.EntityData.Children = make(map[string]types.YChild)
    portSlots.EntityData.Children["port-slot"] = types.YChild{"PortSlot", nil}
    for i := range portSlots.PortSlot {
        portSlots.EntityData.Children[types.GetSegmentPath(&portSlots.PortSlot[i])] = types.YChild{"PortSlot", &portSlots.PortSlot[i]}
    }
    portSlots.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(portSlots.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot
// PortSlot number in the SPA/PLIM
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. portslot number. The type is interface{} with
    // range: -2147483648..2147483647.
    Number interface{}

    // Port string.
    Port Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port

    // Attributes.
    BasicAttributes Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes
}

func (portSlot *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot) GetEntityData() *types.CommonEntityData {
    portSlot.EntityData.YFilter = portSlot.YFilter
    portSlot.EntityData.YangName = "port-slot"
    portSlot.EntityData.BundleName = "cisco_ios_xr"
    portSlot.EntityData.ParentYangName = "port-slots"
    portSlot.EntityData.SegmentPath = "port-slot" + "[number='" + fmt.Sprintf("%v", portSlot.Number) + "']"
    portSlot.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portSlot.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portSlot.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portSlot.EntityData.Children = make(map[string]types.YChild)
    portSlot.EntityData.Children["port"] = types.YChild{"Port", &portSlot.Port}
    portSlot.EntityData.Children["basic-attributes"] = types.YChild{"BasicAttributes", &portSlot.BasicAttributes}
    portSlot.EntityData.Leafs = make(map[string]types.YLeaf)
    portSlot.EntityData.Leafs["number"] = types.YLeaf{"Number", portSlot.Number}
    return &(portSlot.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port
// Port string
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Attributes.
    BasicAttributes Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes
}

func (port *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port) GetEntityData() *types.CommonEntityData {
    port.EntityData.YFilter = port.YFilter
    port.EntityData.YangName = "port"
    port.EntityData.BundleName = "cisco_ios_xr"
    port.EntityData.ParentYangName = "port-slot"
    port.EntityData.SegmentPath = "port"
    port.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    port.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    port.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    port.EntityData.Children = make(map[string]types.YChild)
    port.EntityData.Children["basic-attributes"] = types.YChild{"BasicAttributes", &port.BasicAttributes}
    port.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(port.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes
// Attributes
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Inventory information.
    BasicInfo Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_BasicInfo

    // Field Replaceable Unit (FRU) inventory information.
    FruInfo Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_FruInfo
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes) GetEntityData() *types.CommonEntityData {
    basicAttributes.EntityData.YFilter = basicAttributes.YFilter
    basicAttributes.EntityData.YangName = "basic-attributes"
    basicAttributes.EntityData.BundleName = "cisco_ios_xr"
    basicAttributes.EntityData.ParentYangName = "port"
    basicAttributes.EntityData.SegmentPath = "basic-attributes"
    basicAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicAttributes.EntityData.Children = make(map[string]types.YChild)
    basicAttributes.EntityData.Children["basic-info"] = types.YChild{"BasicInfo", &basicAttributes.BasicInfo}
    basicAttributes.EntityData.Children["fru-info"] = types.YChild{"FruInfo", &basicAttributes.FruInfo}
    basicAttributes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(basicAttributes.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_BasicInfo
// Inventory information
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_BasicInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // describes in user-readable terms       what the entity in question does.
    // The type is string with length: 0..255.
    Description interface{}

    // maps to the vendor OID string. The type is string with length: 0..255.
    VendorType interface{}

    // name string for the entity. The type is string with length: 0..255.
    Name interface{}

    // hw revision string. The type is string with length: 0..255.
    HardwareRevision interface{}

    // firmware revision string. The type is string with length: 0..255.
    FirmwareRevision interface{}

    // software revision string. The type is string with length: 0..255.
    SoftwareRevision interface{}

    // chip module hw revision string. The type is string with length: 0..255.
    ChipHardwareRevision interface{}

    // serial number. The type is string with length: 0..255.
    SerialNumber interface{}

    // manufacturer's name. The type is string with length: 0..255.
    ManufacturerName interface{}

    // model name. The type is string with length: 0..255.
    ModelName interface{}

    // asset Identification string. The type is string with length: 0..255.
    AssetIdStr interface{}

    // asset Identification. The type is interface{} with range:
    // -2147483648..2147483647.
    AssetIdentification interface{}

    // 1 if Field Replaceable Unit 0, if not. The type is bool.
    IsFieldReplaceableUnit interface{}

    // Manufacture Asset Tags. The type is interface{} with range:
    // -2147483648..2147483647.
    ManufacturerAssetTags interface{}

    // Major&minor class of the entity. The type is interface{} with range:
    // -2147483648..2147483647.
    CompositeClassCode interface{}

    // Size of memory associated with       the entity where applicable. The type
    // is interface{} with range: -2147483648..2147483647.
    MemorySize interface{}

    // sysdb name of sensor in the envmon EDM. The type is string with length:
    // 0..255.
    EnvironmentalMonitorPath interface{}

    // useful for storing an entity alias . The type is string with length:
    // 0..255.
    Alias interface{}

    // indicates if this entity is group       or not. The type is bool.
    GroupFlag interface{}

    // integer value for New Deviation Number 0x88. The type is interface{} with
    // range: -2147483648..2147483647.
    NewDeviationNumber interface{}

    // integer value for plim type if     applicable to this entity. The type is
    // interface{} with range: -2147483648..2147483647.
    PhysicalLayerInterfaceModuleType interface{}

    // 1 if UnrecognizedFRU and 0 for recognizedFRU. The type is bool.
    UnrecognizedFru interface{}

    // integer value for Redundancy State if     applicable to this entity. The
    // type is interface{} with range: -2147483648..2147483647.
    Redundancystate interface{}

    // 1 if ce port found, 0 if not. The type is bool.
    Ceport interface{}

    // 1 if xr scoped, 0 if not. The type is bool.
    XrScoped interface{}

    // Unique id for an entity. The type is interface{} with range:
    // -2147483648..2147483647.
    UniqueId interface{}
}

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_BasicInfo) GetEntityData() *types.CommonEntityData {
    basicInfo.EntityData.YFilter = basicInfo.YFilter
    basicInfo.EntityData.YangName = "basic-info"
    basicInfo.EntityData.BundleName = "cisco_ios_xr"
    basicInfo.EntityData.ParentYangName = "basic-attributes"
    basicInfo.EntityData.SegmentPath = "basic-info"
    basicInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicInfo.EntityData.Children = make(map[string]types.YChild)
    basicInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    basicInfo.EntityData.Leafs["description"] = types.YLeaf{"Description", basicInfo.Description}
    basicInfo.EntityData.Leafs["vendor-type"] = types.YLeaf{"VendorType", basicInfo.VendorType}
    basicInfo.EntityData.Leafs["name"] = types.YLeaf{"Name", basicInfo.Name}
    basicInfo.EntityData.Leafs["hardware-revision"] = types.YLeaf{"HardwareRevision", basicInfo.HardwareRevision}
    basicInfo.EntityData.Leafs["firmware-revision"] = types.YLeaf{"FirmwareRevision", basicInfo.FirmwareRevision}
    basicInfo.EntityData.Leafs["software-revision"] = types.YLeaf{"SoftwareRevision", basicInfo.SoftwareRevision}
    basicInfo.EntityData.Leafs["chip-hardware-revision"] = types.YLeaf{"ChipHardwareRevision", basicInfo.ChipHardwareRevision}
    basicInfo.EntityData.Leafs["serial-number"] = types.YLeaf{"SerialNumber", basicInfo.SerialNumber}
    basicInfo.EntityData.Leafs["manufacturer-name"] = types.YLeaf{"ManufacturerName", basicInfo.ManufacturerName}
    basicInfo.EntityData.Leafs["model-name"] = types.YLeaf{"ModelName", basicInfo.ModelName}
    basicInfo.EntityData.Leafs["asset-id-str"] = types.YLeaf{"AssetIdStr", basicInfo.AssetIdStr}
    basicInfo.EntityData.Leafs["asset-identification"] = types.YLeaf{"AssetIdentification", basicInfo.AssetIdentification}
    basicInfo.EntityData.Leafs["is-field-replaceable-unit"] = types.YLeaf{"IsFieldReplaceableUnit", basicInfo.IsFieldReplaceableUnit}
    basicInfo.EntityData.Leafs["manufacturer-asset-tags"] = types.YLeaf{"ManufacturerAssetTags", basicInfo.ManufacturerAssetTags}
    basicInfo.EntityData.Leafs["composite-class-code"] = types.YLeaf{"CompositeClassCode", basicInfo.CompositeClassCode}
    basicInfo.EntityData.Leafs["memory-size"] = types.YLeaf{"MemorySize", basicInfo.MemorySize}
    basicInfo.EntityData.Leafs["environmental-monitor-path"] = types.YLeaf{"EnvironmentalMonitorPath", basicInfo.EnvironmentalMonitorPath}
    basicInfo.EntityData.Leafs["alias"] = types.YLeaf{"Alias", basicInfo.Alias}
    basicInfo.EntityData.Leafs["group-flag"] = types.YLeaf{"GroupFlag", basicInfo.GroupFlag}
    basicInfo.EntityData.Leafs["new-deviation-number"] = types.YLeaf{"NewDeviationNumber", basicInfo.NewDeviationNumber}
    basicInfo.EntityData.Leafs["physical-layer-interface-module-type"] = types.YLeaf{"PhysicalLayerInterfaceModuleType", basicInfo.PhysicalLayerInterfaceModuleType}
    basicInfo.EntityData.Leafs["unrecognized-fru"] = types.YLeaf{"UnrecognizedFru", basicInfo.UnrecognizedFru}
    basicInfo.EntityData.Leafs["redundancystate"] = types.YLeaf{"Redundancystate", basicInfo.Redundancystate}
    basicInfo.EntityData.Leafs["ceport"] = types.YLeaf{"Ceport", basicInfo.Ceport}
    basicInfo.EntityData.Leafs["xr-scoped"] = types.YLeaf{"XrScoped", basicInfo.XrScoped}
    basicInfo.EntityData.Leafs["unique-id"] = types.YLeaf{"UniqueId", basicInfo.UniqueId}
    return &(basicInfo.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_FruInfo
// Field Replaceable Unit (FRU) inventory
// information
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_FruInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // card admin state: shutdown or not. The type is interface{} with range:
    // -2147483648..2147483647.
    CardAdministrativeState interface{}

    // power admin state: up or down. The type is interface{} with range:
    // -2147483648..2147483647.
    PowerAdministrativeState interface{}

    // card operation state. The type is interface{} with range:
    // -2147483648..2147483647.
    CardOperationalState interface{}

    // card is monitored by a manager or left unmonitored. The type is interface{}
    // with range: -2147483648..2147483647.
    CardMonitorState interface{}

    // card reset reason enum. The type is CardResetReason.
    CardResetReason interface{}

    // power current: not implemented. The type is interface{} with range:
    // -2147483648..2147483647.
    PowerCurrentMeasurement interface{}

    // Power operation state. The type is interface{} with range:
    // -2147483648..2147483647.
    PowerOperationalState interface{}

    // last card oper change state.
    LastOperationalStateChange Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_LastOperationalStateChange

    // card up time.
    CardUpTime Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_CardUpTime
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_FruInfo) GetEntityData() *types.CommonEntityData {
    fruInfo.EntityData.YFilter = fruInfo.YFilter
    fruInfo.EntityData.YangName = "fru-info"
    fruInfo.EntityData.BundleName = "cisco_ios_xr"
    fruInfo.EntityData.ParentYangName = "basic-attributes"
    fruInfo.EntityData.SegmentPath = "fru-info"
    fruInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fruInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fruInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fruInfo.EntityData.Children = make(map[string]types.YChild)
    fruInfo.EntityData.Children["last-operational-state-change"] = types.YChild{"LastOperationalStateChange", &fruInfo.LastOperationalStateChange}
    fruInfo.EntityData.Children["card-up-time"] = types.YChild{"CardUpTime", &fruInfo.CardUpTime}
    fruInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    fruInfo.EntityData.Leafs["card-administrative-state"] = types.YLeaf{"CardAdministrativeState", fruInfo.CardAdministrativeState}
    fruInfo.EntityData.Leafs["power-administrative-state"] = types.YLeaf{"PowerAdministrativeState", fruInfo.PowerAdministrativeState}
    fruInfo.EntityData.Leafs["card-operational-state"] = types.YLeaf{"CardOperationalState", fruInfo.CardOperationalState}
    fruInfo.EntityData.Leafs["card-monitor-state"] = types.YLeaf{"CardMonitorState", fruInfo.CardMonitorState}
    fruInfo.EntityData.Leafs["card-reset-reason"] = types.YLeaf{"CardResetReason", fruInfo.CardResetReason}
    fruInfo.EntityData.Leafs["power-current-measurement"] = types.YLeaf{"PowerCurrentMeasurement", fruInfo.PowerCurrentMeasurement}
    fruInfo.EntityData.Leafs["power-operational-state"] = types.YLeaf{"PowerOperationalState", fruInfo.PowerOperationalState}
    return &(fruInfo.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_LastOperationalStateChange
// last card oper change state
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_LastOperationalStateChange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_LastOperationalStateChange) GetEntityData() *types.CommonEntityData {
    lastOperationalStateChange.EntityData.YFilter = lastOperationalStateChange.YFilter
    lastOperationalStateChange.EntityData.YangName = "last-operational-state-change"
    lastOperationalStateChange.EntityData.BundleName = "cisco_ios_xr"
    lastOperationalStateChange.EntityData.ParentYangName = "fru-info"
    lastOperationalStateChange.EntityData.SegmentPath = "last-operational-state-change"
    lastOperationalStateChange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastOperationalStateChange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastOperationalStateChange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastOperationalStateChange.EntityData.Children = make(map[string]types.YChild)
    lastOperationalStateChange.EntityData.Leafs = make(map[string]types.YLeaf)
    lastOperationalStateChange.EntityData.Leafs["time-in-seconds"] = types.YLeaf{"TimeInSeconds", lastOperationalStateChange.TimeInSeconds}
    lastOperationalStateChange.EntityData.Leafs["time-in-nano-seconds"] = types.YLeaf{"TimeInNanoSeconds", lastOperationalStateChange.TimeInNanoSeconds}
    return &(lastOperationalStateChange.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_CardUpTime
// card up time
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_CardUpTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_CardUpTime) GetEntityData() *types.CommonEntityData {
    cardUpTime.EntityData.YFilter = cardUpTime.YFilter
    cardUpTime.EntityData.YangName = "card-up-time"
    cardUpTime.EntityData.BundleName = "cisco_ios_xr"
    cardUpTime.EntityData.ParentYangName = "fru-info"
    cardUpTime.EntityData.SegmentPath = "card-up-time"
    cardUpTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cardUpTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cardUpTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cardUpTime.EntityData.Children = make(map[string]types.YChild)
    cardUpTime.EntityData.Leafs = make(map[string]types.YLeaf)
    cardUpTime.EntityData.Leafs["time-in-seconds"] = types.YLeaf{"TimeInSeconds", cardUpTime.TimeInSeconds}
    cardUpTime.EntityData.Leafs["time-in-nano-seconds"] = types.YLeaf{"TimeInNanoSeconds", cardUpTime.TimeInNanoSeconds}
    return &(cardUpTime.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes
// Attributes
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Inventory information.
    BasicInfo Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_BasicInfo

    // Field Replaceable Unit (FRU) inventory information.
    FruInfo Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_FruInfo
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes) GetEntityData() *types.CommonEntityData {
    basicAttributes.EntityData.YFilter = basicAttributes.YFilter
    basicAttributes.EntityData.YangName = "basic-attributes"
    basicAttributes.EntityData.BundleName = "cisco_ios_xr"
    basicAttributes.EntityData.ParentYangName = "port-slot"
    basicAttributes.EntityData.SegmentPath = "basic-attributes"
    basicAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicAttributes.EntityData.Children = make(map[string]types.YChild)
    basicAttributes.EntityData.Children["basic-info"] = types.YChild{"BasicInfo", &basicAttributes.BasicInfo}
    basicAttributes.EntityData.Children["fru-info"] = types.YChild{"FruInfo", &basicAttributes.FruInfo}
    basicAttributes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(basicAttributes.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_BasicInfo
// Inventory information
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_BasicInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // describes in user-readable terms       what the entity in question does.
    // The type is string with length: 0..255.
    Description interface{}

    // maps to the vendor OID string. The type is string with length: 0..255.
    VendorType interface{}

    // name string for the entity. The type is string with length: 0..255.
    Name interface{}

    // hw revision string. The type is string with length: 0..255.
    HardwareRevision interface{}

    // firmware revision string. The type is string with length: 0..255.
    FirmwareRevision interface{}

    // software revision string. The type is string with length: 0..255.
    SoftwareRevision interface{}

    // chip module hw revision string. The type is string with length: 0..255.
    ChipHardwareRevision interface{}

    // serial number. The type is string with length: 0..255.
    SerialNumber interface{}

    // manufacturer's name. The type is string with length: 0..255.
    ManufacturerName interface{}

    // model name. The type is string with length: 0..255.
    ModelName interface{}

    // asset Identification string. The type is string with length: 0..255.
    AssetIdStr interface{}

    // asset Identification. The type is interface{} with range:
    // -2147483648..2147483647.
    AssetIdentification interface{}

    // 1 if Field Replaceable Unit 0, if not. The type is bool.
    IsFieldReplaceableUnit interface{}

    // Manufacture Asset Tags. The type is interface{} with range:
    // -2147483648..2147483647.
    ManufacturerAssetTags interface{}

    // Major&minor class of the entity. The type is interface{} with range:
    // -2147483648..2147483647.
    CompositeClassCode interface{}

    // Size of memory associated with       the entity where applicable. The type
    // is interface{} with range: -2147483648..2147483647.
    MemorySize interface{}

    // sysdb name of sensor in the envmon EDM. The type is string with length:
    // 0..255.
    EnvironmentalMonitorPath interface{}

    // useful for storing an entity alias . The type is string with length:
    // 0..255.
    Alias interface{}

    // indicates if this entity is group       or not. The type is bool.
    GroupFlag interface{}

    // integer value for New Deviation Number 0x88. The type is interface{} with
    // range: -2147483648..2147483647.
    NewDeviationNumber interface{}

    // integer value for plim type if     applicable to this entity. The type is
    // interface{} with range: -2147483648..2147483647.
    PhysicalLayerInterfaceModuleType interface{}

    // 1 if UnrecognizedFRU and 0 for recognizedFRU. The type is bool.
    UnrecognizedFru interface{}

    // integer value for Redundancy State if     applicable to this entity. The
    // type is interface{} with range: -2147483648..2147483647.
    Redundancystate interface{}

    // 1 if ce port found, 0 if not. The type is bool.
    Ceport interface{}

    // 1 if xr scoped, 0 if not. The type is bool.
    XrScoped interface{}

    // Unique id for an entity. The type is interface{} with range:
    // -2147483648..2147483647.
    UniqueId interface{}
}

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_BasicInfo) GetEntityData() *types.CommonEntityData {
    basicInfo.EntityData.YFilter = basicInfo.YFilter
    basicInfo.EntityData.YangName = "basic-info"
    basicInfo.EntityData.BundleName = "cisco_ios_xr"
    basicInfo.EntityData.ParentYangName = "basic-attributes"
    basicInfo.EntityData.SegmentPath = "basic-info"
    basicInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicInfo.EntityData.Children = make(map[string]types.YChild)
    basicInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    basicInfo.EntityData.Leafs["description"] = types.YLeaf{"Description", basicInfo.Description}
    basicInfo.EntityData.Leafs["vendor-type"] = types.YLeaf{"VendorType", basicInfo.VendorType}
    basicInfo.EntityData.Leafs["name"] = types.YLeaf{"Name", basicInfo.Name}
    basicInfo.EntityData.Leafs["hardware-revision"] = types.YLeaf{"HardwareRevision", basicInfo.HardwareRevision}
    basicInfo.EntityData.Leafs["firmware-revision"] = types.YLeaf{"FirmwareRevision", basicInfo.FirmwareRevision}
    basicInfo.EntityData.Leafs["software-revision"] = types.YLeaf{"SoftwareRevision", basicInfo.SoftwareRevision}
    basicInfo.EntityData.Leafs["chip-hardware-revision"] = types.YLeaf{"ChipHardwareRevision", basicInfo.ChipHardwareRevision}
    basicInfo.EntityData.Leafs["serial-number"] = types.YLeaf{"SerialNumber", basicInfo.SerialNumber}
    basicInfo.EntityData.Leafs["manufacturer-name"] = types.YLeaf{"ManufacturerName", basicInfo.ManufacturerName}
    basicInfo.EntityData.Leafs["model-name"] = types.YLeaf{"ModelName", basicInfo.ModelName}
    basicInfo.EntityData.Leafs["asset-id-str"] = types.YLeaf{"AssetIdStr", basicInfo.AssetIdStr}
    basicInfo.EntityData.Leafs["asset-identification"] = types.YLeaf{"AssetIdentification", basicInfo.AssetIdentification}
    basicInfo.EntityData.Leafs["is-field-replaceable-unit"] = types.YLeaf{"IsFieldReplaceableUnit", basicInfo.IsFieldReplaceableUnit}
    basicInfo.EntityData.Leafs["manufacturer-asset-tags"] = types.YLeaf{"ManufacturerAssetTags", basicInfo.ManufacturerAssetTags}
    basicInfo.EntityData.Leafs["composite-class-code"] = types.YLeaf{"CompositeClassCode", basicInfo.CompositeClassCode}
    basicInfo.EntityData.Leafs["memory-size"] = types.YLeaf{"MemorySize", basicInfo.MemorySize}
    basicInfo.EntityData.Leafs["environmental-monitor-path"] = types.YLeaf{"EnvironmentalMonitorPath", basicInfo.EnvironmentalMonitorPath}
    basicInfo.EntityData.Leafs["alias"] = types.YLeaf{"Alias", basicInfo.Alias}
    basicInfo.EntityData.Leafs["group-flag"] = types.YLeaf{"GroupFlag", basicInfo.GroupFlag}
    basicInfo.EntityData.Leafs["new-deviation-number"] = types.YLeaf{"NewDeviationNumber", basicInfo.NewDeviationNumber}
    basicInfo.EntityData.Leafs["physical-layer-interface-module-type"] = types.YLeaf{"PhysicalLayerInterfaceModuleType", basicInfo.PhysicalLayerInterfaceModuleType}
    basicInfo.EntityData.Leafs["unrecognized-fru"] = types.YLeaf{"UnrecognizedFru", basicInfo.UnrecognizedFru}
    basicInfo.EntityData.Leafs["redundancystate"] = types.YLeaf{"Redundancystate", basicInfo.Redundancystate}
    basicInfo.EntityData.Leafs["ceport"] = types.YLeaf{"Ceport", basicInfo.Ceport}
    basicInfo.EntityData.Leafs["xr-scoped"] = types.YLeaf{"XrScoped", basicInfo.XrScoped}
    basicInfo.EntityData.Leafs["unique-id"] = types.YLeaf{"UniqueId", basicInfo.UniqueId}
    return &(basicInfo.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_FruInfo
// Field Replaceable Unit (FRU) inventory
// information
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_FruInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // card admin state: shutdown or not. The type is interface{} with range:
    // -2147483648..2147483647.
    CardAdministrativeState interface{}

    // power admin state: up or down. The type is interface{} with range:
    // -2147483648..2147483647.
    PowerAdministrativeState interface{}

    // card operation state. The type is interface{} with range:
    // -2147483648..2147483647.
    CardOperationalState interface{}

    // card is monitored by a manager or left unmonitored. The type is interface{}
    // with range: -2147483648..2147483647.
    CardMonitorState interface{}

    // card reset reason enum. The type is CardResetReason.
    CardResetReason interface{}

    // power current: not implemented. The type is interface{} with range:
    // -2147483648..2147483647.
    PowerCurrentMeasurement interface{}

    // Power operation state. The type is interface{} with range:
    // -2147483648..2147483647.
    PowerOperationalState interface{}

    // last card oper change state.
    LastOperationalStateChange Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_FruInfo_LastOperationalStateChange

    // card up time.
    CardUpTime Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_FruInfo_CardUpTime
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_FruInfo) GetEntityData() *types.CommonEntityData {
    fruInfo.EntityData.YFilter = fruInfo.YFilter
    fruInfo.EntityData.YangName = "fru-info"
    fruInfo.EntityData.BundleName = "cisco_ios_xr"
    fruInfo.EntityData.ParentYangName = "basic-attributes"
    fruInfo.EntityData.SegmentPath = "fru-info"
    fruInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fruInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fruInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fruInfo.EntityData.Children = make(map[string]types.YChild)
    fruInfo.EntityData.Children["last-operational-state-change"] = types.YChild{"LastOperationalStateChange", &fruInfo.LastOperationalStateChange}
    fruInfo.EntityData.Children["card-up-time"] = types.YChild{"CardUpTime", &fruInfo.CardUpTime}
    fruInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    fruInfo.EntityData.Leafs["card-administrative-state"] = types.YLeaf{"CardAdministrativeState", fruInfo.CardAdministrativeState}
    fruInfo.EntityData.Leafs["power-administrative-state"] = types.YLeaf{"PowerAdministrativeState", fruInfo.PowerAdministrativeState}
    fruInfo.EntityData.Leafs["card-operational-state"] = types.YLeaf{"CardOperationalState", fruInfo.CardOperationalState}
    fruInfo.EntityData.Leafs["card-monitor-state"] = types.YLeaf{"CardMonitorState", fruInfo.CardMonitorState}
    fruInfo.EntityData.Leafs["card-reset-reason"] = types.YLeaf{"CardResetReason", fruInfo.CardResetReason}
    fruInfo.EntityData.Leafs["power-current-measurement"] = types.YLeaf{"PowerCurrentMeasurement", fruInfo.PowerCurrentMeasurement}
    fruInfo.EntityData.Leafs["power-operational-state"] = types.YLeaf{"PowerOperationalState", fruInfo.PowerOperationalState}
    return &(fruInfo.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_FruInfo_LastOperationalStateChange
// last card oper change state
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_FruInfo_LastOperationalStateChange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_FruInfo_LastOperationalStateChange) GetEntityData() *types.CommonEntityData {
    lastOperationalStateChange.EntityData.YFilter = lastOperationalStateChange.YFilter
    lastOperationalStateChange.EntityData.YangName = "last-operational-state-change"
    lastOperationalStateChange.EntityData.BundleName = "cisco_ios_xr"
    lastOperationalStateChange.EntityData.ParentYangName = "fru-info"
    lastOperationalStateChange.EntityData.SegmentPath = "last-operational-state-change"
    lastOperationalStateChange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastOperationalStateChange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastOperationalStateChange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastOperationalStateChange.EntityData.Children = make(map[string]types.YChild)
    lastOperationalStateChange.EntityData.Leafs = make(map[string]types.YLeaf)
    lastOperationalStateChange.EntityData.Leafs["time-in-seconds"] = types.YLeaf{"TimeInSeconds", lastOperationalStateChange.TimeInSeconds}
    lastOperationalStateChange.EntityData.Leafs["time-in-nano-seconds"] = types.YLeaf{"TimeInNanoSeconds", lastOperationalStateChange.TimeInNanoSeconds}
    return &(lastOperationalStateChange.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_FruInfo_CardUpTime
// card up time
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_FruInfo_CardUpTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_FruInfo_CardUpTime) GetEntityData() *types.CommonEntityData {
    cardUpTime.EntityData.YFilter = cardUpTime.YFilter
    cardUpTime.EntityData.YangName = "card-up-time"
    cardUpTime.EntityData.BundleName = "cisco_ios_xr"
    cardUpTime.EntityData.ParentYangName = "fru-info"
    cardUpTime.EntityData.SegmentPath = "card-up-time"
    cardUpTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cardUpTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cardUpTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cardUpTime.EntityData.Children = make(map[string]types.YChild)
    cardUpTime.EntityData.Leafs = make(map[string]types.YLeaf)
    cardUpTime.EntityData.Leafs["time-in-seconds"] = types.YLeaf{"TimeInSeconds", cardUpTime.TimeInSeconds}
    cardUpTime.EntityData.Leafs["time-in-nano-seconds"] = types.YLeaf{"TimeInNanoSeconds", cardUpTime.TimeInNanoSeconds}
    return &(cardUpTime.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes
// Attributes
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Inventory information.
    BasicInfo Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_BasicInfo

    // Field Replaceable Unit (FRU) inventory information.
    FruInfo Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_FruInfo
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes) GetEntityData() *types.CommonEntityData {
    basicAttributes.EntityData.YFilter = basicAttributes.YFilter
    basicAttributes.EntityData.YangName = "basic-attributes"
    basicAttributes.EntityData.BundleName = "cisco_ios_xr"
    basicAttributes.EntityData.ParentYangName = "module"
    basicAttributes.EntityData.SegmentPath = "basic-attributes"
    basicAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicAttributes.EntityData.Children = make(map[string]types.YChild)
    basicAttributes.EntityData.Children["basic-info"] = types.YChild{"BasicInfo", &basicAttributes.BasicInfo}
    basicAttributes.EntityData.Children["fru-info"] = types.YChild{"FruInfo", &basicAttributes.FruInfo}
    basicAttributes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(basicAttributes.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_BasicInfo
// Inventory information
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_BasicInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // describes in user-readable terms       what the entity in question does.
    // The type is string with length: 0..255.
    Description interface{}

    // maps to the vendor OID string. The type is string with length: 0..255.
    VendorType interface{}

    // name string for the entity. The type is string with length: 0..255.
    Name interface{}

    // hw revision string. The type is string with length: 0..255.
    HardwareRevision interface{}

    // firmware revision string. The type is string with length: 0..255.
    FirmwareRevision interface{}

    // software revision string. The type is string with length: 0..255.
    SoftwareRevision interface{}

    // chip module hw revision string. The type is string with length: 0..255.
    ChipHardwareRevision interface{}

    // serial number. The type is string with length: 0..255.
    SerialNumber interface{}

    // manufacturer's name. The type is string with length: 0..255.
    ManufacturerName interface{}

    // model name. The type is string with length: 0..255.
    ModelName interface{}

    // asset Identification string. The type is string with length: 0..255.
    AssetIdStr interface{}

    // asset Identification. The type is interface{} with range:
    // -2147483648..2147483647.
    AssetIdentification interface{}

    // 1 if Field Replaceable Unit 0, if not. The type is bool.
    IsFieldReplaceableUnit interface{}

    // Manufacture Asset Tags. The type is interface{} with range:
    // -2147483648..2147483647.
    ManufacturerAssetTags interface{}

    // Major&minor class of the entity. The type is interface{} with range:
    // -2147483648..2147483647.
    CompositeClassCode interface{}

    // Size of memory associated with       the entity where applicable. The type
    // is interface{} with range: -2147483648..2147483647.
    MemorySize interface{}

    // sysdb name of sensor in the envmon EDM. The type is string with length:
    // 0..255.
    EnvironmentalMonitorPath interface{}

    // useful for storing an entity alias . The type is string with length:
    // 0..255.
    Alias interface{}

    // indicates if this entity is group       or not. The type is bool.
    GroupFlag interface{}

    // integer value for New Deviation Number 0x88. The type is interface{} with
    // range: -2147483648..2147483647.
    NewDeviationNumber interface{}

    // integer value for plim type if     applicable to this entity. The type is
    // interface{} with range: -2147483648..2147483647.
    PhysicalLayerInterfaceModuleType interface{}

    // 1 if UnrecognizedFRU and 0 for recognizedFRU. The type is bool.
    UnrecognizedFru interface{}

    // integer value for Redundancy State if     applicable to this entity. The
    // type is interface{} with range: -2147483648..2147483647.
    Redundancystate interface{}

    // 1 if ce port found, 0 if not. The type is bool.
    Ceport interface{}

    // 1 if xr scoped, 0 if not. The type is bool.
    XrScoped interface{}

    // Unique id for an entity. The type is interface{} with range:
    // -2147483648..2147483647.
    UniqueId interface{}
}

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_BasicInfo) GetEntityData() *types.CommonEntityData {
    basicInfo.EntityData.YFilter = basicInfo.YFilter
    basicInfo.EntityData.YangName = "basic-info"
    basicInfo.EntityData.BundleName = "cisco_ios_xr"
    basicInfo.EntityData.ParentYangName = "basic-attributes"
    basicInfo.EntityData.SegmentPath = "basic-info"
    basicInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicInfo.EntityData.Children = make(map[string]types.YChild)
    basicInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    basicInfo.EntityData.Leafs["description"] = types.YLeaf{"Description", basicInfo.Description}
    basicInfo.EntityData.Leafs["vendor-type"] = types.YLeaf{"VendorType", basicInfo.VendorType}
    basicInfo.EntityData.Leafs["name"] = types.YLeaf{"Name", basicInfo.Name}
    basicInfo.EntityData.Leafs["hardware-revision"] = types.YLeaf{"HardwareRevision", basicInfo.HardwareRevision}
    basicInfo.EntityData.Leafs["firmware-revision"] = types.YLeaf{"FirmwareRevision", basicInfo.FirmwareRevision}
    basicInfo.EntityData.Leafs["software-revision"] = types.YLeaf{"SoftwareRevision", basicInfo.SoftwareRevision}
    basicInfo.EntityData.Leafs["chip-hardware-revision"] = types.YLeaf{"ChipHardwareRevision", basicInfo.ChipHardwareRevision}
    basicInfo.EntityData.Leafs["serial-number"] = types.YLeaf{"SerialNumber", basicInfo.SerialNumber}
    basicInfo.EntityData.Leafs["manufacturer-name"] = types.YLeaf{"ManufacturerName", basicInfo.ManufacturerName}
    basicInfo.EntityData.Leafs["model-name"] = types.YLeaf{"ModelName", basicInfo.ModelName}
    basicInfo.EntityData.Leafs["asset-id-str"] = types.YLeaf{"AssetIdStr", basicInfo.AssetIdStr}
    basicInfo.EntityData.Leafs["asset-identification"] = types.YLeaf{"AssetIdentification", basicInfo.AssetIdentification}
    basicInfo.EntityData.Leafs["is-field-replaceable-unit"] = types.YLeaf{"IsFieldReplaceableUnit", basicInfo.IsFieldReplaceableUnit}
    basicInfo.EntityData.Leafs["manufacturer-asset-tags"] = types.YLeaf{"ManufacturerAssetTags", basicInfo.ManufacturerAssetTags}
    basicInfo.EntityData.Leafs["composite-class-code"] = types.YLeaf{"CompositeClassCode", basicInfo.CompositeClassCode}
    basicInfo.EntityData.Leafs["memory-size"] = types.YLeaf{"MemorySize", basicInfo.MemorySize}
    basicInfo.EntityData.Leafs["environmental-monitor-path"] = types.YLeaf{"EnvironmentalMonitorPath", basicInfo.EnvironmentalMonitorPath}
    basicInfo.EntityData.Leafs["alias"] = types.YLeaf{"Alias", basicInfo.Alias}
    basicInfo.EntityData.Leafs["group-flag"] = types.YLeaf{"GroupFlag", basicInfo.GroupFlag}
    basicInfo.EntityData.Leafs["new-deviation-number"] = types.YLeaf{"NewDeviationNumber", basicInfo.NewDeviationNumber}
    basicInfo.EntityData.Leafs["physical-layer-interface-module-type"] = types.YLeaf{"PhysicalLayerInterfaceModuleType", basicInfo.PhysicalLayerInterfaceModuleType}
    basicInfo.EntityData.Leafs["unrecognized-fru"] = types.YLeaf{"UnrecognizedFru", basicInfo.UnrecognizedFru}
    basicInfo.EntityData.Leafs["redundancystate"] = types.YLeaf{"Redundancystate", basicInfo.Redundancystate}
    basicInfo.EntityData.Leafs["ceport"] = types.YLeaf{"Ceport", basicInfo.Ceport}
    basicInfo.EntityData.Leafs["xr-scoped"] = types.YLeaf{"XrScoped", basicInfo.XrScoped}
    basicInfo.EntityData.Leafs["unique-id"] = types.YLeaf{"UniqueId", basicInfo.UniqueId}
    return &(basicInfo.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_FruInfo
// Field Replaceable Unit (FRU) inventory
// information
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_FruInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // card admin state: shutdown or not. The type is interface{} with range:
    // -2147483648..2147483647.
    CardAdministrativeState interface{}

    // power admin state: up or down. The type is interface{} with range:
    // -2147483648..2147483647.
    PowerAdministrativeState interface{}

    // card operation state. The type is interface{} with range:
    // -2147483648..2147483647.
    CardOperationalState interface{}

    // card is monitored by a manager or left unmonitored. The type is interface{}
    // with range: -2147483648..2147483647.
    CardMonitorState interface{}

    // card reset reason enum. The type is CardResetReason.
    CardResetReason interface{}

    // power current: not implemented. The type is interface{} with range:
    // -2147483648..2147483647.
    PowerCurrentMeasurement interface{}

    // Power operation state. The type is interface{} with range:
    // -2147483648..2147483647.
    PowerOperationalState interface{}

    // last card oper change state.
    LastOperationalStateChange Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_FruInfo_LastOperationalStateChange

    // card up time.
    CardUpTime Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_FruInfo_CardUpTime
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_FruInfo) GetEntityData() *types.CommonEntityData {
    fruInfo.EntityData.YFilter = fruInfo.YFilter
    fruInfo.EntityData.YangName = "fru-info"
    fruInfo.EntityData.BundleName = "cisco_ios_xr"
    fruInfo.EntityData.ParentYangName = "basic-attributes"
    fruInfo.EntityData.SegmentPath = "fru-info"
    fruInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fruInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fruInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fruInfo.EntityData.Children = make(map[string]types.YChild)
    fruInfo.EntityData.Children["last-operational-state-change"] = types.YChild{"LastOperationalStateChange", &fruInfo.LastOperationalStateChange}
    fruInfo.EntityData.Children["card-up-time"] = types.YChild{"CardUpTime", &fruInfo.CardUpTime}
    fruInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    fruInfo.EntityData.Leafs["card-administrative-state"] = types.YLeaf{"CardAdministrativeState", fruInfo.CardAdministrativeState}
    fruInfo.EntityData.Leafs["power-administrative-state"] = types.YLeaf{"PowerAdministrativeState", fruInfo.PowerAdministrativeState}
    fruInfo.EntityData.Leafs["card-operational-state"] = types.YLeaf{"CardOperationalState", fruInfo.CardOperationalState}
    fruInfo.EntityData.Leafs["card-monitor-state"] = types.YLeaf{"CardMonitorState", fruInfo.CardMonitorState}
    fruInfo.EntityData.Leafs["card-reset-reason"] = types.YLeaf{"CardResetReason", fruInfo.CardResetReason}
    fruInfo.EntityData.Leafs["power-current-measurement"] = types.YLeaf{"PowerCurrentMeasurement", fruInfo.PowerCurrentMeasurement}
    fruInfo.EntityData.Leafs["power-operational-state"] = types.YLeaf{"PowerOperationalState", fruInfo.PowerOperationalState}
    return &(fruInfo.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_FruInfo_LastOperationalStateChange
// last card oper change state
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_FruInfo_LastOperationalStateChange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_FruInfo_LastOperationalStateChange) GetEntityData() *types.CommonEntityData {
    lastOperationalStateChange.EntityData.YFilter = lastOperationalStateChange.YFilter
    lastOperationalStateChange.EntityData.YangName = "last-operational-state-change"
    lastOperationalStateChange.EntityData.BundleName = "cisco_ios_xr"
    lastOperationalStateChange.EntityData.ParentYangName = "fru-info"
    lastOperationalStateChange.EntityData.SegmentPath = "last-operational-state-change"
    lastOperationalStateChange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastOperationalStateChange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastOperationalStateChange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastOperationalStateChange.EntityData.Children = make(map[string]types.YChild)
    lastOperationalStateChange.EntityData.Leafs = make(map[string]types.YLeaf)
    lastOperationalStateChange.EntityData.Leafs["time-in-seconds"] = types.YLeaf{"TimeInSeconds", lastOperationalStateChange.TimeInSeconds}
    lastOperationalStateChange.EntityData.Leafs["time-in-nano-seconds"] = types.YLeaf{"TimeInNanoSeconds", lastOperationalStateChange.TimeInNanoSeconds}
    return &(lastOperationalStateChange.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_FruInfo_CardUpTime
// card up time
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_FruInfo_CardUpTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_FruInfo_CardUpTime) GetEntityData() *types.CommonEntityData {
    cardUpTime.EntityData.YFilter = cardUpTime.YFilter
    cardUpTime.EntityData.YangName = "card-up-time"
    cardUpTime.EntityData.BundleName = "cisco_ios_xr"
    cardUpTime.EntityData.ParentYangName = "fru-info"
    cardUpTime.EntityData.SegmentPath = "card-up-time"
    cardUpTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cardUpTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cardUpTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cardUpTime.EntityData.Children = make(map[string]types.YChild)
    cardUpTime.EntityData.Leafs = make(map[string]types.YLeaf)
    cardUpTime.EntityData.Leafs["time-in-seconds"] = types.YLeaf{"TimeInSeconds", cardUpTime.TimeInSeconds}
    cardUpTime.EntityData.Leafs["time-in-nano-seconds"] = types.YLeaf{"TimeInNanoSeconds", cardUpTime.TimeInNanoSeconds}
    return &(cardUpTime.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes
// Attributes
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Inventory information.
    BasicInfo Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_BasicInfo

    // Field Replaceable Unit (FRU) inventory information.
    FruInfo Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_FruInfo
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes) GetEntityData() *types.CommonEntityData {
    basicAttributes.EntityData.YFilter = basicAttributes.YFilter
    basicAttributes.EntityData.YangName = "basic-attributes"
    basicAttributes.EntityData.BundleName = "cisco_ios_xr"
    basicAttributes.EntityData.ParentYangName = "sub-slot"
    basicAttributes.EntityData.SegmentPath = "basic-attributes"
    basicAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicAttributes.EntityData.Children = make(map[string]types.YChild)
    basicAttributes.EntityData.Children["basic-info"] = types.YChild{"BasicInfo", &basicAttributes.BasicInfo}
    basicAttributes.EntityData.Children["fru-info"] = types.YChild{"FruInfo", &basicAttributes.FruInfo}
    basicAttributes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(basicAttributes.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_BasicInfo
// Inventory information
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_BasicInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // describes in user-readable terms       what the entity in question does.
    // The type is string with length: 0..255.
    Description interface{}

    // maps to the vendor OID string. The type is string with length: 0..255.
    VendorType interface{}

    // name string for the entity. The type is string with length: 0..255.
    Name interface{}

    // hw revision string. The type is string with length: 0..255.
    HardwareRevision interface{}

    // firmware revision string. The type is string with length: 0..255.
    FirmwareRevision interface{}

    // software revision string. The type is string with length: 0..255.
    SoftwareRevision interface{}

    // chip module hw revision string. The type is string with length: 0..255.
    ChipHardwareRevision interface{}

    // serial number. The type is string with length: 0..255.
    SerialNumber interface{}

    // manufacturer's name. The type is string with length: 0..255.
    ManufacturerName interface{}

    // model name. The type is string with length: 0..255.
    ModelName interface{}

    // asset Identification string. The type is string with length: 0..255.
    AssetIdStr interface{}

    // asset Identification. The type is interface{} with range:
    // -2147483648..2147483647.
    AssetIdentification interface{}

    // 1 if Field Replaceable Unit 0, if not. The type is bool.
    IsFieldReplaceableUnit interface{}

    // Manufacture Asset Tags. The type is interface{} with range:
    // -2147483648..2147483647.
    ManufacturerAssetTags interface{}

    // Major&minor class of the entity. The type is interface{} with range:
    // -2147483648..2147483647.
    CompositeClassCode interface{}

    // Size of memory associated with       the entity where applicable. The type
    // is interface{} with range: -2147483648..2147483647.
    MemorySize interface{}

    // sysdb name of sensor in the envmon EDM. The type is string with length:
    // 0..255.
    EnvironmentalMonitorPath interface{}

    // useful for storing an entity alias . The type is string with length:
    // 0..255.
    Alias interface{}

    // indicates if this entity is group       or not. The type is bool.
    GroupFlag interface{}

    // integer value for New Deviation Number 0x88. The type is interface{} with
    // range: -2147483648..2147483647.
    NewDeviationNumber interface{}

    // integer value for plim type if     applicable to this entity. The type is
    // interface{} with range: -2147483648..2147483647.
    PhysicalLayerInterfaceModuleType interface{}

    // 1 if UnrecognizedFRU and 0 for recognizedFRU. The type is bool.
    UnrecognizedFru interface{}

    // integer value for Redundancy State if     applicable to this entity. The
    // type is interface{} with range: -2147483648..2147483647.
    Redundancystate interface{}

    // 1 if ce port found, 0 if not. The type is bool.
    Ceport interface{}

    // 1 if xr scoped, 0 if not. The type is bool.
    XrScoped interface{}

    // Unique id for an entity. The type is interface{} with range:
    // -2147483648..2147483647.
    UniqueId interface{}
}

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_BasicInfo) GetEntityData() *types.CommonEntityData {
    basicInfo.EntityData.YFilter = basicInfo.YFilter
    basicInfo.EntityData.YangName = "basic-info"
    basicInfo.EntityData.BundleName = "cisco_ios_xr"
    basicInfo.EntityData.ParentYangName = "basic-attributes"
    basicInfo.EntityData.SegmentPath = "basic-info"
    basicInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicInfo.EntityData.Children = make(map[string]types.YChild)
    basicInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    basicInfo.EntityData.Leafs["description"] = types.YLeaf{"Description", basicInfo.Description}
    basicInfo.EntityData.Leafs["vendor-type"] = types.YLeaf{"VendorType", basicInfo.VendorType}
    basicInfo.EntityData.Leafs["name"] = types.YLeaf{"Name", basicInfo.Name}
    basicInfo.EntityData.Leafs["hardware-revision"] = types.YLeaf{"HardwareRevision", basicInfo.HardwareRevision}
    basicInfo.EntityData.Leafs["firmware-revision"] = types.YLeaf{"FirmwareRevision", basicInfo.FirmwareRevision}
    basicInfo.EntityData.Leafs["software-revision"] = types.YLeaf{"SoftwareRevision", basicInfo.SoftwareRevision}
    basicInfo.EntityData.Leafs["chip-hardware-revision"] = types.YLeaf{"ChipHardwareRevision", basicInfo.ChipHardwareRevision}
    basicInfo.EntityData.Leafs["serial-number"] = types.YLeaf{"SerialNumber", basicInfo.SerialNumber}
    basicInfo.EntityData.Leafs["manufacturer-name"] = types.YLeaf{"ManufacturerName", basicInfo.ManufacturerName}
    basicInfo.EntityData.Leafs["model-name"] = types.YLeaf{"ModelName", basicInfo.ModelName}
    basicInfo.EntityData.Leafs["asset-id-str"] = types.YLeaf{"AssetIdStr", basicInfo.AssetIdStr}
    basicInfo.EntityData.Leafs["asset-identification"] = types.YLeaf{"AssetIdentification", basicInfo.AssetIdentification}
    basicInfo.EntityData.Leafs["is-field-replaceable-unit"] = types.YLeaf{"IsFieldReplaceableUnit", basicInfo.IsFieldReplaceableUnit}
    basicInfo.EntityData.Leafs["manufacturer-asset-tags"] = types.YLeaf{"ManufacturerAssetTags", basicInfo.ManufacturerAssetTags}
    basicInfo.EntityData.Leafs["composite-class-code"] = types.YLeaf{"CompositeClassCode", basicInfo.CompositeClassCode}
    basicInfo.EntityData.Leafs["memory-size"] = types.YLeaf{"MemorySize", basicInfo.MemorySize}
    basicInfo.EntityData.Leafs["environmental-monitor-path"] = types.YLeaf{"EnvironmentalMonitorPath", basicInfo.EnvironmentalMonitorPath}
    basicInfo.EntityData.Leafs["alias"] = types.YLeaf{"Alias", basicInfo.Alias}
    basicInfo.EntityData.Leafs["group-flag"] = types.YLeaf{"GroupFlag", basicInfo.GroupFlag}
    basicInfo.EntityData.Leafs["new-deviation-number"] = types.YLeaf{"NewDeviationNumber", basicInfo.NewDeviationNumber}
    basicInfo.EntityData.Leafs["physical-layer-interface-module-type"] = types.YLeaf{"PhysicalLayerInterfaceModuleType", basicInfo.PhysicalLayerInterfaceModuleType}
    basicInfo.EntityData.Leafs["unrecognized-fru"] = types.YLeaf{"UnrecognizedFru", basicInfo.UnrecognizedFru}
    basicInfo.EntityData.Leafs["redundancystate"] = types.YLeaf{"Redundancystate", basicInfo.Redundancystate}
    basicInfo.EntityData.Leafs["ceport"] = types.YLeaf{"Ceport", basicInfo.Ceport}
    basicInfo.EntityData.Leafs["xr-scoped"] = types.YLeaf{"XrScoped", basicInfo.XrScoped}
    basicInfo.EntityData.Leafs["unique-id"] = types.YLeaf{"UniqueId", basicInfo.UniqueId}
    return &(basicInfo.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_FruInfo
// Field Replaceable Unit (FRU) inventory
// information
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_FruInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // card admin state: shutdown or not. The type is interface{} with range:
    // -2147483648..2147483647.
    CardAdministrativeState interface{}

    // power admin state: up or down. The type is interface{} with range:
    // -2147483648..2147483647.
    PowerAdministrativeState interface{}

    // card operation state. The type is interface{} with range:
    // -2147483648..2147483647.
    CardOperationalState interface{}

    // card is monitored by a manager or left unmonitored. The type is interface{}
    // with range: -2147483648..2147483647.
    CardMonitorState interface{}

    // card reset reason enum. The type is CardResetReason.
    CardResetReason interface{}

    // power current: not implemented. The type is interface{} with range:
    // -2147483648..2147483647.
    PowerCurrentMeasurement interface{}

    // Power operation state. The type is interface{} with range:
    // -2147483648..2147483647.
    PowerOperationalState interface{}

    // last card oper change state.
    LastOperationalStateChange Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_FruInfo_LastOperationalStateChange

    // card up time.
    CardUpTime Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_FruInfo_CardUpTime
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_FruInfo) GetEntityData() *types.CommonEntityData {
    fruInfo.EntityData.YFilter = fruInfo.YFilter
    fruInfo.EntityData.YangName = "fru-info"
    fruInfo.EntityData.BundleName = "cisco_ios_xr"
    fruInfo.EntityData.ParentYangName = "basic-attributes"
    fruInfo.EntityData.SegmentPath = "fru-info"
    fruInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fruInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fruInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fruInfo.EntityData.Children = make(map[string]types.YChild)
    fruInfo.EntityData.Children["last-operational-state-change"] = types.YChild{"LastOperationalStateChange", &fruInfo.LastOperationalStateChange}
    fruInfo.EntityData.Children["card-up-time"] = types.YChild{"CardUpTime", &fruInfo.CardUpTime}
    fruInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    fruInfo.EntityData.Leafs["card-administrative-state"] = types.YLeaf{"CardAdministrativeState", fruInfo.CardAdministrativeState}
    fruInfo.EntityData.Leafs["power-administrative-state"] = types.YLeaf{"PowerAdministrativeState", fruInfo.PowerAdministrativeState}
    fruInfo.EntityData.Leafs["card-operational-state"] = types.YLeaf{"CardOperationalState", fruInfo.CardOperationalState}
    fruInfo.EntityData.Leafs["card-monitor-state"] = types.YLeaf{"CardMonitorState", fruInfo.CardMonitorState}
    fruInfo.EntityData.Leafs["card-reset-reason"] = types.YLeaf{"CardResetReason", fruInfo.CardResetReason}
    fruInfo.EntityData.Leafs["power-current-measurement"] = types.YLeaf{"PowerCurrentMeasurement", fruInfo.PowerCurrentMeasurement}
    fruInfo.EntityData.Leafs["power-operational-state"] = types.YLeaf{"PowerOperationalState", fruInfo.PowerOperationalState}
    return &(fruInfo.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_FruInfo_LastOperationalStateChange
// last card oper change state
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_FruInfo_LastOperationalStateChange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_FruInfo_LastOperationalStateChange) GetEntityData() *types.CommonEntityData {
    lastOperationalStateChange.EntityData.YFilter = lastOperationalStateChange.YFilter
    lastOperationalStateChange.EntityData.YangName = "last-operational-state-change"
    lastOperationalStateChange.EntityData.BundleName = "cisco_ios_xr"
    lastOperationalStateChange.EntityData.ParentYangName = "fru-info"
    lastOperationalStateChange.EntityData.SegmentPath = "last-operational-state-change"
    lastOperationalStateChange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastOperationalStateChange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastOperationalStateChange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastOperationalStateChange.EntityData.Children = make(map[string]types.YChild)
    lastOperationalStateChange.EntityData.Leafs = make(map[string]types.YLeaf)
    lastOperationalStateChange.EntityData.Leafs["time-in-seconds"] = types.YLeaf{"TimeInSeconds", lastOperationalStateChange.TimeInSeconds}
    lastOperationalStateChange.EntityData.Leafs["time-in-nano-seconds"] = types.YLeaf{"TimeInNanoSeconds", lastOperationalStateChange.TimeInNanoSeconds}
    return &(lastOperationalStateChange.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_FruInfo_CardUpTime
// card up time
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_FruInfo_CardUpTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_FruInfo_CardUpTime) GetEntityData() *types.CommonEntityData {
    cardUpTime.EntityData.YFilter = cardUpTime.YFilter
    cardUpTime.EntityData.YangName = "card-up-time"
    cardUpTime.EntityData.BundleName = "cisco_ios_xr"
    cardUpTime.EntityData.ParentYangName = "fru-info"
    cardUpTime.EntityData.SegmentPath = "card-up-time"
    cardUpTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cardUpTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cardUpTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cardUpTime.EntityData.Children = make(map[string]types.YChild)
    cardUpTime.EntityData.Leafs = make(map[string]types.YLeaf)
    cardUpTime.EntityData.Leafs["time-in-seconds"] = types.YLeaf{"TimeInSeconds", cardUpTime.TimeInSeconds}
    cardUpTime.EntityData.Leafs["time-in-nano-seconds"] = types.YLeaf{"TimeInNanoSeconds", cardUpTime.TimeInNanoSeconds}
    return &(cardUpTime.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents
// HWComponent table contains all HW modules
// within the card 
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // HWComponent number. The type is slice of
    // Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent.
    HwComponent []Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent
}

func (hwComponents *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents) GetEntityData() *types.CommonEntityData {
    hwComponents.EntityData.YFilter = hwComponents.YFilter
    hwComponents.EntityData.YangName = "hw-components"
    hwComponents.EntityData.BundleName = "cisco_ios_xr"
    hwComponents.EntityData.ParentYangName = "card"
    hwComponents.EntityData.SegmentPath = "hw-components"
    hwComponents.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hwComponents.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hwComponents.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hwComponents.EntityData.Children = make(map[string]types.YChild)
    hwComponents.EntityData.Children["hw-component"] = types.YChild{"HwComponent", nil}
    for i := range hwComponents.HwComponent {
        hwComponents.EntityData.Children[types.GetSegmentPath(&hwComponents.HwComponent[i])] = types.YChild{"HwComponent", &hwComponents.HwComponent[i]}
    }
    hwComponents.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(hwComponents.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent
// HWComponent number
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. node number. The type is interface{} with range:
    // -2147483648..2147483647.
    Number interface{}

    // ModuleSensorTable contains all sensors in a Module.
    Sensors Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors

    // Attributes.
    BasicAttributes Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes
}

func (hwComponent *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent) GetEntityData() *types.CommonEntityData {
    hwComponent.EntityData.YFilter = hwComponent.YFilter
    hwComponent.EntityData.YangName = "hw-component"
    hwComponent.EntityData.BundleName = "cisco_ios_xr"
    hwComponent.EntityData.ParentYangName = "hw-components"
    hwComponent.EntityData.SegmentPath = "hw-component" + "[number='" + fmt.Sprintf("%v", hwComponent.Number) + "']"
    hwComponent.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hwComponent.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hwComponent.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hwComponent.EntityData.Children = make(map[string]types.YChild)
    hwComponent.EntityData.Children["sensors"] = types.YChild{"Sensors", &hwComponent.Sensors}
    hwComponent.EntityData.Children["basic-attributes"] = types.YChild{"BasicAttributes", &hwComponent.BasicAttributes}
    hwComponent.EntityData.Leafs = make(map[string]types.YLeaf)
    hwComponent.EntityData.Leafs["number"] = types.YLeaf{"Number", hwComponent.Number}
    return &(hwComponent.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors
// ModuleSensorTable contains all sensors in a
// Module.
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sensor number in the Module. The type is slice of
    // Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor.
    Sensor []Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor
}

func (sensors *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors) GetEntityData() *types.CommonEntityData {
    sensors.EntityData.YFilter = sensors.YFilter
    sensors.EntityData.YangName = "sensors"
    sensors.EntityData.BundleName = "cisco_ios_xr"
    sensors.EntityData.ParentYangName = "hw-component"
    sensors.EntityData.SegmentPath = "sensors"
    sensors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sensors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sensors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sensors.EntityData.Children = make(map[string]types.YChild)
    sensors.EntityData.Children["sensor"] = types.YChild{"Sensor", nil}
    for i := range sensors.Sensor {
        sensors.EntityData.Children[types.GetSegmentPath(&sensors.Sensor[i])] = types.YChild{"Sensor", &sensors.Sensor[i]}
    }
    sensors.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sensors.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor
// Sensor number in the Module
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. sensor number. The type is interface{} with range:
    // -2147483648..2147483647.
    Number interface{}

    // Attributes.
    BasicAttributes Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes
}

func (sensor *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor) GetEntityData() *types.CommonEntityData {
    sensor.EntityData.YFilter = sensor.YFilter
    sensor.EntityData.YangName = "sensor"
    sensor.EntityData.BundleName = "cisco_ios_xr"
    sensor.EntityData.ParentYangName = "sensors"
    sensor.EntityData.SegmentPath = "sensor" + "[number='" + fmt.Sprintf("%v", sensor.Number) + "']"
    sensor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sensor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sensor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sensor.EntityData.Children = make(map[string]types.YChild)
    sensor.EntityData.Children["basic-attributes"] = types.YChild{"BasicAttributes", &sensor.BasicAttributes}
    sensor.EntityData.Leafs = make(map[string]types.YLeaf)
    sensor.EntityData.Leafs["number"] = types.YLeaf{"Number", sensor.Number}
    return &(sensor.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes
// Attributes
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Inventory information.
    BasicInfo Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_BasicInfo

    // Field Replaceable Unit (FRU) inventory information.
    FruInfo Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_FruInfo
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes) GetEntityData() *types.CommonEntityData {
    basicAttributes.EntityData.YFilter = basicAttributes.YFilter
    basicAttributes.EntityData.YangName = "basic-attributes"
    basicAttributes.EntityData.BundleName = "cisco_ios_xr"
    basicAttributes.EntityData.ParentYangName = "sensor"
    basicAttributes.EntityData.SegmentPath = "basic-attributes"
    basicAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicAttributes.EntityData.Children = make(map[string]types.YChild)
    basicAttributes.EntityData.Children["basic-info"] = types.YChild{"BasicInfo", &basicAttributes.BasicInfo}
    basicAttributes.EntityData.Children["fru-info"] = types.YChild{"FruInfo", &basicAttributes.FruInfo}
    basicAttributes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(basicAttributes.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_BasicInfo
// Inventory information
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_BasicInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // describes in user-readable terms       what the entity in question does.
    // The type is string with length: 0..255.
    Description interface{}

    // maps to the vendor OID string. The type is string with length: 0..255.
    VendorType interface{}

    // name string for the entity. The type is string with length: 0..255.
    Name interface{}

    // hw revision string. The type is string with length: 0..255.
    HardwareRevision interface{}

    // firmware revision string. The type is string with length: 0..255.
    FirmwareRevision interface{}

    // software revision string. The type is string with length: 0..255.
    SoftwareRevision interface{}

    // chip module hw revision string. The type is string with length: 0..255.
    ChipHardwareRevision interface{}

    // serial number. The type is string with length: 0..255.
    SerialNumber interface{}

    // manufacturer's name. The type is string with length: 0..255.
    ManufacturerName interface{}

    // model name. The type is string with length: 0..255.
    ModelName interface{}

    // asset Identification string. The type is string with length: 0..255.
    AssetIdStr interface{}

    // asset Identification. The type is interface{} with range:
    // -2147483648..2147483647.
    AssetIdentification interface{}

    // 1 if Field Replaceable Unit 0, if not. The type is bool.
    IsFieldReplaceableUnit interface{}

    // Manufacture Asset Tags. The type is interface{} with range:
    // -2147483648..2147483647.
    ManufacturerAssetTags interface{}

    // Major&minor class of the entity. The type is interface{} with range:
    // -2147483648..2147483647.
    CompositeClassCode interface{}

    // Size of memory associated with       the entity where applicable. The type
    // is interface{} with range: -2147483648..2147483647.
    MemorySize interface{}

    // sysdb name of sensor in the envmon EDM. The type is string with length:
    // 0..255.
    EnvironmentalMonitorPath interface{}

    // useful for storing an entity alias . The type is string with length:
    // 0..255.
    Alias interface{}

    // indicates if this entity is group       or not. The type is bool.
    GroupFlag interface{}

    // integer value for New Deviation Number 0x88. The type is interface{} with
    // range: -2147483648..2147483647.
    NewDeviationNumber interface{}

    // integer value for plim type if     applicable to this entity. The type is
    // interface{} with range: -2147483648..2147483647.
    PhysicalLayerInterfaceModuleType interface{}

    // 1 if UnrecognizedFRU and 0 for recognizedFRU. The type is bool.
    UnrecognizedFru interface{}

    // integer value for Redundancy State if     applicable to this entity. The
    // type is interface{} with range: -2147483648..2147483647.
    Redundancystate interface{}

    // 1 if ce port found, 0 if not. The type is bool.
    Ceport interface{}

    // 1 if xr scoped, 0 if not. The type is bool.
    XrScoped interface{}

    // Unique id for an entity. The type is interface{} with range:
    // -2147483648..2147483647.
    UniqueId interface{}
}

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_BasicInfo) GetEntityData() *types.CommonEntityData {
    basicInfo.EntityData.YFilter = basicInfo.YFilter
    basicInfo.EntityData.YangName = "basic-info"
    basicInfo.EntityData.BundleName = "cisco_ios_xr"
    basicInfo.EntityData.ParentYangName = "basic-attributes"
    basicInfo.EntityData.SegmentPath = "basic-info"
    basicInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicInfo.EntityData.Children = make(map[string]types.YChild)
    basicInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    basicInfo.EntityData.Leafs["description"] = types.YLeaf{"Description", basicInfo.Description}
    basicInfo.EntityData.Leafs["vendor-type"] = types.YLeaf{"VendorType", basicInfo.VendorType}
    basicInfo.EntityData.Leafs["name"] = types.YLeaf{"Name", basicInfo.Name}
    basicInfo.EntityData.Leafs["hardware-revision"] = types.YLeaf{"HardwareRevision", basicInfo.HardwareRevision}
    basicInfo.EntityData.Leafs["firmware-revision"] = types.YLeaf{"FirmwareRevision", basicInfo.FirmwareRevision}
    basicInfo.EntityData.Leafs["software-revision"] = types.YLeaf{"SoftwareRevision", basicInfo.SoftwareRevision}
    basicInfo.EntityData.Leafs["chip-hardware-revision"] = types.YLeaf{"ChipHardwareRevision", basicInfo.ChipHardwareRevision}
    basicInfo.EntityData.Leafs["serial-number"] = types.YLeaf{"SerialNumber", basicInfo.SerialNumber}
    basicInfo.EntityData.Leafs["manufacturer-name"] = types.YLeaf{"ManufacturerName", basicInfo.ManufacturerName}
    basicInfo.EntityData.Leafs["model-name"] = types.YLeaf{"ModelName", basicInfo.ModelName}
    basicInfo.EntityData.Leafs["asset-id-str"] = types.YLeaf{"AssetIdStr", basicInfo.AssetIdStr}
    basicInfo.EntityData.Leafs["asset-identification"] = types.YLeaf{"AssetIdentification", basicInfo.AssetIdentification}
    basicInfo.EntityData.Leafs["is-field-replaceable-unit"] = types.YLeaf{"IsFieldReplaceableUnit", basicInfo.IsFieldReplaceableUnit}
    basicInfo.EntityData.Leafs["manufacturer-asset-tags"] = types.YLeaf{"ManufacturerAssetTags", basicInfo.ManufacturerAssetTags}
    basicInfo.EntityData.Leafs["composite-class-code"] = types.YLeaf{"CompositeClassCode", basicInfo.CompositeClassCode}
    basicInfo.EntityData.Leafs["memory-size"] = types.YLeaf{"MemorySize", basicInfo.MemorySize}
    basicInfo.EntityData.Leafs["environmental-monitor-path"] = types.YLeaf{"EnvironmentalMonitorPath", basicInfo.EnvironmentalMonitorPath}
    basicInfo.EntityData.Leafs["alias"] = types.YLeaf{"Alias", basicInfo.Alias}
    basicInfo.EntityData.Leafs["group-flag"] = types.YLeaf{"GroupFlag", basicInfo.GroupFlag}
    basicInfo.EntityData.Leafs["new-deviation-number"] = types.YLeaf{"NewDeviationNumber", basicInfo.NewDeviationNumber}
    basicInfo.EntityData.Leafs["physical-layer-interface-module-type"] = types.YLeaf{"PhysicalLayerInterfaceModuleType", basicInfo.PhysicalLayerInterfaceModuleType}
    basicInfo.EntityData.Leafs["unrecognized-fru"] = types.YLeaf{"UnrecognizedFru", basicInfo.UnrecognizedFru}
    basicInfo.EntityData.Leafs["redundancystate"] = types.YLeaf{"Redundancystate", basicInfo.Redundancystate}
    basicInfo.EntityData.Leafs["ceport"] = types.YLeaf{"Ceport", basicInfo.Ceport}
    basicInfo.EntityData.Leafs["xr-scoped"] = types.YLeaf{"XrScoped", basicInfo.XrScoped}
    basicInfo.EntityData.Leafs["unique-id"] = types.YLeaf{"UniqueId", basicInfo.UniqueId}
    return &(basicInfo.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_FruInfo
// Field Replaceable Unit (FRU) inventory
// information
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_FruInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // card admin state: shutdown or not. The type is interface{} with range:
    // -2147483648..2147483647.
    CardAdministrativeState interface{}

    // power admin state: up or down. The type is interface{} with range:
    // -2147483648..2147483647.
    PowerAdministrativeState interface{}

    // card operation state. The type is interface{} with range:
    // -2147483648..2147483647.
    CardOperationalState interface{}

    // card is monitored by a manager or left unmonitored. The type is interface{}
    // with range: -2147483648..2147483647.
    CardMonitorState interface{}

    // card reset reason enum. The type is CardResetReason.
    CardResetReason interface{}

    // power current: not implemented. The type is interface{} with range:
    // -2147483648..2147483647.
    PowerCurrentMeasurement interface{}

    // Power operation state. The type is interface{} with range:
    // -2147483648..2147483647.
    PowerOperationalState interface{}

    // last card oper change state.
    LastOperationalStateChange Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_FruInfo_LastOperationalStateChange

    // card up time.
    CardUpTime Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_FruInfo_CardUpTime
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_FruInfo) GetEntityData() *types.CommonEntityData {
    fruInfo.EntityData.YFilter = fruInfo.YFilter
    fruInfo.EntityData.YangName = "fru-info"
    fruInfo.EntityData.BundleName = "cisco_ios_xr"
    fruInfo.EntityData.ParentYangName = "basic-attributes"
    fruInfo.EntityData.SegmentPath = "fru-info"
    fruInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fruInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fruInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fruInfo.EntityData.Children = make(map[string]types.YChild)
    fruInfo.EntityData.Children["last-operational-state-change"] = types.YChild{"LastOperationalStateChange", &fruInfo.LastOperationalStateChange}
    fruInfo.EntityData.Children["card-up-time"] = types.YChild{"CardUpTime", &fruInfo.CardUpTime}
    fruInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    fruInfo.EntityData.Leafs["card-administrative-state"] = types.YLeaf{"CardAdministrativeState", fruInfo.CardAdministrativeState}
    fruInfo.EntityData.Leafs["power-administrative-state"] = types.YLeaf{"PowerAdministrativeState", fruInfo.PowerAdministrativeState}
    fruInfo.EntityData.Leafs["card-operational-state"] = types.YLeaf{"CardOperationalState", fruInfo.CardOperationalState}
    fruInfo.EntityData.Leafs["card-monitor-state"] = types.YLeaf{"CardMonitorState", fruInfo.CardMonitorState}
    fruInfo.EntityData.Leafs["card-reset-reason"] = types.YLeaf{"CardResetReason", fruInfo.CardResetReason}
    fruInfo.EntityData.Leafs["power-current-measurement"] = types.YLeaf{"PowerCurrentMeasurement", fruInfo.PowerCurrentMeasurement}
    fruInfo.EntityData.Leafs["power-operational-state"] = types.YLeaf{"PowerOperationalState", fruInfo.PowerOperationalState}
    return &(fruInfo.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_FruInfo_LastOperationalStateChange
// last card oper change state
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_FruInfo_LastOperationalStateChange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_FruInfo_LastOperationalStateChange) GetEntityData() *types.CommonEntityData {
    lastOperationalStateChange.EntityData.YFilter = lastOperationalStateChange.YFilter
    lastOperationalStateChange.EntityData.YangName = "last-operational-state-change"
    lastOperationalStateChange.EntityData.BundleName = "cisco_ios_xr"
    lastOperationalStateChange.EntityData.ParentYangName = "fru-info"
    lastOperationalStateChange.EntityData.SegmentPath = "last-operational-state-change"
    lastOperationalStateChange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastOperationalStateChange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastOperationalStateChange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastOperationalStateChange.EntityData.Children = make(map[string]types.YChild)
    lastOperationalStateChange.EntityData.Leafs = make(map[string]types.YLeaf)
    lastOperationalStateChange.EntityData.Leafs["time-in-seconds"] = types.YLeaf{"TimeInSeconds", lastOperationalStateChange.TimeInSeconds}
    lastOperationalStateChange.EntityData.Leafs["time-in-nano-seconds"] = types.YLeaf{"TimeInNanoSeconds", lastOperationalStateChange.TimeInNanoSeconds}
    return &(lastOperationalStateChange.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_FruInfo_CardUpTime
// card up time
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_FruInfo_CardUpTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_FruInfo_CardUpTime) GetEntityData() *types.CommonEntityData {
    cardUpTime.EntityData.YFilter = cardUpTime.YFilter
    cardUpTime.EntityData.YangName = "card-up-time"
    cardUpTime.EntityData.BundleName = "cisco_ios_xr"
    cardUpTime.EntityData.ParentYangName = "fru-info"
    cardUpTime.EntityData.SegmentPath = "card-up-time"
    cardUpTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cardUpTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cardUpTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cardUpTime.EntityData.Children = make(map[string]types.YChild)
    cardUpTime.EntityData.Leafs = make(map[string]types.YLeaf)
    cardUpTime.EntityData.Leafs["time-in-seconds"] = types.YLeaf{"TimeInSeconds", cardUpTime.TimeInSeconds}
    cardUpTime.EntityData.Leafs["time-in-nano-seconds"] = types.YLeaf{"TimeInNanoSeconds", cardUpTime.TimeInNanoSeconds}
    return &(cardUpTime.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes
// Attributes
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Inventory information.
    BasicInfo Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_BasicInfo

    // Field Replaceable Unit (FRU) inventory information.
    FruInfo Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_FruInfo
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes) GetEntityData() *types.CommonEntityData {
    basicAttributes.EntityData.YFilter = basicAttributes.YFilter
    basicAttributes.EntityData.YangName = "basic-attributes"
    basicAttributes.EntityData.BundleName = "cisco_ios_xr"
    basicAttributes.EntityData.ParentYangName = "hw-component"
    basicAttributes.EntityData.SegmentPath = "basic-attributes"
    basicAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicAttributes.EntityData.Children = make(map[string]types.YChild)
    basicAttributes.EntityData.Children["basic-info"] = types.YChild{"BasicInfo", &basicAttributes.BasicInfo}
    basicAttributes.EntityData.Children["fru-info"] = types.YChild{"FruInfo", &basicAttributes.FruInfo}
    basicAttributes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(basicAttributes.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_BasicInfo
// Inventory information
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_BasicInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // describes in user-readable terms       what the entity in question does.
    // The type is string with length: 0..255.
    Description interface{}

    // maps to the vendor OID string. The type is string with length: 0..255.
    VendorType interface{}

    // name string for the entity. The type is string with length: 0..255.
    Name interface{}

    // hw revision string. The type is string with length: 0..255.
    HardwareRevision interface{}

    // firmware revision string. The type is string with length: 0..255.
    FirmwareRevision interface{}

    // software revision string. The type is string with length: 0..255.
    SoftwareRevision interface{}

    // chip module hw revision string. The type is string with length: 0..255.
    ChipHardwareRevision interface{}

    // serial number. The type is string with length: 0..255.
    SerialNumber interface{}

    // manufacturer's name. The type is string with length: 0..255.
    ManufacturerName interface{}

    // model name. The type is string with length: 0..255.
    ModelName interface{}

    // asset Identification string. The type is string with length: 0..255.
    AssetIdStr interface{}

    // asset Identification. The type is interface{} with range:
    // -2147483648..2147483647.
    AssetIdentification interface{}

    // 1 if Field Replaceable Unit 0, if not. The type is bool.
    IsFieldReplaceableUnit interface{}

    // Manufacture Asset Tags. The type is interface{} with range:
    // -2147483648..2147483647.
    ManufacturerAssetTags interface{}

    // Major&minor class of the entity. The type is interface{} with range:
    // -2147483648..2147483647.
    CompositeClassCode interface{}

    // Size of memory associated with       the entity where applicable. The type
    // is interface{} with range: -2147483648..2147483647.
    MemorySize interface{}

    // sysdb name of sensor in the envmon EDM. The type is string with length:
    // 0..255.
    EnvironmentalMonitorPath interface{}

    // useful for storing an entity alias . The type is string with length:
    // 0..255.
    Alias interface{}

    // indicates if this entity is group       or not. The type is bool.
    GroupFlag interface{}

    // integer value for New Deviation Number 0x88. The type is interface{} with
    // range: -2147483648..2147483647.
    NewDeviationNumber interface{}

    // integer value for plim type if     applicable to this entity. The type is
    // interface{} with range: -2147483648..2147483647.
    PhysicalLayerInterfaceModuleType interface{}

    // 1 if UnrecognizedFRU and 0 for recognizedFRU. The type is bool.
    UnrecognizedFru interface{}

    // integer value for Redundancy State if     applicable to this entity. The
    // type is interface{} with range: -2147483648..2147483647.
    Redundancystate interface{}

    // 1 if ce port found, 0 if not. The type is bool.
    Ceport interface{}

    // 1 if xr scoped, 0 if not. The type is bool.
    XrScoped interface{}

    // Unique id for an entity. The type is interface{} with range:
    // -2147483648..2147483647.
    UniqueId interface{}
}

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_BasicInfo) GetEntityData() *types.CommonEntityData {
    basicInfo.EntityData.YFilter = basicInfo.YFilter
    basicInfo.EntityData.YangName = "basic-info"
    basicInfo.EntityData.BundleName = "cisco_ios_xr"
    basicInfo.EntityData.ParentYangName = "basic-attributes"
    basicInfo.EntityData.SegmentPath = "basic-info"
    basicInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicInfo.EntityData.Children = make(map[string]types.YChild)
    basicInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    basicInfo.EntityData.Leafs["description"] = types.YLeaf{"Description", basicInfo.Description}
    basicInfo.EntityData.Leafs["vendor-type"] = types.YLeaf{"VendorType", basicInfo.VendorType}
    basicInfo.EntityData.Leafs["name"] = types.YLeaf{"Name", basicInfo.Name}
    basicInfo.EntityData.Leafs["hardware-revision"] = types.YLeaf{"HardwareRevision", basicInfo.HardwareRevision}
    basicInfo.EntityData.Leafs["firmware-revision"] = types.YLeaf{"FirmwareRevision", basicInfo.FirmwareRevision}
    basicInfo.EntityData.Leafs["software-revision"] = types.YLeaf{"SoftwareRevision", basicInfo.SoftwareRevision}
    basicInfo.EntityData.Leafs["chip-hardware-revision"] = types.YLeaf{"ChipHardwareRevision", basicInfo.ChipHardwareRevision}
    basicInfo.EntityData.Leafs["serial-number"] = types.YLeaf{"SerialNumber", basicInfo.SerialNumber}
    basicInfo.EntityData.Leafs["manufacturer-name"] = types.YLeaf{"ManufacturerName", basicInfo.ManufacturerName}
    basicInfo.EntityData.Leafs["model-name"] = types.YLeaf{"ModelName", basicInfo.ModelName}
    basicInfo.EntityData.Leafs["asset-id-str"] = types.YLeaf{"AssetIdStr", basicInfo.AssetIdStr}
    basicInfo.EntityData.Leafs["asset-identification"] = types.YLeaf{"AssetIdentification", basicInfo.AssetIdentification}
    basicInfo.EntityData.Leafs["is-field-replaceable-unit"] = types.YLeaf{"IsFieldReplaceableUnit", basicInfo.IsFieldReplaceableUnit}
    basicInfo.EntityData.Leafs["manufacturer-asset-tags"] = types.YLeaf{"ManufacturerAssetTags", basicInfo.ManufacturerAssetTags}
    basicInfo.EntityData.Leafs["composite-class-code"] = types.YLeaf{"CompositeClassCode", basicInfo.CompositeClassCode}
    basicInfo.EntityData.Leafs["memory-size"] = types.YLeaf{"MemorySize", basicInfo.MemorySize}
    basicInfo.EntityData.Leafs["environmental-monitor-path"] = types.YLeaf{"EnvironmentalMonitorPath", basicInfo.EnvironmentalMonitorPath}
    basicInfo.EntityData.Leafs["alias"] = types.YLeaf{"Alias", basicInfo.Alias}
    basicInfo.EntityData.Leafs["group-flag"] = types.YLeaf{"GroupFlag", basicInfo.GroupFlag}
    basicInfo.EntityData.Leafs["new-deviation-number"] = types.YLeaf{"NewDeviationNumber", basicInfo.NewDeviationNumber}
    basicInfo.EntityData.Leafs["physical-layer-interface-module-type"] = types.YLeaf{"PhysicalLayerInterfaceModuleType", basicInfo.PhysicalLayerInterfaceModuleType}
    basicInfo.EntityData.Leafs["unrecognized-fru"] = types.YLeaf{"UnrecognizedFru", basicInfo.UnrecognizedFru}
    basicInfo.EntityData.Leafs["redundancystate"] = types.YLeaf{"Redundancystate", basicInfo.Redundancystate}
    basicInfo.EntityData.Leafs["ceport"] = types.YLeaf{"Ceport", basicInfo.Ceport}
    basicInfo.EntityData.Leafs["xr-scoped"] = types.YLeaf{"XrScoped", basicInfo.XrScoped}
    basicInfo.EntityData.Leafs["unique-id"] = types.YLeaf{"UniqueId", basicInfo.UniqueId}
    return &(basicInfo.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_FruInfo
// Field Replaceable Unit (FRU) inventory
// information
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_FruInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // card admin state: shutdown or not. The type is interface{} with range:
    // -2147483648..2147483647.
    CardAdministrativeState interface{}

    // power admin state: up or down. The type is interface{} with range:
    // -2147483648..2147483647.
    PowerAdministrativeState interface{}

    // card operation state. The type is interface{} with range:
    // -2147483648..2147483647.
    CardOperationalState interface{}

    // card is monitored by a manager or left unmonitored. The type is interface{}
    // with range: -2147483648..2147483647.
    CardMonitorState interface{}

    // card reset reason enum. The type is CardResetReason.
    CardResetReason interface{}

    // power current: not implemented. The type is interface{} with range:
    // -2147483648..2147483647.
    PowerCurrentMeasurement interface{}

    // Power operation state. The type is interface{} with range:
    // -2147483648..2147483647.
    PowerOperationalState interface{}

    // last card oper change state.
    LastOperationalStateChange Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_FruInfo_LastOperationalStateChange

    // card up time.
    CardUpTime Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_FruInfo_CardUpTime
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_FruInfo) GetEntityData() *types.CommonEntityData {
    fruInfo.EntityData.YFilter = fruInfo.YFilter
    fruInfo.EntityData.YangName = "fru-info"
    fruInfo.EntityData.BundleName = "cisco_ios_xr"
    fruInfo.EntityData.ParentYangName = "basic-attributes"
    fruInfo.EntityData.SegmentPath = "fru-info"
    fruInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fruInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fruInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fruInfo.EntityData.Children = make(map[string]types.YChild)
    fruInfo.EntityData.Children["last-operational-state-change"] = types.YChild{"LastOperationalStateChange", &fruInfo.LastOperationalStateChange}
    fruInfo.EntityData.Children["card-up-time"] = types.YChild{"CardUpTime", &fruInfo.CardUpTime}
    fruInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    fruInfo.EntityData.Leafs["card-administrative-state"] = types.YLeaf{"CardAdministrativeState", fruInfo.CardAdministrativeState}
    fruInfo.EntityData.Leafs["power-administrative-state"] = types.YLeaf{"PowerAdministrativeState", fruInfo.PowerAdministrativeState}
    fruInfo.EntityData.Leafs["card-operational-state"] = types.YLeaf{"CardOperationalState", fruInfo.CardOperationalState}
    fruInfo.EntityData.Leafs["card-monitor-state"] = types.YLeaf{"CardMonitorState", fruInfo.CardMonitorState}
    fruInfo.EntityData.Leafs["card-reset-reason"] = types.YLeaf{"CardResetReason", fruInfo.CardResetReason}
    fruInfo.EntityData.Leafs["power-current-measurement"] = types.YLeaf{"PowerCurrentMeasurement", fruInfo.PowerCurrentMeasurement}
    fruInfo.EntityData.Leafs["power-operational-state"] = types.YLeaf{"PowerOperationalState", fruInfo.PowerOperationalState}
    return &(fruInfo.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_FruInfo_LastOperationalStateChange
// last card oper change state
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_FruInfo_LastOperationalStateChange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_FruInfo_LastOperationalStateChange) GetEntityData() *types.CommonEntityData {
    lastOperationalStateChange.EntityData.YFilter = lastOperationalStateChange.YFilter
    lastOperationalStateChange.EntityData.YangName = "last-operational-state-change"
    lastOperationalStateChange.EntityData.BundleName = "cisco_ios_xr"
    lastOperationalStateChange.EntityData.ParentYangName = "fru-info"
    lastOperationalStateChange.EntityData.SegmentPath = "last-operational-state-change"
    lastOperationalStateChange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastOperationalStateChange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastOperationalStateChange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastOperationalStateChange.EntityData.Children = make(map[string]types.YChild)
    lastOperationalStateChange.EntityData.Leafs = make(map[string]types.YLeaf)
    lastOperationalStateChange.EntityData.Leafs["time-in-seconds"] = types.YLeaf{"TimeInSeconds", lastOperationalStateChange.TimeInSeconds}
    lastOperationalStateChange.EntityData.Leafs["time-in-nano-seconds"] = types.YLeaf{"TimeInNanoSeconds", lastOperationalStateChange.TimeInNanoSeconds}
    return &(lastOperationalStateChange.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_FruInfo_CardUpTime
// card up time
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_FruInfo_CardUpTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_FruInfo_CardUpTime) GetEntityData() *types.CommonEntityData {
    cardUpTime.EntityData.YFilter = cardUpTime.YFilter
    cardUpTime.EntityData.YangName = "card-up-time"
    cardUpTime.EntityData.BundleName = "cisco_ios_xr"
    cardUpTime.EntityData.ParentYangName = "fru-info"
    cardUpTime.EntityData.SegmentPath = "card-up-time"
    cardUpTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cardUpTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cardUpTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cardUpTime.EntityData.Children = make(map[string]types.YChild)
    cardUpTime.EntityData.Leafs = make(map[string]types.YLeaf)
    cardUpTime.EntityData.Leafs["time-in-seconds"] = types.YLeaf{"TimeInSeconds", cardUpTime.TimeInSeconds}
    cardUpTime.EntityData.Leafs["time-in-nano-seconds"] = types.YLeaf{"TimeInNanoSeconds", cardUpTime.TimeInNanoSeconds}
    return &(cardUpTime.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors
// ModuleSensorTable contains all sensors in a
// Module.
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sensor number in the Module. The type is slice of
    // Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor.
    Sensor []Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor
}

func (sensors *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors) GetEntityData() *types.CommonEntityData {
    sensors.EntityData.YFilter = sensors.YFilter
    sensors.EntityData.YangName = "sensors"
    sensors.EntityData.BundleName = "cisco_ios_xr"
    sensors.EntityData.ParentYangName = "card"
    sensors.EntityData.SegmentPath = "sensors"
    sensors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sensors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sensors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sensors.EntityData.Children = make(map[string]types.YChild)
    sensors.EntityData.Children["sensor"] = types.YChild{"Sensor", nil}
    for i := range sensors.Sensor {
        sensors.EntityData.Children[types.GetSegmentPath(&sensors.Sensor[i])] = types.YChild{"Sensor", &sensors.Sensor[i]}
    }
    sensors.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sensors.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor
// Sensor number in the Module
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. sensor number. The type is interface{} with range:
    // -2147483648..2147483647.
    Number interface{}

    // Attributes.
    BasicAttributes Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes
}

func (sensor *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor) GetEntityData() *types.CommonEntityData {
    sensor.EntityData.YFilter = sensor.YFilter
    sensor.EntityData.YangName = "sensor"
    sensor.EntityData.BundleName = "cisco_ios_xr"
    sensor.EntityData.ParentYangName = "sensors"
    sensor.EntityData.SegmentPath = "sensor" + "[number='" + fmt.Sprintf("%v", sensor.Number) + "']"
    sensor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sensor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sensor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sensor.EntityData.Children = make(map[string]types.YChild)
    sensor.EntityData.Children["basic-attributes"] = types.YChild{"BasicAttributes", &sensor.BasicAttributes}
    sensor.EntityData.Leafs = make(map[string]types.YLeaf)
    sensor.EntityData.Leafs["number"] = types.YLeaf{"Number", sensor.Number}
    return &(sensor.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes
// Attributes
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Inventory information.
    BasicInfo Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_BasicInfo

    // Field Replaceable Unit (FRU) inventory information.
    FruInfo Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_FruInfo
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes) GetEntityData() *types.CommonEntityData {
    basicAttributes.EntityData.YFilter = basicAttributes.YFilter
    basicAttributes.EntityData.YangName = "basic-attributes"
    basicAttributes.EntityData.BundleName = "cisco_ios_xr"
    basicAttributes.EntityData.ParentYangName = "sensor"
    basicAttributes.EntityData.SegmentPath = "basic-attributes"
    basicAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicAttributes.EntityData.Children = make(map[string]types.YChild)
    basicAttributes.EntityData.Children["basic-info"] = types.YChild{"BasicInfo", &basicAttributes.BasicInfo}
    basicAttributes.EntityData.Children["fru-info"] = types.YChild{"FruInfo", &basicAttributes.FruInfo}
    basicAttributes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(basicAttributes.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_BasicInfo
// Inventory information
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_BasicInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // describes in user-readable terms       what the entity in question does.
    // The type is string with length: 0..255.
    Description interface{}

    // maps to the vendor OID string. The type is string with length: 0..255.
    VendorType interface{}

    // name string for the entity. The type is string with length: 0..255.
    Name interface{}

    // hw revision string. The type is string with length: 0..255.
    HardwareRevision interface{}

    // firmware revision string. The type is string with length: 0..255.
    FirmwareRevision interface{}

    // software revision string. The type is string with length: 0..255.
    SoftwareRevision interface{}

    // chip module hw revision string. The type is string with length: 0..255.
    ChipHardwareRevision interface{}

    // serial number. The type is string with length: 0..255.
    SerialNumber interface{}

    // manufacturer's name. The type is string with length: 0..255.
    ManufacturerName interface{}

    // model name. The type is string with length: 0..255.
    ModelName interface{}

    // asset Identification string. The type is string with length: 0..255.
    AssetIdStr interface{}

    // asset Identification. The type is interface{} with range:
    // -2147483648..2147483647.
    AssetIdentification interface{}

    // 1 if Field Replaceable Unit 0, if not. The type is bool.
    IsFieldReplaceableUnit interface{}

    // Manufacture Asset Tags. The type is interface{} with range:
    // -2147483648..2147483647.
    ManufacturerAssetTags interface{}

    // Major&minor class of the entity. The type is interface{} with range:
    // -2147483648..2147483647.
    CompositeClassCode interface{}

    // Size of memory associated with       the entity where applicable. The type
    // is interface{} with range: -2147483648..2147483647.
    MemorySize interface{}

    // sysdb name of sensor in the envmon EDM. The type is string with length:
    // 0..255.
    EnvironmentalMonitorPath interface{}

    // useful for storing an entity alias . The type is string with length:
    // 0..255.
    Alias interface{}

    // indicates if this entity is group       or not. The type is bool.
    GroupFlag interface{}

    // integer value for New Deviation Number 0x88. The type is interface{} with
    // range: -2147483648..2147483647.
    NewDeviationNumber interface{}

    // integer value for plim type if     applicable to this entity. The type is
    // interface{} with range: -2147483648..2147483647.
    PhysicalLayerInterfaceModuleType interface{}

    // 1 if UnrecognizedFRU and 0 for recognizedFRU. The type is bool.
    UnrecognizedFru interface{}

    // integer value for Redundancy State if     applicable to this entity. The
    // type is interface{} with range: -2147483648..2147483647.
    Redundancystate interface{}

    // 1 if ce port found, 0 if not. The type is bool.
    Ceport interface{}

    // 1 if xr scoped, 0 if not. The type is bool.
    XrScoped interface{}

    // Unique id for an entity. The type is interface{} with range:
    // -2147483648..2147483647.
    UniqueId interface{}
}

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_BasicInfo) GetEntityData() *types.CommonEntityData {
    basicInfo.EntityData.YFilter = basicInfo.YFilter
    basicInfo.EntityData.YangName = "basic-info"
    basicInfo.EntityData.BundleName = "cisco_ios_xr"
    basicInfo.EntityData.ParentYangName = "basic-attributes"
    basicInfo.EntityData.SegmentPath = "basic-info"
    basicInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicInfo.EntityData.Children = make(map[string]types.YChild)
    basicInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    basicInfo.EntityData.Leafs["description"] = types.YLeaf{"Description", basicInfo.Description}
    basicInfo.EntityData.Leafs["vendor-type"] = types.YLeaf{"VendorType", basicInfo.VendorType}
    basicInfo.EntityData.Leafs["name"] = types.YLeaf{"Name", basicInfo.Name}
    basicInfo.EntityData.Leafs["hardware-revision"] = types.YLeaf{"HardwareRevision", basicInfo.HardwareRevision}
    basicInfo.EntityData.Leafs["firmware-revision"] = types.YLeaf{"FirmwareRevision", basicInfo.FirmwareRevision}
    basicInfo.EntityData.Leafs["software-revision"] = types.YLeaf{"SoftwareRevision", basicInfo.SoftwareRevision}
    basicInfo.EntityData.Leafs["chip-hardware-revision"] = types.YLeaf{"ChipHardwareRevision", basicInfo.ChipHardwareRevision}
    basicInfo.EntityData.Leafs["serial-number"] = types.YLeaf{"SerialNumber", basicInfo.SerialNumber}
    basicInfo.EntityData.Leafs["manufacturer-name"] = types.YLeaf{"ManufacturerName", basicInfo.ManufacturerName}
    basicInfo.EntityData.Leafs["model-name"] = types.YLeaf{"ModelName", basicInfo.ModelName}
    basicInfo.EntityData.Leafs["asset-id-str"] = types.YLeaf{"AssetIdStr", basicInfo.AssetIdStr}
    basicInfo.EntityData.Leafs["asset-identification"] = types.YLeaf{"AssetIdentification", basicInfo.AssetIdentification}
    basicInfo.EntityData.Leafs["is-field-replaceable-unit"] = types.YLeaf{"IsFieldReplaceableUnit", basicInfo.IsFieldReplaceableUnit}
    basicInfo.EntityData.Leafs["manufacturer-asset-tags"] = types.YLeaf{"ManufacturerAssetTags", basicInfo.ManufacturerAssetTags}
    basicInfo.EntityData.Leafs["composite-class-code"] = types.YLeaf{"CompositeClassCode", basicInfo.CompositeClassCode}
    basicInfo.EntityData.Leafs["memory-size"] = types.YLeaf{"MemorySize", basicInfo.MemorySize}
    basicInfo.EntityData.Leafs["environmental-monitor-path"] = types.YLeaf{"EnvironmentalMonitorPath", basicInfo.EnvironmentalMonitorPath}
    basicInfo.EntityData.Leafs["alias"] = types.YLeaf{"Alias", basicInfo.Alias}
    basicInfo.EntityData.Leafs["group-flag"] = types.YLeaf{"GroupFlag", basicInfo.GroupFlag}
    basicInfo.EntityData.Leafs["new-deviation-number"] = types.YLeaf{"NewDeviationNumber", basicInfo.NewDeviationNumber}
    basicInfo.EntityData.Leafs["physical-layer-interface-module-type"] = types.YLeaf{"PhysicalLayerInterfaceModuleType", basicInfo.PhysicalLayerInterfaceModuleType}
    basicInfo.EntityData.Leafs["unrecognized-fru"] = types.YLeaf{"UnrecognizedFru", basicInfo.UnrecognizedFru}
    basicInfo.EntityData.Leafs["redundancystate"] = types.YLeaf{"Redundancystate", basicInfo.Redundancystate}
    basicInfo.EntityData.Leafs["ceport"] = types.YLeaf{"Ceport", basicInfo.Ceport}
    basicInfo.EntityData.Leafs["xr-scoped"] = types.YLeaf{"XrScoped", basicInfo.XrScoped}
    basicInfo.EntityData.Leafs["unique-id"] = types.YLeaf{"UniqueId", basicInfo.UniqueId}
    return &(basicInfo.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_FruInfo
// Field Replaceable Unit (FRU) inventory
// information
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_FruInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // card admin state: shutdown or not. The type is interface{} with range:
    // -2147483648..2147483647.
    CardAdministrativeState interface{}

    // power admin state: up or down. The type is interface{} with range:
    // -2147483648..2147483647.
    PowerAdministrativeState interface{}

    // card operation state. The type is interface{} with range:
    // -2147483648..2147483647.
    CardOperationalState interface{}

    // card is monitored by a manager or left unmonitored. The type is interface{}
    // with range: -2147483648..2147483647.
    CardMonitorState interface{}

    // card reset reason enum. The type is CardResetReason.
    CardResetReason interface{}

    // power current: not implemented. The type is interface{} with range:
    // -2147483648..2147483647.
    PowerCurrentMeasurement interface{}

    // Power operation state. The type is interface{} with range:
    // -2147483648..2147483647.
    PowerOperationalState interface{}

    // last card oper change state.
    LastOperationalStateChange Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_FruInfo_LastOperationalStateChange

    // card up time.
    CardUpTime Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_FruInfo_CardUpTime
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_FruInfo) GetEntityData() *types.CommonEntityData {
    fruInfo.EntityData.YFilter = fruInfo.YFilter
    fruInfo.EntityData.YangName = "fru-info"
    fruInfo.EntityData.BundleName = "cisco_ios_xr"
    fruInfo.EntityData.ParentYangName = "basic-attributes"
    fruInfo.EntityData.SegmentPath = "fru-info"
    fruInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fruInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fruInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fruInfo.EntityData.Children = make(map[string]types.YChild)
    fruInfo.EntityData.Children["last-operational-state-change"] = types.YChild{"LastOperationalStateChange", &fruInfo.LastOperationalStateChange}
    fruInfo.EntityData.Children["card-up-time"] = types.YChild{"CardUpTime", &fruInfo.CardUpTime}
    fruInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    fruInfo.EntityData.Leafs["card-administrative-state"] = types.YLeaf{"CardAdministrativeState", fruInfo.CardAdministrativeState}
    fruInfo.EntityData.Leafs["power-administrative-state"] = types.YLeaf{"PowerAdministrativeState", fruInfo.PowerAdministrativeState}
    fruInfo.EntityData.Leafs["card-operational-state"] = types.YLeaf{"CardOperationalState", fruInfo.CardOperationalState}
    fruInfo.EntityData.Leafs["card-monitor-state"] = types.YLeaf{"CardMonitorState", fruInfo.CardMonitorState}
    fruInfo.EntityData.Leafs["card-reset-reason"] = types.YLeaf{"CardResetReason", fruInfo.CardResetReason}
    fruInfo.EntityData.Leafs["power-current-measurement"] = types.YLeaf{"PowerCurrentMeasurement", fruInfo.PowerCurrentMeasurement}
    fruInfo.EntityData.Leafs["power-operational-state"] = types.YLeaf{"PowerOperationalState", fruInfo.PowerOperationalState}
    return &(fruInfo.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_FruInfo_LastOperationalStateChange
// last card oper change state
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_FruInfo_LastOperationalStateChange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_FruInfo_LastOperationalStateChange) GetEntityData() *types.CommonEntityData {
    lastOperationalStateChange.EntityData.YFilter = lastOperationalStateChange.YFilter
    lastOperationalStateChange.EntityData.YangName = "last-operational-state-change"
    lastOperationalStateChange.EntityData.BundleName = "cisco_ios_xr"
    lastOperationalStateChange.EntityData.ParentYangName = "fru-info"
    lastOperationalStateChange.EntityData.SegmentPath = "last-operational-state-change"
    lastOperationalStateChange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastOperationalStateChange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastOperationalStateChange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastOperationalStateChange.EntityData.Children = make(map[string]types.YChild)
    lastOperationalStateChange.EntityData.Leafs = make(map[string]types.YLeaf)
    lastOperationalStateChange.EntityData.Leafs["time-in-seconds"] = types.YLeaf{"TimeInSeconds", lastOperationalStateChange.TimeInSeconds}
    lastOperationalStateChange.EntityData.Leafs["time-in-nano-seconds"] = types.YLeaf{"TimeInNanoSeconds", lastOperationalStateChange.TimeInNanoSeconds}
    return &(lastOperationalStateChange.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_FruInfo_CardUpTime
// card up time
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_FruInfo_CardUpTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_FruInfo_CardUpTime) GetEntityData() *types.CommonEntityData {
    cardUpTime.EntityData.YFilter = cardUpTime.YFilter
    cardUpTime.EntityData.YangName = "card-up-time"
    cardUpTime.EntityData.BundleName = "cisco_ios_xr"
    cardUpTime.EntityData.ParentYangName = "fru-info"
    cardUpTime.EntityData.SegmentPath = "card-up-time"
    cardUpTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cardUpTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cardUpTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cardUpTime.EntityData.Children = make(map[string]types.YChild)
    cardUpTime.EntityData.Leafs = make(map[string]types.YLeaf)
    cardUpTime.EntityData.Leafs["time-in-seconds"] = types.YLeaf{"TimeInSeconds", cardUpTime.TimeInSeconds}
    cardUpTime.EntityData.Leafs["time-in-nano-seconds"] = types.YLeaf{"TimeInNanoSeconds", cardUpTime.TimeInNanoSeconds}
    return &(cardUpTime.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots
// PortSlotTable contains all optics ports in a
// SPA/PLIM.
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PortSlot number in the SPA/PLIM. The type is slice of
    // Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot.
    PortSlot []Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot
}

func (portSlots *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots) GetEntityData() *types.CommonEntityData {
    portSlots.EntityData.YFilter = portSlots.YFilter
    portSlots.EntityData.YangName = "port-slots"
    portSlots.EntityData.BundleName = "cisco_ios_xr"
    portSlots.EntityData.ParentYangName = "card"
    portSlots.EntityData.SegmentPath = "port-slots"
    portSlots.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portSlots.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portSlots.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portSlots.EntityData.Children = make(map[string]types.YChild)
    portSlots.EntityData.Children["port-slot"] = types.YChild{"PortSlot", nil}
    for i := range portSlots.PortSlot {
        portSlots.EntityData.Children[types.GetSegmentPath(&portSlots.PortSlot[i])] = types.YChild{"PortSlot", &portSlots.PortSlot[i]}
    }
    portSlots.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(portSlots.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot
// PortSlot number in the SPA/PLIM
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. portslot number. The type is interface{} with
    // range: -2147483648..2147483647.
    Number interface{}

    // Port string.
    Port Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port

    // Attributes.
    BasicAttributes Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes
}

func (portSlot *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot) GetEntityData() *types.CommonEntityData {
    portSlot.EntityData.YFilter = portSlot.YFilter
    portSlot.EntityData.YangName = "port-slot"
    portSlot.EntityData.BundleName = "cisco_ios_xr"
    portSlot.EntityData.ParentYangName = "port-slots"
    portSlot.EntityData.SegmentPath = "port-slot" + "[number='" + fmt.Sprintf("%v", portSlot.Number) + "']"
    portSlot.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portSlot.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portSlot.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portSlot.EntityData.Children = make(map[string]types.YChild)
    portSlot.EntityData.Children["port"] = types.YChild{"Port", &portSlot.Port}
    portSlot.EntityData.Children["basic-attributes"] = types.YChild{"BasicAttributes", &portSlot.BasicAttributes}
    portSlot.EntityData.Leafs = make(map[string]types.YLeaf)
    portSlot.EntityData.Leafs["number"] = types.YLeaf{"Number", portSlot.Number}
    return &(portSlot.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port
// Port string
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Attributes.
    BasicAttributes Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes
}

func (port *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port) GetEntityData() *types.CommonEntityData {
    port.EntityData.YFilter = port.YFilter
    port.EntityData.YangName = "port"
    port.EntityData.BundleName = "cisco_ios_xr"
    port.EntityData.ParentYangName = "port-slot"
    port.EntityData.SegmentPath = "port"
    port.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    port.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    port.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    port.EntityData.Children = make(map[string]types.YChild)
    port.EntityData.Children["basic-attributes"] = types.YChild{"BasicAttributes", &port.BasicAttributes}
    port.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(port.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes
// Attributes
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Inventory information.
    BasicInfo Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_BasicInfo

    // Field Replaceable Unit (FRU) inventory information.
    FruInfo Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_FruInfo
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes) GetEntityData() *types.CommonEntityData {
    basicAttributes.EntityData.YFilter = basicAttributes.YFilter
    basicAttributes.EntityData.YangName = "basic-attributes"
    basicAttributes.EntityData.BundleName = "cisco_ios_xr"
    basicAttributes.EntityData.ParentYangName = "port"
    basicAttributes.EntityData.SegmentPath = "basic-attributes"
    basicAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicAttributes.EntityData.Children = make(map[string]types.YChild)
    basicAttributes.EntityData.Children["basic-info"] = types.YChild{"BasicInfo", &basicAttributes.BasicInfo}
    basicAttributes.EntityData.Children["fru-info"] = types.YChild{"FruInfo", &basicAttributes.FruInfo}
    basicAttributes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(basicAttributes.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_BasicInfo
// Inventory information
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_BasicInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // describes in user-readable terms       what the entity in question does.
    // The type is string with length: 0..255.
    Description interface{}

    // maps to the vendor OID string. The type is string with length: 0..255.
    VendorType interface{}

    // name string for the entity. The type is string with length: 0..255.
    Name interface{}

    // hw revision string. The type is string with length: 0..255.
    HardwareRevision interface{}

    // firmware revision string. The type is string with length: 0..255.
    FirmwareRevision interface{}

    // software revision string. The type is string with length: 0..255.
    SoftwareRevision interface{}

    // chip module hw revision string. The type is string with length: 0..255.
    ChipHardwareRevision interface{}

    // serial number. The type is string with length: 0..255.
    SerialNumber interface{}

    // manufacturer's name. The type is string with length: 0..255.
    ManufacturerName interface{}

    // model name. The type is string with length: 0..255.
    ModelName interface{}

    // asset Identification string. The type is string with length: 0..255.
    AssetIdStr interface{}

    // asset Identification. The type is interface{} with range:
    // -2147483648..2147483647.
    AssetIdentification interface{}

    // 1 if Field Replaceable Unit 0, if not. The type is bool.
    IsFieldReplaceableUnit interface{}

    // Manufacture Asset Tags. The type is interface{} with range:
    // -2147483648..2147483647.
    ManufacturerAssetTags interface{}

    // Major&minor class of the entity. The type is interface{} with range:
    // -2147483648..2147483647.
    CompositeClassCode interface{}

    // Size of memory associated with       the entity where applicable. The type
    // is interface{} with range: -2147483648..2147483647.
    MemorySize interface{}

    // sysdb name of sensor in the envmon EDM. The type is string with length:
    // 0..255.
    EnvironmentalMonitorPath interface{}

    // useful for storing an entity alias . The type is string with length:
    // 0..255.
    Alias interface{}

    // indicates if this entity is group       or not. The type is bool.
    GroupFlag interface{}

    // integer value for New Deviation Number 0x88. The type is interface{} with
    // range: -2147483648..2147483647.
    NewDeviationNumber interface{}

    // integer value for plim type if     applicable to this entity. The type is
    // interface{} with range: -2147483648..2147483647.
    PhysicalLayerInterfaceModuleType interface{}

    // 1 if UnrecognizedFRU and 0 for recognizedFRU. The type is bool.
    UnrecognizedFru interface{}

    // integer value for Redundancy State if     applicable to this entity. The
    // type is interface{} with range: -2147483648..2147483647.
    Redundancystate interface{}

    // 1 if ce port found, 0 if not. The type is bool.
    Ceport interface{}

    // 1 if xr scoped, 0 if not. The type is bool.
    XrScoped interface{}

    // Unique id for an entity. The type is interface{} with range:
    // -2147483648..2147483647.
    UniqueId interface{}
}

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_BasicInfo) GetEntityData() *types.CommonEntityData {
    basicInfo.EntityData.YFilter = basicInfo.YFilter
    basicInfo.EntityData.YangName = "basic-info"
    basicInfo.EntityData.BundleName = "cisco_ios_xr"
    basicInfo.EntityData.ParentYangName = "basic-attributes"
    basicInfo.EntityData.SegmentPath = "basic-info"
    basicInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicInfo.EntityData.Children = make(map[string]types.YChild)
    basicInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    basicInfo.EntityData.Leafs["description"] = types.YLeaf{"Description", basicInfo.Description}
    basicInfo.EntityData.Leafs["vendor-type"] = types.YLeaf{"VendorType", basicInfo.VendorType}
    basicInfo.EntityData.Leafs["name"] = types.YLeaf{"Name", basicInfo.Name}
    basicInfo.EntityData.Leafs["hardware-revision"] = types.YLeaf{"HardwareRevision", basicInfo.HardwareRevision}
    basicInfo.EntityData.Leafs["firmware-revision"] = types.YLeaf{"FirmwareRevision", basicInfo.FirmwareRevision}
    basicInfo.EntityData.Leafs["software-revision"] = types.YLeaf{"SoftwareRevision", basicInfo.SoftwareRevision}
    basicInfo.EntityData.Leafs["chip-hardware-revision"] = types.YLeaf{"ChipHardwareRevision", basicInfo.ChipHardwareRevision}
    basicInfo.EntityData.Leafs["serial-number"] = types.YLeaf{"SerialNumber", basicInfo.SerialNumber}
    basicInfo.EntityData.Leafs["manufacturer-name"] = types.YLeaf{"ManufacturerName", basicInfo.ManufacturerName}
    basicInfo.EntityData.Leafs["model-name"] = types.YLeaf{"ModelName", basicInfo.ModelName}
    basicInfo.EntityData.Leafs["asset-id-str"] = types.YLeaf{"AssetIdStr", basicInfo.AssetIdStr}
    basicInfo.EntityData.Leafs["asset-identification"] = types.YLeaf{"AssetIdentification", basicInfo.AssetIdentification}
    basicInfo.EntityData.Leafs["is-field-replaceable-unit"] = types.YLeaf{"IsFieldReplaceableUnit", basicInfo.IsFieldReplaceableUnit}
    basicInfo.EntityData.Leafs["manufacturer-asset-tags"] = types.YLeaf{"ManufacturerAssetTags", basicInfo.ManufacturerAssetTags}
    basicInfo.EntityData.Leafs["composite-class-code"] = types.YLeaf{"CompositeClassCode", basicInfo.CompositeClassCode}
    basicInfo.EntityData.Leafs["memory-size"] = types.YLeaf{"MemorySize", basicInfo.MemorySize}
    basicInfo.EntityData.Leafs["environmental-monitor-path"] = types.YLeaf{"EnvironmentalMonitorPath", basicInfo.EnvironmentalMonitorPath}
    basicInfo.EntityData.Leafs["alias"] = types.YLeaf{"Alias", basicInfo.Alias}
    basicInfo.EntityData.Leafs["group-flag"] = types.YLeaf{"GroupFlag", basicInfo.GroupFlag}
    basicInfo.EntityData.Leafs["new-deviation-number"] = types.YLeaf{"NewDeviationNumber", basicInfo.NewDeviationNumber}
    basicInfo.EntityData.Leafs["physical-layer-interface-module-type"] = types.YLeaf{"PhysicalLayerInterfaceModuleType", basicInfo.PhysicalLayerInterfaceModuleType}
    basicInfo.EntityData.Leafs["unrecognized-fru"] = types.YLeaf{"UnrecognizedFru", basicInfo.UnrecognizedFru}
    basicInfo.EntityData.Leafs["redundancystate"] = types.YLeaf{"Redundancystate", basicInfo.Redundancystate}
    basicInfo.EntityData.Leafs["ceport"] = types.YLeaf{"Ceport", basicInfo.Ceport}
    basicInfo.EntityData.Leafs["xr-scoped"] = types.YLeaf{"XrScoped", basicInfo.XrScoped}
    basicInfo.EntityData.Leafs["unique-id"] = types.YLeaf{"UniqueId", basicInfo.UniqueId}
    return &(basicInfo.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_FruInfo
// Field Replaceable Unit (FRU) inventory
// information
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_FruInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // card admin state: shutdown or not. The type is interface{} with range:
    // -2147483648..2147483647.
    CardAdministrativeState interface{}

    // power admin state: up or down. The type is interface{} with range:
    // -2147483648..2147483647.
    PowerAdministrativeState interface{}

    // card operation state. The type is interface{} with range:
    // -2147483648..2147483647.
    CardOperationalState interface{}

    // card is monitored by a manager or left unmonitored. The type is interface{}
    // with range: -2147483648..2147483647.
    CardMonitorState interface{}

    // card reset reason enum. The type is CardResetReason.
    CardResetReason interface{}

    // power current: not implemented. The type is interface{} with range:
    // -2147483648..2147483647.
    PowerCurrentMeasurement interface{}

    // Power operation state. The type is interface{} with range:
    // -2147483648..2147483647.
    PowerOperationalState interface{}

    // last card oper change state.
    LastOperationalStateChange Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_LastOperationalStateChange

    // card up time.
    CardUpTime Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_CardUpTime
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_FruInfo) GetEntityData() *types.CommonEntityData {
    fruInfo.EntityData.YFilter = fruInfo.YFilter
    fruInfo.EntityData.YangName = "fru-info"
    fruInfo.EntityData.BundleName = "cisco_ios_xr"
    fruInfo.EntityData.ParentYangName = "basic-attributes"
    fruInfo.EntityData.SegmentPath = "fru-info"
    fruInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fruInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fruInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fruInfo.EntityData.Children = make(map[string]types.YChild)
    fruInfo.EntityData.Children["last-operational-state-change"] = types.YChild{"LastOperationalStateChange", &fruInfo.LastOperationalStateChange}
    fruInfo.EntityData.Children["card-up-time"] = types.YChild{"CardUpTime", &fruInfo.CardUpTime}
    fruInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    fruInfo.EntityData.Leafs["card-administrative-state"] = types.YLeaf{"CardAdministrativeState", fruInfo.CardAdministrativeState}
    fruInfo.EntityData.Leafs["power-administrative-state"] = types.YLeaf{"PowerAdministrativeState", fruInfo.PowerAdministrativeState}
    fruInfo.EntityData.Leafs["card-operational-state"] = types.YLeaf{"CardOperationalState", fruInfo.CardOperationalState}
    fruInfo.EntityData.Leafs["card-monitor-state"] = types.YLeaf{"CardMonitorState", fruInfo.CardMonitorState}
    fruInfo.EntityData.Leafs["card-reset-reason"] = types.YLeaf{"CardResetReason", fruInfo.CardResetReason}
    fruInfo.EntityData.Leafs["power-current-measurement"] = types.YLeaf{"PowerCurrentMeasurement", fruInfo.PowerCurrentMeasurement}
    fruInfo.EntityData.Leafs["power-operational-state"] = types.YLeaf{"PowerOperationalState", fruInfo.PowerOperationalState}
    return &(fruInfo.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_LastOperationalStateChange
// last card oper change state
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_LastOperationalStateChange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_LastOperationalStateChange) GetEntityData() *types.CommonEntityData {
    lastOperationalStateChange.EntityData.YFilter = lastOperationalStateChange.YFilter
    lastOperationalStateChange.EntityData.YangName = "last-operational-state-change"
    lastOperationalStateChange.EntityData.BundleName = "cisco_ios_xr"
    lastOperationalStateChange.EntityData.ParentYangName = "fru-info"
    lastOperationalStateChange.EntityData.SegmentPath = "last-operational-state-change"
    lastOperationalStateChange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastOperationalStateChange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastOperationalStateChange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastOperationalStateChange.EntityData.Children = make(map[string]types.YChild)
    lastOperationalStateChange.EntityData.Leafs = make(map[string]types.YLeaf)
    lastOperationalStateChange.EntityData.Leafs["time-in-seconds"] = types.YLeaf{"TimeInSeconds", lastOperationalStateChange.TimeInSeconds}
    lastOperationalStateChange.EntityData.Leafs["time-in-nano-seconds"] = types.YLeaf{"TimeInNanoSeconds", lastOperationalStateChange.TimeInNanoSeconds}
    return &(lastOperationalStateChange.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_CardUpTime
// card up time
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_CardUpTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_CardUpTime) GetEntityData() *types.CommonEntityData {
    cardUpTime.EntityData.YFilter = cardUpTime.YFilter
    cardUpTime.EntityData.YangName = "card-up-time"
    cardUpTime.EntityData.BundleName = "cisco_ios_xr"
    cardUpTime.EntityData.ParentYangName = "fru-info"
    cardUpTime.EntityData.SegmentPath = "card-up-time"
    cardUpTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cardUpTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cardUpTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cardUpTime.EntityData.Children = make(map[string]types.YChild)
    cardUpTime.EntityData.Leafs = make(map[string]types.YLeaf)
    cardUpTime.EntityData.Leafs["time-in-seconds"] = types.YLeaf{"TimeInSeconds", cardUpTime.TimeInSeconds}
    cardUpTime.EntityData.Leafs["time-in-nano-seconds"] = types.YLeaf{"TimeInNanoSeconds", cardUpTime.TimeInNanoSeconds}
    return &(cardUpTime.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes
// Attributes
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Inventory information.
    BasicInfo Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_BasicInfo

    // Field Replaceable Unit (FRU) inventory information.
    FruInfo Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_FruInfo
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes) GetEntityData() *types.CommonEntityData {
    basicAttributes.EntityData.YFilter = basicAttributes.YFilter
    basicAttributes.EntityData.YangName = "basic-attributes"
    basicAttributes.EntityData.BundleName = "cisco_ios_xr"
    basicAttributes.EntityData.ParentYangName = "port-slot"
    basicAttributes.EntityData.SegmentPath = "basic-attributes"
    basicAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicAttributes.EntityData.Children = make(map[string]types.YChild)
    basicAttributes.EntityData.Children["basic-info"] = types.YChild{"BasicInfo", &basicAttributes.BasicInfo}
    basicAttributes.EntityData.Children["fru-info"] = types.YChild{"FruInfo", &basicAttributes.FruInfo}
    basicAttributes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(basicAttributes.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_BasicInfo
// Inventory information
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_BasicInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // describes in user-readable terms       what the entity in question does.
    // The type is string with length: 0..255.
    Description interface{}

    // maps to the vendor OID string. The type is string with length: 0..255.
    VendorType interface{}

    // name string for the entity. The type is string with length: 0..255.
    Name interface{}

    // hw revision string. The type is string with length: 0..255.
    HardwareRevision interface{}

    // firmware revision string. The type is string with length: 0..255.
    FirmwareRevision interface{}

    // software revision string. The type is string with length: 0..255.
    SoftwareRevision interface{}

    // chip module hw revision string. The type is string with length: 0..255.
    ChipHardwareRevision interface{}

    // serial number. The type is string with length: 0..255.
    SerialNumber interface{}

    // manufacturer's name. The type is string with length: 0..255.
    ManufacturerName interface{}

    // model name. The type is string with length: 0..255.
    ModelName interface{}

    // asset Identification string. The type is string with length: 0..255.
    AssetIdStr interface{}

    // asset Identification. The type is interface{} with range:
    // -2147483648..2147483647.
    AssetIdentification interface{}

    // 1 if Field Replaceable Unit 0, if not. The type is bool.
    IsFieldReplaceableUnit interface{}

    // Manufacture Asset Tags. The type is interface{} with range:
    // -2147483648..2147483647.
    ManufacturerAssetTags interface{}

    // Major&minor class of the entity. The type is interface{} with range:
    // -2147483648..2147483647.
    CompositeClassCode interface{}

    // Size of memory associated with       the entity where applicable. The type
    // is interface{} with range: -2147483648..2147483647.
    MemorySize interface{}

    // sysdb name of sensor in the envmon EDM. The type is string with length:
    // 0..255.
    EnvironmentalMonitorPath interface{}

    // useful for storing an entity alias . The type is string with length:
    // 0..255.
    Alias interface{}

    // indicates if this entity is group       or not. The type is bool.
    GroupFlag interface{}

    // integer value for New Deviation Number 0x88. The type is interface{} with
    // range: -2147483648..2147483647.
    NewDeviationNumber interface{}

    // integer value for plim type if     applicable to this entity. The type is
    // interface{} with range: -2147483648..2147483647.
    PhysicalLayerInterfaceModuleType interface{}

    // 1 if UnrecognizedFRU and 0 for recognizedFRU. The type is bool.
    UnrecognizedFru interface{}

    // integer value for Redundancy State if     applicable to this entity. The
    // type is interface{} with range: -2147483648..2147483647.
    Redundancystate interface{}

    // 1 if ce port found, 0 if not. The type is bool.
    Ceport interface{}

    // 1 if xr scoped, 0 if not. The type is bool.
    XrScoped interface{}

    // Unique id for an entity. The type is interface{} with range:
    // -2147483648..2147483647.
    UniqueId interface{}
}

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_BasicInfo) GetEntityData() *types.CommonEntityData {
    basicInfo.EntityData.YFilter = basicInfo.YFilter
    basicInfo.EntityData.YangName = "basic-info"
    basicInfo.EntityData.BundleName = "cisco_ios_xr"
    basicInfo.EntityData.ParentYangName = "basic-attributes"
    basicInfo.EntityData.SegmentPath = "basic-info"
    basicInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicInfo.EntityData.Children = make(map[string]types.YChild)
    basicInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    basicInfo.EntityData.Leafs["description"] = types.YLeaf{"Description", basicInfo.Description}
    basicInfo.EntityData.Leafs["vendor-type"] = types.YLeaf{"VendorType", basicInfo.VendorType}
    basicInfo.EntityData.Leafs["name"] = types.YLeaf{"Name", basicInfo.Name}
    basicInfo.EntityData.Leafs["hardware-revision"] = types.YLeaf{"HardwareRevision", basicInfo.HardwareRevision}
    basicInfo.EntityData.Leafs["firmware-revision"] = types.YLeaf{"FirmwareRevision", basicInfo.FirmwareRevision}
    basicInfo.EntityData.Leafs["software-revision"] = types.YLeaf{"SoftwareRevision", basicInfo.SoftwareRevision}
    basicInfo.EntityData.Leafs["chip-hardware-revision"] = types.YLeaf{"ChipHardwareRevision", basicInfo.ChipHardwareRevision}
    basicInfo.EntityData.Leafs["serial-number"] = types.YLeaf{"SerialNumber", basicInfo.SerialNumber}
    basicInfo.EntityData.Leafs["manufacturer-name"] = types.YLeaf{"ManufacturerName", basicInfo.ManufacturerName}
    basicInfo.EntityData.Leafs["model-name"] = types.YLeaf{"ModelName", basicInfo.ModelName}
    basicInfo.EntityData.Leafs["asset-id-str"] = types.YLeaf{"AssetIdStr", basicInfo.AssetIdStr}
    basicInfo.EntityData.Leafs["asset-identification"] = types.YLeaf{"AssetIdentification", basicInfo.AssetIdentification}
    basicInfo.EntityData.Leafs["is-field-replaceable-unit"] = types.YLeaf{"IsFieldReplaceableUnit", basicInfo.IsFieldReplaceableUnit}
    basicInfo.EntityData.Leafs["manufacturer-asset-tags"] = types.YLeaf{"ManufacturerAssetTags", basicInfo.ManufacturerAssetTags}
    basicInfo.EntityData.Leafs["composite-class-code"] = types.YLeaf{"CompositeClassCode", basicInfo.CompositeClassCode}
    basicInfo.EntityData.Leafs["memory-size"] = types.YLeaf{"MemorySize", basicInfo.MemorySize}
    basicInfo.EntityData.Leafs["environmental-monitor-path"] = types.YLeaf{"EnvironmentalMonitorPath", basicInfo.EnvironmentalMonitorPath}
    basicInfo.EntityData.Leafs["alias"] = types.YLeaf{"Alias", basicInfo.Alias}
    basicInfo.EntityData.Leafs["group-flag"] = types.YLeaf{"GroupFlag", basicInfo.GroupFlag}
    basicInfo.EntityData.Leafs["new-deviation-number"] = types.YLeaf{"NewDeviationNumber", basicInfo.NewDeviationNumber}
    basicInfo.EntityData.Leafs["physical-layer-interface-module-type"] = types.YLeaf{"PhysicalLayerInterfaceModuleType", basicInfo.PhysicalLayerInterfaceModuleType}
    basicInfo.EntityData.Leafs["unrecognized-fru"] = types.YLeaf{"UnrecognizedFru", basicInfo.UnrecognizedFru}
    basicInfo.EntityData.Leafs["redundancystate"] = types.YLeaf{"Redundancystate", basicInfo.Redundancystate}
    basicInfo.EntityData.Leafs["ceport"] = types.YLeaf{"Ceport", basicInfo.Ceport}
    basicInfo.EntityData.Leafs["xr-scoped"] = types.YLeaf{"XrScoped", basicInfo.XrScoped}
    basicInfo.EntityData.Leafs["unique-id"] = types.YLeaf{"UniqueId", basicInfo.UniqueId}
    return &(basicInfo.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_FruInfo
// Field Replaceable Unit (FRU) inventory
// information
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_FruInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // card admin state: shutdown or not. The type is interface{} with range:
    // -2147483648..2147483647.
    CardAdministrativeState interface{}

    // power admin state: up or down. The type is interface{} with range:
    // -2147483648..2147483647.
    PowerAdministrativeState interface{}

    // card operation state. The type is interface{} with range:
    // -2147483648..2147483647.
    CardOperationalState interface{}

    // card is monitored by a manager or left unmonitored. The type is interface{}
    // with range: -2147483648..2147483647.
    CardMonitorState interface{}

    // card reset reason enum. The type is CardResetReason.
    CardResetReason interface{}

    // power current: not implemented. The type is interface{} with range:
    // -2147483648..2147483647.
    PowerCurrentMeasurement interface{}

    // Power operation state. The type is interface{} with range:
    // -2147483648..2147483647.
    PowerOperationalState interface{}

    // last card oper change state.
    LastOperationalStateChange Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_FruInfo_LastOperationalStateChange

    // card up time.
    CardUpTime Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_FruInfo_CardUpTime
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_FruInfo) GetEntityData() *types.CommonEntityData {
    fruInfo.EntityData.YFilter = fruInfo.YFilter
    fruInfo.EntityData.YangName = "fru-info"
    fruInfo.EntityData.BundleName = "cisco_ios_xr"
    fruInfo.EntityData.ParentYangName = "basic-attributes"
    fruInfo.EntityData.SegmentPath = "fru-info"
    fruInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fruInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fruInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fruInfo.EntityData.Children = make(map[string]types.YChild)
    fruInfo.EntityData.Children["last-operational-state-change"] = types.YChild{"LastOperationalStateChange", &fruInfo.LastOperationalStateChange}
    fruInfo.EntityData.Children["card-up-time"] = types.YChild{"CardUpTime", &fruInfo.CardUpTime}
    fruInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    fruInfo.EntityData.Leafs["card-administrative-state"] = types.YLeaf{"CardAdministrativeState", fruInfo.CardAdministrativeState}
    fruInfo.EntityData.Leafs["power-administrative-state"] = types.YLeaf{"PowerAdministrativeState", fruInfo.PowerAdministrativeState}
    fruInfo.EntityData.Leafs["card-operational-state"] = types.YLeaf{"CardOperationalState", fruInfo.CardOperationalState}
    fruInfo.EntityData.Leafs["card-monitor-state"] = types.YLeaf{"CardMonitorState", fruInfo.CardMonitorState}
    fruInfo.EntityData.Leafs["card-reset-reason"] = types.YLeaf{"CardResetReason", fruInfo.CardResetReason}
    fruInfo.EntityData.Leafs["power-current-measurement"] = types.YLeaf{"PowerCurrentMeasurement", fruInfo.PowerCurrentMeasurement}
    fruInfo.EntityData.Leafs["power-operational-state"] = types.YLeaf{"PowerOperationalState", fruInfo.PowerOperationalState}
    return &(fruInfo.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_FruInfo_LastOperationalStateChange
// last card oper change state
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_FruInfo_LastOperationalStateChange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_FruInfo_LastOperationalStateChange) GetEntityData() *types.CommonEntityData {
    lastOperationalStateChange.EntityData.YFilter = lastOperationalStateChange.YFilter
    lastOperationalStateChange.EntityData.YangName = "last-operational-state-change"
    lastOperationalStateChange.EntityData.BundleName = "cisco_ios_xr"
    lastOperationalStateChange.EntityData.ParentYangName = "fru-info"
    lastOperationalStateChange.EntityData.SegmentPath = "last-operational-state-change"
    lastOperationalStateChange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastOperationalStateChange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastOperationalStateChange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastOperationalStateChange.EntityData.Children = make(map[string]types.YChild)
    lastOperationalStateChange.EntityData.Leafs = make(map[string]types.YLeaf)
    lastOperationalStateChange.EntityData.Leafs["time-in-seconds"] = types.YLeaf{"TimeInSeconds", lastOperationalStateChange.TimeInSeconds}
    lastOperationalStateChange.EntityData.Leafs["time-in-nano-seconds"] = types.YLeaf{"TimeInNanoSeconds", lastOperationalStateChange.TimeInNanoSeconds}
    return &(lastOperationalStateChange.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_FruInfo_CardUpTime
// card up time
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_FruInfo_CardUpTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_FruInfo_CardUpTime) GetEntityData() *types.CommonEntityData {
    cardUpTime.EntityData.YFilter = cardUpTime.YFilter
    cardUpTime.EntityData.YangName = "card-up-time"
    cardUpTime.EntityData.BundleName = "cisco_ios_xr"
    cardUpTime.EntityData.ParentYangName = "fru-info"
    cardUpTime.EntityData.SegmentPath = "card-up-time"
    cardUpTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cardUpTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cardUpTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cardUpTime.EntityData.Children = make(map[string]types.YChild)
    cardUpTime.EntityData.Leafs = make(map[string]types.YLeaf)
    cardUpTime.EntityData.Leafs["time-in-seconds"] = types.YLeaf{"TimeInSeconds", cardUpTime.TimeInSeconds}
    cardUpTime.EntityData.Leafs["time-in-nano-seconds"] = types.YLeaf{"TimeInNanoSeconds", cardUpTime.TimeInNanoSeconds}
    return &(cardUpTime.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes
// Attributes
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Inventory information.
    BasicInfo Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_BasicInfo

    // Field Replaceable Unit (FRU) inventory information.
    FruInfo Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_FruInfo
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes) GetEntityData() *types.CommonEntityData {
    basicAttributes.EntityData.YFilter = basicAttributes.YFilter
    basicAttributes.EntityData.YangName = "basic-attributes"
    basicAttributes.EntityData.BundleName = "cisco_ios_xr"
    basicAttributes.EntityData.ParentYangName = "card"
    basicAttributes.EntityData.SegmentPath = "basic-attributes"
    basicAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicAttributes.EntityData.Children = make(map[string]types.YChild)
    basicAttributes.EntityData.Children["basic-info"] = types.YChild{"BasicInfo", &basicAttributes.BasicInfo}
    basicAttributes.EntityData.Children["fru-info"] = types.YChild{"FruInfo", &basicAttributes.FruInfo}
    basicAttributes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(basicAttributes.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_BasicInfo
// Inventory information
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_BasicInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // describes in user-readable terms       what the entity in question does.
    // The type is string with length: 0..255.
    Description interface{}

    // maps to the vendor OID string. The type is string with length: 0..255.
    VendorType interface{}

    // name string for the entity. The type is string with length: 0..255.
    Name interface{}

    // hw revision string. The type is string with length: 0..255.
    HardwareRevision interface{}

    // firmware revision string. The type is string with length: 0..255.
    FirmwareRevision interface{}

    // software revision string. The type is string with length: 0..255.
    SoftwareRevision interface{}

    // chip module hw revision string. The type is string with length: 0..255.
    ChipHardwareRevision interface{}

    // serial number. The type is string with length: 0..255.
    SerialNumber interface{}

    // manufacturer's name. The type is string with length: 0..255.
    ManufacturerName interface{}

    // model name. The type is string with length: 0..255.
    ModelName interface{}

    // asset Identification string. The type is string with length: 0..255.
    AssetIdStr interface{}

    // asset Identification. The type is interface{} with range:
    // -2147483648..2147483647.
    AssetIdentification interface{}

    // 1 if Field Replaceable Unit 0, if not. The type is bool.
    IsFieldReplaceableUnit interface{}

    // Manufacture Asset Tags. The type is interface{} with range:
    // -2147483648..2147483647.
    ManufacturerAssetTags interface{}

    // Major&minor class of the entity. The type is interface{} with range:
    // -2147483648..2147483647.
    CompositeClassCode interface{}

    // Size of memory associated with       the entity where applicable. The type
    // is interface{} with range: -2147483648..2147483647.
    MemorySize interface{}

    // sysdb name of sensor in the envmon EDM. The type is string with length:
    // 0..255.
    EnvironmentalMonitorPath interface{}

    // useful for storing an entity alias . The type is string with length:
    // 0..255.
    Alias interface{}

    // indicates if this entity is group       or not. The type is bool.
    GroupFlag interface{}

    // integer value for New Deviation Number 0x88. The type is interface{} with
    // range: -2147483648..2147483647.
    NewDeviationNumber interface{}

    // integer value for plim type if     applicable to this entity. The type is
    // interface{} with range: -2147483648..2147483647.
    PhysicalLayerInterfaceModuleType interface{}

    // 1 if UnrecognizedFRU and 0 for recognizedFRU. The type is bool.
    UnrecognizedFru interface{}

    // integer value for Redundancy State if     applicable to this entity. The
    // type is interface{} with range: -2147483648..2147483647.
    Redundancystate interface{}

    // 1 if ce port found, 0 if not. The type is bool.
    Ceport interface{}

    // 1 if xr scoped, 0 if not. The type is bool.
    XrScoped interface{}

    // Unique id for an entity. The type is interface{} with range:
    // -2147483648..2147483647.
    UniqueId interface{}
}

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_BasicInfo) GetEntityData() *types.CommonEntityData {
    basicInfo.EntityData.YFilter = basicInfo.YFilter
    basicInfo.EntityData.YangName = "basic-info"
    basicInfo.EntityData.BundleName = "cisco_ios_xr"
    basicInfo.EntityData.ParentYangName = "basic-attributes"
    basicInfo.EntityData.SegmentPath = "basic-info"
    basicInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicInfo.EntityData.Children = make(map[string]types.YChild)
    basicInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    basicInfo.EntityData.Leafs["description"] = types.YLeaf{"Description", basicInfo.Description}
    basicInfo.EntityData.Leafs["vendor-type"] = types.YLeaf{"VendorType", basicInfo.VendorType}
    basicInfo.EntityData.Leafs["name"] = types.YLeaf{"Name", basicInfo.Name}
    basicInfo.EntityData.Leafs["hardware-revision"] = types.YLeaf{"HardwareRevision", basicInfo.HardwareRevision}
    basicInfo.EntityData.Leafs["firmware-revision"] = types.YLeaf{"FirmwareRevision", basicInfo.FirmwareRevision}
    basicInfo.EntityData.Leafs["software-revision"] = types.YLeaf{"SoftwareRevision", basicInfo.SoftwareRevision}
    basicInfo.EntityData.Leafs["chip-hardware-revision"] = types.YLeaf{"ChipHardwareRevision", basicInfo.ChipHardwareRevision}
    basicInfo.EntityData.Leafs["serial-number"] = types.YLeaf{"SerialNumber", basicInfo.SerialNumber}
    basicInfo.EntityData.Leafs["manufacturer-name"] = types.YLeaf{"ManufacturerName", basicInfo.ManufacturerName}
    basicInfo.EntityData.Leafs["model-name"] = types.YLeaf{"ModelName", basicInfo.ModelName}
    basicInfo.EntityData.Leafs["asset-id-str"] = types.YLeaf{"AssetIdStr", basicInfo.AssetIdStr}
    basicInfo.EntityData.Leafs["asset-identification"] = types.YLeaf{"AssetIdentification", basicInfo.AssetIdentification}
    basicInfo.EntityData.Leafs["is-field-replaceable-unit"] = types.YLeaf{"IsFieldReplaceableUnit", basicInfo.IsFieldReplaceableUnit}
    basicInfo.EntityData.Leafs["manufacturer-asset-tags"] = types.YLeaf{"ManufacturerAssetTags", basicInfo.ManufacturerAssetTags}
    basicInfo.EntityData.Leafs["composite-class-code"] = types.YLeaf{"CompositeClassCode", basicInfo.CompositeClassCode}
    basicInfo.EntityData.Leafs["memory-size"] = types.YLeaf{"MemorySize", basicInfo.MemorySize}
    basicInfo.EntityData.Leafs["environmental-monitor-path"] = types.YLeaf{"EnvironmentalMonitorPath", basicInfo.EnvironmentalMonitorPath}
    basicInfo.EntityData.Leafs["alias"] = types.YLeaf{"Alias", basicInfo.Alias}
    basicInfo.EntityData.Leafs["group-flag"] = types.YLeaf{"GroupFlag", basicInfo.GroupFlag}
    basicInfo.EntityData.Leafs["new-deviation-number"] = types.YLeaf{"NewDeviationNumber", basicInfo.NewDeviationNumber}
    basicInfo.EntityData.Leafs["physical-layer-interface-module-type"] = types.YLeaf{"PhysicalLayerInterfaceModuleType", basicInfo.PhysicalLayerInterfaceModuleType}
    basicInfo.EntityData.Leafs["unrecognized-fru"] = types.YLeaf{"UnrecognizedFru", basicInfo.UnrecognizedFru}
    basicInfo.EntityData.Leafs["redundancystate"] = types.YLeaf{"Redundancystate", basicInfo.Redundancystate}
    basicInfo.EntityData.Leafs["ceport"] = types.YLeaf{"Ceport", basicInfo.Ceport}
    basicInfo.EntityData.Leafs["xr-scoped"] = types.YLeaf{"XrScoped", basicInfo.XrScoped}
    basicInfo.EntityData.Leafs["unique-id"] = types.YLeaf{"UniqueId", basicInfo.UniqueId}
    return &(basicInfo.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_FruInfo
// Field Replaceable Unit (FRU) inventory
// information
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_FruInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // card admin state: shutdown or not. The type is interface{} with range:
    // -2147483648..2147483647.
    CardAdministrativeState interface{}

    // power admin state: up or down. The type is interface{} with range:
    // -2147483648..2147483647.
    PowerAdministrativeState interface{}

    // card operation state. The type is interface{} with range:
    // -2147483648..2147483647.
    CardOperationalState interface{}

    // card is monitored by a manager or left unmonitored. The type is interface{}
    // with range: -2147483648..2147483647.
    CardMonitorState interface{}

    // card reset reason enum. The type is CardResetReason.
    CardResetReason interface{}

    // power current: not implemented. The type is interface{} with range:
    // -2147483648..2147483647.
    PowerCurrentMeasurement interface{}

    // Power operation state. The type is interface{} with range:
    // -2147483648..2147483647.
    PowerOperationalState interface{}

    // last card oper change state.
    LastOperationalStateChange Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_FruInfo_LastOperationalStateChange

    // card up time.
    CardUpTime Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_FruInfo_CardUpTime
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_FruInfo) GetEntityData() *types.CommonEntityData {
    fruInfo.EntityData.YFilter = fruInfo.YFilter
    fruInfo.EntityData.YangName = "fru-info"
    fruInfo.EntityData.BundleName = "cisco_ios_xr"
    fruInfo.EntityData.ParentYangName = "basic-attributes"
    fruInfo.EntityData.SegmentPath = "fru-info"
    fruInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fruInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fruInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fruInfo.EntityData.Children = make(map[string]types.YChild)
    fruInfo.EntityData.Children["last-operational-state-change"] = types.YChild{"LastOperationalStateChange", &fruInfo.LastOperationalStateChange}
    fruInfo.EntityData.Children["card-up-time"] = types.YChild{"CardUpTime", &fruInfo.CardUpTime}
    fruInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    fruInfo.EntityData.Leafs["card-administrative-state"] = types.YLeaf{"CardAdministrativeState", fruInfo.CardAdministrativeState}
    fruInfo.EntityData.Leafs["power-administrative-state"] = types.YLeaf{"PowerAdministrativeState", fruInfo.PowerAdministrativeState}
    fruInfo.EntityData.Leafs["card-operational-state"] = types.YLeaf{"CardOperationalState", fruInfo.CardOperationalState}
    fruInfo.EntityData.Leafs["card-monitor-state"] = types.YLeaf{"CardMonitorState", fruInfo.CardMonitorState}
    fruInfo.EntityData.Leafs["card-reset-reason"] = types.YLeaf{"CardResetReason", fruInfo.CardResetReason}
    fruInfo.EntityData.Leafs["power-current-measurement"] = types.YLeaf{"PowerCurrentMeasurement", fruInfo.PowerCurrentMeasurement}
    fruInfo.EntityData.Leafs["power-operational-state"] = types.YLeaf{"PowerOperationalState", fruInfo.PowerOperationalState}
    return &(fruInfo.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_FruInfo_LastOperationalStateChange
// last card oper change state
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_FruInfo_LastOperationalStateChange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_FruInfo_LastOperationalStateChange) GetEntityData() *types.CommonEntityData {
    lastOperationalStateChange.EntityData.YFilter = lastOperationalStateChange.YFilter
    lastOperationalStateChange.EntityData.YangName = "last-operational-state-change"
    lastOperationalStateChange.EntityData.BundleName = "cisco_ios_xr"
    lastOperationalStateChange.EntityData.ParentYangName = "fru-info"
    lastOperationalStateChange.EntityData.SegmentPath = "last-operational-state-change"
    lastOperationalStateChange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastOperationalStateChange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastOperationalStateChange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastOperationalStateChange.EntityData.Children = make(map[string]types.YChild)
    lastOperationalStateChange.EntityData.Leafs = make(map[string]types.YLeaf)
    lastOperationalStateChange.EntityData.Leafs["time-in-seconds"] = types.YLeaf{"TimeInSeconds", lastOperationalStateChange.TimeInSeconds}
    lastOperationalStateChange.EntityData.Leafs["time-in-nano-seconds"] = types.YLeaf{"TimeInNanoSeconds", lastOperationalStateChange.TimeInNanoSeconds}
    return &(lastOperationalStateChange.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_FruInfo_CardUpTime
// card up time
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_FruInfo_CardUpTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_FruInfo_CardUpTime) GetEntityData() *types.CommonEntityData {
    cardUpTime.EntityData.YFilter = cardUpTime.YFilter
    cardUpTime.EntityData.YangName = "card-up-time"
    cardUpTime.EntityData.BundleName = "cisco_ios_xr"
    cardUpTime.EntityData.ParentYangName = "fru-info"
    cardUpTime.EntityData.SegmentPath = "card-up-time"
    cardUpTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cardUpTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cardUpTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cardUpTime.EntityData.Children = make(map[string]types.YChild)
    cardUpTime.EntityData.Leafs = make(map[string]types.YLeaf)
    cardUpTime.EntityData.Leafs["time-in-seconds"] = types.YLeaf{"TimeInSeconds", cardUpTime.TimeInSeconds}
    cardUpTime.EntityData.Leafs["time-in-nano-seconds"] = types.YLeaf{"TimeInNanoSeconds", cardUpTime.TimeInNanoSeconds}
    return &(cardUpTime.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_BasicAttributes
// Attributes
type Inventory_Racks_Rack_Slots_Slot_BasicAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Inventory information.
    BasicInfo Inventory_Racks_Rack_Slots_Slot_BasicAttributes_BasicInfo

    // Field Replaceable Unit (FRU) inventory information.
    FruInfo Inventory_Racks_Rack_Slots_Slot_BasicAttributes_FruInfo
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_BasicAttributes) GetEntityData() *types.CommonEntityData {
    basicAttributes.EntityData.YFilter = basicAttributes.YFilter
    basicAttributes.EntityData.YangName = "basic-attributes"
    basicAttributes.EntityData.BundleName = "cisco_ios_xr"
    basicAttributes.EntityData.ParentYangName = "slot"
    basicAttributes.EntityData.SegmentPath = "basic-attributes"
    basicAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicAttributes.EntityData.Children = make(map[string]types.YChild)
    basicAttributes.EntityData.Children["basic-info"] = types.YChild{"BasicInfo", &basicAttributes.BasicInfo}
    basicAttributes.EntityData.Children["fru-info"] = types.YChild{"FruInfo", &basicAttributes.FruInfo}
    basicAttributes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(basicAttributes.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_BasicAttributes_BasicInfo
// Inventory information
type Inventory_Racks_Rack_Slots_Slot_BasicAttributes_BasicInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // describes in user-readable terms       what the entity in question does.
    // The type is string with length: 0..255.
    Description interface{}

    // maps to the vendor OID string. The type is string with length: 0..255.
    VendorType interface{}

    // name string for the entity. The type is string with length: 0..255.
    Name interface{}

    // hw revision string. The type is string with length: 0..255.
    HardwareRevision interface{}

    // firmware revision string. The type is string with length: 0..255.
    FirmwareRevision interface{}

    // software revision string. The type is string with length: 0..255.
    SoftwareRevision interface{}

    // chip module hw revision string. The type is string with length: 0..255.
    ChipHardwareRevision interface{}

    // serial number. The type is string with length: 0..255.
    SerialNumber interface{}

    // manufacturer's name. The type is string with length: 0..255.
    ManufacturerName interface{}

    // model name. The type is string with length: 0..255.
    ModelName interface{}

    // asset Identification string. The type is string with length: 0..255.
    AssetIdStr interface{}

    // asset Identification. The type is interface{} with range:
    // -2147483648..2147483647.
    AssetIdentification interface{}

    // 1 if Field Replaceable Unit 0, if not. The type is bool.
    IsFieldReplaceableUnit interface{}

    // Manufacture Asset Tags. The type is interface{} with range:
    // -2147483648..2147483647.
    ManufacturerAssetTags interface{}

    // Major&minor class of the entity. The type is interface{} with range:
    // -2147483648..2147483647.
    CompositeClassCode interface{}

    // Size of memory associated with       the entity where applicable. The type
    // is interface{} with range: -2147483648..2147483647.
    MemorySize interface{}

    // sysdb name of sensor in the envmon EDM. The type is string with length:
    // 0..255.
    EnvironmentalMonitorPath interface{}

    // useful for storing an entity alias . The type is string with length:
    // 0..255.
    Alias interface{}

    // indicates if this entity is group       or not. The type is bool.
    GroupFlag interface{}

    // integer value for New Deviation Number 0x88. The type is interface{} with
    // range: -2147483648..2147483647.
    NewDeviationNumber interface{}

    // integer value for plim type if     applicable to this entity. The type is
    // interface{} with range: -2147483648..2147483647.
    PhysicalLayerInterfaceModuleType interface{}

    // 1 if UnrecognizedFRU and 0 for recognizedFRU. The type is bool.
    UnrecognizedFru interface{}

    // integer value for Redundancy State if     applicable to this entity. The
    // type is interface{} with range: -2147483648..2147483647.
    Redundancystate interface{}

    // 1 if ce port found, 0 if not. The type is bool.
    Ceport interface{}

    // 1 if xr scoped, 0 if not. The type is bool.
    XrScoped interface{}

    // Unique id for an entity. The type is interface{} with range:
    // -2147483648..2147483647.
    UniqueId interface{}
}

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_BasicAttributes_BasicInfo) GetEntityData() *types.CommonEntityData {
    basicInfo.EntityData.YFilter = basicInfo.YFilter
    basicInfo.EntityData.YangName = "basic-info"
    basicInfo.EntityData.BundleName = "cisco_ios_xr"
    basicInfo.EntityData.ParentYangName = "basic-attributes"
    basicInfo.EntityData.SegmentPath = "basic-info"
    basicInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicInfo.EntityData.Children = make(map[string]types.YChild)
    basicInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    basicInfo.EntityData.Leafs["description"] = types.YLeaf{"Description", basicInfo.Description}
    basicInfo.EntityData.Leafs["vendor-type"] = types.YLeaf{"VendorType", basicInfo.VendorType}
    basicInfo.EntityData.Leafs["name"] = types.YLeaf{"Name", basicInfo.Name}
    basicInfo.EntityData.Leafs["hardware-revision"] = types.YLeaf{"HardwareRevision", basicInfo.HardwareRevision}
    basicInfo.EntityData.Leafs["firmware-revision"] = types.YLeaf{"FirmwareRevision", basicInfo.FirmwareRevision}
    basicInfo.EntityData.Leafs["software-revision"] = types.YLeaf{"SoftwareRevision", basicInfo.SoftwareRevision}
    basicInfo.EntityData.Leafs["chip-hardware-revision"] = types.YLeaf{"ChipHardwareRevision", basicInfo.ChipHardwareRevision}
    basicInfo.EntityData.Leafs["serial-number"] = types.YLeaf{"SerialNumber", basicInfo.SerialNumber}
    basicInfo.EntityData.Leafs["manufacturer-name"] = types.YLeaf{"ManufacturerName", basicInfo.ManufacturerName}
    basicInfo.EntityData.Leafs["model-name"] = types.YLeaf{"ModelName", basicInfo.ModelName}
    basicInfo.EntityData.Leafs["asset-id-str"] = types.YLeaf{"AssetIdStr", basicInfo.AssetIdStr}
    basicInfo.EntityData.Leafs["asset-identification"] = types.YLeaf{"AssetIdentification", basicInfo.AssetIdentification}
    basicInfo.EntityData.Leafs["is-field-replaceable-unit"] = types.YLeaf{"IsFieldReplaceableUnit", basicInfo.IsFieldReplaceableUnit}
    basicInfo.EntityData.Leafs["manufacturer-asset-tags"] = types.YLeaf{"ManufacturerAssetTags", basicInfo.ManufacturerAssetTags}
    basicInfo.EntityData.Leafs["composite-class-code"] = types.YLeaf{"CompositeClassCode", basicInfo.CompositeClassCode}
    basicInfo.EntityData.Leafs["memory-size"] = types.YLeaf{"MemorySize", basicInfo.MemorySize}
    basicInfo.EntityData.Leafs["environmental-monitor-path"] = types.YLeaf{"EnvironmentalMonitorPath", basicInfo.EnvironmentalMonitorPath}
    basicInfo.EntityData.Leafs["alias"] = types.YLeaf{"Alias", basicInfo.Alias}
    basicInfo.EntityData.Leafs["group-flag"] = types.YLeaf{"GroupFlag", basicInfo.GroupFlag}
    basicInfo.EntityData.Leafs["new-deviation-number"] = types.YLeaf{"NewDeviationNumber", basicInfo.NewDeviationNumber}
    basicInfo.EntityData.Leafs["physical-layer-interface-module-type"] = types.YLeaf{"PhysicalLayerInterfaceModuleType", basicInfo.PhysicalLayerInterfaceModuleType}
    basicInfo.EntityData.Leafs["unrecognized-fru"] = types.YLeaf{"UnrecognizedFru", basicInfo.UnrecognizedFru}
    basicInfo.EntityData.Leafs["redundancystate"] = types.YLeaf{"Redundancystate", basicInfo.Redundancystate}
    basicInfo.EntityData.Leafs["ceport"] = types.YLeaf{"Ceport", basicInfo.Ceport}
    basicInfo.EntityData.Leafs["xr-scoped"] = types.YLeaf{"XrScoped", basicInfo.XrScoped}
    basicInfo.EntityData.Leafs["unique-id"] = types.YLeaf{"UniqueId", basicInfo.UniqueId}
    return &(basicInfo.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_BasicAttributes_FruInfo
// Field Replaceable Unit (FRU) inventory
// information
type Inventory_Racks_Rack_Slots_Slot_BasicAttributes_FruInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // card admin state: shutdown or not. The type is interface{} with range:
    // -2147483648..2147483647.
    CardAdministrativeState interface{}

    // power admin state: up or down. The type is interface{} with range:
    // -2147483648..2147483647.
    PowerAdministrativeState interface{}

    // card operation state. The type is interface{} with range:
    // -2147483648..2147483647.
    CardOperationalState interface{}

    // card is monitored by a manager or left unmonitored. The type is interface{}
    // with range: -2147483648..2147483647.
    CardMonitorState interface{}

    // card reset reason enum. The type is CardResetReason.
    CardResetReason interface{}

    // power current: not implemented. The type is interface{} with range:
    // -2147483648..2147483647.
    PowerCurrentMeasurement interface{}

    // Power operation state. The type is interface{} with range:
    // -2147483648..2147483647.
    PowerOperationalState interface{}

    // last card oper change state.
    LastOperationalStateChange Inventory_Racks_Rack_Slots_Slot_BasicAttributes_FruInfo_LastOperationalStateChange

    // card up time.
    CardUpTime Inventory_Racks_Rack_Slots_Slot_BasicAttributes_FruInfo_CardUpTime
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_BasicAttributes_FruInfo) GetEntityData() *types.CommonEntityData {
    fruInfo.EntityData.YFilter = fruInfo.YFilter
    fruInfo.EntityData.YangName = "fru-info"
    fruInfo.EntityData.BundleName = "cisco_ios_xr"
    fruInfo.EntityData.ParentYangName = "basic-attributes"
    fruInfo.EntityData.SegmentPath = "fru-info"
    fruInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fruInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fruInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fruInfo.EntityData.Children = make(map[string]types.YChild)
    fruInfo.EntityData.Children["last-operational-state-change"] = types.YChild{"LastOperationalStateChange", &fruInfo.LastOperationalStateChange}
    fruInfo.EntityData.Children["card-up-time"] = types.YChild{"CardUpTime", &fruInfo.CardUpTime}
    fruInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    fruInfo.EntityData.Leafs["card-administrative-state"] = types.YLeaf{"CardAdministrativeState", fruInfo.CardAdministrativeState}
    fruInfo.EntityData.Leafs["power-administrative-state"] = types.YLeaf{"PowerAdministrativeState", fruInfo.PowerAdministrativeState}
    fruInfo.EntityData.Leafs["card-operational-state"] = types.YLeaf{"CardOperationalState", fruInfo.CardOperationalState}
    fruInfo.EntityData.Leafs["card-monitor-state"] = types.YLeaf{"CardMonitorState", fruInfo.CardMonitorState}
    fruInfo.EntityData.Leafs["card-reset-reason"] = types.YLeaf{"CardResetReason", fruInfo.CardResetReason}
    fruInfo.EntityData.Leafs["power-current-measurement"] = types.YLeaf{"PowerCurrentMeasurement", fruInfo.PowerCurrentMeasurement}
    fruInfo.EntityData.Leafs["power-operational-state"] = types.YLeaf{"PowerOperationalState", fruInfo.PowerOperationalState}
    return &(fruInfo.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_BasicAttributes_FruInfo_LastOperationalStateChange
// last card oper change state
type Inventory_Racks_Rack_Slots_Slot_BasicAttributes_FruInfo_LastOperationalStateChange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_BasicAttributes_FruInfo_LastOperationalStateChange) GetEntityData() *types.CommonEntityData {
    lastOperationalStateChange.EntityData.YFilter = lastOperationalStateChange.YFilter
    lastOperationalStateChange.EntityData.YangName = "last-operational-state-change"
    lastOperationalStateChange.EntityData.BundleName = "cisco_ios_xr"
    lastOperationalStateChange.EntityData.ParentYangName = "fru-info"
    lastOperationalStateChange.EntityData.SegmentPath = "last-operational-state-change"
    lastOperationalStateChange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastOperationalStateChange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastOperationalStateChange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastOperationalStateChange.EntityData.Children = make(map[string]types.YChild)
    lastOperationalStateChange.EntityData.Leafs = make(map[string]types.YLeaf)
    lastOperationalStateChange.EntityData.Leafs["time-in-seconds"] = types.YLeaf{"TimeInSeconds", lastOperationalStateChange.TimeInSeconds}
    lastOperationalStateChange.EntityData.Leafs["time-in-nano-seconds"] = types.YLeaf{"TimeInNanoSeconds", lastOperationalStateChange.TimeInNanoSeconds}
    return &(lastOperationalStateChange.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_BasicAttributes_FruInfo_CardUpTime
// card up time
type Inventory_Racks_Rack_Slots_Slot_BasicAttributes_FruInfo_CardUpTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_BasicAttributes_FruInfo_CardUpTime) GetEntityData() *types.CommonEntityData {
    cardUpTime.EntityData.YFilter = cardUpTime.YFilter
    cardUpTime.EntityData.YangName = "card-up-time"
    cardUpTime.EntityData.BundleName = "cisco_ios_xr"
    cardUpTime.EntityData.ParentYangName = "fru-info"
    cardUpTime.EntityData.SegmentPath = "card-up-time"
    cardUpTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cardUpTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cardUpTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cardUpTime.EntityData.Children = make(map[string]types.YChild)
    cardUpTime.EntityData.Leafs = make(map[string]types.YLeaf)
    cardUpTime.EntityData.Leafs["time-in-seconds"] = types.YLeaf{"TimeInSeconds", cardUpTime.TimeInSeconds}
    cardUpTime.EntityData.Leafs["time-in-nano-seconds"] = types.YLeaf{"TimeInNanoSeconds", cardUpTime.TimeInNanoSeconds}
    return &(cardUpTime.EntityData)
}

