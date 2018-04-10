// This module contains definitions
// for the Calvados model objects.
// 
// This module contains a collection of YANG
// definitions for Cisco IOS-XR SysAdmin configuration.
// 
// This module holds diag data for chassis, card, fan, power.
// 
// Copyright(c) 2012-2017 by Cisco Systems, Inc.
// All rights reserved.
// 
// Copyright (c) 2012-2017 by Cisco Systems, Inc.
// All rights reserved.
package sysadmin_show_diag

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package sysadmin_show_diag"))
    ydk.RegisterEntity("{http://www.cisco.com/ns/yang/Cisco-IOS-XR-sysadmin-show-diag diag}", reflect.TypeOf(Diag{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sysadmin-show-diag:diag", reflect.TypeOf(Diag{}))
}

// Diag
// diag data
type Diag struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Default_ Diag_Default

    
    Fans Diag_Fans

    
    PowerSupply Diag_PowerSupply

    
    Chassis Diag_Chassis

    
    Summary Diag_Summary

    
    Eeprom Diag_Eeprom

    
    Detail Diag_Detail
}

func (diag *Diag) GetEntityData() *types.CommonEntityData {
    diag.EntityData.YFilter = diag.YFilter
    diag.EntityData.YangName = "diag"
    diag.EntityData.BundleName = "cisco_ios_xr"
    diag.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-show-diag"
    diag.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-show-diag:diag"
    diag.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    diag.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    diag.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    diag.EntityData.Children = make(map[string]types.YChild)
    diag.EntityData.Children["default"] = types.YChild{"Default_", &diag.Default_}
    diag.EntityData.Children["fans"] = types.YChild{"Fans", &diag.Fans}
    diag.EntityData.Children["power-supply"] = types.YChild{"PowerSupply", &diag.PowerSupply}
    diag.EntityData.Children["chassis"] = types.YChild{"Chassis", &diag.Chassis}
    diag.EntityData.Children["summary"] = types.YChild{"Summary", &diag.Summary}
    diag.EntityData.Children["eeprom"] = types.YChild{"Eeprom", &diag.Eeprom}
    diag.EntityData.Children["detail"] = types.YChild{"Detail", &diag.Detail}
    diag.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(diag.EntityData)
}

// Diag_Default
type Diag_Default struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Diag_Default_DefaultList.
    DefaultList []Diag_Default_DefaultList
}

func (self *Diag_Default) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "default"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "diag"
    self.EntityData.SegmentPath = "default"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Children["default_list"] = types.YChild{"DefaultList", nil}
    for i := range self.DefaultList {
        self.EntityData.Children[types.GetSegmentPath(&self.DefaultList[i])] = types.YChild{"DefaultList", &self.DefaultList[i]}
    }
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(self.EntityData)
}

// Diag_Default_DefaultList
type Diag_Default_DefaultList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Location interface{}

    
    DefaultData Diag_Default_DefaultList_DefaultData
}

func (defaultList *Diag_Default_DefaultList) GetEntityData() *types.CommonEntityData {
    defaultList.EntityData.YFilter = defaultList.YFilter
    defaultList.EntityData.YangName = "default_list"
    defaultList.EntityData.BundleName = "cisco_ios_xr"
    defaultList.EntityData.ParentYangName = "default"
    defaultList.EntityData.SegmentPath = "default_list" + "[location='" + fmt.Sprintf("%v", defaultList.Location) + "']"
    defaultList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    defaultList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    defaultList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    defaultList.EntityData.Children = make(map[string]types.YChild)
    defaultList.EntityData.Children["default-data"] = types.YChild{"DefaultData", &defaultList.DefaultData}
    defaultList.EntityData.Leafs = make(map[string]types.YLeaf)
    defaultList.EntityData.Leafs["location"] = types.YLeaf{"Location", defaultList.Location}
    return &(defaultList.EntityData)
}

// Diag_Default_DefaultList_DefaultData
type Diag_Default_DefaultList_DefaultData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of string.
    DefaultOutList []interface{}
}

