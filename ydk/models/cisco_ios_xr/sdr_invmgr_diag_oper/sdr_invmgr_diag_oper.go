// This module contains a collection of YANG definitions
// for Cisco IOS-XR sdr-invmgr-diag package operational data.
// 
// This module contains definitions
// for the following management objects:
//   diag: Diag information
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package sdr_invmgr_diag_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package sdr_invmgr_diag_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-sdr-invmgr-diag-oper diag}", reflect.TypeOf(Diag{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sdr-invmgr-diag-oper:diag", reflect.TypeOf(Diag{}))
}

// Diag
// Diag information
type Diag struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Table of racks.
    Racks Diag_Racks
}

func (diag *Diag) GetFilter() yfilter.YFilter { return diag.YFilter }

func (diag *Diag) SetFilter(yf yfilter.YFilter) { diag.YFilter = yf }

func (diag *Diag) GetGoName(yname string) string {
    if yname == "racks" { return "Racks" }
    return ""
}

func (diag *Diag) GetSegmentPath() string {
    return "Cisco-IOS-XR-sdr-invmgr-diag-oper:diag"
}

func (diag *Diag) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "racks" {
        return &diag.Racks
    }
    return nil
}

func (diag *Diag) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["racks"] = &diag.Racks
    return children
}

func (diag *Diag) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (diag *Diag) GetBundleName() string { return "cisco_ios_xr" }

func (diag *Diag) GetYangName() string { return "diag" }

func (diag *Diag) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (diag *Diag) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (diag *Diag) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (diag *Diag) SetParent(parent types.Entity) { diag.parent = parent }

func (diag *Diag) GetParent() types.Entity { return diag.parent }

func (diag *Diag) GetParentYangName() string { return "Cisco-IOS-XR-sdr-invmgr-diag-oper" }

// Diag_Racks
// Table of racks
type Diag_Racks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Rack name. The type is slice of Diag_Racks_Rack.
    Rack []Diag_Racks_Rack
}

func (racks *Diag_Racks) GetFilter() yfilter.YFilter { return racks.YFilter }

func (racks *Diag_Racks) SetFilter(yf yfilter.YFilter) { racks.YFilter = yf }

func (racks *Diag_Racks) GetGoName(yname string) string {
    if yname == "rack" { return "Rack" }
    return ""
}

func (racks *Diag_Racks) GetSegmentPath() string {
    return "racks"
}

func (racks *Diag_Racks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rack" {
        for _, c := range racks.Rack {
            if racks.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Diag_Racks_Rack{}
        racks.Rack = append(racks.Rack, child)
        return &racks.Rack[len(racks.Rack)-1]
    }
    return nil
}

func (racks *Diag_Racks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range racks.Rack {
        children[racks.Rack[i].GetSegmentPath()] = &racks.Rack[i]
    }
    return children
}

func (racks *Diag_Racks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (racks *Diag_Racks) GetBundleName() string { return "cisco_ios_xr" }

func (racks *Diag_Racks) GetYangName() string { return "racks" }

func (racks *Diag_Racks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (racks *Diag_Racks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (racks *Diag_Racks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (racks *Diag_Racks) SetParent(parent types.Entity) { racks.parent = parent }

func (racks *Diag_Racks) GetParent() types.Entity { return racks.parent }

func (racks *Diag_Racks) GetParentYangName() string { return "diag" }

// Diag_Racks_Rack
// Rack name
type Diag_Racks_Rack struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Rack name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    RackName interface{}

    // Table for rack power shelf .
    PowerShelfs Diag_Racks_Rack_PowerShelfs

    // Table for rack fan trays.
    FanTraies Diag_Racks_Rack_FanTraies

    // Table of slots.
    Slots Diag_Racks_Rack_Slots

    // Chassis information.
    Chassis Diag_Racks_Rack_Chassis
}

func (rack *Diag_Racks_Rack) GetFilter() yfilter.YFilter { return rack.YFilter }

func (rack *Diag_Racks_Rack) SetFilter(yf yfilter.YFilter) { rack.YFilter = yf }

func (rack *Diag_Racks_Rack) GetGoName(yname string) string {
    if yname == "rack-name" { return "RackName" }
    if yname == "power-shelfs" { return "PowerShelfs" }
    if yname == "fan-traies" { return "FanTraies" }
    if yname == "slots" { return "Slots" }
    if yname == "chassis" { return "Chassis" }
    return ""
}

func (rack *Diag_Racks_Rack) GetSegmentPath() string {
    return "rack" + "[rack-name='" + fmt.Sprintf("%v", rack.RackName) + "']"
}

func (rack *Diag_Racks_Rack) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "power-shelfs" {
        return &rack.PowerShelfs
    }
    if childYangName == "fan-traies" {
        return &rack.FanTraies
    }
    if childYangName == "slots" {
        return &rack.Slots
    }
    if childYangName == "chassis" {
        return &rack.Chassis
    }
    return nil
}

func (rack *Diag_Racks_Rack) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["power-shelfs"] = &rack.PowerShelfs
    children["fan-traies"] = &rack.FanTraies
    children["slots"] = &rack.Slots
    children["chassis"] = &rack.Chassis
    return children
}

func (rack *Diag_Racks_Rack) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rack-name"] = rack.RackName
    return leafs
}

func (rack *Diag_Racks_Rack) GetBundleName() string { return "cisco_ios_xr" }

func (rack *Diag_Racks_Rack) GetYangName() string { return "rack" }

func (rack *Diag_Racks_Rack) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rack *Diag_Racks_Rack) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rack *Diag_Racks_Rack) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rack *Diag_Racks_Rack) SetParent(parent types.Entity) { rack.parent = parent }

func (rack *Diag_Racks_Rack) GetParent() types.Entity { return rack.parent }

func (rack *Diag_Racks_Rack) GetParentYangName() string { return "racks" }

// Diag_Racks_Rack_PowerShelfs
// Table for rack power shelf 
type Diag_Racks_Rack_PowerShelfs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Power shelf name. The type is slice of
    // Diag_Racks_Rack_PowerShelfs_PowerShelf.
    PowerShelf []Diag_Racks_Rack_PowerShelfs_PowerShelf
}

func (powerShelfs *Diag_Racks_Rack_PowerShelfs) GetFilter() yfilter.YFilter { return powerShelfs.YFilter }

func (powerShelfs *Diag_Racks_Rack_PowerShelfs) SetFilter(yf yfilter.YFilter) { powerShelfs.YFilter = yf }

func (powerShelfs *Diag_Racks_Rack_PowerShelfs) GetGoName(yname string) string {
    if yname == "power-shelf" { return "PowerShelf" }
    return ""
}

func (powerShelfs *Diag_Racks_Rack_PowerShelfs) GetSegmentPath() string {
    return "power-shelfs"
}

