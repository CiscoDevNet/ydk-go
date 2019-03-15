// This module contains a collection of YANG definitions
// for Cisco IOS-XR asr9k-sc-invmgr package operational data.
// 
// This module contains definitions
// for the following management objects:
//   inventory: Logical Router Inventory operational data
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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

// InvResetReason represents Inv reset reason
type InvResetReason string

const (
    // module reset reason unknown
    InvResetReason_module_reset_reason_unknown InvResetReason = "module-reset-reason-unknown"

    // module reset reason powerup
    InvResetReason_module_reset_reason_powerup InvResetReason = "module-reset-reason-powerup"

    // module reset reason user shutdown
    InvResetReason_module_reset_reason_user_shutdown InvResetReason = "module-reset-reason-user-shutdown"

    // module reset reason user reload
    InvResetReason_module_reset_reason_user_reload InvResetReason = "module-reset-reason-user-reload"

    // module reset reason auto reload
    InvResetReason_module_reset_reason_auto_reload InvResetReason = "module-reset-reason-auto-reload"

    // module reset reason environment
    InvResetReason_module_reset_reason_environment InvResetReason = "module-reset-reason-environment"

    // module reset reason user unpower
    InvResetReason_module_reset_reason_user_unpower InvResetReason = "module-reset-reason-user-unpower"
)

// InvAdminState represents Inv admin state
type InvAdminState string

const (
    // admin state invalid
    InvAdminState_admin_state_invalid InvAdminState = "admin-state-invalid"

    // admin up
    InvAdminState_admin_up InvAdminState = "admin-up"

    // admin down
    InvAdminState_admin_down InvAdminState = "admin-down"
)

// InvMonitorState represents Inv monitor state
type InvMonitorState string

const (
    // unmonitored
    InvMonitorState_unmonitored InvMonitorState = "unmonitored"

    // monitored
    InvMonitorState_monitored InvMonitorState = "monitored"
)

// InvPowerAdminState represents Inv power admin state
type InvPowerAdminState string

const (
    // admin power invalid
    InvPowerAdminState_admin_power_invalid InvPowerAdminState = "admin-power-invalid"

    // admin on
    InvPowerAdminState_admin_on InvPowerAdminState = "admin-on"

    // admin off
    InvPowerAdminState_admin_off InvPowerAdminState = "admin-off"
)

// InvCardState represents Inv card state
type InvCardState string

const (
    // inv card not present
    InvCardState_inv_card_not_present InvCardState = "inv-card-not-present"

    // inv card present
    InvCardState_inv_card_present InvCardState = "inv-card-present"

    // inv card reset
    InvCardState_inv_card_reset InvCardState = "inv-card-reset"

    // inv card booting
    InvCardState_inv_card_booting InvCardState = "inv-card-booting"

    // inv card mbi booting
    InvCardState_inv_card_mbi_booting InvCardState = "inv-card-mbi-booting"

    // inv card running mbi
    InvCardState_inv_card_running_mbi InvCardState = "inv-card-running-mbi"

    // inv card running ena
    InvCardState_inv_card_running_ena InvCardState = "inv-card-running-ena"

    // inv card bring down
    InvCardState_inv_card_bring_down InvCardState = "inv-card-bring-down"

    // inv card ena failure
    InvCardState_inv_card_ena_failure InvCardState = "inv-card-ena-failure"

    // inv card f diag run
    InvCardState_inv_card_f_diag_run InvCardState = "inv-card-f-diag-run"

    // inv card f diag failure
    InvCardState_inv_card_f_diag_failure InvCardState = "inv-card-f-diag-failure"

    // inv card powered
    InvCardState_inv_card_powered InvCardState = "inv-card-powered"

    // inv card unpowered
    InvCardState_inv_card_unpowered InvCardState = "inv-card-unpowered"

    // inv card mdr
    InvCardState_inv_card_mdr InvCardState = "inv-card-mdr"

    // inv card mdr running mbi
    InvCardState_inv_card_mdr_running_mbi InvCardState = "inv-card-mdr-running-mbi"

    // inv card main t mode
    InvCardState_inv_card_main_t_mode InvCardState = "inv-card-main-t-mode"

    // inv card admin down
    InvCardState_inv_card_admin_down InvCardState = "inv-card-admin-down"

    // inv card no mon
    InvCardState_inv_card_no_mon InvCardState = "inv-card-no-mon"

    // inv card unknown
    InvCardState_inv_card_unknown InvCardState = "inv-card-unknown"

    // inv card failed
    InvCardState_inv_card_failed InvCardState = "inv-card-failed"

    // inv card ok
    InvCardState_inv_card_ok InvCardState = "inv-card-ok"

    // inv card missing
    InvCardState_inv_card_missing InvCardState = "inv-card-missing"

    // inv card field diag downloading
    InvCardState_inv_card_field_diag_downloading InvCardState = "inv-card-field-diag-downloading"

    // inv card field diag unmonitor
    InvCardState_inv_card_field_diag_unmonitor InvCardState = "inv-card-field-diag-unmonitor"

    // inv card fabric field diag unmonitor
    InvCardState_inv_card_fabric_field_diag_unmonitor InvCardState = "inv-card-fabric-field-diag-unmonitor"

    // inv card field diag rp launching
    InvCardState_inv_card_field_diag_rp_launching InvCardState = "inv-card-field-diag-rp-launching"

    // inv card field diag running
    InvCardState_inv_card_field_diag_running InvCardState = "inv-card-field-diag-running"

    // inv card field diag pass
    InvCardState_inv_card_field_diag_pass InvCardState = "inv-card-field-diag-pass"

    // inv card field diag fail
    InvCardState_inv_card_field_diag_fail InvCardState = "inv-card-field-diag-fail"

    // inv card field diag timeout
    InvCardState_inv_card_field_diag_timeout InvCardState = "inv-card-field-diag-timeout"

    // inv card disabled
    InvCardState_inv_card_disabled InvCardState = "inv-card-disabled"

    // inv card spa booting
    InvCardState_inv_card_spa_booting InvCardState = "inv-card-spa-booting"

    // inv card not allowed online
    InvCardState_inv_card_not_allowed_online InvCardState = "inv-card-not-allowed-online"

    // inv card stopped
    InvCardState_inv_card_stopped InvCardState = "inv-card-stopped"

    // inv card incompatible fw ver
    InvCardState_inv_card_incompatible_fw_ver InvCardState = "inv-card-incompatible-fw-ver"

    // inv card fpd hold
    InvCardState_inv_card_fpd_hold InvCardState = "inv-card-fpd-hold"

    // inv card node prep
    InvCardState_inv_card_node_prep InvCardState = "inv-card-node-prep"

    // inv card updating fpd
    InvCardState_inv_card_updating_fpd InvCardState = "inv-card-updating-fpd"

    // inv card num states
    InvCardState_inv_card_num_states InvCardState = "inv-card-num-states"
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
    inventory.EntityData.AbsolutePath = inventory.EntityData.SegmentPath
    inventory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inventory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inventory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inventory.EntityData.Children = types.NewOrderedMap()
    inventory.EntityData.Children.Append("racks", types.YChild{"Racks", &inventory.Racks})
    inventory.EntityData.Leafs = types.NewOrderedMap()

    inventory.EntityData.YListKeys = []string {}

    return &(inventory.EntityData)
}

// Inventory_Racks
// Table of racks
type Inventory_Racks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Rack number. The type is slice of Inventory_Racks_Rack.
    Rack []*Inventory_Racks_Rack
}

func (racks *Inventory_Racks) GetEntityData() *types.CommonEntityData {
    racks.EntityData.YFilter = racks.YFilter
    racks.EntityData.YangName = "racks"
    racks.EntityData.BundleName = "cisco_ios_xr"
    racks.EntityData.ParentYangName = "inventory"
    racks.EntityData.SegmentPath = "racks"
    racks.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/" + racks.EntityData.SegmentPath
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

// Inventory_Racks_Rack
// Rack number
type Inventory_Racks_Rack struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Rack number. The type is interface{} with range:
    // 0..4294967295.
    Number interface{}

    // Slot table contains all slots in the rack.
    Slots Inventory_Racks_Rack_Slots
}

func (rack *Inventory_Racks_Rack) GetEntityData() *types.CommonEntityData {
    rack.EntityData.YFilter = rack.YFilter
    rack.EntityData.YangName = "rack"
    rack.EntityData.BundleName = "cisco_ios_xr"
    rack.EntityData.ParentYangName = "racks"
    rack.EntityData.SegmentPath = "rack" + types.AddKeyToken(rack.Number, "number")
    rack.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/" + rack.EntityData.SegmentPath
    rack.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rack.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rack.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rack.EntityData.Children = types.NewOrderedMap()
    rack.EntityData.Children.Append("slots", types.YChild{"Slots", &rack.Slots})
    rack.EntityData.Leafs = types.NewOrderedMap()
    rack.EntityData.Leafs.Append("number", types.YLeaf{"Number", rack.Number})

    rack.EntityData.YListKeys = []string {"Number"}

    return &(rack.EntityData)
}

// Inventory_Racks_Rack_Slots
// Slot table contains all slots in the rack
type Inventory_Racks_Rack_Slots struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Slot number. The type is slice of Inventory_Racks_Rack_Slots_Slot.
    Slot []*Inventory_Racks_Rack_Slots_Slot
}