func (defaultData *Diag_Default_DefaultList_DefaultData) GetEntityData() *types.CommonEntityData {
    defaultData.EntityData.YFilter = defaultData.YFilter
    defaultData.EntityData.YangName = "default-data"
    defaultData.EntityData.BundleName = "cisco_ios_xr"
    defaultData.EntityData.ParentYangName = "default_list"
    defaultData.EntityData.SegmentPath = "default-data"
    defaultData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    defaultData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    defaultData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    defaultData.EntityData.Children = make(map[string]types.YChild)
    defaultData.EntityData.Leafs = make(map[string]types.YLeaf)
    defaultData.EntityData.Leafs["default_out_list"] = types.YLeaf{"DefaultOutList", defaultData.DefaultOutList}
    return &(defaultData.EntityData)
}

// Diag_Fans
type Diag_Fans struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Diag_Fans_FansList.
    FansList []Diag_Fans_FansList
}

func (fans *Diag_Fans) GetEntityData() *types.CommonEntityData {
    fans.EntityData.YFilter = fans.YFilter
    fans.EntityData.YangName = "fans"
    fans.EntityData.BundleName = "cisco_ios_xr"
    fans.EntityData.ParentYangName = "diag"
    fans.EntityData.SegmentPath = "fans"
    fans.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fans.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fans.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fans.EntityData.Children = make(map[string]types.YChild)
    fans.EntityData.Children["fans_list"] = types.YChild{"FansList", nil}
    for i := range fans.FansList {
        fans.EntityData.Children[types.GetSegmentPath(&fans.FansList[i])] = types.YChild{"FansList", &fans.FansList[i]}
    }
    fans.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(fans.EntityData)
}

// Diag_Fans_FansList
type Diag_Fans_FansList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Location interface{}

    
    DefaultData Diag_Fans_FansList_DefaultData
}

func (fansList *Diag_Fans_FansList) GetEntityData() *types.CommonEntityData {
    fansList.EntityData.YFilter = fansList.YFilter
    fansList.EntityData.YangName = "fans_list"
    fansList.EntityData.BundleName = "cisco_ios_xr"
    fansList.EntityData.ParentYangName = "fans"
    fansList.EntityData.SegmentPath = "fans_list" + "[location='" + fmt.Sprintf("%v", fansList.Location) + "']"
    fansList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fansList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fansList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fansList.EntityData.Children = make(map[string]types.YChild)
    fansList.EntityData.Children["default-data"] = types.YChild{"DefaultData", &fansList.DefaultData}
    fansList.EntityData.Leafs = make(map[string]types.YLeaf)
    fansList.EntityData.Leafs["location"] = types.YLeaf{"Location", fansList.Location}
    return &(fansList.EntityData)
}

// Diag_Fans_FansList_DefaultData
type Diag_Fans_FansList_DefaultData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of string.
    DefaultOutList []interface{}
}

func (defaultData *Diag_Fans_FansList_DefaultData) GetEntityData() *types.CommonEntityData {
    defaultData.EntityData.YFilter = defaultData.YFilter
    defaultData.EntityData.YangName = "default-data"
    defaultData.EntityData.BundleName = "cisco_ios_xr"
    defaultData.EntityData.ParentYangName = "fans_list"
    defaultData.EntityData.SegmentPath = "default-data"
    defaultData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    defaultData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    defaultData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    defaultData.EntityData.Children = make(map[string]types.YChild)
    defaultData.EntityData.Leafs = make(map[string]types.YLeaf)
    defaultData.EntityData.Leafs["default_out_list"] = types.YLeaf{"DefaultOutList", defaultData.DefaultOutList}
    return &(defaultData.EntityData)
}

// Diag_PowerSupply
type Diag_PowerSupply struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Diag_PowerSupply_PwrList.
    PwrList []Diag_PowerSupply_PwrList
}

