// This module contains definitions
// for the Calvados model objects.
// 
// Copyright (c) 2012-2018 by Cisco Systems, Inc.
// All rights reserved.
package cisco_entity_fru_control_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_entity_fru_control_mib"))
    ydk.RegisterEntity("{http://tail-f.com/ns/mibs/CISCO-ENTITY-FRU-CONTROL-MIB/200311240000Z CISCO-ENTITY-FRU-CONTROL-MIB}", reflect.TypeOf(CISCOENTITYFRUCONTROLMIB{}))
    ydk.RegisterEntity("CISCO-ENTITY-FRU-CONTROL-MIB:CISCO-ENTITY-FRU-CONTROL-MIB", reflect.TypeOf(CISCOENTITYFRUCONTROLMIB{}))
}

// PowerRedundancyType
type PowerRedundancyType string

const (
    PowerRedundancyType_notsupported PowerRedundancyType = "notsupported"

    PowerRedundancyType_redundant PowerRedundancyType = "redundant"

    PowerRedundancyType_combined PowerRedundancyType = "combined"
)

// PowerAdminType
type PowerAdminType string

const (
    PowerAdminType_on PowerAdminType = "on"

    PowerAdminType_off PowerAdminType = "off"

    PowerAdminType_inlineAuto PowerAdminType = "inlineAuto"

    PowerAdminType_inlineOn PowerAdminType = "inlineOn"
)

// PowerOperType
type PowerOperType string

const (
    PowerOperType_offEnvOther PowerOperType = "offEnvOther"

    PowerOperType_on PowerOperType = "on"

    PowerOperType_offAdmin PowerOperType = "offAdmin"

    PowerOperType_offDenied PowerOperType = "offDenied"

    PowerOperType_offEnvPower PowerOperType = "offEnvPower"

    PowerOperType_offEnvTemp PowerOperType = "offEnvTemp"

    PowerOperType_offEnvFan PowerOperType = "offEnvFan"

    PowerOperType_failed PowerOperType = "failed"

    PowerOperType_onButFanFail PowerOperType = "onButFanFail"
)

// ModuleAdminType
type ModuleAdminType string

const (
    ModuleAdminType_enabled ModuleAdminType = "enabled"

    ModuleAdminType_disabled ModuleAdminType = "disabled"

    ModuleAdminType_reset ModuleAdminType = "reset"

    ModuleAdminType_outOfServiceAdmin ModuleAdminType = "outOfServiceAdmin"
)

// ModuleOperType
type ModuleOperType string

const (
    ModuleOperType_unknown ModuleOperType = "unknown"

    ModuleOperType_ok ModuleOperType = "ok"

    ModuleOperType_disabled ModuleOperType = "disabled"

    ModuleOperType_okButDiagFailed ModuleOperType = "okButDiagFailed"

    ModuleOperType_boot ModuleOperType = "boot"

    ModuleOperType_selfTest ModuleOperType = "selfTest"

    ModuleOperType_failed ModuleOperType = "failed"

    ModuleOperType_missing ModuleOperType = "missing"

    ModuleOperType_mismatchWithParent ModuleOperType = "mismatchWithParent"

    ModuleOperType_mismatchConfig ModuleOperType = "mismatchConfig"

    ModuleOperType_diagFailed ModuleOperType = "diagFailed"

    ModuleOperType_dormant ModuleOperType = "dormant"

    ModuleOperType_outOfServiceAdmin ModuleOperType = "outOfServiceAdmin"

    ModuleOperType_outOfServiceEnvTemp ModuleOperType = "outOfServiceEnvTemp"

    ModuleOperType_poweredDown ModuleOperType = "poweredDown"

    ModuleOperType_poweredUp ModuleOperType = "poweredUp"

    ModuleOperType_powerDenied ModuleOperType = "powerDenied"

    ModuleOperType_powerCycled ModuleOperType = "powerCycled"

    ModuleOperType_okButPowerOverWarning ModuleOperType = "okButPowerOverWarning"

    ModuleOperType_okButPowerOverCritical ModuleOperType = "okButPowerOverCritical"

    ModuleOperType_updatingFPD ModuleOperType = "updatingFPD"
)

// ModuleResetReasonType
type ModuleResetReasonType string

const (
    ModuleResetReasonType_unknown ModuleResetReasonType = "unknown"

    ModuleResetReasonType_powerUp ModuleResetReasonType = "powerUp"

    ModuleResetReasonType_parityError ModuleResetReasonType = "parityError"

    ModuleResetReasonType_clearConfigReset ModuleResetReasonType = "clearConfigReset"

    ModuleResetReasonType_manualReset ModuleResetReasonType = "manualReset"

    ModuleResetReasonType_watchDogTimeoutReset ModuleResetReasonType = "watchDogTimeoutReset"

    ModuleResetReasonType_resourceOverflowReset ModuleResetReasonType = "resourceOverflowReset"

    ModuleResetReasonType_missingTaskReset ModuleResetReasonType = "missingTaskReset"

    ModuleResetReasonType_lowVoltageReset ModuleResetReasonType = "lowVoltageReset"

    ModuleResetReasonType_controllerReset ModuleResetReasonType = "controllerReset"

    ModuleResetReasonType_systemReset ModuleResetReasonType = "systemReset"

    ModuleResetReasonType_switchoverReset ModuleResetReasonType = "switchoverReset"

    ModuleResetReasonType_upgradeReset ModuleResetReasonType = "upgradeReset"

    ModuleResetReasonType_downgradeReset ModuleResetReasonType = "downgradeReset"

    ModuleResetReasonType_cacheErrorReset ModuleResetReasonType = "cacheErrorReset"

    ModuleResetReasonType_deviceDriverReset ModuleResetReasonType = "deviceDriverReset"

    ModuleResetReasonType_softwareExceptionReset ModuleResetReasonType = "softwareExceptionReset"

    ModuleResetReasonType_restoreConfigReset ModuleResetReasonType = "restoreConfigReset"

    ModuleResetReasonType_abortRevReset ModuleResetReasonType = "abortRevReset"

    ModuleResetReasonType_burnBootReset ModuleResetReasonType = "burnBootReset"

    ModuleResetReasonType_standbyCdHealthierReset ModuleResetReasonType = "standbyCdHealthierReset"

    ModuleResetReasonType_nonNativeConfigClearReset ModuleResetReasonType = "nonNativeConfigClearReset"

    ModuleResetReasonType_memoryProtectionErrorReset ModuleResetReasonType = "memoryProtectionErrorReset"
)

// CefcFanTrayOperStatusType
type CefcFanTrayOperStatusType string

const (
    CefcFanTrayOperStatusType_unknown CefcFanTrayOperStatusType = "unknown"

    CefcFanTrayOperStatusType_up CefcFanTrayOperStatusType = "up"

    CefcFanTrayOperStatusType_down CefcFanTrayOperStatusType = "down"

    CefcFanTrayOperStatusType_warning CefcFanTrayOperStatusType = "warning"
)

// CefcPhysicalStatusType
type CefcPhysicalStatusType string