func (powerShelfs *Diag_Racks_Rack_PowerShelfs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "power-shelf" {
        for _, c := range powerShelfs.PowerShelf {
            if powerShelfs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Diag_Racks_Rack_PowerShelfs_PowerShelf{}
        powerShelfs.PowerShelf = append(powerShelfs.PowerShelf, child)
        return &powerShelfs.PowerShelf[len(powerShelfs.PowerShelf)-1]
    }
    return nil
}

func (powerShelfs *Diag_Racks_Rack_PowerShelfs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range powerShelfs.PowerShelf {
        children[powerShelfs.PowerShelf[i].GetSegmentPath()] = &powerShelfs.PowerShelf[i]
    }
    return children
}

func (powerShelfs *Diag_Racks_Rack_PowerShelfs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (powerShelfs *Diag_Racks_Rack_PowerShelfs) GetBundleName() string { return "cisco_ios_xr" }

func (powerShelfs *Diag_Racks_Rack_PowerShelfs) GetYangName() string { return "power-shelfs" }

func (powerShelfs *Diag_Racks_Rack_PowerShelfs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (powerShelfs *Diag_Racks_Rack_PowerShelfs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (powerShelfs *Diag_Racks_Rack_PowerShelfs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (powerShelfs *Diag_Racks_Rack_PowerShelfs) SetParent(parent types.Entity) { powerShelfs.parent = parent }

func (powerShelfs *Diag_Racks_Rack_PowerShelfs) GetParent() types.Entity { return powerShelfs.parent }

func (powerShelfs *Diag_Racks_Rack_PowerShelfs) GetParentYangName() string { return "rack" }

// Diag_Racks_Rack_PowerShelfs_PowerShelf
// Power shelf name
type Diag_Racks_Rack_PowerShelfs_PowerShelf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Power Shelf name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    PowerShelfName interface{}

    // Table for rack power supply .
    PowerSupplies Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies
}

func (powerShelf *Diag_Racks_Rack_PowerShelfs_PowerShelf) GetFilter() yfilter.YFilter { return powerShelf.YFilter }

func (powerShelf *Diag_Racks_Rack_PowerShelfs_PowerShelf) SetFilter(yf yfilter.YFilter) { powerShelf.YFilter = yf }

func (powerShelf *Diag_Racks_Rack_PowerShelfs_PowerShelf) GetGoName(yname string) string {
    if yname == "power-shelf-name" { return "PowerShelfName" }
    if yname == "power-supplies" { return "PowerSupplies" }
    return ""
}

func (powerShelf *Diag_Racks_Rack_PowerShelfs_PowerShelf) GetSegmentPath() string {
    return "power-shelf" + "[power-shelf-name='" + fmt.Sprintf("%v", powerShelf.PowerShelfName) + "']"
}

func (powerShelf *Diag_Racks_Rack_PowerShelfs_PowerShelf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "power-supplies" {
        return &powerShelf.PowerSupplies
    }
    return nil
}

func (powerShelf *Diag_Racks_Rack_PowerShelfs_PowerShelf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["power-supplies"] = &powerShelf.PowerSupplies
    return children
}

func (powerShelf *Diag_Racks_Rack_PowerShelfs_PowerShelf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["power-shelf-name"] = powerShelf.PowerShelfName
    return leafs
}

func (powerShelf *Diag_Racks_Rack_PowerShelfs_PowerShelf) GetBundleName() string { return "cisco_ios_xr" }

func (powerShelf *Diag_Racks_Rack_PowerShelfs_PowerShelf) GetYangName() string { return "power-shelf" }

func (powerShelf *Diag_Racks_Rack_PowerShelfs_PowerShelf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (powerShelf *Diag_Racks_Rack_PowerShelfs_PowerShelf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (powerShelf *Diag_Racks_Rack_PowerShelfs_PowerShelf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (powerShelf *Diag_Racks_Rack_PowerShelfs_PowerShelf) SetParent(parent types.Entity) { powerShelf.parent = parent }

func (powerShelf *Diag_Racks_Rack_PowerShelfs_PowerShelf) GetParent() types.Entity { return powerShelf.parent }

func (powerShelf *Diag_Racks_Rack_PowerShelfs_PowerShelf) GetParentYangName() string { return "power-shelfs" }

// Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies
// Table for rack power supply 
type Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Power Supply name. The type is slice of
    // Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies_PowerSupply.
    PowerSupply []Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies_PowerSupply
}

func (powerSupplies *Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies) GetFilter() yfilter.YFilter { return powerSupplies.YFilter }

func (powerSupplies *Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies) SetFilter(yf yfilter.YFilter) { powerSupplies.YFilter = yf }

func (powerSupplies *Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies) GetGoName(yname string) string {
    if yname == "power-supply" { return "PowerSupply" }
    return ""
}

func (powerSupplies *Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies) GetSegmentPath() string {
    return "power-supplies"
}

func (powerSupplies *Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "power-supply" {
        for _, c := range powerSupplies.PowerSupply {
            if powerSupplies.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies_PowerSupply{}
        powerSupplies.PowerSupply = append(powerSupplies.PowerSupply, child)
        return &powerSupplies.PowerSupply[len(powerSupplies.PowerSupply)-1]
    }
    return nil
}

func (powerSupplies *Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range powerSupplies.PowerSupply {
        children[powerSupplies.PowerSupply[i].GetSegmentPath()] = &powerSupplies.PowerSupply[i]
    }
    return children
}

func (powerSupplies *Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (powerSupplies *Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies) GetBundleName() string { return "cisco_ios_xr" }

func (powerSupplies *Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies) GetYangName() string { return "power-supplies" }

func (powerSupplies *Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (powerSupplies *Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (powerSupplies *Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (powerSupplies *Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies) SetParent(parent types.Entity) { powerSupplies.parent = parent }

func (powerSupplies *Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies) GetParent() types.Entity { return powerSupplies.parent }

func (powerSupplies *Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies) GetParentYangName() string { return "power-shelf" }

// Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies_PowerSupply
// Power Supply name
type Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies_PowerSupply struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Power Supply name. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    PowerSupplyName interface{}

    // Basic information.
    Information Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies_PowerSupply_Information
}

func (powerSupply *Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies_PowerSupply) GetFilter() yfilter.YFilter { return powerSupply.YFilter }

func (powerSupply *Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies_PowerSupply) SetFilter(yf yfilter.YFilter) { powerSupply.YFilter = yf }

func (powerSupply *Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies_PowerSupply) GetGoName(yname string) string {
    if yname == "power-supply-name" { return "PowerSupplyName" }
    if yname == "information" { return "Information" }
    return ""
}

func (powerSupply *Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies_PowerSupply) GetSegmentPath() string {
    return "power-supply" + "[power-supply-name='" + fmt.Sprintf("%v", powerSupply.PowerSupplyName) + "']"
}

func (powerSupply *Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies_PowerSupply) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "information" {
        return &powerSupply.Information
    }
    return nil
}

func (powerSupply *Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies_PowerSupply) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["information"] = &powerSupply.Information
    return children
}

func (powerSupply *Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies_PowerSupply) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["power-supply-name"] = powerSupply.PowerSupplyName
    return leafs
}

func (powerSupply *Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies_PowerSupply) GetBundleName() string { return "cisco_ios_xr" }

func (powerSupply *Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies_PowerSupply) GetYangName() string { return "power-supply" }

func (powerSupply *Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies_PowerSupply) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (powerSupply *Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies_PowerSupply) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (powerSupply *Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies_PowerSupply) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (powerSupply *Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies_PowerSupply) SetParent(parent types.Entity) { powerSupply.parent = parent }

func (powerSupply *Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies_PowerSupply) GetParent() types.Entity { return powerSupply.parent }

func (powerSupply *Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies_PowerSupply) GetParentYangName() string { return "power-supplies" }

// Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies_PowerSupply_Information
// Basic information
type Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies_PowerSupply_Information struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A textual description of physical entity. The type is string with length:
    // 0..255.
    Description interface{}

    // IDPROM Format Revision. The type is string with length: 0..255.
    IdpromFormatRev interface{}

    // Controller family. The type is string with length: 0..255.
    ControllerFamily interface{}

    // Controller type. The type is string with length: 0..255.
    ControllerType interface{}

    // Version ID. The type is string with length: 0..255.
    Vid interface{}

    // Hardware Revision. The type is string with length: 0..255.
    Hwid interface{}

    // Product ID. The type is string with length: 0..255.
    Pid interface{}

    // UDI description. The type is string with length: 0..255.
    UdiDescription interface{}

    // UDI name. The type is string with length: 0..255.
    UdiName interface{}

    // Common Language Equipment Identifier (CLEI) code. The type is string with
    // length: 0..255.
    Clei interface{}

    // Equipment Catalog Item (ECI) number. The type is string with length:
    // 0..255.
    Eci interface{}

    // Top assembly part number. The type is string with length: 0..255.
    TopAssemPartNum interface{}

    // Top assembly revision number. The type is string with length: 0..255.
    TopAssemVid interface{}

    // PCA number. The type is string with length: 0..255.
    PcaNum interface{}

    // PCA revision ID. The type is string with length: 0..255.
    Pcavid interface{}

    // Chassis serial number. The type is string with length: 0..255.
    ChassisSid interface{}

    // Deviation Number # 1. The type is string with length: 0..255.
    DevNum1 interface{}

    // Deviation Number # 2. The type is string with length: 0..255.
    DevNum2 interface{}

    // Deviation Number # 3. The type is string with length: 0..255.
    DevNum3 interface{}

    // Deviation Number # 4. The type is string with length: 0..255.
    DevNum4 interface{}

    // Deviation Number # 5. The type is string with length: 0..255.
    DevNum5 interface{}

    // Deviation Number # 6. The type is string with length: 0..255.
    DevNum6 interface{}

    // Deviation Number # 7. The type is string with length: 0..255.
    DevNum7 interface{}

    // Manufacturing Test Data. The type is string with length: 0..255.
    ManuTestData interface{}

    // Asset ID. The type is string with length: 0..255.
    AssetId interface{}

    // Asset Alias. The type is string with length: 0..255.
    AssetAlias interface{}

    // Base Mac Address #1. The type is string with length: 0..255.
    BaseMacAddress1 interface{}

    // Mac Address Block Size #1. The type is string with length: 0..255.
    MacAddBlkSize1 interface{}

    // Base Mac Address #2. The type is string with length: 0..255.
    BaseMacAddress2 interface{}

    // Mac Address Block Size #2. The type is string with length: 0..255.
    MacAddBlkSize2 interface{}

    // Base Mac Address #3. The type is string with length: 0..255.
    BaseMacAddress3 interface{}

    // Mac Address Block Size #3. The type is string with length: 0..255.
    MacAddBlkSize3 interface{}

    // Base Mac Address #4. The type is string with length: 0..255.
    BaseMacAddress4 interface{}

    // Mac Address Block Size #4. The type is string with length: 0..255.
    MacAddBlkSize4 interface{}

    // PCB Serial Number. The type is string with length: 0..255.
    PcbSerialNum interface{}

    // Power Supply Type. The type is string with length: 0..255.
    PowerSupplyType interface{}

    // Power Consumption. The type is string with length: 0..255.
    PowerConsumption interface{}

    // Block Signature. The type is string with length: 0..255.
    BlockSignature interface{}

    // Block Version. The type is string with length: 0..255.
    BlockVersion interface{}

    // Block Length. The type is string with length: 0..255.
    BlockLength interface{}

    // Block Checksum. The type is string with length: 0..255.
    BlockChecksum interface{}

    // EEPROM Size. The type is string with length: 0..255.
    EepromSize interface{}

    // Block Count. The type is string with length: 0..255.
    BlockCount interface{}

    // FRU Major Type. The type is string with length: 0..255.
    FruMajorType interface{}

    // FRU Minor Type. The type is string with length: 0..255.
    FruMinorType interface{}

    // OEM String. The type is string with length: 0..255.
    OemString interface{}

    // Product ID. The type is string with length: 0..255.
    ProductId interface{}

    // Serial Number. The type is string with length: 0..255.
    SerialNumber interface{}

    // Part Number. The type is string with length: 0..255.
    PartNumber interface{}

    // Part Revision. The type is string with length: 0..255.
    PartRevision interface{}

    // MFG Deviation. The type is string with length: 0..255.
    MfgDeviation interface{}

    // Hardware Version. The type is string with length: 0..255.
    HwVersion interface{}

    // MFG Bits. The type is string with length: 0..255.
    MfgBits interface{}

    // Engineer Use. The type is string with length: 0..255.
    EngineerUse interface{}

    // SNMP OID. The type is string with length: 0..255.
    Snmpoid interface{}

    // RMA Code. The type is string with length: 0..255.
    RmaCode interface{}

    // RMA Data.
    Rma Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies_PowerSupply_Information_Rma
}

func (information *Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies_PowerSupply_Information) GetFilter() yfilter.YFilter { return information.YFilter }

func (information *Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies_PowerSupply_Information) SetFilter(yf yfilter.YFilter) { information.YFilter = yf }

func (information *Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies_PowerSupply_Information) GetGoName(yname string) string {
    if yname == "description" { return "Description" }
    if yname == "idprom-format-rev" { return "IdpromFormatRev" }
    if yname == "controller-family" { return "ControllerFamily" }
    if yname == "controller-type" { return "ControllerType" }
    if yname == "vid" { return "Vid" }
    if yname == "hwid" { return "Hwid" }
    if yname == "pid" { return "Pid" }
    if yname == "udi-description" { return "UdiDescription" }
    if yname == "udi-name" { return "UdiName" }
    if yname == "clei" { return "Clei" }
    if yname == "eci" { return "Eci" }
    if yname == "top-assem-part-num" { return "TopAssemPartNum" }
    if yname == "top-assem-vid" { return "TopAssemVid" }
    if yname == "pca-num" { return "PcaNum" }
    if yname == "pcavid" { return "Pcavid" }
    if yname == "chassis-sid" { return "ChassisSid" }
    if yname == "dev-num1" { return "DevNum1" }
    if yname == "dev-num2" { return "DevNum2" }
    if yname == "dev-num3" { return "DevNum3" }
    if yname == "dev-num4" { return "DevNum4" }
    if yname == "dev-num5" { return "DevNum5" }
    if yname == "dev-num6" { return "DevNum6" }
    if yname == "dev-num7" { return "DevNum7" }
    if yname == "manu-test-data" { return "ManuTestData" }
    if yname == "asset-id" { return "AssetId" }
    if yname == "asset-alias" { return "AssetAlias" }
    if yname == "base-mac-address1" { return "BaseMacAddress1" }
    if yname == "mac-add-blk-size1" { return "MacAddBlkSize1" }
    if yname == "base-mac-address2" { return "BaseMacAddress2" }
    if yname == "mac-add-blk-size2" { return "MacAddBlkSize2" }
    if yname == "base-mac-address3" { return "BaseMacAddress3" }
    if yname == "mac-add-blk-size3" { return "MacAddBlkSize3" }
    if yname == "base-mac-address4" { return "BaseMacAddress4" }
    if yname == "mac-add-blk-size4" { return "MacAddBlkSize4" }
    if yname == "pcb-serial-num" { return "PcbSerialNum" }
    if yname == "power-supply-type" { return "PowerSupplyType" }
    if yname == "power-consumption" { return "PowerConsumption" }
    if yname == "block-signature" { return "BlockSignature" }
    if yname == "block-version" { return "BlockVersion" }
    if yname == "block-length" { return "BlockLength" }
    if yname == "block-checksum" { return "BlockChecksum" }
    if yname == "eeprom-size" { return "EepromSize" }
    if yname == "block-count" { return "BlockCount" }
    if yname == "fru-major-type" { return "FruMajorType" }
    if yname == "fru-minor-type" { return "FruMinorType" }
    if yname == "oem-string" { return "OemString" }
    if yname == "product-id" { return "ProductId" }
    if yname == "serial-number" { return "SerialNumber" }
    if yname == "part-number" { return "PartNumber" }
    if yname == "part-revision" { return "PartRevision" }
    if yname == "mfg-deviation" { return "MfgDeviation" }
    if yname == "hw-version" { return "HwVersion" }
    if yname == "mfg-bits" { return "MfgBits" }
    if yname == "engineer-use" { return "EngineerUse" }
    if yname == "snmpoid" { return "Snmpoid" }
    if yname == "rma-code" { return "RmaCode" }
    if yname == "rma" { return "Rma" }
    return ""
}

func (information *Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies_PowerSupply_Information) GetSegmentPath() string {
    return "information"
}

func (information *Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies_PowerSupply_Information) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rma" {
        return &information.Rma
    }
    return nil
}

func (information *Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies_PowerSupply_Information) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["rma"] = &information.Rma
    return children
}

func (information *Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies_PowerSupply_Information) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["description"] = information.Description
    leafs["idprom-format-rev"] = information.IdpromFormatRev
    leafs["controller-family"] = information.ControllerFamily
    leafs["controller-type"] = information.ControllerType
    leafs["vid"] = information.Vid
    leafs["hwid"] = information.Hwid
    leafs["pid"] = information.Pid
    leafs["udi-description"] = information.UdiDescription
    leafs["udi-name"] = information.UdiName
    leafs["clei"] = information.Clei
    leafs["eci"] = information.Eci
    leafs["top-assem-part-num"] = information.TopAssemPartNum
    leafs["top-assem-vid"] = information.TopAssemVid
    leafs["pca-num"] = information.PcaNum
    leafs["pcavid"] = information.Pcavid
    leafs["chassis-sid"] = information.ChassisSid
    leafs["dev-num1"] = information.DevNum1
    leafs["dev-num2"] = information.DevNum2
    leafs["dev-num3"] = information.DevNum3
    leafs["dev-num4"] = information.DevNum4
    leafs["dev-num5"] = information.DevNum5
    leafs["dev-num6"] = information.DevNum6
    leafs["dev-num7"] = information.DevNum7
    leafs["manu-test-data"] = information.ManuTestData
    leafs["asset-id"] = information.AssetId
    leafs["asset-alias"] = information.AssetAlias
    leafs["base-mac-address1"] = information.BaseMacAddress1
    leafs["mac-add-blk-size1"] = information.MacAddBlkSize1
    leafs["base-mac-address2"] = information.BaseMacAddress2
    leafs["mac-add-blk-size2"] = information.MacAddBlkSize2
    leafs["base-mac-address3"] = information.BaseMacAddress3
    leafs["mac-add-blk-size3"] = information.MacAddBlkSize3
    leafs["base-mac-address4"] = information.BaseMacAddress4
    leafs["mac-add-blk-size4"] = information.MacAddBlkSize4
    leafs["pcb-serial-num"] = information.PcbSerialNum
    leafs["power-supply-type"] = information.PowerSupplyType
    leafs["power-consumption"] = information.PowerConsumption
    leafs["block-signature"] = information.BlockSignature
    leafs["block-version"] = information.BlockVersion
    leafs["block-length"] = information.BlockLength
    leafs["block-checksum"] = information.BlockChecksum
    leafs["eeprom-size"] = information.EepromSize
    leafs["block-count"] = information.BlockCount
    leafs["fru-major-type"] = information.FruMajorType
    leafs["fru-minor-type"] = information.FruMinorType
    leafs["oem-string"] = information.OemString
    leafs["product-id"] = information.ProductId
    leafs["serial-number"] = information.SerialNumber
    leafs["part-number"] = information.PartNumber
    leafs["part-revision"] = information.PartRevision
    leafs["mfg-deviation"] = information.MfgDeviation
    leafs["hw-version"] = information.HwVersion
    leafs["mfg-bits"] = information.MfgBits
    leafs["engineer-use"] = information.EngineerUse
    leafs["snmpoid"] = information.Snmpoid
    leafs["rma-code"] = information.RmaCode
    return leafs
}

func (information *Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies_PowerSupply_Information) GetBundleName() string { return "cisco_ios_xr" }

func (information *Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies_PowerSupply_Information) GetYangName() string { return "information" }

func (information *Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies_PowerSupply_Information) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (information *Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies_PowerSupply_Information) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (information *Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies_PowerSupply_Information) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (information *Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies_PowerSupply_Information) SetParent(parent types.Entity) { information.parent = parent }

func (information *Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies_PowerSupply_Information) GetParent() types.Entity { return information.parent }

func (information *Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies_PowerSupply_Information) GetParentYangName() string { return "power-supply" }

// Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies_PowerSupply_Information_Rma
// RMA Data
type Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies_PowerSupply_Information_Rma struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Test history. The type is string with length: 0..255.
    TestHistory interface{}

    // RMA tracking number format is N-N-N. The type is string with length:
    // 0..255.
    RmaNumber interface{}

    // RMA history. The type is string with length: 0..255.
    RmaHistory interface{}
}

func (rma *Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies_PowerSupply_Information_Rma) GetFilter() yfilter.YFilter { return rma.YFilter }

func (rma *Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies_PowerSupply_Information_Rma) SetFilter(yf yfilter.YFilter) { rma.YFilter = yf }

func (rma *Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies_PowerSupply_Information_Rma) GetGoName(yname string) string {
    if yname == "test-history" { return "TestHistory" }
    if yname == "rma-number" { return "RmaNumber" }
    if yname == "rma-history" { return "RmaHistory" }
    return ""
}

func (rma *Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies_PowerSupply_Information_Rma) GetSegmentPath() string {
    return "rma"
}

func (rma *Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies_PowerSupply_Information_Rma) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rma *Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies_PowerSupply_Information_Rma) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rma *Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies_PowerSupply_Information_Rma) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["test-history"] = rma.TestHistory
    leafs["rma-number"] = rma.RmaNumber
    leafs["rma-history"] = rma.RmaHistory
    return leafs
}

