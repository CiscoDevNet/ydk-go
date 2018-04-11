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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of racks.
    Racks Diag_Racks
}

func (diag *Diag) GetEntityData() *types.CommonEntityData {
    diag.EntityData.YFilter = diag.YFilter
    diag.EntityData.YangName = "diag"
    diag.EntityData.BundleName = "cisco_ios_xr"
    diag.EntityData.ParentYangName = "Cisco-IOS-XR-sdr-invmgr-diag-oper"
    diag.EntityData.SegmentPath = "Cisco-IOS-XR-sdr-invmgr-diag-oper:diag"
    diag.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    diag.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    diag.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    diag.EntityData.Children = make(map[string]types.YChild)
    diag.EntityData.Children["racks"] = types.YChild{"Racks", &diag.Racks}
    diag.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(diag.EntityData)
}

// Diag_Racks
// Table of racks
type Diag_Racks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Rack name. The type is slice of Diag_Racks_Rack.
    Rack []Diag_Racks_Rack
}

func (racks *Diag_Racks) GetEntityData() *types.CommonEntityData {
    racks.EntityData.YFilter = racks.YFilter
    racks.EntityData.YangName = "racks"
    racks.EntityData.BundleName = "cisco_ios_xr"
    racks.EntityData.ParentYangName = "diag"
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

// Diag_Racks_Rack
// Rack name
type Diag_Racks_Rack struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Rack name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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

func (rack *Diag_Racks_Rack) GetEntityData() *types.CommonEntityData {
    rack.EntityData.YFilter = rack.YFilter
    rack.EntityData.YangName = "rack"
    rack.EntityData.BundleName = "cisco_ios_xr"
    rack.EntityData.ParentYangName = "racks"
    rack.EntityData.SegmentPath = "rack" + "[rack-name='" + fmt.Sprintf("%v", rack.RackName) + "']"
    rack.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rack.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rack.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rack.EntityData.Children = make(map[string]types.YChild)
    rack.EntityData.Children["power-shelfs"] = types.YChild{"PowerShelfs", &rack.PowerShelfs}
    rack.EntityData.Children["fan-traies"] = types.YChild{"FanTraies", &rack.FanTraies}
    rack.EntityData.Children["slots"] = types.YChild{"Slots", &rack.Slots}
    rack.EntityData.Children["chassis"] = types.YChild{"Chassis", &rack.Chassis}
    rack.EntityData.Leafs = make(map[string]types.YLeaf)
    rack.EntityData.Leafs["rack-name"] = types.YLeaf{"RackName", rack.RackName}
    return &(rack.EntityData)
}

// Diag_Racks_Rack_PowerShelfs
// Table for rack power shelf 
type Diag_Racks_Rack_PowerShelfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Power shelf name. The type is slice of
    // Diag_Racks_Rack_PowerShelfs_PowerShelf.
    PowerShelf []Diag_Racks_Rack_PowerShelfs_PowerShelf
}

func (powerShelfs *Diag_Racks_Rack_PowerShelfs) GetEntityData() *types.CommonEntityData {
    powerShelfs.EntityData.YFilter = powerShelfs.YFilter
    powerShelfs.EntityData.YangName = "power-shelfs"
    powerShelfs.EntityData.BundleName = "cisco_ios_xr"
    powerShelfs.EntityData.ParentYangName = "rack"
    powerShelfs.EntityData.SegmentPath = "power-shelfs"
    powerShelfs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    powerShelfs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    powerShelfs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    powerShelfs.EntityData.Children = make(map[string]types.YChild)
    powerShelfs.EntityData.Children["power-shelf"] = types.YChild{"PowerShelf", nil}
    for i := range powerShelfs.PowerShelf {
        powerShelfs.EntityData.Children[types.GetSegmentPath(&powerShelfs.PowerShelf[i])] = types.YChild{"PowerShelf", &powerShelfs.PowerShelf[i]}
    }
    powerShelfs.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(powerShelfs.EntityData)
}

// Diag_Racks_Rack_PowerShelfs_PowerShelf
// Power shelf name
type Diag_Racks_Rack_PowerShelfs_PowerShelf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Power Shelf name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    PowerShelfName interface{}

    // Table for rack power supply .
    PowerSupplies Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies
}

func (powerShelf *Diag_Racks_Rack_PowerShelfs_PowerShelf) GetEntityData() *types.CommonEntityData {
    powerShelf.EntityData.YFilter = powerShelf.YFilter
    powerShelf.EntityData.YangName = "power-shelf"
    powerShelf.EntityData.BundleName = "cisco_ios_xr"
    powerShelf.EntityData.ParentYangName = "power-shelfs"
    powerShelf.EntityData.SegmentPath = "power-shelf" + "[power-shelf-name='" + fmt.Sprintf("%v", powerShelf.PowerShelfName) + "']"
    powerShelf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    powerShelf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    powerShelf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    powerShelf.EntityData.Children = make(map[string]types.YChild)
    powerShelf.EntityData.Children["power-supplies"] = types.YChild{"PowerSupplies", &powerShelf.PowerSupplies}
    powerShelf.EntityData.Leafs = make(map[string]types.YLeaf)
    powerShelf.EntityData.Leafs["power-shelf-name"] = types.YLeaf{"PowerShelfName", powerShelf.PowerShelfName}
    return &(powerShelf.EntityData)
}

// Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies
// Table for rack power supply 
type Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Power Supply name. The type is slice of
    // Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies_PowerSupply.
    PowerSupply []Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies_PowerSupply
}

func (powerSupplies *Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies) GetEntityData() *types.CommonEntityData {
    powerSupplies.EntityData.YFilter = powerSupplies.YFilter
    powerSupplies.EntityData.YangName = "power-supplies"
    powerSupplies.EntityData.BundleName = "cisco_ios_xr"
    powerSupplies.EntityData.ParentYangName = "power-shelf"
    powerSupplies.EntityData.SegmentPath = "power-supplies"
    powerSupplies.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    powerSupplies.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    powerSupplies.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    powerSupplies.EntityData.Children = make(map[string]types.YChild)
    powerSupplies.EntityData.Children["power-supply"] = types.YChild{"PowerSupply", nil}
    for i := range powerSupplies.PowerSupply {
        powerSupplies.EntityData.Children[types.GetSegmentPath(&powerSupplies.PowerSupply[i])] = types.YChild{"PowerSupply", &powerSupplies.PowerSupply[i]}
    }
    powerSupplies.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(powerSupplies.EntityData)
}

// Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies_PowerSupply
// Power Supply name
type Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies_PowerSupply struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Power Supply name. The type is string with
    // pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    PowerSupplyName interface{}

    // Basic information.
    Information Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies_PowerSupply_Information
}