func (powerSupply *Diag_PowerSupply) GetEntityData() *types.CommonEntityData {
    powerSupply.EntityData.YFilter = powerSupply.YFilter
    powerSupply.EntityData.YangName = "power-supply"
    powerSupply.EntityData.BundleName = "cisco_ios_xr"
    powerSupply.EntityData.ParentYangName = "diag"
    powerSupply.EntityData.SegmentPath = "power-supply"
    powerSupply.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    powerSupply.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    powerSupply.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    powerSupply.EntityData.Children = make(map[string]types.YChild)
    powerSupply.EntityData.Children["pwr_list"] = types.YChild{"PwrList", nil}
    for i := range powerSupply.PwrList {
        powerSupply.EntityData.Children[types.GetSegmentPath(&powerSupply.PwrList[i])] = types.YChild{"PwrList", &powerSupply.PwrList[i]}
    }
    powerSupply.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(powerSupply.EntityData)
}

// Diag_PowerSupply_PwrList
type Diag_PowerSupply_PwrList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Location interface{}

    
    DefaultData Diag_PowerSupply_PwrList_DefaultData
}

func (pwrList *Diag_PowerSupply_PwrList) GetEntityData() *types.CommonEntityData {
    pwrList.EntityData.YFilter = pwrList.YFilter
    pwrList.EntityData.YangName = "pwr_list"
    pwrList.EntityData.BundleName = "cisco_ios_xr"
    pwrList.EntityData.ParentYangName = "power-supply"
    pwrList.EntityData.SegmentPath = "pwr_list" + "[location='" + fmt.Sprintf("%v", pwrList.Location) + "']"
    pwrList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pwrList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pwrList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pwrList.EntityData.Children = make(map[string]types.YChild)
    pwrList.EntityData.Children["default-data"] = types.YChild{"DefaultData", &pwrList.DefaultData}
    pwrList.EntityData.Leafs = make(map[string]types.YLeaf)
    pwrList.EntityData.Leafs["location"] = types.YLeaf{"Location", pwrList.Location}
    return &(pwrList.EntityData)
}

// Diag_PowerSupply_PwrList_DefaultData
type Diag_PowerSupply_PwrList_DefaultData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of string.
    DefaultOutList []interface{}
}

func (defaultData *Diag_PowerSupply_PwrList_DefaultData) GetEntityData() *types.CommonEntityData {
    defaultData.EntityData.YFilter = defaultData.YFilter
    defaultData.EntityData.YangName = "default-data"
    defaultData.EntityData.BundleName = "cisco_ios_xr"
    defaultData.EntityData.ParentYangName = "pwr_list"
    defaultData.EntityData.SegmentPath = "default-data"
    defaultData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    defaultData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    defaultData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    defaultData.EntityData.Children = make(map[string]types.YChild)
    defaultData.EntityData.Leafs = make(map[string]types.YLeaf)
    defaultData.EntityData.Leafs["default_out_list"] = types.YLeaf{"DefaultOutList", defaultData.DefaultOutList}
    return &(defaultData.EntityData)
}

// Diag_Chassis
type Diag_Chassis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    ChassisCnt Diag_Chassis_ChassisCnt

    
    ChassisEepromCnt Diag_Chassis_ChassisEepromCnt
}

func (chassis *Diag_Chassis) GetEntityData() *types.CommonEntityData {
    chassis.EntityData.YFilter = chassis.YFilter
    chassis.EntityData.YangName = "chassis"
    chassis.EntityData.BundleName = "cisco_ios_xr"
    chassis.EntityData.ParentYangName = "diag"
    chassis.EntityData.SegmentPath = "chassis"
    chassis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    chassis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    chassis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    chassis.EntityData.Children = make(map[string]types.YChild)
    chassis.EntityData.Children["chassis_cnt"] = types.YChild{"ChassisCnt", &chassis.ChassisCnt}
    chassis.EntityData.Children["chassis_eeprom_cnt"] = types.YChild{"ChassisEepromCnt", &chassis.ChassisEepromCnt}
    chassis.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(chassis.EntityData)
}

// Diag_Chassis_ChassisCnt
type Diag_Chassis_ChassisCnt struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Diag_Chassis_ChassisCnt_ChassisList.
    ChassisList []Diag_Chassis_ChassisCnt_ChassisList
}