const (
    CefcPhysicalStatusType_other CefcPhysicalStatusType = "other"

    CefcPhysicalStatusType_supported CefcPhysicalStatusType = "supported"

    CefcPhysicalStatusType_unsupported CefcPhysicalStatusType = "unsupported"

    CefcPhysicalStatusType_incompatible CefcPhysicalStatusType = "incompatible"
)

// CISCOENTITYFRUCONTROLMIB
type CISCOENTITYFRUCONTROLMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    CefcFRUPower CISCOENTITYFRUCONTROLMIB_CefcFRUPower

    
    CefcMIBNotificationEnables CISCOENTITYFRUCONTROLMIB_CefcMIBNotificationEnables

    
    CefcFRUPowerSupplyGroupTable CISCOENTITYFRUCONTROLMIB_CefcFRUPowerSupplyGroupTable

    
    CefcFRUPowerStatusTable CISCOENTITYFRUCONTROLMIB_CefcFRUPowerStatusTable

    
    CefcFRUPowerSupplyValueTable CISCOENTITYFRUCONTROLMIB_CefcFRUPowerSupplyValueTable

    
    CefcModuleTable CISCOENTITYFRUCONTROLMIB_CefcModuleTable

    
    CefcIntelliModuleTable CISCOENTITYFRUCONTROLMIB_CefcIntelliModuleTable

    
    CefcFanTrayStatusTable CISCOENTITYFRUCONTROLMIB_CefcFanTrayStatusTable

    
    CefcPhysicalTable CISCOENTITYFRUCONTROLMIB_CefcPhysicalTable
}

func (cISCOENTITYFRUCONTROLMIB *CISCOENTITYFRUCONTROLMIB) GetEntityData() *types.CommonEntityData {
    cISCOENTITYFRUCONTROLMIB.EntityData.YFilter = cISCOENTITYFRUCONTROLMIB.YFilter
    cISCOENTITYFRUCONTROLMIB.EntityData.YangName = "CISCO-ENTITY-FRU-CONTROL-MIB"
    cISCOENTITYFRUCONTROLMIB.EntityData.BundleName = "cisco_ios_xr"
    cISCOENTITYFRUCONTROLMIB.EntityData.ParentYangName = "CISCO-ENTITY-FRU-CONTROL-MIB"
    cISCOENTITYFRUCONTROLMIB.EntityData.SegmentPath = "CISCO-ENTITY-FRU-CONTROL-MIB:CISCO-ENTITY-FRU-CONTROL-MIB"
    cISCOENTITYFRUCONTROLMIB.EntityData.AbsolutePath = cISCOENTITYFRUCONTROLMIB.EntityData.SegmentPath
    cISCOENTITYFRUCONTROLMIB.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cISCOENTITYFRUCONTROLMIB.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cISCOENTITYFRUCONTROLMIB.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cISCOENTITYFRUCONTROLMIB.EntityData.Children = types.NewOrderedMap()
    cISCOENTITYFRUCONTROLMIB.EntityData.Children.Append("cefcFRUPower", types.YChild{"CefcFRUPower", &cISCOENTITYFRUCONTROLMIB.CefcFRUPower})
    cISCOENTITYFRUCONTROLMIB.EntityData.Children.Append("cefcMIBNotificationEnables", types.YChild{"CefcMIBNotificationEnables", &cISCOENTITYFRUCONTROLMIB.CefcMIBNotificationEnables})
    cISCOENTITYFRUCONTROLMIB.EntityData.Children.Append("cefcFRUPowerSupplyGroupTable", types.YChild{"CefcFRUPowerSupplyGroupTable", &cISCOENTITYFRUCONTROLMIB.CefcFRUPowerSupplyGroupTable})
    cISCOENTITYFRUCONTROLMIB.EntityData.Children.Append("cefcFRUPowerStatusTable", types.YChild{"CefcFRUPowerStatusTable", &cISCOENTITYFRUCONTROLMIB.CefcFRUPowerStatusTable})
    cISCOENTITYFRUCONTROLMIB.EntityData.Children.Append("cefcFRUPowerSupplyValueTable", types.YChild{"CefcFRUPowerSupplyValueTable", &cISCOENTITYFRUCONTROLMIB.CefcFRUPowerSupplyValueTable})
    cISCOENTITYFRUCONTROLMIB.EntityData.Children.Append("cefcModuleTable", types.YChild{"CefcModuleTable", &cISCOENTITYFRUCONTROLMIB.CefcModuleTable})
    cISCOENTITYFRUCONTROLMIB.EntityData.Children.Append("cefcIntelliModuleTable", types.YChild{"CefcIntelliModuleTable", &cISCOENTITYFRUCONTROLMIB.CefcIntelliModuleTable})
    cISCOENTITYFRUCONTROLMIB.EntityData.Children.Append("cefcFanTrayStatusTable", types.YChild{"CefcFanTrayStatusTable", &cISCOENTITYFRUCONTROLMIB.CefcFanTrayStatusTable})
    cISCOENTITYFRUCONTROLMIB.EntityData.Children.Append("cefcPhysicalTable", types.YChild{"CefcPhysicalTable", &cISCOENTITYFRUCONTROLMIB.CefcPhysicalTable})
    cISCOENTITYFRUCONTROLMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOENTITYFRUCONTROLMIB.EntityData.YListKeys = []string {}

    return &(cISCOENTITYFRUCONTROLMIB.EntityData)
}

// CISCOENTITYFRUCONTROLMIB_CefcFRUPower
type CISCOENTITYFRUCONTROLMIB_CefcFRUPower struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..12500. The default value is 12500.
    CefcMaxDefaultInLinePower interface{}

    // The type is interface{} with range: 0..4294967295.
    CefcMaxDefaultHighInLinePower interface{}
}

func (cefcFRUPower *CISCOENTITYFRUCONTROLMIB_CefcFRUPower) GetEntityData() *types.CommonEntityData {
    cefcFRUPower.EntityData.YFilter = cefcFRUPower.YFilter
    cefcFRUPower.EntityData.YangName = "cefcFRUPower"
    cefcFRUPower.EntityData.BundleName = "cisco_ios_xr"
    cefcFRUPower.EntityData.ParentYangName = "CISCO-ENTITY-FRU-CONTROL-MIB"
    cefcFRUPower.EntityData.SegmentPath = "cefcFRUPower"
    cefcFRUPower.EntityData.AbsolutePath = "CISCO-ENTITY-FRU-CONTROL-MIB:CISCO-ENTITY-FRU-CONTROL-MIB/" + cefcFRUPower.EntityData.SegmentPath
    cefcFRUPower.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cefcFRUPower.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cefcFRUPower.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cefcFRUPower.EntityData.Children = types.NewOrderedMap()
    cefcFRUPower.EntityData.Leafs = types.NewOrderedMap()
    cefcFRUPower.EntityData.Leafs.Append("cefcMaxDefaultInLinePower", types.YLeaf{"CefcMaxDefaultInLinePower", cefcFRUPower.CefcMaxDefaultInLinePower})
    cefcFRUPower.EntityData.Leafs.Append("cefcMaxDefaultHighInLinePower", types.YLeaf{"CefcMaxDefaultHighInLinePower", cefcFRUPower.CefcMaxDefaultHighInLinePower})

    cefcFRUPower.EntityData.YListKeys = []string {}

    return &(cefcFRUPower.EntityData)
}

