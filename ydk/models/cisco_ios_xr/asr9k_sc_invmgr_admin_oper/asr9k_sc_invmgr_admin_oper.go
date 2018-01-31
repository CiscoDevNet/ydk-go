// This module contains a collection of YANG definitions
// for Cisco IOS-XR asr9k-sc-invmgr package
// admin-plane operational data.
// 
// This module contains definitions
// for the following management objects:
//   inventory: Inventory operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package asr9k_sc_invmgr_admin_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package asr9k_sc_invmgr_admin_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-asr9k-sc-invmgr-admin-oper inventory}", reflect.TypeOf(Inventory{}))
    ydk.RegisterEntity("Cisco-IOS-XR-asr9k-sc-invmgr-admin-oper:inventory", reflect.TypeOf(Inventory{}))
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
// Inventory operational data
type Inventory struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Table of racks.
    Racks Inventory_Racks
}

func (inventory *Inventory) GetFilter() yfilter.YFilter { return inventory.YFilter }

func (inventory *Inventory) SetFilter(yf yfilter.YFilter) { inventory.YFilter = yf }

func (inventory *Inventory) GetGoName(yname string) string {
    if yname == "racks" { return "Racks" }
    return ""
}

func (inventory *Inventory) GetSegmentPath() string {
    return "Cisco-IOS-XR-asr9k-sc-invmgr-admin-oper:inventory"
}

func (inventory *Inventory) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "racks" {
        return &inventory.Racks
    }
    return nil
}

func (inventory *Inventory) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["racks"] = &inventory.Racks
    return children
}

func (inventory *Inventory) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (inventory *Inventory) GetBundleName() string { return "cisco_ios_xr" }

func (inventory *Inventory) GetYangName() string { return "inventory" }

func (inventory *Inventory) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inventory *Inventory) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inventory *Inventory) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inventory *Inventory) SetParent(parent types.Entity) { inventory.parent = parent }

func (inventory *Inventory) GetParent() types.Entity { return inventory.parent }

func (inventory *Inventory) GetParentYangName() string { return "Cisco-IOS-XR-asr9k-sc-invmgr-admin-oper" }

// Inventory_Racks
// Table of racks
type Inventory_Racks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Rack number. The type is slice of Inventory_Racks_Rack.
    Rack []Inventory_Racks_Rack
}

func (racks *Inventory_Racks) GetFilter() yfilter.YFilter { return racks.YFilter }

func (racks *Inventory_Racks) SetFilter(yf yfilter.YFilter) { racks.YFilter = yf }

func (racks *Inventory_Racks) GetGoName(yname string) string {
    if yname == "rack" { return "Rack" }
    return ""
}

func (racks *Inventory_Racks) GetSegmentPath() string {
    return "racks"
}

func (racks *Inventory_Racks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rack" {
        for _, c := range racks.Rack {
            if racks.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Inventory_Racks_Rack{}
        racks.Rack = append(racks.Rack, child)
        return &racks.Rack[len(racks.Rack)-1]
    }
    return nil
}

func (racks *Inventory_Racks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range racks.Rack {
        children[racks.Rack[i].GetSegmentPath()] = &racks.Rack[i]
    }
    return children
}

func (racks *Inventory_Racks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (racks *Inventory_Racks) GetBundleName() string { return "cisco_ios_xr" }

func (racks *Inventory_Racks) GetYangName() string { return "racks" }

func (racks *Inventory_Racks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (racks *Inventory_Racks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (racks *Inventory_Racks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (racks *Inventory_Racks) SetParent(parent types.Entity) { racks.parent = parent }

func (racks *Inventory_Racks) GetParent() types.Entity { return racks.parent }

func (racks *Inventory_Racks) GetParentYangName() string { return "inventory" }

// Inventory_Racks_Rack
// Rack number
type Inventory_Racks_Rack struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Rack number. The type is interface{} with range:
    // -2147483648..2147483647.
    Number interface{}

    // Table for rack power supply shelves .
    PowerSupplyShelfs Inventory_Racks_Rack_PowerSupplyShelfs

    // Slot table contains all slots in the rack.
    Slots Inventory_Racks_Rack_Slots

    // Table for rack fan trays.
    FanTraies Inventory_Racks_Rack_FanTraies

    // Table for defining rack power supply zones .
    PowerSupplyZones Inventory_Racks_Rack_PowerSupplyZones

    // Attributes.
    BasicAttributes Inventory_Racks_Rack_BasicAttributes
}

func (rack *Inventory_Racks_Rack) GetFilter() yfilter.YFilter { return rack.YFilter }

func (rack *Inventory_Racks_Rack) SetFilter(yf yfilter.YFilter) { rack.YFilter = yf }

func (rack *Inventory_Racks_Rack) GetGoName(yname string) string {
    if yname == "number" { return "Number" }
    if yname == "power-supply-shelfs" { return "PowerSupplyShelfs" }
    if yname == "slots" { return "Slots" }
    if yname == "fan-traies" { return "FanTraies" }
    if yname == "power-supply-zones" { return "PowerSupplyZones" }
    if yname == "basic-attributes" { return "BasicAttributes" }
    return ""
}

func (rack *Inventory_Racks_Rack) GetSegmentPath() string {
    return "rack" + "[number='" + fmt.Sprintf("%v", rack.Number) + "']"
}

func (rack *Inventory_Racks_Rack) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "power-supply-shelfs" {
        return &rack.PowerSupplyShelfs
    }
    if childYangName == "slots" {
        return &rack.Slots
    }
    if childYangName == "fan-traies" {
        return &rack.FanTraies
    }
    if childYangName == "power-supply-zones" {
        return &rack.PowerSupplyZones
    }
    if childYangName == "basic-attributes" {
        return &rack.BasicAttributes
    }
    return nil
}

func (rack *Inventory_Racks_Rack) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["power-supply-shelfs"] = &rack.PowerSupplyShelfs
    children["slots"] = &rack.Slots
    children["fan-traies"] = &rack.FanTraies
    children["power-supply-zones"] = &rack.PowerSupplyZones
    children["basic-attributes"] = &rack.BasicAttributes
    return children
}

func (rack *Inventory_Racks_Rack) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["number"] = rack.Number
    return leafs
}

func (rack *Inventory_Racks_Rack) GetBundleName() string { return "cisco_ios_xr" }

func (rack *Inventory_Racks_Rack) GetYangName() string { return "rack" }

func (rack *Inventory_Racks_Rack) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rack *Inventory_Racks_Rack) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rack *Inventory_Racks_Rack) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rack *Inventory_Racks_Rack) SetParent(parent types.Entity) { rack.parent = parent }

func (rack *Inventory_Racks_Rack) GetParent() types.Entity { return rack.parent }

func (rack *Inventory_Racks_Rack) GetParentYangName() string { return "racks" }

// Inventory_Racks_Rack_PowerSupplyShelfs
// Table for rack power supply shelves 
type Inventory_Racks_Rack_PowerSupplyShelfs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Power Supply Shelf number. The type is slice of
    // Inventory_Racks_Rack_PowerSupplyShelfs_PowerSupplyShelf.
    PowerSupplyShelf []Inventory_Racks_Rack_PowerSupplyShelfs_PowerSupplyShelf
}

func (powerSupplyShelfs *Inventory_Racks_Rack_PowerSupplyShelfs) GetFilter() yfilter.YFilter { return powerSupplyShelfs.YFilter }

func (powerSupplyShelfs *Inventory_Racks_Rack_PowerSupplyShelfs) SetFilter(yf yfilter.YFilter) { powerSupplyShelfs.YFilter = yf }

func (powerSupplyShelfs *Inventory_Racks_Rack_PowerSupplyShelfs) GetGoName(yname string) string {
    if yname == "power-supply-shelf" { return "PowerSupplyShelf" }
    return ""
}

func (powerSupplyShelfs *Inventory_Racks_Rack_PowerSupplyShelfs) GetSegmentPath() string {
    return "power-supply-shelfs"
}

func (powerSupplyShelfs *Inventory_Racks_Rack_PowerSupplyShelfs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "power-supply-shelf" {
        for _, c := range powerSupplyShelfs.PowerSupplyShelf {
            if powerSupplyShelfs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Inventory_Racks_Rack_PowerSupplyShelfs_PowerSupplyShelf{}
        powerSupplyShelfs.PowerSupplyShelf = append(powerSupplyShelfs.PowerSupplyShelf, child)
        return &powerSupplyShelfs.PowerSupplyShelf[len(powerSupplyShelfs.PowerSupplyShelf)-1]
    }
    return nil
}

func (powerSupplyShelfs *Inventory_Racks_Rack_PowerSupplyShelfs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range powerSupplyShelfs.PowerSupplyShelf {
        children[powerSupplyShelfs.PowerSupplyShelf[i].GetSegmentPath()] = &powerSupplyShelfs.PowerSupplyShelf[i]
    }
    return children
}

func (powerSupplyShelfs *Inventory_Racks_Rack_PowerSupplyShelfs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (powerSupplyShelfs *Inventory_Racks_Rack_PowerSupplyShelfs) GetBundleName() string { return "cisco_ios_xr" }

func (powerSupplyShelfs *Inventory_Racks_Rack_PowerSupplyShelfs) GetYangName() string { return "power-supply-shelfs" }

func (powerSupplyShelfs *Inventory_Racks_Rack_PowerSupplyShelfs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (powerSupplyShelfs *Inventory_Racks_Rack_PowerSupplyShelfs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (powerSupplyShelfs *Inventory_Racks_Rack_PowerSupplyShelfs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (powerSupplyShelfs *Inventory_Racks_Rack_PowerSupplyShelfs) SetParent(parent types.Entity) { powerSupplyShelfs.parent = parent }

func (powerSupplyShelfs *Inventory_Racks_Rack_PowerSupplyShelfs) GetParent() types.Entity { return powerSupplyShelfs.parent }

func (powerSupplyShelfs *Inventory_Racks_Rack_PowerSupplyShelfs) GetParentYangName() string { return "rack" }

// Inventory_Racks_Rack_PowerSupplyShelfs_PowerSupplyShelf
// Power Supply Shelf number
type Inventory_Racks_Rack_PowerSupplyShelfs_PowerSupplyShelf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Power Shelf number. The type is interface{} with
    // range: -2147483648..2147483647.
    Number interface{}

    // Attributes.
    BasicAttributes Inventory_Racks_Rack_PowerSupplyShelfs_PowerSupplyShelf_BasicAttributes
}

func (powerSupplyShelf *Inventory_Racks_Rack_PowerSupplyShelfs_PowerSupplyShelf) GetFilter() yfilter.YFilter { return powerSupplyShelf.YFilter }

func (powerSupplyShelf *Inventory_Racks_Rack_PowerSupplyShelfs_PowerSupplyShelf) SetFilter(yf yfilter.YFilter) { powerSupplyShelf.YFilter = yf }

func (powerSupplyShelf *Inventory_Racks_Rack_PowerSupplyShelfs_PowerSupplyShelf) GetGoName(yname string) string {
    if yname == "number" { return "Number" }
    if yname == "basic-attributes" { return "BasicAttributes" }
    return ""
}

func (powerSupplyShelf *Inventory_Racks_Rack_PowerSupplyShelfs_PowerSupplyShelf) GetSegmentPath() string {
    return "power-supply-shelf" + "[number='" + fmt.Sprintf("%v", powerSupplyShelf.Number) + "']"
}

func (powerSupplyShelf *Inventory_Racks_Rack_PowerSupplyShelfs_PowerSupplyShelf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "basic-attributes" {
        return &powerSupplyShelf.BasicAttributes
    }
    return nil
}

func (powerSupplyShelf *Inventory_Racks_Rack_PowerSupplyShelfs_PowerSupplyShelf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["basic-attributes"] = &powerSupplyShelf.BasicAttributes
    return children
}

func (powerSupplyShelf *Inventory_Racks_Rack_PowerSupplyShelfs_PowerSupplyShelf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["number"] = powerSupplyShelf.Number
    return leafs
}

func (powerSupplyShelf *Inventory_Racks_Rack_PowerSupplyShelfs_PowerSupplyShelf) GetBundleName() string { return "cisco_ios_xr" }

func (powerSupplyShelf *Inventory_Racks_Rack_PowerSupplyShelfs_PowerSupplyShelf) GetYangName() string { return "power-supply-shelf" }

func (powerSupplyShelf *Inventory_Racks_Rack_PowerSupplyShelfs_PowerSupplyShelf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (powerSupplyShelf *Inventory_Racks_Rack_PowerSupplyShelfs_PowerSupplyShelf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (powerSupplyShelf *Inventory_Racks_Rack_PowerSupplyShelfs_PowerSupplyShelf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (powerSupplyShelf *Inventory_Racks_Rack_PowerSupplyShelfs_PowerSupplyShelf) SetParent(parent types.Entity) { powerSupplyShelf.parent = parent }

func (powerSupplyShelf *Inventory_Racks_Rack_PowerSupplyShelfs_PowerSupplyShelf) GetParent() types.Entity { return powerSupplyShelf.parent }

func (powerSupplyShelf *Inventory_Racks_Rack_PowerSupplyShelfs_PowerSupplyShelf) GetParentYangName() string { return "power-supply-shelfs" }

// Inventory_Racks_Rack_PowerSupplyShelfs_PowerSupplyShelf_BasicAttributes
// Attributes
type Inventory_Racks_Rack_PowerSupplyShelfs_PowerSupplyShelf_BasicAttributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Inventory information.
    BasicInfo Inventory_Racks_Rack_PowerSupplyShelfs_PowerSupplyShelf_BasicAttributes_BasicInfo
}

func (basicAttributes *Inventory_Racks_Rack_PowerSupplyShelfs_PowerSupplyShelf_BasicAttributes) GetFilter() yfilter.YFilter { return basicAttributes.YFilter }

func (basicAttributes *Inventory_Racks_Rack_PowerSupplyShelfs_PowerSupplyShelf_BasicAttributes) SetFilter(yf yfilter.YFilter) { basicAttributes.YFilter = yf }

func (basicAttributes *Inventory_Racks_Rack_PowerSupplyShelfs_PowerSupplyShelf_BasicAttributes) GetGoName(yname string) string {
    if yname == "basic-info" { return "BasicInfo" }
    return ""
}

func (basicAttributes *Inventory_Racks_Rack_PowerSupplyShelfs_PowerSupplyShelf_BasicAttributes) GetSegmentPath() string {
    return "basic-attributes"
}

func (basicAttributes *Inventory_Racks_Rack_PowerSupplyShelfs_PowerSupplyShelf_BasicAttributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "basic-info" {
        return &basicAttributes.BasicInfo
    }
    return nil
}

func (basicAttributes *Inventory_Racks_Rack_PowerSupplyShelfs_PowerSupplyShelf_BasicAttributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["basic-info"] = &basicAttributes.BasicInfo
    return children
}

func (basicAttributes *Inventory_Racks_Rack_PowerSupplyShelfs_PowerSupplyShelf_BasicAttributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (basicAttributes *Inventory_Racks_Rack_PowerSupplyShelfs_PowerSupplyShelf_BasicAttributes) GetBundleName() string { return "cisco_ios_xr" }

func (basicAttributes *Inventory_Racks_Rack_PowerSupplyShelfs_PowerSupplyShelf_BasicAttributes) GetYangName() string { return "basic-attributes" }

func (basicAttributes *Inventory_Racks_Rack_PowerSupplyShelfs_PowerSupplyShelf_BasicAttributes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (basicAttributes *Inventory_Racks_Rack_PowerSupplyShelfs_PowerSupplyShelf_BasicAttributes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (basicAttributes *Inventory_Racks_Rack_PowerSupplyShelfs_PowerSupplyShelf_BasicAttributes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (basicAttributes *Inventory_Racks_Rack_PowerSupplyShelfs_PowerSupplyShelf_BasicAttributes) SetParent(parent types.Entity) { basicAttributes.parent = parent }

func (basicAttributes *Inventory_Racks_Rack_PowerSupplyShelfs_PowerSupplyShelf_BasicAttributes) GetParent() types.Entity { return basicAttributes.parent }

func (basicAttributes *Inventory_Racks_Rack_PowerSupplyShelfs_PowerSupplyShelf_BasicAttributes) GetParentYangName() string { return "power-supply-shelf" }

// Inventory_Racks_Rack_PowerSupplyShelfs_PowerSupplyShelf_BasicAttributes_BasicInfo
// Inventory information
type Inventory_Racks_Rack_PowerSupplyShelfs_PowerSupplyShelf_BasicAttributes_BasicInfo struct {
    parent types.Entity
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

func (basicInfo *Inventory_Racks_Rack_PowerSupplyShelfs_PowerSupplyShelf_BasicAttributes_BasicInfo) GetFilter() yfilter.YFilter { return basicInfo.YFilter }

func (basicInfo *Inventory_Racks_Rack_PowerSupplyShelfs_PowerSupplyShelf_BasicAttributes_BasicInfo) SetFilter(yf yfilter.YFilter) { basicInfo.YFilter = yf }

func (basicInfo *Inventory_Racks_Rack_PowerSupplyShelfs_PowerSupplyShelf_BasicAttributes_BasicInfo) GetGoName(yname string) string {
    if yname == "description" { return "Description" }
    if yname == "vendor-type" { return "VendorType" }
    if yname == "name" { return "Name" }
    if yname == "hardware-revision" { return "HardwareRevision" }
    if yname == "firmware-revision" { return "FirmwareRevision" }
    if yname == "software-revision" { return "SoftwareRevision" }
    if yname == "chip-hardware-revision" { return "ChipHardwareRevision" }
    if yname == "serial-number" { return "SerialNumber" }
    if yname == "manufacturer-name" { return "ManufacturerName" }
    if yname == "model-name" { return "ModelName" }
    if yname == "asset-id-str" { return "AssetIdStr" }
    if yname == "asset-identification" { return "AssetIdentification" }
    if yname == "is-field-replaceable-unit" { return "IsFieldReplaceableUnit" }
    if yname == "manufacturer-asset-tags" { return "ManufacturerAssetTags" }
    if yname == "composite-class-code" { return "CompositeClassCode" }
    if yname == "memory-size" { return "MemorySize" }
    if yname == "environmental-monitor-path" { return "EnvironmentalMonitorPath" }
    if yname == "alias" { return "Alias" }
    if yname == "group-flag" { return "GroupFlag" }
    if yname == "new-deviation-number" { return "NewDeviationNumber" }
    if yname == "physical-layer-interface-module-type" { return "PhysicalLayerInterfaceModuleType" }
    if yname == "unrecognized-fru" { return "UnrecognizedFru" }
    if yname == "redundancystate" { return "Redundancystate" }
    if yname == "ceport" { return "Ceport" }
    if yname == "xr-scoped" { return "XrScoped" }
    if yname == "unique-id" { return "UniqueId" }
    return ""
}

func (basicInfo *Inventory_Racks_Rack_PowerSupplyShelfs_PowerSupplyShelf_BasicAttributes_BasicInfo) GetSegmentPath() string {
    return "basic-info"
}

func (basicInfo *Inventory_Racks_Rack_PowerSupplyShelfs_PowerSupplyShelf_BasicAttributes_BasicInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (basicInfo *Inventory_Racks_Rack_PowerSupplyShelfs_PowerSupplyShelf_BasicAttributes_BasicInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (basicInfo *Inventory_Racks_Rack_PowerSupplyShelfs_PowerSupplyShelf_BasicAttributes_BasicInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["description"] = basicInfo.Description
    leafs["vendor-type"] = basicInfo.VendorType
    leafs["name"] = basicInfo.Name
    leafs["hardware-revision"] = basicInfo.HardwareRevision
    leafs["firmware-revision"] = basicInfo.FirmwareRevision
    leafs["software-revision"] = basicInfo.SoftwareRevision
    leafs["chip-hardware-revision"] = basicInfo.ChipHardwareRevision
    leafs["serial-number"] = basicInfo.SerialNumber
    leafs["manufacturer-name"] = basicInfo.ManufacturerName
    leafs["model-name"] = basicInfo.ModelName
    leafs["asset-id-str"] = basicInfo.AssetIdStr
    leafs["asset-identification"] = basicInfo.AssetIdentification
    leafs["is-field-replaceable-unit"] = basicInfo.IsFieldReplaceableUnit
    leafs["manufacturer-asset-tags"] = basicInfo.ManufacturerAssetTags
    leafs["composite-class-code"] = basicInfo.CompositeClassCode
    leafs["memory-size"] = basicInfo.MemorySize
    leafs["environmental-monitor-path"] = basicInfo.EnvironmentalMonitorPath
    leafs["alias"] = basicInfo.Alias
    leafs["group-flag"] = basicInfo.GroupFlag
    leafs["new-deviation-number"] = basicInfo.NewDeviationNumber
    leafs["physical-layer-interface-module-type"] = basicInfo.PhysicalLayerInterfaceModuleType
    leafs["unrecognized-fru"] = basicInfo.UnrecognizedFru
    leafs["redundancystate"] = basicInfo.Redundancystate
    leafs["ceport"] = basicInfo.Ceport
    leafs["xr-scoped"] = basicInfo.XrScoped
    leafs["unique-id"] = basicInfo.UniqueId
    return leafs
}

func (basicInfo *Inventory_Racks_Rack_PowerSupplyShelfs_PowerSupplyShelf_BasicAttributes_BasicInfo) GetBundleName() string { return "cisco_ios_xr" }

func (basicInfo *Inventory_Racks_Rack_PowerSupplyShelfs_PowerSupplyShelf_BasicAttributes_BasicInfo) GetYangName() string { return "basic-info" }

func (basicInfo *Inventory_Racks_Rack_PowerSupplyShelfs_PowerSupplyShelf_BasicAttributes_BasicInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (basicInfo *Inventory_Racks_Rack_PowerSupplyShelfs_PowerSupplyShelf_BasicAttributes_BasicInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (basicInfo *Inventory_Racks_Rack_PowerSupplyShelfs_PowerSupplyShelf_BasicAttributes_BasicInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (basicInfo *Inventory_Racks_Rack_PowerSupplyShelfs_PowerSupplyShelf_BasicAttributes_BasicInfo) SetParent(parent types.Entity) { basicInfo.parent = parent }

func (basicInfo *Inventory_Racks_Rack_PowerSupplyShelfs_PowerSupplyShelf_BasicAttributes_BasicInfo) GetParent() types.Entity { return basicInfo.parent }

func (basicInfo *Inventory_Racks_Rack_PowerSupplyShelfs_PowerSupplyShelf_BasicAttributes_BasicInfo) GetParentYangName() string { return "basic-attributes" }

// Inventory_Racks_Rack_Slots
// Slot table contains all slots in the rack
type Inventory_Racks_Rack_Slots struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Slot number. The type is slice of Inventory_Racks_Rack_Slots_Slot.
    Slot []Inventory_Racks_Rack_Slots_Slot
}

func (slots *Inventory_Racks_Rack_Slots) GetFilter() yfilter.YFilter { return slots.YFilter }

func (slots *Inventory_Racks_Rack_Slots) SetFilter(yf yfilter.YFilter) { slots.YFilter = yf }

func (slots *Inventory_Racks_Rack_Slots) GetGoName(yname string) string {
    if yname == "slot" { return "Slot" }
    return ""
}

func (slots *Inventory_Racks_Rack_Slots) GetSegmentPath() string {
    return "slots"
}

func (slots *Inventory_Racks_Rack_Slots) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "slot" {
        for _, c := range slots.Slot {
            if slots.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Inventory_Racks_Rack_Slots_Slot{}
        slots.Slot = append(slots.Slot, child)
        return &slots.Slot[len(slots.Slot)-1]
    }
    return nil
}

func (slots *Inventory_Racks_Rack_Slots) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range slots.Slot {
        children[slots.Slot[i].GetSegmentPath()] = &slots.Slot[i]
    }
    return children
}

func (slots *Inventory_Racks_Rack_Slots) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (slots *Inventory_Racks_Rack_Slots) GetBundleName() string { return "cisco_ios_xr" }

func (slots *Inventory_Racks_Rack_Slots) GetYangName() string { return "slots" }

func (slots *Inventory_Racks_Rack_Slots) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (slots *Inventory_Racks_Rack_Slots) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (slots *Inventory_Racks_Rack_Slots) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (slots *Inventory_Racks_Rack_Slots) SetParent(parent types.Entity) { slots.parent = parent }

func (slots *Inventory_Racks_Rack_Slots) GetParent() types.Entity { return slots.parent }

func (slots *Inventory_Racks_Rack_Slots) GetParentYangName() string { return "rack" }

// Inventory_Racks_Rack_Slots_Slot
// Slot number
type Inventory_Racks_Rack_Slots_Slot struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Slot number. The type is interface{} with range:
    // -2147483648..2147483647.
    Number interface{}

    // Card table contains all cards in the slot.
    Cards Inventory_Racks_Rack_Slots_Slot_Cards

    // Attributes.
    BasicAttributes Inventory_Racks_Rack_Slots_Slot_BasicAttributes
}

func (slot *Inventory_Racks_Rack_Slots_Slot) GetFilter() yfilter.YFilter { return slot.YFilter }

func (slot *Inventory_Racks_Rack_Slots_Slot) SetFilter(yf yfilter.YFilter) { slot.YFilter = yf }

func (slot *Inventory_Racks_Rack_Slots_Slot) GetGoName(yname string) string {
    if yname == "number" { return "Number" }
    if yname == "cards" { return "Cards" }
    if yname == "basic-attributes" { return "BasicAttributes" }
    return ""
}

func (slot *Inventory_Racks_Rack_Slots_Slot) GetSegmentPath() string {
    return "slot" + "[number='" + fmt.Sprintf("%v", slot.Number) + "']"
}

func (slot *Inventory_Racks_Rack_Slots_Slot) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cards" {
        return &slot.Cards
    }
    if childYangName == "basic-attributes" {
        return &slot.BasicAttributes
    }
    return nil
}

func (slot *Inventory_Racks_Rack_Slots_Slot) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cards"] = &slot.Cards
    children["basic-attributes"] = &slot.BasicAttributes
    return children
}

func (slot *Inventory_Racks_Rack_Slots_Slot) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["number"] = slot.Number
    return leafs
}

func (slot *Inventory_Racks_Rack_Slots_Slot) GetBundleName() string { return "cisco_ios_xr" }

func (slot *Inventory_Racks_Rack_Slots_Slot) GetYangName() string { return "slot" }

func (slot *Inventory_Racks_Rack_Slots_Slot) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (slot *Inventory_Racks_Rack_Slots_Slot) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (slot *Inventory_Racks_Rack_Slots_Slot) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (slot *Inventory_Racks_Rack_Slots_Slot) SetParent(parent types.Entity) { slot.parent = parent }

func (slot *Inventory_Racks_Rack_Slots_Slot) GetParent() types.Entity { return slot.parent }

func (slot *Inventory_Racks_Rack_Slots_Slot) GetParentYangName() string { return "slots" }

// Inventory_Racks_Rack_Slots_Slot_Cards
// Card table contains all cards in the slot
type Inventory_Racks_Rack_Slots_Slot_Cards struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Card number. The type is slice of
    // Inventory_Racks_Rack_Slots_Slot_Cards_Card.
    Card []Inventory_Racks_Rack_Slots_Slot_Cards_Card
}

func (cards *Inventory_Racks_Rack_Slots_Slot_Cards) GetFilter() yfilter.YFilter { return cards.YFilter }

func (cards *Inventory_Racks_Rack_Slots_Slot_Cards) SetFilter(yf yfilter.YFilter) { cards.YFilter = yf }

func (cards *Inventory_Racks_Rack_Slots_Slot_Cards) GetGoName(yname string) string {
    if yname == "card" { return "Card" }
    return ""
}

func (cards *Inventory_Racks_Rack_Slots_Slot_Cards) GetSegmentPath() string {
    return "cards"
}

func (cards *Inventory_Racks_Rack_Slots_Slot_Cards) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "card" {
        for _, c := range cards.Card {
            if cards.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Inventory_Racks_Rack_Slots_Slot_Cards_Card{}
        cards.Card = append(cards.Card, child)
        return &cards.Card[len(cards.Card)-1]
    }
    return nil
}

func (cards *Inventory_Racks_Rack_Slots_Slot_Cards) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cards.Card {
        children[cards.Card[i].GetSegmentPath()] = &cards.Card[i]
    }
    return children
}

func (cards *Inventory_Racks_Rack_Slots_Slot_Cards) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cards *Inventory_Racks_Rack_Slots_Slot_Cards) GetBundleName() string { return "cisco_ios_xr" }

func (cards *Inventory_Racks_Rack_Slots_Slot_Cards) GetYangName() string { return "cards" }

func (cards *Inventory_Racks_Rack_Slots_Slot_Cards) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cards *Inventory_Racks_Rack_Slots_Slot_Cards) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cards *Inventory_Racks_Rack_Slots_Slot_Cards) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cards *Inventory_Racks_Rack_Slots_Slot_Cards) SetParent(parent types.Entity) { cards.parent = parent }

func (cards *Inventory_Racks_Rack_Slots_Slot_Cards) GetParent() types.Entity { return cards.parent }

func (cards *Inventory_Racks_Rack_Slots_Slot_Cards) GetParentYangName() string { return "slot" }

// Inventory_Racks_Rack_Slots_Slot_Cards_Card
// Card number
type Inventory_Racks_Rack_Slots_Slot_Cards_Card struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. card number. The type is interface{} with range:
    // -2147483648..2147483647.
    Number interface{}

    // SubSlotTable contains all subslots in a Slot.
    SubSlots Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots

    // HWComponent table contains all HW modules within the card .
    HwComponents Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents

    // PortSlotTable contains all optics ports in a SPA/PLIM.
    PortSlots Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots

    // ModuleSensorTable contains all sensors in a Module.
    Sensors Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors

    // Attributes.
    BasicAttributes Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes
}

func (card *Inventory_Racks_Rack_Slots_Slot_Cards_Card) GetFilter() yfilter.YFilter { return card.YFilter }

func (card *Inventory_Racks_Rack_Slots_Slot_Cards_Card) SetFilter(yf yfilter.YFilter) { card.YFilter = yf }

func (card *Inventory_Racks_Rack_Slots_Slot_Cards_Card) GetGoName(yname string) string {
    if yname == "number" { return "Number" }
    if yname == "sub-slots" { return "SubSlots" }
    if yname == "hw-components" { return "HwComponents" }
    if yname == "port-slots" { return "PortSlots" }
    if yname == "sensors" { return "Sensors" }
    if yname == "basic-attributes" { return "BasicAttributes" }
    return ""
}

func (card *Inventory_Racks_Rack_Slots_Slot_Cards_Card) GetSegmentPath() string {
    return "card" + "[number='" + fmt.Sprintf("%v", card.Number) + "']"
}

func (card *Inventory_Racks_Rack_Slots_Slot_Cards_Card) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sub-slots" {
        return &card.SubSlots
    }
    if childYangName == "hw-components" {
        return &card.HwComponents
    }
    if childYangName == "port-slots" {
        return &card.PortSlots
    }
    if childYangName == "sensors" {
        return &card.Sensors
    }
    if childYangName == "basic-attributes" {
        return &card.BasicAttributes
    }
    return nil
}