func (chassisCnt *Diag_Chassis_ChassisCnt) GetEntityData() *types.CommonEntityData {
    chassisCnt.EntityData.YFilter = chassisCnt.YFilter
    chassisCnt.EntityData.YangName = "chassis_cnt"
    chassisCnt.EntityData.BundleName = "cisco_ios_xr"
    chassisCnt.EntityData.ParentYangName = "chassis"
    chassisCnt.EntityData.SegmentPath = "chassis_cnt"
    chassisCnt.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    chassisCnt.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    chassisCnt.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    chassisCnt.EntityData.Children = make(map[string]types.YChild)
    chassisCnt.EntityData.Children["chassis_list"] = types.YChild{"ChassisList", nil}
    for i := range chassisCnt.ChassisList {
        chassisCnt.EntityData.Children[types.GetSegmentPath(&chassisCnt.ChassisList[i])] = types.YChild{"ChassisList", &chassisCnt.ChassisList[i]}
    }
    chassisCnt.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(chassisCnt.EntityData)
}

// Diag_Chassis_ChassisCnt_ChassisList
type Diag_Chassis_ChassisCnt_ChassisList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Location interface{}

    
    DefaultData Diag_Chassis_ChassisCnt_ChassisList_DefaultData
}

func (chassisList *Diag_Chassis_ChassisCnt_ChassisList) GetEntityData() *types.CommonEntityData {
    chassisList.EntityData.YFilter = chassisList.YFilter
    chassisList.EntityData.YangName = "chassis_list"
    chassisList.EntityData.BundleName = "cisco_ios_xr"
    chassisList.EntityData.ParentYangName = "chassis_cnt"
    chassisList.EntityData.SegmentPath = "chassis_list" + "[location='" + fmt.Sprintf("%v", chassisList.Location) + "']"
    chassisList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    chassisList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    chassisList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    chassisList.EntityData.Children = make(map[string]types.YChild)
    chassisList.EntityData.Children["default-data"] = types.YChild{"DefaultData", &chassisList.DefaultData}
    chassisList.EntityData.Leafs = make(map[string]types.YLeaf)
    chassisList.EntityData.Leafs["location"] = types.YLeaf{"Location", chassisList.Location}
    return &(chassisList.EntityData)
}

// Diag_Chassis_ChassisCnt_ChassisList_DefaultData
type Diag_Chassis_ChassisCnt_ChassisList_DefaultData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of string.
    DefaultOutList []interface{}
}

func (defaultData *Diag_Chassis_ChassisCnt_ChassisList_DefaultData) GetEntityData() *types.CommonEntityData {
    defaultData.EntityData.YFilter = defaultData.YFilter
    defaultData.EntityData.YangName = "default-data"
    defaultData.EntityData.BundleName = "cisco_ios_xr"
    defaultData.EntityData.ParentYangName = "chassis_list"
    defaultData.EntityData.SegmentPath = "default-data"
    defaultData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    defaultData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    defaultData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    defaultData.EntityData.Children = make(map[string]types.YChild)
    defaultData.EntityData.Leafs = make(map[string]types.YLeaf)
    defaultData.EntityData.Leafs["default_out_list"] = types.YLeaf{"DefaultOutList", defaultData.DefaultOutList}
    return &(defaultData.EntityData)
}

// Diag_Chassis_ChassisEepromCnt
type Diag_Chassis_ChassisEepromCnt struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Diag_Chassis_ChassisEepromCnt_ChassisEepromList.
    ChassisEepromList []Diag_Chassis_ChassisEepromCnt_ChassisEepromList
}