// CISCOENTITYFRUCONTROLMIB_CefcMIBNotificationEnables
type CISCOENTITYFRUCONTROLMIB_CefcMIBNotificationEnables struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is TruthValue. The default value is false.
    CefcMIBEnableStatusNotification interface{}
}

func (cefcMIBNotificationEnables *CISCOENTITYFRUCONTROLMIB_CefcMIBNotificationEnables) GetEntityData() *types.CommonEntityData {
    cefcMIBNotificationEnables.EntityData.YFilter = cefcMIBNotificationEnables.YFilter
    cefcMIBNotificationEnables.EntityData.YangName = "cefcMIBNotificationEnables"
    cefcMIBNotificationEnables.EntityData.BundleName = "cisco_ios_xr"
    cefcMIBNotificationEnables.EntityData.ParentYangName = "CISCO-ENTITY-FRU-CONTROL-MIB"
    cefcMIBNotificationEnables.EntityData.SegmentPath = "cefcMIBNotificationEnables"
    cefcMIBNotificationEnables.EntityData.AbsolutePath = "CISCO-ENTITY-FRU-CONTROL-MIB:CISCO-ENTITY-FRU-CONTROL-MIB/" + cefcMIBNotificationEnables.EntityData.SegmentPath
    cefcMIBNotificationEnables.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cefcMIBNotificationEnables.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cefcMIBNotificationEnables.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cefcMIBNotificationEnables.EntityData.Children = types.NewOrderedMap()
    cefcMIBNotificationEnables.EntityData.Leafs = types.NewOrderedMap()
    cefcMIBNotificationEnables.EntityData.Leafs.Append("cefcMIBEnableStatusNotification", types.YLeaf{"CefcMIBEnableStatusNotification", cefcMIBNotificationEnables.CefcMIBEnableStatusNotification})

    cefcMIBNotificationEnables.EntityData.YListKeys = []string {}

    return &(cefcMIBNotificationEnables.EntityData)
}

// CISCOENTITYFRUCONTROLMIB_CefcFRUPowerSupplyGroupTable
type CISCOENTITYFRUCONTROLMIB_CefcFRUPowerSupplyGroupTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // CISCOENTITYFRUCONTROLMIB_CefcFRUPowerSupplyGroupTable_CefcFRUPowerSupplyGroupEntry.
    CefcFRUPowerSupplyGroupEntry []*CISCOENTITYFRUCONTROLMIB_CefcFRUPowerSupplyGroupTable_CefcFRUPowerSupplyGroupEntry
}

func (cefcFRUPowerSupplyGroupTable *CISCOENTITYFRUCONTROLMIB_CefcFRUPowerSupplyGroupTable) GetEntityData() *types.CommonEntityData {
    cefcFRUPowerSupplyGroupTable.EntityData.YFilter = cefcFRUPowerSupplyGroupTable.YFilter
    cefcFRUPowerSupplyGroupTable.EntityData.YangName = "cefcFRUPowerSupplyGroupTable"
    cefcFRUPowerSupplyGroupTable.EntityData.BundleName = "cisco_ios_xr"
    cefcFRUPowerSupplyGroupTable.EntityData.ParentYangName = "CISCO-ENTITY-FRU-CONTROL-MIB"
    cefcFRUPowerSupplyGroupTable.EntityData.SegmentPath = "cefcFRUPowerSupplyGroupTable"
    cefcFRUPowerSupplyGroupTable.EntityData.AbsolutePath = "CISCO-ENTITY-FRU-CONTROL-MIB:CISCO-ENTITY-FRU-CONTROL-MIB/" + cefcFRUPowerSupplyGroupTable.EntityData.SegmentPath
    cefcFRUPowerSupplyGroupTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cefcFRUPowerSupplyGroupTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cefcFRUPowerSupplyGroupTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cefcFRUPowerSupplyGroupTable.EntityData.Children = types.NewOrderedMap()
    cefcFRUPowerSupplyGroupTable.EntityData.Children.Append("cefcFRUPowerSupplyGroupEntry", types.YChild{"CefcFRUPowerSupplyGroupEntry", nil})
    for i := range cefcFRUPowerSupplyGroupTable.CefcFRUPowerSupplyGroupEntry {
        cefcFRUPowerSupplyGroupTable.EntityData.Children.Append(types.GetSegmentPath(cefcFRUPowerSupplyGroupTable.CefcFRUPowerSupplyGroupEntry[i]), types.YChild{"CefcFRUPowerSupplyGroupEntry", cefcFRUPowerSupplyGroupTable.CefcFRUPowerSupplyGroupEntry[i]})
    }
    cefcFRUPowerSupplyGroupTable.EntityData.Leafs = types.NewOrderedMap()

    cefcFRUPowerSupplyGroupTable.EntityData.YListKeys = []string {}

    return &(cefcFRUPowerSupplyGroupTable.EntityData)
}

// CISCOENTITYFRUCONTROLMIB_CefcFRUPowerSupplyGroupTable_CefcFRUPowerSupplyGroupEntry
type CISCOENTITYFRUCONTROLMIB_CefcFRUPowerSupplyGroupTable_CefcFRUPowerSupplyGroupEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is interface{} with range: 1..2147483647.
    EntPhysicalIndex interface{}

    // The type is PowerRedundancyType.
    CefcPowerRedundancyMode interface{}

    // The type is string with length: 0..255.
    CefcPowerUnits interface{}

    // The type is interface{} with range: -1000000000..1000000000.
    CefcTotalAvailableCurrent interface{}

    // The type is interface{} with range: -1000000000..1000000000.
    CefcTotalDrawnCurrent interface{}

    // The type is PowerRedundancyType.
    CefcPowerRedundancyOperMode interface{}
}