func (card *Inventory_Racks_Rack_Slots_Slot_Cards_Card) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sub-slots"] = &card.SubSlots
    children["hw-components"] = &card.HwComponents
    children["port-slots"] = &card.PortSlots
    children["sensors"] = &card.Sensors
    children["basic-attributes"] = &card.BasicAttributes
    return children
}

func (card *Inventory_Racks_Rack_Slots_Slot_Cards_Card) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["number"] = card.Number
    return leafs
}

func (card *Inventory_Racks_Rack_Slots_Slot_Cards_Card) GetBundleName() string { return "cisco_ios_xr" }

func (card *Inventory_Racks_Rack_Slots_Slot_Cards_Card) GetYangName() string { return "card" }

func (card *Inventory_Racks_Rack_Slots_Slot_Cards_Card) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (card *Inventory_Racks_Rack_Slots_Slot_Cards_Card) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (card *Inventory_Racks_Rack_Slots_Slot_Cards_Card) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (card *Inventory_Racks_Rack_Slots_Slot_Cards_Card) SetParent(parent types.Entity) { card.parent = parent }

func (card *Inventory_Racks_Rack_Slots_Slot_Cards_Card) GetParent() types.Entity { return card.parent }

func (card *Inventory_Racks_Rack_Slots_Slot_Cards_Card) GetParentYangName() string { return "cards" }

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots
// SubSlotTable contains all subslots in a
// Slot
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SubSlot number. The type is slice of
    // Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot.
    SubSlot []Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot
}

func (subSlots *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots) GetFilter() yfilter.YFilter { return subSlots.YFilter }

func (subSlots *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots) SetFilter(yf yfilter.YFilter) { subSlots.YFilter = yf }

func (subSlots *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots) GetGoName(yname string) string {
    if yname == "sub-slot" { return "SubSlot" }
    return ""
}

func (subSlots *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots) GetSegmentPath() string {
    return "sub-slots"
}

func (subSlots *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sub-slot" {
        for _, c := range subSlots.SubSlot {
            if subSlots.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot{}
        subSlots.SubSlot = append(subSlots.SubSlot, child)
        return &subSlots.SubSlot[len(subSlots.SubSlot)-1]
    }
    return nil
}

func (subSlots *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range subSlots.SubSlot {
        children[subSlots.SubSlot[i].GetSegmentPath()] = &subSlots.SubSlot[i]
    }
    return children
}

func (subSlots *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (subSlots *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots) GetBundleName() string { return "cisco_ios_xr" }

func (subSlots *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots) GetYangName() string { return "sub-slots" }

func (subSlots *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (subSlots *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (subSlots *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (subSlots *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots) SetParent(parent types.Entity) { subSlots.parent = parent }

func (subSlots *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots) GetParent() types.Entity { return subSlots.parent }

func (subSlots *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots) GetParentYangName() string { return "card" }

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot
// SubSlot number
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. subslot number. The type is interface{} with
    // range: -2147483648..2147483647.
    Number interface{}

    // Module string.
    Module Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module

    // Attributes.
    BasicAttributes Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes
}

func (subSlot *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot) GetFilter() yfilter.YFilter { return subSlot.YFilter }

func (subSlot *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot) SetFilter(yf yfilter.YFilter) { subSlot.YFilter = yf }

func (subSlot *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot) GetGoName(yname string) string {
    if yname == "number" { return "Number" }
    if yname == "module" { return "Module" }
    if yname == "basic-attributes" { return "BasicAttributes" }
    return ""
}

func (subSlot *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot) GetSegmentPath() string {
    return "sub-slot" + "[number='" + fmt.Sprintf("%v", subSlot.Number) + "']"
}

func (subSlot *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "module" {
        return &subSlot.Module
    }
    if childYangName == "basic-attributes" {
        return &subSlot.BasicAttributes
    }
    return nil
}

func (subSlot *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["module"] = &subSlot.Module
    children["basic-attributes"] = &subSlot.BasicAttributes
    return children
}

func (subSlot *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["number"] = subSlot.Number
    return leafs
}

func (subSlot *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot) GetBundleName() string { return "cisco_ios_xr" }

func (subSlot *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot) GetYangName() string { return "sub-slot" }

func (subSlot *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (subSlot *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (subSlot *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (subSlot *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot) SetParent(parent types.Entity) { subSlot.parent = parent }

func (subSlot *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot) GetParent() types.Entity { return subSlot.parent }

func (subSlot *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot) GetParentYangName() string { return "sub-slots" }

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module
// Module string
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PortSlotTable contains all optics ports in a SPA/PLIM.
    PortSlots Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots

    // ModuleSensorTable contains all sensors in a Module.
    Sensors Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors

    // Attributes.
    BasicAttributes Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes
}

func (module *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module) GetFilter() yfilter.YFilter { return module.YFilter }

func (module *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module) SetFilter(yf yfilter.YFilter) { module.YFilter = yf }

func (module *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module) GetGoName(yname string) string {
    if yname == "port-slots" { return "PortSlots" }
    if yname == "sensors" { return "Sensors" }
    if yname == "basic-attributes" { return "BasicAttributes" }
    return ""
}

func (module *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module) GetSegmentPath() string {
    return "module"
}

func (module *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "port-slots" {
        return &module.PortSlots
    }
    if childYangName == "sensors" {
        return &module.Sensors
    }
    if childYangName == "basic-attributes" {
        return &module.BasicAttributes
    }
    return nil
}

func (module *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["port-slots"] = &module.PortSlots
    children["sensors"] = &module.Sensors
    children["basic-attributes"] = &module.BasicAttributes
    return children
}

func (module *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (module *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module) GetBundleName() string { return "cisco_ios_xr" }

func (module *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module) GetYangName() string { return "module" }

func (module *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (module *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (module *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (module *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module) SetParent(parent types.Entity) { module.parent = parent }

func (module *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module) GetParent() types.Entity { return module.parent }

func (module *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module) GetParentYangName() string { return "sub-slot" }

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots
// PortSlotTable contains all optics ports in a
// SPA/PLIM.
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PortSlot number in the SPA/PLIM. The type is slice of
    // Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot.
    PortSlot []Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot
}

func (portSlots *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots) GetFilter() yfilter.YFilter { return portSlots.YFilter }

func (portSlots *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots) SetFilter(yf yfilter.YFilter) { portSlots.YFilter = yf }

func (portSlots *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots) GetGoName(yname string) string {
    if yname == "port-slot" { return "PortSlot" }
    return ""
}

func (portSlots *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots) GetSegmentPath() string {
    return "port-slots"
}

func (portSlots *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "port-slot" {
        for _, c := range portSlots.PortSlot {
            if portSlots.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot{}
        portSlots.PortSlot = append(portSlots.PortSlot, child)
        return &portSlots.PortSlot[len(portSlots.PortSlot)-1]
    }
    return nil
}

func (portSlots *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range portSlots.PortSlot {
        children[portSlots.PortSlot[i].GetSegmentPath()] = &portSlots.PortSlot[i]
    }
    return children
}

func (portSlots *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (portSlots *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots) GetBundleName() string { return "cisco_ios_xr" }

func (portSlots *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots) GetYangName() string { return "port-slots" }

func (portSlots *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (portSlots *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (portSlots *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (portSlots *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots) SetParent(parent types.Entity) { portSlots.parent = parent }

func (portSlots *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots) GetParent() types.Entity { return portSlots.parent }

func (portSlots *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots) GetParentYangName() string { return "module" }

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot
// PortSlot number in the SPA/PLIM
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. portslot number. The type is interface{} with
    // range: -2147483648..2147483647.
    Number interface{}

    // Port string.
    Port Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port

    // Attributes.
    BasicAttributes Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes
}

func (portSlot *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot) GetFilter() yfilter.YFilter { return portSlot.YFilter }

func (portSlot *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot) SetFilter(yf yfilter.YFilter) { portSlot.YFilter = yf }

func (portSlot *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot) GetGoName(yname string) string {
    if yname == "number" { return "Number" }
    if yname == "port" { return "Port" }
    if yname == "basic-attributes" { return "BasicAttributes" }
    return ""
}

func (portSlot *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot) GetSegmentPath() string {
    return "port-slot" + "[number='" + fmt.Sprintf("%v", portSlot.Number) + "']"
}

func (portSlot *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "port" {
        return &portSlot.Port
    }
    if childYangName == "basic-attributes" {
        return &portSlot.BasicAttributes
    }
    return nil
}

func (portSlot *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["port"] = &portSlot.Port
    children["basic-attributes"] = &portSlot.BasicAttributes
    return children
}

func (portSlot *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["number"] = portSlot.Number
    return leafs
}

func (portSlot *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot) GetBundleName() string { return "cisco_ios_xr" }

func (portSlot *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot) GetYangName() string { return "port-slot" }

func (portSlot *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (portSlot *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (portSlot *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (portSlot *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot) SetParent(parent types.Entity) { portSlot.parent = parent }

func (portSlot *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot) GetParent() types.Entity { return portSlot.parent }

func (portSlot *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot) GetParentYangName() string { return "port-slots" }

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port
// Port string
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Attributes.
    BasicAttributes Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes
}

func (port *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port) GetFilter() yfilter.YFilter { return port.YFilter }

func (port *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port) SetFilter(yf yfilter.YFilter) { port.YFilter = yf }

func (port *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port) GetGoName(yname string) string {
    if yname == "basic-attributes" { return "BasicAttributes" }
    return ""
}

func (port *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port) GetSegmentPath() string {
    return "port"
}

func (port *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "basic-attributes" {
        return &port.BasicAttributes
    }
    return nil
}

func (port *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["basic-attributes"] = &port.BasicAttributes
    return children
}

func (port *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (port *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port) GetBundleName() string { return "cisco_ios_xr" }

func (port *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port) GetYangName() string { return "port" }

func (port *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (port *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (port *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (port *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port) SetParent(parent types.Entity) { port.parent = parent }

func (port *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port) GetParent() types.Entity { return port.parent }

func (port *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port) GetParentYangName() string { return "port-slot" }

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes
// Attributes
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Field Replaceable Unit (FRU) inventory information.
    FruInfo Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_FruInfo

    // Inventory information.
    BasicInfo Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_BasicInfo
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes) GetFilter() yfilter.YFilter { return basicAttributes.YFilter }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes) SetFilter(yf yfilter.YFilter) { basicAttributes.YFilter = yf }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes) GetGoName(yname string) string {
    if yname == "fru-info" { return "FruInfo" }
    if yname == "basic-info" { return "BasicInfo" }
    return ""
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes) GetSegmentPath() string {
    return "basic-attributes"
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "fru-info" {
        return &basicAttributes.FruInfo
    }
    if childYangName == "basic-info" {
        return &basicAttributes.BasicInfo
    }
    return nil
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["fru-info"] = &basicAttributes.FruInfo
    children["basic-info"] = &basicAttributes.BasicInfo
    return children
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes) GetBundleName() string { return "cisco_ios_xr" }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes) GetYangName() string { return "basic-attributes" }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes) SetParent(parent types.Entity) { basicAttributes.parent = parent }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes) GetParent() types.Entity { return basicAttributes.parent }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes) GetParentYangName() string { return "port" }

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_FruInfo
// Field Replaceable Unit (FRU) inventory
// information
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_FruInfo struct {
    parent types.Entity
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

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_FruInfo) GetFilter() yfilter.YFilter { return fruInfo.YFilter }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_FruInfo) SetFilter(yf yfilter.YFilter) { fruInfo.YFilter = yf }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_FruInfo) GetGoName(yname string) string {
    if yname == "card-administrative-state" { return "CardAdministrativeState" }
    if yname == "power-administrative-state" { return "PowerAdministrativeState" }
    if yname == "card-operational-state" { return "CardOperationalState" }
    if yname == "card-monitor-state" { return "CardMonitorState" }
    if yname == "card-reset-reason" { return "CardResetReason" }
    if yname == "power-current-measurement" { return "PowerCurrentMeasurement" }
    if yname == "power-operational-state" { return "PowerOperationalState" }
    if yname == "last-operational-state-change" { return "LastOperationalStateChange" }
    if yname == "card-up-time" { return "CardUpTime" }
    return ""
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_FruInfo) GetSegmentPath() string {
    return "fru-info"
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_FruInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "last-operational-state-change" {
        return &fruInfo.LastOperationalStateChange
    }
    if childYangName == "card-up-time" {
        return &fruInfo.CardUpTime
    }
    return nil
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_FruInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["last-operational-state-change"] = &fruInfo.LastOperationalStateChange
    children["card-up-time"] = &fruInfo.CardUpTime
    return children
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_FruInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["card-administrative-state"] = fruInfo.CardAdministrativeState
    leafs["power-administrative-state"] = fruInfo.PowerAdministrativeState
    leafs["card-operational-state"] = fruInfo.CardOperationalState
    leafs["card-monitor-state"] = fruInfo.CardMonitorState
    leafs["card-reset-reason"] = fruInfo.CardResetReason
    leafs["power-current-measurement"] = fruInfo.PowerCurrentMeasurement
    leafs["power-operational-state"] = fruInfo.PowerOperationalState
    return leafs
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_FruInfo) GetBundleName() string { return "cisco_ios_xr" }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_FruInfo) GetYangName() string { return "fru-info" }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_FruInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_FruInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_FruInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_FruInfo) SetParent(parent types.Entity) { fruInfo.parent = parent }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_FruInfo) GetParent() types.Entity { return fruInfo.parent }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_FruInfo) GetParentYangName() string { return "basic-attributes" }

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_LastOperationalStateChange
// last card oper change state
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_LastOperationalStateChange struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_LastOperationalStateChange) GetFilter() yfilter.YFilter { return lastOperationalStateChange.YFilter }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_LastOperationalStateChange) SetFilter(yf yfilter.YFilter) { lastOperationalStateChange.YFilter = yf }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_LastOperationalStateChange) GetGoName(yname string) string {
    if yname == "time-in-seconds" { return "TimeInSeconds" }
    if yname == "time-in-nano-seconds" { return "TimeInNanoSeconds" }
    return ""
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_LastOperationalStateChange) GetSegmentPath() string {
    return "last-operational-state-change"
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_LastOperationalStateChange) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_LastOperationalStateChange) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_LastOperationalStateChange) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-in-seconds"] = lastOperationalStateChange.TimeInSeconds
    leafs["time-in-nano-seconds"] = lastOperationalStateChange.TimeInNanoSeconds
    return leafs
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_LastOperationalStateChange) GetBundleName() string { return "cisco_ios_xr" }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_LastOperationalStateChange) GetYangName() string { return "last-operational-state-change" }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_LastOperationalStateChange) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_LastOperationalStateChange) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_LastOperationalStateChange) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_LastOperationalStateChange) SetParent(parent types.Entity) { lastOperationalStateChange.parent = parent }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_LastOperationalStateChange) GetParent() types.Entity { return lastOperationalStateChange.parent }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_LastOperationalStateChange) GetParentYangName() string { return "fru-info" }

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_CardUpTime
// card up time
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_CardUpTime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_CardUpTime) GetFilter() yfilter.YFilter { return cardUpTime.YFilter }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_CardUpTime) SetFilter(yf yfilter.YFilter) { cardUpTime.YFilter = yf }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_CardUpTime) GetGoName(yname string) string {
    if yname == "time-in-seconds" { return "TimeInSeconds" }
    if yname == "time-in-nano-seconds" { return "TimeInNanoSeconds" }
    return ""
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_CardUpTime) GetSegmentPath() string {
    return "card-up-time"
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_CardUpTime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_CardUpTime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_CardUpTime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-in-seconds"] = cardUpTime.TimeInSeconds
    leafs["time-in-nano-seconds"] = cardUpTime.TimeInNanoSeconds
    return leafs
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_CardUpTime) GetBundleName() string { return "cisco_ios_xr" }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_CardUpTime) GetYangName() string { return "card-up-time" }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_CardUpTime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_CardUpTime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_CardUpTime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_CardUpTime) SetParent(parent types.Entity) { cardUpTime.parent = parent }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_CardUpTime) GetParent() types.Entity { return cardUpTime.parent }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_CardUpTime) GetParentYangName() string { return "fru-info" }

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_BasicInfo
// Inventory information
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_BasicInfo struct {
    parent types.Entity
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

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_BasicInfo) GetFilter() yfilter.YFilter { return basicInfo.YFilter }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_BasicInfo) SetFilter(yf yfilter.YFilter) { basicInfo.YFilter = yf }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_BasicInfo) GetGoName(yname string) string {
    if yname == "description" { return "Description" }
    if yname == "vendor-type" { return "VendorType" }
    if yname == "name" { return "Name" }
    if yname == "hardware-revision" { return "HardwareRevision" }
    if yname == "firmware-revision" { return "FirmwareRevision" }
    if yname == "software-revision" { return "SoftwareRevision" }
    if yname == "chip-hardware-revision" { return "ChipHardwareRevision" }
    if yname == "serial-number" { return "SerialNumber" }
    if yname == "manufacturer-name" { return "ManufacturerName" }
    if yname == "model-name" { return "ModelName" }
    if yname == "asset-id-str" { return "AssetIdStr" }
    if yname == "asset-identification" { return "AssetIdentification" }
    if yname == "is-field-replaceable-unit" { return "IsFieldReplaceableUnit" }
    if yname == "manufacturer-asset-tags" { return "ManufacturerAssetTags" }
    if yname == "composite-class-code" { return "CompositeClassCode" }
    if yname == "memory-size" { return "MemorySize" }
    if yname == "environmental-monitor-path" { return "EnvironmentalMonitorPath" }
    if yname == "alias" { return "Alias" }
    if yname == "group-flag" { return "GroupFlag" }
    if yname == "new-deviation-number" { return "NewDeviationNumber" }
    if yname == "physical-layer-interface-module-type" { return "PhysicalLayerInterfaceModuleType" }
    if yname == "unrecognized-fru" { return "UnrecognizedFru" }
    if yname == "redundancystate" { return "Redundancystate" }
    if yname == "ceport" { return "Ceport" }
    if yname == "xr-scoped" { return "XrScoped" }
    if yname == "unique-id" { return "UniqueId" }
    return ""
}

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_BasicInfo) GetSegmentPath() string {
    return "basic-info"
}

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_BasicInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_BasicInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_BasicInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["description"] = basicInfo.Description
    leafs["vendor-type"] = basicInfo.VendorType
    leafs["name"] = basicInfo.Name
    leafs["hardware-revision"] = basicInfo.HardwareRevision
    leafs["firmware-revision"] = basicInfo.FirmwareRevision
    leafs["software-revision"] = basicInfo.SoftwareRevision
    leafs["chip-hardware-revision"] = basicInfo.ChipHardwareRevision
    leafs["serial-number"] = basicInfo.SerialNumber
    leafs["manufacturer-name"] = basicInfo.ManufacturerName
    leafs["model-name"] = basicInfo.ModelName
    leafs["asset-id-str"] = basicInfo.AssetIdStr
    leafs["asset-identification"] = basicInfo.AssetIdentification
    leafs["is-field-replaceable-unit"] = basicInfo.IsFieldReplaceableUnit
    leafs["manufacturer-asset-tags"] = basicInfo.ManufacturerAssetTags
    leafs["composite-class-code"] = basicInfo.CompositeClassCode
    leafs["memory-size"] = basicInfo.MemorySize
    leafs["environmental-monitor-path"] = basicInfo.EnvironmentalMonitorPath
    leafs["alias"] = basicInfo.Alias
    leafs["group-flag"] = basicInfo.GroupFlag
    leafs["new-deviation-number"] = basicInfo.NewDeviationNumber
    leafs["physical-layer-interface-module-type"] = basicInfo.PhysicalLayerInterfaceModuleType
    leafs["unrecognized-fru"] = basicInfo.UnrecognizedFru
    leafs["redundancystate"] = basicInfo.Redundancystate
    leafs["ceport"] = basicInfo.Ceport
    leafs["xr-scoped"] = basicInfo.XrScoped
    leafs["unique-id"] = basicInfo.UniqueId
    return leafs
}

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_BasicInfo) GetBundleName() string { return "cisco_ios_xr" }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_BasicInfo) GetYangName() string { return "basic-info" }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_BasicInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_BasicInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_BasicInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_BasicInfo) SetParent(parent types.Entity) { basicInfo.parent = parent }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_BasicInfo) GetParent() types.Entity { return basicInfo.parent }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Port_BasicAttributes_BasicInfo) GetParentYangName() string { return "basic-attributes" }

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes
// Attributes
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Field Replaceable Unit (FRU) inventory information.
    FruInfo Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_FruInfo

    // Inventory information.
    BasicInfo Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_BasicInfo
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes) GetFilter() yfilter.YFilter { return basicAttributes.YFilter }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes) SetFilter(yf yfilter.YFilter) { basicAttributes.YFilter = yf }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes) GetGoName(yname string) string {
    if yname == "fru-info" { return "FruInfo" }
    if yname == "basic-info" { return "BasicInfo" }
    return ""
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes) GetSegmentPath() string {
    return "basic-attributes"
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "fru-info" {
        return &basicAttributes.FruInfo
    }
    if childYangName == "basic-info" {
        return &basicAttributes.BasicInfo
    }
    return nil
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["fru-info"] = &basicAttributes.FruInfo
    children["basic-info"] = &basicAttributes.BasicInfo
    return children
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes) GetBundleName() string { return "cisco_ios_xr" }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes) GetYangName() string { return "basic-attributes" }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes) SetParent(parent types.Entity) { basicAttributes.parent = parent }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes) GetParent() types.Entity { return basicAttributes.parent }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes) GetParentYangName() string { return "port-slot" }

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_FruInfo
// Field Replaceable Unit (FRU) inventory
// information
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_FruInfo struct {
    parent types.Entity
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

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_FruInfo) GetFilter() yfilter.YFilter { return fruInfo.YFilter }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_FruInfo) SetFilter(yf yfilter.YFilter) { fruInfo.YFilter = yf }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_FruInfo) GetGoName(yname string) string {
    if yname == "card-administrative-state" { return "CardAdministrativeState" }
    if yname == "power-administrative-state" { return "PowerAdministrativeState" }
    if yname == "card-operational-state" { return "CardOperationalState" }
    if yname == "card-monitor-state" { return "CardMonitorState" }
    if yname == "card-reset-reason" { return "CardResetReason" }
    if yname == "power-current-measurement" { return "PowerCurrentMeasurement" }
    if yname == "power-operational-state" { return "PowerOperationalState" }
    if yname == "last-operational-state-change" { return "LastOperationalStateChange" }
    if yname == "card-up-time" { return "CardUpTime" }
    return ""
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_FruInfo) GetSegmentPath() string {
    return "fru-info"
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_FruInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "last-operational-state-change" {
        return &fruInfo.LastOperationalStateChange
    }
    if childYangName == "card-up-time" {
        return &fruInfo.CardUpTime
    }
    return nil
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_FruInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["last-operational-state-change"] = &fruInfo.LastOperationalStateChange
    children["card-up-time"] = &fruInfo.CardUpTime
    return children
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_FruInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["card-administrative-state"] = fruInfo.CardAdministrativeState
    leafs["power-administrative-state"] = fruInfo.PowerAdministrativeState
    leafs["card-operational-state"] = fruInfo.CardOperationalState
    leafs["card-monitor-state"] = fruInfo.CardMonitorState
    leafs["card-reset-reason"] = fruInfo.CardResetReason
    leafs["power-current-measurement"] = fruInfo.PowerCurrentMeasurement
    leafs["power-operational-state"] = fruInfo.PowerOperationalState
    return leafs
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_FruInfo) GetBundleName() string { return "cisco_ios_xr" }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_FruInfo) GetYangName() string { return "fru-info" }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_FruInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_FruInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_FruInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_FruInfo) SetParent(parent types.Entity) { fruInfo.parent = parent }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_FruInfo) GetParent() types.Entity { return fruInfo.parent }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_FruInfo) GetParentYangName() string { return "basic-attributes" }

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_FruInfo_LastOperationalStateChange
// last card oper change state
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_FruInfo_LastOperationalStateChange struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_FruInfo_LastOperationalStateChange) GetFilter() yfilter.YFilter { return lastOperationalStateChange.YFilter }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_FruInfo_LastOperationalStateChange) SetFilter(yf yfilter.YFilter) { lastOperationalStateChange.YFilter = yf }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_FruInfo_LastOperationalStateChange) GetGoName(yname string) string {
    if yname == "time-in-seconds" { return "TimeInSeconds" }
    if yname == "time-in-nano-seconds" { return "TimeInNanoSeconds" }
    return ""
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_FruInfo_LastOperationalStateChange) GetSegmentPath() string {
    return "last-operational-state-change"
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_FruInfo_LastOperationalStateChange) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_FruInfo_LastOperationalStateChange) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_FruInfo_LastOperationalStateChange) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-in-seconds"] = lastOperationalStateChange.TimeInSeconds
    leafs["time-in-nano-seconds"] = lastOperationalStateChange.TimeInNanoSeconds
    return leafs
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_FruInfo_LastOperationalStateChange) GetBundleName() string { return "cisco_ios_xr" }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_FruInfo_LastOperationalStateChange) GetYangName() string { return "last-operational-state-change" }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_FruInfo_LastOperationalStateChange) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_FruInfo_LastOperationalStateChange) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_FruInfo_LastOperationalStateChange) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_FruInfo_LastOperationalStateChange) SetParent(parent types.Entity) { lastOperationalStateChange.parent = parent }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_FruInfo_LastOperationalStateChange) GetParent() types.Entity { return lastOperationalStateChange.parent }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_FruInfo_LastOperationalStateChange) GetParentYangName() string { return "fru-info" }

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_FruInfo_CardUpTime
// card up time
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_FruInfo_CardUpTime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_FruInfo_CardUpTime) GetFilter() yfilter.YFilter { return cardUpTime.YFilter }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_FruInfo_CardUpTime) SetFilter(yf yfilter.YFilter) { cardUpTime.YFilter = yf }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_FruInfo_CardUpTime) GetGoName(yname string) string {
    if yname == "time-in-seconds" { return "TimeInSeconds" }
    if yname == "time-in-nano-seconds" { return "TimeInNanoSeconds" }
    return ""
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_FruInfo_CardUpTime) GetSegmentPath() string {
    return "card-up-time"
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_FruInfo_CardUpTime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_FruInfo_CardUpTime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_FruInfo_CardUpTime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-in-seconds"] = cardUpTime.TimeInSeconds
    leafs["time-in-nano-seconds"] = cardUpTime.TimeInNanoSeconds
    return leafs
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_FruInfo_CardUpTime) GetBundleName() string { return "cisco_ios_xr" }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_FruInfo_CardUpTime) GetYangName() string { return "card-up-time" }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_FruInfo_CardUpTime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_FruInfo_CardUpTime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_FruInfo_CardUpTime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_FruInfo_CardUpTime) SetParent(parent types.Entity) { cardUpTime.parent = parent }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_FruInfo_CardUpTime) GetParent() types.Entity { return cardUpTime.parent }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_FruInfo_CardUpTime) GetParentYangName() string { return "fru-info" }

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_BasicInfo
// Inventory information
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_BasicInfo struct {
    parent types.Entity
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

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_BasicInfo) GetFilter() yfilter.YFilter { return basicInfo.YFilter }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_BasicInfo) SetFilter(yf yfilter.YFilter) { basicInfo.YFilter = yf }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_BasicInfo) GetGoName(yname string) string {
    if yname == "description" { return "Description" }
    if yname == "vendor-type" { return "VendorType" }
    if yname == "name" { return "Name" }
    if yname == "hardware-revision" { return "HardwareRevision" }
    if yname == "firmware-revision" { return "FirmwareRevision" }
    if yname == "software-revision" { return "SoftwareRevision" }
    if yname == "chip-hardware-revision" { return "ChipHardwareRevision" }
    if yname == "serial-number" { return "SerialNumber" }
    if yname == "manufacturer-name" { return "ManufacturerName" }
    if yname == "model-name" { return "ModelName" }
    if yname == "asset-id-str" { return "AssetIdStr" }
    if yname == "asset-identification" { return "AssetIdentification" }
    if yname == "is-field-replaceable-unit" { return "IsFieldReplaceableUnit" }
    if yname == "manufacturer-asset-tags" { return "ManufacturerAssetTags" }
    if yname == "composite-class-code" { return "CompositeClassCode" }
    if yname == "memory-size" { return "MemorySize" }
    if yname == "environmental-monitor-path" { return "EnvironmentalMonitorPath" }
    if yname == "alias" { return "Alias" }
    if yname == "group-flag" { return "GroupFlag" }
    if yname == "new-deviation-number" { return "NewDeviationNumber" }
    if yname == "physical-layer-interface-module-type" { return "PhysicalLayerInterfaceModuleType" }
    if yname == "unrecognized-fru" { return "UnrecognizedFru" }
    if yname == "redundancystate" { return "Redundancystate" }
    if yname == "ceport" { return "Ceport" }
    if yname == "xr-scoped" { return "XrScoped" }
    if yname == "unique-id" { return "UniqueId" }
    return ""
}

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_BasicInfo) GetSegmentPath() string {
    return "basic-info"
}

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_BasicInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_BasicInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_BasicInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["description"] = basicInfo.Description
    leafs["vendor-type"] = basicInfo.VendorType
    leafs["name"] = basicInfo.Name
    leafs["hardware-revision"] = basicInfo.HardwareRevision
    leafs["firmware-revision"] = basicInfo.FirmwareRevision
    leafs["software-revision"] = basicInfo.SoftwareRevision
    leafs["chip-hardware-revision"] = basicInfo.ChipHardwareRevision
    leafs["serial-number"] = basicInfo.SerialNumber
    leafs["manufacturer-name"] = basicInfo.ManufacturerName
    leafs["model-name"] = basicInfo.ModelName
    leafs["asset-id-str"] = basicInfo.AssetIdStr
    leafs["asset-identification"] = basicInfo.AssetIdentification
    leafs["is-field-replaceable-unit"] = basicInfo.IsFieldReplaceableUnit
    leafs["manufacturer-asset-tags"] = basicInfo.ManufacturerAssetTags
    leafs["composite-class-code"] = basicInfo.CompositeClassCode
    leafs["memory-size"] = basicInfo.MemorySize
    leafs["environmental-monitor-path"] = basicInfo.EnvironmentalMonitorPath
    leafs["alias"] = basicInfo.Alias
    leafs["group-flag"] = basicInfo.GroupFlag
    leafs["new-deviation-number"] = basicInfo.NewDeviationNumber
    leafs["physical-layer-interface-module-type"] = basicInfo.PhysicalLayerInterfaceModuleType
    leafs["unrecognized-fru"] = basicInfo.UnrecognizedFru
    leafs["redundancystate"] = basicInfo.Redundancystate
    leafs["ceport"] = basicInfo.Ceport
    leafs["xr-scoped"] = basicInfo.XrScoped
    leafs["unique-id"] = basicInfo.UniqueId
    return leafs
}

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_BasicInfo) GetBundleName() string { return "cisco_ios_xr" }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_BasicInfo) GetYangName() string { return "basic-info" }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_BasicInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_BasicInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_BasicInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_BasicInfo) SetParent(parent types.Entity) { basicInfo.parent = parent }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_BasicInfo) GetParent() types.Entity { return basicInfo.parent }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_BasicAttributes_BasicInfo) GetParentYangName() string { return "basic-attributes" }

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors
// ModuleSensorTable contains all sensors in a
// Module.
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Sensor number in the Module. The type is slice of
    // Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor.
    Sensor []Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor
}