func (chassisEepromCnt *Diag_Chassis_ChassisEepromCnt) GetEntityData() *types.CommonEntityData {
    chassisEepromCnt.EntityData.YFilter = chassisEepromCnt.YFilter
    chassisEepromCnt.EntityData.YangName = "chassis_eeprom_cnt"
    chassisEepromCnt.EntityData.BundleName = "cisco_ios_xr"
    chassisEepromCnt.EntityData.ParentYangName = "chassis"
    chassisEepromCnt.EntityData.SegmentPath = "chassis_eeprom_cnt"
    chassisEepromCnt.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    chassisEepromCnt.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    chassisEepromCnt.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    chassisEepromCnt.EntityData.Children = make(map[string]types.YChild)
    chassisEepromCnt.EntityData.Children["chassis_eeprom_list"] = types.YChild{"ChassisEepromList", nil}
    for i := range chassisEepromCnt.ChassisEepromList {
        chassisEepromCnt.EntityData.Children[types.GetSegmentPath(&chassisEepromCnt.ChassisEepromList[i])] = types.YChild{"ChassisEepromList", &chassisEepromCnt.ChassisEepromList[i]}
    }
    chassisEepromCnt.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(chassisEepromCnt.EntityData)
}

// Diag_Chassis_ChassisEepromCnt_ChassisEepromList
type Diag_Chassis_ChassisEepromCnt_ChassisEepromList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Location interface{}

    
    EepromData Diag_Chassis_ChassisEepromCnt_ChassisEepromList_EepromData
}

func (chassisEepromList *Diag_Chassis_ChassisEepromCnt_ChassisEepromList) GetEntityData() *types.CommonEntityData {
    chassisEepromList.EntityData.YFilter = chassisEepromList.YFilter
    chassisEepromList.EntityData.YangName = "chassis_eeprom_list"
    chassisEepromList.EntityData.BundleName = "cisco_ios_xr"
    chassisEepromList.EntityData.ParentYangName = "chassis_eeprom_cnt"
    chassisEepromList.EntityData.SegmentPath = "chassis_eeprom_list" + "[location='" + fmt.Sprintf("%v", chassisEepromList.Location) + "']"
    chassisEepromList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    chassisEepromList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    chassisEepromList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    chassisEepromList.EntityData.Children = make(map[string]types.YChild)
    chassisEepromList.EntityData.Children["eeprom-data"] = types.YChild{"EepromData", &chassisEepromList.EepromData}
    chassisEepromList.EntityData.Leafs = make(map[string]types.YLeaf)
    chassisEepromList.EntityData.Leafs["location"] = types.YLeaf{"Location", chassisEepromList.Location}
    return &(chassisEepromList.EntityData)
}

// Diag_Chassis_ChassisEepromCnt_ChassisEepromList_EepromData
type Diag_Chassis_ChassisEepromCnt_ChassisEepromList_EepromData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of string.
    RawList []interface{}
}

func (eepromData *Diag_Chassis_ChassisEepromCnt_ChassisEepromList_EepromData) GetEntityData() *types.CommonEntityData {
    eepromData.EntityData.YFilter = eepromData.YFilter
    eepromData.EntityData.YangName = "eeprom-data"
    eepromData.EntityData.BundleName = "cisco_ios_xr"
    eepromData.EntityData.ParentYangName = "chassis_eeprom_list"
    eepromData.EntityData.SegmentPath = "eeprom-data"
    eepromData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    eepromData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    eepromData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    eepromData.EntityData.Children = make(map[string]types.YChild)
    eepromData.EntityData.Leafs = make(map[string]types.YLeaf)
    eepromData.EntityData.Leafs["raw_list"] = types.YLeaf{"RawList", eepromData.RawList}
    return &(eepromData.EntityData)
}

// Diag_Summary
type Diag_Summary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Diag_Summary_SummaryList.
    SummaryList []Diag_Summary_SummaryList
}

func (summary *Diag_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "cisco_ios_xr"
    summary.EntityData.ParentYangName = "diag"
    summary.EntityData.SegmentPath = "summary"
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = make(map[string]types.YChild)
    summary.EntityData.Children["summary_list"] = types.YChild{"SummaryList", nil}
    for i := range summary.SummaryList {
        summary.EntityData.Children[types.GetSegmentPath(&summary.SummaryList[i])] = types.YChild{"SummaryList", &summary.SummaryList[i]}
    }
    summary.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(summary.EntityData)
}