func (cefcFRUPowerSupplyGroupEntry *CISCOENTITYFRUCONTROLMIB_CefcFRUPowerSupplyGroupTable_CefcFRUPowerSupplyGroupEntry) GetEntityData() *types.CommonEntityData {
    cefcFRUPowerSupplyGroupEntry.EntityData.YFilter = cefcFRUPowerSupplyGroupEntry.YFilter
    cefcFRUPowerSupplyGroupEntry.EntityData.YangName = "cefcFRUPowerSupplyGroupEntry"
    cefcFRUPowerSupplyGroupEntry.EntityData.BundleName = "cisco_ios_xr"
    cefcFRUPowerSupplyGroupEntry.EntityData.ParentYangName = "cefcFRUPowerSupplyGroupTable"
    cefcFRUPowerSupplyGroupEntry.EntityData.SegmentPath = "cefcFRUPowerSupplyGroupEntry" + types.AddKeyToken(cefcFRUPowerSupplyGroupEntry.EntPhysicalIndex, "entPhysicalIndex")
    cefcFRUPowerSupplyGroupEntry.EntityData.AbsolutePath = "CISCO-ENTITY-FRU-CONTROL-MIB:CISCO-ENTITY-FRU-CONTROL-MIB/cefcFRUPowerSupplyGroupTable/" + cefcFRUPowerSupplyGroupEntry.EntityData.SegmentPath
    cefcFRUPowerSupplyGroupEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cefcFRUPowerSupplyGroupEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cefcFRUPowerSupplyGroupEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cefcFRUPowerSupplyGroupEntry.EntityData.Children = types.NewOrderedMap()
    cefcFRUPowerSupplyGroupEntry.EntityData.Leafs = types.NewOrderedMap()
    cefcFRUPowerSupplyGroupEntry.EntityData.Leafs.Append("entPhysicalIndex", types.YLeaf{"EntPhysicalIndex", cefcFRUPowerSupplyGroupEntry.EntPhysicalIndex})
    cefcFRUPowerSupplyGroupEntry.EntityData.Leafs.Append("cefcPowerRedundancyMode", types.YLeaf{"CefcPowerRedundancyMode", cefcFRUPowerSupplyGroupEntry.CefcPowerRedundancyMode})
    cefcFRUPowerSupplyGroupEntry.EntityData.Leafs.Append("cefcPowerUnits", types.YLeaf{"CefcPowerUnits", cefcFRUPowerSupplyGroupEntry.CefcPowerUnits})
    cefcFRUPowerSupplyGroupEntry.EntityData.Leafs.Append("cefcTotalAvailableCurrent", types.YLeaf{"CefcTotalAvailableCurrent", cefcFRUPowerSupplyGroupEntry.CefcTotalAvailableCurrent})
    cefcFRUPowerSupplyGroupEntry.EntityData.Leafs.Append("cefcTotalDrawnCurrent", types.YLeaf{"CefcTotalDrawnCurrent", cefcFRUPowerSupplyGroupEntry.CefcTotalDrawnCurrent})
    cefcFRUPowerSupplyGroupEntry.EntityData.Leafs.Append("cefcPowerRedundancyOperMode", types.YLeaf{"CefcPowerRedundancyOperMode", cefcFRUPowerSupplyGroupEntry.CefcPowerRedundancyOperMode})

    cefcFRUPowerSupplyGroupEntry.EntityData.YListKeys = []string {"EntPhysicalIndex"}

    return &(cefcFRUPowerSupplyGroupEntry.EntityData)
}

// CISCOENTITYFRUCONTROLMIB_CefcFRUPowerStatusTable
type CISCOENTITYFRUCONTROLMIB_CefcFRUPowerStatusTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // CISCOENTITYFRUCONTROLMIB_CefcFRUPowerStatusTable_CefcFRUPowerStatusEntry.
    CefcFRUPowerStatusEntry []*CISCOENTITYFRUCONTROLMIB_CefcFRUPowerStatusTable_CefcFRUPowerStatusEntry
}

func (cefcFRUPowerStatusTable *CISCOENTITYFRUCONTROLMIB_CefcFRUPowerStatusTable) GetEntityData() *types.CommonEntityData {
    cefcFRUPowerStatusTable.EntityData.YFilter = cefcFRUPowerStatusTable.YFilter
    cefcFRUPowerStatusTable.EntityData.YangName = "cefcFRUPowerStatusTable"
    cefcFRUPowerStatusTable.EntityData.BundleName = "cisco_ios_xr"
    cefcFRUPowerStatusTable.EntityData.ParentYangName = "CISCO-ENTITY-FRU-CONTROL-MIB"
    cefcFRUPowerStatusTable.EntityData.SegmentPath = "cefcFRUPowerStatusTable"
    cefcFRUPowerStatusTable.EntityData.AbsolutePath = "CISCO-ENTITY-FRU-CONTROL-MIB:CISCO-ENTITY-FRU-CONTROL-MIB/" + cefcFRUPowerStatusTable.EntityData.SegmentPath
    cefcFRUPowerStatusTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cefcFRUPowerStatusTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cefcFRUPowerStatusTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cefcFRUPowerStatusTable.EntityData.Children = types.NewOrderedMap()
    cefcFRUPowerStatusTable.EntityData.Children.Append("cefcFRUPowerStatusEntry", types.YChild{"CefcFRUPowerStatusEntry", nil})
    for i := range cefcFRUPowerStatusTable.CefcFRUPowerStatusEntry {
        cefcFRUPowerStatusTable.EntityData.Children.Append(types.GetSegmentPath(cefcFRUPowerStatusTable.CefcFRUPowerStatusEntry[i]), types.YChild{"CefcFRUPowerStatusEntry", cefcFRUPowerStatusTable.CefcFRUPowerStatusEntry[i]})
    }
    cefcFRUPowerStatusTable.EntityData.Leafs = types.NewOrderedMap()

    cefcFRUPowerStatusTable.EntityData.YListKeys = []string {}

    return &(cefcFRUPowerStatusTable.EntityData)
}

// CISCOENTITYFRUCONTROLMIB_CefcFRUPowerStatusTable_CefcFRUPowerStatusEntry
type CISCOENTITYFRUCONTROLMIB_CefcFRUPowerStatusTable_CefcFRUPowerStatusEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is interface{} with range: 1..2147483647.
    EntPhysicalIndex interface{}

    // The type is PowerAdminType.
    CefcFRUPowerAdminStatus interface{}

    // The type is PowerOperType.
    CefcFRUPowerOperStatus interface{}

    // The type is interface{} with range: -1000000000..1000000000.
    CefcFRUCurrent interface{}
}

func (cefcFRUPowerStatusEntry *CISCOENTITYFRUCONTROLMIB_CefcFRUPowerStatusTable_CefcFRUPowerStatusEntry) GetEntityData() *types.CommonEntityData {
    cefcFRUPowerStatusEntry.EntityData.YFilter = cefcFRUPowerStatusEntry.YFilter
    cefcFRUPowerStatusEntry.EntityData.YangName = "cefcFRUPowerStatusEntry"
    cefcFRUPowerStatusEntry.EntityData.BundleName = "cisco_ios_xr"
    cefcFRUPowerStatusEntry.EntityData.ParentYangName = "cefcFRUPowerStatusTable"
    cefcFRUPowerStatusEntry.EntityData.SegmentPath = "cefcFRUPowerStatusEntry" + types.AddKeyToken(cefcFRUPowerStatusEntry.EntPhysicalIndex, "entPhysicalIndex")
    cefcFRUPowerStatusEntry.EntityData.AbsolutePath = "CISCO-ENTITY-FRU-CONTROL-MIB:CISCO-ENTITY-FRU-CONTROL-MIB/cefcFRUPowerStatusTable/" + cefcFRUPowerStatusEntry.EntityData.SegmentPath
    cefcFRUPowerStatusEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cefcFRUPowerStatusEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cefcFRUPowerStatusEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cefcFRUPowerStatusEntry.EntityData.Children = types.NewOrderedMap()
    cefcFRUPowerStatusEntry.EntityData.Leafs = types.NewOrderedMap()
    cefcFRUPowerStatusEntry.EntityData.Leafs.Append("entPhysicalIndex", types.YLeaf{"EntPhysicalIndex", cefcFRUPowerStatusEntry.EntPhysicalIndex})
    cefcFRUPowerStatusEntry.EntityData.Leafs.Append("cefcFRUPowerAdminStatus", types.YLeaf{"CefcFRUPowerAdminStatus", cefcFRUPowerStatusEntry.CefcFRUPowerAdminStatus})
    cefcFRUPowerStatusEntry.EntityData.Leafs.Append("cefcFRUPowerOperStatus", types.YLeaf{"CefcFRUPowerOperStatus", cefcFRUPowerStatusEntry.CefcFRUPowerOperStatus})
    cefcFRUPowerStatusEntry.EntityData.Leafs.Append("cefcFRUCurrent", types.YLeaf{"CefcFRUCurrent", cefcFRUPowerStatusEntry.CefcFRUCurrent})

    cefcFRUPowerStatusEntry.EntityData.YListKeys = []string {"EntPhysicalIndex"}

    return &(cefcFRUPowerStatusEntry.EntityData)
}