func (powerSupply *Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies_PowerSupply) GetEntityData() *types.CommonEntityData {
    powerSupply.EntityData.YFilter = powerSupply.YFilter
    powerSupply.EntityData.YangName = "power-supply"
    powerSupply.EntityData.BundleName = "cisco_ios_xr"
    powerSupply.EntityData.ParentYangName = "power-supplies"
    powerSupply.EntityData.SegmentPath = "power-supply" + "[power-supply-name='" + fmt.Sprintf("%v", powerSupply.PowerSupplyName) + "']"
    powerSupply.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    powerSupply.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    powerSupply.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    powerSupply.EntityData.Children = make(map[string]types.YChild)
    powerSupply.EntityData.Children["information"] = types.YChild{"Information", &powerSupply.Information}
    powerSupply.EntityData.Leafs = make(map[string]types.YLeaf)
    powerSupply.EntityData.Leafs["power-supply-name"] = types.YLeaf{"PowerSupplyName", powerSupply.PowerSupplyName}
    return &(powerSupply.EntityData)
}

// Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies_PowerSupply_Information
// Basic information
type Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies_PowerSupply_Information struct {
    EntityData types.CommonEntityData
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

func (information *Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies_PowerSupply_Information) GetEntityData() *types.CommonEntityData {
    information.EntityData.YFilter = information.YFilter
    information.EntityData.YangName = "information"
    information.EntityData.BundleName = "cisco_ios_xr"
    information.EntityData.ParentYangName = "power-supply"
    information.EntityData.SegmentPath = "information"
    information.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    information.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    information.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    information.EntityData.Children = make(map[string]types.YChild)
    information.EntityData.Children["rma"] = types.YChild{"Rma", &information.Rma}
    information.EntityData.Leafs = make(map[string]types.YLeaf)
    information.EntityData.Leafs["description"] = types.YLeaf{"Description", information.Description}
    information.EntityData.Leafs["idprom-format-rev"] = types.YLeaf{"IdpromFormatRev", information.IdpromFormatRev}
    information.EntityData.Leafs["controller-family"] = types.YLeaf{"ControllerFamily", information.ControllerFamily}
    information.EntityData.Leafs["controller-type"] = types.YLeaf{"ControllerType", information.ControllerType}
    information.EntityData.Leafs["vid"] = types.YLeaf{"Vid", information.Vid}
    information.EntityData.Leafs["hwid"] = types.YLeaf{"Hwid", information.Hwid}
    information.EntityData.Leafs["pid"] = types.YLeaf{"Pid", information.Pid}
    information.EntityData.Leafs["udi-description"] = types.YLeaf{"UdiDescription", information.UdiDescription}
    information.EntityData.Leafs["udi-name"] = types.YLeaf{"UdiName", information.UdiName}
    information.EntityData.Leafs["clei"] = types.YLeaf{"Clei", information.Clei}
    information.EntityData.Leafs["eci"] = types.YLeaf{"Eci", information.Eci}
    information.EntityData.Leafs["top-assem-part-num"] = types.YLeaf{"TopAssemPartNum", information.TopAssemPartNum}
    information.EntityData.Leafs["top-assem-vid"] = types.YLeaf{"TopAssemVid", information.TopAssemVid}
    information.EntityData.Leafs["pca-num"] = types.YLeaf{"PcaNum", information.PcaNum}
    information.EntityData.Leafs["pcavid"] = types.YLeaf{"Pcavid", information.Pcavid}
    information.EntityData.Leafs["chassis-sid"] = types.YLeaf{"ChassisSid", information.ChassisSid}
    information.EntityData.Leafs["dev-num1"] = types.YLeaf{"DevNum1", information.DevNum1}
    information.EntityData.Leafs["dev-num2"] = types.YLeaf{"DevNum2", information.DevNum2}
    information.EntityData.Leafs["dev-num3"] = types.YLeaf{"DevNum3", information.DevNum3}
    information.EntityData.Leafs["dev-num4"] = types.YLeaf{"DevNum4", information.DevNum4}
    information.EntityData.Leafs["dev-num5"] = types.YLeaf{"DevNum5", information.DevNum5}
    information.EntityData.Leafs["dev-num6"] = types.YLeaf{"DevNum6", information.DevNum6}
    information.EntityData.Leafs["dev-num7"] = types.YLeaf{"DevNum7", information.DevNum7}
    information.EntityData.Leafs["manu-test-data"] = types.YLeaf{"ManuTestData", information.ManuTestData}
    information.EntityData.Leafs["asset-id"] = types.YLeaf{"AssetId", information.AssetId}
    information.EntityData.Leafs["asset-alias"] = types.YLeaf{"AssetAlias", information.AssetAlias}
    information.EntityData.Leafs["base-mac-address1"] = types.YLeaf{"BaseMacAddress1", information.BaseMacAddress1}
    information.EntityData.Leafs["mac-add-blk-size1"] = types.YLeaf{"MacAddBlkSize1", information.MacAddBlkSize1}
    information.EntityData.Leafs["base-mac-address2"] = types.YLeaf{"BaseMacAddress2", information.BaseMacAddress2}
    information.EntityData.Leafs["mac-add-blk-size2"] = types.YLeaf{"MacAddBlkSize2", information.MacAddBlkSize2}
    information.EntityData.Leafs["base-mac-address3"] = types.YLeaf{"BaseMacAddress3", information.BaseMacAddress3}
    information.EntityData.Leafs["mac-add-blk-size3"] = types.YLeaf{"MacAddBlkSize3", information.MacAddBlkSize3}
    information.EntityData.Leafs["base-mac-address4"] = types.YLeaf{"BaseMacAddress4", information.BaseMacAddress4}
    information.EntityData.Leafs["mac-add-blk-size4"] = types.YLeaf{"MacAddBlkSize4", information.MacAddBlkSize4}
    information.EntityData.Leafs["pcb-serial-num"] = types.YLeaf{"PcbSerialNum", information.PcbSerialNum}
    information.EntityData.Leafs["power-supply-type"] = types.YLeaf{"PowerSupplyType", information.PowerSupplyType}
    information.EntityData.Leafs["power-consumption"] = types.YLeaf{"PowerConsumption", information.PowerConsumption}
    information.EntityData.Leafs["block-signature"] = types.YLeaf{"BlockSignature", information.BlockSignature}
    information.EntityData.Leafs["block-version"] = types.YLeaf{"BlockVersion", information.BlockVersion}
    information.EntityData.Leafs["block-length"] = types.YLeaf{"BlockLength", information.BlockLength}
    information.EntityData.Leafs["block-checksum"] = types.YLeaf{"BlockChecksum", information.BlockChecksum}
    information.EntityData.Leafs["eeprom-size"] = types.YLeaf{"EepromSize", information.EepromSize}
    information.EntityData.Leafs["block-count"] = types.YLeaf{"BlockCount", information.BlockCount}
    information.EntityData.Leafs["fru-major-type"] = types.YLeaf{"FruMajorType", information.FruMajorType}
    information.EntityData.Leafs["fru-minor-type"] = types.YLeaf{"FruMinorType", information.FruMinorType}
    information.EntityData.Leafs["oem-string"] = types.YLeaf{"OemString", information.OemString}
    information.EntityData.Leafs["product-id"] = types.YLeaf{"ProductId", information.ProductId}
    information.EntityData.Leafs["serial-number"] = types.YLeaf{"SerialNumber", information.SerialNumber}
    information.EntityData.Leafs["part-number"] = types.YLeaf{"PartNumber", information.PartNumber}
    information.EntityData.Leafs["part-revision"] = types.YLeaf{"PartRevision", information.PartRevision}
    information.EntityData.Leafs["mfg-deviation"] = types.YLeaf{"MfgDeviation", information.MfgDeviation}
    information.EntityData.Leafs["hw-version"] = types.YLeaf{"HwVersion", information.HwVersion}
    information.EntityData.Leafs["mfg-bits"] = types.YLeaf{"MfgBits", information.MfgBits}
    information.EntityData.Leafs["engineer-use"] = types.YLeaf{"EngineerUse", information.EngineerUse}
    information.EntityData.Leafs["snmpoid"] = types.YLeaf{"Snmpoid", information.Snmpoid}
    information.EntityData.Leafs["rma-code"] = types.YLeaf{"RmaCode", information.RmaCode}
    return &(information.EntityData)
}

// Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies_PowerSupply_Information_Rma
// RMA Data
type Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies_PowerSupply_Information_Rma struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Test history. The type is string with length: 0..255.
    TestHistory interface{}

    // RMA tracking number format is N-N-N. The type is string with length:
    // 0..255.
    RmaNumber interface{}

    // RMA history. The type is string with length: 0..255.
    RmaHistory interface{}
}