func (rma *Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies_PowerSupply_Information_Rma) GetBundleName() string { return "cisco_ios_xr" }

func (rma *Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies_PowerSupply_Information_Rma) GetYangName() string { return "rma" }

func (rma *Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies_PowerSupply_Information_Rma) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rma *Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies_PowerSupply_Information_Rma) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rma *Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies_PowerSupply_Information_Rma) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rma *Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies_PowerSupply_Information_Rma) SetParent(parent types.Entity) { rma.parent = parent }

func (rma *Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies_PowerSupply_Information_Rma) GetParent() types.Entity { return rma.parent }

func (rma *Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies_PowerSupply_Information_Rma) GetParentYangName() string { return "information" }

// Diag_Racks_Rack_FanTraies
// Table for rack fan trays
type Diag_Racks_Rack_FanTraies struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Fan tray name. The type is slice of Diag_Racks_Rack_FanTraies_FanTray.
    FanTray []Diag_Racks_Rack_FanTraies_FanTray
}

func (fanTraies *Diag_Racks_Rack_FanTraies) GetFilter() yfilter.YFilter { return fanTraies.YFilter }

func (fanTraies *Diag_Racks_Rack_FanTraies) SetFilter(yf yfilter.YFilter) { fanTraies.YFilter = yf }

func (fanTraies *Diag_Racks_Rack_FanTraies) GetGoName(yname string) string {
    if yname == "fan-tray" { return "FanTray" }
    return ""
}

func (fanTraies *Diag_Racks_Rack_FanTraies) GetSegmentPath() string {
    return "fan-traies"
}

func (fanTraies *Diag_Racks_Rack_FanTraies) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "fan-tray" {
        for _, c := range fanTraies.FanTray {
            if fanTraies.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Diag_Racks_Rack_FanTraies_FanTray{}
        fanTraies.FanTray = append(fanTraies.FanTray, child)
        return &fanTraies.FanTray[len(fanTraies.FanTray)-1]
    }
    return nil
}

func (fanTraies *Diag_Racks_Rack_FanTraies) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range fanTraies.FanTray {
        children[fanTraies.FanTray[i].GetSegmentPath()] = &fanTraies.FanTray[i]
    }
    return children
}

func (fanTraies *Diag_Racks_Rack_FanTraies) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (fanTraies *Diag_Racks_Rack_FanTraies) GetBundleName() string { return "cisco_ios_xr" }

func (fanTraies *Diag_Racks_Rack_FanTraies) GetYangName() string { return "fan-traies" }

func (fanTraies *Diag_Racks_Rack_FanTraies) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fanTraies *Diag_Racks_Rack_FanTraies) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fanTraies *Diag_Racks_Rack_FanTraies) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fanTraies *Diag_Racks_Rack_FanTraies) SetParent(parent types.Entity) { fanTraies.parent = parent }

func (fanTraies *Diag_Racks_Rack_FanTraies) GetParent() types.Entity { return fanTraies.parent }

func (fanTraies *Diag_Racks_Rack_FanTraies) GetParentYangName() string { return "rack" }

// Diag_Racks_Rack_FanTraies_FanTray
// Fan tray name
type Diag_Racks_Rack_FanTraies_FanTray struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Fan tray name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    FanTrayName interface{}

    // Table for rack fans .
    Fanses Diag_Racks_Rack_FanTraies_FanTray_Fanses
}

func (fanTray *Diag_Racks_Rack_FanTraies_FanTray) GetFilter() yfilter.YFilter { return fanTray.YFilter }

func (fanTray *Diag_Racks_Rack_FanTraies_FanTray) SetFilter(yf yfilter.YFilter) { fanTray.YFilter = yf }

func (fanTray *Diag_Racks_Rack_FanTraies_FanTray) GetGoName(yname string) string {
    if yname == "fan-tray-name" { return "FanTrayName" }
    if yname == "fanses" { return "Fanses" }
    return ""
}

func (fanTray *Diag_Racks_Rack_FanTraies_FanTray) GetSegmentPath() string {
    return "fan-tray" + "[fan-tray-name='" + fmt.Sprintf("%v", fanTray.FanTrayName) + "']"
}

func (fanTray *Diag_Racks_Rack_FanTraies_FanTray) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "fanses" {
        return &fanTray.Fanses
    }
    return nil
}

func (fanTray *Diag_Racks_Rack_FanTraies_FanTray) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["fanses"] = &fanTray.Fanses
    return children
}

func (fanTray *Diag_Racks_Rack_FanTraies_FanTray) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["fan-tray-name"] = fanTray.FanTrayName
    return leafs
}

func (fanTray *Diag_Racks_Rack_FanTraies_FanTray) GetBundleName() string { return "cisco_ios_xr" }

func (fanTray *Diag_Racks_Rack_FanTraies_FanTray) GetYangName() string { return "fan-tray" }

func (fanTray *Diag_Racks_Rack_FanTraies_FanTray) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fanTray *Diag_Racks_Rack_FanTraies_FanTray) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fanTray *Diag_Racks_Rack_FanTraies_FanTray) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fanTray *Diag_Racks_Rack_FanTraies_FanTray) SetParent(parent types.Entity) { fanTray.parent = parent }

func (fanTray *Diag_Racks_Rack_FanTraies_FanTray) GetParent() types.Entity { return fanTray.parent }

func (fanTray *Diag_Racks_Rack_FanTraies_FanTray) GetParentYangName() string { return "fan-traies" }

// Diag_Racks_Rack_FanTraies_FanTray_Fanses
// Table for rack fans 
type Diag_Racks_Rack_FanTraies_FanTray_Fanses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Fan name. The type is slice of
    // Diag_Racks_Rack_FanTraies_FanTray_Fanses_Fans.
    Fans []Diag_Racks_Rack_FanTraies_FanTray_Fanses_Fans
}

func (fanses *Diag_Racks_Rack_FanTraies_FanTray_Fanses) GetFilter() yfilter.YFilter { return fanses.YFilter }

func (fanses *Diag_Racks_Rack_FanTraies_FanTray_Fanses) SetFilter(yf yfilter.YFilter) { fanses.YFilter = yf }

func (fanses *Diag_Racks_Rack_FanTraies_FanTray_Fanses) GetGoName(yname string) string {
    if yname == "fans" { return "Fans" }
    return ""
}

func (fanses *Diag_Racks_Rack_FanTraies_FanTray_Fanses) GetSegmentPath() string {
    return "fanses"
}