// CISCOENTITYFRUCONTROLMIB_CefcFRUPowerSupplyValueTable
type CISCOENTITYFRUCONTROLMIB_CefcFRUPowerSupplyValueTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // CISCOENTITYFRUCONTROLMIB_CefcFRUPowerSupplyValueTable_CefcFRUPowerSupplyValueEntry.
    CefcFRUPowerSupplyValueEntry []*CISCOENTITYFRUCONTROLMIB_CefcFRUPowerSupplyValueTable_CefcFRUPowerSupplyValueEntry
}

func (cefcFRUPowerSupplyValueTable *CISCOENTITYFRUCONTROLMIB_CefcFRUPowerSupplyValueTable) GetEntityData() *types.CommonEntityData {
    cefcFRUPowerSupplyValueTable.EntityData.YFilter = cefcFRUPowerSupplyValueTable.YFilter
    cefcFRUPowerSupplyValueTable.EntityData.YangName = "cefcFRUPowerSupplyValueTable"
    cefcFRUPowerSupplyValueTable.EntityData.BundleName = "cisco_ios_xr"
    cefcFRUPowerSupplyValueTable.EntityData.ParentYangName = "CISCO-ENTITY-FRU-CONTROL-MIB"
    cefcFRUPowerSupplyValueTable.EntityData.SegmentPath = "cefcFRUPowerSupplyValueTable"
    cefcFRUPowerSupplyValueTable.EntityData.AbsolutePath = "CISCO-ENTITY-FRU-CONTROL-MIB:CISCO-ENTITY-FRU-CONTROL-MIB/" + cefcFRUPowerSupplyValueTable.EntityData.SegmentPath
    cefcFRUPowerSupplyValueTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cefcFRUPowerSupplyValueTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cefcFRUPowerSupplyValueTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cefcFRUPowerSupplyValueTable.EntityData.Children = types.NewOrderedMap()
    cefcFRUPowerSupplyValueTable.EntityData.Children.Append("cefcFRUPowerSupplyValueEntry", types.YChild{"CefcFRUPowerSupplyValueEntry", nil})
    for i := range cefcFRUPowerSupplyValueTable.CefcFRUPowerSupplyValueEntry {
        cefcFRUPowerSupplyValueTable.EntityData.Children.Append(types.GetSegmentPath(cefcFRUPowerSupplyValueTable.CefcFRUPowerSupplyValueEntry[i]), types.YChild{"CefcFRUPowerSupplyValueEntry", cefcFRUPowerSupplyValueTable.CefcFRUPowerSupplyValueEntry[i]})
    }
    cefcFRUPowerSupplyValueTable.EntityData.Leafs = types.NewOrderedMap()

    cefcFRUPowerSupplyValueTable.EntityData.YListKeys = []string {}

    return &(cefcFRUPowerSupplyValueTable.EntityData)
}

// CISCOENTITYFRUCONTROLMIB_CefcFRUPowerSupplyValueTable_CefcFRUPowerSupplyValueEntry
type CISCOENTITYFRUCONTROLMIB_CefcFRUPowerSupplyValueTable_CefcFRUPowerSupplyValueEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is interface{} with range: 1..2147483647.
    EntPhysicalIndex interface{}

    // The type is interface{} with range: -1000000000..1000000000.
    CefcFRUTotalSystemCurrent interface{}

    // The type is interface{} with range: -1000000000..1000000000.
    CefcFRUDrawnSystemCurrent interface{}

    // The type is interface{} with range: -1000000000..1000000000.
    CefcFRUTotalInlineCurrent interface{}

    // The type is interface{} with range: -1000000000..1000000000.
    CefcFRUDrawnInlineCurrent interface{}
}

func (cefcFRUPowerSupplyValueEntry *CISCOENTITYFRUCONTROLMIB_CefcFRUPowerSupplyValueTable_CefcFRUPowerSupplyValueEntry) GetEntityData() *types.CommonEntityData {
    cefcFRUPowerSupplyValueEntry.EntityData.YFilter = cefcFRUPowerSupplyValueEntry.YFilter
    cefcFRUPowerSupplyValueEntry.EntityData.YangName = "cefcFRUPowerSupplyValueEntry"
    cefcFRUPowerSupplyValueEntry.EntityData.BundleName = "cisco_ios_xr"
    cefcFRUPowerSupplyValueEntry.EntityData.ParentYangName = "cefcFRUPowerSupplyValueTable"
    cefcFRUPowerSupplyValueEntry.EntityData.SegmentPath = "cefcFRUPowerSupplyValueEntry" + types.AddKeyToken(cefcFRUPowerSupplyValueEntry.EntPhysicalIndex, "entPhysicalIndex")
    cefcFRUPowerSupplyValueEntry.EntityData.AbsolutePath = "CISCO-ENTITY-FRU-CONTROL-MIB:CISCO-ENTITY-FRU-CONTROL-MIB/cefcFRUPowerSupplyValueTable/" + cefcFRUPowerSupplyValueEntry.EntityData.SegmentPath
    cefcFRUPowerSupplyValueEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cefcFRUPowerSupplyValueEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cefcFRUPowerSupplyValueEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cefcFRUPowerSupplyValueEntry.EntityData.Children = types.NewOrderedMap()
    cefcFRUPowerSupplyValueEntry.EntityData.Leafs = types.NewOrderedMap()
    cefcFRUPowerSupplyValueEntry.EntityData.Leafs.Append("entPhysicalIndex", types.YLeaf{"EntPhysicalIndex", cefcFRUPowerSupplyValueEntry.EntPhysicalIndex})
    cefcFRUPowerSupplyValueEntry.EntityData.Leafs.Append("cefcFRUTotalSystemCurrent", types.YLeaf{"CefcFRUTotalSystemCurrent", cefcFRUPowerSupplyValueEntry.CefcFRUTotalSystemCurrent})
    cefcFRUPowerSupplyValueEntry.EntityData.Leafs.Append("cefcFRUDrawnSystemCurrent", types.YLeaf{"CefcFRUDrawnSystemCurrent", cefcFRUPowerSupplyValueEntry.CefcFRUDrawnSystemCurrent})
    cefcFRUPowerSupplyValueEntry.EntityData.Leafs.Append("cefcFRUTotalInlineCurrent", types.YLeaf{"CefcFRUTotalInlineCurrent", cefcFRUPowerSupplyValueEntry.CefcFRUTotalInlineCurrent})
    cefcFRUPowerSupplyValueEntry.EntityData.Leafs.Append("cefcFRUDrawnInlineCurrent", types.YLeaf{"CefcFRUDrawnInlineCurrent", cefcFRUPowerSupplyValueEntry.CefcFRUDrawnInlineCurrent})

    cefcFRUPowerSupplyValueEntry.EntityData.YListKeys = []string {"EntPhysicalIndex"}

    return &(cefcFRUPowerSupplyValueEntry.EntityData)
}