func (sensors *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors) GetFilter() yfilter.YFilter { return sensors.YFilter }

func (sensors *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors) SetFilter(yf yfilter.YFilter) { sensors.YFilter = yf }

func (sensors *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors) GetGoName(yname string) string {
    if yname == "sensor" { return "Sensor" }
    return ""
}

func (sensors *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors) GetSegmentPath() string {
    return "sensors"
}

func (sensors *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sensor" {
        for _, c := range sensors.Sensor {
            if sensors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor{}
        sensors.Sensor = append(sensors.Sensor, child)
        return &sensors.Sensor[len(sensors.Sensor)-1]
    }
    return nil
}

func (sensors *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sensors.Sensor {
        children[sensors.Sensor[i].GetSegmentPath()] = &sensors.Sensor[i]
    }
    return children
}

func (sensors *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sensors *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors) GetBundleName() string { return "cisco_ios_xr" }

func (sensors *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors) GetYangName() string { return "sensors" }

func (sensors *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sensors *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sensors *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sensors *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors) SetParent(parent types.Entity) { sensors.parent = parent }

func (sensors *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors) GetParent() types.Entity { return sensors.parent }

func (sensors *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors) GetParentYangName() string { return "module" }

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor
// Sensor number in the Module
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. sensor number. The type is interface{} with range:
    // -2147483648..2147483647.
    Number interface{}

    // Attributes.
    BasicAttributes Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes
}

func (sensor *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor) GetFilter() yfilter.YFilter { return sensor.YFilter }

func (sensor *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor) SetFilter(yf yfilter.YFilter) { sensor.YFilter = yf }

func (sensor *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor) GetGoName(yname string) string {
    if yname == "number" { return "Number" }
    if yname == "basic-attributes" { return "BasicAttributes" }
    return ""
}

func (sensor *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor) GetSegmentPath() string {
    return "sensor" + "[number='" + fmt.Sprintf("%v", sensor.Number) + "']"
}

func (sensor *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "basic-attributes" {
        return &sensor.BasicAttributes
    }
    return nil
}

func (sensor *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["basic-attributes"] = &sensor.BasicAttributes
    return children
}

func (sensor *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["number"] = sensor.Number
    return leafs
}

func (sensor *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor) GetBundleName() string { return "cisco_ios_xr" }

func (sensor *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor) GetYangName() string { return "sensor" }

func (sensor *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sensor *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sensor *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sensor *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor) SetParent(parent types.Entity) { sensor.parent = parent }

func (sensor *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor) GetParent() types.Entity { return sensor.parent }

func (sensor *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor) GetParentYangName() string { return "sensors" }

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes
// Attributes
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Field Replaceable Unit (FRU) inventory information.
    FruInfo Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_FruInfo

    // Inventory information.
    BasicInfo Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_BasicInfo
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes) GetFilter() yfilter.YFilter { return basicAttributes.YFilter }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes) SetFilter(yf yfilter.YFilter) { basicAttributes.YFilter = yf }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes) GetGoName(yname string) string {
    if yname == "fru-info" { return "FruInfo" }
    if yname == "basic-info" { return "BasicInfo" }
    return ""
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes) GetSegmentPath() string {
    return "basic-attributes"
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "fru-info" {
        return &basicAttributes.FruInfo
    }
    if childYangName == "basic-info" {
        return &basicAttributes.BasicInfo
    }
    return nil
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["fru-info"] = &basicAttributes.FruInfo
    children["basic-info"] = &basicAttributes.BasicInfo
    return children
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes) GetBundleName() string { return "cisco_ios_xr" }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes) GetYangName() string { return "basic-attributes" }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes) SetParent(parent types.Entity) { basicAttributes.parent = parent }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes) GetParent() types.Entity { return basicAttributes.parent }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes) GetParentYangName() string { return "sensor" }

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_FruInfo
// Field Replaceable Unit (FRU) inventory
// information
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_FruInfo struct {
    parent types.Entity
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

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_FruInfo) GetFilter() yfilter.YFilter { return fruInfo.YFilter }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_FruInfo) SetFilter(yf yfilter.YFilter) { fruInfo.YFilter = yf }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_FruInfo) GetGoName(yname string) string {
    if yname == "card-administrative-state" { return "CardAdministrativeState" }
    if yname == "power-administrative-state" { return "PowerAdministrativeState" }
    if yname == "card-operational-state" { return "CardOperationalState" }
    if yname == "card-monitor-state" { return "CardMonitorState" }
    if yname == "card-reset-reason" { return "CardResetReason" }
    if yname == "power-current-measurement" { return "PowerCurrentMeasurement" }
    if yname == "power-operational-state" { return "PowerOperationalState" }
    if yname == "last-operational-state-change" { return "LastOperationalStateChange" }
    if yname == "card-up-time" { return "CardUpTime" }
    return ""
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_FruInfo) GetSegmentPath() string {
    return "fru-info"
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_FruInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "last-operational-state-change" {
        return &fruInfo.LastOperationalStateChange
    }
    if childYangName == "card-up-time" {
        return &fruInfo.CardUpTime
    }
    return nil
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_FruInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["last-operational-state-change"] = &fruInfo.LastOperationalStateChange
    children["card-up-time"] = &fruInfo.CardUpTime
    return children
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_FruInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["card-administrative-state"] = fruInfo.CardAdministrativeState
    leafs["power-administrative-state"] = fruInfo.PowerAdministrativeState
    leafs["card-operational-state"] = fruInfo.CardOperationalState
    leafs["card-monitor-state"] = fruInfo.CardMonitorState
    leafs["card-reset-reason"] = fruInfo.CardResetReason
    leafs["power-current-measurement"] = fruInfo.PowerCurrentMeasurement
    leafs["power-operational-state"] = fruInfo.PowerOperationalState
    return leafs
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_FruInfo) GetBundleName() string { return "cisco_ios_xr" }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_FruInfo) GetYangName() string { return "fru-info" }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_FruInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_FruInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_FruInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_FruInfo) SetParent(parent types.Entity) { fruInfo.parent = parent }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_FruInfo) GetParent() types.Entity { return fruInfo.parent }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_FruInfo) GetParentYangName() string { return "basic-attributes" }

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_FruInfo_LastOperationalStateChange
// last card oper change state
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_FruInfo_LastOperationalStateChange struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_FruInfo_LastOperationalStateChange) GetFilter() yfilter.YFilter { return lastOperationalStateChange.YFilter }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_FruInfo_LastOperationalStateChange) SetFilter(yf yfilter.YFilter) { lastOperationalStateChange.YFilter = yf }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_FruInfo_LastOperationalStateChange) GetGoName(yname string) string {
    if yname == "time-in-seconds" { return "TimeInSeconds" }
    if yname == "time-in-nano-seconds" { return "TimeInNanoSeconds" }
    return ""
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_FruInfo_LastOperationalStateChange) GetSegmentPath() string {
    return "last-operational-state-change"
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_FruInfo_LastOperationalStateChange) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_FruInfo_LastOperationalStateChange) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_FruInfo_LastOperationalStateChange) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-in-seconds"] = lastOperationalStateChange.TimeInSeconds
    leafs["time-in-nano-seconds"] = lastOperationalStateChange.TimeInNanoSeconds
    return leafs
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_FruInfo_LastOperationalStateChange) GetBundleName() string { return "cisco_ios_xr" }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_FruInfo_LastOperationalStateChange) GetYangName() string { return "last-operational-state-change" }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_FruInfo_LastOperationalStateChange) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_FruInfo_LastOperationalStateChange) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_FruInfo_LastOperationalStateChange) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_FruInfo_LastOperationalStateChange) SetParent(parent types.Entity) { lastOperationalStateChange.parent = parent }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_FruInfo_LastOperationalStateChange) GetParent() types.Entity { return lastOperationalStateChange.parent }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_FruInfo_LastOperationalStateChange) GetParentYangName() string { return "fru-info" }

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_FruInfo_CardUpTime
// card up time
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_FruInfo_CardUpTime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_FruInfo_CardUpTime) GetFilter() yfilter.YFilter { return cardUpTime.YFilter }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_FruInfo_CardUpTime) SetFilter(yf yfilter.YFilter) { cardUpTime.YFilter = yf }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_FruInfo_CardUpTime) GetGoName(yname string) string {
    if yname == "time-in-seconds" { return "TimeInSeconds" }
    if yname == "time-in-nano-seconds" { return "TimeInNanoSeconds" }
    return ""
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_FruInfo_CardUpTime) GetSegmentPath() string {
    return "card-up-time"
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_FruInfo_CardUpTime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_FruInfo_CardUpTime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_FruInfo_CardUpTime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-in-seconds"] = cardUpTime.TimeInSeconds
    leafs["time-in-nano-seconds"] = cardUpTime.TimeInNanoSeconds
    return leafs
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_FruInfo_CardUpTime) GetBundleName() string { return "cisco_ios_xr" }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_FruInfo_CardUpTime) GetYangName() string { return "card-up-time" }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_FruInfo_CardUpTime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_FruInfo_CardUpTime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_FruInfo_CardUpTime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_FruInfo_CardUpTime) SetParent(parent types.Entity) { cardUpTime.parent = parent }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_FruInfo_CardUpTime) GetParent() types.Entity { return cardUpTime.parent }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_FruInfo_CardUpTime) GetParentYangName() string { return "fru-info" }

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_BasicInfo
// Inventory information
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_BasicInfo struct {
    parent types.Entity
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

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_BasicInfo) GetFilter() yfilter.YFilter { return basicInfo.YFilter }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_BasicInfo) SetFilter(yf yfilter.YFilter) { basicInfo.YFilter = yf }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_BasicInfo) GetGoName(yname string) string {
    if yname == "description" { return "Description" }
    if yname == "vendor-type" { return "VendorType" }
    if yname == "name" { return "Name" }
    if yname == "hardware-revision" { return "HardwareRevision" }
    if yname == "firmware-revision" { return "FirmwareRevision" }
    if yname == "software-revision" { return "SoftwareRevision" }
    if yname == "chip-hardware-revision" { return "ChipHardwareRevision" }
    if yname == "serial-number" { return "SerialNumber" }
    if yname == "manufacturer-name" { return "ManufacturerName" }
    if yname == "model-name" { return "ModelName" }
    if yname == "asset-id-str" { return "AssetIdStr" }
    if yname == "asset-identification" { return "AssetIdentification" }
    if yname == "is-field-replaceable-unit" { return "IsFieldReplaceableUnit" }
    if yname == "manufacturer-asset-tags" { return "ManufacturerAssetTags" }
    if yname == "composite-class-code" { return "CompositeClassCode" }
    if yname == "memory-size" { return "MemorySize" }
    if yname == "environmental-monitor-path" { return "EnvironmentalMonitorPath" }
    if yname == "alias" { return "Alias" }
    if yname == "group-flag" { return "GroupFlag" }
    if yname == "new-deviation-number" { return "NewDeviationNumber" }
    if yname == "physical-layer-interface-module-type" { return "PhysicalLayerInterfaceModuleType" }
    if yname == "unrecognized-fru" { return "UnrecognizedFru" }
    if yname == "redundancystate" { return "Redundancystate" }
    if yname == "ceport" { return "Ceport" }
    if yname == "xr-scoped" { return "XrScoped" }
    if yname == "unique-id" { return "UniqueId" }
    return ""
}

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_BasicInfo) GetSegmentPath() string {
    return "basic-info"
}

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_BasicInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_BasicInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_BasicInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["description"] = basicInfo.Description
    leafs["vendor-type"] = basicInfo.VendorType
    leafs["name"] = basicInfo.Name
    leafs["hardware-revision"] = basicInfo.HardwareRevision
    leafs["firmware-revision"] = basicInfo.FirmwareRevision
    leafs["software-revision"] = basicInfo.SoftwareRevision
    leafs["chip-hardware-revision"] = basicInfo.ChipHardwareRevision
    leafs["serial-number"] = basicInfo.SerialNumber
    leafs["manufacturer-name"] = basicInfo.ManufacturerName
    leafs["model-name"] = basicInfo.ModelName
    leafs["asset-id-str"] = basicInfo.AssetIdStr
    leafs["asset-identification"] = basicInfo.AssetIdentification
    leafs["is-field-replaceable-unit"] = basicInfo.IsFieldReplaceableUnit
    leafs["manufacturer-asset-tags"] = basicInfo.ManufacturerAssetTags
    leafs["composite-class-code"] = basicInfo.CompositeClassCode
    leafs["memory-size"] = basicInfo.MemorySize
    leafs["environmental-monitor-path"] = basicInfo.EnvironmentalMonitorPath
    leafs["alias"] = basicInfo.Alias
    leafs["group-flag"] = basicInfo.GroupFlag
    leafs["new-deviation-number"] = basicInfo.NewDeviationNumber
    leafs["physical-layer-interface-module-type"] = basicInfo.PhysicalLayerInterfaceModuleType
    leafs["unrecognized-fru"] = basicInfo.UnrecognizedFru
    leafs["redundancystate"] = basicInfo.Redundancystate
    leafs["ceport"] = basicInfo.Ceport
    leafs["xr-scoped"] = basicInfo.XrScoped
    leafs["unique-id"] = basicInfo.UniqueId
    return leafs
}

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_BasicInfo) GetBundleName() string { return "cisco_ios_xr" }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_BasicInfo) GetYangName() string { return "basic-info" }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_BasicInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_BasicInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_BasicInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_BasicInfo) SetParent(parent types.Entity) { basicInfo.parent = parent }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_BasicInfo) GetParent() types.Entity { return basicInfo.parent }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_BasicAttributes_BasicInfo) GetParentYangName() string { return "basic-attributes" }

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes
// Attributes
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Field Replaceable Unit (FRU) inventory information.
    FruInfo Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_FruInfo

    // Inventory information.
    BasicInfo Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_BasicInfo
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes) GetFilter() yfilter.YFilter { return basicAttributes.YFilter }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes) SetFilter(yf yfilter.YFilter) { basicAttributes.YFilter = yf }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes) GetGoName(yname string) string {
    if yname == "fru-info" { return "FruInfo" }
    if yname == "basic-info" { return "BasicInfo" }
    return ""
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes) GetSegmentPath() string {
    return "basic-attributes"
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "fru-info" {
        return &basicAttributes.FruInfo
    }
    if childYangName == "basic-info" {
        return &basicAttributes.BasicInfo
    }
    return nil
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["fru-info"] = &basicAttributes.FruInfo
    children["basic-info"] = &basicAttributes.BasicInfo
    return children
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes) GetBundleName() string { return "cisco_ios_xr" }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes) GetYangName() string { return "basic-attributes" }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes) SetParent(parent types.Entity) { basicAttributes.parent = parent }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes) GetParent() types.Entity { return basicAttributes.parent }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes) GetParentYangName() string { return "module" }

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_FruInfo
// Field Replaceable Unit (FRU) inventory
// information
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_FruInfo struct {
    parent types.Entity
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

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_FruInfo) GetFilter() yfilter.YFilter { return fruInfo.YFilter }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_FruInfo) SetFilter(yf yfilter.YFilter) { fruInfo.YFilter = yf }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_FruInfo) GetGoName(yname string) string {
    if yname == "card-administrative-state" { return "CardAdministrativeState" }
    if yname == "power-administrative-state" { return "PowerAdministrativeState" }
    if yname == "card-operational-state" { return "CardOperationalState" }
    if yname == "card-monitor-state" { return "CardMonitorState" }
    if yname == "card-reset-reason" { return "CardResetReason" }
    if yname == "power-current-measurement" { return "PowerCurrentMeasurement" }
    if yname == "power-operational-state" { return "PowerOperationalState" }
    if yname == "last-operational-state-change" { return "LastOperationalStateChange" }
    if yname == "card-up-time" { return "CardUpTime" }
    return ""
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_FruInfo) GetSegmentPath() string {
    return "fru-info"
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_FruInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "last-operational-state-change" {
        return &fruInfo.LastOperationalStateChange
    }
    if childYangName == "card-up-time" {
        return &fruInfo.CardUpTime
    }
    return nil
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_FruInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["last-operational-state-change"] = &fruInfo.LastOperationalStateChange
    children["card-up-time"] = &fruInfo.CardUpTime
    return children
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_FruInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["card-administrative-state"] = fruInfo.CardAdministrativeState
    leafs["power-administrative-state"] = fruInfo.PowerAdministrativeState
    leafs["card-operational-state"] = fruInfo.CardOperationalState
    leafs["card-monitor-state"] = fruInfo.CardMonitorState
    leafs["card-reset-reason"] = fruInfo.CardResetReason
    leafs["power-current-measurement"] = fruInfo.PowerCurrentMeasurement
    leafs["power-operational-state"] = fruInfo.PowerOperationalState
    return leafs
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_FruInfo) GetBundleName() string { return "cisco_ios_xr" }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_FruInfo) GetYangName() string { return "fru-info" }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_FruInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_FruInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_FruInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_FruInfo) SetParent(parent types.Entity) { fruInfo.parent = parent }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_FruInfo) GetParent() types.Entity { return fruInfo.parent }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_FruInfo) GetParentYangName() string { return "basic-attributes" }

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_FruInfo_LastOperationalStateChange
// last card oper change state
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_FruInfo_LastOperationalStateChange struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_FruInfo_LastOperationalStateChange) GetFilter() yfilter.YFilter { return lastOperationalStateChange.YFilter }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_FruInfo_LastOperationalStateChange) SetFilter(yf yfilter.YFilter) { lastOperationalStateChange.YFilter = yf }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_FruInfo_LastOperationalStateChange) GetGoName(yname string) string {
    if yname == "time-in-seconds" { return "TimeInSeconds" }
    if yname == "time-in-nano-seconds" { return "TimeInNanoSeconds" }
    return ""
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_FruInfo_LastOperationalStateChange) GetSegmentPath() string {
    return "last-operational-state-change"
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_FruInfo_LastOperationalStateChange) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_FruInfo_LastOperationalStateChange) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_FruInfo_LastOperationalStateChange) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-in-seconds"] = lastOperationalStateChange.TimeInSeconds
    leafs["time-in-nano-seconds"] = lastOperationalStateChange.TimeInNanoSeconds
    return leafs
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_FruInfo_LastOperationalStateChange) GetBundleName() string { return "cisco_ios_xr" }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_FruInfo_LastOperationalStateChange) GetYangName() string { return "last-operational-state-change" }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_FruInfo_LastOperationalStateChange) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_FruInfo_LastOperationalStateChange) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_FruInfo_LastOperationalStateChange) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_FruInfo_LastOperationalStateChange) SetParent(parent types.Entity) { lastOperationalStateChange.parent = parent }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_FruInfo_LastOperationalStateChange) GetParent() types.Entity { return lastOperationalStateChange.parent }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_FruInfo_LastOperationalStateChange) GetParentYangName() string { return "fru-info" }

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_FruInfo_CardUpTime
// card up time
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_FruInfo_CardUpTime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_FruInfo_CardUpTime) GetFilter() yfilter.YFilter { return cardUpTime.YFilter }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_FruInfo_CardUpTime) SetFilter(yf yfilter.YFilter) { cardUpTime.YFilter = yf }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_FruInfo_CardUpTime) GetGoName(yname string) string {
    if yname == "time-in-seconds" { return "TimeInSeconds" }
    if yname == "time-in-nano-seconds" { return "TimeInNanoSeconds" }
    return ""
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_FruInfo_CardUpTime) GetSegmentPath() string {
    return "card-up-time"
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_FruInfo_CardUpTime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_FruInfo_CardUpTime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_FruInfo_CardUpTime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-in-seconds"] = cardUpTime.TimeInSeconds
    leafs["time-in-nano-seconds"] = cardUpTime.TimeInNanoSeconds
    return leafs
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_FruInfo_CardUpTime) GetBundleName() string { return "cisco_ios_xr" }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_FruInfo_CardUpTime) GetYangName() string { return "card-up-time" }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_FruInfo_CardUpTime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_FruInfo_CardUpTime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_FruInfo_CardUpTime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_FruInfo_CardUpTime) SetParent(parent types.Entity) { cardUpTime.parent = parent }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_FruInfo_CardUpTime) GetParent() types.Entity { return cardUpTime.parent }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_FruInfo_CardUpTime) GetParentYangName() string { return "fru-info" }

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_BasicInfo
// Inventory information
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_BasicInfo struct {
    parent types.Entity
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

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_BasicInfo) GetFilter() yfilter.YFilter { return basicInfo.YFilter }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_BasicInfo) SetFilter(yf yfilter.YFilter) { basicInfo.YFilter = yf }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_BasicInfo) GetGoName(yname string) string {
    if yname == "description" { return "Description" }
    if yname == "vendor-type" { return "VendorType" }
    if yname == "name" { return "Name" }
    if yname == "hardware-revision" { return "HardwareRevision" }
    if yname == "firmware-revision" { return "FirmwareRevision" }
    if yname == "software-revision" { return "SoftwareRevision" }
    if yname == "chip-hardware-revision" { return "ChipHardwareRevision" }
    if yname == "serial-number" { return "SerialNumber" }
    if yname == "manufacturer-name" { return "ManufacturerName" }
    if yname == "model-name" { return "ModelName" }
    if yname == "asset-id-str" { return "AssetIdStr" }
    if yname == "asset-identification" { return "AssetIdentification" }
    if yname == "is-field-replaceable-unit" { return "IsFieldReplaceableUnit" }
    if yname == "manufacturer-asset-tags" { return "ManufacturerAssetTags" }
    if yname == "composite-class-code" { return "CompositeClassCode" }
    if yname == "memory-size" { return "MemorySize" }
    if yname == "environmental-monitor-path" { return "EnvironmentalMonitorPath" }
    if yname == "alias" { return "Alias" }
    if yname == "group-flag" { return "GroupFlag" }
    if yname == "new-deviation-number" { return "NewDeviationNumber" }
    if yname == "physical-layer-interface-module-type" { return "PhysicalLayerInterfaceModuleType" }
    if yname == "unrecognized-fru" { return "UnrecognizedFru" }
    if yname == "redundancystate" { return "Redundancystate" }
    if yname == "ceport" { return "Ceport" }
    if yname == "xr-scoped" { return "XrScoped" }
    if yname == "unique-id" { return "UniqueId" }
    return ""
}

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_BasicInfo) GetSegmentPath() string {
    return "basic-info"
}

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_BasicInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_BasicInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_BasicInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["description"] = basicInfo.Description
    leafs["vendor-type"] = basicInfo.VendorType
    leafs["name"] = basicInfo.Name
    leafs["hardware-revision"] = basicInfo.HardwareRevision
    leafs["firmware-revision"] = basicInfo.FirmwareRevision
    leafs["software-revision"] = basicInfo.SoftwareRevision
    leafs["chip-hardware-revision"] = basicInfo.ChipHardwareRevision
    leafs["serial-number"] = basicInfo.SerialNumber
    leafs["manufacturer-name"] = basicInfo.ManufacturerName
    leafs["model-name"] = basicInfo.ModelName
    leafs["asset-id-str"] = basicInfo.AssetIdStr
    leafs["asset-identification"] = basicInfo.AssetIdentification
    leafs["is-field-replaceable-unit"] = basicInfo.IsFieldReplaceableUnit
    leafs["manufacturer-asset-tags"] = basicInfo.ManufacturerAssetTags
    leafs["composite-class-code"] = basicInfo.CompositeClassCode
    leafs["memory-size"] = basicInfo.MemorySize
    leafs["environmental-monitor-path"] = basicInfo.EnvironmentalMonitorPath
    leafs["alias"] = basicInfo.Alias
    leafs["group-flag"] = basicInfo.GroupFlag
    leafs["new-deviation-number"] = basicInfo.NewDeviationNumber
    leafs["physical-layer-interface-module-type"] = basicInfo.PhysicalLayerInterfaceModuleType
    leafs["unrecognized-fru"] = basicInfo.UnrecognizedFru
    leafs["redundancystate"] = basicInfo.Redundancystate
    leafs["ceport"] = basicInfo.Ceport
    leafs["xr-scoped"] = basicInfo.XrScoped
    leafs["unique-id"] = basicInfo.UniqueId
    return leafs
}

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_BasicInfo) GetBundleName() string { return "cisco_ios_xr" }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_BasicInfo) GetYangName() string { return "basic-info" }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_BasicInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_BasicInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_BasicInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_BasicInfo) SetParent(parent types.Entity) { basicInfo.parent = parent }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_BasicInfo) GetParent() types.Entity { return basicInfo.parent }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_BasicAttributes_BasicInfo) GetParentYangName() string { return "basic-attributes" }

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes
// Attributes
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Field Replaceable Unit (FRU) inventory information.
    FruInfo Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_FruInfo

    // Inventory information.
    BasicInfo Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_BasicInfo
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes) GetFilter() yfilter.YFilter { return basicAttributes.YFilter }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes) SetFilter(yf yfilter.YFilter) { basicAttributes.YFilter = yf }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes) GetGoName(yname string) string {
    if yname == "fru-info" { return "FruInfo" }
    if yname == "basic-info" { return "BasicInfo" }
    return ""
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes) GetSegmentPath() string {
    return "basic-attributes"
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "fru-info" {
        return &basicAttributes.FruInfo
    }
    if childYangName == "basic-info" {
        return &basicAttributes.BasicInfo
    }
    return nil
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["fru-info"] = &basicAttributes.FruInfo
    children["basic-info"] = &basicAttributes.BasicInfo
    return children
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes) GetBundleName() string { return "cisco_ios_xr" }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes) GetYangName() string { return "basic-attributes" }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes) SetParent(parent types.Entity) { basicAttributes.parent = parent }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes) GetParent() types.Entity { return basicAttributes.parent }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes) GetParentYangName() string { return "sub-slot" }

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_FruInfo
// Field Replaceable Unit (FRU) inventory
// information
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_FruInfo struct {
    parent types.Entity
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

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_FruInfo) GetFilter() yfilter.YFilter { return fruInfo.YFilter }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_FruInfo) SetFilter(yf yfilter.YFilter) { fruInfo.YFilter = yf }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_FruInfo) GetGoName(yname string) string {
    if yname == "card-administrative-state" { return "CardAdministrativeState" }
    if yname == "power-administrative-state" { return "PowerAdministrativeState" }
    if yname == "card-operational-state" { return "CardOperationalState" }
    if yname == "card-monitor-state" { return "CardMonitorState" }
    if yname == "card-reset-reason" { return "CardResetReason" }
    if yname == "power-current-measurement" { return "PowerCurrentMeasurement" }
    if yname == "power-operational-state" { return "PowerOperationalState" }
    if yname == "last-operational-state-change" { return "LastOperationalStateChange" }
    if yname == "card-up-time" { return "CardUpTime" }
    return ""
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_FruInfo) GetSegmentPath() string {
    return "fru-info"
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_FruInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "last-operational-state-change" {
        return &fruInfo.LastOperationalStateChange
    }
    if childYangName == "card-up-time" {
        return &fruInfo.CardUpTime
    }
    return nil
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_FruInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["last-operational-state-change"] = &fruInfo.LastOperationalStateChange
    children["card-up-time"] = &fruInfo.CardUpTime
    return children
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_FruInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["card-administrative-state"] = fruInfo.CardAdministrativeState
    leafs["power-administrative-state"] = fruInfo.PowerAdministrativeState
    leafs["card-operational-state"] = fruInfo.CardOperationalState
    leafs["card-monitor-state"] = fruInfo.CardMonitorState
    leafs["card-reset-reason"] = fruInfo.CardResetReason
    leafs["power-current-measurement"] = fruInfo.PowerCurrentMeasurement
    leafs["power-operational-state"] = fruInfo.PowerOperationalState
    return leafs
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_FruInfo) GetBundleName() string { return "cisco_ios_xr" }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_FruInfo) GetYangName() string { return "fru-info" }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_FruInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_FruInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_FruInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_FruInfo) SetParent(parent types.Entity) { fruInfo.parent = parent }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_FruInfo) GetParent() types.Entity { return fruInfo.parent }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_FruInfo) GetParentYangName() string { return "basic-attributes" }

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_FruInfo_LastOperationalStateChange
// last card oper change state
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_FruInfo_LastOperationalStateChange struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_FruInfo_LastOperationalStateChange) GetFilter() yfilter.YFilter { return lastOperationalStateChange.YFilter }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_FruInfo_LastOperationalStateChange) SetFilter(yf yfilter.YFilter) { lastOperationalStateChange.YFilter = yf }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_FruInfo_LastOperationalStateChange) GetGoName(yname string) string {
    if yname == "time-in-seconds" { return "TimeInSeconds" }
    if yname == "time-in-nano-seconds" { return "TimeInNanoSeconds" }
    return ""
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_FruInfo_LastOperationalStateChange) GetSegmentPath() string {
    return "last-operational-state-change"
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_FruInfo_LastOperationalStateChange) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_FruInfo_LastOperationalStateChange) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_FruInfo_LastOperationalStateChange) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-in-seconds"] = lastOperationalStateChange.TimeInSeconds
    leafs["time-in-nano-seconds"] = lastOperationalStateChange.TimeInNanoSeconds
    return leafs
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_FruInfo_LastOperationalStateChange) GetBundleName() string { return "cisco_ios_xr" }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_FruInfo_LastOperationalStateChange) GetYangName() string { return "last-operational-state-change" }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_FruInfo_LastOperationalStateChange) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_FruInfo_LastOperationalStateChange) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_FruInfo_LastOperationalStateChange) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_FruInfo_LastOperationalStateChange) SetParent(parent types.Entity) { lastOperationalStateChange.parent = parent }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_FruInfo_LastOperationalStateChange) GetParent() types.Entity { return lastOperationalStateChange.parent }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_FruInfo_LastOperationalStateChange) GetParentYangName() string { return "fru-info" }

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_FruInfo_CardUpTime
// card up time
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_FruInfo_CardUpTime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_FruInfo_CardUpTime) GetFilter() yfilter.YFilter { return cardUpTime.YFilter }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_FruInfo_CardUpTime) SetFilter(yf yfilter.YFilter) { cardUpTime.YFilter = yf }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_FruInfo_CardUpTime) GetGoName(yname string) string {
    if yname == "time-in-seconds" { return "TimeInSeconds" }
    if yname == "time-in-nano-seconds" { return "TimeInNanoSeconds" }
    return ""
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_FruInfo_CardUpTime) GetSegmentPath() string {
    return "card-up-time"
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_FruInfo_CardUpTime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_FruInfo_CardUpTime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_FruInfo_CardUpTime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-in-seconds"] = cardUpTime.TimeInSeconds
    leafs["time-in-nano-seconds"] = cardUpTime.TimeInNanoSeconds
    return leafs
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_FruInfo_CardUpTime) GetBundleName() string { return "cisco_ios_xr" }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_FruInfo_CardUpTime) GetYangName() string { return "card-up-time" }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_FruInfo_CardUpTime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_FruInfo_CardUpTime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_FruInfo_CardUpTime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_FruInfo_CardUpTime) SetParent(parent types.Entity) { cardUpTime.parent = parent }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_FruInfo_CardUpTime) GetParent() types.Entity { return cardUpTime.parent }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_FruInfo_CardUpTime) GetParentYangName() string { return "fru-info" }

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_BasicInfo
// Inventory information
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_BasicInfo struct {
    parent types.Entity
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

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_BasicInfo) GetFilter() yfilter.YFilter { return basicInfo.YFilter }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_BasicInfo) SetFilter(yf yfilter.YFilter) { basicInfo.YFilter = yf }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_BasicInfo) GetGoName(yname string) string {
    if yname == "description" { return "Description" }
    if yname == "vendor-type" { return "VendorType" }
    if yname == "name" { return "Name" }
    if yname == "hardware-revision" { return "HardwareRevision" }
    if yname == "firmware-revision" { return "FirmwareRevision" }
    if yname == "software-revision" { return "SoftwareRevision" }
    if yname == "chip-hardware-revision" { return "ChipHardwareRevision" }
    if yname == "serial-number" { return "SerialNumber" }
    if yname == "manufacturer-name" { return "ManufacturerName" }
    if yname == "model-name" { return "ModelName" }
    if yname == "asset-id-str" { return "AssetIdStr" }
    if yname == "asset-identification" { return "AssetIdentification" }
    if yname == "is-field-replaceable-unit" { return "IsFieldReplaceableUnit" }
    if yname == "manufacturer-asset-tags" { return "ManufacturerAssetTags" }
    if yname == "composite-class-code" { return "CompositeClassCode" }
    if yname == "memory-size" { return "MemorySize" }
    if yname == "environmental-monitor-path" { return "EnvironmentalMonitorPath" }
    if yname == "alias" { return "Alias" }
    if yname == "group-flag" { return "GroupFlag" }
    if yname == "new-deviation-number" { return "NewDeviationNumber" }
    if yname == "physical-layer-interface-module-type" { return "PhysicalLayerInterfaceModuleType" }
    if yname == "unrecognized-fru" { return "UnrecognizedFru" }
    if yname == "redundancystate" { return "Redundancystate" }
    if yname == "ceport" { return "Ceport" }
    if yname == "xr-scoped" { return "XrScoped" }
    if yname == "unique-id" { return "UniqueId" }
    return ""
}

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_BasicInfo) GetSegmentPath() string {
    return "basic-info"
}

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_BasicInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_BasicInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_BasicInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["description"] = basicInfo.Description
    leafs["vendor-type"] = basicInfo.VendorType
    leafs["name"] = basicInfo.Name
    leafs["hardware-revision"] = basicInfo.HardwareRevision
    leafs["firmware-revision"] = basicInfo.FirmwareRevision
    leafs["software-revision"] = basicInfo.SoftwareRevision
    leafs["chip-hardware-revision"] = basicInfo.ChipHardwareRevision
    leafs["serial-number"] = basicInfo.SerialNumber
    leafs["manufacturer-name"] = basicInfo.ManufacturerName
    leafs["model-name"] = basicInfo.ModelName
    leafs["asset-id-str"] = basicInfo.AssetIdStr
    leafs["asset-identification"] = basicInfo.AssetIdentification
    leafs["is-field-replaceable-unit"] = basicInfo.IsFieldReplaceableUnit
    leafs["manufacturer-asset-tags"] = basicInfo.ManufacturerAssetTags
    leafs["composite-class-code"] = basicInfo.CompositeClassCode
    leafs["memory-size"] = basicInfo.MemorySize
    leafs["environmental-monitor-path"] = basicInfo.EnvironmentalMonitorPath
    leafs["alias"] = basicInfo.Alias
    leafs["group-flag"] = basicInfo.GroupFlag
    leafs["new-deviation-number"] = basicInfo.NewDeviationNumber
    leafs["physical-layer-interface-module-type"] = basicInfo.PhysicalLayerInterfaceModuleType
    leafs["unrecognized-fru"] = basicInfo.UnrecognizedFru
    leafs["redundancystate"] = basicInfo.Redundancystate
    leafs["ceport"] = basicInfo.Ceport
    leafs["xr-scoped"] = basicInfo.XrScoped
    leafs["unique-id"] = basicInfo.UniqueId
    return leafs
}

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_BasicInfo) GetBundleName() string { return "cisco_ios_xr" }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_BasicInfo) GetYangName() string { return "basic-info" }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_BasicInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_BasicInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_BasicInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_BasicInfo) SetParent(parent types.Entity) { basicInfo.parent = parent }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_BasicInfo) GetParent() types.Entity { return basicInfo.parent }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_BasicAttributes_BasicInfo) GetParentYangName() string { return "basic-attributes" }

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents
// HWComponent table contains all HW modules
// within the card 
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // HWComponent number. The type is slice of
    // Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent.
    HwComponent []Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent
}