func (fanses *Diag_Racks_Rack_FanTraies_FanTray_Fanses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "fans" {
        for _, c := range fanses.Fans {
            if fanses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Diag_Racks_Rack_FanTraies_FanTray_Fanses_Fans{}
        fanses.Fans = append(fanses.Fans, child)
        return &fanses.Fans[len(fanses.Fans)-1]
    }
    return nil
}

func (fanses *Diag_Racks_Rack_FanTraies_FanTray_Fanses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range fanses.Fans {
        children[fanses.Fans[i].GetSegmentPath()] = &fanses.Fans[i]
    }
    return children
}

func (fanses *Diag_Racks_Rack_FanTraies_FanTray_Fanses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (fanses *Diag_Racks_Rack_FanTraies_FanTray_Fanses) GetBundleName() string { return "cisco_ios_xr" }

func (fanses *Diag_Racks_Rack_FanTraies_FanTray_Fanses) GetYangName() string { return "fanses" }

func (fanses *Diag_Racks_Rack_FanTraies_FanTray_Fanses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fanses *Diag_Racks_Rack_FanTraies_FanTray_Fanses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fanses *Diag_Racks_Rack_FanTraies_FanTray_Fanses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fanses *Diag_Racks_Rack_FanTraies_FanTray_Fanses) SetParent(parent types.Entity) { fanses.parent = parent }

func (fanses *Diag_Racks_Rack_FanTraies_FanTray_Fanses) GetParent() types.Entity { return fanses.parent }

func (fanses *Diag_Racks_Rack_FanTraies_FanTray_Fanses) GetParentYangName() string { return "fan-tray" }

// Diag_Racks_Rack_FanTraies_FanTray_Fanses_Fans
// Fan name
type Diag_Racks_Rack_FanTraies_FanTray_Fanses_Fans struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Fans name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    FansName interface{}

    // Basic information.
    Information Diag_Racks_Rack_FanTraies_FanTray_Fanses_Fans_Information
}

func (fans *Diag_Racks_Rack_FanTraies_FanTray_Fanses_Fans) GetFilter() yfilter.YFilter { return fans.YFilter }

func (fans *Diag_Racks_Rack_FanTraies_FanTray_Fanses_Fans) SetFilter(yf yfilter.YFilter) { fans.YFilter = yf }

func (fans *Diag_Racks_Rack_FanTraies_FanTray_Fanses_Fans) GetGoName(yname string) string {
    if yname == "fans-name" { return "FansName" }
    if yname == "information" { return "Information" }
    return ""
}

func (fans *Diag_Racks_Rack_FanTraies_FanTray_Fanses_Fans) GetSegmentPath() string {
    return "fans" + "[fans-name='" + fmt.Sprintf("%v", fans.FansName) + "']"
}

func (fans *Diag_Racks_Rack_FanTraies_FanTray_Fanses_Fans) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "information" {
        return &fans.Information
    }
    return nil
}

func (fans *Diag_Racks_Rack_FanTraies_FanTray_Fanses_Fans) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["information"] = &fans.Information
    return children
}

func (fans *Diag_Racks_Rack_FanTraies_FanTray_Fanses_Fans) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["fans-name"] = fans.FansName
    return leafs
}

func (fans *Diag_Racks_Rack_FanTraies_FanTray_Fanses_Fans) GetBundleName() string { return "cisco_ios_xr" }

func (fans *Diag_Racks_Rack_FanTraies_FanTray_Fanses_Fans) GetYangName() string { return "fans" }

func (fans *Diag_Racks_Rack_FanTraies_FanTray_Fanses_Fans) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fans *Diag_Racks_Rack_FanTraies_FanTray_Fanses_Fans) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fans *Diag_Racks_Rack_FanTraies_FanTray_Fanses_Fans) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fans *Diag_Racks_Rack_FanTraies_FanTray_Fanses_Fans) SetParent(parent types.Entity) { fans.parent = parent }

func (fans *Diag_Racks_Rack_FanTraies_FanTray_Fanses_Fans) GetParent() types.Entity { return fans.parent }

func (fans *Diag_Racks_Rack_FanTraies_FanTray_Fanses_Fans) GetParentYangName() string { return "fanses" }

// Diag_Racks_Rack_FanTraies_FanTray_Fanses_Fans_Information
// Basic information
type Diag_Racks_Rack_FanTraies_FanTray_Fanses_Fans_Information struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A textual description of physical entity. The type is string with length:
    // 0..255.
    Description interface{}

    // IDPROM Format Revision. The type is string with length: 0..255.
    IdpromFormatRev interface{}

    // Controller family. The type is string with length: 0..255.
    ControllerFamily interface{}

    // Controller type. The type is string with length: 0..255.
    ControllerType interface{}

    // Version ID. The type is string with length: 0..255.
    Vid interface{}

    // Hardware Revision. The type is string with length: 0..255.
    Hwid interface{}

    // Product ID. The type is string with length: 0..255.
    Pid interface{}

    // UDI description. The type is string with length: 0..255.
    UdiDescription interface{}

    // UDI name. The type is string with length: 0..255.
    UdiName interface{}

    // Common Language Equipment Identifier (CLEI) code. The type is string with
    // length: 0..255.
    Clei interface{}

    // Equipment Catalog Item (ECI) number. The type is string with length:
    // 0..255.
    Eci interface{}

    // Top assembly part number. The type is string with length: 0..255.
    TopAssemPartNum interface{}

    // Top assembly revision number. The type is string with length: 0..255.
    TopAssemVid interface{}

    // PCA number. The type is string with length: 0..255.
    PcaNum interface{}

    // PCA revision ID. The type is string with length: 0..255.
    Pcavid interface{}

    // Chassis serial number. The type is string with length: 0..255.
    ChassisSid interface{}

    // Deviation Number # 1. The type is string with length: 0..255.
    DevNum1 interface{}

    // Deviation Number # 2. The type is string with length: 0..255.
    DevNum2 interface{}

    // Deviation Number # 3. The type is string with length: 0..255.
    DevNum3 interface{}

    // Deviation Number # 4. The type is string with length: 0..255.
    DevNum4 interface{}

    // Deviation Number # 5. The type is string with length: 0..255.
    DevNum5 interface{}

    // Deviation Number # 6. The type is string with length: 0..255.
    DevNum6 interface{}

    // Deviation Number # 7. The type is string with length: 0..255.
    DevNum7 interface{}

    // Manufacturing Test Data. The type is string with length: 0..255.
    ManuTestData interface{}

    // Asset ID. The type is string with length: 0..255.
    AssetId interface{}

    // Asset Alias. The type is string with length: 0..255.
    AssetAlias interface{}

    // Base Mac Address #1. The type is string with length: 0..255.
    BaseMacAddress1 interface{}

    // Mac Address Block Size #1. The type is string with length: 0..255.
    MacAddBlkSize1 interface{}

    // Base Mac Address #2. The type is string with length: 0..255.
    BaseMacAddress2 interface{}

    // Mac Address Block Size #2. The type is string with length: 0..255.
    MacAddBlkSize2 interface{}

    // Base Mac Address #3. The type is string with length: 0..255.
    BaseMacAddress3 interface{}

    // Mac Address Block Size #3. The type is string with length: 0..255.
    MacAddBlkSize3 interface{}

    // Base Mac Address #4. The type is string with length: 0..255.
    BaseMacAddress4 interface{}

    // Mac Address Block Size #4. The type is string with length: 0..255.
    MacAddBlkSize4 interface{}

    // PCB Serial Number. The type is string with length: 0..255.
    PcbSerialNum interface{}

    // Power Supply Type. The type is string with length: 0..255.
    PowerSupplyType interface{}

    // Power Consumption. The type is string with length: 0..255.
    PowerConsumption interface{}

    // Block Signature. The type is string with length: 0..255.
    BlockSignature interface{}

    // Block Version. The type is string with length: 0..255.
    BlockVersion interface{}

    // Block Length. The type is string with length: 0..255.
    BlockLength interface{}

    // Block Checksum. The type is string with length: 0..255.
    BlockChecksum interface{}

    // EEPROM Size. The type is string with length: 0..255.
    EepromSize interface{}

    // Block Count. The type is string with length: 0..255.
    BlockCount interface{}

    // FRU Major Type. The type is string with length: 0..255.
    FruMajorType interface{}

    // FRU Minor Type. The type is string with length: 0..255.
    FruMinorType interface{}

    // OEM String. The type is string with length: 0..255.
    OemString interface{}

    // Product ID. The type is string with length: 0..255.
    ProductId interface{}

    // Serial Number. The type is string with length: 0..255.
    SerialNumber interface{}

    // Part Number. The type is string with length: 0..255.
    PartNumber interface{}

    // Part Revision. The type is string with length: 0..255.
    PartRevision interface{}

    // MFG Deviation. The type is string with length: 0..255.
    MfgDeviation interface{}

    // Hardware Version. The type is string with length: 0..255.
    HwVersion interface{}

    // MFG Bits. The type is string with length: 0..255.
    MfgBits interface{}

    // Engineer Use. The type is string with length: 0..255.
    EngineerUse interface{}

    // SNMP OID. The type is string with length: 0..255.
    Snmpoid interface{}

    // RMA Code. The type is string with length: 0..255.
    RmaCode interface{}

    // RMA Data.
    Rma Diag_Racks_Rack_FanTraies_FanTray_Fanses_Fans_Information_Rma
}

func (information *Diag_Racks_Rack_FanTraies_FanTray_Fanses_Fans_Information) GetFilter() yfilter.YFilter { return information.YFilter }

func (information *Diag_Racks_Rack_FanTraies_FanTray_Fanses_Fans_Information) SetFilter(yf yfilter.YFilter) { information.YFilter = yf }

func (information *Diag_Racks_Rack_FanTraies_FanTray_Fanses_Fans_Information) GetGoName(yname string) string {
    if yname == "description" { return "Description" }
    if yname == "idprom-format-rev" { return "IdpromFormatRev" }
    if yname == "controller-family" { return "ControllerFamily" }
    if yname == "controller-type" { return "ControllerType" }
    if yname == "vid" { return "Vid" }
    if yname == "hwid" { return "Hwid" }
    if yname == "pid" { return "Pid" }
    if yname == "udi-description" { return "UdiDescription" }
    if yname == "udi-name" { return "UdiName" }
    if yname == "clei" { return "Clei" }
    if yname == "eci" { return "Eci" }
    if yname == "top-assem-part-num" { return "TopAssemPartNum" }
    if yname == "top-assem-vid" { return "TopAssemVid" }
    if yname == "pca-num" { return "PcaNum" }
    if yname == "pcavid" { return "Pcavid" }
    if yname == "chassis-sid" { return "ChassisSid" }
    if yname == "dev-num1" { return "DevNum1" }
    if yname == "dev-num2" { return "DevNum2" }
    if yname == "dev-num3" { return "DevNum3" }
    if yname == "dev-num4" { return "DevNum4" }
    if yname == "dev-num5" { return "DevNum5" }
    if yname == "dev-num6" { return "DevNum6" }
    if yname == "dev-num7" { return "DevNum7" }
    if yname == "manu-test-data" { return "ManuTestData" }
    if yname == "asset-id" { return "AssetId" }
    if yname == "asset-alias" { return "AssetAlias" }
    if yname == "base-mac-address1" { return "BaseMacAddress1" }
    if yname == "mac-add-blk-size1" { return "MacAddBlkSize1" }
    if yname == "base-mac-address2" { return "BaseMacAddress2" }
    if yname == "mac-add-blk-size2" { return "MacAddBlkSize2" }
    if yname == "base-mac-address3" { return "BaseMacAddress3" }
    if yname == "mac-add-blk-size3" { return "MacAddBlkSize3" }
    if yname == "base-mac-address4" { return "BaseMacAddress4" }
    if yname == "mac-add-blk-size4" { return "MacAddBlkSize4" }
    if yname == "pcb-serial-num" { return "PcbSerialNum" }
    if yname == "power-supply-type" { return "PowerSupplyType" }
    if yname == "power-consumption" { return "PowerConsumption" }
    if yname == "block-signature" { return "BlockSignature" }
    if yname == "block-version" { return "BlockVersion" }
    if yname == "block-length" { return "BlockLength" }
    if yname == "block-checksum" { return "BlockChecksum" }
    if yname == "eeprom-size" { return "EepromSize" }
    if yname == "block-count" { return "BlockCount" }
    if yname == "fru-major-type" { return "FruMajorType" }
    if yname == "fru-minor-type" { return "FruMinorType" }
    if yname == "oem-string" { return "OemString" }
    if yname == "product-id" { return "ProductId" }
    if yname == "serial-number" { return "SerialNumber" }
    if yname == "part-number" { return "PartNumber" }
    if yname == "part-revision" { return "PartRevision" }
    if yname == "mfg-deviation" { return "MfgDeviation" }
    if yname == "hw-version" { return "HwVersion" }
    if yname == "mfg-bits" { return "MfgBits" }
    if yname == "engineer-use" { return "EngineerUse" }
    if yname == "snmpoid" { return "Snmpoid" }
    if yname == "rma-code" { return "RmaCode" }
    if yname == "rma" { return "Rma" }
    return ""
}