// CISCOENTITYFRUCONTROLMIB_CefcModuleTable
type CISCOENTITYFRUCONTROLMIB_CefcModuleTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // CISCOENTITYFRUCONTROLMIB_CefcModuleTable_CefcModuleEntry.
    CefcModuleEntry []*CISCOENTITYFRUCONTROLMIB_CefcModuleTable_CefcModuleEntry
}

func (cefcModuleTable *CISCOENTITYFRUCONTROLMIB_CefcModuleTable) GetEntityData() *types.CommonEntityData {
    cefcModuleTable.EntityData.YFilter = cefcModuleTable.YFilter
    cefcModuleTable.EntityData.YangName = "cefcModuleTable"
    cefcModuleTable.EntityData.BundleName = "cisco_ios_xr"
    cefcModuleTable.EntityData.ParentYangName = "CISCO-ENTITY-FRU-CONTROL-MIB"
    cefcModuleTable.EntityData.SegmentPath = "cefcModuleTable"
    cefcModuleTable.EntityData.AbsolutePath = "CISCO-ENTITY-FRU-CONTROL-MIB:CISCO-ENTITY-FRU-CONTROL-MIB/" + cefcModuleTable.EntityData.SegmentPath
    cefcModuleTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cefcModuleTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cefcModuleTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cefcModuleTable.EntityData.Children = types.NewOrderedMap()
    cefcModuleTable.EntityData.Children.Append("cefcModuleEntry", types.YChild{"CefcModuleEntry", nil})
    for i := range cefcModuleTable.CefcModuleEntry {
        cefcModuleTable.EntityData.Children.Append(types.GetSegmentPath(cefcModuleTable.CefcModuleEntry[i]), types.YChild{"CefcModuleEntry", cefcModuleTable.CefcModuleEntry[i]})
    }
    cefcModuleTable.EntityData.Leafs = types.NewOrderedMap()

    cefcModuleTable.EntityData.YListKeys = []string {}

    return &(cefcModuleTable.EntityData)
}

// CISCOENTITYFRUCONTROLMIB_CefcModuleTable_CefcModuleEntry
type CISCOENTITYFRUCONTROLMIB_CefcModuleTable_CefcModuleEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is interface{} with range: 1..2147483647.
    EntPhysicalIndex interface{}

    // The type is ModuleAdminType.
    CefcModuleAdminStatus interface{}

    // The type is ModuleOperType.
    CefcModuleOperStatus interface{}

    // The type is ModuleResetReasonType.
    CefcModuleResetReason interface{}

    // The type is interface{} with range: 0..4294967295.
    CefcModuleStatusLastChangeTime interface{}

    // The type is interface{} with range: 0..4294967295.
    CefcModuleLastClearConfigTime interface{}

    // The type is string with length: 0..255.
    CefcModuleResetReasonDescription interface{}

    // The type is string with length: 0..255.
    CefcModuleStateChangeReasonDescr interface{}

    // The type is interface{} with range: 0..4294967295.
    CefcModuleUpTime interface{}
}

func (cefcModuleEntry *CISCOENTITYFRUCONTROLMIB_CefcModuleTable_CefcModuleEntry) GetEntityData() *types.CommonEntityData {
    cefcModuleEntry.EntityData.YFilter = cefcModuleEntry.YFilter
    cefcModuleEntry.EntityData.YangName = "cefcModuleEntry"
    cefcModuleEntry.EntityData.BundleName = "cisco_ios_xr"
    cefcModuleEntry.EntityData.ParentYangName = "cefcModuleTable"
    cefcModuleEntry.EntityData.SegmentPath = "cefcModuleEntry" + types.AddKeyToken(cefcModuleEntry.EntPhysicalIndex, "entPhysicalIndex")
    cefcModuleEntry.EntityData.AbsolutePath = "CISCO-ENTITY-FRU-CONTROL-MIB:CISCO-ENTITY-FRU-CONTROL-MIB/cefcModuleTable/" + cefcModuleEntry.EntityData.SegmentPath
    cefcModuleEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cefcModuleEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cefcModuleEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cefcModuleEntry.EntityData.Children = types.NewOrderedMap()
    cefcModuleEntry.EntityData.Leafs = types.NewOrderedMap()
    cefcModuleEntry.EntityData.Leafs.Append("entPhysicalIndex", types.YLeaf{"EntPhysicalIndex", cefcModuleEntry.EntPhysicalIndex})
    cefcModuleEntry.EntityData.Leafs.Append("cefcModuleAdminStatus", types.YLeaf{"CefcModuleAdminStatus", cefcModuleEntry.CefcModuleAdminStatus})
    cefcModuleEntry.EntityData.Leafs.Append("cefcModuleOperStatus", types.YLeaf{"CefcModuleOperStatus", cefcModuleEntry.CefcModuleOperStatus})
    cefcModuleEntry.EntityData.Leafs.Append("cefcModuleResetReason", types.YLeaf{"CefcModuleResetReason", cefcModuleEntry.CefcModuleResetReason})
    cefcModuleEntry.EntityData.Leafs.Append("cefcModuleStatusLastChangeTime", types.YLeaf{"CefcModuleStatusLastChangeTime", cefcModuleEntry.CefcModuleStatusLastChangeTime})
    cefcModuleEntry.EntityData.Leafs.Append("cefcModuleLastClearConfigTime", types.YLeaf{"CefcModuleLastClearConfigTime", cefcModuleEntry.CefcModuleLastClearConfigTime})
    cefcModuleEntry.EntityData.Leafs.Append("cefcModuleResetReasonDescription", types.YLeaf{"CefcModuleResetReasonDescription", cefcModuleEntry.CefcModuleResetReasonDescription})
    cefcModuleEntry.EntityData.Leafs.Append("cefcModuleStateChangeReasonDescr", types.YLeaf{"CefcModuleStateChangeReasonDescr", cefcModuleEntry.CefcModuleStateChangeReasonDescr})
    cefcModuleEntry.EntityData.Leafs.Append("cefcModuleUpTime", types.YLeaf{"CefcModuleUpTime", cefcModuleEntry.CefcModuleUpTime})

    cefcModuleEntry.EntityData.YListKeys = []string {"EntPhysicalIndex"}

    return &(cefcModuleEntry.EntityData)
}

// CISCOENTITYFRUCONTROLMIB_CefcIntelliModuleTable
type CISCOENTITYFRUCONTROLMIB_CefcIntelliModuleTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // CISCOENTITYFRUCONTROLMIB_CefcIntelliModuleTable_CefcIntelliModuleEntry.
    CefcIntelliModuleEntry []*CISCOENTITYFRUCONTROLMIB_CefcIntelliModuleTable_CefcIntelliModuleEntry
}