func (hwComponents *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents) GetFilter() yfilter.YFilter { return hwComponents.YFilter }

func (hwComponents *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents) SetFilter(yf yfilter.YFilter) { hwComponents.YFilter = yf }

func (hwComponents *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents) GetGoName(yname string) string {
    if yname == "hw-component" { return "HwComponent" }
    return ""
}

func (hwComponents *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents) GetSegmentPath() string {
    return "hw-components"
}

func (hwComponents *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "hw-component" {
        for _, c := range hwComponents.HwComponent {
            if hwComponents.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent{}
        hwComponents.HwComponent = append(hwComponents.HwComponent, child)
        return &hwComponents.HwComponent[len(hwComponents.HwComponent)-1]
    }
    return nil
}

func (hwComponents *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range hwComponents.HwComponent {
        children[hwComponents.HwComponent[i].GetSegmentPath()] = &hwComponents.HwComponent[i]
    }
    return children
}

func (hwComponents *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (hwComponents *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents) GetBundleName() string { return "cisco_ios_xr" }

func (hwComponents *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents) GetYangName() string { return "hw-components" }

func (hwComponents *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hwComponents *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hwComponents *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hwComponents *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents) SetParent(parent types.Entity) { hwComponents.parent = parent }

func (hwComponents *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents) GetParent() types.Entity { return hwComponents.parent }

func (hwComponents *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents) GetParentYangName() string { return "card" }

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent
// HWComponent number
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. node number. The type is interface{} with range:
    // -2147483648..2147483647.
    Number interface{}

    // ModuleSensorTable contains all sensors in a Module.
    Sensors Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors

    // Attributes.
    BasicAttributes Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes
}

func (hwComponent *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent) GetFilter() yfilter.YFilter { return hwComponent.YFilter }

func (hwComponent *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent) SetFilter(yf yfilter.YFilter) { hwComponent.YFilter = yf }

func (hwComponent *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent) GetGoName(yname string) string {
    if yname == "number" { return "Number" }
    if yname == "sensors" { return "Sensors" }
    if yname == "basic-attributes" { return "BasicAttributes" }
    return ""
}

func (hwComponent *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent) GetSegmentPath() string {
    return "hw-component" + "[number='" + fmt.Sprintf("%v", hwComponent.Number) + "']"
}

func (hwComponent *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sensors" {
        return &hwComponent.Sensors
    }
    if childYangName == "basic-attributes" {
        return &hwComponent.BasicAttributes
    }
    return nil
}

func (hwComponent *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sensors"] = &hwComponent.Sensors
    children["basic-attributes"] = &hwComponent.BasicAttributes
    return children
}

func (hwComponent *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["number"] = hwComponent.Number
    return leafs
}

func (hwComponent *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent) GetBundleName() string { return "cisco_ios_xr" }

func (hwComponent *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent) GetYangName() string { return "hw-component" }

func (hwComponent *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hwComponent *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hwComponent *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hwComponent *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent) SetParent(parent types.Entity) { hwComponent.parent = parent }

func (hwComponent *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent) GetParent() types.Entity { return hwComponent.parent }

func (hwComponent *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent) GetParentYangName() string { return "hw-components" }

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors
// ModuleSensorTable contains all sensors in a
// Module.
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Sensor number in the Module. The type is slice of
    // Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor.
    Sensor []Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor
}

func (sensors *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors) GetFilter() yfilter.YFilter { return sensors.YFilter }

func (sensors *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors) SetFilter(yf yfilter.YFilter) { sensors.YFilter = yf }

func (sensors *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors) GetGoName(yname string) string {
    if yname == "sensor" { return "Sensor" }
    return ""
}

func (sensors *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors) GetSegmentPath() string {
    return "sensors"
}

func (sensors *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sensor" {
        for _, c := range sensors.Sensor {
            if sensors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor{}
        sensors.Sensor = append(sensors.Sensor, child)
        return &sensors.Sensor[len(sensors.Sensor)-1]
    }
    return nil
}

func (sensors *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sensors.Sensor {
        children[sensors.Sensor[i].GetSegmentPath()] = &sensors.Sensor[i]
    }
    return children
}

func (sensors *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sensors *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors) GetBundleName() string { return "cisco_ios_xr" }

func (sensors *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors) GetYangName() string { return "sensors" }

func (sensors *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sensors *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sensors *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sensors *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors) SetParent(parent types.Entity) { sensors.parent = parent }

func (sensors *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors) GetParent() types.Entity { return sensors.parent }

func (sensors *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors) GetParentYangName() string { return "hw-component" }

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor
// Sensor number in the Module
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. sensor number. The type is interface{} with range:
    // -2147483648..2147483647.
    Number interface{}

    // Attributes.
    BasicAttributes Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes
}

func (sensor *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor) GetFilter() yfilter.YFilter { return sensor.YFilter }

func (sensor *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor) SetFilter(yf yfilter.YFilter) { sensor.YFilter = yf }

func (sensor *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor) GetGoName(yname string) string {
    if yname == "number" { return "Number" }
    if yname == "basic-attributes" { return "BasicAttributes" }
    return ""
}

func (sensor *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor) GetSegmentPath() string {
    return "sensor" + "[number='" + fmt.Sprintf("%v", sensor.Number) + "']"
}

func (sensor *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "basic-attributes" {
        return &sensor.BasicAttributes
    }
    return nil
}

func (sensor *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["basic-attributes"] = &sensor.BasicAttributes
    return children
}

func (sensor *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["number"] = sensor.Number
    return leafs
}

func (sensor *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor) GetBundleName() string { return "cisco_ios_xr" }

func (sensor *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor) GetYangName() string { return "sensor" }

func (sensor *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sensor *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sensor *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sensor *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor) SetParent(parent types.Entity) { sensor.parent = parent }

func (sensor *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor) GetParent() types.Entity { return sensor.parent }