func (information *Diag_Racks_Rack_FanTraies_FanTray_Fanses_Fans_Information) GetSegmentPath() string {
    return "information"
}

func (information *Diag_Racks_Rack_FanTraies_FanTray_Fanses_Fans_Information) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rma" {
        return &information.Rma
    }
    return nil
}

func (information *Diag_Racks_Rack_FanTraies_FanTray_Fanses_Fans_Information) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["rma"] = &information.Rma
    return children
}

func (information *Diag_Racks_Rack_FanTraies_FanTray_Fanses_Fans_Information) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["description"] = information.Description
    leafs["idprom-format-rev"] = information.IdpromFormatRev
    leafs["controller-family"] = information.ControllerFamily
    leafs["controller-type"] = information.ControllerType
    leafs["vid"] = information.Vid
    leafs["hwid"] = information.Hwid
    leafs["pid"] = information.Pid
    leafs["udi-description"] = information.UdiDescription
    leafs["udi-name"] = information.UdiName
    leafs["clei"] = information.Clei
    leafs["eci"] = information.Eci
    leafs["top-assem-part-num"] = information.TopAssemPartNum
    leafs["top-assem-vid"] = information.TopAssemVid
    leafs["pca-num"] = information.PcaNum
    leafs["pcavid"] = information.Pcavid
    leafs["chassis-sid"] = information.ChassisSid
    leafs["dev-num1"] = information.DevNum1
    leafs["dev-num2"] = information.DevNum2
    leafs["dev-num3"] = information.DevNum3
    leafs["dev-num4"] = information.DevNum4
    leafs["dev-num5"] = information.DevNum5
    leafs["dev-num6"] = information.DevNum6
    leafs["dev-num7"] = information.DevNum7
    leafs["manu-test-data"] = information.ManuTestData
    leafs["asset-id"] = information.AssetId
    leafs["asset-alias"] = information.AssetAlias
    leafs["base-mac-address1"] = information.BaseMacAddress1
    leafs["mac-add-blk-size1"] = information.MacAddBlkSize1
    leafs["base-mac-address2"] = information.BaseMacAddress2
    leafs["mac-add-blk-size2"] = information.MacAddBlkSize2
    leafs["base-mac-address3"] = information.BaseMacAddress3
    leafs["mac-add-blk-size3"] = information.MacAddBlkSize3
    leafs["base-mac-address4"] = information.BaseMacAddress4
    leafs["mac-add-blk-size4"] = information.MacAddBlkSize4
    leafs["pcb-serial-num"] = information.PcbSerialNum
    leafs["power-supply-type"] = information.PowerSupplyType
    leafs["power-consumption"] = information.PowerConsumption
    leafs["block-signature"] = information.BlockSignature
    leafs["block-version"] = information.BlockVersion
    leafs["block-length"] = information.BlockLength
    leafs["block-checksum"] = information.BlockChecksum
    leafs["eeprom-size"] = information.EepromSize
    leafs["block-count"] = information.BlockCount
    leafs["fru-major-type"] = information.FruMajorType
    leafs["fru-minor-type"] = information.FruMinorType
    leafs["oem-string"] = information.OemString
    leafs["product-id"] = information.ProductId
    leafs["serial-number"] = information.SerialNumber
    leafs["part-number"] = information.PartNumber
    leafs["part-revision"] = information.PartRevision
    leafs["mfg-deviation"] = information.MfgDeviation
    leafs["hw-version"] = information.HwVersion
    leafs["mfg-bits"] = information.MfgBits
    leafs["engineer-use"] = information.EngineerUse
    leafs["snmpoid"] = information.Snmpoid
    leafs["rma-code"] = information.RmaCode
    return leafs
}

func (information *Diag_Racks_Rack_FanTraies_FanTray_Fanses_Fans_Information) GetBundleName() string { return "cisco_ios_xr" }

func (information *Diag_Racks_Rack_FanTraies_FanTray_Fanses_Fans_Information) GetYangName() string { return "information" }

func (information *Diag_Racks_Rack_FanTraies_FanTray_Fanses_Fans_Information) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (information *Diag_Racks_Rack_FanTraies_FanTray_Fanses_Fans_Information) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (information *Diag_Racks_Rack_FanTraies_FanTray_Fanses_Fans_Information) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (information *Diag_Racks_Rack_FanTraies_FanTray_Fanses_Fans_Information) SetParent(parent types.Entity) { information.parent = parent }

func (information *Diag_Racks_Rack_FanTraies_FanTray_Fanses_Fans_Information) GetParent() types.Entity { return information.parent }

func (information *Diag_Racks_Rack_FanTraies_FanTray_Fanses_Fans_Information) GetParentYangName() string { return "fans" }

// Diag_Racks_Rack_FanTraies_FanTray_Fanses_Fans_Information_Rma
// RMA Data
type Diag_Racks_Rack_FanTraies_FanTray_Fanses_Fans_Information_Rma struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Test history. The type is string with length: 0..255.
    TestHistory interface{}

    // RMA tracking number format is N-N-N. The type is string with length:
    // 0..255.
    RmaNumber interface{}

    // RMA history. The type is string with length: 0..255.
    RmaHistory interface{}
}

func (rma *Diag_Racks_Rack_FanTraies_FanTray_Fanses_Fans_Information_Rma) GetFilter() yfilter.YFilter { return rma.YFilter }

func (rma *Diag_Racks_Rack_FanTraies_FanTray_Fanses_Fans_Information_Rma) SetFilter(yf yfilter.YFilter) { rma.YFilter = yf }

func (rma *Diag_Racks_Rack_FanTraies_FanTray_Fanses_Fans_Information_Rma) GetGoName(yname string) string {
    if yname == "test-history" { return "TestHistory" }
    if yname == "rma-number" { return "RmaNumber" }
    if yname == "rma-history" { return "RmaHistory" }
    return ""
}

func (rma *Diag_Racks_Rack_FanTraies_FanTray_Fanses_Fans_Information_Rma) GetSegmentPath() string {
    return "rma"
}

func (rma *Diag_Racks_Rack_FanTraies_FanTray_Fanses_Fans_Information_Rma) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rma *Diag_Racks_Rack_FanTraies_FanTray_Fanses_Fans_Information_Rma) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rma *Diag_Racks_Rack_FanTraies_FanTray_Fanses_Fans_Information_Rma) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["test-history"] = rma.TestHistory
    leafs["rma-number"] = rma.RmaNumber
    leafs["rma-history"] = rma.RmaHistory
    return leafs
}

func (rma *Diag_Racks_Rack_FanTraies_FanTray_Fanses_Fans_Information_Rma) GetBundleName() string { return "cisco_ios_xr" }

func (rma *Diag_Racks_Rack_FanTraies_FanTray_Fanses_Fans_Information_Rma) GetYangName() string { return "rma" }

func (rma *Diag_Racks_Rack_FanTraies_FanTray_Fanses_Fans_Information_Rma) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rma *Diag_Racks_Rack_FanTraies_FanTray_Fanses_Fans_Information_Rma) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rma *Diag_Racks_Rack_FanTraies_FanTray_Fanses_Fans_Information_Rma) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rma *Diag_Racks_Rack_FanTraies_FanTray_Fanses_Fans_Information_Rma) SetParent(parent types.Entity) { rma.parent = parent }

func (rma *Diag_Racks_Rack_FanTraies_FanTray_Fanses_Fans_Information_Rma) GetParent() types.Entity { return rma.parent }

func (rma *Diag_Racks_Rack_FanTraies_FanTray_Fanses_Fans_Information_Rma) GetParentYangName() string { return "information" }

// Diag_Racks_Rack_Slots
// Table of slots
type Diag_Racks_Rack_Slots struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Slot name. The type is slice of Diag_Racks_Rack_Slots_Slot.
    Slot []Diag_Racks_Rack_Slots_Slot
}

func (slots *Diag_Racks_Rack_Slots) GetFilter() yfilter.YFilter { return slots.YFilter }

func (slots *Diag_Racks_Rack_Slots) SetFilter(yf yfilter.YFilter) { slots.YFilter = yf }

func (slots *Diag_Racks_Rack_Slots) GetGoName(yname string) string {
    if yname == "slot" { return "Slot" }
    return ""
}

func (slots *Diag_Racks_Rack_Slots) GetSegmentPath() string {
    return "slots"
}

func (slots *Diag_Racks_Rack_Slots) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "slot" {
        for _, c := range slots.Slot {
            if slots.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Diag_Racks_Rack_Slots_Slot{}
        slots.Slot = append(slots.Slot, child)
        return &slots.Slot[len(slots.Slot)-1]
    }
    return nil
}

func (slots *Diag_Racks_Rack_Slots) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range slots.Slot {
        children[slots.Slot[i].GetSegmentPath()] = &slots.Slot[i]
    }
    return children
}

func (slots *Diag_Racks_Rack_Slots) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (slots *Diag_Racks_Rack_Slots) GetBundleName() string { return "cisco_ios_xr" }

func (slots *Diag_Racks_Rack_Slots) GetYangName() string { return "slots" }

func (slots *Diag_Racks_Rack_Slots) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (slots *Diag_Racks_Rack_Slots) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (slots *Diag_Racks_Rack_Slots) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (slots *Diag_Racks_Rack_Slots) SetParent(parent types.Entity) { slots.parent = parent }

func (slots *Diag_Racks_Rack_Slots) GetParent() types.Entity { return slots.parent }

func (slots *Diag_Racks_Rack_Slots) GetParentYangName() string { return "rack" }

// Diag_Racks_Rack_Slots_Slot
// Slot name
type Diag_Racks_Rack_Slots_Slot struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Slot name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SlotName interface{}

    // Table of instances.
    Instances Diag_Racks_Rack_Slots_Slot_Instances
}

func (slot *Diag_Racks_Rack_Slots_Slot) GetFilter() yfilter.YFilter { return slot.YFilter }

func (slot *Diag_Racks_Rack_Slots_Slot) SetFilter(yf yfilter.YFilter) { slot.YFilter = yf }

func (slot *Diag_Racks_Rack_Slots_Slot) GetGoName(yname string) string {
    if yname == "slot-name" { return "SlotName" }
    if yname == "instances" { return "Instances" }
    return ""
}

func (slot *Diag_Racks_Rack_Slots_Slot) GetSegmentPath() string {
    return "slot" + "[slot-name='" + fmt.Sprintf("%v", slot.SlotName) + "']"
}

func (slot *Diag_Racks_Rack_Slots_Slot) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "instances" {
        return &slot.Instances
    }
    return nil
}

func (slot *Diag_Racks_Rack_Slots_Slot) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["instances"] = &slot.Instances
    return children
}

func (slot *Diag_Racks_Rack_Slots_Slot) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["slot-name"] = slot.SlotName
    return leafs
}

func (slot *Diag_Racks_Rack_Slots_Slot) GetBundleName() string { return "cisco_ios_xr" }

func (slot *Diag_Racks_Rack_Slots_Slot) GetYangName() string { return "slot" }

func (slot *Diag_Racks_Rack_Slots_Slot) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (slot *Diag_Racks_Rack_Slots_Slot) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (slot *Diag_Racks_Rack_Slots_Slot) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (slot *Diag_Racks_Rack_Slots_Slot) SetParent(parent types.Entity) { slot.parent = parent }

func (slot *Diag_Racks_Rack_Slots_Slot) GetParent() types.Entity { return slot.parent }

func (slot *Diag_Racks_Rack_Slots_Slot) GetParentYangName() string { return "slots" }

// Diag_Racks_Rack_Slots_Slot_Instances
// Table of instances
type Diag_Racks_Rack_Slots_Slot_Instances struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // instance number. The type is slice of
    // Diag_Racks_Rack_Slots_Slot_Instances_Instance.
    Instance []Diag_Racks_Rack_Slots_Slot_Instances_Instance
}