// Diag_Summary_SummaryList
type Diag_Summary_SummaryList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Location interface{}

    
    SummaryData Diag_Summary_SummaryList_SummaryData
}

func (summaryList *Diag_Summary_SummaryList) GetEntityData() *types.CommonEntityData {
    summaryList.EntityData.YFilter = summaryList.YFilter
    summaryList.EntityData.YangName = "summary_list"
    summaryList.EntityData.BundleName = "cisco_ios_xr"
    summaryList.EntityData.ParentYangName = "summary"
    summaryList.EntityData.SegmentPath = "summary_list" + "[location='" + fmt.Sprintf("%v", summaryList.Location) + "']"
    summaryList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summaryList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summaryList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summaryList.EntityData.Children = make(map[string]types.YChild)
    summaryList.EntityData.Children["summary-data"] = types.YChild{"SummaryData", &summaryList.SummaryData}
    summaryList.EntityData.Leafs = make(map[string]types.YLeaf)
    summaryList.EntityData.Leafs["location"] = types.YLeaf{"Location", summaryList.Location}
    return &(summaryList.EntityData)
}

// Diag_Summary_SummaryList_SummaryData
type Diag_Summary_SummaryList_SummaryData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of string.
    SummaryOutList []interface{}
}

func (summaryData *Diag_Summary_SummaryList_SummaryData) GetEntityData() *types.CommonEntityData {
    summaryData.EntityData.YFilter = summaryData.YFilter
    summaryData.EntityData.YangName = "summary-data"
    summaryData.EntityData.BundleName = "cisco_ios_xr"
    summaryData.EntityData.ParentYangName = "summary_list"
    summaryData.EntityData.SegmentPath = "summary-data"
    summaryData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summaryData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summaryData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summaryData.EntityData.Children = make(map[string]types.YChild)
    summaryData.EntityData.Leafs = make(map[string]types.YLeaf)
    summaryData.EntityData.Leafs["summary_out_list"] = types.YLeaf{"SummaryOutList", summaryData.SummaryOutList}
    return &(summaryData.EntityData)
}

// Diag_Eeprom
type Diag_Eeprom struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Diag_Eeprom_EepromList.
    EepromList []Diag_Eeprom_EepromList
}

func (eeprom *Diag_Eeprom) GetEntityData() *types.CommonEntityData {
    eeprom.EntityData.YFilter = eeprom.YFilter
    eeprom.EntityData.YangName = "eeprom"
    eeprom.EntityData.BundleName = "cisco_ios_xr"
    eeprom.EntityData.ParentYangName = "diag"
    eeprom.EntityData.SegmentPath = "eeprom"
    eeprom.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    eeprom.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    eeprom.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    eeprom.EntityData.Children = make(map[string]types.YChild)
    eeprom.EntityData.Children["eeprom_list"] = types.YChild{"EepromList", nil}
    for i := range eeprom.EepromList {
        eeprom.EntityData.Children[types.GetSegmentPath(&eeprom.EepromList[i])] = types.YChild{"EepromList", &eeprom.EepromList[i]}
    }
    eeprom.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(eeprom.EntityData)
}

// Diag_Eeprom_EepromList
type Diag_Eeprom_EepromList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Location interface{}

    
    EepromData Diag_Eeprom_EepromList_EepromData
}

func (eepromList *Diag_Eeprom_EepromList) GetEntityData() *types.CommonEntityData {
    eepromList.EntityData.YFilter = eepromList.YFilter
    eepromList.EntityData.YangName = "eeprom_list"
    eepromList.EntityData.BundleName = "cisco_ios_xr"
    eepromList.EntityData.ParentYangName = "eeprom"
    eepromList.EntityData.SegmentPath = "eeprom_list" + "[location='" + fmt.Sprintf("%v", eepromList.Location) + "']"
    eepromList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    eepromList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    eepromList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    eepromList.EntityData.Children = make(map[string]types.YChild)
    eepromList.EntityData.Children["eeprom-data"] = types.YChild{"EepromData", &eepromList.EepromData}
    eepromList.EntityData.Leafs = make(map[string]types.YLeaf)
    eepromList.EntityData.Leafs["location"] = types.YLeaf{"Location", eepromList.Location}
    return &(eepromList.EntityData)
}