func (sensor *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor) GetParentYangName() string { return "sensors" }

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes
// Attributes
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Field Replaceable Unit (FRU) inventory information.
    FruInfo Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_FruInfo

    // Inventory information.
    BasicInfo Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_BasicInfo
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes) GetFilter() yfilter.YFilter { return basicAttributes.YFilter }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes) SetFilter(yf yfilter.YFilter) { basicAttributes.YFilter = yf }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes) GetGoName(yname string) string {
    if yname == "fru-info" { return "FruInfo" }
    if yname == "basic-info" { return "BasicInfo" }
    return ""
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes) GetSegmentPath() string {
    return "basic-attributes"
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "fru-info" {
        return &basicAttributes.FruInfo
    }
    if childYangName == "basic-info" {
        return &basicAttributes.BasicInfo
    }
    return nil
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["fru-info"] = &basicAttributes.FruInfo
    children["basic-info"] = &basicAttributes.BasicInfo
    return children
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes) GetBundleName() string { return "cisco_ios_xr" }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes) GetYangName() string { return "basic-attributes" }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes) SetParent(parent types.Entity) { basicAttributes.parent = parent }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes) GetParent() types.Entity { return basicAttributes.parent }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes) GetParentYangName() string { return "sensor" }

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_FruInfo
// Field Replaceable Unit (FRU) inventory
// information
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_FruInfo struct {
    parent types.Entity
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

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_FruInfo) GetFilter() yfilter.YFilter { return fruInfo.YFilter }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_FruInfo) SetFilter(yf yfilter.YFilter) { fruInfo.YFilter = yf }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_FruInfo) GetGoName(yname string) string {
    if yname == "card-administrative-state" { return "CardAdministrativeState" }
    if yname == "power-administrative-state" { return "PowerAdministrativeState" }
    if yname == "card-operational-state" { return "CardOperationalState" }
    if yname == "card-monitor-state" { return "CardMonitorState" }
    if yname == "card-reset-reason" { return "CardResetReason" }
    if yname == "power-current-measurement" { return "PowerCurrentMeasurement" }
    if yname == "power-operational-state" { return "PowerOperationalState" }
    if yname == "last-operational-state-change" { return "LastOperationalStateChange" }
    if yname == "card-up-time" { return "CardUpTime" }
    return ""
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_FruInfo) GetSegmentPath() string {
    return "fru-info"
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_FruInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "last-operational-state-change" {
        return &fruInfo.LastOperationalStateChange
    }
    if childYangName == "card-up-time" {
        return &fruInfo.CardUpTime
    }
    return nil
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_FruInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["last-operational-state-change"] = &fruInfo.LastOperationalStateChange
    children["card-up-time"] = &fruInfo.CardUpTime
    return children
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_FruInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["card-administrative-state"] = fruInfo.CardAdministrativeState
    leafs["power-administrative-state"] = fruInfo.PowerAdministrativeState
    leafs["card-operational-state"] = fruInfo.CardOperationalState
    leafs["card-monitor-state"] = fruInfo.CardMonitorState
    leafs["card-reset-reason"] = fruInfo.CardResetReason
    leafs["power-current-measurement"] = fruInfo.PowerCurrentMeasurement
    leafs["power-operational-state"] = fruInfo.PowerOperationalState
    return leafs
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_FruInfo) GetBundleName() string { return "cisco_ios_xr" }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_FruInfo) GetYangName() string { return "fru-info" }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_FruInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_FruInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_FruInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_FruInfo) SetParent(parent types.Entity) { fruInfo.parent = parent }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_FruInfo) GetParent() types.Entity { return fruInfo.parent }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_FruInfo) GetParentYangName() string { return "basic-attributes" }

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_FruInfo_LastOperationalStateChange
// last card oper change state
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_FruInfo_LastOperationalStateChange struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_FruInfo_LastOperationalStateChange) GetFilter() yfilter.YFilter { return lastOperationalStateChange.YFilter }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_FruInfo_LastOperationalStateChange) SetFilter(yf yfilter.YFilter) { lastOperationalStateChange.YFilter = yf }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_FruInfo_LastOperationalStateChange) GetGoName(yname string) string {
    if yname == "time-in-seconds" { return "TimeInSeconds" }
    if yname == "time-in-nano-seconds" { return "TimeInNanoSeconds" }
    return ""
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_FruInfo_LastOperationalStateChange) GetSegmentPath() string {
    return "last-operational-state-change"
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_FruInfo_LastOperationalStateChange) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_FruInfo_LastOperationalStateChange) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_FruInfo_LastOperationalStateChange) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-in-seconds"] = lastOperationalStateChange.TimeInSeconds
    leafs["time-in-nano-seconds"] = lastOperationalStateChange.TimeInNanoSeconds
    return leafs
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_FruInfo_LastOperationalStateChange) GetBundleName() string { return "cisco_ios_xr" }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_FruInfo_LastOperationalStateChange) GetYangName() string { return "last-operational-state-change" }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_FruInfo_LastOperationalStateChange) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_FruInfo_LastOperationalStateChange) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_FruInfo_LastOperationalStateChange) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_FruInfo_LastOperationalStateChange) SetParent(parent types.Entity) { lastOperationalStateChange.parent = parent }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_FruInfo_LastOperationalStateChange) GetParent() types.Entity { return lastOperationalStateChange.parent }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_FruInfo_LastOperationalStateChange) GetParentYangName() string { return "fru-info" }

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_FruInfo_CardUpTime
// card up time
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_FruInfo_CardUpTime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_FruInfo_CardUpTime) GetFilter() yfilter.YFilter { return cardUpTime.YFilter }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_FruInfo_CardUpTime) SetFilter(yf yfilter.YFilter) { cardUpTime.YFilter = yf }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_FruInfo_CardUpTime) GetGoName(yname string) string {
    if yname == "time-in-seconds" { return "TimeInSeconds" }
    if yname == "time-in-nano-seconds" { return "TimeInNanoSeconds" }
    return ""
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_FruInfo_CardUpTime) GetSegmentPath() string {
    return "card-up-time"
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_FruInfo_CardUpTime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_FruInfo_CardUpTime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_FruInfo_CardUpTime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-in-seconds"] = cardUpTime.TimeInSeconds
    leafs["time-in-nano-seconds"] = cardUpTime.TimeInNanoSeconds
    return leafs
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_FruInfo_CardUpTime) GetBundleName() string { return "cisco_ios_xr" }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_FruInfo_CardUpTime) GetYangName() string { return "card-up-time" }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_FruInfo_CardUpTime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_FruInfo_CardUpTime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_FruInfo_CardUpTime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_FruInfo_CardUpTime) SetParent(parent types.Entity) { cardUpTime.parent = parent }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_FruInfo_CardUpTime) GetParent() types.Entity { return cardUpTime.parent }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_FruInfo_CardUpTime) GetParentYangName() string { return "fru-info" }

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_BasicInfo
// Inventory information
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_BasicInfo struct {
    parent types.Entity
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

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_BasicInfo) GetFilter() yfilter.YFilter { return basicInfo.YFilter }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_BasicInfo) SetFilter(yf yfilter.YFilter) { basicInfo.YFilter = yf }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_BasicInfo) GetGoName(yname string) string {
    if yname == "description" { return "Description" }
    if yname == "vendor-type" { return "VendorType" }
    if yname == "name" { return "Name" }
    if yname == "hardware-revision" { return "HardwareRevision" }
    if yname == "firmware-revision" { return "FirmwareRevision" }
    if yname == "software-revision" { return "SoftwareRevision" }
    if yname == "chip-hardware-revision" { return "ChipHardwareRevision" }
    if yname == "serial-number" { return "SerialNumber" }
    if yname == "manufacturer-name" { return "ManufacturerName" }
    if yname == "model-name" { return "ModelName" }
    if yname == "asset-id-str" { return "AssetIdStr" }
    if yname == "asset-identification" { return "AssetIdentification" }
    if yname == "is-field-replaceable-unit" { return "IsFieldReplaceableUnit" }
    if yname == "manufacturer-asset-tags" { return "ManufacturerAssetTags" }
    if yname == "composite-class-code" { return "CompositeClassCode" }
    if yname == "memory-size" { return "MemorySize" }
    if yname == "environmental-monitor-path" { return "EnvironmentalMonitorPath" }
    if yname == "alias" { return "Alias" }
    if yname == "group-flag" { return "GroupFlag" }
    if yname == "new-deviation-number" { return "NewDeviationNumber" }
    if yname == "physical-layer-interface-module-type" { return "PhysicalLayerInterfaceModuleType" }
    if yname == "unrecognized-fru" { return "UnrecognizedFru" }
    if yname == "redundancystate" { return "Redundancystate" }
    if yname == "ceport" { return "Ceport" }
    if yname == "xr-scoped" { return "XrScoped" }
    if yname == "unique-id" { return "UniqueId" }
    return ""
}

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_BasicInfo) GetSegmentPath() string {
    return "basic-info"
}

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_BasicInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_BasicInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_BasicInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["description"] = basicInfo.Description
    leafs["vendor-type"] = basicInfo.VendorType
    leafs["name"] = basicInfo.Name
    leafs["hardware-revision"] = basicInfo.HardwareRevision
    leafs["firmware-revision"] = basicInfo.FirmwareRevision
    leafs["software-revision"] = basicInfo.SoftwareRevision
    leafs["chip-hardware-revision"] = basicInfo.ChipHardwareRevision
    leafs["serial-number"] = basicInfo.SerialNumber
    leafs["manufacturer-name"] = basicInfo.ManufacturerName
    leafs["model-name"] = basicInfo.ModelName
    leafs["asset-id-str"] = basicInfo.AssetIdStr
    leafs["asset-identification"] = basicInfo.AssetIdentification
    leafs["is-field-replaceable-unit"] = basicInfo.IsFieldReplaceableUnit
    leafs["manufacturer-asset-tags"] = basicInfo.ManufacturerAssetTags
    leafs["composite-class-code"] = basicInfo.CompositeClassCode
    leafs["memory-size"] = basicInfo.MemorySize
    leafs["environmental-monitor-path"] = basicInfo.EnvironmentalMonitorPath
    leafs["alias"] = basicInfo.Alias
    leafs["group-flag"] = basicInfo.GroupFlag
    leafs["new-deviation-number"] = basicInfo.NewDeviationNumber
    leafs["physical-layer-interface-module-type"] = basicInfo.PhysicalLayerInterfaceModuleType
    leafs["unrecognized-fru"] = basicInfo.UnrecognizedFru
    leafs["redundancystate"] = basicInfo.Redundancystate
    leafs["ceport"] = basicInfo.Ceport
    leafs["xr-scoped"] = basicInfo.XrScoped
    leafs["unique-id"] = basicInfo.UniqueId
    return leafs
}

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_BasicInfo) GetBundleName() string { return "cisco_ios_xr" }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_BasicInfo) GetYangName() string { return "basic-info" }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_BasicInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_BasicInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_BasicInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_BasicInfo) SetParent(parent types.Entity) { basicInfo.parent = parent }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_BasicInfo) GetParent() types.Entity { return basicInfo.parent }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_BasicAttributes_BasicInfo) GetParentYangName() string { return "basic-attributes" }

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes
// Attributes
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Field Replaceable Unit (FRU) inventory information.
    FruInfo Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_FruInfo

    // Inventory information.
    BasicInfo Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_BasicInfo
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes) GetFilter() yfilter.YFilter { return basicAttributes.YFilter }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes) SetFilter(yf yfilter.YFilter) { basicAttributes.YFilter = yf }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes) GetGoName(yname string) string {
    if yname == "fru-info" { return "FruInfo" }
    if yname == "basic-info" { return "BasicInfo" }
    return ""
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes) GetSegmentPath() string {
    return "basic-attributes"
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "fru-info" {
        return &basicAttributes.FruInfo
    }
    if childYangName == "basic-info" {
        return &basicAttributes.BasicInfo
    }
    return nil
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["fru-info"] = &basicAttributes.FruInfo
    children["basic-info"] = &basicAttributes.BasicInfo
    return children
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes) GetBundleName() string { return "cisco_ios_xr" }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes) GetYangName() string { return "basic-attributes" }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes) SetParent(parent types.Entity) { basicAttributes.parent = parent }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes) GetParent() types.Entity { return basicAttributes.parent }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes) GetParentYangName() string { return "hw-component" }

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_FruInfo
// Field Replaceable Unit (FRU) inventory
// information
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_FruInfo struct {
    parent types.Entity
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

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_FruInfo) GetFilter() yfilter.YFilter { return fruInfo.YFilter }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_FruInfo) SetFilter(yf yfilter.YFilter) { fruInfo.YFilter = yf }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_FruInfo) GetGoName(yname string) string {
    if yname == "card-administrative-state" { return "CardAdministrativeState" }
    if yname == "power-administrative-state" { return "PowerAdministrativeState" }
    if yname == "card-operational-state" { return "CardOperationalState" }
    if yname == "card-monitor-state" { return "CardMonitorState" }
    if yname == "card-reset-reason" { return "CardResetReason" }
    if yname == "power-current-measurement" { return "PowerCurrentMeasurement" }
    if yname == "power-operational-state" { return "PowerOperationalState" }
    if yname == "last-operational-state-change" { return "LastOperationalStateChange" }
    if yname == "card-up-time" { return "CardUpTime" }
    return ""
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_FruInfo) GetSegmentPath() string {
    return "fru-info"
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_FruInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "last-operational-state-change" {
        return &fruInfo.LastOperationalStateChange
    }
    if childYangName == "card-up-time" {
        return &fruInfo.CardUpTime
    }
    return nil
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_FruInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["last-operational-state-change"] = &fruInfo.LastOperationalStateChange
    children["card-up-time"] = &fruInfo.CardUpTime
    return children
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_FruInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["card-administrative-state"] = fruInfo.CardAdministrativeState
    leafs["power-administrative-state"] = fruInfo.PowerAdministrativeState
    leafs["card-operational-state"] = fruInfo.CardOperationalState
    leafs["card-monitor-state"] = fruInfo.CardMonitorState
    leafs["card-reset-reason"] = fruInfo.CardResetReason
    leafs["power-current-measurement"] = fruInfo.PowerCurrentMeasurement
    leafs["power-operational-state"] = fruInfo.PowerOperationalState
    return leafs
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_FruInfo) GetBundleName() string { return "cisco_ios_xr" }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_FruInfo) GetYangName() string { return "fru-info" }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_FruInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_FruInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_FruInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_FruInfo) SetParent(parent types.Entity) { fruInfo.parent = parent }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_FruInfo) GetParent() types.Entity { return fruInfo.parent }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_FruInfo) GetParentYangName() string { return "basic-attributes" }

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_FruInfo_LastOperationalStateChange
// last card oper change state
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_FruInfo_LastOperationalStateChange struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_FruInfo_LastOperationalStateChange) GetFilter() yfilter.YFilter { return lastOperationalStateChange.YFilter }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_FruInfo_LastOperationalStateChange) SetFilter(yf yfilter.YFilter) { lastOperationalStateChange.YFilter = yf }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_FruInfo_LastOperationalStateChange) GetGoName(yname string) string {
    if yname == "time-in-seconds" { return "TimeInSeconds" }
    if yname == "time-in-nano-seconds" { return "TimeInNanoSeconds" }
    return ""
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_FruInfo_LastOperationalStateChange) GetSegmentPath() string {
    return "last-operational-state-change"
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_FruInfo_LastOperationalStateChange) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_FruInfo_LastOperationalStateChange) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_FruInfo_LastOperationalStateChange) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-in-seconds"] = lastOperationalStateChange.TimeInSeconds
    leafs["time-in-nano-seconds"] = lastOperationalStateChange.TimeInNanoSeconds
    return leafs
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_FruInfo_LastOperationalStateChange) GetBundleName() string { return "cisco_ios_xr" }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_FruInfo_LastOperationalStateChange) GetYangName() string { return "last-operational-state-change" }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_FruInfo_LastOperationalStateChange) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_FruInfo_LastOperationalStateChange) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_FruInfo_LastOperationalStateChange) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_FruInfo_LastOperationalStateChange) SetParent(parent types.Entity) { lastOperationalStateChange.parent = parent }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_FruInfo_LastOperationalStateChange) GetParent() types.Entity { return lastOperationalStateChange.parent }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_FruInfo_LastOperationalStateChange) GetParentYangName() string { return "fru-info" }

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_FruInfo_CardUpTime
// card up time
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_FruInfo_CardUpTime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_FruInfo_CardUpTime) GetFilter() yfilter.YFilter { return cardUpTime.YFilter }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_FruInfo_CardUpTime) SetFilter(yf yfilter.YFilter) { cardUpTime.YFilter = yf }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_FruInfo_CardUpTime) GetGoName(yname string) string {
    if yname == "time-in-seconds" { return "TimeInSeconds" }
    if yname == "time-in-nano-seconds" { return "TimeInNanoSeconds" }
    return ""
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_FruInfo_CardUpTime) GetSegmentPath() string {
    return "card-up-time"
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_FruInfo_CardUpTime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_FruInfo_CardUpTime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_FruInfo_CardUpTime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-in-seconds"] = cardUpTime.TimeInSeconds
    leafs["time-in-nano-seconds"] = cardUpTime.TimeInNanoSeconds
    return leafs
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_FruInfo_CardUpTime) GetBundleName() string { return "cisco_ios_xr" }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_FruInfo_CardUpTime) GetYangName() string { return "card-up-time" }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_FruInfo_CardUpTime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_FruInfo_CardUpTime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_FruInfo_CardUpTime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_FruInfo_CardUpTime) SetParent(parent types.Entity) { cardUpTime.parent = parent }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_FruInfo_CardUpTime) GetParent() types.Entity { return cardUpTime.parent }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_FruInfo_CardUpTime) GetParentYangName() string { return "fru-info" }

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_BasicInfo
// Inventory information
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_BasicInfo struct {
    parent types.Entity
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

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_BasicInfo) GetFilter() yfilter.YFilter { return basicInfo.YFilter }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_BasicInfo) SetFilter(yf yfilter.YFilter) { basicInfo.YFilter = yf }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_BasicInfo) GetGoName(yname string) string {
    if yname == "description" { return "Description" }
    if yname == "vendor-type" { return "VendorType" }
    if yname == "name" { return "Name" }
    if yname == "hardware-revision" { return "HardwareRevision" }
    if yname == "firmware-revision" { return "FirmwareRevision" }
    if yname == "software-revision" { return "SoftwareRevision" }
    if yname == "chip-hardware-revision" { return "ChipHardwareRevision" }
    if yname == "serial-number" { return "SerialNumber" }
    if yname == "manufacturer-name" { return "ManufacturerName" }
    if yname == "model-name" { return "ModelName" }
    if yname == "asset-id-str" { return "AssetIdStr" }
    if yname == "asset-identification" { return "AssetIdentification" }
    if yname == "is-field-replaceable-unit" { return "IsFieldReplaceableUnit" }
    if yname == "manufacturer-asset-tags" { return "ManufacturerAssetTags" }
    if yname == "composite-class-code" { return "CompositeClassCode" }
    if yname == "memory-size" { return "MemorySize" }
    if yname == "environmental-monitor-path" { return "EnvironmentalMonitorPath" }
    if yname == "alias" { return "Alias" }
    if yname == "group-flag" { return "GroupFlag" }
    if yname == "new-deviation-number" { return "NewDeviationNumber" }
    if yname == "physical-layer-interface-module-type" { return "PhysicalLayerInterfaceModuleType" }
    if yname == "unrecognized-fru" { return "UnrecognizedFru" }
    if yname == "redundancystate" { return "Redundancystate" }
    if yname == "ceport" { return "Ceport" }
    if yname == "xr-scoped" { return "XrScoped" }
    if yname == "unique-id" { return "UniqueId" }
    return ""
}

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_BasicInfo) GetSegmentPath() string {
    return "basic-info"
}

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_BasicInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_BasicInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_BasicInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["description"] = basicInfo.Description
    leafs["vendor-type"] = basicInfo.VendorType
    leafs["name"] = basicInfo.Name
    leafs["hardware-revision"] = basicInfo.HardwareRevision
    leafs["firmware-revision"] = basicInfo.FirmwareRevision
    leafs["software-revision"] = basicInfo.SoftwareRevision
    leafs["chip-hardware-revision"] = basicInfo.ChipHardwareRevision
    leafs["serial-number"] = basicInfo.SerialNumber
    leafs["manufacturer-name"] = basicInfo.ManufacturerName
    leafs["model-name"] = basicInfo.ModelName
    leafs["asset-id-str"] = basicInfo.AssetIdStr
    leafs["asset-identification"] = basicInfo.AssetIdentification
    leafs["is-field-replaceable-unit"] = basicInfo.IsFieldReplaceableUnit
    leafs["manufacturer-asset-tags"] = basicInfo.ManufacturerAssetTags
    leafs["composite-class-code"] = basicInfo.CompositeClassCode
    leafs["memory-size"] = basicInfo.MemorySize
    leafs["environmental-monitor-path"] = basicInfo.EnvironmentalMonitorPath
    leafs["alias"] = basicInfo.Alias
    leafs["group-flag"] = basicInfo.GroupFlag
    leafs["new-deviation-number"] = basicInfo.NewDeviationNumber
    leafs["physical-layer-interface-module-type"] = basicInfo.PhysicalLayerInterfaceModuleType
    leafs["unrecognized-fru"] = basicInfo.UnrecognizedFru
    leafs["redundancystate"] = basicInfo.Redundancystate
    leafs["ceport"] = basicInfo.Ceport
    leafs["xr-scoped"] = basicInfo.XrScoped
    leafs["unique-id"] = basicInfo.UniqueId
    return leafs
}

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_BasicInfo) GetBundleName() string { return "cisco_ios_xr" }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_BasicInfo) GetYangName() string { return "basic-info" }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_BasicInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_BasicInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_BasicInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_BasicInfo) SetParent(parent types.Entity) { basicInfo.parent = parent }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_BasicInfo) GetParent() types.Entity { return basicInfo.parent }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_BasicAttributes_BasicInfo) GetParentYangName() string { return "basic-attributes" }

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots
// PortSlotTable contains all optics ports in a
// SPA/PLIM.
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PortSlot number in the SPA/PLIM. The type is slice of
    // Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot.
    PortSlot []Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot
}

func (portSlots *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots) GetFilter() yfilter.YFilter { return portSlots.YFilter }

func (portSlots *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots) SetFilter(yf yfilter.YFilter) { portSlots.YFilter = yf }

func (portSlots *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots) GetGoName(yname string) string {
    if yname == "port-slot" { return "PortSlot" }
    return ""
}

func (portSlots *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots) GetSegmentPath() string {
    return "port-slots"
}

func (portSlots *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "port-slot" {
        for _, c := range portSlots.PortSlot {
            if portSlots.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot{}
        portSlots.PortSlot = append(portSlots.PortSlot, child)
        return &portSlots.PortSlot[len(portSlots.PortSlot)-1]
    }
    return nil
}

func (portSlots *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range portSlots.PortSlot {
        children[portSlots.PortSlot[i].GetSegmentPath()] = &portSlots.PortSlot[i]
    }
    return children
}

func (portSlots *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (portSlots *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots) GetBundleName() string { return "cisco_ios_xr" }

func (portSlots *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots) GetYangName() string { return "port-slots" }

func (portSlots *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (portSlots *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (portSlots *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (portSlots *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots) SetParent(parent types.Entity) { portSlots.parent = parent }

func (portSlots *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots) GetParent() types.Entity { return portSlots.parent }

func (portSlots *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots) GetParentYangName() string { return "card" }

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot
// PortSlot number in the SPA/PLIM
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. portslot number. The type is interface{} with
    // range: -2147483648..2147483647.
    Number interface{}

    // Port string.
    Port Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port

    // Attributes.
    BasicAttributes Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes
}

func (portSlot *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot) GetFilter() yfilter.YFilter { return portSlot.YFilter }

func (portSlot *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot) SetFilter(yf yfilter.YFilter) { portSlot.YFilter = yf }

func (portSlot *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot) GetGoName(yname string) string {
    if yname == "number" { return "Number" }
    if yname == "port" { return "Port" }
    if yname == "basic-attributes" { return "BasicAttributes" }
    return ""
}

func (portSlot *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot) GetSegmentPath() string {
    return "port-slot" + "[number='" + fmt.Sprintf("%v", portSlot.Number) + "']"
}

func (portSlot *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "port" {
        return &portSlot.Port
    }
    if childYangName == "basic-attributes" {
        return &portSlot.BasicAttributes
    }
    return nil
}

func (portSlot *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["port"] = &portSlot.Port
    children["basic-attributes"] = &portSlot.BasicAttributes
    return children
}

func (portSlot *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["number"] = portSlot.Number
    return leafs
}

func (portSlot *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot) GetBundleName() string { return "cisco_ios_xr" }

func (portSlot *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot) GetYangName() string { return "port-slot" }

func (portSlot *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (portSlot *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (portSlot *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (portSlot *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot) SetParent(parent types.Entity) { portSlot.parent = parent }

func (portSlot *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot) GetParent() types.Entity { return portSlot.parent }

func (portSlot *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot) GetParentYangName() string { return "port-slots" }

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port
// Port string
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Attributes.
    BasicAttributes Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes
}

func (port *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port) GetFilter() yfilter.YFilter { return port.YFilter }

func (port *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port) SetFilter(yf yfilter.YFilter) { port.YFilter = yf }

func (port *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port) GetGoName(yname string) string {
    if yname == "basic-attributes" { return "BasicAttributes" }
    return ""
}

func (port *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port) GetSegmentPath() string {
    return "port"
}

func (port *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "basic-attributes" {
        return &port.BasicAttributes
    }
    return nil
}

func (port *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["basic-attributes"] = &port.BasicAttributes
    return children
}

func (port *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (port *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port) GetBundleName() string { return "cisco_ios_xr" }

func (port *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port) GetYangName() string { return "port" }

func (port *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (port *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (port *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (port *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port) SetParent(parent types.Entity) { port.parent = parent }

func (port *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port) GetParent() types.Entity { return port.parent }

func (port *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port) GetParentYangName() string { return "port-slot" }

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes
// Attributes
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Field Replaceable Unit (FRU) inventory information.
    FruInfo Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_FruInfo

    // Inventory information.
    BasicInfo Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_BasicInfo
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes) GetFilter() yfilter.YFilter { return basicAttributes.YFilter }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes) SetFilter(yf yfilter.YFilter) { basicAttributes.YFilter = yf }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes) GetGoName(yname string) string {
    if yname == "fru-info" { return "FruInfo" }
    if yname == "basic-info" { return "BasicInfo" }
    return ""
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes) GetSegmentPath() string {
    return "basic-attributes"
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "fru-info" {
        return &basicAttributes.FruInfo
    }
    if childYangName == "basic-info" {
        return &basicAttributes.BasicInfo
    }
    return nil
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["fru-info"] = &basicAttributes.FruInfo
    children["basic-info"] = &basicAttributes.BasicInfo
    return children
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes) GetBundleName() string { return "cisco_ios_xr" }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes) GetYangName() string { return "basic-attributes" }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes) SetParent(parent types.Entity) { basicAttributes.parent = parent }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes) GetParent() types.Entity { return basicAttributes.parent }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes) GetParentYangName() string { return "port" }

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_FruInfo
// Field Replaceable Unit (FRU) inventory
// information
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_FruInfo struct {
    parent types.Entity
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

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_FruInfo) GetFilter() yfilter.YFilter { return fruInfo.YFilter }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_FruInfo) SetFilter(yf yfilter.YFilter) { fruInfo.YFilter = yf }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_FruInfo) GetGoName(yname string) string {
    if yname == "card-administrative-state" { return "CardAdministrativeState" }
    if yname == "power-administrative-state" { return "PowerAdministrativeState" }
    if yname == "card-operational-state" { return "CardOperationalState" }
    if yname == "card-monitor-state" { return "CardMonitorState" }
    if yname == "card-reset-reason" { return "CardResetReason" }
    if yname == "power-current-measurement" { return "PowerCurrentMeasurement" }
    if yname == "power-operational-state" { return "PowerOperationalState" }
    if yname == "last-operational-state-change" { return "LastOperationalStateChange" }
    if yname == "card-up-time" { return "CardUpTime" }
    return ""
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_FruInfo) GetSegmentPath() string {
    return "fru-info"
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_FruInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "last-operational-state-change" {
        return &fruInfo.LastOperationalStateChange
    }
    if childYangName == "card-up-time" {
        return &fruInfo.CardUpTime
    }
    return nil
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_FruInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["last-operational-state-change"] = &fruInfo.LastOperationalStateChange
    children["card-up-time"] = &fruInfo.CardUpTime
    return children
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_FruInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["card-administrative-state"] = fruInfo.CardAdministrativeState
    leafs["power-administrative-state"] = fruInfo.PowerAdministrativeState
    leafs["card-operational-state"] = fruInfo.CardOperationalState
    leafs["card-monitor-state"] = fruInfo.CardMonitorState
    leafs["card-reset-reason"] = fruInfo.CardResetReason
    leafs["power-current-measurement"] = fruInfo.PowerCurrentMeasurement
    leafs["power-operational-state"] = fruInfo.PowerOperationalState
    return leafs
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_FruInfo) GetBundleName() string { return "cisco_ios_xr" }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_FruInfo) GetYangName() string { return "fru-info" }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_FruInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_FruInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_FruInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_FruInfo) SetParent(parent types.Entity) { fruInfo.parent = parent }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_FruInfo) GetParent() types.Entity { return fruInfo.parent }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_FruInfo) GetParentYangName() string { return "basic-attributes" }

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_LastOperationalStateChange
// last card oper change state
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_LastOperationalStateChange struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_LastOperationalStateChange) GetFilter() yfilter.YFilter { return lastOperationalStateChange.YFilter }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_LastOperationalStateChange) SetFilter(yf yfilter.YFilter) { lastOperationalStateChange.YFilter = yf }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_LastOperationalStateChange) GetGoName(yname string) string {
    if yname == "time-in-seconds" { return "TimeInSeconds" }
    if yname == "time-in-nano-seconds" { return "TimeInNanoSeconds" }
    return ""
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_LastOperationalStateChange) GetSegmentPath() string {
    return "last-operational-state-change"
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_LastOperationalStateChange) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_LastOperationalStateChange) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_LastOperationalStateChange) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-in-seconds"] = lastOperationalStateChange.TimeInSeconds
    leafs["time-in-nano-seconds"] = lastOperationalStateChange.TimeInNanoSeconds
    return leafs
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_LastOperationalStateChange) GetBundleName() string { return "cisco_ios_xr" }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_LastOperationalStateChange) GetYangName() string { return "last-operational-state-change" }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_LastOperationalStateChange) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_LastOperationalStateChange) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_LastOperationalStateChange) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_LastOperationalStateChange) SetParent(parent types.Entity) { lastOperationalStateChange.parent = parent }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_LastOperationalStateChange) GetParent() types.Entity { return lastOperationalStateChange.parent }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_LastOperationalStateChange) GetParentYangName() string { return "fru-info" }

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_CardUpTime
// card up time
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_CardUpTime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_CardUpTime) GetFilter() yfilter.YFilter { return cardUpTime.YFilter }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_CardUpTime) SetFilter(yf yfilter.YFilter) { cardUpTime.YFilter = yf }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_CardUpTime) GetGoName(yname string) string {
    if yname == "time-in-seconds" { return "TimeInSeconds" }
    if yname == "time-in-nano-seconds" { return "TimeInNanoSeconds" }
    return ""
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_CardUpTime) GetSegmentPath() string {
    return "card-up-time"
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_CardUpTime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_CardUpTime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_CardUpTime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-in-seconds"] = cardUpTime.TimeInSeconds
    leafs["time-in-nano-seconds"] = cardUpTime.TimeInNanoSeconds
    return leafs
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_CardUpTime) GetBundleName() string { return "cisco_ios_xr" }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_CardUpTime) GetYangName() string { return "card-up-time" }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_CardUpTime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_CardUpTime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_CardUpTime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_CardUpTime) SetParent(parent types.Entity) { cardUpTime.parent = parent }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_CardUpTime) GetParent() types.Entity { return cardUpTime.parent }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_FruInfo_CardUpTime) GetParentYangName() string { return "fru-info" }

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_BasicInfo
// Inventory information
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_BasicInfo struct {
    parent types.Entity
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

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_BasicInfo) GetFilter() yfilter.YFilter { return basicInfo.YFilter }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_BasicInfo) SetFilter(yf yfilter.YFilter) { basicInfo.YFilter = yf }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_BasicInfo) GetGoName(yname string) string {
    if yname == "description" { return "Description" }
    if yname == "vendor-type" { return "VendorType" }
    if yname == "name" { return "Name" }
    if yname == "hardware-revision" { return "HardwareRevision" }
    if yname == "firmware-revision" { return "FirmwareRevision" }
    if yname == "software-revision" { return "SoftwareRevision" }
    if yname == "chip-hardware-revision" { return "ChipHardwareRevision" }
    if yname == "serial-number" { return "SerialNumber" }
    if yname == "manufacturer-name" { return "ManufacturerName" }
    if yname == "model-name" { return "ModelName" }
    if yname == "asset-id-str" { return "AssetIdStr" }
    if yname == "asset-identification" { return "AssetIdentification" }
    if yname == "is-field-replaceable-unit" { return "IsFieldReplaceableUnit" }
    if yname == "manufacturer-asset-tags" { return "ManufacturerAssetTags" }
    if yname == "composite-class-code" { return "CompositeClassCode" }
    if yname == "memory-size" { return "MemorySize" }
    if yname == "environmental-monitor-path" { return "EnvironmentalMonitorPath" }
    if yname == "alias" { return "Alias" }
    if yname == "group-flag" { return "GroupFlag" }
    if yname == "new-deviation-number" { return "NewDeviationNumber" }
    if yname == "physical-layer-interface-module-type" { return "PhysicalLayerInterfaceModuleType" }
    if yname == "unrecognized-fru" { return "UnrecognizedFru" }
    if yname == "redundancystate" { return "Redundancystate" }
    if yname == "ceport" { return "Ceport" }
    if yname == "xr-scoped" { return "XrScoped" }
    if yname == "unique-id" { return "UniqueId" }
    return ""
}

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_BasicInfo) GetSegmentPath() string {
    return "basic-info"
}

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_BasicInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_BasicInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_BasicInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["description"] = basicInfo.Description
    leafs["vendor-type"] = basicInfo.VendorType
    leafs["name"] = basicInfo.Name
    leafs["hardware-revision"] = basicInfo.HardwareRevision
    leafs["firmware-revision"] = basicInfo.FirmwareRevision
    leafs["software-revision"] = basicInfo.SoftwareRevision
    leafs["chip-hardware-revision"] = basicInfo.ChipHardwareRevision
    leafs["serial-number"] = basicInfo.SerialNumber
    leafs["manufacturer-name"] = basicInfo.ManufacturerName
    leafs["model-name"] = basicInfo.ModelName
    leafs["asset-id-str"] = basicInfo.AssetIdStr
    leafs["asset-identification"] = basicInfo.AssetIdentification
    leafs["is-field-replaceable-unit"] = basicInfo.IsFieldReplaceableUnit
    leafs["manufacturer-asset-tags"] = basicInfo.ManufacturerAssetTags
    leafs["composite-class-code"] = basicInfo.CompositeClassCode
    leafs["memory-size"] = basicInfo.MemorySize
    leafs["environmental-monitor-path"] = basicInfo.EnvironmentalMonitorPath
    leafs["alias"] = basicInfo.Alias
    leafs["group-flag"] = basicInfo.GroupFlag
    leafs["new-deviation-number"] = basicInfo.NewDeviationNumber
    leafs["physical-layer-interface-module-type"] = basicInfo.PhysicalLayerInterfaceModuleType
    leafs["unrecognized-fru"] = basicInfo.UnrecognizedFru
    leafs["redundancystate"] = basicInfo.Redundancystate
    leafs["ceport"] = basicInfo.Ceport
    leafs["xr-scoped"] = basicInfo.XrScoped
    leafs["unique-id"] = basicInfo.UniqueId
    return leafs
}

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_BasicInfo) GetBundleName() string { return "cisco_ios_xr" }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_BasicInfo) GetYangName() string { return "basic-info" }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_BasicInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_BasicInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_BasicInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_BasicInfo) SetParent(parent types.Entity) { basicInfo.parent = parent }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_BasicInfo) GetParent() types.Entity { return basicInfo.parent }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Port_BasicAttributes_BasicInfo) GetParentYangName() string { return "basic-attributes" }

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes
// Attributes
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Field Replaceable Unit (FRU) inventory information.
    FruInfo Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_FruInfo

    // Inventory information.
    BasicInfo Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_BasicInfo
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes) GetFilter() yfilter.YFilter { return basicAttributes.YFilter }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes) SetFilter(yf yfilter.YFilter) { basicAttributes.YFilter = yf }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes) GetGoName(yname string) string {
    if yname == "fru-info" { return "FruInfo" }
    if yname == "basic-info" { return "BasicInfo" }
    return ""
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes) GetSegmentPath() string {
    return "basic-attributes"
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "fru-info" {
        return &basicAttributes.FruInfo
    }
    if childYangName == "basic-info" {
        return &basicAttributes.BasicInfo
    }
    return nil
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["fru-info"] = &basicAttributes.FruInfo
    children["basic-info"] = &basicAttributes.BasicInfo
    return children
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes) GetBundleName() string { return "cisco_ios_xr" }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes) GetYangName() string { return "basic-attributes" }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes) SetParent(parent types.Entity) { basicAttributes.parent = parent }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes) GetParent() types.Entity { return basicAttributes.parent }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes) GetParentYangName() string { return "port-slot" }

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_FruInfo
// Field Replaceable Unit (FRU) inventory
// information
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_FruInfo struct {
    parent types.Entity
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

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_FruInfo) GetFilter() yfilter.YFilter { return fruInfo.YFilter }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_FruInfo) SetFilter(yf yfilter.YFilter) { fruInfo.YFilter = yf }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_FruInfo) GetGoName(yname string) string {
    if yname == "card-administrative-state" { return "CardAdministrativeState" }
    if yname == "power-administrative-state" { return "PowerAdministrativeState" }
    if yname == "card-operational-state" { return "CardOperationalState" }
    if yname == "card-monitor-state" { return "CardMonitorState" }
    if yname == "card-reset-reason" { return "CardResetReason" }
    if yname == "power-current-measurement" { return "PowerCurrentMeasurement" }
    if yname == "power-operational-state" { return "PowerOperationalState" }
    if yname == "last-operational-state-change" { return "LastOperationalStateChange" }
    if yname == "card-up-time" { return "CardUpTime" }
    return ""
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_FruInfo) GetSegmentPath() string {
    return "fru-info"
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_FruInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "last-operational-state-change" {
        return &fruInfo.LastOperationalStateChange
    }
    if childYangName == "card-up-time" {
        return &fruInfo.CardUpTime
    }
    return nil
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_FruInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["last-operational-state-change"] = &fruInfo.LastOperationalStateChange
    children["card-up-time"] = &fruInfo.CardUpTime
    return children
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_FruInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["card-administrative-state"] = fruInfo.CardAdministrativeState
    leafs["power-administrative-state"] = fruInfo.PowerAdministrativeState
    leafs["card-operational-state"] = fruInfo.CardOperationalState
    leafs["card-monitor-state"] = fruInfo.CardMonitorState
    leafs["card-reset-reason"] = fruInfo.CardResetReason
    leafs["power-current-measurement"] = fruInfo.PowerCurrentMeasurement
    leafs["power-operational-state"] = fruInfo.PowerOperationalState
    return leafs
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_FruInfo) GetBundleName() string { return "cisco_ios_xr" }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_FruInfo) GetYangName() string { return "fru-info" }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_FruInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_FruInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_FruInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_FruInfo) SetParent(parent types.Entity) { fruInfo.parent = parent }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_FruInfo) GetParent() types.Entity { return fruInfo.parent }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_FruInfo) GetParentYangName() string { return "basic-attributes" }

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_FruInfo_LastOperationalStateChange
// last card oper change state
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_FruInfo_LastOperationalStateChange struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_FruInfo_LastOperationalStateChange) GetFilter() yfilter.YFilter { return lastOperationalStateChange.YFilter }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_FruInfo_LastOperationalStateChange) SetFilter(yf yfilter.YFilter) { lastOperationalStateChange.YFilter = yf }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_FruInfo_LastOperationalStateChange) GetGoName(yname string) string {
    if yname == "time-in-seconds" { return "TimeInSeconds" }
    if yname == "time-in-nano-seconds" { return "TimeInNanoSeconds" }
    return ""
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_FruInfo_LastOperationalStateChange) GetSegmentPath() string {
    return "last-operational-state-change"
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_FruInfo_LastOperationalStateChange) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_FruInfo_LastOperationalStateChange) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_FruInfo_LastOperationalStateChange) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-in-seconds"] = lastOperationalStateChange.TimeInSeconds
    leafs["time-in-nano-seconds"] = lastOperationalStateChange.TimeInNanoSeconds
    return leafs
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_FruInfo_LastOperationalStateChange) GetBundleName() string { return "cisco_ios_xr" }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_FruInfo_LastOperationalStateChange) GetYangName() string { return "last-operational-state-change" }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_FruInfo_LastOperationalStateChange) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_FruInfo_LastOperationalStateChange) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_FruInfo_LastOperationalStateChange) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_FruInfo_LastOperationalStateChange) SetParent(parent types.Entity) { lastOperationalStateChange.parent = parent }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_FruInfo_LastOperationalStateChange) GetParent() types.Entity { return lastOperationalStateChange.parent }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_FruInfo_LastOperationalStateChange) GetParentYangName() string { return "fru-info" }

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_FruInfo_CardUpTime
// card up time
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_FruInfo_CardUpTime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_FruInfo_CardUpTime) GetFilter() yfilter.YFilter { return cardUpTime.YFilter }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_FruInfo_CardUpTime) SetFilter(yf yfilter.YFilter) { cardUpTime.YFilter = yf }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_FruInfo_CardUpTime) GetGoName(yname string) string {
    if yname == "time-in-seconds" { return "TimeInSeconds" }
    if yname == "time-in-nano-seconds" { return "TimeInNanoSeconds" }
    return ""
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_FruInfo_CardUpTime) GetSegmentPath() string {
    return "card-up-time"
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_FruInfo_CardUpTime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_FruInfo_CardUpTime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_FruInfo_CardUpTime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-in-seconds"] = cardUpTime.TimeInSeconds
    leafs["time-in-nano-seconds"] = cardUpTime.TimeInNanoSeconds
    return leafs
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_FruInfo_CardUpTime) GetBundleName() string { return "cisco_ios_xr" }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_FruInfo_CardUpTime) GetYangName() string { return "card-up-time" }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_FruInfo_CardUpTime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_FruInfo_CardUpTime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_FruInfo_CardUpTime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_FruInfo_CardUpTime) SetParent(parent types.Entity) { cardUpTime.parent = parent }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_FruInfo_CardUpTime) GetParent() types.Entity { return cardUpTime.parent }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_FruInfo_CardUpTime) GetParentYangName() string { return "fru-info" }

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_BasicInfo
// Inventory information
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_BasicInfo struct {
    parent types.Entity
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

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_BasicInfo) GetFilter() yfilter.YFilter { return basicInfo.YFilter }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_BasicInfo) SetFilter(yf yfilter.YFilter) { basicInfo.YFilter = yf }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_BasicInfo) GetGoName(yname string) string {
    if yname == "description" { return "Description" }
    if yname == "vendor-type" { return "VendorType" }
    if yname == "name" { return "Name" }
    if yname == "hardware-revision" { return "HardwareRevision" }
    if yname == "firmware-revision" { return "FirmwareRevision" }
    if yname == "software-revision" { return "SoftwareRevision" }
    if yname == "chip-hardware-revision" { return "ChipHardwareRevision" }
    if yname == "serial-number" { return "SerialNumber" }
    if yname == "manufacturer-name" { return "ManufacturerName" }
    if yname == "model-name" { return "ModelName" }
    if yname == "asset-id-str" { return "AssetIdStr" }
    if yname == "asset-identification" { return "AssetIdentification" }
    if yname == "is-field-replaceable-unit" { return "IsFieldReplaceableUnit" }
    if yname == "manufacturer-asset-tags" { return "ManufacturerAssetTags" }
    if yname == "composite-class-code" { return "CompositeClassCode" }
    if yname == "memory-size" { return "MemorySize" }
    if yname == "environmental-monitor-path" { return "EnvironmentalMonitorPath" }
    if yname == "alias" { return "Alias" }
    if yname == "group-flag" { return "GroupFlag" }
    if yname == "new-deviation-number" { return "NewDeviationNumber" }
    if yname == "physical-layer-interface-module-type" { return "PhysicalLayerInterfaceModuleType" }
    if yname == "unrecognized-fru" { return "UnrecognizedFru" }
    if yname == "redundancystate" { return "Redundancystate" }
    if yname == "ceport" { return "Ceport" }
    if yname == "xr-scoped" { return "XrScoped" }
    if yname == "unique-id" { return "UniqueId" }
    return ""
}

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_BasicInfo) GetSegmentPath() string {
    return "basic-info"
}

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_BasicInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_BasicInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_BasicInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["description"] = basicInfo.Description
    leafs["vendor-type"] = basicInfo.VendorType
    leafs["name"] = basicInfo.Name
    leafs["hardware-revision"] = basicInfo.HardwareRevision
    leafs["firmware-revision"] = basicInfo.FirmwareRevision
    leafs["software-revision"] = basicInfo.SoftwareRevision
    leafs["chip-hardware-revision"] = basicInfo.ChipHardwareRevision
    leafs["serial-number"] = basicInfo.SerialNumber
    leafs["manufacturer-name"] = basicInfo.ManufacturerName
    leafs["model-name"] = basicInfo.ModelName
    leafs["asset-id-str"] = basicInfo.AssetIdStr
    leafs["asset-identification"] = basicInfo.AssetIdentification
    leafs["is-field-replaceable-unit"] = basicInfo.IsFieldReplaceableUnit
    leafs["manufacturer-asset-tags"] = basicInfo.ManufacturerAssetTags
    leafs["composite-class-code"] = basicInfo.CompositeClassCode
    leafs["memory-size"] = basicInfo.MemorySize
    leafs["environmental-monitor-path"] = basicInfo.EnvironmentalMonitorPath
    leafs["alias"] = basicInfo.Alias
    leafs["group-flag"] = basicInfo.GroupFlag
    leafs["new-deviation-number"] = basicInfo.NewDeviationNumber
    leafs["physical-layer-interface-module-type"] = basicInfo.PhysicalLayerInterfaceModuleType
    leafs["unrecognized-fru"] = basicInfo.UnrecognizedFru
    leafs["redundancystate"] = basicInfo.Redundancystate
    leafs["ceport"] = basicInfo.Ceport
    leafs["xr-scoped"] = basicInfo.XrScoped
    leafs["unique-id"] = basicInfo.UniqueId
    return leafs
}

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_BasicInfo) GetBundleName() string { return "cisco_ios_xr" }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_BasicInfo) GetYangName() string { return "basic-info" }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_BasicInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_BasicInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_BasicInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_BasicInfo) SetParent(parent types.Entity) { basicInfo.parent = parent }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_BasicInfo) GetParent() types.Entity { return basicInfo.parent }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_BasicAttributes_BasicInfo) GetParentYangName() string { return "basic-attributes" }

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors
// ModuleSensorTable contains all sensors in a
// Module.
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Sensor number in the Module. The type is slice of
    // Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor.
    Sensor []Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor
}