func (instances *Diag_Racks_Rack_Slots_Slot_Instances) GetFilter() yfilter.YFilter { return instances.YFilter }

func (instances *Diag_Racks_Rack_Slots_Slot_Instances) SetFilter(yf yfilter.YFilter) { instances.YFilter = yf }

func (instances *Diag_Racks_Rack_Slots_Slot_Instances) GetGoName(yname string) string {
    if yname == "instance" { return "Instance" }
    return ""
}

func (instances *Diag_Racks_Rack_Slots_Slot_Instances) GetSegmentPath() string {
    return "instances"
}

func (instances *Diag_Racks_Rack_Slots_Slot_Instances) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "instance" {
        for _, c := range instances.Instance {
            if instances.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Diag_Racks_Rack_Slots_Slot_Instances_Instance{}
        instances.Instance = append(instances.Instance, child)
        return &instances.Instance[len(instances.Instance)-1]
    }
    return nil
}

func (instances *Diag_Racks_Rack_Slots_Slot_Instances) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range instances.Instance {
        children[instances.Instance[i].GetSegmentPath()] = &instances.Instance[i]
    }
    return children
}

func (instances *Diag_Racks_Rack_Slots_Slot_Instances) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (instances *Diag_Racks_Rack_Slots_Slot_Instances) GetBundleName() string { return "cisco_ios_xr" }

func (instances *Diag_Racks_Rack_Slots_Slot_Instances) GetYangName() string { return "instances" }

func (instances *Diag_Racks_Rack_Slots_Slot_Instances) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (instances *Diag_Racks_Rack_Slots_Slot_Instances) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (instances *Diag_Racks_Rack_Slots_Slot_Instances) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (instances *Diag_Racks_Rack_Slots_Slot_Instances) SetParent(parent types.Entity) { instances.parent = parent }

func (instances *Diag_Racks_Rack_Slots_Slot_Instances) GetParent() types.Entity { return instances.parent }

func (instances *Diag_Racks_Rack_Slots_Slot_Instances) GetParentYangName() string { return "slot" }

// Diag_Racks_Rack_Slots_Slot_Instances_Instance
// instance number
type Diag_Racks_Rack_Slots_Slot_Instances_Instance struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Instance name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // Detail information.
    Detail Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail
}

func (instance *Diag_Racks_Rack_Slots_Slot_Instances_Instance) GetFilter() yfilter.YFilter { return instance.YFilter }

func (instance *Diag_Racks_Rack_Slots_Slot_Instances_Instance) SetFilter(yf yfilter.YFilter) { instance.YFilter = yf }

func (instance *Diag_Racks_Rack_Slots_Slot_Instances_Instance) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "detail" { return "Detail" }
    return ""
}

func (instance *Diag_Racks_Rack_Slots_Slot_Instances_Instance) GetSegmentPath() string {
    return "instance" + "[name='" + fmt.Sprintf("%v", instance.Name) + "']"
}

func (instance *Diag_Racks_Rack_Slots_Slot_Instances_Instance) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "detail" {
        return &instance.Detail
    }
    return nil
}

func (instance *Diag_Racks_Rack_Slots_Slot_Instances_Instance) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["detail"] = &instance.Detail
    return children
}

func (instance *Diag_Racks_Rack_Slots_Slot_Instances_Instance) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = instance.Name
    return leafs
}

func (instance *Diag_Racks_Rack_Slots_Slot_Instances_Instance) GetBundleName() string { return "cisco_ios_xr" }

func (instance *Diag_Racks_Rack_Slots_Slot_Instances_Instance) GetYangName() string { return "instance" }

func (instance *Diag_Racks_Rack_Slots_Slot_Instances_Instance) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (instance *Diag_Racks_Rack_Slots_Slot_Instances_Instance) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (instance *Diag_Racks_Rack_Slots_Slot_Instances_Instance) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (instance *Diag_Racks_Rack_Slots_Slot_Instances_Instance) SetParent(parent types.Entity) { instance.parent = parent }

func (instance *Diag_Racks_Rack_Slots_Slot_Instances_Instance) GetParent() types.Entity { return instance.parent }

func (instance *Diag_Racks_Rack_Slots_Slot_Instances_Instance) GetParentYangName() string { return "instances" }

// Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail
// Detail information
type Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node operational state . The type is string with length: 0..255.
    NodeOperationalState interface{}

    // Card instance.
    CardInstance Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_CardInstance
}

func (detail *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail) GetFilter() yfilter.YFilter { return detail.YFilter }

func (detail *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail) SetFilter(yf yfilter.YFilter) { detail.YFilter = yf }

func (detail *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail) GetGoName(yname string) string {
    if yname == "node-operational-state" { return "NodeOperationalState" }
    if yname == "card-instance" { return "CardInstance" }
    return ""
}

func (detail *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail) GetSegmentPath() string {
    return "detail"
}

func (detail *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "card-instance" {
        return &detail.CardInstance
    }
    return nil
}

func (detail *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["card-instance"] = &detail.CardInstance
    return children
}

func (detail *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-operational-state"] = detail.NodeOperationalState
    return leafs
}

func (detail *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail) GetBundleName() string { return "cisco_ios_xr" }

func (detail *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail) GetYangName() string { return "detail" }

func (detail *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (detail *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (detail *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (detail *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail) SetParent(parent types.Entity) { detail.parent = parent }

func (detail *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail) GetParent() types.Entity { return detail.parent }

func (detail *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail) GetParentYangName() string { return "instance" }

// Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_CardInstance
// Card instance
type Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_CardInstance struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A textual description of physical entity. The type is string with length:
    // 0..255.
    Description interface{}

    // IDPROM Format Revision. The type is string with length: 0..255.
    IdpromFormatRev interface{}

    // Controller family. The type is string with length: 0..255.
    ControllerFamily interface{}

    // Controller type. The type is string with length: 0..255.
    ControllerType interface{}

    // Version ID. The type is string with length: 0..255.
    Vid interface{}

    // Hardware Revision. The type is string with length: 0..255.
    Hwid interface{}

    // Product ID. The type is string with length: 0..255.
    Pid interface{}

    // UDI description. The type is string with length: 0..255.
    UdiDescription interface{}

    // UDI name. The type is string with length: 0..255.
    UdiName interface{}

    // Common Language Equipment Identifier (CLEI) code. The type is string with
    // length: 0..255.
    Clei interface{}

    // Equipment Catalog Item (ECI) number. The type is string with length:
    // 0..255.
    Eci interface{}

    // Top assembly part number. The type is string with length: 0..255.
    TopAssemPartNum interface{}

    // Top assembly revision number. The type is string with length: 0..255.
    TopAssemVid interface{}

    // PCA number. The type is string with length: 0..255.
    PcaNum interface{}

    // PCA revision ID. The type is string with length: 0..255.
    Pcavid interface{}

    // Chassis serial number. The type is string with length: 0..255.
    ChassisSid interface{}

    // Deviation Number # 1. The type is string with length: 0..255.
    DevNum1 interface{}

    // Deviation Number # 2. The type is string with length: 0..255.
    DevNum2 interface{}

    // Deviation Number # 3. The type is string with length: 0..255.
    DevNum3 interface{}

    // Deviation Number # 4. The type is string with length: 0..255.
    DevNum4 interface{}

    // Deviation Number # 5. The type is string with length: 0..255.
    DevNum5 interface{}

    // Deviation Number # 6. The type is string with length: 0..255.
    DevNum6 interface{}

    // Deviation Number # 7. The type is string with length: 0..255.
    DevNum7 interface{}

    // Manufacturing Test Data. The type is string with length: 0..255.
    ManuTestData interface{}

    // Asset ID. The type is string with length: 0..255.
    AssetId interface{}

    // Asset Alias. The type is string with length: 0..255.
    AssetAlias interface{}

    // Base Mac Address #1. The type is string with length: 0..255.
    BaseMacAddress1 interface{}

    // Mac Address Block Size #1. The type is string with length: 0..255.
    MacAddBlkSize1 interface{}

    // Base Mac Address #2. The type is string with length: 0..255.
    BaseMacAddress2 interface{}

    // Mac Address Block Size #2. The type is string with length: 0..255.
    MacAddBlkSize2 interface{}

    // Base Mac Address #3. The type is string with length: 0..255.
    BaseMacAddress3 interface{}

    // Mac Address Block Size #3. The type is string with length: 0..255.
    MacAddBlkSize3 interface{}

    // Base Mac Address #4. The type is string with length: 0..255.
    BaseMacAddress4 interface{}

    // Mac Address Block Size #4. The type is string with length: 0..255.
    MacAddBlkSize4 interface{}

    // PCB Serial Number. The type is string with length: 0..255.
    PcbSerialNum interface{}

    // Power Supply Type. The type is string with length: 0..255.
    PowerSupplyType interface{}

    // Power Consumption. The type is string with length: 0..255.
    PowerConsumption interface{}

    // Block Signature. The type is string with length: 0..255.
    BlockSignature interface{}

    // Block Version. The type is string with length: 0..255.
    BlockVersion interface{}

    // Block Length. The type is string with length: 0..255.
    BlockLength interface{}

    // Block Checksum. The type is string with length: 0..255.
    BlockChecksum interface{}

    // EEPROM Size. The type is string with length: 0..255.
    EepromSize interface{}

    // Block Count. The type is string with length: 0..255.
    BlockCount interface{}

    // FRU Major Type. The type is string with length: 0..255.
    FruMajorType interface{}

    // FRU Minor Type. The type is string with length: 0..255.
    FruMinorType interface{}

    // OEM String. The type is string with length: 0..255.
    OemString interface{}

    // Product ID. The type is string with length: 0..255.
    ProductId interface{}

    // Serial Number. The type is string with length: 0..255.
    SerialNumber interface{}

    // Part Number. The type is string with length: 0..255.
    PartNumber interface{}

    // Part Revision. The type is string with length: 0..255.
    PartRevision interface{}

    // MFG Deviation. The type is string with length: 0..255.
    MfgDeviation interface{}

    // Hardware Version. The type is string with length: 0..255.
    HwVersion interface{}

    // MFG Bits. The type is string with length: 0..255.
    MfgBits interface{}

    // Engineer Use. The type is string with length: 0..255.
    EngineerUse interface{}

    // SNMP OID. The type is string with length: 0..255.
    Snmpoid interface{}

    // RMA Code. The type is string with length: 0..255.
    RmaCode interface{}

    // RMA Data.
    Rma Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_CardInstance_Rma
}

func (cardInstance *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_CardInstance) GetFilter() yfilter.YFilter { return cardInstance.YFilter }

func (cardInstance *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_CardInstance) SetFilter(yf yfilter.YFilter) { cardInstance.YFilter = yf }

func (cardInstance *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_CardInstance) GetGoName(yname string) string {
    if yname == "description" { return "Description" }
    if yname == "idprom-format-rev" { return "IdpromFormatRev" }
    if yname == "controller-family" { return "ControllerFamily" }
    if yname == "controller-type" { return "ControllerType" }
    if yname == "vid" { return "Vid" }
    if yname == "hwid" { return "Hwid" }
    if yname == "pid" { return "Pid" }
    if yname == "udi-description" { return "UdiDescription" }
    if yname == "udi-name" { return "UdiName" }
    if yname == "clei" { return "Clei" }
    if yname == "eci" { return "Eci" }
    if yname == "top-assem-part-num" { return "TopAssemPartNum" }
    if yname == "top-assem-vid" { return "TopAssemVid" }
    if yname == "pca-num" { return "PcaNum" }
    if yname == "pcavid" { return "Pcavid" }
    if yname == "chassis-sid" { return "ChassisSid" }
    if yname == "dev-num1" { return "DevNum1" }
    if yname == "dev-num2" { return "DevNum2" }
    if yname == "dev-num3" { return "DevNum3" }
    if yname == "dev-num4" { return "DevNum4" }
    if yname == "dev-num5" { return "DevNum5" }
    if yname == "dev-num6" { return "DevNum6" }
    if yname == "dev-num7" { return "DevNum7" }
    if yname == "manu-test-data" { return "ManuTestData" }
    if yname == "asset-id" { return "AssetId" }
    if yname == "asset-alias" { return "AssetAlias" }
    if yname == "base-mac-address1" { return "BaseMacAddress1" }
    if yname == "mac-add-blk-size1" { return "MacAddBlkSize1" }
    if yname == "base-mac-address2" { return "BaseMacAddress2" }
    if yname == "mac-add-blk-size2" { return "MacAddBlkSize2" }
    if yname == "base-mac-address3" { return "BaseMacAddress3" }
    if yname == "mac-add-blk-size3" { return "MacAddBlkSize3" }
    if yname == "base-mac-address4" { return "BaseMacAddress4" }
    if yname == "mac-add-blk-size4" { return "MacAddBlkSize4" }
    if yname == "pcb-serial-num" { return "PcbSerialNum" }
    if yname == "power-supply-type" { return "PowerSupplyType" }
    if yname == "power-consumption" { return "PowerConsumption" }
    if yname == "block-signature" { return "BlockSignature" }
    if yname == "block-version" { return "BlockVersion" }
    if yname == "block-length" { return "BlockLength" }
    if yname == "block-checksum" { return "BlockChecksum" }
    if yname == "eeprom-size" { return "EepromSize" }
    if yname == "block-count" { return "BlockCount" }
    if yname == "fru-major-type" { return "FruMajorType" }
    if yname == "fru-minor-type" { return "FruMinorType" }
    if yname == "oem-string" { return "OemString" }
    if yname == "product-id" { return "ProductId" }
    if yname == "serial-number" { return "SerialNumber" }
    if yname == "part-number" { return "PartNumber" }
    if yname == "part-revision" { return "PartRevision" }
    if yname == "mfg-deviation" { return "MfgDeviation" }
    if yname == "hw-version" { return "HwVersion" }
    if yname == "mfg-bits" { return "MfgBits" }
    if yname == "engineer-use" { return "EngineerUse" }
    if yname == "snmpoid" { return "Snmpoid" }
    if yname == "rma-code" { return "RmaCode" }
    if yname == "rma" { return "Rma" }
    return ""
}

func (cardInstance *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_CardInstance) GetSegmentPath() string {
    return "card-instance"
}

func (cardInstance *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_CardInstance) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rma" {
        return &cardInstance.Rma
    }
    return nil
}