func (cefcIntelliModuleTable *CISCOENTITYFRUCONTROLMIB_CefcIntelliModuleTable) GetEntityData() *types.CommonEntityData {
    cefcIntelliModuleTable.EntityData.YFilter = cefcIntelliModuleTable.YFilter
    cefcIntelliModuleTable.EntityData.YangName = "cefcIntelliModuleTable"
    cefcIntelliModuleTable.EntityData.BundleName = "cisco_ios_xr"
    cefcIntelliModuleTable.EntityData.ParentYangName = "CISCO-ENTITY-FRU-CONTROL-MIB"
    cefcIntelliModuleTable.EntityData.SegmentPath = "cefcIntelliModuleTable"
    cefcIntelliModuleTable.EntityData.AbsolutePath = "CISCO-ENTITY-FRU-CONTROL-MIB:CISCO-ENTITY-FRU-CONTROL-MIB/" + cefcIntelliModuleTable.EntityData.SegmentPath
    cefcIntelliModuleTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cefcIntelliModuleTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cefcIntelliModuleTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cefcIntelliModuleTable.EntityData.Children = types.NewOrderedMap()
    cefcIntelliModuleTable.EntityData.Children.Append("cefcIntelliModuleEntry", types.YChild{"CefcIntelliModuleEntry", nil})
    for i := range cefcIntelliModuleTable.CefcIntelliModuleEntry {
        cefcIntelliModuleTable.EntityData.Children.Append(types.GetSegmentPath(cefcIntelliModuleTable.CefcIntelliModuleEntry[i]), types.YChild{"CefcIntelliModuleEntry", cefcIntelliModuleTable.CefcIntelliModuleEntry[i]})
    }
    cefcIntelliModuleTable.EntityData.Leafs = types.NewOrderedMap()

    cefcIntelliModuleTable.EntityData.YListKeys = []string {}

    return &(cefcIntelliModuleTable.EntityData)
}

// CISCOENTITYFRUCONTROLMIB_CefcIntelliModuleTable_CefcIntelliModuleEntry
type CISCOENTITYFRUCONTROLMIB_CefcIntelliModuleTable_CefcIntelliModuleEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is interface{} with range: 1..2147483647.
    EntPhysicalIndex interface{}

    // The type is InetAddressType.
    CefcIntelliModuleIPAddrType interface{}

    // The type is string with pattern:
    // b'(([0-9a-fA-F]){2}(:([0-9a-fA-F]){2})*)?'.
    CefcIntelliModuleIPAddr interface{}
}

func (cefcIntelliModuleEntry *CISCOENTITYFRUCONTROLMIB_CefcIntelliModuleTable_CefcIntelliModuleEntry) GetEntityData() *types.CommonEntityData {
    cefcIntelliModuleEntry.EntityData.YFilter = cefcIntelliModuleEntry.YFilter
    cefcIntelliModuleEntry.EntityData.YangName = "cefcIntelliModuleEntry"
    cefcIntelliModuleEntry.EntityData.BundleName = "cisco_ios_xr"
    cefcIntelliModuleEntry.EntityData.ParentYangName = "cefcIntelliModuleTable"
    cefcIntelliModuleEntry.EntityData.SegmentPath = "cefcIntelliModuleEntry" + types.AddKeyToken(cefcIntelliModuleEntry.EntPhysicalIndex, "entPhysicalIndex")
    cefcIntelliModuleEntry.EntityData.AbsolutePath = "CISCO-ENTITY-FRU-CONTROL-MIB:CISCO-ENTITY-FRU-CONTROL-MIB/cefcIntelliModuleTable/" + cefcIntelliModuleEntry.EntityData.SegmentPath
    cefcIntelliModuleEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cefcIntelliModuleEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cefcIntelliModuleEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cefcIntelliModuleEntry.EntityData.Children = types.NewOrderedMap()
    cefcIntelliModuleEntry.EntityData.Leafs = types.NewOrderedMap()
    cefcIntelliModuleEntry.EntityData.Leafs.Append("entPhysicalIndex", types.YLeaf{"EntPhysicalIndex", cefcIntelliModuleEntry.EntPhysicalIndex})
    cefcIntelliModuleEntry.EntityData.Leafs.Append("cefcIntelliModuleIPAddrType", types.YLeaf{"CefcIntelliModuleIPAddrType", cefcIntelliModuleEntry.CefcIntelliModuleIPAddrType})
    cefcIntelliModuleEntry.EntityData.Leafs.Append("cefcIntelliModuleIPAddr", types.YLeaf{"CefcIntelliModuleIPAddr", cefcIntelliModuleEntry.CefcIntelliModuleIPAddr})

    cefcIntelliModuleEntry.EntityData.YListKeys = []string {"EntPhysicalIndex"}

    return &(cefcIntelliModuleEntry.EntityData)
}

// CISCOENTITYFRUCONTROLMIB_CefcFanTrayStatusTable
type CISCOENTITYFRUCONTROLMIB_CefcFanTrayStatusTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // CISCOENTITYFRUCONTROLMIB_CefcFanTrayStatusTable_CefcFanTrayStatusEntry.
    CefcFanTrayStatusEntry []*CISCOENTITYFRUCONTROLMIB_CefcFanTrayStatusTable_CefcFanTrayStatusEntry
}

func (cefcFanTrayStatusTable *CISCOENTITYFRUCONTROLMIB_CefcFanTrayStatusTable) GetEntityData() *types.CommonEntityData {
    cefcFanTrayStatusTable.EntityData.YFilter = cefcFanTrayStatusTable.YFilter
    cefcFanTrayStatusTable.EntityData.YangName = "cefcFanTrayStatusTable"
    cefcFanTrayStatusTable.EntityData.BundleName = "cisco_ios_xr"
    cefcFanTrayStatusTable.EntityData.ParentYangName = "CISCO-ENTITY-FRU-CONTROL-MIB"
    cefcFanTrayStatusTable.EntityData.SegmentPath = "cefcFanTrayStatusTable"
    cefcFanTrayStatusTable.EntityData.AbsolutePath = "CISCO-ENTITY-FRU-CONTROL-MIB:CISCO-ENTITY-FRU-CONTROL-MIB/" + cefcFanTrayStatusTable.EntityData.SegmentPath
    cefcFanTrayStatusTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cefcFanTrayStatusTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cefcFanTrayStatusTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cefcFanTrayStatusTable.EntityData.Children = types.NewOrderedMap()
    cefcFanTrayStatusTable.EntityData.Children.Append("cefcFanTrayStatusEntry", types.YChild{"CefcFanTrayStatusEntry", nil})
    for i := range cefcFanTrayStatusTable.CefcFanTrayStatusEntry {
        cefcFanTrayStatusTable.EntityData.Children.Append(types.GetSegmentPath(cefcFanTrayStatusTable.CefcFanTrayStatusEntry[i]), types.YChild{"CefcFanTrayStatusEntry", cefcFanTrayStatusTable.CefcFanTrayStatusEntry[i]})
    }
    cefcFanTrayStatusTable.EntityData.Leafs = types.NewOrderedMap()

    cefcFanTrayStatusTable.EntityData.YListKeys = []string {}

    return &(cefcFanTrayStatusTable.EntityData)
}