func (sensors *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors) GetFilter() yfilter.YFilter { return sensors.YFilter }

func (sensors *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors) SetFilter(yf yfilter.YFilter) { sensors.YFilter = yf }

func (sensors *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors) GetGoName(yname string) string {
    if yname == "sensor" { return "Sensor" }
    return ""
}

func (sensors *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors) GetSegmentPath() string {
    return "sensors"
}

func (sensors *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sensor" {
        for _, c := range sensors.Sensor {
            if sensors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor{}
        sensors.Sensor = append(sensors.Sensor, child)
        return &sensors.Sensor[len(sensors.Sensor)-1]
    }
    return nil
}

func (sensors *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sensors.Sensor {
        children[sensors.Sensor[i].GetSegmentPath()] = &sensors.Sensor[i]
    }
    return children
}

func (sensors *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sensors *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors) GetBundleName() string { return "cisco_ios_xr" }

func (sensors *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors) GetYangName() string { return "sensors" }

func (sensors *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sensors *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sensors *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sensors *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors) SetParent(parent types.Entity) { sensors.parent = parent }

func (sensors *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors) GetParent() types.Entity { return sensors.parent }

func (sensors *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors) GetParentYangName() string { return "card" }

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor
// Sensor number in the Module
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. sensor number. The type is interface{} with range:
    // -2147483648..2147483647.
    Number interface{}

    // Attributes.
    BasicAttributes Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes
}

func (sensor *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor) GetFilter() yfilter.YFilter { return sensor.YFilter }

func (sensor *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor) SetFilter(yf yfilter.YFilter) { sensor.YFilter = yf }

func (sensor *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor) GetGoName(yname string) string {
    if yname == "number" { return "Number" }
    if yname == "basic-attributes" { return "BasicAttributes" }
    return ""
}

func (sensor *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor) GetSegmentPath() string {
    return "sensor" + "[number='" + fmt.Sprintf("%v", sensor.Number) + "']"
}

func (sensor *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "basic-attributes" {
        return &sensor.BasicAttributes
    }
    return nil
}

func (sensor *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["basic-attributes"] = &sensor.BasicAttributes
    return children
}

func (sensor *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["number"] = sensor.Number
    return leafs
}

func (sensor *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor) GetBundleName() string { return "cisco_ios_xr" }

func (sensor *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor) GetYangName() string { return "sensor" }

func (sensor *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sensor *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sensor *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sensor *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor) SetParent(parent types.Entity) { sensor.parent = parent }

func (sensor *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor) GetParent() types.Entity { return sensor.parent }

func (sensor *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor) GetParentYangName() string { return "sensors" }

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes
// Attributes
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Field Replaceable Unit (FRU) inventory information.
    FruInfo Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_FruInfo

    // Inventory information.
    BasicInfo Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_BasicInfo
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes) GetFilter() yfilter.YFilter { return basicAttributes.YFilter }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes) SetFilter(yf yfilter.YFilter) { basicAttributes.YFilter = yf }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes) GetGoName(yname string) string {
    if yname == "fru-info" { return "FruInfo" }
    if yname == "basic-info" { return "BasicInfo" }
    return ""
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes) GetSegmentPath() string {
    return "basic-attributes"
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "fru-info" {
        return &basicAttributes.FruInfo
    }
    if childYangName == "basic-info" {
        return &basicAttributes.BasicInfo
    }
    return nil
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["fru-info"] = &basicAttributes.FruInfo
    children["basic-info"] = &basicAttributes.BasicInfo
    return children
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes) GetBundleName() string { return "cisco_ios_xr" }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes) GetYangName() string { return "basic-attributes" }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes) SetParent(parent types.Entity) { basicAttributes.parent = parent }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes) GetParent() types.Entity { return basicAttributes.parent }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes) GetParentYangName() string { return "sensor" }

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_FruInfo
// Field Replaceable Unit (FRU) inventory
// information
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_FruInfo struct {
    parent types.Entity
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

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_FruInfo) GetFilter() yfilter.YFilter { return fruInfo.YFilter }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_FruInfo) SetFilter(yf yfilter.YFilter) { fruInfo.YFilter = yf }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_FruInfo) GetGoName(yname string) string {
    if yname == "card-administrative-state" { return "CardAdministrativeState" }
    if yname == "power-administrative-state" { return "PowerAdministrativeState" }
    if yname == "card-operational-state" { return "CardOperationalState" }
    if yname == "card-monitor-state" { return "CardMonitorState" }
    if yname == "card-reset-reason" { return "CardResetReason" }
    if yname == "power-current-measurement" { return "PowerCurrentMeasurement" }
    if yname == "power-operational-state" { return "PowerOperationalState" }
    if yname == "last-operational-state-change" { return "LastOperationalStateChange" }
    if yname == "card-up-time" { return "CardUpTime" }
    return ""
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_FruInfo) GetSegmentPath() string {
    return "fru-info"
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_FruInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "last-operational-state-change" {
        return &fruInfo.LastOperationalStateChange
    }
    if childYangName == "card-up-time" {
        return &fruInfo.CardUpTime
    }
    return nil
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_FruInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["last-operational-state-change"] = &fruInfo.LastOperationalStateChange
    children["card-up-time"] = &fruInfo.CardUpTime
    return children
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_FruInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["card-administrative-state"] = fruInfo.CardAdministrativeState
    leafs["power-administrative-state"] = fruInfo.PowerAdministrativeState
    leafs["card-operational-state"] = fruInfo.CardOperationalState
    leafs["card-monitor-state"] = fruInfo.CardMonitorState
    leafs["card-reset-reason"] = fruInfo.CardResetReason
    leafs["power-current-measurement"] = fruInfo.PowerCurrentMeasurement
    leafs["power-operational-state"] = fruInfo.PowerOperationalState
    return leafs
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_FruInfo) GetBundleName() string { return "cisco_ios_xr" }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_FruInfo) GetYangName() string { return "fru-info" }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_FruInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_FruInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_FruInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_FruInfo) SetParent(parent types.Entity) { fruInfo.parent = parent }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_FruInfo) GetParent() types.Entity { return fruInfo.parent }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_FruInfo) GetParentYangName() string { return "basic-attributes" }

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_FruInfo_LastOperationalStateChange
// last card oper change state
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_FruInfo_LastOperationalStateChange struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_FruInfo_LastOperationalStateChange) GetFilter() yfilter.YFilter { return lastOperationalStateChange.YFilter }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_FruInfo_LastOperationalStateChange) SetFilter(yf yfilter.YFilter) { lastOperationalStateChange.YFilter = yf }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_FruInfo_LastOperationalStateChange) GetGoName(yname string) string {
    if yname == "time-in-seconds" { return "TimeInSeconds" }
    if yname == "time-in-nano-seconds" { return "TimeInNanoSeconds" }
    return ""
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_FruInfo_LastOperationalStateChange) GetSegmentPath() string {
    return "last-operational-state-change"
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_FruInfo_LastOperationalStateChange) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_FruInfo_LastOperationalStateChange) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_FruInfo_LastOperationalStateChange) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-in-seconds"] = lastOperationalStateChange.TimeInSeconds
    leafs["time-in-nano-seconds"] = lastOperationalStateChange.TimeInNanoSeconds
    return leafs
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_FruInfo_LastOperationalStateChange) GetBundleName() string { return "cisco_ios_xr" }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_FruInfo_LastOperationalStateChange) GetYangName() string { return "last-operational-state-change" }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_FruInfo_LastOperationalStateChange) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_FruInfo_LastOperationalStateChange) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_FruInfo_LastOperationalStateChange) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_FruInfo_LastOperationalStateChange) SetParent(parent types.Entity) { lastOperationalStateChange.parent = parent }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_FruInfo_LastOperationalStateChange) GetParent() types.Entity { return lastOperationalStateChange.parent }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_FruInfo_LastOperationalStateChange) GetParentYangName() string { return "fru-info" }

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_FruInfo_CardUpTime
// card up time
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_FruInfo_CardUpTime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_FruInfo_CardUpTime) GetFilter() yfilter.YFilter { return cardUpTime.YFilter }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_FruInfo_CardUpTime) SetFilter(yf yfilter.YFilter) { cardUpTime.YFilter = yf }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_FruInfo_CardUpTime) GetGoName(yname string) string {
    if yname == "time-in-seconds" { return "TimeInSeconds" }
    if yname == "time-in-nano-seconds" { return "TimeInNanoSeconds" }
    return ""
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_FruInfo_CardUpTime) GetSegmentPath() string {
    return "card-up-time"
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_FruInfo_CardUpTime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_FruInfo_CardUpTime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_FruInfo_CardUpTime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-in-seconds"] = cardUpTime.TimeInSeconds
    leafs["time-in-nano-seconds"] = cardUpTime.TimeInNanoSeconds
    return leafs
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_FruInfo_CardUpTime) GetBundleName() string { return "cisco_ios_xr" }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_FruInfo_CardUpTime) GetYangName() string { return "card-up-time" }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_FruInfo_CardUpTime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_FruInfo_CardUpTime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_FruInfo_CardUpTime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_FruInfo_CardUpTime) SetParent(parent types.Entity) { cardUpTime.parent = parent }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_FruInfo_CardUpTime) GetParent() types.Entity { return cardUpTime.parent }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_FruInfo_CardUpTime) GetParentYangName() string { return "fru-info" }

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_BasicInfo
// Inventory information
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_BasicInfo struct {
    parent types.Entity
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

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_BasicInfo) GetFilter() yfilter.YFilter { return basicInfo.YFilter }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_BasicInfo) SetFilter(yf yfilter.YFilter) { basicInfo.YFilter = yf }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_BasicInfo) GetGoName(yname string) string {
    if yname == "description" { return "Description" }
    if yname == "vendor-type" { return "VendorType" }
    if yname == "name" { return "Name" }
    if yname == "hardware-revision" { return "HardwareRevision" }
    if yname == "firmware-revision" { return "FirmwareRevision" }
    if yname == "software-revision" { return "SoftwareRevision" }
    if yname == "chip-hardware-revision" { return "ChipHardwareRevision" }
    if yname == "serial-number" { return "SerialNumber" }
    if yname == "manufacturer-name" { return "ManufacturerName" }
    if yname == "model-name" { return "ModelName" }
    if yname == "asset-id-str" { return "AssetIdStr" }
    if yname == "asset-identification" { return "AssetIdentification" }
    if yname == "is-field-replaceable-unit" { return "IsFieldReplaceableUnit" }
    if yname == "manufacturer-asset-tags" { return "ManufacturerAssetTags" }
    if yname == "composite-class-code" { return "CompositeClassCode" }
    if yname == "memory-size" { return "MemorySize" }
    if yname == "environmental-monitor-path" { return "EnvironmentalMonitorPath" }
    if yname == "alias" { return "Alias" }
    if yname == "group-flag" { return "GroupFlag" }
    if yname == "new-deviation-number" { return "NewDeviationNumber" }
    if yname == "physical-layer-interface-module-type" { return "PhysicalLayerInterfaceModuleType" }
    if yname == "unrecognized-fru" { return "UnrecognizedFru" }
    if yname == "redundancystate" { return "Redundancystate" }
    if yname == "ceport" { return "Ceport" }
    if yname == "xr-scoped" { return "XrScoped" }
    if yname == "unique-id" { return "UniqueId" }
    return ""
}

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_BasicInfo) GetSegmentPath() string {
    return "basic-info"
}

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_BasicInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_BasicInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_BasicInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["description"] = basicInfo.Description
    leafs["vendor-type"] = basicInfo.VendorType
    leafs["name"] = basicInfo.Name
    leafs["hardware-revision"] = basicInfo.HardwareRevision
    leafs["firmware-revision"] = basicInfo.FirmwareRevision
    leafs["software-revision"] = basicInfo.SoftwareRevision
    leafs["chip-hardware-revision"] = basicInfo.ChipHardwareRevision
    leafs["serial-number"] = basicInfo.SerialNumber
    leafs["manufacturer-name"] = basicInfo.ManufacturerName
    leafs["model-name"] = basicInfo.ModelName
    leafs["asset-id-str"] = basicInfo.AssetIdStr
    leafs["asset-identification"] = basicInfo.AssetIdentification
    leafs["is-field-replaceable-unit"] = basicInfo.IsFieldReplaceableUnit
    leafs["manufacturer-asset-tags"] = basicInfo.ManufacturerAssetTags
    leafs["composite-class-code"] = basicInfo.CompositeClassCode
    leafs["memory-size"] = basicInfo.MemorySize
    leafs["environmental-monitor-path"] = basicInfo.EnvironmentalMonitorPath
    leafs["alias"] = basicInfo.Alias
    leafs["group-flag"] = basicInfo.GroupFlag
    leafs["new-deviation-number"] = basicInfo.NewDeviationNumber
    leafs["physical-layer-interface-module-type"] = basicInfo.PhysicalLayerInterfaceModuleType
    leafs["unrecognized-fru"] = basicInfo.UnrecognizedFru
    leafs["redundancystate"] = basicInfo.Redundancystate
    leafs["ceport"] = basicInfo.Ceport
    leafs["xr-scoped"] = basicInfo.XrScoped
    leafs["unique-id"] = basicInfo.UniqueId
    return leafs
}

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_BasicInfo) GetBundleName() string { return "cisco_ios_xr" }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_BasicInfo) GetYangName() string { return "basic-info" }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_BasicInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_BasicInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_BasicInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_BasicInfo) SetParent(parent types.Entity) { basicInfo.parent = parent }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_BasicInfo) GetParent() types.Entity { return basicInfo.parent }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_BasicAttributes_BasicInfo) GetParentYangName() string { return "basic-attributes" }

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes
// Attributes
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Field Replaceable Unit (FRU) inventory information.
    FruInfo Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_FruInfo

    // Inventory information.
    BasicInfo Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_BasicInfo
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes) GetFilter() yfilter.YFilter { return basicAttributes.YFilter }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes) SetFilter(yf yfilter.YFilter) { basicAttributes.YFilter = yf }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes) GetGoName(yname string) string {
    if yname == "fru-info" { return "FruInfo" }
    if yname == "basic-info" { return "BasicInfo" }
    return ""
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes) GetSegmentPath() string {
    return "basic-attributes"
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "fru-info" {
        return &basicAttributes.FruInfo
    }
    if childYangName == "basic-info" {
        return &basicAttributes.BasicInfo
    }
    return nil
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["fru-info"] = &basicAttributes.FruInfo
    children["basic-info"] = &basicAttributes.BasicInfo
    return children
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes) GetBundleName() string { return "cisco_ios_xr" }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes) GetYangName() string { return "basic-attributes" }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes) SetParent(parent types.Entity) { basicAttributes.parent = parent }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes) GetParent() types.Entity { return basicAttributes.parent }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes) GetParentYangName() string { return "card" }

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_FruInfo
// Field Replaceable Unit (FRU) inventory
// information
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_FruInfo struct {
    parent types.Entity
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

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_FruInfo) GetFilter() yfilter.YFilter { return fruInfo.YFilter }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_FruInfo) SetFilter(yf yfilter.YFilter) { fruInfo.YFilter = yf }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_FruInfo) GetGoName(yname string) string {
    if yname == "card-administrative-state" { return "CardAdministrativeState" }
    if yname == "power-administrative-state" { return "PowerAdministrativeState" }
    if yname == "card-operational-state" { return "CardOperationalState" }
    if yname == "card-monitor-state" { return "CardMonitorState" }
    if yname == "card-reset-reason" { return "CardResetReason" }
    if yname == "power-current-measurement" { return "PowerCurrentMeasurement" }
    if yname == "power-operational-state" { return "PowerOperationalState" }
    if yname == "last-operational-state-change" { return "LastOperationalStateChange" }
    if yname == "card-up-time" { return "CardUpTime" }
    return ""
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_FruInfo) GetSegmentPath() string {
    return "fru-info"
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_FruInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "last-operational-state-change" {
        return &fruInfo.LastOperationalStateChange
    }
    if childYangName == "card-up-time" {
        return &fruInfo.CardUpTime
    }
    return nil
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_FruInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["last-operational-state-change"] = &fruInfo.LastOperationalStateChange
    children["card-up-time"] = &fruInfo.CardUpTime
    return children
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_FruInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["card-administrative-state"] = fruInfo.CardAdministrativeState
    leafs["power-administrative-state"] = fruInfo.PowerAdministrativeState
    leafs["card-operational-state"] = fruInfo.CardOperationalState
    leafs["card-monitor-state"] = fruInfo.CardMonitorState
    leafs["card-reset-reason"] = fruInfo.CardResetReason
    leafs["power-current-measurement"] = fruInfo.PowerCurrentMeasurement
    leafs["power-operational-state"] = fruInfo.PowerOperationalState
    return leafs
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_FruInfo) GetBundleName() string { return "cisco_ios_xr" }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_FruInfo) GetYangName() string { return "fru-info" }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_FruInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_FruInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_FruInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_FruInfo) SetParent(parent types.Entity) { fruInfo.parent = parent }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_FruInfo) GetParent() types.Entity { return fruInfo.parent }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_FruInfo) GetParentYangName() string { return "basic-attributes" }

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_FruInfo_LastOperationalStateChange
// last card oper change state
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_FruInfo_LastOperationalStateChange struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_FruInfo_LastOperationalStateChange) GetFilter() yfilter.YFilter { return lastOperationalStateChange.YFilter }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_FruInfo_LastOperationalStateChange) SetFilter(yf yfilter.YFilter) { lastOperationalStateChange.YFilter = yf }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_FruInfo_LastOperationalStateChange) GetGoName(yname string) string {
    if yname == "time-in-seconds" { return "TimeInSeconds" }
    if yname == "time-in-nano-seconds" { return "TimeInNanoSeconds" }
    return ""
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_FruInfo_LastOperationalStateChange) GetSegmentPath() string {
    return "last-operational-state-change"
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_FruInfo_LastOperationalStateChange) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_FruInfo_LastOperationalStateChange) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_FruInfo_LastOperationalStateChange) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-in-seconds"] = lastOperationalStateChange.TimeInSeconds
    leafs["time-in-nano-seconds"] = lastOperationalStateChange.TimeInNanoSeconds
    return leafs
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_FruInfo_LastOperationalStateChange) GetBundleName() string { return "cisco_ios_xr" }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_FruInfo_LastOperationalStateChange) GetYangName() string { return "last-operational-state-change" }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_FruInfo_LastOperationalStateChange) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_FruInfo_LastOperationalStateChange) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_FruInfo_LastOperationalStateChange) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_FruInfo_LastOperationalStateChange) SetParent(parent types.Entity) { lastOperationalStateChange.parent = parent }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_FruInfo_LastOperationalStateChange) GetParent() types.Entity { return lastOperationalStateChange.parent }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_FruInfo_LastOperationalStateChange) GetParentYangName() string { return "fru-info" }

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_FruInfo_CardUpTime
// card up time
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_FruInfo_CardUpTime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_FruInfo_CardUpTime) GetFilter() yfilter.YFilter { return cardUpTime.YFilter }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_FruInfo_CardUpTime) SetFilter(yf yfilter.YFilter) { cardUpTime.YFilter = yf }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_FruInfo_CardUpTime) GetGoName(yname string) string {
    if yname == "time-in-seconds" { return "TimeInSeconds" }
    if yname == "time-in-nano-seconds" { return "TimeInNanoSeconds" }
    return ""
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_FruInfo_CardUpTime) GetSegmentPath() string {
    return "card-up-time"
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_FruInfo_CardUpTime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_FruInfo_CardUpTime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_FruInfo_CardUpTime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-in-seconds"] = cardUpTime.TimeInSeconds
    leafs["time-in-nano-seconds"] = cardUpTime.TimeInNanoSeconds
    return leafs
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_FruInfo_CardUpTime) GetBundleName() string { return "cisco_ios_xr" }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_FruInfo_CardUpTime) GetYangName() string { return "card-up-time" }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_FruInfo_CardUpTime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_FruInfo_CardUpTime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_FruInfo_CardUpTime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_FruInfo_CardUpTime) SetParent(parent types.Entity) { cardUpTime.parent = parent }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_FruInfo_CardUpTime) GetParent() types.Entity { return cardUpTime.parent }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_FruInfo_CardUpTime) GetParentYangName() string { return "fru-info" }

// Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_BasicInfo
// Inventory information
type Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_BasicInfo struct {
    parent types.Entity
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

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_BasicInfo) GetFilter() yfilter.YFilter { return basicInfo.YFilter }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_BasicInfo) SetFilter(yf yfilter.YFilter) { basicInfo.YFilter = yf }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_BasicInfo) GetGoName(yname string) string {
    if yname == "description" { return "Description" }
    if yname == "vendor-type" { return "VendorType" }
    if yname == "name" { return "Name" }
    if yname == "hardware-revision" { return "HardwareRevision" }
    if yname == "firmware-revision" { return "FirmwareRevision" }
    if yname == "software-revision" { return "SoftwareRevision" }
    if yname == "chip-hardware-revision" { return "ChipHardwareRevision" }
    if yname == "serial-number" { return "SerialNumber" }
    if yname == "manufacturer-name" { return "ManufacturerName" }
    if yname == "model-name" { return "ModelName" }
    if yname == "asset-id-str" { return "AssetIdStr" }
    if yname == "asset-identification" { return "AssetIdentification" }
    if yname == "is-field-replaceable-unit" { return "IsFieldReplaceableUnit" }
    if yname == "manufacturer-asset-tags" { return "ManufacturerAssetTags" }
    if yname == "composite-class-code" { return "CompositeClassCode" }
    if yname == "memory-size" { return "MemorySize" }
    if yname == "environmental-monitor-path" { return "EnvironmentalMonitorPath" }
    if yname == "alias" { return "Alias" }
    if yname == "group-flag" { return "GroupFlag" }
    if yname == "new-deviation-number" { return "NewDeviationNumber" }
    if yname == "physical-layer-interface-module-type" { return "PhysicalLayerInterfaceModuleType" }
    if yname == "unrecognized-fru" { return "UnrecognizedFru" }
    if yname == "redundancystate" { return "Redundancystate" }
    if yname == "ceport" { return "Ceport" }
    if yname == "xr-scoped" { return "XrScoped" }
    if yname == "unique-id" { return "UniqueId" }
    return ""
}

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_BasicInfo) GetSegmentPath() string {
    return "basic-info"
}

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_BasicInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_BasicInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_BasicInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["description"] = basicInfo.Description
    leafs["vendor-type"] = basicInfo.VendorType
    leafs["name"] = basicInfo.Name
    leafs["hardware-revision"] = basicInfo.HardwareRevision
    leafs["firmware-revision"] = basicInfo.FirmwareRevision
    leafs["software-revision"] = basicInfo.SoftwareRevision
    leafs["chip-hardware-revision"] = basicInfo.ChipHardwareRevision
    leafs["serial-number"] = basicInfo.SerialNumber
    leafs["manufacturer-name"] = basicInfo.ManufacturerName
    leafs["model-name"] = basicInfo.ModelName
    leafs["asset-id-str"] = basicInfo.AssetIdStr
    leafs["asset-identification"] = basicInfo.AssetIdentification
    leafs["is-field-replaceable-unit"] = basicInfo.IsFieldReplaceableUnit
    leafs["manufacturer-asset-tags"] = basicInfo.ManufacturerAssetTags
    leafs["composite-class-code"] = basicInfo.CompositeClassCode
    leafs["memory-size"] = basicInfo.MemorySize
    leafs["environmental-monitor-path"] = basicInfo.EnvironmentalMonitorPath
    leafs["alias"] = basicInfo.Alias
    leafs["group-flag"] = basicInfo.GroupFlag
    leafs["new-deviation-number"] = basicInfo.NewDeviationNumber
    leafs["physical-layer-interface-module-type"] = basicInfo.PhysicalLayerInterfaceModuleType
    leafs["unrecognized-fru"] = basicInfo.UnrecognizedFru
    leafs["redundancystate"] = basicInfo.Redundancystate
    leafs["ceport"] = basicInfo.Ceport
    leafs["xr-scoped"] = basicInfo.XrScoped
    leafs["unique-id"] = basicInfo.UniqueId
    return leafs
}

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_BasicInfo) GetBundleName() string { return "cisco_ios_xr" }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_BasicInfo) GetYangName() string { return "basic-info" }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_BasicInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_BasicInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_BasicInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_BasicInfo) SetParent(parent types.Entity) { basicInfo.parent = parent }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_BasicInfo) GetParent() types.Entity { return basicInfo.parent }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_Cards_Card_BasicAttributes_BasicInfo) GetParentYangName() string { return "basic-attributes" }

// Inventory_Racks_Rack_Slots_Slot_BasicAttributes
// Attributes
type Inventory_Racks_Rack_Slots_Slot_BasicAttributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Field Replaceable Unit (FRU) inventory information.
    FruInfo Inventory_Racks_Rack_Slots_Slot_BasicAttributes_FruInfo

    // Inventory information.
    BasicInfo Inventory_Racks_Rack_Slots_Slot_BasicAttributes_BasicInfo
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_BasicAttributes) GetFilter() yfilter.YFilter { return basicAttributes.YFilter }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_BasicAttributes) SetFilter(yf yfilter.YFilter) { basicAttributes.YFilter = yf }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_BasicAttributes) GetGoName(yname string) string {
    if yname == "fru-info" { return "FruInfo" }
    if yname == "basic-info" { return "BasicInfo" }
    return ""
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_BasicAttributes) GetSegmentPath() string {
    return "basic-attributes"
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_BasicAttributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "fru-info" {
        return &basicAttributes.FruInfo
    }
    if childYangName == "basic-info" {
        return &basicAttributes.BasicInfo
    }
    return nil
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_BasicAttributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["fru-info"] = &basicAttributes.FruInfo
    children["basic-info"] = &basicAttributes.BasicInfo
    return children
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_BasicAttributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_BasicAttributes) GetBundleName() string { return "cisco_ios_xr" }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_BasicAttributes) GetYangName() string { return "basic-attributes" }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_BasicAttributes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_BasicAttributes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_BasicAttributes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_BasicAttributes) SetParent(parent types.Entity) { basicAttributes.parent = parent }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_BasicAttributes) GetParent() types.Entity { return basicAttributes.parent }

func (basicAttributes *Inventory_Racks_Rack_Slots_Slot_BasicAttributes) GetParentYangName() string { return "slot" }

// Inventory_Racks_Rack_Slots_Slot_BasicAttributes_FruInfo
// Field Replaceable Unit (FRU) inventory
// information
type Inventory_Racks_Rack_Slots_Slot_BasicAttributes_FruInfo struct {
    parent types.Entity
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

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_BasicAttributes_FruInfo) GetFilter() yfilter.YFilter { return fruInfo.YFilter }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_BasicAttributes_FruInfo) SetFilter(yf yfilter.YFilter) { fruInfo.YFilter = yf }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_BasicAttributes_FruInfo) GetGoName(yname string) string {
    if yname == "card-administrative-state" { return "CardAdministrativeState" }
    if yname == "power-administrative-state" { return "PowerAdministrativeState" }
    if yname == "card-operational-state" { return "CardOperationalState" }
    if yname == "card-monitor-state" { return "CardMonitorState" }
    if yname == "card-reset-reason" { return "CardResetReason" }
    if yname == "power-current-measurement" { return "PowerCurrentMeasurement" }
    if yname == "power-operational-state" { return "PowerOperationalState" }
    if yname == "last-operational-state-change" { return "LastOperationalStateChange" }
    if yname == "card-up-time" { return "CardUpTime" }
    return ""
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_BasicAttributes_FruInfo) GetSegmentPath() string {
    return "fru-info"
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_BasicAttributes_FruInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "last-operational-state-change" {
        return &fruInfo.LastOperationalStateChange
    }
    if childYangName == "card-up-time" {
        return &fruInfo.CardUpTime
    }
    return nil
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_BasicAttributes_FruInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["last-operational-state-change"] = &fruInfo.LastOperationalStateChange
    children["card-up-time"] = &fruInfo.CardUpTime
    return children
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_BasicAttributes_FruInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["card-administrative-state"] = fruInfo.CardAdministrativeState
    leafs["power-administrative-state"] = fruInfo.PowerAdministrativeState
    leafs["card-operational-state"] = fruInfo.CardOperationalState
    leafs["card-monitor-state"] = fruInfo.CardMonitorState
    leafs["card-reset-reason"] = fruInfo.CardResetReason
    leafs["power-current-measurement"] = fruInfo.PowerCurrentMeasurement
    leafs["power-operational-state"] = fruInfo.PowerOperationalState
    return leafs
}

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_BasicAttributes_FruInfo) GetBundleName() string { return "cisco_ios_xr" }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_BasicAttributes_FruInfo) GetYangName() string { return "fru-info" }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_BasicAttributes_FruInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_BasicAttributes_FruInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_BasicAttributes_FruInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_BasicAttributes_FruInfo) SetParent(parent types.Entity) { fruInfo.parent = parent }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_BasicAttributes_FruInfo) GetParent() types.Entity { return fruInfo.parent }

func (fruInfo *Inventory_Racks_Rack_Slots_Slot_BasicAttributes_FruInfo) GetParentYangName() string { return "basic-attributes" }

// Inventory_Racks_Rack_Slots_Slot_BasicAttributes_FruInfo_LastOperationalStateChange
// last card oper change state
type Inventory_Racks_Rack_Slots_Slot_BasicAttributes_FruInfo_LastOperationalStateChange struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_BasicAttributes_FruInfo_LastOperationalStateChange) GetFilter() yfilter.YFilter { return lastOperationalStateChange.YFilter }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_BasicAttributes_FruInfo_LastOperationalStateChange) SetFilter(yf yfilter.YFilter) { lastOperationalStateChange.YFilter = yf }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_BasicAttributes_FruInfo_LastOperationalStateChange) GetGoName(yname string) string {
    if yname == "time-in-seconds" { return "TimeInSeconds" }
    if yname == "time-in-nano-seconds" { return "TimeInNanoSeconds" }
    return ""
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_BasicAttributes_FruInfo_LastOperationalStateChange) GetSegmentPath() string {
    return "last-operational-state-change"
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_BasicAttributes_FruInfo_LastOperationalStateChange) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_BasicAttributes_FruInfo_LastOperationalStateChange) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_BasicAttributes_FruInfo_LastOperationalStateChange) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-in-seconds"] = lastOperationalStateChange.TimeInSeconds
    leafs["time-in-nano-seconds"] = lastOperationalStateChange.TimeInNanoSeconds
    return leafs
}

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_BasicAttributes_FruInfo_LastOperationalStateChange) GetBundleName() string { return "cisco_ios_xr" }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_BasicAttributes_FruInfo_LastOperationalStateChange) GetYangName() string { return "last-operational-state-change" }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_BasicAttributes_FruInfo_LastOperationalStateChange) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_BasicAttributes_FruInfo_LastOperationalStateChange) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_BasicAttributes_FruInfo_LastOperationalStateChange) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_BasicAttributes_FruInfo_LastOperationalStateChange) SetParent(parent types.Entity) { lastOperationalStateChange.parent = parent }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_BasicAttributes_FruInfo_LastOperationalStateChange) GetParent() types.Entity { return lastOperationalStateChange.parent }

func (lastOperationalStateChange *Inventory_Racks_Rack_Slots_Slot_BasicAttributes_FruInfo_LastOperationalStateChange) GetParentYangName() string { return "fru-info" }

// Inventory_Racks_Rack_Slots_Slot_BasicAttributes_FruInfo_CardUpTime
// card up time
type Inventory_Racks_Rack_Slots_Slot_BasicAttributes_FruInfo_CardUpTime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_BasicAttributes_FruInfo_CardUpTime) GetFilter() yfilter.YFilter { return cardUpTime.YFilter }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_BasicAttributes_FruInfo_CardUpTime) SetFilter(yf yfilter.YFilter) { cardUpTime.YFilter = yf }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_BasicAttributes_FruInfo_CardUpTime) GetGoName(yname string) string {
    if yname == "time-in-seconds" { return "TimeInSeconds" }
    if yname == "time-in-nano-seconds" { return "TimeInNanoSeconds" }
    return ""
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_BasicAttributes_FruInfo_CardUpTime) GetSegmentPath() string {
    return "card-up-time"
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_BasicAttributes_FruInfo_CardUpTime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_BasicAttributes_FruInfo_CardUpTime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_BasicAttributes_FruInfo_CardUpTime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-in-seconds"] = cardUpTime.TimeInSeconds
    leafs["time-in-nano-seconds"] = cardUpTime.TimeInNanoSeconds
    return leafs
}

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_BasicAttributes_FruInfo_CardUpTime) GetBundleName() string { return "cisco_ios_xr" }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_BasicAttributes_FruInfo_CardUpTime) GetYangName() string { return "card-up-time" }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_BasicAttributes_FruInfo_CardUpTime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_BasicAttributes_FruInfo_CardUpTime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_BasicAttributes_FruInfo_CardUpTime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_BasicAttributes_FruInfo_CardUpTime) SetParent(parent types.Entity) { cardUpTime.parent = parent }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_BasicAttributes_FruInfo_CardUpTime) GetParent() types.Entity { return cardUpTime.parent }

func (cardUpTime *Inventory_Racks_Rack_Slots_Slot_BasicAttributes_FruInfo_CardUpTime) GetParentYangName() string { return "fru-info" }

// Inventory_Racks_Rack_Slots_Slot_BasicAttributes_BasicInfo
// Inventory information
type Inventory_Racks_Rack_Slots_Slot_BasicAttributes_BasicInfo struct {
    parent types.Entity
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

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_BasicAttributes_BasicInfo) GetFilter() yfilter.YFilter { return basicInfo.YFilter }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_BasicAttributes_BasicInfo) SetFilter(yf yfilter.YFilter) { basicInfo.YFilter = yf }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_BasicAttributes_BasicInfo) GetGoName(yname string) string {
    if yname == "description" { return "Description" }
    if yname == "vendor-type" { return "VendorType" }
    if yname == "name" { return "Name" }
    if yname == "hardware-revision" { return "HardwareRevision" }
    if yname == "firmware-revision" { return "FirmwareRevision" }
    if yname == "software-revision" { return "SoftwareRevision" }
    if yname == "chip-hardware-revision" { return "ChipHardwareRevision" }
    if yname == "serial-number" { return "SerialNumber" }
    if yname == "manufacturer-name" { return "ManufacturerName" }
    if yname == "model-name" { return "ModelName" }
    if yname == "asset-id-str" { return "AssetIdStr" }
    if yname == "asset-identification" { return "AssetIdentification" }
    if yname == "is-field-replaceable-unit" { return "IsFieldReplaceableUnit" }
    if yname == "manufacturer-asset-tags" { return "ManufacturerAssetTags" }
    if yname == "composite-class-code" { return "CompositeClassCode" }
    if yname == "memory-size" { return "MemorySize" }
    if yname == "environmental-monitor-path" { return "EnvironmentalMonitorPath" }
    if yname == "alias" { return "Alias" }
    if yname == "group-flag" { return "GroupFlag" }
    if yname == "new-deviation-number" { return "NewDeviationNumber" }
    if yname == "physical-layer-interface-module-type" { return "PhysicalLayerInterfaceModuleType" }
    if yname == "unrecognized-fru" { return "UnrecognizedFru" }
    if yname == "redundancystate" { return "Redundancystate" }
    if yname == "ceport" { return "Ceport" }
    if yname == "xr-scoped" { return "XrScoped" }
    if yname == "unique-id" { return "UniqueId" }
    return ""
}

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_BasicAttributes_BasicInfo) GetSegmentPath() string {
    return "basic-info"
}

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_BasicAttributes_BasicInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_BasicAttributes_BasicInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_BasicAttributes_BasicInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["description"] = basicInfo.Description
    leafs["vendor-type"] = basicInfo.VendorType
    leafs["name"] = basicInfo.Name
    leafs["hardware-revision"] = basicInfo.HardwareRevision
    leafs["firmware-revision"] = basicInfo.FirmwareRevision
    leafs["software-revision"] = basicInfo.SoftwareRevision
    leafs["chip-hardware-revision"] = basicInfo.ChipHardwareRevision
    leafs["serial-number"] = basicInfo.SerialNumber
    leafs["manufacturer-name"] = basicInfo.ManufacturerName
    leafs["model-name"] = basicInfo.ModelName
    leafs["asset-id-str"] = basicInfo.AssetIdStr
    leafs["asset-identification"] = basicInfo.AssetIdentification
    leafs["is-field-replaceable-unit"] = basicInfo.IsFieldReplaceableUnit
    leafs["manufacturer-asset-tags"] = basicInfo.ManufacturerAssetTags
    leafs["composite-class-code"] = basicInfo.CompositeClassCode
    leafs["memory-size"] = basicInfo.MemorySize
    leafs["environmental-monitor-path"] = basicInfo.EnvironmentalMonitorPath
    leafs["alias"] = basicInfo.Alias
    leafs["group-flag"] = basicInfo.GroupFlag
    leafs["new-deviation-number"] = basicInfo.NewDeviationNumber
    leafs["physical-layer-interface-module-type"] = basicInfo.PhysicalLayerInterfaceModuleType
    leafs["unrecognized-fru"] = basicInfo.UnrecognizedFru
    leafs["redundancystate"] = basicInfo.Redundancystate
    leafs["ceport"] = basicInfo.Ceport
    leafs["xr-scoped"] = basicInfo.XrScoped
    leafs["unique-id"] = basicInfo.UniqueId
    return leafs
}

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_BasicAttributes_BasicInfo) GetBundleName() string { return "cisco_ios_xr" }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_BasicAttributes_BasicInfo) GetYangName() string { return "basic-info" }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_BasicAttributes_BasicInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_BasicAttributes_BasicInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_BasicAttributes_BasicInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_BasicAttributes_BasicInfo) SetParent(parent types.Entity) { basicInfo.parent = parent }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_BasicAttributes_BasicInfo) GetParent() types.Entity { return basicInfo.parent }

func (basicInfo *Inventory_Racks_Rack_Slots_Slot_BasicAttributes_BasicInfo) GetParentYangName() string { return "basic-attributes" }

// Inventory_Racks_Rack_FanTraies
// Table for rack fan trays
type Inventory_Racks_Rack_FanTraies struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Fan tray number. The type is slice of
    // Inventory_Racks_Rack_FanTraies_FanTray.
    FanTray []Inventory_Racks_Rack_FanTraies_FanTray
}

func (fanTraies *Inventory_Racks_Rack_FanTraies) GetFilter() yfilter.YFilter { return fanTraies.YFilter }

func (fanTraies *Inventory_Racks_Rack_FanTraies) SetFilter(yf yfilter.YFilter) { fanTraies.YFilter = yf }

func (fanTraies *Inventory_Racks_Rack_FanTraies) GetGoName(yname string) string {
    if yname == "fan-tray" { return "FanTray" }
    return ""
}

func (fanTraies *Inventory_Racks_Rack_FanTraies) GetSegmentPath() string {
    return "fan-traies"
}

func (fanTraies *Inventory_Racks_Rack_FanTraies) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "fan-tray" {
        for _, c := range fanTraies.FanTray {
            if fanTraies.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Inventory_Racks_Rack_FanTraies_FanTray{}
        fanTraies.FanTray = append(fanTraies.FanTray, child)
        return &fanTraies.FanTray[len(fanTraies.FanTray)-1]
    }
    return nil
}

func (fanTraies *Inventory_Racks_Rack_FanTraies) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range fanTraies.FanTray {
        children[fanTraies.FanTray[i].GetSegmentPath()] = &fanTraies.FanTray[i]
    }
    return children
}

func (fanTraies *Inventory_Racks_Rack_FanTraies) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (fanTraies *Inventory_Racks_Rack_FanTraies) GetBundleName() string { return "cisco_ios_xr" }

func (fanTraies *Inventory_Racks_Rack_FanTraies) GetYangName() string { return "fan-traies" }

func (fanTraies *Inventory_Racks_Rack_FanTraies) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fanTraies *Inventory_Racks_Rack_FanTraies) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fanTraies *Inventory_Racks_Rack_FanTraies) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fanTraies *Inventory_Racks_Rack_FanTraies) SetParent(parent types.Entity) { fanTraies.parent = parent }

func (fanTraies *Inventory_Racks_Rack_FanTraies) GetParent() types.Entity { return fanTraies.parent }

func (fanTraies *Inventory_Racks_Rack_FanTraies) GetParentYangName() string { return "rack" }

// Inventory_Racks_Rack_FanTraies_FanTray
// Fan tray number
type Inventory_Racks_Rack_FanTraies_FanTray struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Fan tray number. The type is interface{} with
    // range: -2147483648..2147483647.
    Number interface{}

    // Attributes.
    BasicAttributes Inventory_Racks_Rack_FanTraies_FanTray_BasicAttributes
}

func (fanTray *Inventory_Racks_Rack_FanTraies_FanTray) GetFilter() yfilter.YFilter { return fanTray.YFilter }

func (fanTray *Inventory_Racks_Rack_FanTraies_FanTray) SetFilter(yf yfilter.YFilter) { fanTray.YFilter = yf }

func (fanTray *Inventory_Racks_Rack_FanTraies_FanTray) GetGoName(yname string) string {
    if yname == "number" { return "Number" }
    if yname == "basic-attributes" { return "BasicAttributes" }
    return ""
}

func (fanTray *Inventory_Racks_Rack_FanTraies_FanTray) GetSegmentPath() string {
    return "fan-tray" + "[number='" + fmt.Sprintf("%v", fanTray.Number) + "']"
}

func (fanTray *Inventory_Racks_Rack_FanTraies_FanTray) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "basic-attributes" {
        return &fanTray.BasicAttributes
    }
    return nil
}

func (fanTray *Inventory_Racks_Rack_FanTraies_FanTray) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["basic-attributes"] = &fanTray.BasicAttributes
    return children
}

func (fanTray *Inventory_Racks_Rack_FanTraies_FanTray) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["number"] = fanTray.Number
    return leafs
}

func (fanTray *Inventory_Racks_Rack_FanTraies_FanTray) GetBundleName() string { return "cisco_ios_xr" }

func (fanTray *Inventory_Racks_Rack_FanTraies_FanTray) GetYangName() string { return "fan-tray" }

func (fanTray *Inventory_Racks_Rack_FanTraies_FanTray) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fanTray *Inventory_Racks_Rack_FanTraies_FanTray) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fanTray *Inventory_Racks_Rack_FanTraies_FanTray) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fanTray *Inventory_Racks_Rack_FanTraies_FanTray) SetParent(parent types.Entity) { fanTray.parent = parent }

func (fanTray *Inventory_Racks_Rack_FanTraies_FanTray) GetParent() types.Entity { return fanTray.parent }

func (fanTray *Inventory_Racks_Rack_FanTraies_FanTray) GetParentYangName() string { return "fan-traies" }

// Inventory_Racks_Rack_FanTraies_FanTray_BasicAttributes
// Attributes
type Inventory_Racks_Rack_FanTraies_FanTray_BasicAttributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Inventory information.
    BasicInfo Inventory_Racks_Rack_FanTraies_FanTray_BasicAttributes_BasicInfo
}

func (basicAttributes *Inventory_Racks_Rack_FanTraies_FanTray_BasicAttributes) GetFilter() yfilter.YFilter { return basicAttributes.YFilter }

func (basicAttributes *Inventory_Racks_Rack_FanTraies_FanTray_BasicAttributes) SetFilter(yf yfilter.YFilter) { basicAttributes.YFilter = yf }

func (basicAttributes *Inventory_Racks_Rack_FanTraies_FanTray_BasicAttributes) GetGoName(yname string) string {
    if yname == "basic-info" { return "BasicInfo" }
    return ""
}

func (basicAttributes *Inventory_Racks_Rack_FanTraies_FanTray_BasicAttributes) GetSegmentPath() string {
    return "basic-attributes"
}

func (basicAttributes *Inventory_Racks_Rack_FanTraies_FanTray_BasicAttributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "basic-info" {
        return &basicAttributes.BasicInfo
    }
    return nil
}

func (basicAttributes *Inventory_Racks_Rack_FanTraies_FanTray_BasicAttributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["basic-info"] = &basicAttributes.BasicInfo
    return children
}