func (rma *Diag_Racks_Rack_PowerShelfs_PowerShelf_PowerSupplies_PowerSupply_Information_Rma) GetEntityData() *types.CommonEntityData {
    rma.EntityData.YFilter = rma.YFilter
    rma.EntityData.YangName = "rma"
    rma.EntityData.BundleName = "cisco_ios_xr"
    rma.EntityData.ParentYangName = "information"
    rma.EntityData.SegmentPath = "rma"
    rma.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rma.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rma.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rma.EntityData.Children = make(map[string]types.YChild)
    rma.EntityData.Leafs = make(map[string]types.YLeaf)
    rma.EntityData.Leafs["test-history"] = types.YLeaf{"TestHistory", rma.TestHistory}
    rma.EntityData.Leafs["rma-number"] = types.YLeaf{"RmaNumber", rma.RmaNumber}
    rma.EntityData.Leafs["rma-history"] = types.YLeaf{"RmaHistory", rma.RmaHistory}
    return &(rma.EntityData)
}

// Diag_Racks_Rack_FanTraies
// Table for rack fan trays
type Diag_Racks_Rack_FanTraies struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Fan tray name. The type is slice of Diag_Racks_Rack_FanTraies_FanTray.
    FanTray []Diag_Racks_Rack_FanTraies_FanTray
}

func (fanTraies *Diag_Racks_Rack_FanTraies) GetEntityData() *types.CommonEntityData {
    fanTraies.EntityData.YFilter = fanTraies.YFilter
    fanTraies.EntityData.YangName = "fan-traies"
    fanTraies.EntityData.BundleName = "cisco_ios_xr"
    fanTraies.EntityData.ParentYangName = "rack"
    fanTraies.EntityData.SegmentPath = "fan-traies"
    fanTraies.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fanTraies.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fanTraies.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fanTraies.EntityData.Children = make(map[string]types.YChild)
    fanTraies.EntityData.Children["fan-tray"] = types.YChild{"FanTray", nil}
    for i := range fanTraies.FanTray {
        fanTraies.EntityData.Children[types.GetSegmentPath(&fanTraies.FanTray[i])] = types.YChild{"FanTray", &fanTraies.FanTray[i]}
    }
    fanTraies.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(fanTraies.EntityData)
}

// Diag_Racks_Rack_FanTraies_FanTray
// Fan tray name
type Diag_Racks_Rack_FanTraies_FanTray struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Fan tray name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    FanTrayName interface{}

    // Table for rack fans .
    Fanses Diag_Racks_Rack_FanTraies_FanTray_Fanses
}

func (fanTray *Diag_Racks_Rack_FanTraies_FanTray) GetEntityData() *types.CommonEntityData {
    fanTray.EntityData.YFilter = fanTray.YFilter
    fanTray.EntityData.YangName = "fan-tray"
    fanTray.EntityData.BundleName = "cisco_ios_xr"
    fanTray.EntityData.ParentYangName = "fan-traies"
    fanTray.EntityData.SegmentPath = "fan-tray" + "[fan-tray-name='" + fmt.Sprintf("%v", fanTray.FanTrayName) + "']"
    fanTray.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fanTray.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fanTray.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fanTray.EntityData.Children = make(map[string]types.YChild)
    fanTray.EntityData.Children["fanses"] = types.YChild{"Fanses", &fanTray.Fanses}
    fanTray.EntityData.Leafs = make(map[string]types.YLeaf)
    fanTray.EntityData.Leafs["fan-tray-name"] = types.YLeaf{"FanTrayName", fanTray.FanTrayName}
    return &(fanTray.EntityData)
}

// Diag_Racks_Rack_FanTraies_FanTray_Fanses
// Table for rack fans 
type Diag_Racks_Rack_FanTraies_FanTray_Fanses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Fan name. The type is slice of
    // Diag_Racks_Rack_FanTraies_FanTray_Fanses_Fans.
    Fans []Diag_Racks_Rack_FanTraies_FanTray_Fanses_Fans
}

func (fanses *Diag_Racks_Rack_FanTraies_FanTray_Fanses) GetEntityData() *types.CommonEntityData {
    fanses.EntityData.YFilter = fanses.YFilter
    fanses.EntityData.YangName = "fanses"
    fanses.EntityData.BundleName = "cisco_ios_xr"
    fanses.EntityData.ParentYangName = "fan-tray"
    fanses.EntityData.SegmentPath = "fanses"
    fanses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fanses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fanses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fanses.EntityData.Children = make(map[string]types.YChild)
    fanses.EntityData.Children["fans"] = types.YChild{"Fans", nil}
    for i := range fanses.Fans {
        fanses.EntityData.Children[types.GetSegmentPath(&fanses.Fans[i])] = types.YChild{"Fans", &fanses.Fans[i]}
    }
    fanses.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(fanses.EntityData)
}