// CISCOENTITYFRUCONTROLMIB_CefcFanTrayStatusTable_CefcFanTrayStatusEntry
type CISCOENTITYFRUCONTROLMIB_CefcFanTrayStatusTable_CefcFanTrayStatusEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is interface{} with range: 1..2147483647.
    EntPhysicalIndex interface{}

    // The type is CefcFanTrayOperStatusType.
    CefcFanTrayOperStatus interface{}
}

func (cefcFanTrayStatusEntry *CISCOENTITYFRUCONTROLMIB_CefcFanTrayStatusTable_CefcFanTrayStatusEntry) GetEntityData() *types.CommonEntityData {
    cefcFanTrayStatusEntry.EntityData.YFilter = cefcFanTrayStatusEntry.YFilter
    cefcFanTrayStatusEntry.EntityData.YangName = "cefcFanTrayStatusEntry"
    cefcFanTrayStatusEntry.EntityData.BundleName = "cisco_ios_xr"
    cefcFanTrayStatusEntry.EntityData.ParentYangName = "cefcFanTrayStatusTable"
    cefcFanTrayStatusEntry.EntityData.SegmentPath = "cefcFanTrayStatusEntry" + types.AddKeyToken(cefcFanTrayStatusEntry.EntPhysicalIndex, "entPhysicalIndex")
    cefcFanTrayStatusEntry.EntityData.AbsolutePath = "CISCO-ENTITY-FRU-CONTROL-MIB:CISCO-ENTITY-FRU-CONTROL-MIB/cefcFanTrayStatusTable/" + cefcFanTrayStatusEntry.EntityData.SegmentPath
    cefcFanTrayStatusEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cefcFanTrayStatusEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cefcFanTrayStatusEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cefcFanTrayStatusEntry.EntityData.Children = types.NewOrderedMap()
    cefcFanTrayStatusEntry.EntityData.Leafs = types.NewOrderedMap()
    cefcFanTrayStatusEntry.EntityData.Leafs.Append("entPhysicalIndex", types.YLeaf{"EntPhysicalIndex", cefcFanTrayStatusEntry.EntPhysicalIndex})
    cefcFanTrayStatusEntry.EntityData.Leafs.Append("cefcFanTrayOperStatus", types.YLeaf{"CefcFanTrayOperStatus", cefcFanTrayStatusEntry.CefcFanTrayOperStatus})

    cefcFanTrayStatusEntry.EntityData.YListKeys = []string {"EntPhysicalIndex"}

    return &(cefcFanTrayStatusEntry.EntityData)
}

// CISCOENTITYFRUCONTROLMIB_CefcPhysicalTable
type CISCOENTITYFRUCONTROLMIB_CefcPhysicalTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // CISCOENTITYFRUCONTROLMIB_CefcPhysicalTable_CefcPhysicalEntry.
    CefcPhysicalEntry []*CISCOENTITYFRUCONTROLMIB_CefcPhysicalTable_CefcPhysicalEntry
}

func (cefcPhysicalTable *CISCOENTITYFRUCONTROLMIB_CefcPhysicalTable) GetEntityData() *types.CommonEntityData {
    cefcPhysicalTable.EntityData.YFilter = cefcPhysicalTable.YFilter
    cefcPhysicalTable.EntityData.YangName = "cefcPhysicalTable"
    cefcPhysicalTable.EntityData.BundleName = "cisco_ios_xr"
    cefcPhysicalTable.EntityData.ParentYangName = "CISCO-ENTITY-FRU-CONTROL-MIB"
    cefcPhysicalTable.EntityData.SegmentPath = "cefcPhysicalTable"
    cefcPhysicalTable.EntityData.AbsolutePath = "CISCO-ENTITY-FRU-CONTROL-MIB:CISCO-ENTITY-FRU-CONTROL-MIB/" + cefcPhysicalTable.EntityData.SegmentPath
    cefcPhysicalTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cefcPhysicalTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cefcPhysicalTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cefcPhysicalTable.EntityData.Children = types.NewOrderedMap()
    cefcPhysicalTable.EntityData.Children.Append("cefcPhysicalEntry", types.YChild{"CefcPhysicalEntry", nil})
    for i := range cefcPhysicalTable.CefcPhysicalEntry {
        cefcPhysicalTable.EntityData.Children.Append(types.GetSegmentPath(cefcPhysicalTable.CefcPhysicalEntry[i]), types.YChild{"CefcPhysicalEntry", cefcPhysicalTable.CefcPhysicalEntry[i]})
    }
    cefcPhysicalTable.EntityData.Leafs = types.NewOrderedMap()

    cefcPhysicalTable.EntityData.YListKeys = []string {}

    return &(cefcPhysicalTable.EntityData)
}

// CISCOENTITYFRUCONTROLMIB_CefcPhysicalTable_CefcPhysicalEntry
type CISCOENTITYFRUCONTROLMIB_CefcPhysicalTable_CefcPhysicalEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is interface{} with range: 1..2147483647.
    EntPhysicalIndex interface{}

    // The type is CefcPhysicalStatusType.
    CefcPhysicalStatus interface{}
}

func (cefcPhysicalEntry *CISCOENTITYFRUCONTROLMIB_CefcPhysicalTable_CefcPhysicalEntry) GetEntityData() *types.CommonEntityData {
    cefcPhysicalEntry.EntityData.YFilter = cefcPhysicalEntry.YFilter
    cefcPhysicalEntry.EntityData.YangName = "cefcPhysicalEntry"
    cefcPhysicalEntry.EntityData.BundleName = "cisco_ios_xr"
    cefcPhysicalEntry.EntityData.ParentYangName = "cefcPhysicalTable"
    cefcPhysicalEntry.EntityData.SegmentPath = "cefcPhysicalEntry" + types.AddKeyToken(cefcPhysicalEntry.EntPhysicalIndex, "entPhysicalIndex")
    cefcPhysicalEntry.EntityData.AbsolutePath = "CISCO-ENTITY-FRU-CONTROL-MIB:CISCO-ENTITY-FRU-CONTROL-MIB/cefcPhysicalTable/" + cefcPhysicalEntry.EntityData.SegmentPath
    cefcPhysicalEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cefcPhysicalEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cefcPhysicalEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cefcPhysicalEntry.EntityData.Children = types.NewOrderedMap()
    cefcPhysicalEntry.EntityData.Leafs = types.NewOrderedMap()
    cefcPhysicalEntry.EntityData.Leafs.Append("entPhysicalIndex", types.YLeaf{"EntPhysicalIndex", cefcPhysicalEntry.EntPhysicalIndex})
    cefcPhysicalEntry.EntityData.Leafs.Append("cefcPhysicalStatus", types.YLeaf{"CefcPhysicalStatus", cefcPhysicalEntry.CefcPhysicalStatus})

    cefcPhysicalEntry.EntityData.YListKeys = []string {"EntPhysicalIndex"}

    return &(cefcPhysicalEntry.EntityData)
}