func (cardInstance *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_CardInstance) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["rma"] = &cardInstance.Rma
    return children
}

func (cardInstance *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_CardInstance) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["description"] = cardInstance.Description
    leafs["idprom-format-rev"] = cardInstance.IdpromFormatRev
    leafs["controller-family"] = cardInstance.ControllerFamily
    leafs["controller-type"] = cardInstance.ControllerType
    leafs["vid"] = cardInstance.Vid
    leafs["hwid"] = cardInstance.Hwid
    leafs["pid"] = cardInstance.Pid
    leafs["udi-description"] = cardInstance.UdiDescription
    leafs["udi-name"] = cardInstance.UdiName
    leafs["clei"] = cardInstance.Clei
    leafs["eci"] = cardInstance.Eci
    leafs["top-assem-part-num"] = cardInstance.TopAssemPartNum
    leafs["top-assem-vid"] = cardInstance.TopAssemVid
    leafs["pca-num"] = cardInstance.PcaNum
    leafs["pcavid"] = cardInstance.Pcavid
    leafs["chassis-sid"] = cardInstance.ChassisSid
    leafs["dev-num1"] = cardInstance.DevNum1
    leafs["dev-num2"] = cardInstance.DevNum2
    leafs["dev-num3"] = cardInstance.DevNum3
    leafs["dev-num4"] = cardInstance.DevNum4
    leafs["dev-num5"] = cardInstance.DevNum5
    leafs["dev-num6"] = cardInstance.DevNum6
    leafs["dev-num7"] = cardInstance.DevNum7
    leafs["manu-test-data"] = cardInstance.ManuTestData
    leafs["asset-id"] = cardInstance.AssetId
    leafs["asset-alias"] = cardInstance.AssetAlias
    leafs["base-mac-address1"] = cardInstance.BaseMacAddress1
    leafs["mac-add-blk-size1"] = cardInstance.MacAddBlkSize1
    leafs["base-mac-address2"] = cardInstance.BaseMacAddress2
    leafs["mac-add-blk-size2"] = cardInstance.MacAddBlkSize2
    leafs["base-mac-address3"] = cardInstance.BaseMacAddress3
    leafs["mac-add-blk-size3"] = cardInstance.MacAddBlkSize3
    leafs["base-mac-address4"] = cardInstance.BaseMacAddress4
    leafs["mac-add-blk-size4"] = cardInstance.MacAddBlkSize4
    leafs["pcb-serial-num"] = cardInstance.PcbSerialNum
    leafs["power-supply-type"] = cardInstance.PowerSupplyType
    leafs["power-consumption"] = cardInstance.PowerConsumption
    leafs["block-signature"] = cardInstance.BlockSignature
    leafs["block-version"] = cardInstance.BlockVersion
    leafs["block-length"] = cardInstance.BlockLength
    leafs["block-checksum"] = cardInstance.BlockChecksum
    leafs["eeprom-size"] = cardInstance.EepromSize
    leafs["block-count"] = cardInstance.BlockCount
    leafs["fru-major-type"] = cardInstance.FruMajorType
    leafs["fru-minor-type"] = cardInstance.FruMinorType
    leafs["oem-string"] = cardInstance.OemString
    leafs["product-id"] = cardInstance.ProductId
    leafs["serial-number"] = cardInstance.SerialNumber
    leafs["part-number"] = cardInstance.PartNumber
    leafs["part-revision"] = cardInstance.PartRevision
    leafs["mfg-deviation"] = cardInstance.MfgDeviation
    leafs["hw-version"] = cardInstance.HwVersion
    leafs["mfg-bits"] = cardInstance.MfgBits
    leafs["engineer-use"] = cardInstance.EngineerUse
    leafs["snmpoid"] = cardInstance.Snmpoid
    leafs["rma-code"] = cardInstance.RmaCode
    return leafs
}

func (cardInstance *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_CardInstance) GetBundleName() string { return "cisco_ios_xr" }

func (cardInstance *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_CardInstance) GetYangName() string { return "card-instance" }

func (cardInstance *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_CardInstance) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cardInstance *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_CardInstance) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cardInstance *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_CardInstance) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cardInstance *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_CardInstance) SetParent(parent types.Entity) { cardInstance.parent = parent }

func (cardInstance *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_CardInstance) GetParent() types.Entity { return cardInstance.parent }

func (cardInstance *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_CardInstance) GetParentYangName() string { return "detail" }

// Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_CardInstance_Rma
// RMA Data
type Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_CardInstance_Rma struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Test history. The type is string with length: 0..255.
    TestHistory interface{}

    // RMA tracking number format is N-N-N. The type is string with length:
    // 0..255.
    RmaNumber interface{}

    // RMA history. The type is string with length: 0..255.
    RmaHistory interface{}
}

func (rma *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_CardInstance_Rma) GetFilter() yfilter.YFilter { return rma.YFilter }

func (rma *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_CardInstance_Rma) SetFilter(yf yfilter.YFilter) { rma.YFilter = yf }

func (rma *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_CardInstance_Rma) GetGoName(yname string) string {
    if yname == "test-history" { return "TestHistory" }
    if yname == "rma-number" { return "RmaNumber" }
    if yname == "rma-history" { return "RmaHistory" }
    return ""
}

func (rma *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_CardInstance_Rma) GetSegmentPath() string {
    return "rma"
}

func (rma *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_CardInstance_Rma) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rma *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_CardInstance_Rma) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rma *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_CardInstance_Rma) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["test-history"] = rma.TestHistory
    leafs["rma-number"] = rma.RmaNumber
    leafs["rma-history"] = rma.RmaHistory
    return leafs
}

func (rma *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_CardInstance_Rma) GetBundleName() string { return "cisco_ios_xr" }

func (rma *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_CardInstance_Rma) GetYangName() string { return "rma" }

func (rma *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_CardInstance_Rma) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rma *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_CardInstance_Rma) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rma *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_CardInstance_Rma) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rma *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_CardInstance_Rma) SetParent(parent types.Entity) { rma.parent = parent }

func (rma *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_CardInstance_Rma) GetParent() types.Entity { return rma.parent }

func (rma *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_CardInstance_Rma) GetParentYangName() string { return "card-instance" }

// Diag_Racks_Rack_Chassis
// Chassis information
type Diag_Racks_Rack_Chassis struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A textual description of physical entity. The type is string with length:
    // 0..255.
    Description interface{}

    // IDPROM Format Revision. The type is string with length: 0..255.
    IdpromFormatRev interface{}

    // Controller family. The type is string with length: 0..255.
    ControllerFamily interface{}

    // Controller type. The type is string with length: 0..255.
    ControllerType interface{}

    // Version ID. The type is string with length: 0..255.
    Vid interface{}

    // Hardware Revision. The type is string with length: 0..255.
    Hwid interface{}

    // Product ID. The type is string with length: 0..255.
    Pid interface{}

    // UDI description. The type is string with length: 0..255.
    UdiDescription interface{}

    // UDI name. The type is string with length: 0..255.
    UdiName interface{}

    // Common Language Equipment Identifier (CLEI) code. The type is string with
    // length: 0..255.
    Clei interface{}

    // Equipment Catalog Item (ECI) number. The type is string with length:
    // 0..255.
    Eci interface{}

    // Top assembly part number. The type is string with length: 0..255.
    TopAssemPartNum interface{}

    // Top assembly revision number. The type is string with length: 0..255.
    TopAssemVid interface{}

    // PCA number. The type is string with length: 0..255.
    PcaNum interface{}

    // PCA revision ID. The type is string with length: 0..255.
    Pcavid interface{}

    // Chassis serial number. The type is string with length: 0..255.
    ChassisSid interface{}

    // Deviation Number # 1. The type is string with length: 0..255.
    DevNum1 interface{}

    // Deviation Number # 2. The type is string with length: 0..255.
    DevNum2 interface{}

    // Deviation Number # 3. The type is string with length: 0..255.
    DevNum3 interface{}

    // Deviation Number # 4. The type is string with length: 0..255.
    DevNum4 interface{}

    // Deviation Number # 5. The type is string with length: 0..255.
    DevNum5 interface{}

    // Deviation Number # 6. The type is string with length: 0..255.
    DevNum6 interface{}

    // Deviation Number # 7. The type is string with length: 0..255.
    DevNum7 interface{}

    // Manufacturing Test Data. The type is string with length: 0..255.
    ManuTestData interface{}

    // Asset ID. The type is string with length: 0..255.
    AssetId interface{}

    // Asset Alias. The type is string with length: 0..255.
    AssetAlias interface{}

    // Base Mac Address #1. The type is string with length: 0..255.
    BaseMacAddress1 interface{}

    // Mac Address Block Size #1. The type is string with length: 0..255.
    MacAddBlkSize1 interface{}

    // Base Mac Address #2. The type is string with length: 0..255.
    BaseMacAddress2 interface{}

    // Mac Address Block Size #2. The type is string with length: 0..255.
    MacAddBlkSize2 interface{}

    // Base Mac Address #3. The type is string with length: 0..255.
    BaseMacAddress3 interface{}

    // Mac Address Block Size #3. The type is string with length: 0..255.
    MacAddBlkSize3 interface{}

    // Base Mac Address #4. The type is string with length: 0..255.
    BaseMacAddress4 interface{}

    // Mac Address Block Size #4. The type is string with length: 0..255.
    MacAddBlkSize4 interface{}

    // PCB Serial Number. The type is string with length: 0..255.
    PcbSerialNum interface{}

    // Power Supply Type. The type is string with length: 0..255.
    PowerSupplyType interface{}

    // Power Consumption. The type is string with length: 0..255.
    PowerConsumption interface{}

    // Block Signature. The type is string with length: 0..255.
    BlockSignature interface{}

    // Block Version. The type is string with length: 0..255.
    BlockVersion interface{}

    // Block Length. The type is string with length: 0..255.
    BlockLength interface{}

    // Block Checksum. The type is string with length: 0..255.
    BlockChecksum interface{}

    // EEPROM Size. The type is string with length: 0..255.
    EepromSize interface{}

    // Block Count. The type is string with length: 0..255.
    BlockCount interface{}

    // FRU Major Type. The type is string with length: 0..255.
    FruMajorType interface{}

    // FRU Minor Type. The type is string with length: 0..255.
    FruMinorType interface{}

    // OEM String. The type is string with length: 0..255.
    OemString interface{}

    // Product ID. The type is string with length: 0..255.
    ProductId interface{}

    // Serial Number. The type is string with length: 0..255.
    SerialNumber interface{}

    // Part Number. The type is string with length: 0..255.
    PartNumber interface{}

    // Part Revision. The type is string with length: 0..255.
    PartRevision interface{}

    // MFG Deviation. The type is string with length: 0..255.
    MfgDeviation interface{}

    // Hardware Version. The type is string with length: 0..255.
    HwVersion interface{}

    // MFG Bits. The type is string with length: 0..255.
    MfgBits interface{}

    // Engineer Use. The type is string with length: 0..255.
    EngineerUse interface{}

    // SNMP OID. The type is string with length: 0..255.
    Snmpoid interface{}

    // RMA Code. The type is string with length: 0..255.
    RmaCode interface{}

    // RMA Data.
    Rma Diag_Racks_Rack_Chassis_Rma
}