// Diag_Racks_Rack_FanTraies_FanTray_Fanses_Fans
// Fan name
type Diag_Racks_Rack_FanTraies_FanTray_Fanses_Fans struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Fans name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    FansName interface{}

    // Basic information.
    Information Diag_Racks_Rack_FanTraies_FanTray_Fanses_Fans_Information
}

func (fans *Diag_Racks_Rack_FanTraies_FanTray_Fanses_Fans) GetEntityData() *types.CommonEntityData {
    fans.EntityData.YFilter = fans.YFilter
    fans.EntityData.YangName = "fans"
    fans.EntityData.BundleName = "cisco_ios_xr"
    fans.EntityData.ParentYangName = "fanses"
    fans.EntityData.SegmentPath = "fans" + "[fans-name='" + fmt.Sprintf("%v", fans.FansName) + "']"
    fans.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fans.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fans.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fans.EntityData.Children = make(map[string]types.YChild)
    fans.EntityData.Children["information"] = types.YChild{"Information", &fans.Information}
    fans.EntityData.Leafs = make(map[string]types.YLeaf)
    fans.EntityData.Leafs["fans-name"] = types.YLeaf{"FansName", fans.FansName}
    return &(fans.EntityData)
}

// Diag_Racks_Rack_FanTraies_FanTray_Fanses_Fans_Information
// Basic information
type Diag_Racks_Rack_FanTraies_FanTray_Fanses_Fans_Information struct {
    EntityData types.CommonEntityData
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

func (information *Diag_Racks_Rack_FanTraies_FanTray_Fanses_Fans_Information) GetEntityData() *types.CommonEntityData {
    information.EntityData.YFilter = information.YFilter
    information.EntityData.YangName = "information"
    information.EntityData.BundleName = "cisco_ios_xr"
    information.EntityData.ParentYangName = "fans"
    information.EntityData.SegmentPath = "information"
    information.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    information.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    information.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    information.EntityData.Children = make(map[string]types.YChild)
    information.EntityData.Children["rma"] = types.YChild{"Rma", &information.Rma}
    information.EntityData.Leafs = make(map[string]types.YLeaf)
    information.EntityData.Leafs["description"] = types.YLeaf{"Description", information.Description}
    information.EntityData.Leafs["idprom-format-rev"] = types.YLeaf{"IdpromFormatRev", information.IdpromFormatRev}
    information.EntityData.Leafs["controller-family"] = types.YLeaf{"ControllerFamily", information.ControllerFamily}
    information.EntityData.Leafs["controller-type"] = types.YLeaf{"ControllerType", information.ControllerType}
    information.EntityData.Leafs["vid"] = types.YLeaf{"Vid", information.Vid}
    information.EntityData.Leafs["hwid"] = types.YLeaf{"Hwid", information.Hwid}
    information.EntityData.Leafs["pid"] = types.YLeaf{"Pid", information.Pid}
    information.EntityData.Leafs["udi-description"] = types.YLeaf{"UdiDescription", information.UdiDescription}
    information.EntityData.Leafs["udi-name"] = types.YLeaf{"UdiName", information.UdiName}
    information.EntityData.Leafs["clei"] = types.YLeaf{"Clei", information.Clei}
    information.EntityData.Leafs["eci"] = types.YLeaf{"Eci", information.Eci}
    information.EntityData.Leafs["top-assem-part-num"] = types.YLeaf{"TopAssemPartNum", information.TopAssemPartNum}
    information.EntityData.Leafs["top-assem-vid"] = types.YLeaf{"TopAssemVid", information.TopAssemVid}
    information.EntityData.Leafs["pca-num"] = types.YLeaf{"PcaNum", information.PcaNum}
    information.EntityData.Leafs["pcavid"] = types.YLeaf{"Pcavid", information.Pcavid}
    information.EntityData.Leafs["chassis-sid"] = types.YLeaf{"ChassisSid", information.ChassisSid}
    information.EntityData.Leafs["dev-num1"] = types.YLeaf{"DevNum1", information.DevNum1}
    information.EntityData.Leafs["dev-num2"] = types.YLeaf{"DevNum2", information.DevNum2}
    information.EntityData.Leafs["dev-num3"] = types.YLeaf{"DevNum3", information.DevNum3}
    information.EntityData.Leafs["dev-num4"] = types.YLeaf{"DevNum4", information.DevNum4}
    information.EntityData.Leafs["dev-num5"] = types.YLeaf{"DevNum5", information.DevNum5}
    information.EntityData.Leafs["dev-num6"] = types.YLeaf{"DevNum6", information.DevNum6}
    information.EntityData.Leafs["dev-num7"] = types.YLeaf{"DevNum7", information.DevNum7}
    information.EntityData.Leafs["manu-test-data"] = types.YLeaf{"ManuTestData", information.ManuTestData}
    information.EntityData.Leafs["asset-id"] = types.YLeaf{"AssetId", information.AssetId}
    information.EntityData.Leafs["asset-alias"] = types.YLeaf{"AssetAlias", information.AssetAlias}
    information.EntityData.Leafs["base-mac-address1"] = types.YLeaf{"BaseMacAddress1", information.BaseMacAddress1}
    information.EntityData.Leafs["mac-add-blk-size1"] = types.YLeaf{"MacAddBlkSize1", information.MacAddBlkSize1}
    information.EntityData.Leafs["base-mac-address2"] = types.YLeaf{"BaseMacAddress2", information.BaseMacAddress2}
    information.EntityData.Leafs["mac-add-blk-size2"] = types.YLeaf{"MacAddBlkSize2", information.MacAddBlkSize2}
    information.EntityData.Leafs["base-mac-address3"] = types.YLeaf{"BaseMacAddress3", information.BaseMacAddress3}
    information.EntityData.Leafs["mac-add-blk-size3"] = types.YLeaf{"MacAddBlkSize3", information.MacAddBlkSize3}
    information.EntityData.Leafs["base-mac-address4"] = types.YLeaf{"BaseMacAddress4", information.BaseMacAddress4}
    information.EntityData.Leafs["mac-add-blk-size4"] = types.YLeaf{"MacAddBlkSize4", information.MacAddBlkSize4}
    information.EntityData.Leafs["pcb-serial-num"] = types.YLeaf{"PcbSerialNum", information.PcbSerialNum}
    information.EntityData.Leafs["power-supply-type"] = types.YLeaf{"PowerSupplyType", information.PowerSupplyType}
    information.EntityData.Leafs["power-consumption"] = types.YLeaf{"PowerConsumption", information.PowerConsumption}
    information.EntityData.Leafs["block-signature"] = types.YLeaf{"BlockSignature", information.BlockSignature}
    information.EntityData.Leafs["block-version"] = types.YLeaf{"BlockVersion", information.BlockVersion}
    information.EntityData.Leafs["block-length"] = types.YLeaf{"BlockLength", information.BlockLength}
    information.EntityData.Leafs["block-checksum"] = types.YLeaf{"BlockChecksum", information.BlockChecksum}
    information.EntityData.Leafs["eeprom-size"] = types.YLeaf{"EepromSize", information.EepromSize}
    information.EntityData.Leafs["block-count"] = types.YLeaf{"BlockCount", information.BlockCount}
    information.EntityData.Leafs["fru-major-type"] = types.YLeaf{"FruMajorType", information.FruMajorType}
    information.EntityData.Leafs["fru-minor-type"] = types.YLeaf{"FruMinorType", information.FruMinorType}
    information.EntityData.Leafs["oem-string"] = types.YLeaf{"OemString", information.OemString}
    information.EntityData.Leafs["product-id"] = types.YLeaf{"ProductId", information.ProductId}
    information.EntityData.Leafs["serial-number"] = types.YLeaf{"SerialNumber", information.SerialNumber}
    information.EntityData.Leafs["part-number"] = types.YLeaf{"PartNumber", information.PartNumber}
    information.EntityData.Leafs["part-revision"] = types.YLeaf{"PartRevision", information.PartRevision}
    information.EntityData.Leafs["mfg-deviation"] = types.YLeaf{"MfgDeviation", information.MfgDeviation}
    information.EntityData.Leafs["hw-version"] = types.YLeaf{"HwVersion", information.HwVersion}
    information.EntityData.Leafs["mfg-bits"] = types.YLeaf{"MfgBits", information.MfgBits}
    information.EntityData.Leafs["engineer-use"] = types.YLeaf{"EngineerUse", information.EngineerUse}
    information.EntityData.Leafs["snmpoid"] = types.YLeaf{"Snmpoid", information.Snmpoid}
    information.EntityData.Leafs["rma-code"] = types.YLeaf{"RmaCode", information.RmaCode}
    return &(information.EntityData)
}

// Diag_Racks_Rack_FanTraies_FanTray_Fanses_Fans_Information_Rma
// RMA Data
type Diag_Racks_Rack_FanTraies_FanTray_Fanses_Fans_Information_Rma struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Test history. The type is string with length: 0..255.
    TestHistory interface{}

    // RMA tracking number format is N-N-N. The type is string with length:
    // 0..255.
    RmaNumber interface{}

    // RMA history. The type is string with length: 0..255.
    RmaHistory interface{}
}