func (basicAttributes *Inventory_Racks_Rack_FanTraies_FanTray_BasicAttributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (basicAttributes *Inventory_Racks_Rack_FanTraies_FanTray_BasicAttributes) GetBundleName() string { return "cisco_ios_xr" }

func (basicAttributes *Inventory_Racks_Rack_FanTraies_FanTray_BasicAttributes) GetYangName() string { return "basic-attributes" }

func (basicAttributes *Inventory_Racks_Rack_FanTraies_FanTray_BasicAttributes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (basicAttributes *Inventory_Racks_Rack_FanTraies_FanTray_BasicAttributes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (basicAttributes *Inventory_Racks_Rack_FanTraies_FanTray_BasicAttributes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (basicAttributes *Inventory_Racks_Rack_FanTraies_FanTray_BasicAttributes) SetParent(parent types.Entity) { basicAttributes.parent = parent }

func (basicAttributes *Inventory_Racks_Rack_FanTraies_FanTray_BasicAttributes) GetParent() types.Entity { return basicAttributes.parent }

func (basicAttributes *Inventory_Racks_Rack_FanTraies_FanTray_BasicAttributes) GetParentYangName() string { return "fan-tray" }

// Inventory_Racks_Rack_FanTraies_FanTray_BasicAttributes_BasicInfo
// Inventory information
type Inventory_Racks_Rack_FanTraies_FanTray_BasicAttributes_BasicInfo struct {
    parent types.Entity
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

func (basicInfo *Inventory_Racks_Rack_FanTraies_FanTray_BasicAttributes_BasicInfo) GetFilter() yfilter.YFilter { return basicInfo.YFilter }

func (basicInfo *Inventory_Racks_Rack_FanTraies_FanTray_BasicAttributes_BasicInfo) SetFilter(yf yfilter.YFilter) { basicInfo.YFilter = yf }

func (basicInfo *Inventory_Racks_Rack_FanTraies_FanTray_BasicAttributes_BasicInfo) GetGoName(yname string) string {
    if yname == "description" { return "Description" }
    if yname == "vendor-type" { return "VendorType" }
    if yname == "name" { return "Name" }
    if yname == "hardware-revision" { return "HardwareRevision" }
    if yname == "firmware-revision" { return "FirmwareRevision" }
    if yname == "software-revision" { return "SoftwareRevision" }
    if yname == "chip-hardware-revision" { return "ChipHardwareRevision" }
    if yname == "serial-number" { return "SerialNumber" }
    if yname == "manufacturer-name" { return "ManufacturerName" }
    if yname == "model-name" { return "ModelName" }
    if yname == "asset-id-str" { return "AssetIdStr" }
    if yname == "asset-identification" { return "AssetIdentification" }
    if yname == "is-field-replaceable-unit" { return "IsFieldReplaceableUnit" }
    if yname == "manufacturer-asset-tags" { return "ManufacturerAssetTags" }
    if yname == "composite-class-code" { return "CompositeClassCode" }
    if yname == "memory-size" { return "MemorySize" }
    if yname == "environmental-monitor-path" { return "EnvironmentalMonitorPath" }
    if yname == "alias" { return "Alias" }
    if yname == "group-flag" { return "GroupFlag" }
    if yname == "new-deviation-number" { return "NewDeviationNumber" }
    if yname == "physical-layer-interface-module-type" { return "PhysicalLayerInterfaceModuleType" }
    if yname == "unrecognized-fru" { return "UnrecognizedFru" }
    if yname == "redundancystate" { return "Redundancystate" }
    if yname == "ceport" { return "Ceport" }
    if yname == "xr-scoped" { return "XrScoped" }
    if yname == "unique-id" { return "UniqueId" }
    return ""
}

func (basicInfo *Inventory_Racks_Rack_FanTraies_FanTray_BasicAttributes_BasicInfo) GetSegmentPath() string {
    return "basic-info"
}

func (basicInfo *Inventory_Racks_Rack_FanTraies_FanTray_BasicAttributes_BasicInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (basicInfo *Inventory_Racks_Rack_FanTraies_FanTray_BasicAttributes_BasicInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (basicInfo *Inventory_Racks_Rack_FanTraies_FanTray_BasicAttributes_BasicInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["description"] = basicInfo.Description
    leafs["vendor-type"] = basicInfo.VendorType
    leafs["name"] = basicInfo.Name
    leafs["hardware-revision"] = basicInfo.HardwareRevision
    leafs["firmware-revision"] = basicInfo.FirmwareRevision
    leafs["software-revision"] = basicInfo.SoftwareRevision
    leafs["chip-hardware-revision"] = basicInfo.ChipHardwareRevision
    leafs["serial-number"] = basicInfo.SerialNumber
    leafs["manufacturer-name"] = basicInfo.ManufacturerName
    leafs["model-name"] = basicInfo.ModelName
    leafs["asset-id-str"] = basicInfo.AssetIdStr
    leafs["asset-identification"] = basicInfo.AssetIdentification
    leafs["is-field-replaceable-unit"] = basicInfo.IsFieldReplaceableUnit
    leafs["manufacturer-asset-tags"] = basicInfo.ManufacturerAssetTags
    leafs["composite-class-code"] = basicInfo.CompositeClassCode
    leafs["memory-size"] = basicInfo.MemorySize
    leafs["environmental-monitor-path"] = basicInfo.EnvironmentalMonitorPath
    leafs["alias"] = basicInfo.Alias
    leafs["group-flag"] = basicInfo.GroupFlag
    leafs["new-deviation-number"] = basicInfo.NewDeviationNumber
    leafs["physical-layer-interface-module-type"] = basicInfo.PhysicalLayerInterfaceModuleType
    leafs["unrecognized-fru"] = basicInfo.UnrecognizedFru
    leafs["redundancystate"] = basicInfo.Redundancystate
    leafs["ceport"] = basicInfo.Ceport
    leafs["xr-scoped"] = basicInfo.XrScoped
    leafs["unique-id"] = basicInfo.UniqueId
    return leafs
}

func (basicInfo *Inventory_Racks_Rack_FanTraies_FanTray_BasicAttributes_BasicInfo) GetBundleName() string { return "cisco_ios_xr" }

func (basicInfo *Inventory_Racks_Rack_FanTraies_FanTray_BasicAttributes_BasicInfo) GetYangName() string { return "basic-info" }

func (basicInfo *Inventory_Racks_Rack_FanTraies_FanTray_BasicAttributes_BasicInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (basicInfo *Inventory_Racks_Rack_FanTraies_FanTray_BasicAttributes_BasicInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (basicInfo *Inventory_Racks_Rack_FanTraies_FanTray_BasicAttributes_BasicInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (basicInfo *Inventory_Racks_Rack_FanTraies_FanTray_BasicAttributes_BasicInfo) SetParent(parent types.Entity) { basicInfo.parent = parent }

func (basicInfo *Inventory_Racks_Rack_FanTraies_FanTray_BasicAttributes_BasicInfo) GetParent() types.Entity { return basicInfo.parent }

func (basicInfo *Inventory_Racks_Rack_FanTraies_FanTray_BasicAttributes_BasicInfo) GetParentYangName() string { return "basic-attributes" }

// Inventory_Racks_Rack_PowerSupplyZones
// Table for defining rack power supply zones 
type Inventory_Racks_Rack_PowerSupplyZones struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Power Supply Zone number. The type is slice of
    // Inventory_Racks_Rack_PowerSupplyZones_PowerSupplyZone.
    PowerSupplyZone []Inventory_Racks_Rack_PowerSupplyZones_PowerSupplyZone
}

func (powerSupplyZones *Inventory_Racks_Rack_PowerSupplyZones) GetFilter() yfilter.YFilter { return powerSupplyZones.YFilter }

func (powerSupplyZones *Inventory_Racks_Rack_PowerSupplyZones) SetFilter(yf yfilter.YFilter) { powerSupplyZones.YFilter = yf }

func (powerSupplyZones *Inventory_Racks_Rack_PowerSupplyZones) GetGoName(yname string) string {
    if yname == "power-supply-zone" { return "PowerSupplyZone" }
    return ""
}

func (powerSupplyZones *Inventory_Racks_Rack_PowerSupplyZones) GetSegmentPath() string {
    return "power-supply-zones"
}

func (powerSupplyZones *Inventory_Racks_Rack_PowerSupplyZones) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "power-supply-zone" {
        for _, c := range powerSupplyZones.PowerSupplyZone {
            if powerSupplyZones.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Inventory_Racks_Rack_PowerSupplyZones_PowerSupplyZone{}
        powerSupplyZones.PowerSupplyZone = append(powerSupplyZones.PowerSupplyZone, child)
        return &powerSupplyZones.PowerSupplyZone[len(powerSupplyZones.PowerSupplyZone)-1]
    }
    return nil
}

func (powerSupplyZones *Inventory_Racks_Rack_PowerSupplyZones) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range powerSupplyZones.PowerSupplyZone {
        children[powerSupplyZones.PowerSupplyZone[i].GetSegmentPath()] = &powerSupplyZones.PowerSupplyZone[i]
    }
    return children
}

func (powerSupplyZones *Inventory_Racks_Rack_PowerSupplyZones) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (powerSupplyZones *Inventory_Racks_Rack_PowerSupplyZones) GetBundleName() string { return "cisco_ios_xr" }

func (powerSupplyZones *Inventory_Racks_Rack_PowerSupplyZones) GetYangName() string { return "power-supply-zones" }

func (powerSupplyZones *Inventory_Racks_Rack_PowerSupplyZones) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (powerSupplyZones *Inventory_Racks_Rack_PowerSupplyZones) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (powerSupplyZones *Inventory_Racks_Rack_PowerSupplyZones) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (powerSupplyZones *Inventory_Racks_Rack_PowerSupplyZones) SetParent(parent types.Entity) { powerSupplyZones.parent = parent }

func (powerSupplyZones *Inventory_Racks_Rack_PowerSupplyZones) GetParent() types.Entity { return powerSupplyZones.parent }

func (powerSupplyZones *Inventory_Racks_Rack_PowerSupplyZones) GetParentYangName() string { return "rack" }

// Inventory_Racks_Rack_PowerSupplyZones_PowerSupplyZone
// Power Supply Zone number
type Inventory_Racks_Rack_PowerSupplyZones_PowerSupplyZone struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Power Zone number. The type is interface{} with
    // range: -2147483648..2147483647.
    Number interface{}

    // Power suppy zone attributes .
    PowerSupplyZoneAttributes Inventory_Racks_Rack_PowerSupplyZones_PowerSupplyZone_PowerSupplyZoneAttributes
}

func (powerSupplyZone *Inventory_Racks_Rack_PowerSupplyZones_PowerSupplyZone) GetFilter() yfilter.YFilter { return powerSupplyZone.YFilter }

func (powerSupplyZone *Inventory_Racks_Rack_PowerSupplyZones_PowerSupplyZone) SetFilter(yf yfilter.YFilter) { powerSupplyZone.YFilter = yf }

func (powerSupplyZone *Inventory_Racks_Rack_PowerSupplyZones_PowerSupplyZone) GetGoName(yname string) string {
    if yname == "number" { return "Number" }
    if yname == "power-supply-zone-attributes" { return "PowerSupplyZoneAttributes" }
    return ""
}

func (powerSupplyZone *Inventory_Racks_Rack_PowerSupplyZones_PowerSupplyZone) GetSegmentPath() string {
    return "power-supply-zone" + "[number='" + fmt.Sprintf("%v", powerSupplyZone.Number) + "']"
}

func (powerSupplyZone *Inventory_Racks_Rack_PowerSupplyZones_PowerSupplyZone) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "power-supply-zone-attributes" {
        return &powerSupplyZone.PowerSupplyZoneAttributes
    }
    return nil
}

func (powerSupplyZone *Inventory_Racks_Rack_PowerSupplyZones_PowerSupplyZone) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["power-supply-zone-attributes"] = &powerSupplyZone.PowerSupplyZoneAttributes
    return children
}

func (powerSupplyZone *Inventory_Racks_Rack_PowerSupplyZones_PowerSupplyZone) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["number"] = powerSupplyZone.Number
    return leafs
}

func (powerSupplyZone *Inventory_Racks_Rack_PowerSupplyZones_PowerSupplyZone) GetBundleName() string { return "cisco_ios_xr" }

func (powerSupplyZone *Inventory_Racks_Rack_PowerSupplyZones_PowerSupplyZone) GetYangName() string { return "power-supply-zone" }

func (powerSupplyZone *Inventory_Racks_Rack_PowerSupplyZones_PowerSupplyZone) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (powerSupplyZone *Inventory_Racks_Rack_PowerSupplyZones_PowerSupplyZone) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (powerSupplyZone *Inventory_Racks_Rack_PowerSupplyZones_PowerSupplyZone) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (powerSupplyZone *Inventory_Racks_Rack_PowerSupplyZones_PowerSupplyZone) SetParent(parent types.Entity) { powerSupplyZone.parent = parent }

func (powerSupplyZone *Inventory_Racks_Rack_PowerSupplyZones_PowerSupplyZone) GetParent() types.Entity { return powerSupplyZone.parent }

func (powerSupplyZone *Inventory_Racks_Rack_PowerSupplyZones_PowerSupplyZone) GetParentYangName() string { return "power-supply-zones" }

// Inventory_Racks_Rack_PowerSupplyZones_PowerSupplyZone_PowerSupplyZoneAttributes
// Power suppy zone attributes 
type Inventory_Racks_Rack_PowerSupplyZones_PowerSupplyZone_PowerSupplyZoneAttributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Power supply group information.
    PowerSupplyGroupInfo Inventory_Racks_Rack_PowerSupplyZones_PowerSupplyZone_PowerSupplyZoneAttributes_PowerSupplyGroupInfo
}

func (powerSupplyZoneAttributes *Inventory_Racks_Rack_PowerSupplyZones_PowerSupplyZone_PowerSupplyZoneAttributes) GetFilter() yfilter.YFilter { return powerSupplyZoneAttributes.YFilter }

func (powerSupplyZoneAttributes *Inventory_Racks_Rack_PowerSupplyZones_PowerSupplyZone_PowerSupplyZoneAttributes) SetFilter(yf yfilter.YFilter) { powerSupplyZoneAttributes.YFilter = yf }

func (powerSupplyZoneAttributes *Inventory_Racks_Rack_PowerSupplyZones_PowerSupplyZone_PowerSupplyZoneAttributes) GetGoName(yname string) string {
    if yname == "power-supply-group-info" { return "PowerSupplyGroupInfo" }
    return ""
}

func (powerSupplyZoneAttributes *Inventory_Racks_Rack_PowerSupplyZones_PowerSupplyZone_PowerSupplyZoneAttributes) GetSegmentPath() string {
    return "power-supply-zone-attributes"
}

func (powerSupplyZoneAttributes *Inventory_Racks_Rack_PowerSupplyZones_PowerSupplyZone_PowerSupplyZoneAttributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "power-supply-group-info" {
        return &powerSupplyZoneAttributes.PowerSupplyGroupInfo
    }
    return nil
}

func (powerSupplyZoneAttributes *Inventory_Racks_Rack_PowerSupplyZones_PowerSupplyZone_PowerSupplyZoneAttributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["power-supply-group-info"] = &powerSupplyZoneAttributes.PowerSupplyGroupInfo
    return children
}

func (powerSupplyZoneAttributes *Inventory_Racks_Rack_PowerSupplyZones_PowerSupplyZone_PowerSupplyZoneAttributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (powerSupplyZoneAttributes *Inventory_Racks_Rack_PowerSupplyZones_PowerSupplyZone_PowerSupplyZoneAttributes) GetBundleName() string { return "cisco_ios_xr" }

func (powerSupplyZoneAttributes *Inventory_Racks_Rack_PowerSupplyZones_PowerSupplyZone_PowerSupplyZoneAttributes) GetYangName() string { return "power-supply-zone-attributes" }

func (powerSupplyZoneAttributes *Inventory_Racks_Rack_PowerSupplyZones_PowerSupplyZone_PowerSupplyZoneAttributes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (powerSupplyZoneAttributes *Inventory_Racks_Rack_PowerSupplyZones_PowerSupplyZone_PowerSupplyZoneAttributes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (powerSupplyZoneAttributes *Inventory_Racks_Rack_PowerSupplyZones_PowerSupplyZone_PowerSupplyZoneAttributes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (powerSupplyZoneAttributes *Inventory_Racks_Rack_PowerSupplyZones_PowerSupplyZone_PowerSupplyZoneAttributes) SetParent(parent types.Entity) { powerSupplyZoneAttributes.parent = parent }

func (powerSupplyZoneAttributes *Inventory_Racks_Rack_PowerSupplyZones_PowerSupplyZone_PowerSupplyZoneAttributes) GetParent() types.Entity { return powerSupplyZoneAttributes.parent }

func (powerSupplyZoneAttributes *Inventory_Racks_Rack_PowerSupplyZones_PowerSupplyZone_PowerSupplyZoneAttributes) GetParentYangName() string { return "power-supply-zone" }

// Inventory_Racks_Rack_PowerSupplyZones_PowerSupplyZone_PowerSupplyZoneAttributes_PowerSupplyGroupInfo
// Power supply group information
type Inventory_Racks_Rack_PowerSupplyZones_PowerSupplyZone_PowerSupplyZoneAttributes_PowerSupplyGroupInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // redundancy mode. The type is interface{} with range:
    // -2147483648..2147483647.
    PowerGroupRedundancyMode interface{}

    // power units. The type is string.
    PowerGroupPowerUnits interface{}

    // available current. The type is interface{} with range:
    // -2147483648..2147483647.
    PowerGroupAvailableCurrent interface{}

    // drawn current. The type is interface{} with range: -2147483648..2147483647.
    PowerGroupDrawnCurrent interface{}
}

func (powerSupplyGroupInfo *Inventory_Racks_Rack_PowerSupplyZones_PowerSupplyZone_PowerSupplyZoneAttributes_PowerSupplyGroupInfo) GetFilter() yfilter.YFilter { return powerSupplyGroupInfo.YFilter }

func (powerSupplyGroupInfo *Inventory_Racks_Rack_PowerSupplyZones_PowerSupplyZone_PowerSupplyZoneAttributes_PowerSupplyGroupInfo) SetFilter(yf yfilter.YFilter) { powerSupplyGroupInfo.YFilter = yf }

func (powerSupplyGroupInfo *Inventory_Racks_Rack_PowerSupplyZones_PowerSupplyZone_PowerSupplyZoneAttributes_PowerSupplyGroupInfo) GetGoName(yname string) string {
    if yname == "power-group-redundancy-mode" { return "PowerGroupRedundancyMode" }
    if yname == "power-group-power-units" { return "PowerGroupPowerUnits" }
    if yname == "power-group-available-current" { return "PowerGroupAvailableCurrent" }
    if yname == "power-group-drawn-current" { return "PowerGroupDrawnCurrent" }
    return ""
}

func (powerSupplyGroupInfo *Inventory_Racks_Rack_PowerSupplyZones_PowerSupplyZone_PowerSupplyZoneAttributes_PowerSupplyGroupInfo) GetSegmentPath() string {
    return "power-supply-group-info"
}

func (powerSupplyGroupInfo *Inventory_Racks_Rack_PowerSupplyZones_PowerSupplyZone_PowerSupplyZoneAttributes_PowerSupplyGroupInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (powerSupplyGroupInfo *Inventory_Racks_Rack_PowerSupplyZones_PowerSupplyZone_PowerSupplyZoneAttributes_PowerSupplyGroupInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (powerSupplyGroupInfo *Inventory_Racks_Rack_PowerSupplyZones_PowerSupplyZone_PowerSupplyZoneAttributes_PowerSupplyGroupInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["power-group-redundancy-mode"] = powerSupplyGroupInfo.PowerGroupRedundancyMode
    leafs["power-group-power-units"] = powerSupplyGroupInfo.PowerGroupPowerUnits
    leafs["power-group-available-current"] = powerSupplyGroupInfo.PowerGroupAvailableCurrent
    leafs["power-group-drawn-current"] = powerSupplyGroupInfo.PowerGroupDrawnCurrent
    return leafs
}

func (powerSupplyGroupInfo *Inventory_Racks_Rack_PowerSupplyZones_PowerSupplyZone_PowerSupplyZoneAttributes_PowerSupplyGroupInfo) GetBundleName() string { return "cisco_ios_xr" }

func (powerSupplyGroupInfo *Inventory_Racks_Rack_PowerSupplyZones_PowerSupplyZone_PowerSupplyZoneAttributes_PowerSupplyGroupInfo) GetYangName() string { return "power-supply-group-info" }

func (powerSupplyGroupInfo *Inventory_Racks_Rack_PowerSupplyZones_PowerSupplyZone_PowerSupplyZoneAttributes_PowerSupplyGroupInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (powerSupplyGroupInfo *Inventory_Racks_Rack_PowerSupplyZones_PowerSupplyZone_PowerSupplyZoneAttributes_PowerSupplyGroupInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (powerSupplyGroupInfo *Inventory_Racks_Rack_PowerSupplyZones_PowerSupplyZone_PowerSupplyZoneAttributes_PowerSupplyGroupInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (powerSupplyGroupInfo *Inventory_Racks_Rack_PowerSupplyZones_PowerSupplyZone_PowerSupplyZoneAttributes_PowerSupplyGroupInfo) SetParent(parent types.Entity) { powerSupplyGroupInfo.parent = parent }

func (powerSupplyGroupInfo *Inventory_Racks_Rack_PowerSupplyZones_PowerSupplyZone_PowerSupplyZoneAttributes_PowerSupplyGroupInfo) GetParent() types.Entity { return powerSupplyGroupInfo.parent }

func (powerSupplyGroupInfo *Inventory_Racks_Rack_PowerSupplyZones_PowerSupplyZone_PowerSupplyZoneAttributes_PowerSupplyGroupInfo) GetParentYangName() string { return "power-supply-zone-attributes" }

// Inventory_Racks_Rack_BasicAttributes
// Attributes
type Inventory_Racks_Rack_BasicAttributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Inventory information.
    BasicInfo Inventory_Racks_Rack_BasicAttributes_BasicInfo
}

func (basicAttributes *Inventory_Racks_Rack_BasicAttributes) GetFilter() yfilter.YFilter { return basicAttributes.YFilter }

func (basicAttributes *Inventory_Racks_Rack_BasicAttributes) SetFilter(yf yfilter.YFilter) { basicAttributes.YFilter = yf }

func (basicAttributes *Inventory_Racks_Rack_BasicAttributes) GetGoName(yname string) string {
    if yname == "basic-info" { return "BasicInfo" }
    return ""
}

func (basicAttributes *Inventory_Racks_Rack_BasicAttributes) GetSegmentPath() string {
    return "basic-attributes"
}

func (basicAttributes *Inventory_Racks_Rack_BasicAttributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "basic-info" {
        return &basicAttributes.BasicInfo
    }
    return nil
}

func (basicAttributes *Inventory_Racks_Rack_BasicAttributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["basic-info"] = &basicAttributes.BasicInfo
    return children
}

func (basicAttributes *Inventory_Racks_Rack_BasicAttributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (basicAttributes *Inventory_Racks_Rack_BasicAttributes) GetBundleName() string { return "cisco_ios_xr" }

func (basicAttributes *Inventory_Racks_Rack_BasicAttributes) GetYangName() string { return "basic-attributes" }

func (basicAttributes *Inventory_Racks_Rack_BasicAttributes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (basicAttributes *Inventory_Racks_Rack_BasicAttributes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (basicAttributes *Inventory_Racks_Rack_BasicAttributes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (basicAttributes *Inventory_Racks_Rack_BasicAttributes) SetParent(parent types.Entity) { basicAttributes.parent = parent }

func (basicAttributes *Inventory_Racks_Rack_BasicAttributes) GetParent() types.Entity { return basicAttributes.parent }

func (basicAttributes *Inventory_Racks_Rack_BasicAttributes) GetParentYangName() string { return "rack" }

// Inventory_Racks_Rack_BasicAttributes_BasicInfo
// Inventory information
type Inventory_Racks_Rack_BasicAttributes_BasicInfo struct {
    parent types.Entity
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

func (basicInfo *Inventory_Racks_Rack_BasicAttributes_BasicInfo) GetFilter() yfilter.YFilter { return basicInfo.YFilter }

func (basicInfo *Inventory_Racks_Rack_BasicAttributes_BasicInfo) SetFilter(yf yfilter.YFilter) { basicInfo.YFilter = yf }

func (basicInfo *Inventory_Racks_Rack_BasicAttributes_BasicInfo) GetGoName(yname string) string {
    if yname == "description" { return "Description" }
    if yname == "vendor-type" { return "VendorType" }
    if yname == "name" { return "Name" }
    if yname == "hardware-revision" { return "HardwareRevision" }
    if yname == "firmware-revision" { return "FirmwareRevision" }
    if yname == "software-revision" { return "SoftwareRevision" }
    if yname == "chip-hardware-revision" { return "ChipHardwareRevision" }
    if yname == "serial-number" { return "SerialNumber" }
    if yname == "manufacturer-name" { return "ManufacturerName" }
    if yname == "model-name" { return "ModelName" }
    if yname == "asset-id-str" { return "AssetIdStr" }
    if yname == "asset-identification" { return "AssetIdentification" }
    if yname == "is-field-replaceable-unit" { return "IsFieldReplaceableUnit" }
    if yname == "manufacturer-asset-tags" { return "ManufacturerAssetTags" }
    if yname == "composite-class-code" { return "CompositeClassCode" }
    if yname == "memory-size" { return "MemorySize" }
    if yname == "environmental-monitor-path" { return "EnvironmentalMonitorPath" }
    if yname == "alias" { return "Alias" }
    if yname == "group-flag" { return "GroupFlag" }
    if yname == "new-deviation-number" { return "NewDeviationNumber" }
    if yname == "physical-layer-interface-module-type" { return "PhysicalLayerInterfaceModuleType" }
    if yname == "unrecognized-fru" { return "UnrecognizedFru" }
    if yname == "redundancystate" { return "Redundancystate" }
    if yname == "ceport" { return "Ceport" }
    if yname == "xr-scoped" { return "XrScoped" }
    if yname == "unique-id" { return "UniqueId" }
    return ""
}

func (basicInfo *Inventory_Racks_Rack_BasicAttributes_BasicInfo) GetSegmentPath() string {
    return "basic-info"
}

func (basicInfo *Inventory_Racks_Rack_BasicAttributes_BasicInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (basicInfo *Inventory_Racks_Rack_BasicAttributes_BasicInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (basicInfo *Inventory_Racks_Rack_BasicAttributes_BasicInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["description"] = basicInfo.Description
    leafs["vendor-type"] = basicInfo.VendorType
    leafs["name"] = basicInfo.Name
    leafs["hardware-revision"] = basicInfo.HardwareRevision
    leafs["firmware-revision"] = basicInfo.FirmwareRevision
    leafs["software-revision"] = basicInfo.SoftwareRevision
    leafs["chip-hardware-revision"] = basicInfo.ChipHardwareRevision
    leafs["serial-number"] = basicInfo.SerialNumber
    leafs["manufacturer-name"] = basicInfo.ManufacturerName
    leafs["model-name"] = basicInfo.ModelName
    leafs["asset-id-str"] = basicInfo.AssetIdStr
    leafs["asset-identification"] = basicInfo.AssetIdentification
    leafs["is-field-replaceable-unit"] = basicInfo.IsFieldReplaceableUnit
    leafs["manufacturer-asset-tags"] = basicInfo.ManufacturerAssetTags
    leafs["composite-class-code"] = basicInfo.CompositeClassCode
    leafs["memory-size"] = basicInfo.MemorySize
    leafs["environmental-monitor-path"] = basicInfo.EnvironmentalMonitorPath
    leafs["alias"] = basicInfo.Alias
    leafs["group-flag"] = basicInfo.GroupFlag
    leafs["new-deviation-number"] = basicInfo.NewDeviationNumber
    leafs["physical-layer-interface-module-type"] = basicInfo.PhysicalLayerInterfaceModuleType
    leafs["unrecognized-fru"] = basicInfo.UnrecognizedFru
    leafs["redundancystate"] = basicInfo.Redundancystate
    leafs["ceport"] = basicInfo.Ceport
    leafs["xr-scoped"] = basicInfo.XrScoped
    leafs["unique-id"] = basicInfo.UniqueId
    return leafs
}

func (basicInfo *Inventory_Racks_Rack_BasicAttributes_BasicInfo) GetBundleName() string { return "cisco_ios_xr" }

func (basicInfo *Inventory_Racks_Rack_BasicAttributes_BasicInfo) GetYangName() string { return "basic-info" }

func (basicInfo *Inventory_Racks_Rack_BasicAttributes_BasicInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (basicInfo *Inventory_Racks_Rack_BasicAttributes_BasicInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (basicInfo *Inventory_Racks_Rack_BasicAttributes_BasicInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (basicInfo *Inventory_Racks_Rack_BasicAttributes_BasicInfo) SetParent(parent types.Entity) { basicInfo.parent = parent }

func (basicInfo *Inventory_Racks_Rack_BasicAttributes_BasicInfo) GetParent() types.Entity { return basicInfo.parent }

func (basicInfo *Inventory_Racks_Rack_BasicAttributes_BasicInfo) GetParentYangName() string { return "basic-attributes" }