func (chassis *Diag_Racks_Rack_Chassis) GetFilter() yfilter.YFilter { return chassis.YFilter }

func (chassis *Diag_Racks_Rack_Chassis) SetFilter(yf yfilter.YFilter) { chassis.YFilter = yf }

func (chassis *Diag_Racks_Rack_Chassis) GetGoName(yname string) string {
    if yname == "description" { return "Description" }
    if yname == "idprom-format-rev" { return "IdpromFormatRev" }
    if yname == "controller-family" { return "ControllerFamily" }
    if yname == "controller-type" { return "ControllerType" }
    if yname == "vid" { return "Vid" }
    if yname == "hwid" { return "Hwid" }
    if yname == "pid" { return "Pid" }
    if yname == "udi-description" { return "UdiDescription" }
    if yname == "udi-name" { return "UdiName" }
    if yname == "clei" { return "Clei" }
    if yname == "eci" { return "Eci" }
    if yname == "top-assem-part-num" { return "TopAssemPartNum" }
    if yname == "top-assem-vid" { return "TopAssemVid" }
    if yname == "pca-num" { return "PcaNum" }
    if yname == "pcavid" { return "Pcavid" }
    if yname == "chassis-sid" { return "ChassisSid" }
    if yname == "dev-num1" { return "DevNum1" }
    if yname == "dev-num2" { return "DevNum2" }
    if yname == "dev-num3" { return "DevNum3" }
    if yname == "dev-num4" { return "DevNum4" }
    if yname == "dev-num5" { return "DevNum5" }
    if yname == "dev-num6" { return "DevNum6" }
    if yname == "dev-num7" { return "DevNum7" }
    if yname == "manu-test-data" { return "ManuTestData" }
    if yname == "asset-id" { return "AssetId" }
    if yname == "asset-alias" { return "AssetAlias" }
    if yname == "base-mac-address1" { return "BaseMacAddress1" }
    if yname == "mac-add-blk-size1" { return "MacAddBlkSize1" }
    if yname == "base-mac-address2" { return "BaseMacAddress2" }
    if yname == "mac-add-blk-size2" { return "MacAddBlkSize2" }
    if yname == "base-mac-address3" { return "BaseMacAddress3" }
    if yname == "mac-add-blk-size3" { return "MacAddBlkSize3" }
    if yname == "base-mac-address4" { return "BaseMacAddress4" }
    if yname == "mac-add-blk-size4" { return "MacAddBlkSize4" }
    if yname == "pcb-serial-num" { return "PcbSerialNum" }
    if yname == "power-supply-type" { return "PowerSupplyType" }
    if yname == "power-consumption" { return "PowerConsumption" }
    if yname == "block-signature" { return "BlockSignature" }
    if yname == "block-version" { return "BlockVersion" }
    if yname == "block-length" { return "BlockLength" }
    if yname == "block-checksum" { return "BlockChecksum" }
    if yname == "eeprom-size" { return "EepromSize" }
    if yname == "block-count" { return "BlockCount" }
    if yname == "fru-major-type" { return "FruMajorType" }
    if yname == "fru-minor-type" { return "FruMinorType" }
    if yname == "oem-string" { return "OemString" }
    if yname == "product-id" { return "ProductId" }
    if yname == "serial-number" { return "SerialNumber" }
    if yname == "part-number" { return "PartNumber" }
    if yname == "part-revision" { return "PartRevision" }
    if yname == "mfg-deviation" { return "MfgDeviation" }
    if yname == "hw-version" { return "HwVersion" }
    if yname == "mfg-bits" { return "MfgBits" }
    if yname == "engineer-use" { return "EngineerUse" }
    if yname == "snmpoid" { return "Snmpoid" }
    if yname == "rma-code" { return "RmaCode" }
    if yname == "rma" { return "Rma" }
    return ""
}

func (chassis *Diag_Racks_Rack_Chassis) GetSegmentPath() string {
    return "chassis"
}

func (chassis *Diag_Racks_Rack_Chassis) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rma" {
        return &chassis.Rma
    }
    return nil
}

func (chassis *Diag_Racks_Rack_Chassis) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["rma"] = &chassis.Rma
    return children
}

func (chassis *Diag_Racks_Rack_Chassis) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["description"] = chassis.Description
    leafs["idprom-format-rev"] = chassis.IdpromFormatRev
    leafs["controller-family"] = chassis.ControllerFamily
    leafs["controller-type"] = chassis.ControllerType
    leafs["vid"] = chassis.Vid
    leafs["hwid"] = chassis.Hwid
    leafs["pid"] = chassis.Pid
    leafs["udi-description"] = chassis.UdiDescription
    leafs["udi-name"] = chassis.UdiName
    leafs["clei"] = chassis.Clei
    leafs["eci"] = chassis.Eci
    leafs["top-assem-part-num"] = chassis.TopAssemPartNum
    leafs["top-assem-vid"] = chassis.TopAssemVid
    leafs["pca-num"] = chassis.PcaNum
    leafs["pcavid"] = chassis.Pcavid
    leafs["chassis-sid"] = chassis.ChassisSid
    leafs["dev-num1"] = chassis.DevNum1
    leafs["dev-num2"] = chassis.DevNum2
    leafs["dev-num3"] = chassis.DevNum3
    leafs["dev-num4"] = chassis.DevNum4
    leafs["dev-num5"] = chassis.DevNum5
    leafs["dev-num6"] = chassis.DevNum6
    leafs["dev-num7"] = chassis.DevNum7
    leafs["manu-test-data"] = chassis.ManuTestData
    leafs["asset-id"] = chassis.AssetId
    leafs["asset-alias"] = chassis.AssetAlias
    leafs["base-mac-address1"] = chassis.BaseMacAddress1
    leafs["mac-add-blk-size1"] = chassis.MacAddBlkSize1
    leafs["base-mac-address2"] = chassis.BaseMacAddress2
    leafs["mac-add-blk-size2"] = chassis.MacAddBlkSize2
    leafs["base-mac-address3"] = chassis.BaseMacAddress3
    leafs["mac-add-blk-size3"] = chassis.MacAddBlkSize3
    leafs["base-mac-address4"] = chassis.BaseMacAddress4
    leafs["mac-add-blk-size4"] = chassis.MacAddBlkSize4
    leafs["pcb-serial-num"] = chassis.PcbSerialNum
    leafs["power-supply-type"] = chassis.PowerSupplyType
    leafs["power-consumption"] = chassis.PowerConsumption
    leafs["block-signature"] = chassis.BlockSignature
    leafs["block-version"] = chassis.BlockVersion
    leafs["block-length"] = chassis.BlockLength
    leafs["block-checksum"] = chassis.BlockChecksum
    leafs["eeprom-size"] = chassis.EepromSize
    leafs["block-count"] = chassis.BlockCount
    leafs["fru-major-type"] = chassis.FruMajorType
    leafs["fru-minor-type"] = chassis.FruMinorType
    leafs["oem-string"] = chassis.OemString
    leafs["product-id"] = chassis.ProductId
    leafs["serial-number"] = chassis.SerialNumber
    leafs["part-number"] = chassis.PartNumber
    leafs["part-revision"] = chassis.PartRevision
    leafs["mfg-deviation"] = chassis.MfgDeviation
    leafs["hw-version"] = chassis.HwVersion
    leafs["mfg-bits"] = chassis.MfgBits
    leafs["engineer-use"] = chassis.EngineerUse
    leafs["snmpoid"] = chassis.Snmpoid
    leafs["rma-code"] = chassis.RmaCode
    return leafs
}

func (chassis *Diag_Racks_Rack_Chassis) GetBundleName() string { return "cisco_ios_xr" }

func (chassis *Diag_Racks_Rack_Chassis) GetYangName() string { return "chassis" }

func (chassis *Diag_Racks_Rack_Chassis) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (chassis *Diag_Racks_Rack_Chassis) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (chassis *Diag_Racks_Rack_Chassis) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (chassis *Diag_Racks_Rack_Chassis) SetParent(parent types.Entity) { chassis.parent = parent }

func (chassis *Diag_Racks_Rack_Chassis) GetParent() types.Entity { return chassis.parent }

func (chassis *Diag_Racks_Rack_Chassis) GetParentYangName() string { return "rack" }

// Diag_Racks_Rack_Chassis_Rma
// RMA Data
type Diag_Racks_Rack_Chassis_Rma struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Test history. The type is string with length: 0..255.
    TestHistory interface{}

    // RMA tracking number format is N-N-N. The type is string with length:
    // 0..255.
    RmaNumber interface{}

    // RMA history. The type is string with length: 0..255.
    RmaHistory interface{}
}

func (rma *Diag_Racks_Rack_Chassis_Rma) GetFilter() yfilter.YFilter { return rma.YFilter }

func (rma *Diag_Racks_Rack_Chassis_Rma) SetFilter(yf yfilter.YFilter) { rma.YFilter = yf }

func (rma *Diag_Racks_Rack_Chassis_Rma) GetGoName(yname string) string {
    if yname == "test-history" { return "TestHistory" }
    if yname == "rma-number" { return "RmaNumber" }
    if yname == "rma-history" { return "RmaHistory" }
    return ""
}

func (rma *Diag_Racks_Rack_Chassis_Rma) GetSegmentPath() string {
    return "rma"
}

func (rma *Diag_Racks_Rack_Chassis_Rma) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rma *Diag_Racks_Rack_Chassis_Rma) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rma *Diag_Racks_Rack_Chassis_Rma) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["test-history"] = rma.TestHistory
    leafs["rma-number"] = rma.RmaNumber
    leafs["rma-history"] = rma.RmaHistory
    return leafs
}

func (rma *Diag_Racks_Rack_Chassis_Rma) GetBundleName() string { return "cisco_ios_xr" }

func (rma *Diag_Racks_Rack_Chassis_Rma) GetYangName() string { return "rma" }

func (rma *Diag_Racks_Rack_Chassis_Rma) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rma *Diag_Racks_Rack_Chassis_Rma) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rma *Diag_Racks_Rack_Chassis_Rma) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rma *Diag_Racks_Rack_Chassis_Rma) SetParent(parent types.Entity) { rma.parent = parent }

func (rma *Diag_Racks_Rack_Chassis_Rma) GetParent() types.Entity { return rma.parent }

func (rma *Diag_Racks_Rack_Chassis_Rma) GetParentYangName() string { return "chassis" }