func (rma *Diag_Racks_Rack_FanTraies_FanTray_Fanses_Fans_Information_Rma) GetEntityData() *types.CommonEntityData {
    rma.EntityData.YFilter = rma.YFilter
    rma.EntityData.YangName = "rma"
    rma.EntityData.BundleName = "cisco_ios_xr"
    rma.EntityData.ParentYangName = "information"
    rma.EntityData.SegmentPath = "rma"
    rma.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rma.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rma.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rma.EntityData.Children = make(map[string]types.YChild)
    rma.EntityData.Leafs = make(map[string]types.YLeaf)
    rma.EntityData.Leafs["test-history"] = types.YLeaf{"TestHistory", rma.TestHistory}
    rma.EntityData.Leafs["rma-number"] = types.YLeaf{"RmaNumber", rma.RmaNumber}
    rma.EntityData.Leafs["rma-history"] = types.YLeaf{"RmaHistory", rma.RmaHistory}
    return &(rma.EntityData)
}

// Diag_Racks_Rack_Slots
// Table of slots
type Diag_Racks_Rack_Slots struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Slot name. The type is slice of Diag_Racks_Rack_Slots_Slot.
    Slot []Diag_Racks_Rack_Slots_Slot
}

func (slots *Diag_Racks_Rack_Slots) GetEntityData() *types.CommonEntityData {
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

// Diag_Racks_Rack_Slots_Slot
// Slot name
type Diag_Racks_Rack_Slots_Slot struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Slot name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    SlotName interface{}

    // Table of instances.
    Instances Diag_Racks_Rack_Slots_Slot_Instances
}

func (slot *Diag_Racks_Rack_Slots_Slot) GetEntityData() *types.CommonEntityData {
    slot.EntityData.YFilter = slot.YFilter
    slot.EntityData.YangName = "slot"
    slot.EntityData.BundleName = "cisco_ios_xr"
    slot.EntityData.ParentYangName = "slots"
    slot.EntityData.SegmentPath = "slot" + "[slot-name='" + fmt.Sprintf("%v", slot.SlotName) + "']"
    slot.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    slot.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    slot.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    slot.EntityData.Children = make(map[string]types.YChild)
    slot.EntityData.Children["instances"] = types.YChild{"Instances", &slot.Instances}
    slot.EntityData.Leafs = make(map[string]types.YLeaf)
    slot.EntityData.Leafs["slot-name"] = types.YLeaf{"SlotName", slot.SlotName}
    return &(slot.EntityData)
}

// Diag_Racks_Rack_Slots_Slot_Instances
// Table of instances
type Diag_Racks_Rack_Slots_Slot_Instances struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // instance number. The type is slice of
    // Diag_Racks_Rack_Slots_Slot_Instances_Instance.
    Instance []Diag_Racks_Rack_Slots_Slot_Instances_Instance
}

func (instances *Diag_Racks_Rack_Slots_Slot_Instances) GetEntityData() *types.CommonEntityData {
    instances.EntityData.YFilter = instances.YFilter
    instances.EntityData.YangName = "instances"
    instances.EntityData.BundleName = "cisco_ios_xr"
    instances.EntityData.ParentYangName = "slot"
    instances.EntityData.SegmentPath = "instances"
    instances.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    instances.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    instances.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    instances.EntityData.Children = make(map[string]types.YChild)
    instances.EntityData.Children["instance"] = types.YChild{"Instance", nil}
    for i := range instances.Instance {
        instances.EntityData.Children[types.GetSegmentPath(&instances.Instance[i])] = types.YChild{"Instance", &instances.Instance[i]}
    }
    instances.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(instances.EntityData)
}

// Diag_Racks_Rack_Slots_Slot_Instances_Instance
// instance number
type Diag_Racks_Rack_Slots_Slot_Instances_Instance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Instance name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Name interface{}

    // Detail information.
    Detail Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail
}

func (instance *Diag_Racks_Rack_Slots_Slot_Instances_Instance) GetEntityData() *types.CommonEntityData {
    instance.EntityData.YFilter = instance.YFilter
    instance.EntityData.YangName = "instance"
    instance.EntityData.BundleName = "cisco_ios_xr"
    instance.EntityData.ParentYangName = "instances"
    instance.EntityData.SegmentPath = "instance" + "[name='" + fmt.Sprintf("%v", instance.Name) + "']"
    instance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    instance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    instance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    instance.EntityData.Children = make(map[string]types.YChild)
    instance.EntityData.Children["detail"] = types.YChild{"Detail", &instance.Detail}
    instance.EntityData.Leafs = make(map[string]types.YLeaf)
    instance.EntityData.Leafs["name"] = types.YLeaf{"Name", instance.Name}
    return &(instance.EntityData)
}

// Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail
// Detail information
type Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node operational state . The type is string with length: 0..255.
    NodeOperationalState interface{}

    // Card instance.
    CardInstance Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_CardInstance
}

func (detail *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail) GetEntityData() *types.CommonEntityData {
    detail.EntityData.YFilter = detail.YFilter
    detail.EntityData.YangName = "detail"
    detail.EntityData.BundleName = "cisco_ios_xr"
    detail.EntityData.ParentYangName = "instance"
    detail.EntityData.SegmentPath = "detail"
    detail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detail.EntityData.Children = make(map[string]types.YChild)
    detail.EntityData.Children["card-instance"] = types.YChild{"CardInstance", &detail.CardInstance}
    detail.EntityData.Leafs = make(map[string]types.YLeaf)
    detail.EntityData.Leafs["node-operational-state"] = types.YLeaf{"NodeOperationalState", detail.NodeOperationalState}
    return &(detail.EntityData)
}

// Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_CardInstance
// Card instance
type Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_CardInstance struct {
    EntityData types.CommonEntityData
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

func (cardInstance *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_CardInstance) GetEntityData() *types.CommonEntityData {
    cardInstance.EntityData.YFilter = cardInstance.YFilter
    cardInstance.EntityData.YangName = "card-instance"
    cardInstance.EntityData.BundleName = "cisco_ios_xr"
    cardInstance.EntityData.ParentYangName = "detail"
    cardInstance.EntityData.SegmentPath = "card-instance"
    cardInstance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cardInstance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cardInstance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cardInstance.EntityData.Children = make(map[string]types.YChild)
    cardInstance.EntityData.Children["rma"] = types.YChild{"Rma", &cardInstance.Rma}
    cardInstance.EntityData.Leafs = make(map[string]types.YLeaf)
    cardInstance.EntityData.Leafs["description"] = types.YLeaf{"Description", cardInstance.Description}
    cardInstance.EntityData.Leafs["idprom-format-rev"] = types.YLeaf{"IdpromFormatRev", cardInstance.IdpromFormatRev}
    cardInstance.EntityData.Leafs["controller-family"] = types.YLeaf{"ControllerFamily", cardInstance.ControllerFamily}
    cardInstance.EntityData.Leafs["controller-type"] = types.YLeaf{"ControllerType", cardInstance.ControllerType}
    cardInstance.EntityData.Leafs["vid"] = types.YLeaf{"Vid", cardInstance.Vid}
    cardInstance.EntityData.Leafs["hwid"] = types.YLeaf{"Hwid", cardInstance.Hwid}
    cardInstance.EntityData.Leafs["pid"] = types.YLeaf{"Pid", cardInstance.Pid}
    cardInstance.EntityData.Leafs["udi-description"] = types.YLeaf{"UdiDescription", cardInstance.UdiDescription}
    cardInstance.EntityData.Leafs["udi-name"] = types.YLeaf{"UdiName", cardInstance.UdiName}
    cardInstance.EntityData.Leafs["clei"] = types.YLeaf{"Clei", cardInstance.Clei}
    cardInstance.EntityData.Leafs["eci"] = types.YLeaf{"Eci", cardInstance.Eci}
    cardInstance.EntityData.Leafs["top-assem-part-num"] = types.YLeaf{"TopAssemPartNum", cardInstance.TopAssemPartNum}
    cardInstance.EntityData.Leafs["top-assem-vid"] = types.YLeaf{"TopAssemVid", cardInstance.TopAssemVid}
    cardInstance.EntityData.Leafs["pca-num"] = types.YLeaf{"PcaNum", cardInstance.PcaNum}
    cardInstance.EntityData.Leafs["pcavid"] = types.YLeaf{"Pcavid", cardInstance.Pcavid}
    cardInstance.EntityData.Leafs["chassis-sid"] = types.YLeaf{"ChassisSid", cardInstance.ChassisSid}
    cardInstance.EntityData.Leafs["dev-num1"] = types.YLeaf{"DevNum1", cardInstance.DevNum1}
    cardInstance.EntityData.Leafs["dev-num2"] = types.YLeaf{"DevNum2", cardInstance.DevNum2}
    cardInstance.EntityData.Leafs["dev-num3"] = types.YLeaf{"DevNum3", cardInstance.DevNum3}
    cardInstance.EntityData.Leafs["dev-num4"] = types.YLeaf{"DevNum4", cardInstance.DevNum4}
    cardInstance.EntityData.Leafs["dev-num5"] = types.YLeaf{"DevNum5", cardInstance.DevNum5}
    cardInstance.EntityData.Leafs["dev-num6"] = types.YLeaf{"DevNum6", cardInstance.DevNum6}
    cardInstance.EntityData.Leafs["dev-num7"] = types.YLeaf{"DevNum7", cardInstance.DevNum7}
    cardInstance.EntityData.Leafs["manu-test-data"] = types.YLeaf{"ManuTestData", cardInstance.ManuTestData}
    cardInstance.EntityData.Leafs["asset-id"] = types.YLeaf{"AssetId", cardInstance.AssetId}
    cardInstance.EntityData.Leafs["asset-alias"] = types.YLeaf{"AssetAlias", cardInstance.AssetAlias}
    cardInstance.EntityData.Leafs["base-mac-address1"] = types.YLeaf{"BaseMacAddress1", cardInstance.BaseMacAddress1}
    cardInstance.EntityData.Leafs["mac-add-blk-size1"] = types.YLeaf{"MacAddBlkSize1", cardInstance.MacAddBlkSize1}
    cardInstance.EntityData.Leafs["base-mac-address2"] = types.YLeaf{"BaseMacAddress2", cardInstance.BaseMacAddress2}
    cardInstance.EntityData.Leafs["mac-add-blk-size2"] = types.YLeaf{"MacAddBlkSize2", cardInstance.MacAddBlkSize2}
    cardInstance.EntityData.Leafs["base-mac-address3"] = types.YLeaf{"BaseMacAddress3", cardInstance.BaseMacAddress3}
    cardInstance.EntityData.Leafs["mac-add-blk-size3"] = types.YLeaf{"MacAddBlkSize3", cardInstance.MacAddBlkSize3}
    cardInstance.EntityData.Leafs["base-mac-address4"] = types.YLeaf{"BaseMacAddress4", cardInstance.BaseMacAddress4}
    cardInstance.EntityData.Leafs["mac-add-blk-size4"] = types.YLeaf{"MacAddBlkSize4", cardInstance.MacAddBlkSize4}
    cardInstance.EntityData.Leafs["pcb-serial-num"] = types.YLeaf{"PcbSerialNum", cardInstance.PcbSerialNum}
    cardInstance.EntityData.Leafs["power-supply-type"] = types.YLeaf{"PowerSupplyType", cardInstance.PowerSupplyType}
    cardInstance.EntityData.Leafs["power-consumption"] = types.YLeaf{"PowerConsumption", cardInstance.PowerConsumption}
    cardInstance.EntityData.Leafs["block-signature"] = types.YLeaf{"BlockSignature", cardInstance.BlockSignature}
    cardInstance.EntityData.Leafs["block-version"] = types.YLeaf{"BlockVersion", cardInstance.BlockVersion}
    cardInstance.EntityData.Leafs["block-length"] = types.YLeaf{"BlockLength", cardInstance.BlockLength}
    cardInstance.EntityData.Leafs["block-checksum"] = types.YLeaf{"BlockChecksum", cardInstance.BlockChecksum}
    cardInstance.EntityData.Leafs["eeprom-size"] = types.YLeaf{"EepromSize", cardInstance.EepromSize}
    cardInstance.EntityData.Leafs["block-count"] = types.YLeaf{"BlockCount", cardInstance.BlockCount}
    cardInstance.EntityData.Leafs["fru-major-type"] = types.YLeaf{"FruMajorType", cardInstance.FruMajorType}
    cardInstance.EntityData.Leafs["fru-minor-type"] = types.YLeaf{"FruMinorType", cardInstance.FruMinorType}
    cardInstance.EntityData.Leafs["oem-string"] = types.YLeaf{"OemString", cardInstance.OemString}
    cardInstance.EntityData.Leafs["product-id"] = types.YLeaf{"ProductId", cardInstance.ProductId}
    cardInstance.EntityData.Leafs["serial-number"] = types.YLeaf{"SerialNumber", cardInstance.SerialNumber}
    cardInstance.EntityData.Leafs["part-number"] = types.YLeaf{"PartNumber", cardInstance.PartNumber}
    cardInstance.EntityData.Leafs["part-revision"] = types.YLeaf{"PartRevision", cardInstance.PartRevision}
    cardInstance.EntityData.Leafs["mfg-deviation"] = types.YLeaf{"MfgDeviation", cardInstance.MfgDeviation}
    cardInstance.EntityData.Leafs["hw-version"] = types.YLeaf{"HwVersion", cardInstance.HwVersion}
    cardInstance.EntityData.Leafs["mfg-bits"] = types.YLeaf{"MfgBits", cardInstance.MfgBits}
    cardInstance.EntityData.Leafs["engineer-use"] = types.YLeaf{"EngineerUse", cardInstance.EngineerUse}
    cardInstance.EntityData.Leafs["snmpoid"] = types.YLeaf{"Snmpoid", cardInstance.Snmpoid}
    cardInstance.EntityData.Leafs["rma-code"] = types.YLeaf{"RmaCode", cardInstance.RmaCode}
    return &(cardInstance.EntityData)
}

// Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_CardInstance_Rma
// RMA Data
type Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_CardInstance_Rma struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Test history. The type is string with length: 0..255.
    TestHistory interface{}

    // RMA tracking number format is N-N-N. The type is string with length:
    // 0..255.
    RmaNumber interface{}

    // RMA history. The type is string with length: 0..255.
    RmaHistory interface{}
}