func (slots *Inventory_Racks_Rack_Slots) GetEntityData() *types.CommonEntityData {
    slots.EntityData.YFilter = slots.YFilter
    slots.EntityData.YangName = "slots"
    slots.EntityData.BundleName = "cisco_ios_xr"
    slots.EntityData.ParentYangName = "rack"
    slots.EntityData.SegmentPath = "slots"
    slots.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/" + slots.EntityData.SegmentPath
    slots.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    slots.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    slots.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    slots.EntityData.Children = types.NewOrderedMap()
    slots.EntityData.Children.Append("slot", types.YChild{"Slot", nil})
    for i := range slots.Slot {
        slots.EntityData.Children.Append(types.GetSegmentPath(slots.Slot[i]), types.YChild{"Slot", slots.Slot[i]})
    }
    slots.EntityData.Leafs = types.NewOrderedMap()

    slots.EntityData.YListKeys = []string {}

    return &(slots.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot
// Slot number
type Inventory_Racks_Rack_Slots_Slot struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Slot number. The type is interface{} with range:
    // 0..4294967295.
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
    slot.EntityData.SegmentPath = "slot" + types.AddKeyToken(slot.Number, "number")
    slot.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/" + slot.EntityData.SegmentPath
    slot.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    slot.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    slot.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    slot.EntityData.Children = types.NewOrderedMap()
    slot.EntityData.Children.Append("cards", types.YChild{"Cards", &slot.Cards})
    slot.EntityData.Children.Append("basic-attributes", types.YChild{"BasicAttributes", &slot.BasicAttributes})
    slot.EntityData.Leafs = types.NewOrderedMap()
    slot.EntityData.Leafs.Append("number", types.YLeaf{"Number", slot.Number})

    slot.EntityData.YListKeys = []string {"Number"}

    return &(slot.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards
// Card table contains all cards in the slot
type Inventory_Racks_Rack_Slots_Slot_Cards struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Card number. The type is slice of
    // Inventory_Racks_Rack_Slots_Slot_Cards_Card.
    Card []*Inventory_Racks_Rack_Slots_Slot_Cards_Card
}

func (cards *Inventory_Racks_Rack_Slots_Slot_Cards) GetEntityData() *types.CommonEntityData {
    cards.EntityData.YFilter = cards.YFilter
    cards.EntityData.YangName = "cards"
    cards.EntityData.BundleName = "cisco_ios_xr"
    cards.EntityData.ParentYangName = "slot"
    cards.EntityData.SegmentPath = "cards"
    cards.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/" + cards.EntityData.SegmentPath
    cards.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cards.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cards.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cards.EntityData.Children = types.NewOrderedMap()
    cards.EntityData.Children.Append("card", types.YChild{"Card", nil})
    for i := range cards.Card {
        cards.EntityData.Children.Append(types.GetSegmentPath(cards.Card[i]), types.YChild{"Card", cards.Card[i]})
    }
    cards.EntityData.Leafs = types.NewOrderedMap()

    cards.EntityData.YListKeys = []string {}

    return &(cards.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card
// Card number
type Inventory_Racks_Rack_Slots_Slot_Cards_Card struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. card number. The type is interface{} with range:
    // 0..4294967295.
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
    card.EntityData.SegmentPath = "card" + types.AddKeyToken(card.Number, "number")
    card.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/cards/" + card.EntityData.SegmentPath
    card.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    card.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    card.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    card.EntityData.Children = types.NewOrderedMap()
    card.EntityData.Children.Append("sub-slots", types.YChild{"SubSlots", &card.SubSlots})
    card.EntityData.Children.Append("hw-components", types.YChild{"HwComponents", &card.HwComponents})
    card.EntityData.Children.Append("sensors", types.YChild{"Sensors", &card.Sensors})
    card.EntityData.Children.Append("port-slots", types.YChild{"PortSlots", &card.PortSlots})
    card.EntityData.Children.Append("basic-attributes", types.YChild{"BasicAttributes", &card.BasicAttributes})
    card.EntityData.Leafs = types.NewOrderedMap()
    card.EntityData.Leafs.Append("number", types.YLeaf{"Number", card.Number})

    card.EntityData.YListKeys = []string {"Number"}

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
    SubSlot []*Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot
}

func (subSlots *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots) GetEntityData() *types.CommonEntityData {
    subSlots.EntityData.YFilter = subSlots.YFilter
    subSlots.EntityData.YangName = "sub-slots"
    subSlots.EntityData.BundleName = "cisco_ios_xr"
    subSlots.EntityData.ParentYangName = "card"
    subSlots.EntityData.SegmentPath = "sub-slots"
    subSlots.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/cards/card/" + subSlots.EntityData.SegmentPath
    subSlots.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    subSlots.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    subSlots.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    subSlots.EntityData.Children = types.NewOrderedMap()
    subSlots.EntityData.Children.Append("sub-slot", types.YChild{"SubSlot", nil})
    for i := range subSlots.SubSlot {
        subSlots.EntityData.Children.Append(types.GetSegmentPath(subSlots.SubSlot[i]), types.YChild{"SubSlot", subSlots.SubSlot[i]})
    }
    subSlots.EntityData.Leafs = types.NewOrderedMap()

    subSlots.EntityData.YListKeys = []string {}

    return &(subSlots.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot
// SubSlot number
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. subslot number. The type is interface{} with
    // range: 0..4294967295.
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
    subSlot.EntityData.SegmentPath = "sub-slot" + types.AddKeyToken(subSlot.Number, "number")
    subSlot.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/cards/card/sub-slots/" + subSlot.EntityData.SegmentPath
    subSlot.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    subSlot.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    subSlot.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    subSlot.EntityData.Children = types.NewOrderedMap()
    subSlot.EntityData.Children.Append("module", types.YChild{"Module", &subSlot.Module})
    subSlot.EntityData.Children.Append("basic-attributes", types.YChild{"BasicAttributes", &subSlot.BasicAttributes})
    subSlot.EntityData.Leafs = types.NewOrderedMap()
    subSlot.EntityData.Leafs.Append("number", types.YLeaf{"Number", subSlot.Number})

    subSlot.EntityData.YListKeys = []string {"Number"}

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
    module.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/" + module.EntityData.SegmentPath
    module.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    module.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    module.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    module.EntityData.Children = types.NewOrderedMap()
    module.EntityData.Children.Append("sensors", types.YChild{"Sensors", &module.Sensors})
    module.EntityData.Children.Append("port-slots", types.YChild{"PortSlots", &module.PortSlots})
    module.EntityData.Children.Append("basic-attributes", types.YChild{"BasicAttributes", &module.BasicAttributes})
    module.EntityData.Leafs = types.NewOrderedMap()

    module.EntityData.YListKeys = []string {}

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
    Sensor []*Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor
}

func (sensors *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors) GetEntityData() *types.CommonEntityData {
    sensors.EntityData.YFilter = sensors.YFilter
    sensors.EntityData.YangName = "sensors"
    sensors.EntityData.BundleName = "cisco_ios_xr"
    sensors.EntityData.ParentYangName = "module"
    sensors.EntityData.SegmentPath = "sensors"
    sensors.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/" + sensors.EntityData.SegmentPath
    sensors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sensors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sensors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sensors.EntityData.Children = types.NewOrderedMap()
    sensors.EntityData.Children.Append("sensor", types.YChild{"Sensor", nil})
    for i := range sensors.Sensor {
        sensors.EntityData.Children.Append(types.GetSegmentPath(sensors.Sensor[i]), types.YChild{"Sensor", sensors.Sensor[i]})
    }
    sensors.EntityData.Leafs = types.NewOrderedMap()

    sensors.EntityData.YListKeys = []string {}

    return &(sensors.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor
// Sensor number in the Module
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. sensor number. The type is interface{} with range:
    // 0..4294967295.
    Number interface{}

    // Attributes.
    BasicAttributes Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes
}

func (sensor *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor) GetEntityData() *types.CommonEntityData {
    sensor.EntityData.YFilter = sensor.YFilter
    sensor.EntityData.YangName = "sensor"
    sensor.EntityData.BundleName = "cisco_ios_xr"
    sensor.EntityData.ParentYangName = "sensors"
    sensor.EntityData.SegmentPath = "sensor" + types.AddKeyToken(sensor.Number, "number")
    sensor.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/sensors/" + sensor.EntityData.SegmentPath
    sensor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sensor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sensor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sensor.EntityData.Children = types.NewOrderedMap()
    sensor.EntityData.Children.Append("basic-attributes", types.YChild{"BasicAttributes", &sensor.BasicAttributes})
    sensor.EntityData.Leafs = types.NewOrderedMap()
    sensor.EntityData.Leafs.Append("number", types.YLeaf{"Number", sensor.Number})

    sensor.EntityData.YListKeys = []string {"Number"}

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
    basicAttributes.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/sensors/sensor/" + basicAttributes.EntityData.SegmentPath
    basicAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicAttributes.EntityData.Children = types.NewOrderedMap()
    basicAttributes.EntityData.Children.Append("basic-info", types.YChild{"BasicInfo", &basicAttributes.BasicInfo})
    basicAttributes.EntityData.Children.Append("fru-info", types.YChild{"FruInfo", &basicAttributes.FruInfo})
    basicAttributes.EntityData.Leafs = types.NewOrderedMap()

    basicAttributes.EntityData.YListKeys = []string {}

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

    // integer value for Redundancy State if applicable to this entity. The type
    // is interface{} with range: -2147483648..2147483647.
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
    basicInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/sensors/sensor/basic-attributes/" + basicInfo.EntityData.SegmentPath
    basicInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicInfo.EntityData.Children = types.NewOrderedMap()
    basicInfo.EntityData.Leafs = types.NewOrderedMap()
    basicInfo.EntityData.Leafs.Append("description", types.YLeaf{"Description", basicInfo.Description})
    basicInfo.EntityData.Leafs.Append("vendor-type", types.YLeaf{"VendorType", basicInfo.VendorType})
    basicInfo.EntityData.Leafs.Append("name", types.YLeaf{"Name", basicInfo.Name})
    basicInfo.EntityData.Leafs.Append("hardware-revision", types.YLeaf{"HardwareRevision", basicInfo.HardwareRevision})
    basicInfo.EntityData.Leafs.Append("firmware-revision", types.YLeaf{"FirmwareRevision", basicInfo.FirmwareRevision})
    basicInfo.EntityData.Leafs.Append("software-revision", types.YLeaf{"SoftwareRevision", basicInfo.SoftwareRevision})
    basicInfo.EntityData.Leafs.Append("chip-hardware-revision", types.YLeaf{"ChipHardwareRevision", basicInfo.ChipHardwareRevision})
    basicInfo.EntityData.Leafs.Append("serial-number", types.YLeaf{"SerialNumber", basicInfo.SerialNumber})
    basicInfo.EntityData.Leafs.Append("manufacturer-name", types.YLeaf{"ManufacturerName", basicInfo.ManufacturerName})
    basicInfo.EntityData.Leafs.Append("model-name", types.YLeaf{"ModelName", basicInfo.ModelName})
    basicInfo.EntityData.Leafs.Append("asset-id-str", types.YLeaf{"AssetIdStr", basicInfo.AssetIdStr})
    basicInfo.EntityData.Leafs.Append("asset-identification", types.YLeaf{"AssetIdentification", basicInfo.AssetIdentification})
    basicInfo.EntityData.Leafs.Append("is-field-replaceable-unit", types.YLeaf{"IsFieldReplaceableUnit", basicInfo.IsFieldReplaceableUnit})
    basicInfo.EntityData.Leafs.Append("manufacturer-asset-tags", types.YLeaf{"ManufacturerAssetTags", basicInfo.ManufacturerAssetTags})
    basicInfo.EntityData.Leafs.Append("composite-class-code", types.YLeaf{"CompositeClassCode", basicInfo.CompositeClassCode})
    basicInfo.EntityData.Leafs.Append("memory-size", types.YLeaf{"MemorySize", basicInfo.MemorySize})
    basicInfo.EntityData.Leafs.Append("environmental-monitor-path", types.YLeaf{"EnvironmentalMonitorPath", basicInfo.EnvironmentalMonitorPath})
    basicInfo.EntityData.Leafs.Append("alias", types.YLeaf{"Alias", basicInfo.Alias})
    basicInfo.EntityData.Leafs.Append("group-flag", types.YLeaf{"GroupFlag", basicInfo.GroupFlag})
    basicInfo.EntityData.Leafs.Append("new-deviation-number", types.YLeaf{"NewDeviationNumber", basicInfo.NewDeviationNumber})
    basicInfo.EntityData.Leafs.Append("physical-layer-interface-module-type", types.YLeaf{"PhysicalLayerInterfaceModuleType", basicInfo.PhysicalLayerInterfaceModuleType})
    basicInfo.EntityData.Leafs.Append("unrecognized-fru", types.YLeaf{"UnrecognizedFru", basicInfo.UnrecognizedFru})
    basicInfo.EntityData.Leafs.Append("redundancystate", types.YLeaf{"Redundancystate", basicInfo.Redundancystate})
    basicInfo.EntityData.Leafs.Append("ceport", types.YLeaf{"Ceport", basicInfo.Ceport})
    basicInfo.EntityData.Leafs.Append("xr-scoped", types.YLeaf{"XrScoped", basicInfo.XrScoped})
    basicInfo.EntityData.Leafs.Append("unique-id", types.YLeaf{"UniqueId", basicInfo.UniqueId})

    basicInfo.EntityData.YListKeys = []string {}

    return &(basicInfo.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_FruInfo
// Field Replaceable Unit (FRU) inventory
// information
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_FruInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Administrative    state. The type is InvAdminState.
    ModuleAdministrativeState interface{}

    // Power administrative state. The type is InvPowerAdminState.
    ModulePowerAdministrativeState interface{}

    // Operation state. The type is InvCardState.
    ModuleOperationalState interface{}

    // Monitor state. The type is InvMonitorState.
    ModuleMonitorState interface{}

    // Reset reason. The type is InvResetReason.
    ModuleResetReason interface{}

    // Time operational state is   last changed.
    LastOperationalStateChange Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_FruInfo_LastOperationalStateChange

    // Module up time.
    ModuleUpTime Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_FruInfo_ModuleUpTime
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_FruInfo) GetEntityData() *types.CommonEntityData {
    fruInfo.EntityData.YFilter = fruInfo.YFilter
    fruInfo.EntityData.YangName = "fru-info"
    fruInfo.EntityData.BundleName = "cisco_ios_xr"
    fruInfo.EntityData.ParentYangName = "basic-attributes"
    fruInfo.EntityData.SegmentPath = "fru-info"
    fruInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/sensors/sensor/basic-attributes/" + fruInfo.EntityData.SegmentPath
    fruInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fruInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fruInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fruInfo.EntityData.Children = types.NewOrderedMap()
    fruInfo.EntityData.Children.Append("last-operational-state-change", types.YChild{"LastOperationalStateChange", &fruInfo.LastOperationalStateChange})
    fruInfo.EntityData.Children.Append("module-up-time", types.YChild{"ModuleUpTime", &fruInfo.ModuleUpTime})
    fruInfo.EntityData.Leafs = types.NewOrderedMap()
    fruInfo.EntityData.Leafs.Append("module-administrative-state", types.YLeaf{"ModuleAdministrativeState", fruInfo.ModuleAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-power-administrative-state", types.YLeaf{"ModulePowerAdministrativeState", fruInfo.ModulePowerAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-operational-state", types.YLeaf{"ModuleOperationalState", fruInfo.ModuleOperationalState})
    fruInfo.EntityData.Leafs.Append("module-monitor-state", types.YLeaf{"ModuleMonitorState", fruInfo.ModuleMonitorState})
    fruInfo.EntityData.Leafs.Append("module-reset-reason", types.YLeaf{"ModuleResetReason", fruInfo.ModuleResetReason})

    fruInfo.EntityData.YListKeys = []string {}

    return &(fruInfo.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_FruInfo_LastOperationalStateChange
// Time operational state is   last changed
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
    lastOperationalStateChange.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/sensors/sensor/basic-attributes/fru-info/" + lastOperationalStateChange.EntityData.SegmentPath
    lastOperationalStateChange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastOperationalStateChange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastOperationalStateChange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastOperationalStateChange.EntityData.Children = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", lastOperationalStateChange.TimeInSeconds})
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", lastOperationalStateChange.TimeInNanoSeconds})

    lastOperationalStateChange.EntityData.YListKeys = []string {}

    return &(lastOperationalStateChange.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_FruInfo_ModuleUpTime
// Module up time
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_FruInfo_ModuleUpTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (moduleUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_FruInfo_ModuleUpTime) GetEntityData() *types.CommonEntityData {
    moduleUpTime.EntityData.YFilter = moduleUpTime.YFilter
    moduleUpTime.EntityData.YangName = "module-up-time"
    moduleUpTime.EntityData.BundleName = "cisco_ios_xr"
    moduleUpTime.EntityData.ParentYangName = "fru-info"
    moduleUpTime.EntityData.SegmentPath = "module-up-time"
    moduleUpTime.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/sensors/sensor/basic-attributes/fru-info/" + moduleUpTime.EntityData.SegmentPath
    moduleUpTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    moduleUpTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    moduleUpTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    moduleUpTime.EntityData.Children = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", moduleUpTime.TimeInSeconds})
    moduleUpTime.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", moduleUpTime.TimeInNanoSeconds})

    moduleUpTime.EntityData.YListKeys = []string {}

    return &(moduleUpTime.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots
// PortSlotTable contains all optics ports in a
// SPA/PLIM.
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PortSlot number in the SPA/PLIM. The type is slice of
    // Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot.
    PortSlot []*Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot
}

func (portSlots *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots) GetEntityData() *types.CommonEntityData {
    portSlots.EntityData.YFilter = portSlots.YFilter
    portSlots.EntityData.YangName = "port-slots"
    portSlots.EntityData.BundleName = "cisco_ios_xr"
    portSlots.EntityData.ParentYangName = "module"
    portSlots.EntityData.SegmentPath = "port-slots"
    portSlots.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/" + portSlots.EntityData.SegmentPath
    portSlots.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portSlots.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portSlots.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portSlots.EntityData.Children = types.NewOrderedMap()
    portSlots.EntityData.Children.Append("port-slot", types.YChild{"PortSlot", nil})
    for i := range portSlots.PortSlot {
        portSlots.EntityData.Children.Append(types.GetSegmentPath(portSlots.PortSlot[i]), types.YChild{"PortSlot", portSlots.PortSlot[i]})
    }
    portSlots.EntityData.Leafs = types.NewOrderedMap()

    portSlots.EntityData.YListKeys = []string {}

    return &(portSlots.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot
// PortSlot number in the SPA/PLIM
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. portslot number. The type is interface{} with
    // range: 0..4294967295.
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
    portSlot.EntityData.SegmentPath = "port-slot" + types.AddKeyToken(portSlot.Number, "number")
    portSlot.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/port-slots/" + portSlot.EntityData.SegmentPath
    portSlot.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portSlot.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portSlot.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portSlot.EntityData.Children = types.NewOrderedMap()
    portSlot.EntityData.Children.Append("port", types.YChild{"Port", &portSlot.Port})
    portSlot.EntityData.Children.Append("basic-attributes", types.YChild{"BasicAttributes", &portSlot.BasicAttributes})
    portSlot.EntityData.Leafs = types.NewOrderedMap()
    portSlot.EntityData.Leafs.Append("number", types.YLeaf{"Number", portSlot.Number})

    portSlot.EntityData.YListKeys = []string {"Number"}

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
    port.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/port-slots/port-slot/" + port.EntityData.SegmentPath
    port.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    port.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    port.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    port.EntityData.Children = types.NewOrderedMap()
    port.EntityData.Children.Append("basic-attributes", types.YChild{"BasicAttributes", &port.BasicAttributes})
    port.EntityData.Leafs = types.NewOrderedMap()

    port.EntityData.YListKeys = []string {}

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
    basicAttributes.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/port-slots/port-slot/port/" + basicAttributes.EntityData.SegmentPath
    basicAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicAttributes.EntityData.Children = types.NewOrderedMap()
    basicAttributes.EntityData.Children.Append("basic-info", types.YChild{"BasicInfo", &basicAttributes.BasicInfo})
    basicAttributes.EntityData.Children.Append("fru-info", types.YChild{"FruInfo", &basicAttributes.FruInfo})
    basicAttributes.EntityData.Leafs = types.NewOrderedMap()

    basicAttributes.EntityData.YListKeys = []string {}

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

    // integer value for Redundancy State if applicable to this entity. The type
    // is interface{} with range: -2147483648..2147483647.
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
    basicInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/port-slots/port-slot/port/basic-attributes/" + basicInfo.EntityData.SegmentPath
    basicInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicInfo.EntityData.Children = types.NewOrderedMap()
    basicInfo.EntityData.Leafs = types.NewOrderedMap()
    basicInfo.EntityData.Leafs.Append("description", types.YLeaf{"Description", basicInfo.Description})
    basicInfo.EntityData.Leafs.Append("vendor-type", types.YLeaf{"VendorType", basicInfo.VendorType})
    basicInfo.EntityData.Leafs.Append("name", types.YLeaf{"Name", basicInfo.Name})
    basicInfo.EntityData.Leafs.Append("hardware-revision", types.YLeaf{"HardwareRevision", basicInfo.HardwareRevision})
    basicInfo.EntityData.Leafs.Append("firmware-revision", types.YLeaf{"FirmwareRevision", basicInfo.FirmwareRevision})
    basicInfo.EntityData.Leafs.Append("software-revision", types.YLeaf{"SoftwareRevision", basicInfo.SoftwareRevision})
    basicInfo.EntityData.Leafs.Append("chip-hardware-revision", types.YLeaf{"ChipHardwareRevision", basicInfo.ChipHardwareRevision})
    basicInfo.EntityData.Leafs.Append("serial-number", types.YLeaf{"SerialNumber", basicInfo.SerialNumber})
    basicInfo.EntityData.Leafs.Append("manufacturer-name", types.YLeaf{"ManufacturerName", basicInfo.ManufacturerName})
    basicInfo.EntityData.Leafs.Append("model-name", types.YLeaf{"ModelName", basicInfo.ModelName})
    basicInfo.EntityData.Leafs.Append("asset-id-str", types.YLeaf{"AssetIdStr", basicInfo.AssetIdStr})
    basicInfo.EntityData.Leafs.Append("asset-identification", types.YLeaf{"AssetIdentification", basicInfo.AssetIdentification})
    basicInfo.EntityData.Leafs.Append("is-field-replaceable-unit", types.YLeaf{"IsFieldReplaceableUnit", basicInfo.IsFieldReplaceableUnit})
    basicInfo.EntityData.Leafs.Append("manufacturer-asset-tags", types.YLeaf{"ManufacturerAssetTags", basicInfo.ManufacturerAssetTags})
    basicInfo.EntityData.Leafs.Append("composite-class-code", types.YLeaf{"CompositeClassCode", basicInfo.CompositeClassCode})
    basicInfo.EntityData.Leafs.Append("memory-size", types.YLeaf{"MemorySize", basicInfo.MemorySize})
    basicInfo.EntityData.Leafs.Append("environmental-monitor-path", types.YLeaf{"EnvironmentalMonitorPath", basicInfo.EnvironmentalMonitorPath})
    basicInfo.EntityData.Leafs.Append("alias", types.YLeaf{"Alias", basicInfo.Alias})
    basicInfo.EntityData.Leafs.Append("group-flag", types.YLeaf{"GroupFlag", basicInfo.GroupFlag})
    basicInfo.EntityData.Leafs.Append("new-deviation-number", types.YLeaf{"NewDeviationNumber", basicInfo.NewDeviationNumber})
    basicInfo.EntityData.Leafs.Append("physical-layer-interface-module-type", types.YLeaf{"PhysicalLayerInterfaceModuleType", basicInfo.PhysicalLayerInterfaceModuleType})
    basicInfo.EntityData.Leafs.Append("unrecognized-fru", types.YLeaf{"UnrecognizedFru", basicInfo.UnrecognizedFru})
    basicInfo.EntityData.Leafs.Append("redundancystate", types.YLeaf{"Redundancystate", basicInfo.Redundancystate})
    basicInfo.EntityData.Leafs.Append("ceport", types.YLeaf{"Ceport", basicInfo.Ceport})
    basicInfo.EntityData.Leafs.Append("xr-scoped", types.YLeaf{"XrScoped", basicInfo.XrScoped})
    basicInfo.EntityData.Leafs.Append("unique-id", types.YLeaf{"UniqueId", basicInfo.UniqueId})

    basicInfo.EntityData.YListKeys = []string {}

    return &(basicInfo.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_FruInfo
// Field Replaceable Unit (FRU) inventory
// information
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_FruInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Administrative    state. The type is InvAdminState.
    ModuleAdministrativeState interface{}

    // Power administrative state. The type is InvPowerAdminState.
    ModulePowerAdministrativeState interface{}

    // Operation state. The type is InvCardState.
    ModuleOperationalState interface{}

    // Monitor state. The type is InvMonitorState.
    ModuleMonitorState interface{}

    // Reset reason. The type is InvResetReason.
    ModuleResetReason interface{}

    // Time operational state is   last changed.
    LastOperationalStateChange Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_LastOperationalStateChange

    // Module up time.
    ModuleUpTime Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_ModuleUpTime
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_FruInfo) GetEntityData() *types.CommonEntityData {
    fruInfo.EntityData.YFilter = fruInfo.YFilter
    fruInfo.EntityData.YangName = "fru-info"
    fruInfo.EntityData.BundleName = "cisco_ios_xr"
    fruInfo.EntityData.ParentYangName = "basic-attributes"
    fruInfo.EntityData.SegmentPath = "fru-info"
    fruInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/port-slots/port-slot/port/basic-attributes/" + fruInfo.EntityData.SegmentPath
    fruInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fruInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fruInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fruInfo.EntityData.Children = types.NewOrderedMap()
    fruInfo.EntityData.Children.Append("last-operational-state-change", types.YChild{"LastOperationalStateChange", &fruInfo.LastOperationalStateChange})
    fruInfo.EntityData.Children.Append("module-up-time", types.YChild{"ModuleUpTime", &fruInfo.ModuleUpTime})
    fruInfo.EntityData.Leafs = types.NewOrderedMap()
    fruInfo.EntityData.Leafs.Append("module-administrative-state", types.YLeaf{"ModuleAdministrativeState", fruInfo.ModuleAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-power-administrative-state", types.YLeaf{"ModulePowerAdministrativeState", fruInfo.ModulePowerAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-operational-state", types.YLeaf{"ModuleOperationalState", fruInfo.ModuleOperationalState})
    fruInfo.EntityData.Leafs.Append("module-monitor-state", types.YLeaf{"ModuleMonitorState", fruInfo.ModuleMonitorState})
    fruInfo.EntityData.Leafs.Append("module-reset-reason", types.YLeaf{"ModuleResetReason", fruInfo.ModuleResetReason})

    fruInfo.EntityData.YListKeys = []string {}

    return &(fruInfo.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_LastOperationalStateChange
// Time operational state is   last changed
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
    lastOperationalStateChange.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/port-slots/port-slot/port/basic-attributes/fru-info/" + lastOperationalStateChange.EntityData.SegmentPath
    lastOperationalStateChange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastOperationalStateChange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastOperationalStateChange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastOperationalStateChange.EntityData.Children = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", lastOperationalStateChange.TimeInSeconds})
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", lastOperationalStateChange.TimeInNanoSeconds})

    lastOperationalStateChange.EntityData.YListKeys = []string {}

    return &(lastOperationalStateChange.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_ModuleUpTime
// Module up time
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_ModuleUpTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (moduleUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_ModuleUpTime) GetEntityData() *types.CommonEntityData {
    moduleUpTime.EntityData.YFilter = moduleUpTime.YFilter
    moduleUpTime.EntityData.YangName = "module-up-time"
    moduleUpTime.EntityData.BundleName = "cisco_ios_xr"
    moduleUpTime.EntityData.ParentYangName = "fru-info"
    moduleUpTime.EntityData.SegmentPath = "module-up-time"
    moduleUpTime.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/port-slots/port-slot/port/basic-attributes/fru-info/" + moduleUpTime.EntityData.SegmentPath
    moduleUpTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    moduleUpTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    moduleUpTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    moduleUpTime.EntityData.Children = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", moduleUpTime.TimeInSeconds})
    moduleUpTime.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", moduleUpTime.TimeInNanoSeconds})

    moduleUpTime.EntityData.YListKeys = []string {}

    return &(moduleUpTime.EntityData)
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
    basicAttributes.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/port-slots/port-slot/" + basicAttributes.EntityData.SegmentPath
    basicAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicAttributes.EntityData.Children = types.NewOrderedMap()
    basicAttributes.EntityData.Children.Append("basic-info", types.YChild{"BasicInfo", &basicAttributes.BasicInfo})
    basicAttributes.EntityData.Children.Append("fru-info", types.YChild{"FruInfo", &basicAttributes.FruInfo})
    basicAttributes.EntityData.Leafs = types.NewOrderedMap()

    basicAttributes.EntityData.YListKeys = []string {}

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

    // integer value for Redundancy State if applicable to this entity. The type
    // is interface{} with range: -2147483648..2147483647.
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
    basicInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/port-slots/port-slot/basic-attributes/" + basicInfo.EntityData.SegmentPath
    basicInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicInfo.EntityData.Children = types.NewOrderedMap()
    basicInfo.EntityData.Leafs = types.NewOrderedMap()
    basicInfo.EntityData.Leafs.Append("description", types.YLeaf{"Description", basicInfo.Description})
    basicInfo.EntityData.Leafs.Append("vendor-type", types.YLeaf{"VendorType", basicInfo.VendorType})
    basicInfo.EntityData.Leafs.Append("name", types.YLeaf{"Name", basicInfo.Name})
    basicInfo.EntityData.Leafs.Append("hardware-revision", types.YLeaf{"HardwareRevision", basicInfo.HardwareRevision})
    basicInfo.EntityData.Leafs.Append("firmware-revision", types.YLeaf{"FirmwareRevision", basicInfo.FirmwareRevision})
    basicInfo.EntityData.Leafs.Append("software-revision", types.YLeaf{"SoftwareRevision", basicInfo.SoftwareRevision})
    basicInfo.EntityData.Leafs.Append("chip-hardware-revision", types.YLeaf{"ChipHardwareRevision", basicInfo.ChipHardwareRevision})
    basicInfo.EntityData.Leafs.Append("serial-number", types.YLeaf{"SerialNumber", basicInfo.SerialNumber})
    basicInfo.EntityData.Leafs.Append("manufacturer-name", types.YLeaf{"ManufacturerName", basicInfo.ManufacturerName})
    basicInfo.EntityData.Leafs.Append("model-name", types.YLeaf{"ModelName", basicInfo.ModelName})
    basicInfo.EntityData.Leafs.Append("asset-id-str", types.YLeaf{"AssetIdStr", basicInfo.AssetIdStr})
    basicInfo.EntityData.Leafs.Append("asset-identification", types.YLeaf{"AssetIdentification", basicInfo.AssetIdentification})
    basicInfo.EntityData.Leafs.Append("is-field-replaceable-unit", types.YLeaf{"IsFieldReplaceableUnit", basicInfo.IsFieldReplaceableUnit})
    basicInfo.EntityData.Leafs.Append("manufacturer-asset-tags", types.YLeaf{"ManufacturerAssetTags", basicInfo.ManufacturerAssetTags})
    basicInfo.EntityData.Leafs.Append("composite-class-code", types.YLeaf{"CompositeClassCode", basicInfo.CompositeClassCode})
    basicInfo.EntityData.Leafs.Append("memory-size", types.YLeaf{"MemorySize", basicInfo.MemorySize})
    basicInfo.EntityData.Leafs.Append("environmental-monitor-path", types.YLeaf{"EnvironmentalMonitorPath", basicInfo.EnvironmentalMonitorPath})
    basicInfo.EntityData.Leafs.Append("alias", types.YLeaf{"Alias", basicInfo.Alias})
    basicInfo.EntityData.Leafs.Append("group-flag", types.YLeaf{"GroupFlag", basicInfo.GroupFlag})
    basicInfo.EntityData.Leafs.Append("new-deviation-number", types.YLeaf{"NewDeviationNumber", basicInfo.NewDeviationNumber})
    basicInfo.EntityData.Leafs.Append("physical-layer-interface-module-type", types.YLeaf{"PhysicalLayerInterfaceModuleType", basicInfo.PhysicalLayerInterfaceModuleType})
    basicInfo.EntityData.Leafs.Append("unrecognized-fru", types.YLeaf{"UnrecognizedFru", basicInfo.UnrecognizedFru})
    basicInfo.EntityData.Leafs.Append("redundancystate", types.YLeaf{"Redundancystate", basicInfo.Redundancystate})
    basicInfo.EntityData.Leafs.Append("ceport", types.YLeaf{"Ceport", basicInfo.Ceport})
    basicInfo.EntityData.Leafs.Append("xr-scoped", types.YLeaf{"XrScoped", basicInfo.XrScoped})
    basicInfo.EntityData.Leafs.Append("unique-id", types.YLeaf{"UniqueId", basicInfo.UniqueId})

    basicInfo.EntityData.YListKeys = []string {}

    return &(basicInfo.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_FruInfo
// Field Replaceable Unit (FRU) inventory
// information
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_FruInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Administrative    state. The type is InvAdminState.
    ModuleAdministrativeState interface{}

    // Power administrative state. The type is InvPowerAdminState.
    ModulePowerAdministrativeState interface{}

    // Operation state. The type is InvCardState.
    ModuleOperationalState interface{}

    // Monitor state. The type is InvMonitorState.
    ModuleMonitorState interface{}

    // Reset reason. The type is InvResetReason.
    ModuleResetReason interface{}

    // Time operational state is   last changed.
    LastOperationalStateChange Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_FruInfo_LastOperationalStateChange

    // Module up time.
    ModuleUpTime Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_FruInfo_ModuleUpTime
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_FruInfo) GetEntityData() *types.CommonEntityData {
    fruInfo.EntityData.YFilter = fruInfo.YFilter
    fruInfo.EntityData.YangName = "fru-info"
    fruInfo.EntityData.BundleName = "cisco_ios_xr"
    fruInfo.EntityData.ParentYangName = "basic-attributes"
    fruInfo.EntityData.SegmentPath = "fru-info"
    fruInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/port-slots/port-slot/basic-attributes/" + fruInfo.EntityData.SegmentPath
    fruInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fruInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fruInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fruInfo.EntityData.Children = types.NewOrderedMap()
    fruInfo.EntityData.Children.Append("last-operational-state-change", types.YChild{"LastOperationalStateChange", &fruInfo.LastOperationalStateChange})
    fruInfo.EntityData.Children.Append("module-up-time", types.YChild{"ModuleUpTime", &fruInfo.ModuleUpTime})
    fruInfo.EntityData.Leafs = types.NewOrderedMap()
    fruInfo.EntityData.Leafs.Append("module-administrative-state", types.YLeaf{"ModuleAdministrativeState", fruInfo.ModuleAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-power-administrative-state", types.YLeaf{"ModulePowerAdministrativeState", fruInfo.ModulePowerAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-operational-state", types.YLeaf{"ModuleOperationalState", fruInfo.ModuleOperationalState})
    fruInfo.EntityData.Leafs.Append("module-monitor-state", types.YLeaf{"ModuleMonitorState", fruInfo.ModuleMonitorState})
    fruInfo.EntityData.Leafs.Append("module-reset-reason", types.YLeaf{"ModuleResetReason", fruInfo.ModuleResetReason})

    fruInfo.EntityData.YListKeys = []string {}

    return &(fruInfo.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_FruInfo_LastOperationalStateChange
// Time operational state is   last changed
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
    lastOperationalStateChange.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/port-slots/port-slot/basic-attributes/fru-info/" + lastOperationalStateChange.EntityData.SegmentPath
    lastOperationalStateChange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastOperationalStateChange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastOperationalStateChange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastOperationalStateChange.EntityData.Children = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", lastOperationalStateChange.TimeInSeconds})
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", lastOperationalStateChange.TimeInNanoSeconds})

    lastOperationalStateChange.EntityData.YListKeys = []string {}

    return &(lastOperationalStateChange.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_FruInfo_ModuleUpTime
// Module up time
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_FruInfo_ModuleUpTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (moduleUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_FruInfo_ModuleUpTime) GetEntityData() *types.CommonEntityData {
    moduleUpTime.EntityData.YFilter = moduleUpTime.YFilter
    moduleUpTime.EntityData.YangName = "module-up-time"
    moduleUpTime.EntityData.BundleName = "cisco_ios_xr"
    moduleUpTime.EntityData.ParentYangName = "fru-info"
    moduleUpTime.EntityData.SegmentPath = "module-up-time"
    moduleUpTime.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/port-slots/port-slot/basic-attributes/fru-info/" + moduleUpTime.EntityData.SegmentPath
    moduleUpTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    moduleUpTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    moduleUpTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    moduleUpTime.EntityData.Children = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", moduleUpTime.TimeInSeconds})
    moduleUpTime.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", moduleUpTime.TimeInNanoSeconds})

    moduleUpTime.EntityData.YListKeys = []string {}

    return &(moduleUpTime.EntityData)
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
    basicAttributes.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/" + basicAttributes.EntityData.SegmentPath
    basicAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicAttributes.EntityData.Children = types.NewOrderedMap()
    basicAttributes.EntityData.Children.Append("basic-info", types.YChild{"BasicInfo", &basicAttributes.BasicInfo})
    basicAttributes.EntityData.Children.Append("fru-info", types.YChild{"FruInfo", &basicAttributes.FruInfo})
    basicAttributes.EntityData.Leafs = types.NewOrderedMap()

    basicAttributes.EntityData.YListKeys = []string {}

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

    // integer value for Redundancy State if applicable to this entity. The type
    // is interface{} with range: -2147483648..2147483647.
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
    basicInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/basic-attributes/" + basicInfo.EntityData.SegmentPath
    basicInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicInfo.EntityData.Children = types.NewOrderedMap()
    basicInfo.EntityData.Leafs = types.NewOrderedMap()
    basicInfo.EntityData.Leafs.Append("description", types.YLeaf{"Description", basicInfo.Description})
    basicInfo.EntityData.Leafs.Append("vendor-type", types.YLeaf{"VendorType", basicInfo.VendorType})
    basicInfo.EntityData.Leafs.Append("name", types.YLeaf{"Name", basicInfo.Name})
    basicInfo.EntityData.Leafs.Append("hardware-revision", types.YLeaf{"HardwareRevision", basicInfo.HardwareRevision})
    basicInfo.EntityData.Leafs.Append("firmware-revision", types.YLeaf{"FirmwareRevision", basicInfo.FirmwareRevision})
    basicInfo.EntityData.Leafs.Append("software-revision", types.YLeaf{"SoftwareRevision", basicInfo.SoftwareRevision})
    basicInfo.EntityData.Leafs.Append("chip-hardware-revision", types.YLeaf{"ChipHardwareRevision", basicInfo.ChipHardwareRevision})
    basicInfo.EntityData.Leafs.Append("serial-number", types.YLeaf{"SerialNumber", basicInfo.SerialNumber})
    basicInfo.EntityData.Leafs.Append("manufacturer-name", types.YLeaf{"ManufacturerName", basicInfo.ManufacturerName})
    basicInfo.EntityData.Leafs.Append("model-name", types.YLeaf{"ModelName", basicInfo.ModelName})
    basicInfo.EntityData.Leafs.Append("asset-id-str", types.YLeaf{"AssetIdStr", basicInfo.AssetIdStr})
    basicInfo.EntityData.Leafs.Append("asset-identification", types.YLeaf{"AssetIdentification", basicInfo.AssetIdentification})
    basicInfo.EntityData.Leafs.Append("is-field-replaceable-unit", types.YLeaf{"IsFieldReplaceableUnit", basicInfo.IsFieldReplaceableUnit})
    basicInfo.EntityData.Leafs.Append("manufacturer-asset-tags", types.YLeaf{"ManufacturerAssetTags", basicInfo.ManufacturerAssetTags})
    basicInfo.EntityData.Leafs.Append("composite-class-code", types.YLeaf{"CompositeClassCode", basicInfo.CompositeClassCode})
    basicInfo.EntityData.Leafs.Append("memory-size", types.YLeaf{"MemorySize", basicInfo.MemorySize})
    basicInfo.EntityData.Leafs.Append("environmental-monitor-path", types.YLeaf{"EnvironmentalMonitorPath", basicInfo.EnvironmentalMonitorPath})
    basicInfo.EntityData.Leafs.Append("alias", types.YLeaf{"Alias", basicInfo.Alias})
    basicInfo.EntityData.Leafs.Append("group-flag", types.YLeaf{"GroupFlag", basicInfo.GroupFlag})
    basicInfo.EntityData.Leafs.Append("new-deviation-number", types.YLeaf{"NewDeviationNumber", basicInfo.NewDeviationNumber})
    basicInfo.EntityData.Leafs.Append("physical-layer-interface-module-type", types.YLeaf{"PhysicalLayerInterfaceModuleType", basicInfo.PhysicalLayerInterfaceModuleType})
    basicInfo.EntityData.Leafs.Append("unrecognized-fru", types.YLeaf{"UnrecognizedFru", basicInfo.UnrecognizedFru})
    basicInfo.EntityData.Leafs.Append("redundancystate", types.YLeaf{"Redundancystate", basicInfo.Redundancystate})
    basicInfo.EntityData.Leafs.Append("ceport", types.YLeaf{"Ceport", basicInfo.Ceport})
    basicInfo.EntityData.Leafs.Append("xr-scoped", types.YLeaf{"XrScoped", basicInfo.XrScoped})
    basicInfo.EntityData.Leafs.Append("unique-id", types.YLeaf{"UniqueId", basicInfo.UniqueId})

    basicInfo.EntityData.YListKeys = []string {}

    return &(basicInfo.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_FruInfo
// Field Replaceable Unit (FRU) inventory
// information
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_FruInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Administrative    state. The type is InvAdminState.
    ModuleAdministrativeState interface{}

    // Power administrative state. The type is InvPowerAdminState.
    ModulePowerAdministrativeState interface{}

    // Operation state. The type is InvCardState.
    ModuleOperationalState interface{}

    // Monitor state. The type is InvMonitorState.
    ModuleMonitorState interface{}

    // Reset reason. The type is InvResetReason.
    ModuleResetReason interface{}

    // Time operational state is   last changed.
    LastOperationalStateChange Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_FruInfo_LastOperationalStateChange

    // Module up time.
    ModuleUpTime Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_FruInfo_ModuleUpTime
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_FruInfo) GetEntityData() *types.CommonEntityData {
    fruInfo.EntityData.YFilter = fruInfo.YFilter
    fruInfo.EntityData.YangName = "fru-info"
    fruInfo.EntityData.BundleName = "cisco_ios_xr"
    fruInfo.EntityData.ParentYangName = "basic-attributes"
    fruInfo.EntityData.SegmentPath = "fru-info"
    fruInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/basic-attributes/" + fruInfo.EntityData.SegmentPath
    fruInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fruInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fruInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fruInfo.EntityData.Children = types.NewOrderedMap()
    fruInfo.EntityData.Children.Append("last-operational-state-change", types.YChild{"LastOperationalStateChange", &fruInfo.LastOperationalStateChange})
    fruInfo.EntityData.Children.Append("module-up-time", types.YChild{"ModuleUpTime", &fruInfo.ModuleUpTime})
    fruInfo.EntityData.Leafs = types.NewOrderedMap()
    fruInfo.EntityData.Leafs.Append("module-administrative-state", types.YLeaf{"ModuleAdministrativeState", fruInfo.ModuleAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-power-administrative-state", types.YLeaf{"ModulePowerAdministrativeState", fruInfo.ModulePowerAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-operational-state", types.YLeaf{"ModuleOperationalState", fruInfo.ModuleOperationalState})
    fruInfo.EntityData.Leafs.Append("module-monitor-state", types.YLeaf{"ModuleMonitorState", fruInfo.ModuleMonitorState})
    fruInfo.EntityData.Leafs.Append("module-reset-reason", types.YLeaf{"ModuleResetReason", fruInfo.ModuleResetReason})

    fruInfo.EntityData.YListKeys = []string {}

    return &(fruInfo.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_FruInfo_LastOperationalStateChange
// Time operational state is   last changed
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
    lastOperationalStateChange.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/basic-attributes/fru-info/" + lastOperationalStateChange.EntityData.SegmentPath
    lastOperationalStateChange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastOperationalStateChange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastOperationalStateChange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastOperationalStateChange.EntityData.Children = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", lastOperationalStateChange.TimeInSeconds})
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", lastOperationalStateChange.TimeInNanoSeconds})

    lastOperationalStateChange.EntityData.YListKeys = []string {}

    return &(lastOperationalStateChange.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_FruInfo_ModuleUpTime
// Module up time
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_FruInfo_ModuleUpTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (moduleUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_FruInfo_ModuleUpTime) GetEntityData() *types.CommonEntityData {
    moduleUpTime.EntityData.YFilter = moduleUpTime.YFilter
    moduleUpTime.EntityData.YangName = "module-up-time"
    moduleUpTime.EntityData.BundleName = "cisco_ios_xr"
    moduleUpTime.EntityData.ParentYangName = "fru-info"
    moduleUpTime.EntityData.SegmentPath = "module-up-time"
    moduleUpTime.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/basic-attributes/fru-info/" + moduleUpTime.EntityData.SegmentPath
    moduleUpTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    moduleUpTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    moduleUpTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    moduleUpTime.EntityData.Children = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", moduleUpTime.TimeInSeconds})
    moduleUpTime.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", moduleUpTime.TimeInNanoSeconds})

    moduleUpTime.EntityData.YListKeys = []string {}

    return &(moduleUpTime.EntityData)
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
    basicAttributes.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/" + basicAttributes.EntityData.SegmentPath
    basicAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicAttributes.EntityData.Children = types.NewOrderedMap()
    basicAttributes.EntityData.Children.Append("basic-info", types.YChild{"BasicInfo", &basicAttributes.BasicInfo})
    basicAttributes.EntityData.Children.Append("fru-info", types.YChild{"FruInfo", &basicAttributes.FruInfo})
    basicAttributes.EntityData.Leafs = types.NewOrderedMap()

    basicAttributes.EntityData.YListKeys = []string {}

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

    // integer value for Redundancy State if applicable to this entity. The type
    // is interface{} with range: -2147483648..2147483647.
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
    basicInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/basic-attributes/" + basicInfo.EntityData.SegmentPath
    basicInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicInfo.EntityData.Children = types.NewOrderedMap()
    basicInfo.EntityData.Leafs = types.NewOrderedMap()
    basicInfo.EntityData.Leafs.Append("description", types.YLeaf{"Description", basicInfo.Description})
    basicInfo.EntityData.Leafs.Append("vendor-type", types.YLeaf{"VendorType", basicInfo.VendorType})
    basicInfo.EntityData.Leafs.Append("name", types.YLeaf{"Name", basicInfo.Name})
    basicInfo.EntityData.Leafs.Append("hardware-revision", types.YLeaf{"HardwareRevision", basicInfo.HardwareRevision})
    basicInfo.EntityData.Leafs.Append("firmware-revision", types.YLeaf{"FirmwareRevision", basicInfo.FirmwareRevision})
    basicInfo.EntityData.Leafs.Append("software-revision", types.YLeaf{"SoftwareRevision", basicInfo.SoftwareRevision})
    basicInfo.EntityData.Leafs.Append("chip-hardware-revision", types.YLeaf{"ChipHardwareRevision", basicInfo.ChipHardwareRevision})
    basicInfo.EntityData.Leafs.Append("serial-number", types.YLeaf{"SerialNumber", basicInfo.SerialNumber})
    basicInfo.EntityData.Leafs.Append("manufacturer-name", types.YLeaf{"ManufacturerName", basicInfo.ManufacturerName})
    basicInfo.EntityData.Leafs.Append("model-name", types.YLeaf{"ModelName", basicInfo.ModelName})
    basicInfo.EntityData.Leafs.Append("asset-id-str", types.YLeaf{"AssetIdStr", basicInfo.AssetIdStr})
    basicInfo.EntityData.Leafs.Append("asset-identification", types.YLeaf{"AssetIdentification", basicInfo.AssetIdentification})
    basicInfo.EntityData.Leafs.Append("is-field-replaceable-unit", types.YLeaf{"IsFieldReplaceableUnit", basicInfo.IsFieldReplaceableUnit})
    basicInfo.EntityData.Leafs.Append("manufacturer-asset-tags", types.YLeaf{"ManufacturerAssetTags", basicInfo.ManufacturerAssetTags})
    basicInfo.EntityData.Leafs.Append("composite-class-code", types.YLeaf{"CompositeClassCode", basicInfo.CompositeClassCode})
    basicInfo.EntityData.Leafs.Append("memory-size", types.YLeaf{"MemorySize", basicInfo.MemorySize})
    basicInfo.EntityData.Leafs.Append("environmental-monitor-path", types.YLeaf{"EnvironmentalMonitorPath", basicInfo.EnvironmentalMonitorPath})
    basicInfo.EntityData.Leafs.Append("alias", types.YLeaf{"Alias", basicInfo.Alias})
    basicInfo.EntityData.Leafs.Append("group-flag", types.YLeaf{"GroupFlag", basicInfo.GroupFlag})
    basicInfo.EntityData.Leafs.Append("new-deviation-number", types.YLeaf{"NewDeviationNumber", basicInfo.NewDeviationNumber})
    basicInfo.EntityData.Leafs.Append("physical-layer-interface-module-type", types.YLeaf{"PhysicalLayerInterfaceModuleType", basicInfo.PhysicalLayerInterfaceModuleType})
    basicInfo.EntityData.Leafs.Append("unrecognized-fru", types.YLeaf{"UnrecognizedFru", basicInfo.UnrecognizedFru})
    basicInfo.EntityData.Leafs.Append("redundancystate", types.YLeaf{"Redundancystate", basicInfo.Redundancystate})
    basicInfo.EntityData.Leafs.Append("ceport", types.YLeaf{"Ceport", basicInfo.Ceport})
    basicInfo.EntityData.Leafs.Append("xr-scoped", types.YLeaf{"XrScoped", basicInfo.XrScoped})
    basicInfo.EntityData.Leafs.Append("unique-id", types.YLeaf{"UniqueId", basicInfo.UniqueId})

    basicInfo.EntityData.YListKeys = []string {}

    return &(basicInfo.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_FruInfo
// Field Replaceable Unit (FRU) inventory
// information
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_FruInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Administrative    state. The type is InvAdminState.
    ModuleAdministrativeState interface{}

    // Power administrative state. The type is InvPowerAdminState.
    ModulePowerAdministrativeState interface{}

    // Operation state. The type is InvCardState.
    ModuleOperationalState interface{}

    // Monitor state. The type is InvMonitorState.
    ModuleMonitorState interface{}

    // Reset reason. The type is InvResetReason.
    ModuleResetReason interface{}

    // Time operational state is   last changed.
    LastOperationalStateChange Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_FruInfo_LastOperationalStateChange

    // Module up time.
    ModuleUpTime Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_FruInfo_ModuleUpTime
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_FruInfo) GetEntityData() *types.CommonEntityData {
    fruInfo.EntityData.YFilter = fruInfo.YFilter
    fruInfo.EntityData.YangName = "fru-info"
    fruInfo.EntityData.BundleName = "cisco_ios_xr"
    fruInfo.EntityData.ParentYangName = "basic-attributes"
    fruInfo.EntityData.SegmentPath = "fru-info"
    fruInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/basic-attributes/" + fruInfo.EntityData.SegmentPath
    fruInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fruInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fruInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fruInfo.EntityData.Children = types.NewOrderedMap()
    fruInfo.EntityData.Children.Append("last-operational-state-change", types.YChild{"LastOperationalStateChange", &fruInfo.LastOperationalStateChange})
    fruInfo.EntityData.Children.Append("module-up-time", types.YChild{"ModuleUpTime", &fruInfo.ModuleUpTime})
    fruInfo.EntityData.Leafs = types.NewOrderedMap()
    fruInfo.EntityData.Leafs.Append("module-administrative-state", types.YLeaf{"ModuleAdministrativeState", fruInfo.ModuleAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-power-administrative-state", types.YLeaf{"ModulePowerAdministrativeState", fruInfo.ModulePowerAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-operational-state", types.YLeaf{"ModuleOperationalState", fruInfo.ModuleOperationalState})
    fruInfo.EntityData.Leafs.Append("module-monitor-state", types.YLeaf{"ModuleMonitorState", fruInfo.ModuleMonitorState})
    fruInfo.EntityData.Leafs.Append("module-reset-reason", types.YLeaf{"ModuleResetReason", fruInfo.ModuleResetReason})

    fruInfo.EntityData.YListKeys = []string {}

    return &(fruInfo.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_FruInfo_LastOperationalStateChange
// Time operational state is   last changed
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
    lastOperationalStateChange.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/basic-attributes/fru-info/" + lastOperationalStateChange.EntityData.SegmentPath
    lastOperationalStateChange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastOperationalStateChange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastOperationalStateChange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastOperationalStateChange.EntityData.Children = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", lastOperationalStateChange.TimeInSeconds})
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", lastOperationalStateChange.TimeInNanoSeconds})

    lastOperationalStateChange.EntityData.YListKeys = []string {}

    return &(lastOperationalStateChange.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_FruInfo_ModuleUpTime
// Module up time
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_FruInfo_ModuleUpTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (moduleUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_FruInfo_ModuleUpTime) GetEntityData() *types.CommonEntityData {
    moduleUpTime.EntityData.YFilter = moduleUpTime.YFilter
    moduleUpTime.EntityData.YangName = "module-up-time"
    moduleUpTime.EntityData.BundleName = "cisco_ios_xr"
    moduleUpTime.EntityData.ParentYangName = "fru-info"
    moduleUpTime.EntityData.SegmentPath = "module-up-time"
    moduleUpTime.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/basic-attributes/fru-info/" + moduleUpTime.EntityData.SegmentPath
    moduleUpTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    moduleUpTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    moduleUpTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    moduleUpTime.EntityData.Children = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", moduleUpTime.TimeInSeconds})
    moduleUpTime.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", moduleUpTime.TimeInNanoSeconds})

    moduleUpTime.EntityData.YListKeys = []string {}

    return &(moduleUpTime.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents
// HWComponent table contains all HW modules
// within the card 
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // HWComponent number. The type is slice of
    // Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent.
    HwComponent []*Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent
}

func (hwComponents *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents) GetEntityData() *types.CommonEntityData {
    hwComponents.EntityData.YFilter = hwComponents.YFilter
    hwComponents.EntityData.YangName = "hw-components"
    hwComponents.EntityData.BundleName = "cisco_ios_xr"
    hwComponents.EntityData.ParentYangName = "card"
    hwComponents.EntityData.SegmentPath = "hw-components"
    hwComponents.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/cards/card/" + hwComponents.EntityData.SegmentPath
    hwComponents.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hwComponents.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hwComponents.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hwComponents.EntityData.Children = types.NewOrderedMap()
    hwComponents.EntityData.Children.Append("hw-component", types.YChild{"HwComponent", nil})
    for i := range hwComponents.HwComponent {
        hwComponents.EntityData.Children.Append(types.GetSegmentPath(hwComponents.HwComponent[i]), types.YChild{"HwComponent", hwComponents.HwComponent[i]})
    }
    hwComponents.EntityData.Leafs = types.NewOrderedMap()

    hwComponents.EntityData.YListKeys = []string {}

    return &(hwComponents.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent
// HWComponent number
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. node number. The type is interface{} with range:
    // 0..4294967295.
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
    hwComponent.EntityData.SegmentPath = "hw-component" + types.AddKeyToken(hwComponent.Number, "number")
    hwComponent.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/cards/card/hw-components/" + hwComponent.EntityData.SegmentPath
    hwComponent.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hwComponent.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hwComponent.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hwComponent.EntityData.Children = types.NewOrderedMap()
    hwComponent.EntityData.Children.Append("sensors", types.YChild{"Sensors", &hwComponent.Sensors})
    hwComponent.EntityData.Children.Append("basic-attributes", types.YChild{"BasicAttributes", &hwComponent.BasicAttributes})
    hwComponent.EntityData.Leafs = types.NewOrderedMap()
    hwComponent.EntityData.Leafs.Append("number", types.YLeaf{"Number", hwComponent.Number})

    hwComponent.EntityData.YListKeys = []string {"Number"}

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
    Sensor []*Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor
}

func (sensors *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors) GetEntityData() *types.CommonEntityData {
    sensors.EntityData.YFilter = sensors.YFilter
    sensors.EntityData.YangName = "sensors"
    sensors.EntityData.BundleName = "cisco_ios_xr"
    sensors.EntityData.ParentYangName = "hw-component"
    sensors.EntityData.SegmentPath = "sensors"
    sensors.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/cards/card/hw-components/hw-component/" + sensors.EntityData.SegmentPath
    sensors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sensors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sensors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sensors.EntityData.Children = types.NewOrderedMap()
    sensors.EntityData.Children.Append("sensor", types.YChild{"Sensor", nil})
    for i := range sensors.Sensor {
        sensors.EntityData.Children.Append(types.GetSegmentPath(sensors.Sensor[i]), types.YChild{"Sensor", sensors.Sensor[i]})
    }
    sensors.EntityData.Leafs = types.NewOrderedMap()

    sensors.EntityData.YListKeys = []string {}

    return &(sensors.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor
// Sensor number in the Module
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. sensor number. The type is interface{} with range:
    // 0..4294967295.
    Number interface{}

    // Attributes.
    BasicAttributes Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes
}

func (sensor *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor) GetEntityData() *types.CommonEntityData {
    sensor.EntityData.YFilter = sensor.YFilter
    sensor.EntityData.YangName = "sensor"
    sensor.EntityData.BundleName = "cisco_ios_xr"
    sensor.EntityData.ParentYangName = "sensors"
    sensor.EntityData.SegmentPath = "sensor" + types.AddKeyToken(sensor.Number, "number")
    sensor.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/cards/card/hw-components/hw-component/sensors/" + sensor.EntityData.SegmentPath
    sensor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sensor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sensor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sensor.EntityData.Children = types.NewOrderedMap()
    sensor.EntityData.Children.Append("basic-attributes", types.YChild{"BasicAttributes", &sensor.BasicAttributes})
    sensor.EntityData.Leafs = types.NewOrderedMap()
    sensor.EntityData.Leafs.Append("number", types.YLeaf{"Number", sensor.Number})

    sensor.EntityData.YListKeys = []string {"Number"}

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
    basicAttributes.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/cards/card/hw-components/hw-component/sensors/sensor/" + basicAttributes.EntityData.SegmentPath
    basicAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicAttributes.EntityData.Children = types.NewOrderedMap()
    basicAttributes.EntityData.Children.Append("basic-info", types.YChild{"BasicInfo", &basicAttributes.BasicInfo})
    basicAttributes.EntityData.Children.Append("fru-info", types.YChild{"FruInfo", &basicAttributes.FruInfo})
    basicAttributes.EntityData.Leafs = types.NewOrderedMap()

    basicAttributes.EntityData.YListKeys = []string {}

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

    // integer value for Redundancy State if applicable to this entity. The type
    // is interface{} with range: -2147483648..2147483647.
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
    basicInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/cards/card/hw-components/hw-component/sensors/sensor/basic-attributes/" + basicInfo.EntityData.SegmentPath
    basicInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicInfo.EntityData.Children = types.NewOrderedMap()
    basicInfo.EntityData.Leafs = types.NewOrderedMap()
    basicInfo.EntityData.Leafs.Append("description", types.YLeaf{"Description", basicInfo.Description})
    basicInfo.EntityData.Leafs.Append("vendor-type", types.YLeaf{"VendorType", basicInfo.VendorType})
    basicInfo.EntityData.Leafs.Append("name", types.YLeaf{"Name", basicInfo.Name})
    basicInfo.EntityData.Leafs.Append("hardware-revision", types.YLeaf{"HardwareRevision", basicInfo.HardwareRevision})
    basicInfo.EntityData.Leafs.Append("firmware-revision", types.YLeaf{"FirmwareRevision", basicInfo.FirmwareRevision})
    basicInfo.EntityData.Leafs.Append("software-revision", types.YLeaf{"SoftwareRevision", basicInfo.SoftwareRevision})
    basicInfo.EntityData.Leafs.Append("chip-hardware-revision", types.YLeaf{"ChipHardwareRevision", basicInfo.ChipHardwareRevision})
    basicInfo.EntityData.Leafs.Append("serial-number", types.YLeaf{"SerialNumber", basicInfo.SerialNumber})
    basicInfo.EntityData.Leafs.Append("manufacturer-name", types.YLeaf{"ManufacturerName", basicInfo.ManufacturerName})
    basicInfo.EntityData.Leafs.Append("model-name", types.YLeaf{"ModelName", basicInfo.ModelName})
    basicInfo.EntityData.Leafs.Append("asset-id-str", types.YLeaf{"AssetIdStr", basicInfo.AssetIdStr})
    basicInfo.EntityData.Leafs.Append("asset-identification", types.YLeaf{"AssetIdentification", basicInfo.AssetIdentification})
    basicInfo.EntityData.Leafs.Append("is-field-replaceable-unit", types.YLeaf{"IsFieldReplaceableUnit", basicInfo.IsFieldReplaceableUnit})
    basicInfo.EntityData.Leafs.Append("manufacturer-asset-tags", types.YLeaf{"ManufacturerAssetTags", basicInfo.ManufacturerAssetTags})
    basicInfo.EntityData.Leafs.Append("composite-class-code", types.YLeaf{"CompositeClassCode", basicInfo.CompositeClassCode})
    basicInfo.EntityData.Leafs.Append("memory-size", types.YLeaf{"MemorySize", basicInfo.MemorySize})
    basicInfo.EntityData.Leafs.Append("environmental-monitor-path", types.YLeaf{"EnvironmentalMonitorPath", basicInfo.EnvironmentalMonitorPath})
    basicInfo.EntityData.Leafs.Append("alias", types.YLeaf{"Alias", basicInfo.Alias})
    basicInfo.EntityData.Leafs.Append("group-flag", types.YLeaf{"GroupFlag", basicInfo.GroupFlag})
    basicInfo.EntityData.Leafs.Append("new-deviation-number", types.YLeaf{"NewDeviationNumber", basicInfo.NewDeviationNumber})
    basicInfo.EntityData.Leafs.Append("physical-layer-interface-module-type", types.YLeaf{"PhysicalLayerInterfaceModuleType", basicInfo.PhysicalLayerInterfaceModuleType})
    basicInfo.EntityData.Leafs.Append("unrecognized-fru", types.YLeaf{"UnrecognizedFru", basicInfo.UnrecognizedFru})
    basicInfo.EntityData.Leafs.Append("redundancystate", types.YLeaf{"Redundancystate", basicInfo.Redundancystate})
    basicInfo.EntityData.Leafs.Append("ceport", types.YLeaf{"Ceport", basicInfo.Ceport})
    basicInfo.EntityData.Leafs.Append("xr-scoped", types.YLeaf{"XrScoped", basicInfo.XrScoped})
    basicInfo.EntityData.Leafs.Append("unique-id", types.YLeaf{"UniqueId", basicInfo.UniqueId})

    basicInfo.EntityData.YListKeys = []string {}

    return &(basicInfo.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_FruInfo
// Field Replaceable Unit (FRU) inventory
// information
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_FruInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Administrative    state. The type is InvAdminState.
    ModuleAdministrativeState interface{}

    // Power administrative state. The type is InvPowerAdminState.
    ModulePowerAdministrativeState interface{}

    // Operation state. The type is InvCardState.
    ModuleOperationalState interface{}

    // Monitor state. The type is InvMonitorState.
    ModuleMonitorState interface{}

    // Reset reason. The type is InvResetReason.
    ModuleResetReason interface{}

    // Time operational state is   last changed.
    LastOperationalStateChange Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_FruInfo_LastOperationalStateChange

    // Module up time.
    ModuleUpTime Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_FruInfo_ModuleUpTime
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_FruInfo) GetEntityData() *types.CommonEntityData {
    fruInfo.EntityData.YFilter = fruInfo.YFilter
    fruInfo.EntityData.YangName = "fru-info"
    fruInfo.EntityData.BundleName = "cisco_ios_xr"
    fruInfo.EntityData.ParentYangName = "basic-attributes"
    fruInfo.EntityData.SegmentPath = "fru-info"
    fruInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/cards/card/hw-components/hw-component/sensors/sensor/basic-attributes/" + fruInfo.EntityData.SegmentPath
    fruInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fruInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fruInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fruInfo.EntityData.Children = types.NewOrderedMap()
    fruInfo.EntityData.Children.Append("last-operational-state-change", types.YChild{"LastOperationalStateChange", &fruInfo.LastOperationalStateChange})
    fruInfo.EntityData.Children.Append("module-up-time", types.YChild{"ModuleUpTime", &fruInfo.ModuleUpTime})
    fruInfo.EntityData.Leafs = types.NewOrderedMap()
    fruInfo.EntityData.Leafs.Append("module-administrative-state", types.YLeaf{"ModuleAdministrativeState", fruInfo.ModuleAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-power-administrative-state", types.YLeaf{"ModulePowerAdministrativeState", fruInfo.ModulePowerAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-operational-state", types.YLeaf{"ModuleOperationalState", fruInfo.ModuleOperationalState})
    fruInfo.EntityData.Leafs.Append("module-monitor-state", types.YLeaf{"ModuleMonitorState", fruInfo.ModuleMonitorState})
    fruInfo.EntityData.Leafs.Append("module-reset-reason", types.YLeaf{"ModuleResetReason", fruInfo.ModuleResetReason})

    fruInfo.EntityData.YListKeys = []string {}

    return &(fruInfo.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_FruInfo_LastOperationalStateChange
// Time operational state is   last changed
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
    lastOperationalStateChange.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/cards/card/hw-components/hw-component/sensors/sensor/basic-attributes/fru-info/" + lastOperationalStateChange.EntityData.SegmentPath
    lastOperationalStateChange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastOperationalStateChange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastOperationalStateChange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastOperationalStateChange.EntityData.Children = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", lastOperationalStateChange.TimeInSeconds})
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", lastOperationalStateChange.TimeInNanoSeconds})

    lastOperationalStateChange.EntityData.YListKeys = []string {}

    return &(lastOperationalStateChange.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_FruInfo_ModuleUpTime
// Module up time
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_FruInfo_ModuleUpTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (moduleUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_FruInfo_ModuleUpTime) GetEntityData() *types.CommonEntityData {
    moduleUpTime.EntityData.YFilter = moduleUpTime.YFilter
    moduleUpTime.EntityData.YangName = "module-up-time"
    moduleUpTime.EntityData.BundleName = "cisco_ios_xr"
    moduleUpTime.EntityData.ParentYangName = "fru-info"
    moduleUpTime.EntityData.SegmentPath = "module-up-time"
    moduleUpTime.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/cards/card/hw-components/hw-component/sensors/sensor/basic-attributes/fru-info/" + moduleUpTime.EntityData.SegmentPath
    moduleUpTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    moduleUpTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    moduleUpTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    moduleUpTime.EntityData.Children = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", moduleUpTime.TimeInSeconds})
    moduleUpTime.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", moduleUpTime.TimeInNanoSeconds})

    moduleUpTime.EntityData.YListKeys = []string {}

    return &(moduleUpTime.EntityData)
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
    basicAttributes.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/cards/card/hw-components/hw-component/" + basicAttributes.EntityData.SegmentPath
    basicAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicAttributes.EntityData.Children = types.NewOrderedMap()
    basicAttributes.EntityData.Children.Append("basic-info", types.YChild{"BasicInfo", &basicAttributes.BasicInfo})
    basicAttributes.EntityData.Children.Append("fru-info", types.YChild{"FruInfo", &basicAttributes.FruInfo})
    basicAttributes.EntityData.Leafs = types.NewOrderedMap()

    basicAttributes.EntityData.YListKeys = []string {}

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

    // integer value for Redundancy State if applicable to this entity. The type
    // is interface{} with range: -2147483648..2147483647.
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
    basicInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/cards/card/hw-components/hw-component/basic-attributes/" + basicInfo.EntityData.SegmentPath
    basicInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicInfo.EntityData.Children = types.NewOrderedMap()
    basicInfo.EntityData.Leafs = types.NewOrderedMap()
    basicInfo.EntityData.Leafs.Append("description", types.YLeaf{"Description", basicInfo.Description})
    basicInfo.EntityData.Leafs.Append("vendor-type", types.YLeaf{"VendorType", basicInfo.VendorType})
    basicInfo.EntityData.Leafs.Append("name", types.YLeaf{"Name", basicInfo.Name})
    basicInfo.EntityData.Leafs.Append("hardware-revision", types.YLeaf{"HardwareRevision", basicInfo.HardwareRevision})
    basicInfo.EntityData.Leafs.Append("firmware-revision", types.YLeaf{"FirmwareRevision", basicInfo.FirmwareRevision})
    basicInfo.EntityData.Leafs.Append("software-revision", types.YLeaf{"SoftwareRevision", basicInfo.SoftwareRevision})
    basicInfo.EntityData.Leafs.Append("chip-hardware-revision", types.YLeaf{"ChipHardwareRevision", basicInfo.ChipHardwareRevision})
    basicInfo.EntityData.Leafs.Append("serial-number", types.YLeaf{"SerialNumber", basicInfo.SerialNumber})
    basicInfo.EntityData.Leafs.Append("manufacturer-name", types.YLeaf{"ManufacturerName", basicInfo.ManufacturerName})
    basicInfo.EntityData.Leafs.Append("model-name", types.YLeaf{"ModelName", basicInfo.ModelName})
    basicInfo.EntityData.Leafs.Append("asset-id-str", types.YLeaf{"AssetIdStr", basicInfo.AssetIdStr})
    basicInfo.EntityData.Leafs.Append("asset-identification", types.YLeaf{"AssetIdentification", basicInfo.AssetIdentification})
    basicInfo.EntityData.Leafs.Append("is-field-replaceable-unit", types.YLeaf{"IsFieldReplaceableUnit", basicInfo.IsFieldReplaceableUnit})
    basicInfo.EntityData.Leafs.Append("manufacturer-asset-tags", types.YLeaf{"ManufacturerAssetTags", basicInfo.ManufacturerAssetTags})
    basicInfo.EntityData.Leafs.Append("composite-class-code", types.YLeaf{"CompositeClassCode", basicInfo.CompositeClassCode})
    basicInfo.EntityData.Leafs.Append("memory-size", types.YLeaf{"MemorySize", basicInfo.MemorySize})
    basicInfo.EntityData.Leafs.Append("environmental-monitor-path", types.YLeaf{"EnvironmentalMonitorPath", basicInfo.EnvironmentalMonitorPath})
    basicInfo.EntityData.Leafs.Append("alias", types.YLeaf{"Alias", basicInfo.Alias})
    basicInfo.EntityData.Leafs.Append("group-flag", types.YLeaf{"GroupFlag", basicInfo.GroupFlag})
    basicInfo.EntityData.Leafs.Append("new-deviation-number", types.YLeaf{"NewDeviationNumber", basicInfo.NewDeviationNumber})
    basicInfo.EntityData.Leafs.Append("physical-layer-interface-module-type", types.YLeaf{"PhysicalLayerInterfaceModuleType", basicInfo.PhysicalLayerInterfaceModuleType})
    basicInfo.EntityData.Leafs.Append("unrecognized-fru", types.YLeaf{"UnrecognizedFru", basicInfo.UnrecognizedFru})
    basicInfo.EntityData.Leafs.Append("redundancystate", types.YLeaf{"Redundancystate", basicInfo.Redundancystate})
    basicInfo.EntityData.Leafs.Append("ceport", types.YLeaf{"Ceport", basicInfo.Ceport})
    basicInfo.EntityData.Leafs.Append("xr-scoped", types.YLeaf{"XrScoped", basicInfo.XrScoped})
    basicInfo.EntityData.Leafs.Append("unique-id", types.YLeaf{"UniqueId", basicInfo.UniqueId})

    basicInfo.EntityData.YListKeys = []string {}

    return &(basicInfo.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_FruInfo
// Field Replaceable Unit (FRU) inventory
// information
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_FruInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Administrative    state. The type is InvAdminState.
    ModuleAdministrativeState interface{}

    // Power administrative state. The type is InvPowerAdminState.
    ModulePowerAdministrativeState interface{}

    // Operation state. The type is InvCardState.
    ModuleOperationalState interface{}

    // Monitor state. The type is InvMonitorState.
    ModuleMonitorState interface{}

    // Reset reason. The type is InvResetReason.
    ModuleResetReason interface{}

    // Time operational state is   last changed.
    LastOperationalStateChange Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_FruInfo_LastOperationalStateChange

    // Module up time.
    ModuleUpTime Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_FruInfo_ModuleUpTime
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_FruInfo) GetEntityData() *types.CommonEntityData {
    fruInfo.EntityData.YFilter = fruInfo.YFilter
    fruInfo.EntityData.YangName = "fru-info"
    fruInfo.EntityData.BundleName = "cisco_ios_xr"
    fruInfo.EntityData.ParentYangName = "basic-attributes"
    fruInfo.EntityData.SegmentPath = "fru-info"
    fruInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/cards/card/hw-components/hw-component/basic-attributes/" + fruInfo.EntityData.SegmentPath
    fruInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fruInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fruInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fruInfo.EntityData.Children = types.NewOrderedMap()
    fruInfo.EntityData.Children.Append("last-operational-state-change", types.YChild{"LastOperationalStateChange", &fruInfo.LastOperationalStateChange})
    fruInfo.EntityData.Children.Append("module-up-time", types.YChild{"ModuleUpTime", &fruInfo.ModuleUpTime})
    fruInfo.EntityData.Leafs = types.NewOrderedMap()
    fruInfo.EntityData.Leafs.Append("module-administrative-state", types.YLeaf{"ModuleAdministrativeState", fruInfo.ModuleAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-power-administrative-state", types.YLeaf{"ModulePowerAdministrativeState", fruInfo.ModulePowerAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-operational-state", types.YLeaf{"ModuleOperationalState", fruInfo.ModuleOperationalState})
    fruInfo.EntityData.Leafs.Append("module-monitor-state", types.YLeaf{"ModuleMonitorState", fruInfo.ModuleMonitorState})
    fruInfo.EntityData.Leafs.Append("module-reset-reason", types.YLeaf{"ModuleResetReason", fruInfo.ModuleResetReason})

    fruInfo.EntityData.YListKeys = []string {}

    return &(fruInfo.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_FruInfo_LastOperationalStateChange
// Time operational state is   last changed
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
    lastOperationalStateChange.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/cards/card/hw-components/hw-component/basic-attributes/fru-info/" + lastOperationalStateChange.EntityData.SegmentPath
    lastOperationalStateChange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastOperationalStateChange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastOperationalStateChange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastOperationalStateChange.EntityData.Children = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", lastOperationalStateChange.TimeInSeconds})
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", lastOperationalStateChange.TimeInNanoSeconds})

    lastOperationalStateChange.EntityData.YListKeys = []string {}

    return &(lastOperationalStateChange.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_FruInfo_ModuleUpTime
// Module up time
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_FruInfo_ModuleUpTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (moduleUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_FruInfo_ModuleUpTime) GetEntityData() *types.CommonEntityData {
    moduleUpTime.EntityData.YFilter = moduleUpTime.YFilter
    moduleUpTime.EntityData.YangName = "module-up-time"
    moduleUpTime.EntityData.BundleName = "cisco_ios_xr"
    moduleUpTime.EntityData.ParentYangName = "fru-info"
    moduleUpTime.EntityData.SegmentPath = "module-up-time"
    moduleUpTime.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/cards/card/hw-components/hw-component/basic-attributes/fru-info/" + moduleUpTime.EntityData.SegmentPath
    moduleUpTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    moduleUpTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    moduleUpTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    moduleUpTime.EntityData.Children = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", moduleUpTime.TimeInSeconds})
    moduleUpTime.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", moduleUpTime.TimeInNanoSeconds})

    moduleUpTime.EntityData.YListKeys = []string {}

    return &(moduleUpTime.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors
// ModuleSensorTable contains all sensors in a
// Module.
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sensor number in the Module. The type is slice of
    // Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor.
    Sensor []*Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor
}

func (sensors *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors) GetEntityData() *types.CommonEntityData {
    sensors.EntityData.YFilter = sensors.YFilter
    sensors.EntityData.YangName = "sensors"
    sensors.EntityData.BundleName = "cisco_ios_xr"
    sensors.EntityData.ParentYangName = "card"
    sensors.EntityData.SegmentPath = "sensors"
    sensors.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/cards/card/" + sensors.EntityData.SegmentPath
    sensors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sensors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sensors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sensors.EntityData.Children = types.NewOrderedMap()
    sensors.EntityData.Children.Append("sensor", types.YChild{"Sensor", nil})
    for i := range sensors.Sensor {
        sensors.EntityData.Children.Append(types.GetSegmentPath(sensors.Sensor[i]), types.YChild{"Sensor", sensors.Sensor[i]})
    }
    sensors.EntityData.Leafs = types.NewOrderedMap()

    sensors.EntityData.YListKeys = []string {}

    return &(sensors.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor
// Sensor number in the Module
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. sensor number. The type is interface{} with range:
    // 0..4294967295.
    Number interface{}

    // Attributes.
    BasicAttributes Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes
}

func (sensor *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor) GetEntityData() *types.CommonEntityData {
    sensor.EntityData.YFilter = sensor.YFilter
    sensor.EntityData.YangName = "sensor"
    sensor.EntityData.BundleName = "cisco_ios_xr"
    sensor.EntityData.ParentYangName = "sensors"
    sensor.EntityData.SegmentPath = "sensor" + types.AddKeyToken(sensor.Number, "number")
    sensor.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/cards/card/sensors/" + sensor.EntityData.SegmentPath
    sensor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sensor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sensor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sensor.EntityData.Children = types.NewOrderedMap()
    sensor.EntityData.Children.Append("basic-attributes", types.YChild{"BasicAttributes", &sensor.BasicAttributes})
    sensor.EntityData.Leafs = types.NewOrderedMap()
    sensor.EntityData.Leafs.Append("number", types.YLeaf{"Number", sensor.Number})

    sensor.EntityData.YListKeys = []string {"Number"}

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
    basicAttributes.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/cards/card/sensors/sensor/" + basicAttributes.EntityData.SegmentPath
    basicAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicAttributes.EntityData.Children = types.NewOrderedMap()
    basicAttributes.EntityData.Children.Append("basic-info", types.YChild{"BasicInfo", &basicAttributes.BasicInfo})
    basicAttributes.EntityData.Children.Append("fru-info", types.YChild{"FruInfo", &basicAttributes.FruInfo})
    basicAttributes.EntityData.Leafs = types.NewOrderedMap()

    basicAttributes.EntityData.YListKeys = []string {}

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

    // integer value for Redundancy State if applicable to this entity. The type
    // is interface{} with range: -2147483648..2147483647.
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
    basicInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/cards/card/sensors/sensor/basic-attributes/" + basicInfo.EntityData.SegmentPath
    basicInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicInfo.EntityData.Children = types.NewOrderedMap()
    basicInfo.EntityData.Leafs = types.NewOrderedMap()
    basicInfo.EntityData.Leafs.Append("description", types.YLeaf{"Description", basicInfo.Description})
    basicInfo.EntityData.Leafs.Append("vendor-type", types.YLeaf{"VendorType", basicInfo.VendorType})
    basicInfo.EntityData.Leafs.Append("name", types.YLeaf{"Name", basicInfo.Name})
    basicInfo.EntityData.Leafs.Append("hardware-revision", types.YLeaf{"HardwareRevision", basicInfo.HardwareRevision})
    basicInfo.EntityData.Leafs.Append("firmware-revision", types.YLeaf{"FirmwareRevision", basicInfo.FirmwareRevision})
    basicInfo.EntityData.Leafs.Append("software-revision", types.YLeaf{"SoftwareRevision", basicInfo.SoftwareRevision})
    basicInfo.EntityData.Leafs.Append("chip-hardware-revision", types.YLeaf{"ChipHardwareRevision", basicInfo.ChipHardwareRevision})
    basicInfo.EntityData.Leafs.Append("serial-number", types.YLeaf{"SerialNumber", basicInfo.SerialNumber})
    basicInfo.EntityData.Leafs.Append("manufacturer-name", types.YLeaf{"ManufacturerName", basicInfo.ManufacturerName})
    basicInfo.EntityData.Leafs.Append("model-name", types.YLeaf{"ModelName", basicInfo.ModelName})
    basicInfo.EntityData.Leafs.Append("asset-id-str", types.YLeaf{"AssetIdStr", basicInfo.AssetIdStr})
    basicInfo.EntityData.Leafs.Append("asset-identification", types.YLeaf{"AssetIdentification", basicInfo.AssetIdentification})
    basicInfo.EntityData.Leafs.Append("is-field-replaceable-unit", types.YLeaf{"IsFieldReplaceableUnit", basicInfo.IsFieldReplaceableUnit})
    basicInfo.EntityData.Leafs.Append("manufacturer-asset-tags", types.YLeaf{"ManufacturerAssetTags", basicInfo.ManufacturerAssetTags})
    basicInfo.EntityData.Leafs.Append("composite-class-code", types.YLeaf{"CompositeClassCode", basicInfo.CompositeClassCode})
    basicInfo.EntityData.Leafs.Append("memory-size", types.YLeaf{"MemorySize", basicInfo.MemorySize})
    basicInfo.EntityData.Leafs.Append("environmental-monitor-path", types.YLeaf{"EnvironmentalMonitorPath", basicInfo.EnvironmentalMonitorPath})
    basicInfo.EntityData.Leafs.Append("alias", types.YLeaf{"Alias", basicInfo.Alias})
    basicInfo.EntityData.Leafs.Append("group-flag", types.YLeaf{"GroupFlag", basicInfo.GroupFlag})
    basicInfo.EntityData.Leafs.Append("new-deviation-number", types.YLeaf{"NewDeviationNumber", basicInfo.NewDeviationNumber})
    basicInfo.EntityData.Leafs.Append("physical-layer-interface-module-type", types.YLeaf{"PhysicalLayerInterfaceModuleType", basicInfo.PhysicalLayerInterfaceModuleType})
    basicInfo.EntityData.Leafs.Append("unrecognized-fru", types.YLeaf{"UnrecognizedFru", basicInfo.UnrecognizedFru})
    basicInfo.EntityData.Leafs.Append("redundancystate", types.YLeaf{"Redundancystate", basicInfo.Redundancystate})
    basicInfo.EntityData.Leafs.Append("ceport", types.YLeaf{"Ceport", basicInfo.Ceport})
    basicInfo.EntityData.Leafs.Append("xr-scoped", types.YLeaf{"XrScoped", basicInfo.XrScoped})
    basicInfo.EntityData.Leafs.Append("unique-id", types.YLeaf{"UniqueId", basicInfo.UniqueId})

    basicInfo.EntityData.YListKeys = []string {}

    return &(basicInfo.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_FruInfo
// Field Replaceable Unit (FRU) inventory
// information
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_FruInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Administrative    state. The type is InvAdminState.
    ModuleAdministrativeState interface{}

    // Power administrative state. The type is InvPowerAdminState.
    ModulePowerAdministrativeState interface{}

    // Operation state. The type is InvCardState.
    ModuleOperationalState interface{}

    // Monitor state. The type is InvMonitorState.
    ModuleMonitorState interface{}

    // Reset reason. The type is InvResetReason.
    ModuleResetReason interface{}

    // Time operational state is   last changed.
    LastOperationalStateChange Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_FruInfo_LastOperationalStateChange

    // Module up time.
    ModuleUpTime Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_FruInfo_ModuleUpTime
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_FruInfo) GetEntityData() *types.CommonEntityData {
    fruInfo.EntityData.YFilter = fruInfo.YFilter
    fruInfo.EntityData.YangName = "fru-info"
    fruInfo.EntityData.BundleName = "cisco_ios_xr"
    fruInfo.EntityData.ParentYangName = "basic-attributes"
    fruInfo.EntityData.SegmentPath = "fru-info"
    fruInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/cards/card/sensors/sensor/basic-attributes/" + fruInfo.EntityData.SegmentPath
    fruInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fruInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fruInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fruInfo.EntityData.Children = types.NewOrderedMap()
    fruInfo.EntityData.Children.Append("last-operational-state-change", types.YChild{"LastOperationalStateChange", &fruInfo.LastOperationalStateChange})
    fruInfo.EntityData.Children.Append("module-up-time", types.YChild{"ModuleUpTime", &fruInfo.ModuleUpTime})
    fruInfo.EntityData.Leafs = types.NewOrderedMap()
    fruInfo.EntityData.Leafs.Append("module-administrative-state", types.YLeaf{"ModuleAdministrativeState", fruInfo.ModuleAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-power-administrative-state", types.YLeaf{"ModulePowerAdministrativeState", fruInfo.ModulePowerAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-operational-state", types.YLeaf{"ModuleOperationalState", fruInfo.ModuleOperationalState})
    fruInfo.EntityData.Leafs.Append("module-monitor-state", types.YLeaf{"ModuleMonitorState", fruInfo.ModuleMonitorState})
    fruInfo.EntityData.Leafs.Append("module-reset-reason", types.YLeaf{"ModuleResetReason", fruInfo.ModuleResetReason})

    fruInfo.EntityData.YListKeys = []string {}

    return &(fruInfo.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_FruInfo_LastOperationalStateChange
// Time operational state is   last changed
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
    lastOperationalStateChange.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/cards/card/sensors/sensor/basic-attributes/fru-info/" + lastOperationalStateChange.EntityData.SegmentPath
    lastOperationalStateChange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastOperationalStateChange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastOperationalStateChange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastOperationalStateChange.EntityData.Children = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", lastOperationalStateChange.TimeInSeconds})
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", lastOperationalStateChange.TimeInNanoSeconds})

    lastOperationalStateChange.EntityData.YListKeys = []string {}

    return &(lastOperationalStateChange.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_FruInfo_ModuleUpTime
// Module up time
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_FruInfo_ModuleUpTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (moduleUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_FruInfo_ModuleUpTime) GetEntityData() *types.CommonEntityData {
    moduleUpTime.EntityData.YFilter = moduleUpTime.YFilter
    moduleUpTime.EntityData.YangName = "module-up-time"
    moduleUpTime.EntityData.BundleName = "cisco_ios_xr"
    moduleUpTime.EntityData.ParentYangName = "fru-info"
    moduleUpTime.EntityData.SegmentPath = "module-up-time"
    moduleUpTime.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/cards/card/sensors/sensor/basic-attributes/fru-info/" + moduleUpTime.EntityData.SegmentPath
    moduleUpTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    moduleUpTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    moduleUpTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    moduleUpTime.EntityData.Children = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", moduleUpTime.TimeInSeconds})
    moduleUpTime.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", moduleUpTime.TimeInNanoSeconds})

    moduleUpTime.EntityData.YListKeys = []string {}

    return &(moduleUpTime.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots
// PortSlotTable contains all optics ports in a
// SPA/PLIM.
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PortSlot number in the SPA/PLIM. The type is slice of
    // Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot.
    PortSlot []*Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot
}

func (portSlots *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots) GetEntityData() *types.CommonEntityData {
    portSlots.EntityData.YFilter = portSlots.YFilter
    portSlots.EntityData.YangName = "port-slots"
    portSlots.EntityData.BundleName = "cisco_ios_xr"
    portSlots.EntityData.ParentYangName = "card"
    portSlots.EntityData.SegmentPath = "port-slots"
    portSlots.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/cards/card/" + portSlots.EntityData.SegmentPath
    portSlots.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portSlots.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portSlots.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portSlots.EntityData.Children = types.NewOrderedMap()
    portSlots.EntityData.Children.Append("port-slot", types.YChild{"PortSlot", nil})
    for i := range portSlots.PortSlot {
        portSlots.EntityData.Children.Append(types.GetSegmentPath(portSlots.PortSlot[i]), types.YChild{"PortSlot", portSlots.PortSlot[i]})
    }
    portSlots.EntityData.Leafs = types.NewOrderedMap()

    portSlots.EntityData.YListKeys = []string {}

    return &(portSlots.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot
// PortSlot number in the SPA/PLIM
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. portslot number. The type is interface{} with
    // range: 0..4294967295.
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
    portSlot.EntityData.SegmentPath = "port-slot" + types.AddKeyToken(portSlot.Number, "number")
    portSlot.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/cards/card/port-slots/" + portSlot.EntityData.SegmentPath
    portSlot.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portSlot.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portSlot.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portSlot.EntityData.Children = types.NewOrderedMap()
    portSlot.EntityData.Children.Append("port", types.YChild{"Port", &portSlot.Port})
    portSlot.EntityData.Children.Append("basic-attributes", types.YChild{"BasicAttributes", &portSlot.BasicAttributes})
    portSlot.EntityData.Leafs = types.NewOrderedMap()
    portSlot.EntityData.Leafs.Append("number", types.YLeaf{"Number", portSlot.Number})

    portSlot.EntityData.YListKeys = []string {"Number"}

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
    port.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/cards/card/port-slots/port-slot/" + port.EntityData.SegmentPath
    port.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    port.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    port.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    port.EntityData.Children = types.NewOrderedMap()
    port.EntityData.Children.Append("basic-attributes", types.YChild{"BasicAttributes", &port.BasicAttributes})
    port.EntityData.Leafs = types.NewOrderedMap()

    port.EntityData.YListKeys = []string {}

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
    basicAttributes.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/cards/card/port-slots/port-slot/port/" + basicAttributes.EntityData.SegmentPath
    basicAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicAttributes.EntityData.Children = types.NewOrderedMap()
    basicAttributes.EntityData.Children.Append("basic-info", types.YChild{"BasicInfo", &basicAttributes.BasicInfo})
    basicAttributes.EntityData.Children.Append("fru-info", types.YChild{"FruInfo", &basicAttributes.FruInfo})
    basicAttributes.EntityData.Leafs = types.NewOrderedMap()

    basicAttributes.EntityData.YListKeys = []string {}

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

    // integer value for Redundancy State if applicable to this entity. The type
    // is interface{} with range: -2147483648..2147483647.
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
    basicInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/cards/card/port-slots/port-slot/port/basic-attributes/" + basicInfo.EntityData.SegmentPath
    basicInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicInfo.EntityData.Children = types.NewOrderedMap()
    basicInfo.EntityData.Leafs = types.NewOrderedMap()
    basicInfo.EntityData.Leafs.Append("description", types.YLeaf{"Description", basicInfo.Description})
    basicInfo.EntityData.Leafs.Append("vendor-type", types.YLeaf{"VendorType", basicInfo.VendorType})
    basicInfo.EntityData.Leafs.Append("name", types.YLeaf{"Name", basicInfo.Name})
    basicInfo.EntityData.Leafs.Append("hardware-revision", types.YLeaf{"HardwareRevision", basicInfo.HardwareRevision})
    basicInfo.EntityData.Leafs.Append("firmware-revision", types.YLeaf{"FirmwareRevision", basicInfo.FirmwareRevision})
    basicInfo.EntityData.Leafs.Append("software-revision", types.YLeaf{"SoftwareRevision", basicInfo.SoftwareRevision})
    basicInfo.EntityData.Leafs.Append("chip-hardware-revision", types.YLeaf{"ChipHardwareRevision", basicInfo.ChipHardwareRevision})
    basicInfo.EntityData.Leafs.Append("serial-number", types.YLeaf{"SerialNumber", basicInfo.SerialNumber})
    basicInfo.EntityData.Leafs.Append("manufacturer-name", types.YLeaf{"ManufacturerName", basicInfo.ManufacturerName})
    basicInfo.EntityData.Leafs.Append("model-name", types.YLeaf{"ModelName", basicInfo.ModelName})
    basicInfo.EntityData.Leafs.Append("asset-id-str", types.YLeaf{"AssetIdStr", basicInfo.AssetIdStr})
    basicInfo.EntityData.Leafs.Append("asset-identification", types.YLeaf{"AssetIdentification", basicInfo.AssetIdentification})
    basicInfo.EntityData.Leafs.Append("is-field-replaceable-unit", types.YLeaf{"IsFieldReplaceableUnit", basicInfo.IsFieldReplaceableUnit})
    basicInfo.EntityData.Leafs.Append("manufacturer-asset-tags", types.YLeaf{"ManufacturerAssetTags", basicInfo.ManufacturerAssetTags})
    basicInfo.EntityData.Leafs.Append("composite-class-code", types.YLeaf{"CompositeClassCode", basicInfo.CompositeClassCode})
    basicInfo.EntityData.Leafs.Append("memory-size", types.YLeaf{"MemorySize", basicInfo.MemorySize})
    basicInfo.EntityData.Leafs.Append("environmental-monitor-path", types.YLeaf{"EnvironmentalMonitorPath", basicInfo.EnvironmentalMonitorPath})
    basicInfo.EntityData.Leafs.Append("alias", types.YLeaf{"Alias", basicInfo.Alias})
    basicInfo.EntityData.Leafs.Append("group-flag", types.YLeaf{"GroupFlag", basicInfo.GroupFlag})
    basicInfo.EntityData.Leafs.Append("new-deviation-number", types.YLeaf{"NewDeviationNumber", basicInfo.NewDeviationNumber})
    basicInfo.EntityData.Leafs.Append("physical-layer-interface-module-type", types.YLeaf{"PhysicalLayerInterfaceModuleType", basicInfo.PhysicalLayerInterfaceModuleType})
    basicInfo.EntityData.Leafs.Append("unrecognized-fru", types.YLeaf{"UnrecognizedFru", basicInfo.UnrecognizedFru})
    basicInfo.EntityData.Leafs.Append("redundancystate", types.YLeaf{"Redundancystate", basicInfo.Redundancystate})
    basicInfo.EntityData.Leafs.Append("ceport", types.YLeaf{"Ceport", basicInfo.Ceport})
    basicInfo.EntityData.Leafs.Append("xr-scoped", types.YLeaf{"XrScoped", basicInfo.XrScoped})
    basicInfo.EntityData.Leafs.Append("unique-id", types.YLeaf{"UniqueId", basicInfo.UniqueId})

    basicInfo.EntityData.YListKeys = []string {}

    return &(basicInfo.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_FruInfo
// Field Replaceable Unit (FRU) inventory
// information
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_FruInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Administrative    state. The type is InvAdminState.
    ModuleAdministrativeState interface{}

    // Power administrative state. The type is InvPowerAdminState.
    ModulePowerAdministrativeState interface{}

    // Operation state. The type is InvCardState.
    ModuleOperationalState interface{}

    // Monitor state. The type is InvMonitorState.
    ModuleMonitorState interface{}

    // Reset reason. The type is InvResetReason.
    ModuleResetReason interface{}

    // Time operational state is   last changed.
    LastOperationalStateChange Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_LastOperationalStateChange

    // Module up time.
    ModuleUpTime Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_ModuleUpTime
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_FruInfo) GetEntityData() *types.CommonEntityData {
    fruInfo.EntityData.YFilter = fruInfo.YFilter
    fruInfo.EntityData.YangName = "fru-info"
    fruInfo.EntityData.BundleName = "cisco_ios_xr"
    fruInfo.EntityData.ParentYangName = "basic-attributes"
    fruInfo.EntityData.SegmentPath = "fru-info"
    fruInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/cards/card/port-slots/port-slot/port/basic-attributes/" + fruInfo.EntityData.SegmentPath
    fruInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fruInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fruInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fruInfo.EntityData.Children = types.NewOrderedMap()
    fruInfo.EntityData.Children.Append("last-operational-state-change", types.YChild{"LastOperationalStateChange", &fruInfo.LastOperationalStateChange})
    fruInfo.EntityData.Children.Append("module-up-time", types.YChild{"ModuleUpTime", &fruInfo.ModuleUpTime})
    fruInfo.EntityData.Leafs = types.NewOrderedMap()
    fruInfo.EntityData.Leafs.Append("module-administrative-state", types.YLeaf{"ModuleAdministrativeState", fruInfo.ModuleAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-power-administrative-state", types.YLeaf{"ModulePowerAdministrativeState", fruInfo.ModulePowerAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-operational-state", types.YLeaf{"ModuleOperationalState", fruInfo.ModuleOperationalState})
    fruInfo.EntityData.Leafs.Append("module-monitor-state", types.YLeaf{"ModuleMonitorState", fruInfo.ModuleMonitorState})
    fruInfo.EntityData.Leafs.Append("module-reset-reason", types.YLeaf{"ModuleResetReason", fruInfo.ModuleResetReason})

    fruInfo.EntityData.YListKeys = []string {}

    return &(fruInfo.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_LastOperationalStateChange
// Time operational state is   last changed
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
    lastOperationalStateChange.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/cards/card/port-slots/port-slot/port/basic-attributes/fru-info/" + lastOperationalStateChange.EntityData.SegmentPath
    lastOperationalStateChange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastOperationalStateChange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastOperationalStateChange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastOperationalStateChange.EntityData.Children = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", lastOperationalStateChange.TimeInSeconds})
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", lastOperationalStateChange.TimeInNanoSeconds})

    lastOperationalStateChange.EntityData.YListKeys = []string {}

    return &(lastOperationalStateChange.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_ModuleUpTime
// Module up time
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_ModuleUpTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (moduleUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_ModuleUpTime) GetEntityData() *types.CommonEntityData {
    moduleUpTime.EntityData.YFilter = moduleUpTime.YFilter
    moduleUpTime.EntityData.YangName = "module-up-time"
    moduleUpTime.EntityData.BundleName = "cisco_ios_xr"
    moduleUpTime.EntityData.ParentYangName = "fru-info"
    moduleUpTime.EntityData.SegmentPath = "module-up-time"
    moduleUpTime.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/cards/card/port-slots/port-slot/port/basic-attributes/fru-info/" + moduleUpTime.EntityData.SegmentPath
    moduleUpTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    moduleUpTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    moduleUpTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    moduleUpTime.EntityData.Children = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", moduleUpTime.TimeInSeconds})
    moduleUpTime.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", moduleUpTime.TimeInNanoSeconds})

    moduleUpTime.EntityData.YListKeys = []string {}

    return &(moduleUpTime.EntityData)
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
    basicAttributes.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/cards/card/port-slots/port-slot/" + basicAttributes.EntityData.SegmentPath
    basicAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicAttributes.EntityData.Children = types.NewOrderedMap()
    basicAttributes.EntityData.Children.Append("basic-info", types.YChild{"BasicInfo", &basicAttributes.BasicInfo})
    basicAttributes.EntityData.Children.Append("fru-info", types.YChild{"FruInfo", &basicAttributes.FruInfo})
    basicAttributes.EntityData.Leafs = types.NewOrderedMap()

    basicAttributes.EntityData.YListKeys = []string {}

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

    // integer value for Redundancy State if applicable to this entity. The type
    // is interface{} with range: -2147483648..2147483647.
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
    basicInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/cards/card/port-slots/port-slot/basic-attributes/" + basicInfo.EntityData.SegmentPath
    basicInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicInfo.EntityData.Children = types.NewOrderedMap()
    basicInfo.EntityData.Leafs = types.NewOrderedMap()
    basicInfo.EntityData.Leafs.Append("description", types.YLeaf{"Description", basicInfo.Description})
    basicInfo.EntityData.Leafs.Append("vendor-type", types.YLeaf{"VendorType", basicInfo.VendorType})
    basicInfo.EntityData.Leafs.Append("name", types.YLeaf{"Name", basicInfo.Name})
    basicInfo.EntityData.Leafs.Append("hardware-revision", types.YLeaf{"HardwareRevision", basicInfo.HardwareRevision})
    basicInfo.EntityData.Leafs.Append("firmware-revision", types.YLeaf{"FirmwareRevision", basicInfo.FirmwareRevision})
    basicInfo.EntityData.Leafs.Append("software-revision", types.YLeaf{"SoftwareRevision", basicInfo.SoftwareRevision})
    basicInfo.EntityData.Leafs.Append("chip-hardware-revision", types.YLeaf{"ChipHardwareRevision", basicInfo.ChipHardwareRevision})
    basicInfo.EntityData.Leafs.Append("serial-number", types.YLeaf{"SerialNumber", basicInfo.SerialNumber})
    basicInfo.EntityData.Leafs.Append("manufacturer-name", types.YLeaf{"ManufacturerName", basicInfo.ManufacturerName})
    basicInfo.EntityData.Leafs.Append("model-name", types.YLeaf{"ModelName", basicInfo.ModelName})
    basicInfo.EntityData.Leafs.Append("asset-id-str", types.YLeaf{"AssetIdStr", basicInfo.AssetIdStr})
    basicInfo.EntityData.Leafs.Append("asset-identification", types.YLeaf{"AssetIdentification", basicInfo.AssetIdentification})
    basicInfo.EntityData.Leafs.Append("is-field-replaceable-unit", types.YLeaf{"IsFieldReplaceableUnit", basicInfo.IsFieldReplaceableUnit})
    basicInfo.EntityData.Leafs.Append("manufacturer-asset-tags", types.YLeaf{"ManufacturerAssetTags", basicInfo.ManufacturerAssetTags})
    basicInfo.EntityData.Leafs.Append("composite-class-code", types.YLeaf{"CompositeClassCode", basicInfo.CompositeClassCode})
    basicInfo.EntityData.Leafs.Append("memory-size", types.YLeaf{"MemorySize", basicInfo.MemorySize})
    basicInfo.EntityData.Leafs.Append("environmental-monitor-path", types.YLeaf{"EnvironmentalMonitorPath", basicInfo.EnvironmentalMonitorPath})
    basicInfo.EntityData.Leafs.Append("alias", types.YLeaf{"Alias", basicInfo.Alias})
    basicInfo.EntityData.Leafs.Append("group-flag", types.YLeaf{"GroupFlag", basicInfo.GroupFlag})
    basicInfo.EntityData.Leafs.Append("new-deviation-number", types.YLeaf{"NewDeviationNumber", basicInfo.NewDeviationNumber})
    basicInfo.EntityData.Leafs.Append("physical-layer-interface-module-type", types.YLeaf{"PhysicalLayerInterfaceModuleType", basicInfo.PhysicalLayerInterfaceModuleType})
    basicInfo.EntityData.Leafs.Append("unrecognized-fru", types.YLeaf{"UnrecognizedFru", basicInfo.UnrecognizedFru})
    basicInfo.EntityData.Leafs.Append("redundancystate", types.YLeaf{"Redundancystate", basicInfo.Redundancystate})
    basicInfo.EntityData.Leafs.Append("ceport", types.YLeaf{"Ceport", basicInfo.Ceport})
    basicInfo.EntityData.Leafs.Append("xr-scoped", types.YLeaf{"XrScoped", basicInfo.XrScoped})
    basicInfo.EntityData.Leafs.Append("unique-id", types.YLeaf{"UniqueId", basicInfo.UniqueId})

    basicInfo.EntityData.YListKeys = []string {}

    return &(basicInfo.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_FruInfo
// Field Replaceable Unit (FRU) inventory
// information
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_FruInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Administrative    state. The type is InvAdminState.
    ModuleAdministrativeState interface{}

    // Power administrative state. The type is InvPowerAdminState.
    ModulePowerAdministrativeState interface{}

    // Operation state. The type is InvCardState.
    ModuleOperationalState interface{}

    // Monitor state. The type is InvMonitorState.
    ModuleMonitorState interface{}

    // Reset reason. The type is InvResetReason.
    ModuleResetReason interface{}

    // Time operational state is   last changed.
    LastOperationalStateChange Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_FruInfo_LastOperationalStateChange

    // Module up time.
    ModuleUpTime Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_FruInfo_ModuleUpTime
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_FruInfo) GetEntityData() *types.CommonEntityData {
    fruInfo.EntityData.YFilter = fruInfo.YFilter
    fruInfo.EntityData.YangName = "fru-info"
    fruInfo.EntityData.BundleName = "cisco_ios_xr"
    fruInfo.EntityData.ParentYangName = "basic-attributes"
    fruInfo.EntityData.SegmentPath = "fru-info"
    fruInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/cards/card/port-slots/port-slot/basic-attributes/" + fruInfo.EntityData.SegmentPath
    fruInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fruInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fruInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fruInfo.EntityData.Children = types.NewOrderedMap()
    fruInfo.EntityData.Children.Append("last-operational-state-change", types.YChild{"LastOperationalStateChange", &fruInfo.LastOperationalStateChange})
    fruInfo.EntityData.Children.Append("module-up-time", types.YChild{"ModuleUpTime", &fruInfo.ModuleUpTime})
    fruInfo.EntityData.Leafs = types.NewOrderedMap()
    fruInfo.EntityData.Leafs.Append("module-administrative-state", types.YLeaf{"ModuleAdministrativeState", fruInfo.ModuleAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-power-administrative-state", types.YLeaf{"ModulePowerAdministrativeState", fruInfo.ModulePowerAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-operational-state", types.YLeaf{"ModuleOperationalState", fruInfo.ModuleOperationalState})
    fruInfo.EntityData.Leafs.Append("module-monitor-state", types.YLeaf{"ModuleMonitorState", fruInfo.ModuleMonitorState})
    fruInfo.EntityData.Leafs.Append("module-reset-reason", types.YLeaf{"ModuleResetReason", fruInfo.ModuleResetReason})

    fruInfo.EntityData.YListKeys = []string {}

    return &(fruInfo.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_FruInfo_LastOperationalStateChange
// Time operational state is   last changed
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
    lastOperationalStateChange.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/cards/card/port-slots/port-slot/basic-attributes/fru-info/" + lastOperationalStateChange.EntityData.SegmentPath
    lastOperationalStateChange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastOperationalStateChange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastOperationalStateChange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastOperationalStateChange.EntityData.Children = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", lastOperationalStateChange.TimeInSeconds})
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", lastOperationalStateChange.TimeInNanoSeconds})

    lastOperationalStateChange.EntityData.YListKeys = []string {}

    return &(lastOperationalStateChange.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_FruInfo_ModuleUpTime
// Module up time
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_FruInfo_ModuleUpTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (moduleUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_FruInfo_ModuleUpTime) GetEntityData() *types.CommonEntityData {
    moduleUpTime.EntityData.YFilter = moduleUpTime.YFilter
    moduleUpTime.EntityData.YangName = "module-up-time"
    moduleUpTime.EntityData.BundleName = "cisco_ios_xr"
    moduleUpTime.EntityData.ParentYangName = "fru-info"
    moduleUpTime.EntityData.SegmentPath = "module-up-time"
    moduleUpTime.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/cards/card/port-slots/port-slot/basic-attributes/fru-info/" + moduleUpTime.EntityData.SegmentPath
    moduleUpTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    moduleUpTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    moduleUpTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    moduleUpTime.EntityData.Children = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", moduleUpTime.TimeInSeconds})
    moduleUpTime.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", moduleUpTime.TimeInNanoSeconds})

    moduleUpTime.EntityData.YListKeys = []string {}

    return &(moduleUpTime.EntityData)
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
    basicAttributes.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/cards/card/" + basicAttributes.EntityData.SegmentPath
    basicAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicAttributes.EntityData.Children = types.NewOrderedMap()
    basicAttributes.EntityData.Children.Append("basic-info", types.YChild{"BasicInfo", &basicAttributes.BasicInfo})
    basicAttributes.EntityData.Children.Append("fru-info", types.YChild{"FruInfo", &basicAttributes.FruInfo})
    basicAttributes.EntityData.Leafs = types.NewOrderedMap()

    basicAttributes.EntityData.YListKeys = []string {}

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

    // integer value for Redundancy State if applicable to this entity. The type
    // is interface{} with range: -2147483648..2147483647.
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
    basicInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/cards/card/basic-attributes/" + basicInfo.EntityData.SegmentPath
    basicInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicInfo.EntityData.Children = types.NewOrderedMap()
    basicInfo.EntityData.Leafs = types.NewOrderedMap()
    basicInfo.EntityData.Leafs.Append("description", types.YLeaf{"Description", basicInfo.Description})
    basicInfo.EntityData.Leafs.Append("vendor-type", types.YLeaf{"VendorType", basicInfo.VendorType})
    basicInfo.EntityData.Leafs.Append("name", types.YLeaf{"Name", basicInfo.Name})
    basicInfo.EntityData.Leafs.Append("hardware-revision", types.YLeaf{"HardwareRevision", basicInfo.HardwareRevision})
    basicInfo.EntityData.Leafs.Append("firmware-revision", types.YLeaf{"FirmwareRevision", basicInfo.FirmwareRevision})
    basicInfo.EntityData.Leafs.Append("software-revision", types.YLeaf{"SoftwareRevision", basicInfo.SoftwareRevision})
    basicInfo.EntityData.Leafs.Append("chip-hardware-revision", types.YLeaf{"ChipHardwareRevision", basicInfo.ChipHardwareRevision})
    basicInfo.EntityData.Leafs.Append("serial-number", types.YLeaf{"SerialNumber", basicInfo.SerialNumber})
    basicInfo.EntityData.Leafs.Append("manufacturer-name", types.YLeaf{"ManufacturerName", basicInfo.ManufacturerName})
    basicInfo.EntityData.Leafs.Append("model-name", types.YLeaf{"ModelName", basicInfo.ModelName})
    basicInfo.EntityData.Leafs.Append("asset-id-str", types.YLeaf{"AssetIdStr", basicInfo.AssetIdStr})
    basicInfo.EntityData.Leafs.Append("asset-identification", types.YLeaf{"AssetIdentification", basicInfo.AssetIdentification})
    basicInfo.EntityData.Leafs.Append("is-field-replaceable-unit", types.YLeaf{"IsFieldReplaceableUnit", basicInfo.IsFieldReplaceableUnit})
    basicInfo.EntityData.Leafs.Append("manufacturer-asset-tags", types.YLeaf{"ManufacturerAssetTags", basicInfo.ManufacturerAssetTags})
    basicInfo.EntityData.Leafs.Append("composite-class-code", types.YLeaf{"CompositeClassCode", basicInfo.CompositeClassCode})
    basicInfo.EntityData.Leafs.Append("memory-size", types.YLeaf{"MemorySize", basicInfo.MemorySize})
    basicInfo.EntityData.Leafs.Append("environmental-monitor-path", types.YLeaf{"EnvironmentalMonitorPath", basicInfo.EnvironmentalMonitorPath})
    basicInfo.EntityData.Leafs.Append("alias", types.YLeaf{"Alias", basicInfo.Alias})
    basicInfo.EntityData.Leafs.Append("group-flag", types.YLeaf{"GroupFlag", basicInfo.GroupFlag})
    basicInfo.EntityData.Leafs.Append("new-deviation-number", types.YLeaf{"NewDeviationNumber", basicInfo.NewDeviationNumber})
    basicInfo.EntityData.Leafs.Append("physical-layer-interface-module-type", types.YLeaf{"PhysicalLayerInterfaceModuleType", basicInfo.PhysicalLayerInterfaceModuleType})
    basicInfo.EntityData.Leafs.Append("unrecognized-fru", types.YLeaf{"UnrecognizedFru", basicInfo.UnrecognizedFru})
    basicInfo.EntityData.Leafs.Append("redundancystate", types.YLeaf{"Redundancystate", basicInfo.Redundancystate})
    basicInfo.EntityData.Leafs.Append("ceport", types.YLeaf{"Ceport", basicInfo.Ceport})
    basicInfo.EntityData.Leafs.Append("xr-scoped", types.YLeaf{"XrScoped", basicInfo.XrScoped})
    basicInfo.EntityData.Leafs.Append("unique-id", types.YLeaf{"UniqueId", basicInfo.UniqueId})

    basicInfo.EntityData.YListKeys = []string {}

    return &(basicInfo.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_FruInfo
// Field Replaceable Unit (FRU) inventory
// information
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_FruInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Administrative    state. The type is InvAdminState.
    ModuleAdministrativeState interface{}

    // Power administrative state. The type is InvPowerAdminState.
    ModulePowerAdministrativeState interface{}

    // Operation state. The type is InvCardState.
    ModuleOperationalState interface{}

    // Monitor state. The type is InvMonitorState.
    ModuleMonitorState interface{}

    // Reset reason. The type is InvResetReason.
    ModuleResetReason interface{}

    // Time operational state is   last changed.
    LastOperationalStateChange Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_FruInfo_LastOperationalStateChange

    // Module up time.
    ModuleUpTime Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_FruInfo_ModuleUpTime
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_FruInfo) GetEntityData() *types.CommonEntityData {
    fruInfo.EntityData.YFilter = fruInfo.YFilter
    fruInfo.EntityData.YangName = "fru-info"
    fruInfo.EntityData.BundleName = "cisco_ios_xr"
    fruInfo.EntityData.ParentYangName = "basic-attributes"
    fruInfo.EntityData.SegmentPath = "fru-info"
    fruInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/cards/card/basic-attributes/" + fruInfo.EntityData.SegmentPath
    fruInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fruInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fruInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fruInfo.EntityData.Children = types.NewOrderedMap()
    fruInfo.EntityData.Children.Append("last-operational-state-change", types.YChild{"LastOperationalStateChange", &fruInfo.LastOperationalStateChange})
    fruInfo.EntityData.Children.Append("module-up-time", types.YChild{"ModuleUpTime", &fruInfo.ModuleUpTime})
    fruInfo.EntityData.Leafs = types.NewOrderedMap()
    fruInfo.EntityData.Leafs.Append("module-administrative-state", types.YLeaf{"ModuleAdministrativeState", fruInfo.ModuleAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-power-administrative-state", types.YLeaf{"ModulePowerAdministrativeState", fruInfo.ModulePowerAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-operational-state", types.YLeaf{"ModuleOperationalState", fruInfo.ModuleOperationalState})
    fruInfo.EntityData.Leafs.Append("module-monitor-state", types.YLeaf{"ModuleMonitorState", fruInfo.ModuleMonitorState})
    fruInfo.EntityData.Leafs.Append("module-reset-reason", types.YLeaf{"ModuleResetReason", fruInfo.ModuleResetReason})

    fruInfo.EntityData.YListKeys = []string {}

    return &(fruInfo.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_FruInfo_LastOperationalStateChange
// Time operational state is   last changed
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
    lastOperationalStateChange.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/cards/card/basic-attributes/fru-info/" + lastOperationalStateChange.EntityData.SegmentPath
    lastOperationalStateChange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastOperationalStateChange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastOperationalStateChange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastOperationalStateChange.EntityData.Children = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", lastOperationalStateChange.TimeInSeconds})
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", lastOperationalStateChange.TimeInNanoSeconds})

    lastOperationalStateChange.EntityData.YListKeys = []string {}

    return &(lastOperationalStateChange.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_FruInfo_ModuleUpTime
// Module up time
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_FruInfo_ModuleUpTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (moduleUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_FruInfo_ModuleUpTime) GetEntityData() *types.CommonEntityData {
    moduleUpTime.EntityData.YFilter = moduleUpTime.YFilter
    moduleUpTime.EntityData.YangName = "module-up-time"
    moduleUpTime.EntityData.BundleName = "cisco_ios_xr"
    moduleUpTime.EntityData.ParentYangName = "fru-info"
    moduleUpTime.EntityData.SegmentPath = "module-up-time"
    moduleUpTime.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/cards/card/basic-attributes/fru-info/" + moduleUpTime.EntityData.SegmentPath
    moduleUpTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    moduleUpTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    moduleUpTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    moduleUpTime.EntityData.Children = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", moduleUpTime.TimeInSeconds})
    moduleUpTime.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", moduleUpTime.TimeInNanoSeconds})

    moduleUpTime.EntityData.YListKeys = []string {}

    return &(moduleUpTime.EntityData)
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
    basicAttributes.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/" + basicAttributes.EntityData.SegmentPath
    basicAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicAttributes.EntityData.Children = types.NewOrderedMap()
    basicAttributes.EntityData.Children.Append("basic-info", types.YChild{"BasicInfo", &basicAttributes.BasicInfo})
    basicAttributes.EntityData.Children.Append("fru-info", types.YChild{"FruInfo", &basicAttributes.FruInfo})
    basicAttributes.EntityData.Leafs = types.NewOrderedMap()

    basicAttributes.EntityData.YListKeys = []string {}

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

    // integer value for Redundancy State if applicable to this entity. The type
    // is interface{} with range: -2147483648..2147483647.
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
    basicInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/basic-attributes/" + basicInfo.EntityData.SegmentPath
    basicInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicInfo.EntityData.Children = types.NewOrderedMap()
    basicInfo.EntityData.Leafs = types.NewOrderedMap()
    basicInfo.EntityData.Leafs.Append("description", types.YLeaf{"Description", basicInfo.Description})
    basicInfo.EntityData.Leafs.Append("vendor-type", types.YLeaf{"VendorType", basicInfo.VendorType})
    basicInfo.EntityData.Leafs.Append("name", types.YLeaf{"Name", basicInfo.Name})
    basicInfo.EntityData.Leafs.Append("hardware-revision", types.YLeaf{"HardwareRevision", basicInfo.HardwareRevision})
    basicInfo.EntityData.Leafs.Append("firmware-revision", types.YLeaf{"FirmwareRevision", basicInfo.FirmwareRevision})
    basicInfo.EntityData.Leafs.Append("software-revision", types.YLeaf{"SoftwareRevision", basicInfo.SoftwareRevision})
    basicInfo.EntityData.Leafs.Append("chip-hardware-revision", types.YLeaf{"ChipHardwareRevision", basicInfo.ChipHardwareRevision})
    basicInfo.EntityData.Leafs.Append("serial-number", types.YLeaf{"SerialNumber", basicInfo.SerialNumber})
    basicInfo.EntityData.Leafs.Append("manufacturer-name", types.YLeaf{"ManufacturerName", basicInfo.ManufacturerName})
    basicInfo.EntityData.Leafs.Append("model-name", types.YLeaf{"ModelName", basicInfo.ModelName})
    basicInfo.EntityData.Leafs.Append("asset-id-str", types.YLeaf{"AssetIdStr", basicInfo.AssetIdStr})
    basicInfo.EntityData.Leafs.Append("asset-identification", types.YLeaf{"AssetIdentification", basicInfo.AssetIdentification})
    basicInfo.EntityData.Leafs.Append("is-field-replaceable-unit", types.YLeaf{"IsFieldReplaceableUnit", basicInfo.IsFieldReplaceableUnit})
    basicInfo.EntityData.Leafs.Append("manufacturer-asset-tags", types.YLeaf{"ManufacturerAssetTags", basicInfo.ManufacturerAssetTags})
    basicInfo.EntityData.Leafs.Append("composite-class-code", types.YLeaf{"CompositeClassCode", basicInfo.CompositeClassCode})
    basicInfo.EntityData.Leafs.Append("memory-size", types.YLeaf{"MemorySize", basicInfo.MemorySize})
    basicInfo.EntityData.Leafs.Append("environmental-monitor-path", types.YLeaf{"EnvironmentalMonitorPath", basicInfo.EnvironmentalMonitorPath})
    basicInfo.EntityData.Leafs.Append("alias", types.YLeaf{"Alias", basicInfo.Alias})
    basicInfo.EntityData.Leafs.Append("group-flag", types.YLeaf{"GroupFlag", basicInfo.GroupFlag})
    basicInfo.EntityData.Leafs.Append("new-deviation-number", types.YLeaf{"NewDeviationNumber", basicInfo.NewDeviationNumber})
    basicInfo.EntityData.Leafs.Append("physical-layer-interface-module-type", types.YLeaf{"PhysicalLayerInterfaceModuleType", basicInfo.PhysicalLayerInterfaceModuleType})
    basicInfo.EntityData.Leafs.Append("unrecognized-fru", types.YLeaf{"UnrecognizedFru", basicInfo.UnrecognizedFru})
    basicInfo.EntityData.Leafs.Append("redundancystate", types.YLeaf{"Redundancystate", basicInfo.Redundancystate})
    basicInfo.EntityData.Leafs.Append("ceport", types.YLeaf{"Ceport", basicInfo.Ceport})
    basicInfo.EntityData.Leafs.Append("xr-scoped", types.YLeaf{"XrScoped", basicInfo.XrScoped})
    basicInfo.EntityData.Leafs.Append("unique-id", types.YLeaf{"UniqueId", basicInfo.UniqueId})

    basicInfo.EntityData.YListKeys = []string {}

    return &(basicInfo.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_BasicAttributes_FruInfo
// Field Replaceable Unit (FRU) inventory
// information
type Inventory_Racks_Rack_Slots_Slot_BasicAttributes_FruInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Administrative    state. The type is InvAdminState.
    ModuleAdministrativeState interface{}

    // Power administrative state. The type is InvPowerAdminState.
    ModulePowerAdministrativeState interface{}

    // Operation state. The type is InvCardState.
    ModuleOperationalState interface{}

    // Monitor state. The type is InvMonitorState.
    ModuleMonitorState interface{}

    // Reset reason. The type is InvResetReason.
    ModuleResetReason interface{}

    // Time operational state is   last changed.
    LastOperationalStateChange Inventory_Racks_Rack_Slots_Slot_BasicAttributes_FruInfo_LastOperationalStateChange

    // Module up time.
    ModuleUpTime Inventory_Racks_Rack_Slots_Slot_BasicAttributes_FruInfo_ModuleUpTime
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_BasicAttributes_FruInfo) GetEntityData() *types.CommonEntityData {
    fruInfo.EntityData.YFilter = fruInfo.YFilter
    fruInfo.EntityData.YangName = "fru-info"
    fruInfo.EntityData.BundleName = "cisco_ios_xr"
    fruInfo.EntityData.ParentYangName = "basic-attributes"
    fruInfo.EntityData.SegmentPath = "fru-info"
    fruInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/basic-attributes/" + fruInfo.EntityData.SegmentPath
    fruInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fruInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fruInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fruInfo.EntityData.Children = types.NewOrderedMap()
    fruInfo.EntityData.Children.Append("last-operational-state-change", types.YChild{"LastOperationalStateChange", &fruInfo.LastOperationalStateChange})
    fruInfo.EntityData.Children.Append("module-up-time", types.YChild{"ModuleUpTime", &fruInfo.ModuleUpTime})
    fruInfo.EntityData.Leafs = types.NewOrderedMap()
    fruInfo.EntityData.Leafs.Append("module-administrative-state", types.YLeaf{"ModuleAdministrativeState", fruInfo.ModuleAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-power-administrative-state", types.YLeaf{"ModulePowerAdministrativeState", fruInfo.ModulePowerAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-operational-state", types.YLeaf{"ModuleOperationalState", fruInfo.ModuleOperationalState})
    fruInfo.EntityData.Leafs.Append("module-monitor-state", types.YLeaf{"ModuleMonitorState", fruInfo.ModuleMonitorState})
    fruInfo.EntityData.Leafs.Append("module-reset-reason", types.YLeaf{"ModuleResetReason", fruInfo.ModuleResetReason})

    fruInfo.EntityData.YListKeys = []string {}

    return &(fruInfo.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_BasicAttributes_FruInfo_LastOperationalStateChange
// Time operational state is   last changed
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
    lastOperationalStateChange.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/basic-attributes/fru-info/" + lastOperationalStateChange.EntityData.SegmentPath
    lastOperationalStateChange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastOperationalStateChange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastOperationalStateChange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastOperationalStateChange.EntityData.Children = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", lastOperationalStateChange.TimeInSeconds})
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", lastOperationalStateChange.TimeInNanoSeconds})

    lastOperationalStateChange.EntityData.YListKeys = []string {}

    return &(lastOperationalStateChange.EntityData)
}

// Inventory_Racks_Rack_Slots_Slot_BasicAttributes_FruInfo_ModuleUpTime
// Module up time
type Inventory_Racks_Rack_Slots_Slot_BasicAttributes_FruInfo_ModuleUpTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (moduleUpTime *Inventory_Racks_Rack_Slots_Slot_BasicAttributes_FruInfo_ModuleUpTime) GetEntityData() *types.CommonEntityData {
    moduleUpTime.EntityData.YFilter = moduleUpTime.YFilter
    moduleUpTime.EntityData.YangName = "module-up-time"
    moduleUpTime.EntityData.BundleName = "cisco_ios_xr"
    moduleUpTime.EntityData.ParentYangName = "fru-info"
    moduleUpTime.EntityData.SegmentPath = "module-up-time"
    moduleUpTime.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-invmgr-oper:inventory/racks/rack/slots/slot/basic-attributes/fru-info/" + moduleUpTime.EntityData.SegmentPath
    moduleUpTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    moduleUpTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    moduleUpTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    moduleUpTime.EntityData.Children = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", moduleUpTime.TimeInSeconds})
    moduleUpTime.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", moduleUpTime.TimeInNanoSeconds})

    moduleUpTime.EntityData.YListKeys = []string {}

    return &(moduleUpTime.EntityData)
}