// Diag_Eeprom_EepromList_EepromData
type Diag_Eeprom_EepromList_EepromData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of string.
    RawList []interface{}
}

func (eepromData *Diag_Eeprom_EepromList_EepromData) GetEntityData() *types.CommonEntityData {
    eepromData.EntityData.YFilter = eepromData.YFilter
    eepromData.EntityData.YangName = "eeprom-data"
    eepromData.EntityData.BundleName = "cisco_ios_xr"
    eepromData.EntityData.ParentYangName = "eeprom_list"
    eepromData.EntityData.SegmentPath = "eeprom-data"
    eepromData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    eepromData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    eepromData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    eepromData.EntityData.Children = make(map[string]types.YChild)
    eepromData.EntityData.Leafs = make(map[string]types.YLeaf)
    eepromData.EntityData.Leafs["raw_list"] = types.YLeaf{"RawList", eepromData.RawList}
    return &(eepromData.EntityData)
}

// Diag_Detail
type Diag_Detail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Diag_Detail_DetailList.
    DetailList []Diag_Detail_DetailList
}

func (detail *Diag_Detail) GetEntityData() *types.CommonEntityData {
    detail.EntityData.YFilter = detail.YFilter
    detail.EntityData.YangName = "detail"
    detail.EntityData.BundleName = "cisco_ios_xr"
    detail.EntityData.ParentYangName = "diag"
    detail.EntityData.SegmentPath = "detail"
    detail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detail.EntityData.Children = make(map[string]types.YChild)
    detail.EntityData.Children["detail_list"] = types.YChild{"DetailList", nil}
    for i := range detail.DetailList {
        detail.EntityData.Children[types.GetSegmentPath(&detail.DetailList[i])] = types.YChild{"DetailList", &detail.DetailList[i]}
    }
    detail.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(detail.EntityData)
}

// Diag_Detail_DetailList
type Diag_Detail_DetailList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Location interface{}

    
    DetailData Diag_Detail_DetailList_DetailData
}

func (detailList *Diag_Detail_DetailList) GetEntityData() *types.CommonEntityData {
    detailList.EntityData.YFilter = detailList.YFilter
    detailList.EntityData.YangName = "detail_list"
    detailList.EntityData.BundleName = "cisco_ios_xr"
    detailList.EntityData.ParentYangName = "detail"
    detailList.EntityData.SegmentPath = "detail_list" + "[location='" + fmt.Sprintf("%v", detailList.Location) + "']"
    detailList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detailList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detailList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detailList.EntityData.Children = make(map[string]types.YChild)
    detailList.EntityData.Children["detail-data"] = types.YChild{"DetailData", &detailList.DetailData}
    detailList.EntityData.Leafs = make(map[string]types.YLeaf)
    detailList.EntityData.Leafs["location"] = types.YLeaf{"Location", detailList.Location}
    return &(detailList.EntityData)
}

// Diag_Detail_DetailList_DetailData
type Diag_Detail_DetailList_DetailData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of string.
    DetailOutList []interface{}
}

func (detailData *Diag_Detail_DetailList_DetailData) GetEntityData() *types.CommonEntityData {
    detailData.EntityData.YFilter = detailData.YFilter
    detailData.EntityData.YangName = "detail-data"
    detailData.EntityData.BundleName = "cisco_ios_xr"
    detailData.EntityData.ParentYangName = "detail_list"
    detailData.EntityData.SegmentPath = "detail-data"
    detailData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detailData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detailData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detailData.EntityData.Children = make(map[string]types.YChild)
    detailData.EntityData.Leafs = make(map[string]types.YLeaf)
    detailData.EntityData.Leafs["detail_out_list"] = types.YLeaf{"DetailOutList", detailData.DetailOutList}
    return &(detailData.EntityData)
}