func (rma *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_CardInstance_Rma) GetEntityData() *types.CommonEntityData {
    rma.EntityData.YFilter = rma.YFilter
    rma.EntityData.YangName = "rma"
    rma.EntityData.BundleName = "cisco_ios_xr"
    rma.EntityData.ParentYangName = "card-instance"
    rma.EntityData.SegmentPath = "rma"
    rma.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rma.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rma.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rma.EntityData.Children = make(map[string]types.YChild)
    rma.EntityData.Leafs = make(map[string]types.YLeaf)
    rma.EntityData.Leafs["test-history"] = types.YLeaf{"TestHistory", rma.TestHistory}
    rma.EntityData.Leafs["rma-number"] = types.YLeaf{"RmaNumber", rma.RmaNumber}
    rma.EntityData.Leafs["rma-history"] = types.YLeaf{"RmaHistory", rma.RmaHistory}
    return &(rma.EntityData)
}

// Diag_Racks_Rack_Chassis
// Chassis information
type Diag_Racks_Rack_Chassis struct {
    EntityData types.CommonEntityData
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

func (chassis *Diag_Racks_Rack_Chassis) GetEntityData() *types.CommonEntityData {
    chassis.EntityData.YFilter = chassis.YFilter
    chassis.EntityData.YangName = "chassis"
    chassis.EntityData.BundleName = "cisco_ios_xr"
    chassis.EntityData.ParentYangName = "rack"
    chassis.EntityData.SegmentPath = "chassis"
    chassis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    chassis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    chassis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    chassis.EntityData.Children = make(map[string]types.YChild)
    chassis.EntityData.Children["rma"] = types.YChild{"Rma", &chassis.Rma}
    chassis.EntityData.Leafs = make(map[string]types.YLeaf)
    chassis.EntityData.Leafs["description"] = types.YLeaf{"Description", chassis.Description}
    chassis.EntityData.Leafs["idprom-format-rev"] = types.YLeaf{"IdpromFormatRev", chassis.IdpromFormatRev}
    chassis.EntityData.Leafs["controller-family"] = types.YLeaf{"ControllerFamily", chassis.ControllerFamily}
    chassis.EntityData.Leafs["controller-type"] = types.YLeaf{"ControllerType", chassis.ControllerType}
    chassis.EntityData.Leafs["vid"] = types.YLeaf{"Vid", chassis.Vid}
    chassis.EntityData.Leafs["hwid"] = types.YLeaf{"Hwid", chassis.Hwid}
    chassis.EntityData.Leafs["pid"] = types.YLeaf{"Pid", chassis.Pid}
    chassis.EntityData.Leafs["udi-description"] = types.YLeaf{"UdiDescription", chassis.UdiDescription}
    chassis.EntityData.Leafs["udi-name"] = types.YLeaf{"UdiName", chassis.UdiName}
    chassis.EntityData.Leafs["clei"] = types.YLeaf{"Clei", chassis.Clei}
    chassis.EntityData.Leafs["eci"] = types.YLeaf{"Eci", chassis.Eci}
    chassis.EntityData.Leafs["top-assem-part-num"] = types.YLeaf{"TopAssemPartNum", chassis.TopAssemPartNum}
    chassis.EntityData.Leafs["top-assem-vid"] = types.YLeaf{"TopAssemVid", chassis.TopAssemVid}
    chassis.EntityData.Leafs["pca-num"] = types.YLeaf{"PcaNum", chassis.PcaNum}
    chassis.EntityData.Leafs["pcavid"] = types.YLeaf{"Pcavid", chassis.Pcavid}
    chassis.EntityData.Leafs["chassis-sid"] = types.YLeaf{"ChassisSid", chassis.ChassisSid}
    chassis.EntityData.Leafs["dev-num1"] = types.YLeaf{"DevNum1", chassis.DevNum1}
    chassis.EntityData.Leafs["dev-num2"] = types.YLeaf{"DevNum2", chassis.DevNum2}
    chassis.EntityData.Leafs["dev-num3"] = types.YLeaf{"DevNum3", chassis.DevNum3}
    chassis.EntityData.Leafs["dev-num4"] = types.YLeaf{"DevNum4", chassis.DevNum4}
    chassis.EntityData.Leafs["dev-num5"] = types.YLeaf{"DevNum5", chassis.DevNum5}
    chassis.EntityData.Leafs["dev-num6"] = types.YLeaf{"DevNum6", chassis.DevNum6}
    chassis.EntityData.Leafs["dev-num7"] = types.YLeaf{"DevNum7", chassis.DevNum7}
    chassis.EntityData.Leafs["manu-test-data"] = types.YLeaf{"ManuTestData", chassis.ManuTestData}
    chassis.EntityData.Leafs["asset-id"] = types.YLeaf{"AssetId", chassis.AssetId}
    chassis.EntityData.Leafs["asset-alias"] = types.YLeaf{"AssetAlias", chassis.AssetAlias}
    chassis.EntityData.Leafs["base-mac-address1"] = types.YLeaf{"BaseMacAddress1", chassis.BaseMacAddress1}
    chassis.EntityData.Leafs["mac-add-blk-size1"] = types.YLeaf{"MacAddBlkSize1", chassis.MacAddBlkSize1}
    chassis.EntityData.Leafs["base-mac-address2"] = types.YLeaf{"BaseMacAddress2", chassis.BaseMacAddress2}
    chassis.EntityData.Leafs["mac-add-blk-size2"] = types.YLeaf{"MacAddBlkSize2", chassis.MacAddBlkSize2}
    chassis.EntityData.Leafs["base-mac-address3"] = types.YLeaf{"BaseMacAddress3", chassis.BaseMacAddress3}
    chassis.EntityData.Leafs["mac-add-blk-size3"] = types.YLeaf{"MacAddBlkSize3", chassis.MacAddBlkSize3}
    chassis.EntityData.Leafs["base-mac-address4"] = types.YLeaf{"BaseMacAddress4", chassis.BaseMacAddress4}
    chassis.EntityData.Leafs["mac-add-blk-size4"] = types.YLeaf{"MacAddBlkSize4", chassis.MacAddBlkSize4}
    chassis.EntityData.Leafs["pcb-serial-num"] = types.YLeaf{"PcbSerialNum", chassis.PcbSerialNum}
    chassis.EntityData.Leafs["power-supply-type"] = types.YLeaf{"PowerSupplyType", chassis.PowerSupplyType}
    chassis.EntityData.Leafs["power-consumption"] = types.YLeaf{"PowerConsumption", chassis.PowerConsumption}
    chassis.EntityData.Leafs["block-signature"] = types.YLeaf{"BlockSignature", chassis.BlockSignature}
    chassis.EntityData.Leafs["block-version"] = types.YLeaf{"BlockVersion", chassis.BlockVersion}
    chassis.EntityData.Leafs["block-length"] = types.YLeaf{"BlockLength", chassis.BlockLength}
    chassis.EntityData.Leafs["block-checksum"] = types.YLeaf{"BlockChecksum", chassis.BlockChecksum}
    chassis.EntityData.Leafs["eeprom-size"] = types.YLeaf{"EepromSize", chassis.EepromSize}
    chassis.EntityData.Leafs["block-count"] = types.YLeaf{"BlockCount", chassis.BlockCount}
    chassis.EntityData.Leafs["fru-major-type"] = types.YLeaf{"FruMajorType", chassis.FruMajorType}
    chassis.EntityData.Leafs["fru-minor-type"] = types.YLeaf{"FruMinorType", chassis.FruMinorType}
    chassis.EntityData.Leafs["oem-string"] = types.YLeaf{"OemString", chassis.OemString}
    chassis.EntityData.Leafs["product-id"] = types.YLeaf{"ProductId", chassis.ProductId}
    chassis.EntityData.Leafs["serial-number"] = types.YLeaf{"SerialNumber", chassis.SerialNumber}
    chassis.EntityData.Leafs["part-number"] = types.YLeaf{"PartNumber", chassis.PartNumber}
    chassis.EntityData.Leafs["part-revision"] = types.YLeaf{"PartRevision", chassis.PartRevision}
    chassis.EntityData.Leafs["mfg-deviation"] = types.YLeaf{"MfgDeviation", chassis.MfgDeviation}
    chassis.EntityData.Leafs["hw-version"] = types.YLeaf{"HwVersion", chassis.HwVersion}
    chassis.EntityData.Leafs["mfg-bits"] = types.YLeaf{"MfgBits", chassis.MfgBits}
    chassis.EntityData.Leafs["engineer-use"] = types.YLeaf{"EngineerUse", chassis.EngineerUse}
    chassis.EntityData.Leafs["snmpoid"] = types.YLeaf{"Snmpoid", chassis.Snmpoid}
    chassis.EntityData.Leafs["rma-code"] = types.YLeaf{"RmaCode", chassis.RmaCode}
    return &(chassis.EntityData)
}

// Diag_Racks_Rack_Chassis_Rma
// RMA Data
type Diag_Racks_Rack_Chassis_Rma struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Test history. The type is string with length: 0..255.
    TestHistory interface{}

    // RMA tracking number format is N-N-N. The type is string with length:
    // 0..255.
    RmaNumber interface{}

    // RMA history. The type is string with length: 0..255.
    RmaHistory interface{}
}

func (rma *Diag_Racks_Rack_Chassis_Rma) GetEntityData() *types.CommonEntityData {
    rma.EntityData.YFilter = rma.YFilter
    rma.EntityData.YangName = "rma"
    rma.EntityData.BundleName = "cisco_ios_xr"
    rma.EntityData.ParentYangName = "chassis"
    rma.EntityData.SegmentPath = "rma"
    rma.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rma.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rma.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rma.EntityData.Children = make(map[string]types.YChild)
    rma.EntityData.Leafs = make(map[string]types.YLeaf)
    rma.EntityData.Leafs["test-history"] = types.YLeaf{"TestHistory", rma.TestHistory}
    rma.EntityData.Leafs["rma-number"] = types.YLeaf{"RmaNumber", rma.RmaNumber}
    rma.EntityData.Leafs["rma-history"] = types.YLeaf{"RmaHistory", rma.RmaHistory}
    return &(rma.EntityData)
}

